package sql

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"sqlToKeyValue/src/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

/*
Idea later for generating updates:
I will want a map of all "names" (i.e., IDs) and "parameters" with their values based on the update
Then to each Math (and probably other things), pass all of those and let them access those as needed.
*/

//TODO: We probably want to keep the order of entries in Select.
//This will have to SelectNameable - SelectAsCount

type MyViewSQLListener struct {
	Parse *parser.ViewSQLParser

	ID     string
	Bucket string

	Mode int

	//All nameables found, in all clauses. The idea is to use a single instance of Nameable per different name.
	//NOTE: Both of those below are unused so far.
	AllNameables  map[string]Nameable
	AllParameters map[string]Parameter

	//Select
	//NOTE: The three below are unused so far.
	SelectIDs        map[string]struct{}
	SelectFields     map[string]DateFieldType
	SelectParameters map[string]struct{}
	//SelectParameters map[string]Parameter

	SelectNameable          map[string]Nameable
	SelectAsClause          map[string]AsClause
	SelectAsAggregation     map[string]AsAggregation
	SelectAsCount           map[string]AsCount
	AllSelect               map[string]SelectFields
	NEntriesInSelect        int
	PrimaryKey              Nameable
	SelectByAppearenceOrder []SelectFields

	//From
	Containers []string

	//Where
	WhereIDs        map[string]struct{}
	WhereFields     map[string]DateFieldType
	WhereParameters map[string]struct{}
	Conditions      []Condition
	//IneqParamConds  map[*InequalParameter][]Condition //Built by ReplaceParameters() in viewInfo.go
	IneqParamConds map[Parameter][]Condition //Built by ReplaceParameters() in viewInfo.go
	EqParams       []EqualParameter          //Also need to check these when checking the conditions

	//Second Group By
	Keys []Nameable //After processing, all Parameters will be replaced with *EqualParameter and *InequalParameter

	//Order By
	Sorting []Order

	//Limit
	TopSize int //0 if limit is not defined (aka, shouldn't be a top)

	currentMath Math

	//If there's two group bys, this is set by "view" to true.
	//After the first group by, this is set to false.
	ignoreGroupBy bool

	//Has a parameter by inequality. This requires a different algorithm to build keys.
	//Filled by ReplaceParams() in viewInfo.go
	HasInequalityGrouping bool
}

type SelectFields interface {
	GetCalculable() Calculable
	GetIdentifier() string
}

type AsClause struct {
	Math
	Name string
}

type AsAggregation struct {
	Math
	Name     string
	AggrFunc AggrType
}

type AsCount struct {
	Nameable //if nil then it is '*'
	Name     string
}

func (as AsClause) GetCalculable() Calculable      { return as.Math.GetCalculable() }
func (as AsClause) GetIdentifier() string          { return as.Name }
func (as AsAggregation) GetCalculable() Calculable { return as.Math.GetCalculable() }
func (as AsAggregation) GetIdentifier() string     { return as.Name }
func (as AsCount) GetCalculable() Calculable       { return as.Nameable }
func (as AsCount) GetIdentifier() string           { return as.Name }

/*func (name *Name) GetCalculable() Calculable                     { return name }
func (name *Name) GetName() string                               { return name.Name }
func (field *Field) GetCalculable() Calculable                   { return field }
func (field *Field) GetName() string                             { return field.Name }
func (inP *InParameter) GetName() string                         { return inP.GetIdentifier() }
func (rangeP *RangeParameterInt) GetCalculable() Calculable      { return rangeP }
func (rangeP *RangeParameterFloat) GetCalculable() Calculable    { return rangeP }
func (rangeP *RangeParameterDate) GetCalculable() Calculable     { return rangeP }
func (sparseP *SparseParameterInt) GetCalculable() Calculable    { return sparseP }
func (sparseP *SparseParameterFloat) GetCalculable() Calculable  { return sparseP }
func (sparseP *SparseParameterDate) GetCalculable() Calculable   { return sparseP }
func (sparseP *SparseParameterString) GetCalculable() Calculable { return sparseP }*/

type Condition struct {
	Math
	Nameable
	Op  CondType
	Fun Function
}

type Order struct {
	Nameable
	IsDesc bool
}

type CondType int
type AggrType int

func (c CondType) ToString() string {
	switch c {
	case EQUAL:
		return "="
	case LOWER_EQ:
		return "<="
	case LOWER:
		return "<"
	case HIGHER_EQ:
		return ">="
	case HIGHER:
		return ">"
	case NOT_EQ:
		return "!="
	case FUNCTION: //Should instead call Fun.ToString()
		return "Function"
	}
	return "Unknown"
}

func (a AggrType) ToString() string {
	switch a {
	case SUM:
		return "SUM"
	case AVG:
		return "AVG"
	case MAX:
		return "MAX"
	case MIN:
		return "MIN"
	}
	return "Unknown"
}

func (c AsCount) IsCountAll() bool {
	return c.Nameable == nil
}

func (c Condition) ToString() string {
	return fmt.Sprintf("{Nameable: %s; Math: %s, Op: %d, Fun: %v}", c.Math.ToString(),
		c.Nameable.GetIdentifier(), c.Op, c.Fun)
}

func (c Condition) DoesConditionHold(objValues map[string]interface{}) bool {
	//Need to check if value of EqualParameter is in the valid range
	if eqParam, ok := c.Nameable.(*EqualParameter); ok {
		mathV := c.Math.GetCalculable().CalculateValue(objValues)
		fmt.Printf("[DCH]EqualParameter. MathV %v, EqParameter %v\n", mathV, eqParam)
		return eqParam.IsInRange(mathV)
	}
	if param, ok := c.Nameable.(Parameter); ok {
		mathV := c.Math.GetCalculable().CalculateValue(objValues)
		fmt.Printf("[DCH]Parameter. MathV %v, Parameter %v\n", mathV, param)
		if c.Op == EQUAL {
			return param.IsInRange(mathV)
		} else {
			return param.IsInInequalityRange(mathV, c.Op)
		}
	} else { //Non-parameter
		mathV, nameableV := c.Math.GetCalculable().CalculateValue(objValues), c.Nameable.CalculateValue(objValues)
		fmt.Printf("[DCH]Non-Parameter. MathV %v, NameableV %v\n", mathV, nameableV)
		if c.Op == FUNCTION {
			switch typedFun := c.Fun.(type) {
			case StartsWithFun:
				return typedFun.ApplyFun(nameableV.(string), mathV.(string))
			case EndsWithFun:
				return typedFun.ApplyFun(nameableV.(string), mathV.(string))
			case ContainsFun:
				return typedFun.ApplyFun(nameableV.(string), mathV.(string))
			}
		} else {
			switch mathVTyped := mathV.(type) {
			case int:
				switch nameableVTyped := mathV.(type) {
				case int:
					return compareInts(nameableVTyped, mathVTyped, c.Op)
				case float64:
					return compareFloats(nameableVTyped, float64(mathVTyped), c.Op)
				}
			case float64:
				switch nameableVTyped := mathV.(type) {
				case float64:
					return compareFloats(nameableVTyped, mathVTyped, c.Op)
				case int:
					return compareFloats(float64(nameableVTyped), mathVTyped, c.Op)
				}
			case Date:
				return compareDates(nameableV.(Date), mathVTyped, c.Op)
			case string:
				return compareStrings(nameableV.(string), mathVTyped, c.Op)
			}
		}
	}
	fmt.Println("[DCH]Didn't match anything!")
	return false
}

func compareInts(firstI, secondI int, op CondType) bool {
	switch op {
	case EQUAL:
		return firstI == secondI
	case NOT_EQ:
		return firstI != secondI
	case HIGHER:
		return firstI > secondI
	case HIGHER_EQ:
		return firstI >= secondI
	case LOWER:
		return firstI < secondI
	case LOWER_EQ:
		return firstI <= secondI
	}
	return false
}

func compareFloats(firstI, secondI float64, op CondType) bool {
	switch op {
	case EQUAL:
		return firstI == secondI
	case NOT_EQ:
		return firstI != secondI
	case HIGHER:
		return firstI > secondI
	case HIGHER_EQ:
		return firstI >= secondI
	case LOWER:
		return firstI < secondI
	case LOWER_EQ:
		return firstI <= secondI
	}
	return false
}

func compareDates(firstI, secondI Date, op CondType) bool {
	switch op {
	case EQUAL:
		return firstI == secondI
	case NOT_EQ:
		return firstI != secondI
	case HIGHER:
		return firstI.IsHigher(secondI)
	case HIGHER_EQ:
		return firstI.IsHigherOrEqual(secondI)
	case LOWER:
		return firstI.IsLower(secondI)
	case LOWER_EQ:
		return firstI.IsLowerOrEqual(secondI)
	}
	return false
}

func compareStrings(firstI, secondI string, op CondType) bool {
	switch op {
	case EQUAL:
		return firstI == secondI
	case NOT_EQ:
		return firstI != secondI
	}
	return false
}

//Pre: ReplaceParameters() (on viewInfo.go) was already called
//So now we only have conditions of the kind:
//variable >, >=, <, <=, =, !=, func variable; INEQUALITY_PARAMETER >=, >, <=, <, != variable
//I can do special processing when nameableV is an INEQUALITY_PARAMETER
/*
func (c Condition) DoesConditionHold(objValues map[string]interface{}) bool {
	if ineqParam, ok := c.Nameable.(*InequalParameter); ok {
		mathV := c.Math.GetCalculable().CalculateValue(objValues)
		fmt.Printf("[CONDITIONHOLD]Nameable is ineqParam: %s %s %v\n", ineqParam.Name, c.Op.ToString(), mathV)
		fmt.Printf("[CONDITIONHOLD]Types: %v, %v\n", ineqParam.GetType(), c.Math.GetCalculable().GetType())
		fmt.Printf("[CONDITIONHOLD]%+v, %+v\n", ineqParam, c.Math.GetCalculable())
		switch c.Op {
		case EQUAL:
			return ineqParam.IsInRange(mathV)
		case NOT_EQ:
			return true //It's always different from at least one value...
		}
		//case HIGHER: //param > value -> param.MAX > value;
		//case HIGHER_EQ: //param >= value -> param.MAX >= value;
		//case LOWER: //param < value -> param.MIN < value;
		//case LOWER_EQ: //param <= value -> param.MIN <= value;
		switch ineqParam.Type {
		case INT_TYPE:
			switch mathVTyped := mathV.(type) {
			case int:
				return compIneqParamsInts(c.Op, ineqParam, mathVTyped)
			case float64:
				return compIneqParamsFloatInt(c.Op, ineqParam, mathVTyped)
			}
		case FLOAT_TYPE:
			switch mathVTyped := mathV.(type) {
			case float64:
				return compIneqParamsFloat(c.Op, ineqParam, mathVTyped)
			case int:
				return compIneqParamsFloat(c.Op, ineqParam, float64(mathVTyped))
			}
		case DATE_TYPE:
			return compIneqParamsDate(c.Op, ineqParam, mathV.(Date))
		case STRING_TYPE:
			return compIneqParamsString(c.Op, ineqParam, mathV.(string))
		}
	} else {
		mathV, nameableV := c.Math.GetCalculable().CalculateValue(objValues), c.Nameable.CalculateValue(objValues)
		fmt.Printf("[CONDITIONHOLD]No ineqParam: %s %s %v\n", c.Nameable.GetIdentifier(), c.Op.ToString(), c.Math.ToString())
		fmt.Printf("[CONDITIONHOLD]No ineqParam (values): %v %s %v\n", nameableV, c.Op.ToString(), mathV)

		switch c.Op {
		case EQUAL:
			return mathV == nameableV
		case NOT_EQ:
			return mathV != nameableV
		case FUNCTION: //Must be strings
			switch typedFun := c.Fun.(type) {
			case StartsWithFun:
				return typedFun.ApplyFun(nameableV.(string), mathV.(string))
			case EndsWithFun:
				return typedFun.ApplyFun(nameableV.(string), mathV.(string))
			case ContainsFun:
				return typedFun.ApplyFun(nameableV.(string), mathV.(string))
			}
		}

		switch mathVTyped := mathV.(type) {
		case int:
			switch nameableVTyped := nameableV.(type) {
			case int:
				fmt.Printf("[CONDITIONHOLD]Both ints: %d %s %d\n", nameableVTyped, c.Op.ToString(), mathVTyped)
				switch c.Op {
				case HIGHER:
					return nameableVTyped > mathVTyped
				case HIGHER_EQ:
					return nameableVTyped >= mathVTyped
				case LOWER:
					return nameableVTyped < mathVTyped
				case LOWER_EQ:
					return nameableVTyped <= mathVTyped
				}
			case float64:
				switch c.Op {
				case HIGHER:
					return nameableVTyped > float64(mathVTyped)
				case HIGHER_EQ:
					return nameableVTyped >= float64(mathVTyped)
				case LOWER:
					return nameableVTyped < float64(mathVTyped)
				case LOWER_EQ:
					return nameableVTyped <= float64(mathVTyped)
				}
			}
		case float64:
			switch nameableVTyped := nameableV.(type) {
			case float64:
				switch c.Op {
				case HIGHER:
					return nameableVTyped > mathVTyped
				case HIGHER_EQ:
					return nameableVTyped >= mathVTyped
				case LOWER:
					return nameableVTyped < mathVTyped
				case LOWER_EQ:
					return nameableVTyped <= mathVTyped
				}
			case int:
				switch c.Op {
				case HIGHER:
					return float64(nameableVTyped) > mathVTyped
				case HIGHER_EQ:
					return float64(nameableVTyped) >= mathVTyped
				case LOWER:
					return float64(nameableVTyped) < mathVTyped
				case LOWER_EQ:
					return float64(nameableVTyped) <= mathVTyped
				}
			}
		case Date:
			dateNameable := nameableV.(Date)
			switch c.Op {
			case HIGHER:
				return dateNameable.IsHigher(mathV.(Date))
			case HIGHER_EQ:
				return dateNameable.IsHigherOrEqual(mathV.(Date))
			case LOWER:
				return dateNameable.IsLower(mathV.(Date))
			case LOWER_EQ:
				return dateNameable.IsLowerOrEqual(mathV.(Date))
			}
		case string:
			stringNameable := nameableV.(string)
			switch c.Op {
			case HIGHER:
				return stringNameable > mathV.(string)
			case HIGHER_EQ:
				return stringNameable >= mathV.(string)
			case LOWER:
				return stringNameable < mathV.(string)
			case LOWER_EQ:
				return stringNameable <= mathV.(string)
			}
		default:
			fmt.Printf("How...? %T %v\n", mathVTyped, mathVTyped)
		}
	}
	return false
}

//Note: Doesn't compare NOT_EQ
func compIneqParamsInts(op CondType, ineqParam *InequalParameter, mathV int) bool {
	switch op {
	case HIGHER:
		return ineqParam.MaxValue.(int) > mathV
	case HIGHER_EQ:
		return ineqParam.MaxValue.(int) >= mathV
	case LOWER:
		return ineqParam.MinValue.(int) < mathV
	case LOWER_EQ:
		return ineqParam.MinValue.(int) <= mathV
	}
	return false
}

//In this case the ineqParam is in int but we should convert to float
func compIneqParamsFloatInt(op CondType, ineqParam *InequalParameter, mathV float64) bool {
	switch op {
	case HIGHER:
		return float64(ineqParam.MaxValue.(int)) > mathV
	case HIGHER_EQ:
		return float64(ineqParam.MaxValue.(int)) >= mathV
	case LOWER:
		return float64(ineqParam.MinValue.(int)) < mathV
	case LOWER_EQ:
		return float64(ineqParam.MinValue.(int)) <= mathV
	}
	return false
}

//Note: Doesn't compare NOT_EQ
func compIneqParamsFloat(op CondType, ineqParam *InequalParameter, mathV float64) bool {
	switch op {
	case HIGHER:
		return ineqParam.MaxValue.(float64) > mathV
	case HIGHER_EQ:
		return ineqParam.MaxValue.(float64) >= mathV
	case LOWER:
		return ineqParam.MinValue.(float64) < mathV
	case LOWER_EQ:
		return ineqParam.MinValue.(float64) <= mathV
	}
	return false
}

//Note: Doesn't compare NOT_EQ
func compIneqParamsString(op CondType, ineqParam *InequalParameter, mathV string) bool {
	switch op {
	case HIGHER:
		return ineqParam.MaxValue.(string) > mathV
	case HIGHER_EQ:
		return ineqParam.MaxValue.(string) >= mathV
	case LOWER:
		return ineqParam.MinValue.(string) < mathV
	case LOWER_EQ:
		return ineqParam.MinValue.(string) <= mathV
	}
	return false
}

//Note: Doesn't compare NOT_EQ
func compIneqParamsDate(op CondType, ineqParam *InequalParameter, mathV Date) bool {
	switch op {
	case HIGHER:
		return ineqParam.MaxValue.(Date).IsHigher(mathV)
	case HIGHER_EQ:
		return ineqParam.MaxValue.(Date).IsHigherOrEqual(mathV)
	case LOWER:
		return ineqParam.MinValue.(Date).IsLower(mathV)
	case LOWER_EQ:
		return ineqParam.MinValue.(Date).IsLowerOrEqual(mathV)
	}
	return false
}
*/

const (
	SELECT          = 0
	FROM            = 1
	WHERE           = 2
	GROUP_BY_FIRST  = 3
	GROUP_BY_SECOND = 4
	ORDER_BY        = 5
	LIMIT           = 6

	SUM = AggrType(0)
	AVG = AggrType(1)
	MAX = AggrType(2)
	MIN = AggrType(3)

	EQUAL     = CondType(10)
	HIGHER    = CondType(11)
	HIGHER_EQ = CondType(12)
	LOWER     = CondType(13)
	LOWER_EQ  = CondType(14)
	NOT_EQ    = CondType(15)
	FUNCTION  = CondType(20)
)

func CreateViewSQLListener(parse *parser.ViewSQLParser) *MyViewSQLListener {
	return &MyViewSQLListener{Parse: parse, AllNameables: make(map[string]Nameable), AllParameters: make(map[string]Parameter)}
}

// VisitTerminal is called when a terminal node is visited.
func (listen *MyViewSQLListener) VisitTerminal(node antlr.TerminalNode) {
	//fmt.Println("[Listen]Reached terminal node.")
}

// VisitErrorNode is called when an error node is visited.
func (listen *MyViewSQLListener) VisitErrorNode(node antlr.ErrorNode) {
	//fmt.Println("[Listen]Reached error node.")
}

// EnterEveryRule is called when any rule is entered.
func (listen *MyViewSQLListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	//fmt.Println("[Listen]Entered a rule.")
}

// ExitEveryRule is called when any rule is exited.
func (listen *MyViewSQLListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	//fmt.Println("[Listen]Exitted a rule.")
}

// EnterName is called when production name is entered.
func (listen *MyViewSQLListener) EnterName(ctx *parser.NameContext) {
	//ctx.GetText()		//Returns child description
	//fmt.Printf("[Listen]Entered %s\n", parse.SymbolicNames[ctx.RuleIndex])
	//fmt.Println("[Listen]Entered Name")
}

// ExitName is called when production name is exited.
func (listen *MyViewSQLListener) ExitName(ctx *parser.NameContext) {
	//ctx.GetText()		//Returns child description
	//fmt.Printf("[Listen]Exitted %s\n", parse.SymbolicNames[ctx.RuleIndex])
	//fmt.Println("[Listen]Exit Name")
}

// EnterConstant is called when production constant is entered.
func (listen *MyViewSQLListener) EnterConstant(ctx *parser.ConstantContext) {
	//fmt.Println("[Listen]Entered Constant")
}

// ExitConstant is called when production constant is exited.
func (listen *MyViewSQLListener) ExitConstant(ctx *parser.ConstantContext) {
	//fmt.Println("[Listen]Exit Constant")
}

// EnterField is called when production field is entered.
func (listen *MyViewSQLListener) EnterField(ctx *parser.FieldContext) {
	//fmt.Println("[Listen]Entered Field")
}

// ExitField is called when production field is exited.
func (listen *MyViewSQLListener) ExitField(ctx *parser.FieldContext) {
	//fmt.Println("[Listen]Exit Field")
}

// EnterParameter is called when production parameter is entered.
func (listen *MyViewSQLListener) EnterParameter(ctx *parser.ParameterContext) {
	//fmt.Println("[Listen]Entered Parameter")
}

// ExitParameter is called when production parameter is exited.
func (listen *MyViewSQLListener) ExitParameter(ctx *parser.ParameterContext) {
	//fmt.Println("[Listen]Exit Parameter")
}

// EnterNameable is called when production nameable is entered.
func (listen *MyViewSQLListener) EnterNameable(ctx *parser.NameableContext) {
	//fmt.Println("[Listen]Entered Nameable")
}

// ExitNameable is called when production nameable is exited.
func (listen *MyViewSQLListener) ExitNameable(ctx *parser.NameableContext) {
	//fmt.Println("[Listen]Exit Nameable")
	parent := ctx.GetParent()
	if _, ok := parent.(*parser.CalcContext); ok { //Not a nameable inside something else
		nameable := listen.FullMakeNameableFromContext(ctx)
		listen.SelectNameable[nameable.GetIdentifier()] = nameable
	}
	if listen.Mode == SELECT {
		listen.HandleNameableInSelect(ctx)
	} else if listen.Mode == WHERE {
		listen.HandleNameableInWhere(ctx)
	}
}

// EnterValue is called when production value is entered.
func (listen *MyViewSQLListener) EnterValue(ctx *parser.ValueContext) {
	//fmt.Println("[Listen]Entered Value")
}

// ExitValue is called when production value is exited.
func (listen *MyViewSQLListener) ExitValue(ctx *parser.ValueContext) {
	//fmt.Println("[Listen]Exit Value")

	if nameable := ctx.Nameable(); nameable != nil {
		listen.currentMath.Stack.Push(listen.FullMakeNameableFromContext(nameable))
	} else { //Constant
		constant := MakeConstantFromContext(ctx.Constant())
		listen.currentMath.Stack.Push(constant)
	}
}

// EnterValue is called when production value is entered.
func (listen *MyViewSQLListener) EnterMinus(ctx *parser.MinusContext) {
	//fmt.Println("[Listen]Entered Minus")
}

// ExitValue is called when production value is exited.
func (listen *MyViewSQLListener) ExitMinus(ctx *parser.MinusContext) {
	//fmt.Println("[Listen]Exit Minus")
	//listen.currentMath.Stack.Push(&Minus{ToCalculate: listen.currentMath.Stack.Pop()})
	MinusHelper(ctx, listen.currentMath)
}

// EnterValue is called when production value is entered.
func (listen *MyViewSQLListener) EnterMultOrDiv(ctx *parser.MultOrDivContext) {
	//fmt.Println("[Listen]Entered MultOrDiv")
}

// ExitValue is called when production value is exited.
func (listen *MyViewSQLListener) ExitMultOrDiv(ctx *parser.MultOrDivContext) {
	//fmt.Println("[Listen]Exit MultOrDiv")
	/*right, left := listen.currentMath.Stack.Pop(), listen.currentMath.Stack.Pop()
	if ctx.GetOpType().GetText() == "*" {
		listen.currentMath.Stack.Push(&Mult{Left: left, Right: right})
	} else {
		listen.currentMath.Stack.Push(&Div{Left: left, Right: right})
	}*/
	MultOrDivHelper(ctx, listen.currentMath)
}

// EnterValue is called when production value is entered.
func (listen *MyViewSQLListener) EnterAddOrSub(ctx *parser.AddOrSubContext) {
	//fmt.Println("[Listen]Entered AddOrSub")
}

// ExitValue is called when production value is exited.
func (listen *MyViewSQLListener) ExitAddOrSub(ctx *parser.AddOrSubContext) {
	//fmt.Println("[Listen]Exit AddOrSub")
	/*right, left := listen.currentMath.Stack.Pop(), listen.currentMath.Stack.Pop()
	if ctx.GetOpType().GetText() == "+" {
		listen.currentMath.Stack.Push(&Add{Left: left, Right: right})
	} else {
		listen.currentMath.Stack.Push(&Sub{Left: left, Right: right})
	}
	*/
	AddOrSubHelper(ctx, listen.currentMath)
}

/*
// EnterMath is called when production math is entered.
func (listen *MyViewSQLListener) EnterMath(ctx *parser.MathContext) {
	//fmt.Println("[Listen]Entered Math")
	listen.currentMath = Math{}
}

// ExitMath is called when production math is exited.
func (listen *MyViewSQLListener) ExitMath(ctx *parser.MathContext) {
	//fmt.Println("[Listen]Exit Math")
}
*/

// EnterAsClause is called when production asClause is entered.
func (listen *MyViewSQLListener) EnterAsClause(ctx *parser.AsClauseContext) {
	//fmt.Println("[Listen]Entered AsClause")
	//Need to prepare the stack
	listen.currentMath = CreateMath()
}

// ExitAsClause is called when production asClause is exited.
func (listen *MyViewSQLListener) ExitAsClause(ctx *parser.AsClauseContext) {
	//fmt.Println("[Listen]Exit AsClause")
	clause := AsClause{Math: listen.currentMath, Name: ctx.Name().STRING().GetText()}
	listen.SelectAsClause[clause.Name] = clause
}

// EnterAggregation is called when production aggregation is entered.
func (listen *MyViewSQLListener) EnterAggregation(ctx *parser.AggregationContext) {
	//fmt.Println("[Listen]Entered Aggregation")
	//Need to prepare the stack
	listen.currentMath = CreateMath()
}

// ExitAggregation is called when production aggregation is exited.
func (listen *MyViewSQLListener) ExitAggregation(ctx *parser.AggregationContext) {
	//fmt.Println("[Listen]Exit Aggregation")

	//Handling aggrFunc
	aggr := ctx.AggrFunc()
	var aggrType AggrType
	if sum := aggr.SUM(); sum != nil {
		aggrType = SUM
	} else if avg := aggr.AVG(); avg != nil {
		aggrType = AVG
	} else if max := aggr.MAX(); max != nil {
		aggrType = MAX
	} else { //min
		aggrType = MIN
	}

	clause := AsAggregation{AggrFunc: aggrType, Math: listen.currentMath, Name: ctx.Name().STRING().GetText()}
	listen.SelectAsAggregation[clause.Name] = clause
}

// EnterCount is called when production count is entered.
func (listen *MyViewSQLListener) EnterCount(ctx *parser.CountContext) {
	//fmt.Println("[Listen]Entered Count")
}

// ExitCount is called when production count is exited.
func (listen *MyViewSQLListener) ExitCount(ctx *parser.CountContext) {
	//fmt.Println("[Listen]Exit Count")

	asCount := AsCount{}
	if nameable := ctx.Nameable(); nameable != nil {
		asCount.Nameable = listen.FullMakeNameableFromContext(nameable)
	} //else: '*'. Nothing needs to be done in this case.
	asCount.Name = ctx.Name().STRING().GetText()
	listen.SelectAsCount[asCount.Name] = asCount
}

// EnterComp is called when production comp is entered.
func (listen *MyViewSQLListener) EnterComp(ctx *parser.CompContext) {
	//fmt.Println("[Listen]Entered Comp")
}

// ExitComp is called when production comp is exited.
func (listen *MyViewSQLListener) ExitComp(ctx *parser.CompContext) {
	//fmt.Println("[Listen]Exit Comp")
}

// EnterCondition is called when production condition is entered.
func (listen *MyViewSQLListener) EnterCondition(ctx *parser.ConditionContext) {
	//fmt.Println("[Listen]Entered Condition")
	listen.currentMath = CreateMath()
}

// ExitCondition is called when production condition is exited.
func (listen *MyViewSQLListener) ExitCondition(ctx *parser.ConditionContext) {
	//fmt.Println("[Listen]Exit Condition")

	cond := Condition{
		Math:     listen.currentMath,
		Nameable: listen.FullMakeNameableFromContext(ctx.Nameable()),
	}
	comp := ctx.Comp()
	if opType := comp.GetOpType(); opType != nil {
		cond.Op = StringOperatorToCondOp(opType)
	} else { //function name
		cond.Fun = MakeFunFromName(comp.Name().GetText())
		cond.Op = FUNCTION
	}
	listen.Conditions = append(listen.Conditions, cond)
}

// EnterSortOrder is called when production sortOrder is entered.
func (listen *MyViewSQLListener) EnterSortOrder(ctx *parser.SortOrderContext) {
	//fmt.Println("[Listen]Entered SortOrder")
}

// ExitSortOrder is called when production sortOrder is exited.
func (listen *MyViewSQLListener) ExitSortOrder(ctx *parser.SortOrderContext) {
	//fmt.Println("[Listen]Exit SortOrder")
}

// EnterCreate is called when production create is entered.
func (listen *MyViewSQLListener) EnterCreate(ctx *parser.CreateContext) {
	//fmt.Println("[Listen]Entered Create")
	listen.ID, listen.Bucket = ctx.STRING(0).GetText(), ctx.STRING(1).GetText()
}

// ExitCreate is called when production create is exited.
func (listen *MyViewSQLListener) ExitCreate(ctx *parser.CreateContext) {
	//fmt.Println("[Listen]Exit Create")
}

// EnterSelect is called when production select is entered.
func (listen *MyViewSQLListener) EnterSelect(ctx *parser.SelectContext) {
	//fmt.Println("[Listen]Entered Select")
	listen.Mode = SELECT
	listen.SelectIDs = make(map[string]struct{})
	listen.SelectFields = make(map[string]DateFieldType)
	listen.SelectParameters = make(map[string]struct{})
	//listen.SelectParameters = make(map[string]Parameter)
	listen.SelectNameable = make(map[string]Nameable)
	listen.SelectAsClause = make(map[string]AsClause)
	listen.SelectAsAggregation = make(map[string]AsAggregation)
	listen.SelectAsCount = make(map[string]AsCount)
	listen.AllSelect = make(map[string]SelectFields)
	listen.SelectByAppearenceOrder = make([]SelectFields, 0, 10)
}

// ExitSelect is called when production select is exited.
func (listen *MyViewSQLListener) ExitSelect(ctx *parser.SelectContext) {
	//fmt.Println("[Listen]Exit Select")
}

// EnterFrom is called when production from is entered.
func (listen *MyViewSQLListener) EnterFrom(ctx *parser.FromContext) {
	//fmt.Println("[Listen]Entered From")
	listen.Mode = FROM
}

// ExitFrom is called when production from is exited.
func (listen *MyViewSQLListener) ExitFrom(ctx *parser.FromContext) {
	//fmt.Println("[Listen]Exit From")

	allNames := ctx.AllName()
	listen.Containers = make([]string, len(allNames))
	for i, name := range allNames {
		listen.Containers[i] = name.STRING().GetText()
	}
}

// EnterWhere is called when production where is entered.
func (listen *MyViewSQLListener) EnterWhere(ctx *parser.WhereContext) {
	//fmt.Println("[Listen]Entered Where")

	listen.Mode = WHERE
	listen.Conditions = make([]Condition, 0, len(ctx.AllCondition()))
	listen.WhereIDs = make(map[string]struct{})
	listen.WhereFields = make(map[string]DateFieldType)
	listen.WhereParameters = make(map[string]struct{})
}

// ExitWhere is called when production where is exited.
func (listen *MyViewSQLListener) ExitWhere(ctx *parser.WhereContext) {
	//fmt.Println("[Listen]Exit Where")
}

// EnterGroupby is called when production groupby is entered.
func (listen *MyViewSQLListener) EnterGroupby(ctx *parser.GroupbyContext) {
	//fmt.Println("[Listen]Entered Groupby")
	if listen.ignoreGroupBy {
		listen.Mode = GROUP_BY_FIRST
	} else {
		listen.Mode = GROUP_BY_SECOND
	}
}

// ExitGroupby is called when production groupby is exited.
func (listen *MyViewSQLListener) ExitGroupby(ctx *parser.GroupbyContext) {
	//fmt.Println("[Listen]Exit GroupBy")
	if listen.ignoreGroupBy {
		listen.ignoreGroupBy = false
		return
	}

	//First group by has been ignored, now we must take notes on this one
	nameables := ctx.AllNameable()
	listen.Keys = make([]Nameable, len(nameables))
	for i, nameable := range nameables {
		listen.Keys[i] = listen.FullMakeNameableFromContext(nameable)
	}
}

// EnterOrderby is called when production orderby is entered.
func (listen *MyViewSQLListener) EnterOrderby(ctx *parser.OrderbyContext) {
	//fmt.Println("[Listen]Entered OrderBy")

	listen.Mode = ORDER_BY
}

// ExitOrderby is called when production orderby is exited.
func (listen *MyViewSQLListener) ExitOrderby(ctx *parser.OrderbyContext) {
	//fmt.Println("[Listen]Exit OrderBy")

	nameables, sortOrders := ctx.AllNameable(), ctx.AllSortOrder()
	listen.Sorting = make([]Order, len(nameables))

	for i, nameable := range nameables {
		sortOrder := sortOrders[i]
		isDesc := true
		if asc := sortOrder.ASC(); asc != nil {
			isDesc = false
		}
		listen.Sorting[i] = Order{Nameable: listen.FullMakeNameableFromContext(nameable), IsDesc: isDesc}
	}
}

// EnterLimit is called when production limit is entered.
func (listen *MyViewSQLListener) EnterLimit(ctx *parser.LimitContext) {
	//fmt.Println("[Listen]Entered Limit")

	listen.Mode = LIMIT
}

// ExitLimit is called when production limit is exited.
func (listen *MyViewSQLListener) ExitLimit(ctx *parser.LimitContext) {
	//fmt.Println("[Listen]Exit Limit")

	listen.TopSize, _ = strconv.Atoi(ctx.INT().GetText())
}

// EnterView is called when production view is entered.
func (listen *MyViewSQLListener) EnterView(ctx *parser.ViewContext) {
	//fmt.Println("[Listen]Entered View")
	if first := ctx.GetFirstGroupBy(); first != nil {
		//Must ignore the first group by later when processing
		listen.ignoreGroupBy = true
	}
}

// ExitView is called when production view is exited.
func (listen *MyViewSQLListener) ExitView(ctx *parser.ViewContext) {
	//fmt.Println("[Listen]Exit View")
}

// EnterStart is called when production start is entered.
func (listen *MyViewSQLListener) EnterStart(ctx *parser.StartContext) {
	//fmt.Println("[Listen]Entered Start")
}

// ExitStart is called when production start is exited.
func (listen *MyViewSQLListener) ExitStart(ctx *parser.StartContext) {
	//fmt.Println("[Listen]Exit Start")
	fmt.Println("Finished processing view text")
	fmt.Println("ID, Bucket, Mode:", listen.ID, listen.Bucket, listen.Mode)
	fmt.Println("All select:", listen.AllSelect)
	fmt.Println("Primary key:", listen.PrimaryKey)
	fmt.Println("Containers:", listen.Containers)
	fmt.Println("Where IDs, Where Fields, Where Parameters:", listen.WhereIDs,
		listen.WhereFields, listen.WhereParameters)
	fmt.Println("Conditions:", listen.Conditions)
	fmt.Println("Keys (2nd group by):", listen.Keys)
	fmt.Println("Sorting:", listen.Sorting)
	fmt.Println("TopSize:", listen.TopSize)
}

//Helpers

//Makes nameable if it doesn't already exist; otherwise returns existing instance
func (listen *MyViewSQLListener) FullMakeNameableFromContext(ctx parser.INameableContext) Nameable {
	//ExitNameable already takes care of storing the IDs. So here we only need the "math" part
	if name := ctx.Name(); name != nil {
		text := name.STRING().GetText()
		nameable, has := listen.AllNameables[text]
		if !has {
			nameable = &Name{Name: text}
			listen.AllNameables[text] = nameable
		}
		return nameable
	}
	if parameter := ctx.Parameter(); parameter != nil {
		text := parameter.Name().STRING().GetText()
		nameable := listen.AllNameables[text] //Parameter already exists in both AllNameables and AllParameters
		return nameable
		/*
			nameable, has := listen.AllNameables[text]
			if !has {
				nameable = &Parameter{Name: text}
				listen.AllNameables[text] = nameable
			}
			return nameable
		*/
	}
	field := ctx.Field()
	first, second := field.Name(0).STRING().GetText(), GetDateFieldType(field.Name(1).STRING().GetText())
	tmp := Field{Name: first, FieldType: second}
	nameable, has := listen.AllNameables[tmp.GetIdentifier()]
	if !has {
		nameable = &tmp
		listen.AllNameables[tmp.GetIdentifier()] = nameable
	}
	return nameable
}

func (listen *MyViewSQLListener) HandleNameableInSelect(ctx *parser.NameableContext) {
	//Find out which one is defined
	if ctx.Name() != nil {
		listen.SelectIDs[ctx.Name().STRING().GetText()] = struct{}{}
	} else if ctx.Field() != nil {
		field := ctx.Field()
		tmpF := Field{Name: field.Name(0).STRING().GetText(), FieldType: GetDateFieldType(field.Name(1).STRING().GetText())}
		listen.SelectFields[tmpF.GetIdentifier()] = tmpF.FieldType
	} else { //Parameter
		listen.SelectParameters[ctx.Parameter().Name().STRING().GetText()] = struct{}{}
	}
}

func (listen *MyViewSQLListener) HandleNameableInWhere(ctx *parser.NameableContext) {
	//Find out which one is defined
	if ctx.Name() != nil {
		listen.WhereIDs[ctx.Name().STRING().GetText()] = struct{}{}
	} else if ctx.Field() != nil {
		field := ctx.Field()
		listen.WhereFields[field.Name(0).STRING().GetText()] = GetDateFieldType(field.Name(1).STRING().GetText())
	} else { //Parameter
		listen.WhereParameters[ctx.Parameter().Name().STRING().GetText()] = struct{}{}
	}
}

func (listen *MyViewSQLListener) makeConstantFromContext(ctx parser.IConstantContext) Constant {
	/*if stringValue := ctx.STRING(); stringValue != nil {
	return &StringConstant{Value: stringValue.GetText()}*/
	if allStrings := ctx.AllSTRING(); len(allStrings) > 0 {
		fullString := allStrings[0].GetText()
		for _, str := range allStrings[1:] {
			fullString += " " + str.GetText()
		}
		return &StringConstant{Value: fullString}
	} else if number := ctx.INT(); number != nil {
		value, _ := strconv.ParseInt(number.GetText(), 10, 64)
		return &IntConstant{Value: int(value)}
	} else if number := ctx.FLOAT(); number != nil {
		text := number.GetText()
		value, _ := strconv.ParseFloat(text, 64)
		dotPos := strings.Index(text, ".")
		return &FloatConstant{Value: value, Precision: len(text) - dotPos - 1}
	} else { //Date
		dateText := ctx.DATE().GetText()
		date := &DateConstant{}
		date.FromString(dateText)
		return date
	}
}

// EnterCalc is called when production calc is entered.
func (listen *MyViewSQLListener) EnterCalc(ctx *parser.CalcContext) {
	//fmt.Println("[Listen]Entered Calc")
}

// ExitCalc is called when production calc is exited.
func (listen *MyViewSQLListener) ExitCalc(ctx *parser.CalcContext) {
	//fmt.Println("[Listen]Exit Calc")
	listen.NEntriesInSelect++
	var field SelectFields
	if nameable := ctx.Nameable(); nameable != nil {
		//listen.SelectByAppearenceOrder = append(listen.SelectByAppearenceOrder, nameable.Name().GetText())
		field = listen.SelectNameable[nameable.Name().GetText()]
	} else if asClause := ctx.AsClause(); asClause != nil {
		//listen.SelectByAppearenceOrder = append(listen.SelectByAppearenceOrder, asClause.Name().GetText())
		field = listen.SelectAsClause[asClause.Name().GetText()]
	} else if aggr := ctx.Aggregation(); aggr != nil {
		//listen.SelectByAppearenceOrder = append(listen.SelectByAppearenceOrder, aggr.Name().GetText())
		field = listen.SelectAsAggregation[aggr.Name().GetText()]
	} else if primaryKey := ctx.Key(); primaryKey != nil {
		//listen.SelectByAppearenceOrder = append(listen.SelectByAppearenceOrder, primaryKey.Nameable().Name().GetText())
		field = listen.PrimaryKey
	} else {
		//listen.SelectByAppearenceOrder = append(listen.SelectByAppearenceOrder, ctx.Count().Name().GetText())
		field = listen.SelectAsCount[ctx.Count().Name().GetText()]
	}
	listen.SelectByAppearenceOrder = append(listen.SelectByAppearenceOrder, field)
	listen.AllSelect[field.GetIdentifier()] = field
}

//EnterContinuousRange is called when production continuousRange is entered.
func (listen *MyViewSQLListener) EnterContinuousRange(ctx *parser.ContinuousRangeContext) {

}

//ExitContinuousRange is called when production continuousRange is exited.
func (listen *MyViewSQLListener) ExitContinuousRange(ctx *parser.ContinuousRangeContext) {

}

//EnterSparseRange is called when production sparseRange is entered.
func (listen *MyViewSQLListener) EnterSparseRange(ctx *parser.SparseRangeContext) {

}

//ExitSparseRange is called when production sparseRange is exited.
func (listen *MyViewSQLListener) ExitSparseRange(ctx *parser.SparseRangeContext) {

}

//EnterRange is called when production range is entered.
func (listen *MyViewSQLListener) EnterRange(ctx *parser.RangeContext) {

}

//ExitRange is called when production range is exited.
func (listen *MyViewSQLListener) ExitRange(ctx *parser.RangeContext) {

}

//EnterWith is called when production with is entered.
func (listen *MyViewSQLListener) EnterWith(ctx *parser.WithContext) {

}

//ExitWith is called when production with is exited.
func (listen *MyViewSQLListener) ExitWith(ctx *parser.WithContext) {
	paramsNames, ranges := ctx.AllSTRING(), ctx.AllRange_()
	for i, name := range paramsNames {
		currRange := ranges[i]
		var param Parameter
		inParam := &InParameter{Name: name.GetText()}
		if cRange := currRange.ContinuousRange(); cRange != nil {
			leftInclusive, rightInclusive := true, true
			if cRange.GetLeft().GetText() == "]" {
				leftInclusive = false
			}
			if cRange.GetRight().GetText() == "[" {
				rightInclusive = false
			}
			minC, maxC := cRange.Constant(0), cRange.Constant(1)
			if minI := minC.INT(); minI != nil {
				maxI := maxC.INT()
				min, _ := strconv.Atoi(minI.GetText())
				max, _ := strconv.Atoi(maxI.GetText())
				if !leftInclusive {
					min++
				}
				if !rightInclusive {
					max++
				}
				param = &RangeParameterInt{InParameter: inParam, MinValue: min, MaxValue: max}
			} else if minF := minC.FLOAT(); minI != nil {
				maxF := maxC.FLOAT()
				textMin, textMax := minF.GetText(), maxF.GetText()
				min, _ := strconv.ParseFloat(textMin, 64)
				max, _ := strconv.ParseFloat(textMax, 64)
				dotPosMin, dotPosMax := strings.Index(textMin, "."), strings.Index(textMax, ".")
				floatPMin, floatPMax := len(textMin)-dotPosMin-1, len(textMax)-dotPosMax-1
				floatParam := RangeParameterFloat{InParameter: inParam}
				if floatPMax > floatPMin {
					floatParam.FloatPrecision = floatPMax
				} else {
					floatParam.FloatPrecision = floatPMin
				}
				if !leftInclusive {
					min += floatParam.GetFloatStep()
				}
				if !rightInclusive {
					max -= floatParam.GetFloatStep()
				}
				floatParam.MinValue, floatParam.MaxValue = min, max
				param = &floatParam
			} else if minD := minC.DATE(); minD != nil {
				maxD := maxC.DATE()
				min, max := Date{}.FromString(minD.GetText()), Date{}.FromString(maxD.GetText())
				if !leftInclusive {
					calcDate := time.Date(min.Year, time.Month(min.Month), min.Day+1, 0, 0, 0, 0, time.UTC)
					min = Date{Day: calcDate.Day(), Month: int(calcDate.Month()), Year: calcDate.Year()}
				}
				if !rightInclusive {
					calcDate := time.Date(min.Year, time.Month(min.Month), min.Day-1, 0, 0, 0, 0, time.UTC)
					max = Date{Day: calcDate.Day(), Month: int(calcDate.Month()), Year: calcDate.Year()}
				}
				param = &RangeParameterDate{InParameter: inParam, MinValue: min, MaxValue: max}
			}
			//Note: Decided to not support Strings for continuous range. Sorry.
		} else {
			sRange := currRange.SparseRange()
			values := sRange.AllConstant()
			var typeC DataType
			firstV := values[0]
			//var floatP int
			//if stringV := firstV.STRING(); stringV != nil {
			if stringV := firstV.AllSTRING(); len(stringV) > 0 {
				typeC = STRING_TYPE
			} else if number := firstV.INT(); number != nil {
				typeC = INT_TYPE
			} else if number := firstV.FLOAT(); number != nil {
				typeC = FLOAT_TYPE
				//text := number.GetText()
				//dotPos := strings.Index(text, ".")
				//floatP = len(text) - dotPos - 1
			} else {
				typeC = DATE_TYPE
			}
			switch typeC {
			case INT_TYPE:
				valuesMap := make(map[int]struct{})
				for _, value := range values {
					intV, _ := strconv.Atoi(value.INT().GetText())
					valuesMap[intV] = struct{}{}
				}
				sparseP := &SparseParameterInt{InParameter: inParam, Values: valuesMap}
				sparseP.CalculateMinMax()
				param = sparseP
			case FLOAT_TYPE:
				valuesMap := make(map[float64]struct{})
				for _, value := range values {
					floatV, _ := strconv.ParseFloat(value.INT().GetText(), 64)
					valuesMap[floatV] = struct{}{}
				}
				sparseP := &SparseParameterFloat{InParameter: inParam, Values: valuesMap}
				sparseP.CalculateMinMax()
				param = sparseP
			case STRING_TYPE:
				valuesMap := make(map[string]struct{})
				for _, value := range values {
					//valuesMap[value.STRING().GetText()] = struct{}{}
					allStrings := value.AllSTRING()
					fullS := allStrings[0].GetText()
					for _, str := range allStrings[1:] {
						fullS += " " + str.GetText()
					}
					valuesMap[fullS] = struct{}{}
				}
				param = &SparseParameterString{InParameter: inParam, Values: valuesMap}
			case DATE_TYPE:
				valuesMap := make(map[Date]struct{})
				for _, value := range values {
					date := Date{}
					date = date.FromString(value.DATE().GetText())
					valuesMap[date] = struct{}{}
				}
				sparseP := &SparseParameterDate{InParameter: inParam, Values: valuesMap}
				sparseP.CalculateMinMax()
				param = sparseP
			}
		}
		listen.AllNameables[param.GetIdentifier()] = param
		listen.AllParameters[param.GetIdentifier()] = param
	}
}

//EnterKey is called when production key is entered.
func (listen *MyViewSQLListener) EnterKey(ctx *parser.KeyContext) {

}

//ExitKey is called when production key is exited.
func (listen *MyViewSQLListener) ExitKey(ctx *parser.KeyContext) {
	listen.PrimaryKey = &Name{Name: ctx.Nameable().Name().GetText()}
}

//ANYTHING BELOW HERE SHOULD NOT BE NEEDED TO BE IMPLEMENTED

// EnterValue is called when production value is entered.
func (listen *MyViewSQLListener) EnterParentheses(ctx *parser.ParenthesesContext) {}

// ExitValue is called when production value is exited.
func (listen *MyViewSQLListener) ExitParentheses(ctx *parser.ParenthesesContext) {}

// EnterAggrFunc is called when production aggrFunc is entered.
func (listen *MyViewSQLListener) EnterAggrFunc(ctx *parser.AggrFuncContext) {}

// ExitAggrFunc is called when production aggrFunc is exited.
func (listen *MyViewSQLListener) ExitAggrFunc(ctx *parser.AggrFuncContext) {}

// EnterKey is called when production key is entered.
func (listen *MyViewSQLListener) EnterCheck(ctx *parser.CheckContext) {}

// ExitKey is called when production key is exited.
func (listen *MyViewSQLListener) ExitCheck(ctx *parser.CheckContext) {}

// EnterForeignkey is called when production foreignkey is entered.
func (listen *MyViewSQLListener) EnterForeignkey(ctx *parser.ForeignkeyContext) {}

// ExitForeignkey is called when production foreignkey is exited.
func (listen *MyViewSQLListener) ExitForeignkey(ctx *parser.ForeignkeyContext) {}

// EnterPrimarykey is called when production primarykey is entered.
func (listen *MyViewSQLListener) EnterPrimarykey(ctx *parser.PrimarykeyContext) {}

// ExitPrimarykey is called when production primarykey is exited.
func (listen *MyViewSQLListener) ExitPrimarykey(ctx *parser.PrimarykeyContext) {}

// EnterConstraint is called when production constraint is entered.
func (listen *MyViewSQLListener) EnterConstraint(ctx *parser.ConstraintContext) {}

// ExitConstraint is called when production constraint is exited.
func (listen *MyViewSQLListener) ExitConstraint(ctx *parser.ConstraintContext) {}

// EnterColumns is called when production columns is entered.
func (listen *MyViewSQLListener) EnterColumns(ctx *parser.ColumnsContext) {}

// ExitColumns is called when production columns is exited.
func (listen *MyViewSQLListener) ExitColumns(ctx *parser.ColumnsContext) {}

// EnterCreatetable is called when production createtable is entered.
func (listen *MyViewSQLListener) EnterCreatetable(ctx *parser.CreatetableContext) {}

// ExitCreatetable is called when production createtable is exited.
func (listen *MyViewSQLListener) ExitCreatetable(ctx *parser.CreatetableContext) {}

// EnterCreateindex is called when production createindex is entered.
func (listen *MyViewSQLListener) EnterCreateindex(ctx *parser.CreateindexContext) {}

// ExitCreateindex is called when production createindex is exited.
func (listen *MyViewSQLListener) ExitCreateindex(ctx *parser.CreateindexContext) {}

// EnterDrop is called when production drop is entered.
func (listen *MyViewSQLListener) EnterDrop(ctx *parser.DropContext) {}

// ExitDrop is called when production drop is exited.
func (listen *MyViewSQLListener) ExitDrop(ctx *parser.DropContext) {}

// EnterDelete is called when production delete is entered.
func (listen *MyViewSQLListener) EnterDelete(ctx *parser.DeleteContext) {}

// ExitDelete is called when production delete is exited.
func (listen *MyViewSQLListener) ExitDelete(ctx *parser.DeleteContext) {}

// EnterSet is called when production set is entered.
func (listen *MyViewSQLListener) EnterSet(ctx *parser.SetContext) {}

// ExitSet is called when production set is exited.
func (listen *MyViewSQLListener) ExitSet(ctx *parser.SetContext) {}

// EnterUpdate is called when production update is entered.
func (listen *MyViewSQLListener) EnterUpdate(ctx *parser.UpdateContext) {}

// ExitUpdate is called when production update is exited.
func (listen *MyViewSQLListener) ExitUpdate(ctx *parser.UpdateContext) {}

// EnterValues is called when production values is entered.
func (listen *MyViewSQLListener) EnterValues(ctx *parser.ValuesContext) {}

// ExitValues is called when production values is exited.
func (listen *MyViewSQLListener) ExitValues(ctx *parser.ValuesContext) {}

// EnterColumnNames is called when production columnNames is entered.
func (listen *MyViewSQLListener) EnterColumnNames(ctx *parser.ColumnNamesContext) {}

// ExitColumnNames is called when production columnNames is exited.
func (listen *MyViewSQLListener) ExitColumnNames(ctx *parser.ColumnNamesContext) {}

// EnterQuery is called when production query is entered.
func (listen *MyViewSQLListener) EnterQuery(ctx *parser.QueryContext) {}

// ExitQuery is called when production query is exited.
func (listen *MyViewSQLListener) ExitQuery(ctx *parser.QueryContext) {}

// EnterStatement is called when production statement is entered.
func (listen *MyViewSQLListener) EnterStatement(ctx *parser.StatementContext) {}

// ExitStatement is called when production statement is exited.
func (listen *MyViewSQLListener) ExitStatement(ctx *parser.StatementContext) {}

// EnterInsert is called when production insert is entered.
func (listen *MyViewSQLListener) EnterInsert(ctx *parser.InsertContext) {}

// ExitInsert is called when production insert is exited.
func (listen *MyViewSQLListener) ExitInsert(ctx *parser.InsertContext) {}
