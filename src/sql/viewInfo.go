package sql

import (
	"fmt"
	"strconv"
	"strings"
)

type ViewInfo struct {
	Listen   *MyViewSQLListener
	ViewType int
	MapType  map[string]int //I actually did not need this to build the update
	TopKType TopKType

	FieldsNeeded map[string]DataType //Collects from SELECT and WHERE all data that is needed from the object and what should be their types
}

const (
	TOPK     = 1
	TOPSUM   = 2
	COUNTER  = 3
	AVERAGE  = 4
	REGISTER = 5
	MAP      = 6
	MAX_CRDT = 7
	MIN_CRDT = 8
)

var (
	CrdtTypeToString = []string{"UNDEFINED", "TOPK", "TOPSUM", "COUNTER", "AVERAGE", "REGISTER", "MAP", "MAX_CRDT", "MIN_CRDT"}
)

type TopKType struct {
	Value     string
	Id        string
	ExtraData []string
	//Id    []string
}

func (v *ViewInfo) PrepareInfo() {
	viewType := v.GetViewType()
	if viewType == MAP {
		v.makeViewMapFullType()
	} else if viewType == TOPK || viewType == TOPSUM {
		v.makeViewTopFullType()
	}
	v.CalculateTypes()
	v.CalculateNeededFields()
	v.ReplaceParameters()
	fmt.Println("Conditions:")
	fmt.Print("[")
	for _, cond := range v.Listen.Conditions {
		fmt.Print(cond.ToString() + "; ")
	}
	fmt.Println("]")
}

func (v *ViewInfo) GetViewType() int {
	if v.ViewType < TOPK {
		v.ViewType = v.calculateViewType()
	}
	return v.ViewType
}

func (v *ViewInfo) calculateViewType() int {
	if v.Listen.TopSize > 0 {
		//TOPK or TOPSUM
		if v.HasSumOrCount() {
			return TOPSUM
		}
		return TOPK
	}
	if v.Listen.NEntriesInSelect > 1 {
		return MAP
	}
	if len(v.Listen.SelectAsCount) == 1 {
		return COUNTER
	}

	if len(v.Listen.SelectAsAggregation) == 1 {
		for _, value := range v.Listen.SelectAsAggregation {
			return v.AggrToCrdtType(value.AggrFunc)
		}
	}
	//At this point we know that there's only one entry in select and that said entry is in either asClause or nameable.
	return REGISTER
}

func (v *ViewInfo) AggrToCrdtType(aggrFunc AggrType) int {
	switch aggrFunc {
	case SUM:
		return COUNTER
	case AVG:
		return AVERAGE
	case MAX:
		return MAX_CRDT
	case MIN:
		return MIN_CRDT
	}
	return -1
}

func (v *ViewInfo) GetViewMapFullType() map[string]int {
	if v.MapType == nil {
		v.makeViewMapFullType()
	}
	return v.MapType
}

//Pre: GetViewType() == MAP
//This method calculates the map's full type and stores its information.
func (v *ViewInfo) makeViewMapFullType() {
	v.MapType = make(map[string]int)

	for id, _ := range v.Listen.SelectNameable {
		//If it is a Nameable a register is always OK.
		v.MapType[id] = REGISTER
	}
	for id, _ := range v.Listen.SelectAsClause {
		//mathType := asClause.Math.GetCalculable().GetType()
		//Still OK to be a register: at most we have something like a + 5. Not SUM and such...
		v.MapType[id] = REGISTER
	}
	for id, asAggregation := range v.Listen.SelectAsAggregation {
		v.MapType[id] = v.AggrToCrdtType(asAggregation.AggrFunc)
	}
	for id, _ := range v.Listen.SelectAsCount {
		v.MapType[id] = COUNTER
	}
}

func (v *ViewInfo) GetViewTopFullType() TopKType {
	if v.TopKType.ExtraData == nil {
		v.makeViewTopFullType()
	}
	return v.TopKType
}

//Pre: GetViewType() == TOPK || TOPSUM
func (v *ViewInfo) makeViewTopFullType() {
	//Value: OrderBy (first, due to current limitation)
	//ID: v.Listen.PrimaryKey
	//TODO: Limitation: our TopCRDTs only support ordering by ints and by a single order factor. And Ids must actually be ints.
	v.TopKType = TopKType{Value: v.Listen.Sorting[0].Nameable.GetIdentifier(), Id: v.Listen.PrimaryKey.GetIdentifier(),
		ExtraData: make([]string, v.Listen.NEntriesInSelect-2)} //-2: taking out the value and id
	//It's okay for this to not be efficient as this method will only be called once when preparing the view.
	i := 0
	for _, entry := range v.Listen.SelectByAppearenceOrder {
		name := entry.GetIdentifier()
		if name != v.TopKType.Value && name != v.TopKType.Id {
			v.TopKType.ExtraData[i] = name
			i++
		}
	}
}

func (v *ViewInfo) HasSumOrCount() bool {
	if len(v.Listen.SelectAsCount) > 0 {
		return true
	}
	for _, asAggr := range v.Listen.SelectAsAggregation {
		if asAggr.AggrFunc == SUM {
			return true
		}
	}
	return false
}

func (v *ViewInfo) CalculateTypes() {
	//SelectNameable can't conclude anything about their own typing. Nor can SelectAsCount.
	for _, asClause := range v.Listen.SelectAsClause {
		asClause.Math.GetCalculable().GetType()
	}
	//For all nameables in aggregations we can force their type to be NUMBER
	for _, asAggregation := range v.Listen.SelectAsAggregation {
		//First, call the normal algorithm
		calcType := asAggregation.Math.GetCalculable().GetType()
		if calcType == VARIABLE_TYPE { //If it isn't already defined, set to NUMBER
			asAggregation.Math.GetCalculable().SetType(NUMBER_TYPE)
		}
	}
	//This may need multiple iterations, until no change is observed
	//Note: While multiple iterations are necessary, this could be optimized.
	variableLeft, previousLeft := 1, 0
	for variableLeft > 0 {
		variableLeft = 0
		fmt.Printf("Conditions: %+v\n", v.Listen.Conditions)
		for _, cond := range v.Listen.Conditions {
			calcType := cond.Math.GetCalculable().GetType()
			if cond.Op != FUNCTION {
				nameableType := cond.Nameable.GetType()
				//fmt.Printf("[CALCULATETYPES]Condition: %s, %s, %s. Types before: %s %s\n",
				//cond.Nameable.GetIdentifier(), cond.Op.ToString(), cond.Math.ToString(), nameableType.ToString(), calcType.ToString())
				if nameableType == VARIABLE_TYPE {
					cond.Nameable.SetType(calcType)
					variableLeft++
				} else if nameableType == NUMBER_TYPE && (calcType == INT_TYPE || calcType == FLOAT_TYPE) {
					cond.Nameable.SetType(calcType)
					variableLeft++
				} else if calcType == VARIABLE_TYPE && nameableType != VARIABLE_TYPE {
					cond.Math.GetCalculable().SetType(nameableType)
					variableLeft++
				} else if calcType == NUMBER_TYPE && (nameableType == INT_TYPE || nameableType == FLOAT_TYPE) {
					cond.Math.GetCalculable().SetType(nameableType)
					variableLeft++
				}
				//fmt.Printf("[CALCULATETYPES]Condition: %s, %s, %s. Types after: %s %s\n",
				//cond.Nameable.GetIdentifier(), cond.Op.ToString(), cond.Math.ToString(),
				//cond.Nameable.GetType().ToString(), cond.Math.GetCalculable().GetType().ToString())
			} else {
				cond.Math.GetCalculable().SetType(STRING_TYPE)
				cond.Nameable.SetType(STRING_TYPE)
			}
		}
		if variableLeft == previousLeft { //No progress, some nameables we cannot conclude their type.
			break
		} else {
			previousLeft = variableLeft
		}
	}
}

//Pre: CalculateTypes()
func (v *ViewInfo) CalculateNeededFields() {
	v.FieldsNeeded = make(map[string]DataType)

	for _, nameable := range v.Listen.SelectNameable {
		nameable.RegisterFields(v.FieldsNeeded)
	}
	for _, asClause := range v.Listen.SelectAsClause {
		asClause.Math.GetCalculable().RegisterFields(v.FieldsNeeded)
	}
	for _, asAggr := range v.Listen.SelectAsAggregation {
		asAggr.Math.GetCalculable().RegisterFields(v.FieldsNeeded)
	}
	for _, asCount := range v.Listen.SelectAsCount {
		if !asCount.IsCountAll() {
			asCount.Nameable.RegisterFields(v.FieldsNeeded)
		}
	}
	//This is before conditions have been modified
	for _, cond := range v.Listen.Conditions {
		cond.Nameable.RegisterFields(v.FieldsNeeded)
		cond.Math.GetCalculable().RegisterFields(v.FieldsNeeded)
	}
	for _, key := range v.Listen.Keys {
		key.RegisterFields(v.FieldsNeeded)
	}
	//Order by not needed: it must appear in select

	//I think I need to go through each clause separately,
	//as I do not want Parameters or names generated by AsClause and similar here.
	//I believe I will need to find some way to deal with those.
	/*
		for _, nameable := range v.Listen.AllNameables {
			nameable.RegisterFields(v.FieldsNeeded)
		}
	*/
}

//Tries to find algorithms which are compared to some value by equality
//Any such parameter is replaced by the variable that it is compared to
//Afterwards, for any remaining parameters, it tries to calculate the
//possible range of values.
//TODO: For now I am assuming very simple Parameters: aka no math associated
//to the parameters. I can improve on this later if needed.
//The grammar does not impose such restrictions as of now.

/*
	For equality, possible situations:
	Parameter is in nameable; this should be easy (replace parameter with math)
	Parameter is in math:
		Either math is only parameter; (e.g., shipdate = DATE)
		Math is simple: Parameter is either right away on the right side or on the left side of the first level.
			In this case, it is doable to modify the Math.
	For inequality: same situations as above.
	I should probably rewrite the Conditions to reflect this.

	Pre: all parameters must be defined in Keys
*/

func (v *ViewInfo) ReplaceParameters() {
	//Step 1: write all "=" conditions with a parameter so that the parameter is on the left side (Condition.Nameable)
	//Remove all said conditions from v.Listen.Conditions; save them in replaceEqualParams.
	replaceEqualParams := v.replaceEqualParams()
	//Steps 2-4: inequal parameters
	v.rewriteInequalConds()
	//replaceInequalParams := v.replaceInequalParams(replaceEqualParams)
	//Replace parameters in Select. Note that select only has equality parameters.
	v.replaceParamsInSelect(replaceEqualParams)
	//Where does not need to be updated as of now since it is unused after this method is called.
	//Sorting also does not need to be updated: it cannot have parameters.
	//We do need to update Keys though.
	v.replaceParamsInKeys(replaceEqualParams)
	//v.replaceParamsInKeys(replaceEqualParams, replaceInequalParams)
}

//NOTE: Future optimization: I think the replacement of equal and inequal params could occour in parallel.
//However this optimization is not crucial as this is only done once when the view specification is sent to the server.
//(Aka, this code is not run for each update)
func (v *ViewInfo) replaceEqualParams() (replaceEqualParams map[string]EqualParameter) {
	replaceEqualParams = make(map[string]EqualParameter)
	newConds := make([]Condition, 0, len(v.Listen.Conditions)) //All non-equal conditions and equal conditions not evolving parameters
	for i, cond := range v.Listen.Conditions {
		fmt.Printf("[VIEWINFO][ReplaceEq]Condition: %s %s %s\n", cond.Nameable.ToString(), cond.Op.ToString(), cond.Math.ToString())
		var sb1 strings.Builder
		var sb2 strings.Builder
		cond.Nameable.ToPrintString(&sb1)
		cond.Math.GetCalculable().ToPrintString(&sb2)
		//fmt.Printf("[VIEWINFO][ReplaceEq]Condition: %s %s %s\n", sb1.String(), cond.Op.ToString(), sb2.String())
		if param, ok := cond.Nameable.(Parameter); ok {
			if cond.Op == EQUAL {
				//fmt.Println("[VIEWINFO][ReplaceEq]Condition is EQUAL, has Parameter in Nameable")
				//eqP := EqualParameter{Parameter: param, CalcFrom: cond.Math.GetCalculable()}
				//replaceEqualParams[param.GetIdentifier()] = eqP
				//v.Listen.Conditions[i].Nameable = &eqP
				replaceEqualParams[param.GetIdentifier()] = EqualParameter{Parameter: param, CalcFrom: cond.Math.GetCalculable()}
			} else {
				//fmt.Println("[VIEWINFO][ReplaceEq]Condition isn't EQUAL, adding to newConds")
				newConds = append(newConds, cond)
			}
		} else if cond.Op == EQUAL {
			//Search for parameter in Math
			mathCalc := cond.Math.GetCalculable()
			if param, ok := mathCalc.(Parameter); ok { //Same logic as above.
				///eqP := EqualParameter{Parameter: param, CalcFrom: cond.Nameable}
				//replaceEqualParams[param.GetIdentifier()] = eqP
				//v.Listen.Conditions[i].Math.Stack.ReplaceTop(&eqP)
				replaceEqualParams[param.GetIdentifier()] = EqualParameter{Parameter: param, CalcFrom: cond.Nameable}
				fmt.Printf("[VIEWINFO][ReplaceEq]Condition is EQUAL, Math is Parameter. %v = %v\n", param.ToString(), cond.Nameable.ToString())
			} else if mathCalc.HasParameter() { //If it is a +, -, *, /, MINUS. Need to find in which side and then rewrite.
				//fmt.Println("[VIEWINFO][ReplaceEq]Condition is EQUAL, it has a parameter somewhere in Math")
				nameable := cond.Nameable
				var paramToMove Parameter
				var ok bool
				var newCalc Calculable
				//Idea: the nameable goes to the .Left; the old Math changes to the inverse (e.g: add -> sub.) and is stored on Right.
				switch typedC := mathCalc.(type) {
				case *Add:
					if paramToMove, ok = typedC.Left.(Parameter); ok { //Change right
						newCalc = &Sub{Left: nameable, Right: typedC.Right, Type: typedC.Type}
					} else {
						paramToMove, ok = typedC.Right.(Parameter)
						newCalc = &Sub{Left: nameable, Right: typedC.Left, Type: typedC.Type}
					}
				case *Sub:
					if paramToMove, ok = typedC.Left.(Parameter); ok {
						newCalc = &Add{Left: nameable, Right: typedC.Right, Type: typedC.Type}
					} else {
						paramToMove, ok = typedC.Right.(Parameter)
						newCalc = &Add{Left: nameable, Right: typedC.Left, Type: typedC.Type}
					}
				case *Mult:
					if paramToMove, ok = typedC.Left.(Parameter); ok { //Change right
						newCalc = &Div{Left: nameable, Right: typedC.Right, Type: typedC.Type}
					} else {
						paramToMove, ok = typedC.Right.(Parameter)
						newCalc = &Div{Left: nameable, Right: typedC.Left, Type: typedC.Type}
					}
				case *Div:
					if paramToMove, ok = typedC.Left.(Parameter); ok { //Change right
						newCalc = &Mult{Left: nameable, Right: typedC.Right, Type: typedC.Type}
					} else {
						paramToMove, ok = typedC.Right.(Parameter)
						newCalc = &Mult{Left: nameable, Right: typedC.Left, Type: typedC.Type}
					}
				case *Minus:
					//In theory could re-use
					paramToMove = typedC.ToCalculate.(Parameter)
					newCalc = &Minus{ToCalculate: nameable, Type: typedC.Type}
				}
				//eq := EqualParameter{Parameter: param, CalcFrom: newCalc}
				replaceEqualParams[paramToMove.GetIdentifier()] = EqualParameter{Parameter: param, CalcFrom: newCalc}
				v.Listen.Conditions[i].Math.Stack.ReplaceTop(newCalc)
				v.Listen.Conditions[i].Nameable = paramToMove
				//v.Listen.Conditions[i].Nameable = &eq
			} else {
				//fmt.Println("[VIEWINFO][ReplaceEq]Condition is EQUAL, it does not have a parameter, thus adding to newConds")
				newConds = append(newConds, cond)
			}
		} else {
			//fmt.Println("[VIEWINFO][ReplaceEq]Condition is not EQUAL, thus adding it to newConds")
			newConds = append(newConds, cond)
		}
	}
	v.Listen.Conditions = newConds
	v.Listen.EqParams = make([]EqualParameter, len(replaceEqualParams))
	i := 0
	for _, eqP := range replaceEqualParams {
		v.Listen.EqParams[i] = eqP
		i++
	}
	return
}

/*
func (v *ViewInfo) replaceInequalParams(replaceEqualParams map[string]EqualParameter) (replaceInequalParams map[string]*InequalParameter) {
	//Find which parameters are remaining (i.e., not in replaceEqualParams)
	//Any inequality parameter is for sure in group by (Keys)
	//fmt.Println("[VIEWINFO][ReplaceIneqParams]At start, conds:", v.Listen.Conditions)
	remaining := v.findRemainingParams(replaceEqualParams)
	//All remaining parameters must have limited bounds defined:
	//Thus, we need to find a condition limiting by > (or >=) and another by < (or <=)
	//Where one of the sides is a constant of some sort.
	//Step 2: extract, for all inequality parameters, their comparisons to constants.
	//Each inequality parameter must be compared to two constants, to define its bounds
	//fmt.Println("[VIEWINFO][ReplaceIneqParams]Before findBoundsInequalParams, conds:", v.Listen.Conditions)
	lowerBound, higherBound := v.findBoundsInequalParams(remaining)
	//Step 3: Make Inequal Params
	//fmt.Println("[VIEWINFO][ReplaceIneqParams]Before buildInequalParams, conds:", v.Listen.Conditions)
	replaceInequalParams = v.buildInequalParams(lowerBound, higherBound)
	//Step 4: Replace conditions with Inequal Params; build map of InequalParameters -> []Conditions
	//fmt.Println("[VIEWINFO][ReplaceIneqParams]Before replaceInequalParams, conds:", v.Listen.Conditions)
	v.replaceConditionsInequalParams(replaceInequalParams)
	return
}

func (v *ViewInfo) findRemainingParams(replaceEqualParams map[string]EqualParameter) (remaining map[string]struct{}) {
	remaining = make(map[string]struct{})
	for _, entry := range v.Listen.Keys {
		if _, has := replaceEqualParams[entry.GetIdentifier()]; !has {
			if param, ok := entry.(Parameter); ok {
				remaining[param.GetIdentifier()] = struct{}{}
			}
		}
	}
	return
}*/

//InequalParams: Parameters compared with inequality to some value (e.g.: shiporder > [DATE])
//Updates the expression of non-equality conditions to be of the form <<PARAMETER op Math>>
func (v *ViewInfo) rewriteInequalConds() {
	var param Parameter
	var value Calculable
	ok := false
	v.Listen.IneqParamConds = make(map[Parameter][]Condition)
	for i, cond := range v.Listen.Conditions {
		if cond.Op == EQUAL || cond.Op == NOT_EQ || cond.Op == FUNCTION {
			continue
		}
		if param, ok = cond.Nameable.(Parameter); ok {
			v.Listen.IneqParamConds[param] = append(v.Listen.IneqParamConds[param], cond)
			v.Listen.HasInequalityGrouping = true
			continue //Already OK
		}
		if param, ok = cond.Math.GetCalculable().(Parameter); ok {
			value = cond.Nameable
			//Need to invert the top
			if cond.Op == HIGHER || cond.Op == HIGHER_EQ {
				cond.Op += 2 //LOWER || LOWER_EQ
			} else {
				cond.Op -= 2
			}
			cond.Nameable = param
			cond.Math.Stack.ReplaceTop(value)
			v.Listen.Conditions[i] = cond
			v.Listen.IneqParamConds[param] = append(v.Listen.IneqParamConds[param], cond)
			v.Listen.HasInequalityGrouping = true
		}
		//Else: Neither is a parameter so do nothing.
	}
}

/*
//Updates the existing, non-equality conditions to be of the form <<PARAMETER op Math>>
func (v *ViewInfo) rewriteInequalConds(remaining map[string]struct{}) {
	if len(remaining) > 0 {
		for i, cond := range v.Listen.Conditions {
			var param Parameter
			var value Calculable
			has, ok := false, false //For readability. Could be just one variable
			if cond.Op == EQUAL || cond.Op == NOT_EQ || cond.Op == FUNCTION {
				continue
			}
			if _, has = remaining[cond.Nameable.GetIdentifier()]; !has {
				param, ok = cond.Math.GetCalculable().(Parameter)
				if !ok { //Neither is a parameter so, do not change anything
					continue
				}
				value = cond.Nameable
				//Need to invert the op
				if cond.Op == HIGHER || cond.Op == HIGHER_EQ {
					cond.Op += 2 //LOWER || LOWER_EQ
				} else {
					cond.Op -= 2 //HIGHER || HIGHER_EQ
				}
				cond.Nameable = param
				cond.Math.Stack.ReplaceTop(value)
				v.Listen.Conditions[i] = cond
			}
		}
	}
}
*/

/*
//Also updates the existing conditions to be of the form <<PARAMETER op Math>>
func (v *ViewInfo) findBoundsInequalParams(remaining map[string]struct{}) (lowerBound, higherBound map[string]Condition) {
	lowerBound = make(map[string]Condition)
	higherBound = make(map[string]Condition)
	newConds := make([]Condition, 0, len(v.Listen.Conditions)) //We will remove all conditions that define the bounds of Parameters
	if len(remaining) > 0 {
		for i, cond := range v.Listen.Conditions {
			var param Parameter
			var value Calculable
			has, ok := false, false //For readability. Could be just one variable
			if cond.Op == EQUAL || cond.Op == NOT_EQ || cond.Op == FUNCTION {
				newConds = append(newConds, cond)
				continue
			}
			if _, has = remaining[cond.Nameable.GetIdentifier()]; !has {
				param, ok = cond.Math.GetCalculable().(Parameter)
				if !ok { //Neither is a parameter so, do not change anything
					newConds = append(newConds, cond)
					continue
				}
				value = cond.Nameable
				//Need to invert the op
				if cond.Op == HIGHER || cond.Op == HIGHER_EQ {
					cond.Op += 2 //LOWER || LOWER_EQ
				} else {
					cond.Op -= 2 //HIGHER || HIGHER_EQ
				}
				cond.Nameable = param
				cond.Math.Stack.ReplaceTop(value)
				v.Listen.Conditions[i] = cond
			} else {
				param, value = cond.Nameable.(Parameter), cond.Math.GetCalculable()
			}
			if !value.IsConstant() {
				newConds = append(newConds, cond)
				continue
			}
			//Given at this point, we know we have something of the kind Parameter op Constant
			//With op being >, >=, <, <=.
			//Lets just store it and process it afterwards
			if cond.Op == HIGHER || cond.Op == HIGHER_EQ {
				lowerBound[param.Name] = cond
			} else {
				higherBound[param.Name] = cond
			}
		}
		v.Listen.Conditions = newConds
	}
	if len(lowerBound) > 0 { //There's at least one inequality parameter
		v.Listen.HasInequalityGrouping = true
	}
	return
}

func (v *ViewInfo) buildInequalParams(lowerBound, higherBound map[string]Condition) (replaceInequalParams map[string]*InequalParameter) {
	replaceInequalParams = make(map[string]*InequalParameter)
	//Step 3: now that the lower and higher bounds of each inequality parameter are known, we build an InequalParameter
	for key, lower := range lowerBound {
		higher := higherBound[key]
		switch lower.Math.GetCalculable().GetType() {
		case INT_TYPE:
			minValue, maxValue := lower.Math.GetCalculable().(*IntConstant).Value, higher.Math.GetCalculable().(*IntConstant).Value
			if lower.Op == HIGHER {
				minValue++
			}
			if higher.Op == LOWER {
				maxValue--
			}
			replaceInequalParams[key] = &InequalParameter{Parameter: lower.Nameable.(*Parameter), MinValue: minValue, MaxValue: maxValue}
		case FLOAT_TYPE:
			minC, maxC := lower.Math.GetCalculable().(*FloatConstant), higher.Math.GetCalculable().(*FloatConstant)
			minValue, maxValue := minC.Value, maxC.Value
			if lower.Op == HIGHER {
				minValue += math.Pow(0.1, float64(minC.Precision))
			}
			if higher.Op == LOWER {
				maxValue -= math.Pow(0.1, float64(minC.Precision))
			}
			replaceInequalParams[key] = &InequalParameter{Parameter: lower.Nameable.(*Parameter), MinValue: minValue, MaxValue: maxValue, FloatPrecision: minC.Precision}
		case DATE_TYPE:
			min, max := lower.Math.GetCalculable().(*DateConstant).Date, higher.Math.GetCalculable().(*DateConstant).Date
			if lower.Op == HIGHER {
				tmp := time.Date(min.Year, time.Month(min.Month), min.Day+1, 0, 0, 0, 0, time.UTC)
				min = Date{Year: tmp.Year(), Month: int(tmp.Month()), Day: tmp.Day()}
			}
			if higher.Op == LOWER {
				tmp := time.Date(max.Year, time.Month(max.Month), max.Day+1, 0, 0, 0, 0, time.UTC)
				max = Date{Year: tmp.Year(), Month: int(tmp.Month()), Day: tmp.Day()}
			}
			replaceInequalParams[key] = &InequalParameter{Parameter: lower.Nameable.(*Parameter), MinValue: min, MaxValue: max}
		}
	}
	return
}
*/

/*
func (v *ViewInfo) replaceConditionsInequalParams(replaceInequalParams map[string]*InequalParameter) {
	//Step 4: Replace all Parameters in Conditions with InequalParameters
	//Note: only parameters with inequalities are left in Conditions, and the parameter is on the left side (Condition.Nameable)
	//We also save a map of InequalParameter -> []Condition for ease of access later.
	ineqParamConds := make(map[*InequalParameter][]Condition)
	for i, cond := range v.Listen.Conditions {
		if param, ok := cond.Nameable.(*Parameter); ok {
			ineqP := replaceInequalParams[param.Name]
			cond.Nameable = ineqP
			v.Listen.Conditions[i] = cond
			ineqParamConds[ineqP] = append(ineqParamConds[ineqP], cond)
		}
	}
	v.Listen.IneqParamConds = ineqParamConds
}
*/

func (v *ViewInfo) replaceParamsInSelect(replaceEqualParams map[string]EqualParameter) {
	//For all EqualParams, in Select, I can go to every place where the Parameter occours and replace with its Math.
	//For now, we assume no InequalityParams can occour on Select. No TPC-H query needs this support and it would be odd to cache such a thing.
	for key, _ := range v.Listen.SelectParameters {
		if eqParam, has := replaceEqualParams[key]; has {
			v.Listen.SelectAsClause[key] = AsClause{Name: key, Math: CreateMathFromCalculable(eqParam.CalcFrom)}
			delete(v.Listen.SelectParameters, key)
		}
	}
	for key, clause := range v.Listen.SelectAsClause {
		if param, ok := clause.Math.GetCalculable().(Parameter); ok {
			//It is for sure in replaceEqualParams
			clause.Math.Stack.ReplaceTop(replaceEqualParams[param.GetIdentifier()].CalcFrom)
		} else {
			clause.Math.GetCalculable().ReplaceParams(replaceEqualParams)
		}
		v.Listen.SelectAsClause[key] = clause
	}
	for key, aggr := range v.Listen.SelectAsAggregation {
		if param, ok := aggr.Math.GetCalculable().(Parameter); ok {
			aggr.Math.Stack.ReplaceTop(replaceEqualParams[param.GetIdentifier()].CalcFrom)
			v.Listen.SelectAsAggregation[key] = aggr
		}
	}
	for key, count := range v.Listen.SelectAsCount {
		if param, ok := count.Nameable.(Parameter); ok {
			calc := replaceEqualParams[param.GetIdentifier()].CalcFrom
			if calcP, ok := calc.(Parameter); ok {
				count.Nameable = calcP
			} else {
				//Limitation: for now we just replace it with COUNT(*). (Note: no TPC-H query counts parameters)
				count.Nameable = nil
			}
			v.Listen.SelectAsCount[key] = count
		}
	}
}

/*
func (v *ViewInfo) replaceParamsInKeys(replaceEqualParams map[string]EqualParameter, replaceInequalParams map[string]*InequalParameter) {
	for i, nameable := range v.Listen.Keys {
		//Can have both equality and inequality parameters. Also non-parameters. Proceed accordingly.
		if param, ok := nameable.(*Parameter); ok {
			if eqParam, has := replaceEqualParams[param.Name]; has {
				v.Listen.Keys[i] = &eqParam
			} else {
				v.Listen.Keys[i] = replaceInequalParams[param.Name]
			}
		}
	}
}
*/
//For the parameters with equality, replaces them with their equivalent of EqualParameter
func (v *ViewInfo) replaceParamsInKeys(replaceEqualParams map[string]EqualParameter) {
	for i, nameable := range v.Listen.Keys {
		//Can have both equality and inequality parameters. Also non-parameters. Proceed accordingly.
		if param, ok := nameable.(Parameter); ok {
			if eqParam, has := replaceEqualParams[param.GetIdentifier()]; has {
				v.Listen.Keys[i] = &eqParam
			}
		}
	}
}

func (v *ViewInfo) ShouldBeUpdated(objValues map[string]interface{}) bool {
	for _, cond := range v.Listen.Conditions {
		fmt.Printf("[SHOULDBEUPDATED]Condition %s %s %s\n", cond.Nameable.GetIdentifier(),
			cond.Op.ToString(), cond.Math.ToString())
		if !cond.DoesConditionHold(objValues) {
			return false
		}
	}
	//Need to check EqualParameters too
	for _, eqP := range v.Listen.EqParams {
		if !eqP.IsInRange(eqP.CalculateValue(objValues)) {
			fmt.Printf("[SHOULDBEUPDATED]EqParam %v is not in range.\n", eqP.GetIdentifier())
			return false
		}
	}
	fmt.Println("[SHOULDBEUPDATED]ALL OK")
	return true
}

//Pre: ReplaceParameters() and v.Listen.hasInequalityGrouping == false
func (v *ViewInfo) GetKeyForUpdate(objValues map[string]interface{}) (key string) {
	//Most views will have 1-3 entries in Group By.
	if len(v.Listen.Keys) == 1 {
		return ValueToString(v.Listen.Keys[0].CalculateValue(objValues))
	}
	if len(v.Listen.Keys) == 2 {
		fmt.Printf("%+v, %+v\n", v.Listen.Keys[0], v.Listen.Keys[1])
		equalP0, ok := v.Listen.Keys[0].(*EqualParameter)
		if ok {
			equalP1, _ := v.Listen.Keys[1].(*EqualParameter)
			fmt.Printf("%v, %v, %v, %v\n", equalP0.Parameter.ToString(),
				equalP0.CalcFrom.ToString(), equalP1.Parameter.ToString(), equalP1.CalcFrom.ToString())
		}
		return ValueToString(v.Listen.Keys[0].CalculateValue(objValues)) + "_" +
			ValueToString(v.Listen.Keys[1].CalculateValue(objValues))
	}
	if len(v.Listen.Keys) == 3 {
		return ValueToString(v.Listen.Keys[0].CalculateValue(objValues)) + "_" +
			ValueToString(v.Listen.Keys[1].CalculateValue(objValues)) + "_" +
			ValueToString(v.Listen.Keys[2].CalculateValue(objValues))
	}
	var sb strings.Builder
	for _, key := range v.Listen.Keys[:len(v.Listen.Keys)-1] {
		sb.WriteString(ValueToString(key.CalculateValue(objValues)))
		sb.WriteRune('_')
	}
	sb.WriteString(ValueToString(v.Listen.Keys[len(v.Listen.Keys)-1].CalculateValue(objValues)))
	return sb.String()
}

//Pre: ReplaceParameters() and v.Listen.hasInequalityGrouping == true
func (v *ViewInfo) GetKeysWithInequalityForUpdate(objValues map[string]interface{}) (allKeys []string) {
	//v.Listen.Keys only has the following: *Name, *Field, Parameter
	//*Name and *Field are straightforward. If Parameter is of the type EqualParameter, just call CalculateValue()
	//If Parameter is not EqualParameter, then there are inequal conditions - IneqParramConds will have, for each parameter, its conditions
	//So the idea is... for each Parameter, go through its conditions and calculate what the newMin and newMax will be
	//such that the Condition is respected
	//(Important: we already know that there will be at least one value that respects all conditions)
	//After that, the range of keys possible is what is between that newMin and newMax.
	keys := make([][]string, len(v.Listen.Keys))
	for i, nameable := range v.Listen.Keys {
		if param, ok := nameable.(*EqualParameter); ok {
			keys[i] = []string{ValueToString(param.CalculateValue(objValues))}
		} else if param, ok := nameable.(Parameter); ok {
			ineqConds := v.Listen.IneqParamConds[param]
			keys[i] = param.GetIneqKeys(ineqConds, objValues)
		} else { //*Nameable or *Field
			keys[i] = []string{ValueToString(nameable.CalculateValue(objValues))}
		}
	}
	totalSize := 1
	for _, slice := range keys {
		totalSize += len(slice)
	}
	allKeys = make([]string, totalSize)
	i := 0
	//This method will fill allKeys
	i = v.calculateKeysHelper(keys, allKeys, i, "")
	return allKeys[:i]
}

/*
//Pre: ReplaceParameters() and v.Listen.hasInequalityGrouping == true
func (v *ViewInfo) GetKeysWithInequalityForUpdate(objValues map[string]interface{}) (allKeys []string) {
	//v.Listen.Keys only has the following: *Name; *Field, *EqualParameter and *InequalParameter
	//*Name and *Field are straightforward. *EqualParameter too: can just call CalculateValue().
	//*InequalParameter is the tricky one.
	//We have an auxiliary map IneqParamConds that maps *InequalParameter to []Condition
	//So the idea is... for each Condition, calculate what the newMin and newMax will be
	//such that the Condition is respected.
	//(Important: we know already that there will at least be one value that respects all conditions)
	//After that, the range of keys possible is what is between that newMin and newMax.
	//Then to represent the keys... a new array?
	//But how do I do fors inside fors to generate all the keys possible?
	//Maybe with recursivity?
	keys := make([][]string, len(v.Listen.Keys))
	for i, nameable := range v.Listen.Keys {
		if ineqParam, ok := nameable.(*InequalParameter); ok {
			ineqConds := v.Listen.IneqParamConds[ineqParam]
			keys[i] = v.GetIneqKeys(ineqParam, ineqConds, objValues)
			if len(keys[i]) == 0 { //This parameter cannot be satisfied
				keys = make([][]string, 0)
				return
			}
		} else {
			value := nameable.CalculateValue(objValues)
			keys[i] = []string{ValueToString(value)}
		}
	}
	totalSize := 1
	for _, slice := range keys {
		totalSize += len(slice)
	}
	allKeys = make([]string, totalSize)
	i := 0
	//This method will fill allKeys
	v.calculateKeysHelper(keys, allKeys, i, "")
	return allKeys
}
*/

//It's sad but... I really cannot use strings.Builder, as I would have to reset it after the lowest level is written.
//Best I can do is use a base string and share it to avoid excessive copies.
func (v *ViewInfo) calculateKeysHelper(remainingKeys [][]string, allKeys []string, currPos int, baseString string) (newPos int) {
	if len(remainingKeys) > 1 {
		for _, key := range remainingKeys[0] {
			calcString := baseString + key + "_"
			currPos = v.calculateKeysHelper(remainingKeys[1:], allKeys, currPos, calcString)
		}
	} else { //Last level
		for _, key := range remainingKeys[0] {
			allKeys[currPos] = baseString + key
			currPos++
		}
	}
	return currPos
}

/*
func (v *ViewInfo) GetIneqKeys(param *InequalParameter, conds []Condition, objValues map[string]interface{}) (keys []string) {
	switch param.GetType() {
	case INT_TYPE:
		keys = v.GetIneqIntKeys(param, conds, objValues)
	case FLOAT_TYPE:
		keys = v.GetIneqFloatKeys(param, conds, objValues)
	case STRING_TYPE:
		keys = v.GetIneqStringKeys(param, conds, objValues)
	case DATE_TYPE:
		keys = v.GetIneqDateKeys(param, conds, objValues)
	}
	return keys
}

func (v *ViewInfo) GetIneqIntKeys(param *InequalParameter, conds []Condition, objValues map[string]interface{}) (keys []string) {
	min, max := param.MinValue.(int), param.MaxValue.(int)
	for _, cond := range conds {
		value := cond.Math.GetCalculable().CalculateValue(objValues).(int)
		switch cond.Op {
		case HIGHER:
			min = IntMax(min, value+1)
		case HIGHER_EQ:
			min = IntMax(min, value)
		case LOWER:
			max = IntMin(max, value-1)
		case LOWER_EQ:
			max = IntMin(max, value)
		}
	}
	if min < max { //Not possible to satisfact all conditions
		keys = make([]string, 0)
	} else {
		keys = make([]string, max-min+1)
		curr := min
		for i := 0; i < len(keys); i++ {
			keys[i] = strconv.Itoa(curr)
			curr++
		}
	}
	return
}

func (v *ViewInfo) GetIneqFloatKeys(param *InequalParameter, conds []Condition, objValues map[string]interface{}) (keys []string) {
	min, max := param.MinValue.(float64), param.MaxValue.(float64)
	step, precision := param.GetFloatStep(), param.FloatPrecision
	for _, cond := range conds {
		value := cond.Math.GetCalculable().CalculateValue(objValues).(float64)
		switch cond.Op {
		case HIGHER:
			min = math.Max(min, value+step)
		case HIGHER_EQ:
			min = math.Max(min, value)
		case LOWER:
			max = math.Min(max, value-step)
		case LOWER_EQ:
			max = math.Min(max, value)
		}
	}
	if max < min { //Not possible to satisfact all conditions
		keys = make([]string, 0)
	} else {
		keys = make([]string, int((max-min)/step)+1)
		curr := min
		for i := 0; i < len(keys); i++ {
			keys[i] = strconv.FormatFloat(curr, 'f', precision, 64)
			curr += step
		}
	}
	return
}

func (v *ViewInfo) GetIneqStringKeys(param *InequalParameter, conds []Condition, objValues map[string]interface{}) (keys []string) {
	min, max := param.MinValue.(string), param.MaxValue.(string)
	for _, cond := range conds {
		value := cond.Math.GetCalculable().CalculateValue(objValues).(string)
		switch cond.Op {
		//Does it really make much sense to do > or < on strings?
		//For now the implementation of ">" and "<" will
		case HIGHER:
			if value > min {
				lastChar := value[len(value)-1]
				lastChar += 1
				if len(value) == 1 {
					value = string(lastChar)
				} else {
					value = value[:len(value)-2] + string(lastChar)
				}
				min = value
			}
		case HIGHER_EQ:
			if value > min {
				min = value
			}
		case LOWER:
			if value < max {
				lastChar := value[len(value)-1]
				lastChar -= 1
				if len(value) == 1 {
					value = string(lastChar)
				} else {
					value = value[:len(value)-2] + string(lastChar)
				}
				max = value
			}
		case LOWER_EQ:
			if value < max {
				max = value
			}
		}
	}
	keys = make([]string, 0, 10)
	curr := min
	//I feel like this is a dangerous proposition... we shouldn't support range of strings.
	for curr <= max {
		keys = append(keys, curr)
		lastChar := curr[len(curr)-1]
		lastChar += 1
		if len(curr) == 1 {
			curr = string(lastChar)
		} else {
			curr = curr[:len(curr)-2] + string(lastChar)
		}
	}
	return
}

func (v *ViewInfo) GetIneqDateKeys(param *InequalParameter, conds []Condition, objValues map[string]interface{}) (keys []string) {
	min, max := param.MinValue.(Date), param.MaxValue.(Date)
	for _, cond := range conds {
		value := cond.Math.GetCalculable().CalculateValue(objValues).(Date)
		switch cond.Op {
		case HIGHER:
			if value.IsHigher(min) {
				tmp := time.Date(value.Year, time.Month(value.Month), value.Day+1, 0, 0, 0, 0, time.UTC)
				min = Date{Year: tmp.Year(), Month: int(tmp.Month()), Day: int(tmp.Day())}
			}
		case HIGHER_EQ:
			if value.IsHigher(min) {
				min = value
			}
		case LOWER:
			if value.IsLower(max) {
				tmp := time.Date(value.Year, time.Month(value.Month), value.Day-1, 0, 0, 0, 0, time.UTC)
				max = Date{Year: tmp.Year(), Month: int(tmp.Month()), Day: int(tmp.Day())}
			}
		case LOWER_EQ:
			if value.IsLower(max) {
				max = value
			}
		}
	}

	minD := time.Date(min.Year, time.Month(min.Month), min.Day, 0, 0, 0, 0, time.UTC)
	maxD := time.Date(max.Year, time.Month(max.Month), max.Day, 0, 0, 0, 0, time.UTC)
	diffDays := int(maxD.Sub(minD).Hours()/24) + 1
	if diffDays <= 0 {
		keys = make([]string, 0)
	} else {
		keys = make([]string, diffDays)
		for i := 0; i < len(keys); i++ {
			keys[i] = Date{Year: minD.Year(), Month: int(minD.Month()), Day: int(minD.Day())}.ToString()
			minD = minD.AddDate(0, 0, 1)
		}
	}
	return
}
*/

func (v *ViewInfo) PrintPreparedQuery() {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("CREATE VIEW (%s, %s) AS\n", v.Listen.ID, v.Listen.Bucket))
	sb.WriteString("SELECT ")
	for _, entry := range v.Listen.SelectByAppearenceOrder {
		id := entry.GetIdentifier()
		if nameable, has := v.Listen.SelectNameable[id]; has {
			nameable.ToPrintString(&sb)
		} else if asClause, has := v.Listen.SelectAsClause[id]; has {
			asClause.Math.GetCalculable().ToPrintString(&sb)
			sb.WriteString(" AS ")
			sb.WriteString(asClause.Name)
		} else if asAggr, has := v.Listen.SelectAsAggregation[id]; has {
			aggrFun := asAggr.AggrFunc
			switch aggrFun {
			case SUM:
				sb.WriteString("SUM(")
			case AVG:
				sb.WriteString("AVG(")
			case MAX:
				sb.WriteString("MAX(")
			case MIN:
				sb.WriteString("MIN(")
			}
			asAggr.Math.GetCalculable().ToPrintString(&sb)
			sb.WriteString(") AS ")
			sb.WriteString(asAggr.Name)
		} else {
			asCount := v.Listen.SelectAsCount[id]
			sb.WriteString("COUNT(")
			if asCount.Nameable == nil {
				sb.WriteRune('*')
			} else {
				asCount.Nameable.ToPrintString(&sb)
			}
			sb.WriteString(") AS ")
			sb.WriteString(asCount.Name)
		}
		sb.WriteString(", ")
	}
	sb.WriteRune('\n')

	sb.WriteString("FROM ")
	for _, key := range v.Listen.Containers {
		sb.WriteString(key)
		sb.WriteString(", ")
	}
	sb.WriteRune('\n')

	if len(v.Listen.Conditions) > 0 {
		sb.WriteString("WHERE ")
		for _, cond := range v.Listen.Conditions {
			cond.Nameable.ToPrintString(&sb)
			switch cond.Op {
			case EQUAL:
				sb.WriteString(" = ")
			case HIGHER:
				sb.WriteString(" > ")
			case HIGHER_EQ:
				sb.WriteString(" >= ")
			case LOWER:
				sb.WriteString(" < ")
			case LOWER_EQ:
				sb.WriteString(" <= ")
			case NOT_EQ:
				sb.WriteString(" != ")
			case FUNCTION:
				sb.WriteRune(' ')
				sb.WriteString(cond.Fun.GetFunName())
				sb.WriteRune('(')
			}
			cond.Math.GetCalculable().ToPrintString(&sb)
			if cond.Op == FUNCTION {
				sb.WriteRune(')')
			}
			sb.WriteString(", ")
		}
		sb.WriteRune('\n')
	}

	sb.WriteString("GROUP BY ")
	for _, key := range v.Listen.Keys {
		key.ToPrintString(&sb)
		sb.WriteString(", ")
	}
	sb.WriteRune('\n')

	if len(v.Listen.Sorting) > 0 {
		sb.WriteString("ORDER BY ")
		for _, order := range v.Listen.Sorting {
			order.Nameable.ToPrintString(&sb)
			sb.WriteRune(' ')
			if order.IsDesc {
				sb.WriteString("DESC, ")
			} else {
				sb.WriteString("ASC, ")
			}
		}
		sb.WriteRune('\n')
	}

	if v.Listen.TopSize > 0 {
		sb.WriteString("LIMIT ")
		sb.WriteString(strconv.Itoa(v.Listen.TopSize))
		sb.WriteRune('\n')
	}
	fmt.Println(sb.String())
}

/*
CREATE VIEW (Q3, views) AS
			SELECT OrderKey, SUM(Price * (1 - Discount))
			AS Revenue, OrderDate, ShipPriority
			FROM Lineitem
			WHERE OrderDate < [DATE] AND ShipDate > [DATE] AND [DATE] >= 1995-03-01 AND [DATE] <= 1995-03-31
			GROUP BY OrderKey, OrderDate, ShipPriority
			GROUP BY [DATE], Segment
			ORDER BY Revenue DESC, OrderDate ASC
			LIMIT 10
*/

//Pre: all prepares.
//Calculates the values returned by each select and stores in a map
//TODO: I may actually want to do a specific one per CRDT type.
/*
func (v *ViewInfo) CalculateViewData(objValues map[string]interface{}) (viewData map[string]interface{}) {
	viewData = make(map[string]interface{})
	for key, nameable := range v.Listen.SelectNameable {
		viewData[key] = nameable.CalculateValue(objValues)
	}
	for key, asClause := range v.Listen.SelectAsClause {
		viewData[key] = asClause.Math.GetCalculable().CalculateValue(objValues)
	}
	//TODO: Actually aren't sums, maxs, mins and counts special cases...?
	//I guess it must really be dependent on the CRDT.
	//Oh well. After dinner.
	for key, aggr := range v.Listen.SelectAsAggregation {
		viewData[key] = aggr.Math.GetCalculable().CalculateValue(objValues)
	}
	for key, count := range v.Listen.SelectAsCount {
		viewData[key] = count.Nameable.CalculateValue(objValues)
	}
	return
}*/

func IntMax(first, second int) int {
	if first > second {
		return first
	}
	return second
}

func IntMin(first, second int) int {
	if first < second {
		return first
	}
	return second
}

func ignore(any interface{}) {}
