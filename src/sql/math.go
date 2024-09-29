package sql

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"sqlToKeyValue/src/parser"
)

//TODO: If I forced a calculation of HasParameter() and then cached the result
//it could save computations potencially... I will see after the algorithm is finished.

const (
	NUMBER_TYPE   = DataType(1)
	INT_TYPE      = DataType(2)
	FLOAT_TYPE    = DataType(3)
	STRING_TYPE   = DataType(4)
	DATE_TYPE     = DataType(5)
	VARIABLE_TYPE = DataType(6)

	DATE_FIELD_DAY     = DateFieldType(1)
	DATE_FIELD_MONTH   = DateFieldType(2)
	DATE_FIELD_YEAR    = DateFieldType(3)
	DATE_FIELD_QUARTER = DateFieldType(4)
)

type Math struct {
	Stack CalculableStack
}

func (m Math) GetCalculable() Calculable {
	return m.Stack.Pool()
}

func (m Math) ToString() string {
	var sb strings.Builder
	m.GetCalculable().ToPrintString(&sb)
	return sb.String()
}

func CreateMath() Math {
	return Math{Stack: make(CalculableStack, 0, 10)}
}

func CreateMathFromCalculable(calc Calculable) Math {
	toReturn := Math{Stack: make(CalculableStack, 1, 1)}
	toReturn.Stack[0] = calc
	return toReturn
}

type Calculable interface {
	GetType() DataType
	SetType(dataType DataType) //Used to set the Nameable's type.
	//ReplaceValues(objValues map[string]string) string
	IsConstant() bool
	//Pre: Types should already be calculated previously.
	RegisterFields(fields map[string]DataType)
	CalculateValue(objValues map[string]interface{}) interface{}
	HasParameter() bool                             //For Calculables with other entities inside, it only checks the first level.
	ReplaceParams(equalP map[string]EqualParameter) //Replaces all parameters with the Calculable.
	ToPrintString(sb *strings.Builder)              //Converts to a string representation for printing/debugging/visualization purposes
	ToString() string                               //Same as above but without sb
}

type Nameable interface {
	Calculable
	SelectFields
	//GetIdentifier() string
}

/*
type Value interface {
	Calculable
} */

type Add struct {
	Type  DataType
	Left  Calculable
	Right Calculable
}

type Sub struct {
	Type  DataType
	Left  Calculable
	Right Calculable
}

type Mult struct {
	Type  DataType
	Left  Calculable
	Right Calculable
}

type Div struct {
	Type  DataType
	Left  Calculable
	Right Calculable
}

type Minus struct {
	Type        DataType
	ToCalculate Calculable
}

type Name struct {
	Type DataType
	Name string
}

/*
type Parameter struct {
	Type DataType
	Name string
}

type EqualParameter struct {
	*Parameter
	CalcFrom Calculable
}*/

//TODO: Consider doing a struct per DataType for this one.
/*type InequalParameter struct {
	*Parameter
	MinValue       interface{} //Inclusive
	MaxValue       interface{} //Inclusive
	FloatPrecision int
}

type Parameter interface {
	Nameable
	IsInRange(value interface{}) bool
}

type EqualParameter struct {
	Name     string
	CalcFrom Calculable
}
*/

type Parameter interface {
	Nameable
	IsInRange(value interface{}) bool
	IsInInequalityRange(value interface{}, op CondType) bool
	GetIneqKeys(conds []Condition, objValues map[string]interface{}) (keys []string)
}

//Helper struct that has things common to all kinds of Parameters
type InParameter struct {
	Type           DataType
	Name           string
	FloatPrecision int
}

type RangeParameterInt struct {
	*InParameter
	MinValue, MaxValue int //Inclusive
}

type RangeParameterFloat struct {
	*InParameter
	MinValue, MaxValue float64 //Inclusive
}

type RangeParameterDate struct {
	*InParameter
	MinValue, MaxValue Date //Inclusive
}

/*
type RangeParameter struct {
	*InParameter
	MinValue interface{} //Inclusive
	MaxValue interface{} //Inclusive
}
*/

type SparseParameterInt struct {
	*InParameter
	Values             map[int]struct{}
	MinValue, MaxValue int
}

type SparseParameterFloat struct {
	*InParameter
	Values             map[float64]struct{}
	MinValue, MaxValue float64
}

type SparseParameterDate struct {
	*InParameter
	Values             map[Date]struct{}
	MinValue, MaxValue Date
}

type SparseParameterString struct {
	*InParameter
	Values map[string]struct{}
}

//Note: >, <, >= and <= make no sense with SparseParameters.
/*
type SparseParameter struct {
	*InParameter
	Values map[interface{}]struct{}
}
*/

//Contains either a RangeParameter or SparseParameter
type EqualParameter struct {
	Parameter
	CalcFrom Calculable
}

type Field struct { //Fields of Date: thus type is always an INT
	Name      string
	FieldType DateFieldType
}

type Constant interface {
	Calculable
}

type StringConstant struct {
	Value string
}

type IntConstant struct {
	Value int
}

type FloatConstant struct {
	Value     float64
	Precision int //Number of decimal cases
}

type DateConstant struct {
	Date
}

type Date struct {
	Day   int
	Month int
	Year  int
}

type DateFieldType int

func (dType DateFieldType) ToString() string {
	switch dType {
	case DATE_FIELD_DAY:
		return "DAY"
	case DATE_FIELD_MONTH:
		return "MONTH"
	case DATE_FIELD_YEAR:
		return "YEAR"
	case DATE_FIELD_QUARTER:
		return "QUARTER"
	}
	return "UNKNOWN"
}

type DataType int

func (dType DataType) ToString() string {
	switch dType {
	case NUMBER_TYPE:
		return "NUMBER"
	case INT_TYPE:
		return "INT"
	case FLOAT_TYPE:
		return "FLOAT"
	case STRING_TYPE:
		return "STRING"
	case DATE_TYPE:
		return "DATE"
	case VARIABLE_TYPE:
		return "VARIABLE"
	}
	return "UNKNOWN"
}

func (add *Add) GetType() DataType {
	if add.Type == 0 {
		add.Type = GetTypeOfTwoAux(add.Left, add.Right)
	}
	return add.Type
}

func (sub *Sub) GetType() DataType {
	if sub.Type == 0 {
		sub.Type = GetTypeOfTwoAux(sub.Left, sub.Right)
	}
	return sub.Type
}

func (mult *Mult) GetType() DataType {
	if mult.Type == 0 {
		mult.Type = GetTypeOfTwoAux(mult.Left, mult.Right)
		if mult.Type == VARIABLE_TYPE { //Mult is always of numbers
			mult.Type = NUMBER_TYPE
			mult.Left.(Nameable).SetType(NUMBER_TYPE)
			mult.Right.(Nameable).SetType(NUMBER_TYPE)
		}
	}
	return mult.Type
}

func (div *Div) GetType() DataType {
	if div.Type == 0 {
		div.Type = GetTypeOfTwoAux(div.Left, div.Right)
		if div.Type == VARIABLE_TYPE { //Div is always of numbers
			div.Type = NUMBER_TYPE
			div.Left.(Nameable).SetType(NUMBER_TYPE)
			div.Right.(Nameable).SetType(NUMBER_TYPE)
		}
	}
	return div.Type
}

func (minus *Minus) GetType() DataType {
	if minus.Type == 0 {
		minus.Type = minus.ToCalculate.GetType()
		if minus.Type == VARIABLE_TYPE { //Minus is always of numbers
			minus.Type = NUMBER_TYPE
			minus.ToCalculate.(Nameable).SetType(NUMBER_TYPE)
		}
	}
	return minus.Type
}

func GetTypeOfTwoAux(left, right Calculable) DataType {
	leftT, rightT := left.GetType(), right.GetType()
	if leftT == rightT {
		return leftT
	}
	if leftT == VARIABLE_TYPE { //rightT is constant - OK, update lefT's type
		if rightT == INT_TYPE || rightT == FLOAT_TYPE {
			//We only know the other one is a number
			left.(Nameable).SetType(NUMBER_TYPE)
			return NUMBER_TYPE
		}
		left.(Nameable).SetType(rightT)
		return rightT
	}
	if rightT == VARIABLE_TYPE {
		if leftT == INT_TYPE || leftT == FLOAT_TYPE {
			//We only know the other one is a number
			right.(Nameable).SetType(NUMBER_TYPE)
			return NUMBER_TYPE
		}
		right.(Nameable).SetType(leftT)
		return leftT
	}
	//this is actually not OK - opportunity for an exception/some error.
	//(both are constants but different types)
	return leftT
}

func (sConst *StringConstant) GetType() DataType { return STRING_TYPE }
func (iConst *IntConstant) GetType() DataType    { return INT_TYPE }
func (fConst *FloatConstant) GetType() DataType  { return FLOAT_TYPE }
func (dConst *DateConstant) GetType() DataType   { return DATE_TYPE }

func (name *Name) GetType() DataType {
	if name.Type == 0 {
		return VARIABLE_TYPE
	}
	return name.Type
}

/*func (parameter *Parameter) GetType() DataType {
	if parameter.Type == 0 {
		return VARIABLE_TYPE
	}
	return parameter.Type
}*/
/*func (inP *InParameter) GetType() DataType {
	if inP.Type == 0 {
		return VARIABLE_TYPE
	}
	return inP.Type
}*/

func (rangeP *RangeParameterInt) GetType() DataType     { return INT_TYPE }
func (rangeP *RangeParameterFloat) GetType() DataType   { return FLOAT_TYPE }
func (rangeP *RangeParameterDate) GetType() DataType    { return DATE_TYPE }
func (rangeP *SparseParameterInt) GetType() DataType    { return INT_TYPE }
func (rangeP *SparseParameterFloat) GetType() DataType  { return FLOAT_TYPE }
func (rangeP *SparseParameterDate) GetType() DataType   { return DATE_TYPE }
func (rangeP *SparseParameterString) GetType() DataType { return STRING_TYPE }

func (field *Field) GetType() DataType { return DATE_TYPE }

//This is only called when this is inside an aggregation function.
//Pre: add.Type == VARIABLE_TYPE
func (add *Add) SetType(dataType DataType) {
	add.Left.SetType(dataType)
	add.Right.SetType(dataType)
	add.Type = dataType
}

//This is only called when this is inside an aggregation function.
//Pre: sub.Type == VARIABLE_TYPE
func (sub *Sub) SetType(dataType DataType) {
	sub.Left.SetType(dataType)
	sub.Right.SetType(dataType)
	sub.Type = dataType
}

//No effect: mult, div, minus are already known to be at least a NUMBER_TYPE
func (mult *Mult) SetType(dataType DataType)   {}
func (div *Div) SetType(dataType DataType)     {}
func (minus *Minus) SetType(dataType DataType) {}

//No effect: constant
func (sConst *StringConstant) SetType(dataType DataType) {}
func (iConst *IntConstant) SetType(dataType DataType)    {}
func (fConst *FloatConstant) SetType(dataType DataType)  {}
func (dConst *DateConstant) SetType(dataType DataType)   {}

func (name *Name) SetType(dataType DataType) { name.Type = dataType }

//func (parameter *Parameter) SetType(dataType DataType) { parameter.Type = dataType }
func (inP *InParameter) SetType(dataType DataType) { inP.Type = dataType }
func (field *Field) SetType(dataType DataType)     {}

func (add *Add) IsConstant() bool               { return false }
func (sub *Sub) IsConstant() bool               { return false }
func (mult *Mult) IsConstant() bool             { return false }
func (div *Div) IsConstant() bool               { return false }
func (minus *Minus) IsConstant() bool           { return false }
func (sConst *StringConstant) IsConstant() bool { return true }
func (iConst *IntConstant) IsConstant() bool    { return true }
func (fConst *FloatConstant) IsConstant() bool  { return true }
func (dConst *DateConstant) IsConstant() bool   { return true }
func (name *Name) IsConstant() bool             { return false }

//func (parameter *Parameter) IsConstant() bool   { return false }
func (inP *InParameter) IsConstant() bool { return false }
func (field *Field) IsConstant() bool     { return false }

func (add *Add) RegisterFields(fields map[string]DataType) {
	add.Left.RegisterFields(fields)
	add.Right.RegisterFields(fields)
}

func (sub *Sub) RegisterFields(fields map[string]DataType) {
	sub.Left.RegisterFields(fields)
	sub.Right.RegisterFields(fields)
}

func (mult *Mult) RegisterFields(fields map[string]DataType) {
	mult.Left.RegisterFields(fields)
	mult.Right.RegisterFields(fields)
}

func (div *Div) RegisterFields(fields map[string]DataType) {
	div.Left.RegisterFields(fields)
	div.Right.RegisterFields(fields)
}

func (minus *Minus) RegisterFields(fields map[string]DataType) {
	minus.ToCalculate.RegisterFields(fields)
}

func (sConst *StringConstant) RegisterFields(fields map[string]DataType) {}

func (iConst *IntConstant) RegisterFields(fields map[string]DataType) {}

func (fConst *FloatConstant) RegisterFields(fields map[string]DataType) {}

func (dConst *DateConstant) RegisterFields(fields map[string]DataType) {}

func (name *Name) RegisterFields(fields map[string]DataType) {
	fields[name.Name] = name.GetType()
}

/*func (parameter *Parameter) RegisterFields(fields map[string]DataType) {
	//fields[parameter.Name] = parameter.GetType()
}*/
func (inP *InParameter) RegisterFields(fields map[string]DataType) {
	//fields[inP.Name] = inP.GetType()
}

func (field *Field) RegisterFields(fields map[string]DataType) {
	fields[field.Name] = DATE_TYPE
}

func (add *Add) CalculateValue(objValues map[string]interface{}) interface{} {
	leftC, rightC := add.Left.CalculateValue(objValues), add.Right.CalculateValue(objValues)
	switch leftTyped := leftC.(type) {
	case int:
		switch rightTyped := rightC.(type) {
		case int:
			return leftTyped + rightTyped
		case float64:
			return float64(leftTyped) + rightTyped
		}
	case float64:
		switch rightTyped := rightC.(type) {
		case float64:
			return leftTyped + rightTyped
		case int:
			return leftTyped + float64(rightTyped)
		}
	case Date:
		calcDate := time.Date(leftTyped.Year, time.Month(leftTyped.Month), leftTyped.Day+rightC.(int), 0, 0, 0, 0, time.UTC)
		return Date{Day: calcDate.Day(), Month: int(calcDate.Month()), Year: calcDate.Year()}
	case string:
		return leftTyped + rightC.(string)
	}
	return nil
}

func (sub *Sub) CalculateValue(objValues map[string]interface{}) interface{} {
	leftC, rightC := sub.Left.CalculateValue(objValues), sub.Right.CalculateValue(objValues)
	switch leftTyped := leftC.(type) {
	case int:
		switch rightTyped := rightC.(type) {
		case int:
			return leftTyped - rightTyped
		case float64:
			return float64(leftTyped) - rightTyped
		}
	case float64:
		switch rightTyped := rightC.(type) {
		case float64:
			return leftTyped - rightTyped
		case int:
			return leftTyped - float64(rightTyped)
		}
	case Date:
		calcDate := time.Date(leftTyped.Year, time.Month(leftTyped.Month), leftTyped.Day-rightC.(int), 0, 0, 0, 0, time.UTC)
		return Date{Day: calcDate.Day(), Month: int(calcDate.Month()), Year: calcDate.Year()}
	}
	return nil
}

func (mult *Mult) CalculateValue(objValues map[string]interface{}) interface{} {
	leftC, rightC := mult.Left.CalculateValue(objValues), mult.Right.CalculateValue(objValues)
	switch leftTyped := leftC.(type) {
	case int:
		switch rightTyped := rightC.(type) {
		case int:
			return leftTyped * rightTyped
		case float64:
			return float64(leftTyped) * rightTyped
		}
	case float64:
		switch rightTyped := rightC.(type) {
		case float64:
			return leftTyped * rightTyped
		case int:
			return leftTyped * float64(rightTyped)
		}
	}
	return nil
}

func (div *Div) CalculateValue(objValues map[string]interface{}) interface{} {
	leftC, rightC := div.Left.CalculateValue(objValues), div.Right.CalculateValue(objValues)
	switch leftTyped := leftC.(type) {
	case int:
		switch rightTyped := rightC.(type) {
		case int:
			return leftTyped / rightTyped
		case float64:
			return float64(leftTyped) / rightTyped
		}
	case float64:
		switch rightTyped := rightC.(type) {
		case float64:
			return leftTyped / rightTyped
		case int:
			return leftTyped / float64(rightTyped)
		}
	}
	return nil
}

func (minus *Minus) CalculateValue(objValues map[string]interface{}) interface{} {
	calculated := minus.CalculateValue(objValues)
	switch typed := calculated.(type) {
	case int:
		return -typed
	case float64:
		return -typed
	}
	return calculated
}

func (sConst *StringConstant) CalculateValue(objValues map[string]interface{}) interface{} {
	return sConst.Value
}

func (iConst *IntConstant) CalculateValue(objValues map[string]interface{}) interface{} {
	return iConst.Value
}

func (fConst *FloatConstant) CalculateValue(objValues map[string]interface{}) interface{} {
	return fConst.Value
}

func (dConst *DateConstant) CalculateValue(objValues map[string]interface{}) interface{} {
	return dConst.Date
}

func (name *Name) CalculateValue(objValues map[string]interface{}) interface{} {
	return objValues[name.Name]
}

/*func (parameter *Parameter) CalculateValue(objValues map[string]interface{}) interface{} {
	//Never happens
	return nil
}*/

func (inP *InParameter) CalculateValue(objValues map[string]interface{}) interface{} {
	//Never happens
	return nil
}

func (parameter *EqualParameter) CalculateValue(objValues map[string]interface{}) interface{} {
	return parameter.CalcFrom.CalculateValue(objValues)
}

/*
func (ineqParameter *InequalParameter) CalculateValue(objValues map[string]interface{}) interface{} {

}*/

func (field *Field) CalculateValue(objValues map[string]interface{}) interface{} {
	fmt.Println(*field)
	fmt.Println(objValues[field.Name])
	date := objValues[field.Name].(Date)
	switch field.FieldType {
	case DATE_FIELD_DAY:
		return date.GetDay()
	case DATE_FIELD_MONTH:
		return date.GetMonth()
	case DATE_FIELD_YEAR:
		return date.GetYear()
	case DATE_FIELD_QUARTER:
		return date.GetQuarter()
	}
	return nil
}

func (add *Add) HasParameter() bool {
	_, ok := add.Left.(Parameter)
	if !ok {
		_, ok = add.Right.(Parameter)
	}
	return ok
}

func (sub *Sub) HasParameter() bool {
	_, ok := sub.Left.(Parameter)
	if !ok {
		_, ok = sub.Right.(Parameter)
	}
	return ok
}

func (mult *Mult) HasParameter() bool {
	_, ok := mult.Left.(Parameter)
	if !ok {
		_, ok = mult.Right.(Parameter)
	}
	return ok
}

func (div *Div) HasParameter() bool {
	_, ok := div.Left.(Parameter)
	if !ok {
		_, ok = div.Right.(Parameter)
	}
	return ok
}

func (minus *Minus) HasParameter() bool {
	_, ok := minus.ToCalculate.(Parameter)
	return ok
}

func (sConst *StringConstant) HasParameter() bool { return false }
func (iConst *IntConstant) HasParameter() bool    { return false }
func (fConst *FloatConstant) HasParameter() bool  { return false }
func (dConst *DateConstant) HasParameter() bool   { return false }
func (name *Name) HasParameter() bool             { return false }

//func (parameter *Parameter) HasParameter() bool   { return true }
func (inP *InParameter) HasParameter() bool { return true }
func (field *Field) HasParameter() bool     { return false }

func (add *Add) ReplaceParams(equalP map[string]EqualParameter) {
	add.Left, add.Right = duoReplaceParamsHelper(add.Left, add.Right, equalP)
}

func (sub *Sub) ReplaceParams(equalP map[string]EqualParameter) {
	sub.Left, sub.Right = duoReplaceParamsHelper(sub.Left, sub.Right, equalP)
}

func (mult *Mult) ReplaceParams(equalP map[string]EqualParameter) {
	mult.Left, mult.Right = duoReplaceParamsHelper(mult.Left, mult.Right, equalP)
}

func (div *Div) ReplaceParams(equalP map[string]EqualParameter) {
	div.Left, div.Right = duoReplaceParamsHelper(div.Left, div.Right, equalP)
}

func duoReplaceParamsHelper(left, right Calculable, equalP map[string]EqualParameter) (newL, newR Calculable) {
	if param, ok := left.(Parameter); ok {
		if equal, has := equalP[param.GetIdentifier()]; has {
			left = equal.CalcFrom
		}
	} else {
		left.ReplaceParams(equalP)
	}
	if param, ok := right.(Parameter); ok {
		if equal, has := equalP[param.GetIdentifier()]; has {
			right = equal.CalcFrom
		}
	} else {
		right.ReplaceParams(equalP)
	}
	return left, right
}

func (minus *Minus) ReplaceParams(equalP map[string]EqualParameter) {
	if param, ok := minus.ToCalculate.(Parameter); ok {
		if equal, has := equalP[param.GetIdentifier()]; has {
			minus.ToCalculate = equal.CalcFrom
		}
	} else {
		minus.ToCalculate.ReplaceParams(equalP)
	}
}

func (sConst *StringConstant) ReplaceParams(equalP map[string]EqualParameter) {}
func (sConst *IntConstant) ReplaceParams(equalP map[string]EqualParameter)    {}
func (sConst *FloatConstant) ReplaceParams(equalP map[string]EqualParameter)  {}
func (sConst *DateConstant) ReplaceParams(equalP map[string]EqualParameter)   {}
func (name *Name) ReplaceParams(equalP map[string]EqualParameter)             {}

//func (parameter *Parameter) ReplaceParams(equalP map[string]EqualParameter)   {}
func (inP *InParameter) ReplaceParams(equalP map[string]EqualParameter) {}
func (field *Field) ReplaceParams(equalP map[string]EqualParameter)     {}

func (add *Add) ToPrintString(sb *strings.Builder) {
	duoPrintStringHelper(add.Left, add.Right, sb, '+')
}

func (sub *Sub) ToPrintString(sb *strings.Builder) {
	duoPrintStringHelper(sub.Left, sub.Right, sb, '-')
}

func (mult *Mult) ToPrintString(sb *strings.Builder) {
	duoPrintStringHelper(mult.Left, mult.Right, sb, '*')
}

func (div *Div) ToPrintString(sb *strings.Builder) {
	duoPrintStringHelper(div.Left, div.Right, sb, '/')
}

func duoPrintStringHelper(left, right Calculable, sb *strings.Builder, op rune) {
	sb.WriteRune('(')
	left.ToPrintString(sb)
	sb.WriteString(") ")
	sb.WriteRune(op)
	sb.WriteString("( ")
	right.ToPrintString(sb)
	sb.WriteRune(')')
}

func (minus *Minus) ToPrintString(sb *strings.Builder) {
	sb.WriteString("-(")
	minus.ToCalculate.ToPrintString(sb)
	sb.WriteRune(')')
}

func (sConst *StringConstant) ToPrintString(sb *strings.Builder) {
	sb.WriteRune('"')
	sb.WriteString(sConst.Value)
	sb.WriteRune('"')
}

func (iConst *IntConstant) ToPrintString(sb *strings.Builder) {
	sb.WriteString("int(")
	sb.WriteString(strconv.Itoa(int(iConst.Value)))
	sb.WriteRune(')')
}

func (fConst *FloatConstant) ToPrintString(sb *strings.Builder) {
	sb.WriteString("float(")
	sb.WriteString(strconv.FormatFloat(fConst.Value, 'f', fConst.Precision, 64))
	sb.WriteRune(')')
}

func (dConst *DateConstant) ToPrintString(sb *strings.Builder) {
	sb.WriteString("date(")
	sb.WriteString(dConst.Date.ToString())
	sb.WriteRune(')')
}

func (name *Name) ToPrintString(sb *strings.Builder) {
	sb.WriteString(name.Name)
}

/*func (parameter *Parameter) ToPrintString(sb *strings.Builder) {
	sb.WriteRune('[')
	sb.WriteString(parameter.Name)
	sb.WriteRune(']')
}*/

func (inP *InParameter) ToPrintString(sb *strings.Builder) {
	sb.WriteRune('[')
	sb.WriteString(inP.Name)
	sb.WriteRune(']')
}

func (field *Field) ToPrintString(sb *strings.Builder) {
	sb.WriteString(field.Name)
	sb.WriteRune('.')
	sb.WriteString(field.FieldType.ToString())
}

func (add *Add) ToString() string {
	var sb strings.Builder
	add.ToPrintString(&sb)
	return sb.String()
}

func (sub *Sub) ToString() string {
	var sb strings.Builder
	sub.ToPrintString(&sb)
	return sb.String()
}

func (mult *Mult) ToString() string {
	var sb strings.Builder
	mult.ToPrintString(&sb)
	return sb.String()
}

func (div *Div) ToString() string {
	var sb strings.Builder
	div.ToPrintString(&sb)
	return sb.String()
}

func (minus *Minus) ToString() string {
	var sb strings.Builder
	minus.ToPrintString(&sb)
	return sb.String()
}

func (sConst *StringConstant) ToString() string {
	var sb strings.Builder
	sConst.ToPrintString(&sb)
	return sb.String()
}

func (iConst *IntConstant) ToString() string {
	var sb strings.Builder
	iConst.ToPrintString(&sb)
	return sb.String()
}

func (fConst *FloatConstant) ToString() string {
	var sb strings.Builder
	fConst.ToPrintString(&sb)
	return sb.String()
}

func (dConst *DateConstant) ToString() string {
	var sb strings.Builder
	dConst.ToPrintString(&sb)
	return sb.String()
}

func (name *Name) ToString() string {
	return name.Name
}

/*func (parameter *Parameter) ToString() string {
	return "[" + parameter.Name + "]"
}*/

func (inP *InParameter) ToString() string {
	return "[" + inP.Name + "]"
}

func (field *Field) ToString() string {
	return field.Name + "." + field.FieldType.ToString()
}

func (name *Name) GetIdentifier() string { return name.Name }

//func (parameter *Parameter) GetIdentifier() string { return parameter.Name }
func (inP *InParameter) GetIdentifier() string { return inP.Name }
func (field *Field) GetIdentifier() string     { return field.Name + "." + field.FieldTypeToString() }

func (name *Name) GetCalculable() Calculable                     { return name }
func (field *Field) GetCalculable() Calculable                   { return field }
func (rangeP *RangeParameterInt) GetCalculable() Calculable      { return rangeP }
func (rangeP *RangeParameterFloat) GetCalculable() Calculable    { return rangeP }
func (rangeP *RangeParameterDate) GetCalculable() Calculable     { return rangeP }
func (sparseP *SparseParameterInt) GetCalculable() Calculable    { return sparseP }
func (sparseP *SparseParameterFloat) GetCalculable() Calculable  { return sparseP }
func (sparseP *SparseParameterDate) GetCalculable() Calculable   { return sparseP }
func (sparseP *SparseParameterString) GetCalculable() Calculable { return sparseP }

func (inP *InParameter) GetFloatStep() float64 {
	return math.Pow(0.1, float64(inP.FloatPrecision))
}

func (date *DateConstant) FromString(dateInString string) {
	date.Date = date.Date.FromString(dateInString)
}

func (date Date) ToString() string {
	dayString := strconv.Itoa(date.Day)
	monthString := strconv.Itoa(date.Month)
	yearString := strconv.Itoa(date.Year)
	if date.Day <= 9 {
		dayString = "0" + dayString
	}
	if date.Month <= 9 {
		monthString = "0" + monthString
	}
	return fmt.Sprintf("%s-%s-%s", yearString, monthString, dayString)
}

func (date Date) FromString(dateInString string) Date {
	split := strings.Split(dateInString, "-")
	date.Year, _ = strconv.Atoi(split[0])
	date.Month, _ = strconv.Atoi(split[1])
	date.Day, _ = strconv.Atoi(split[2])
	return date
}

func (date Date) GetDay() int     { return date.Day }
func (date Date) GetMonth() int   { return date.Month }
func (date Date) GetYear() int    { return date.Year }
func (date Date) GetQuarter() int { return (date.Month-1)/3 + 1 }

func (date Date) IsHigherOrEqual(otherDate Date) bool {
	if otherDate == date {
		return true
	}
	if date.Year > otherDate.Year {
		return true
	}
	if date.Year < otherDate.Year {
		return false
	}
	if date.Month > otherDate.Month {
		return true
	}
	if date.Month < otherDate.Month {
		return false
	}
	return date.Day > otherDate.Day
}

func (date Date) IsHigher(otherDate Date) bool {
	if date.Year > otherDate.Year {
		return true
	}
	if date.Year < otherDate.Year {
		return false
	}
	if date.Month > otherDate.Month {
		return true
	}
	if date.Month < otherDate.Month {
		return false
	}
	return date.Day > otherDate.Day
}

func (date Date) IsLowerOrEqual(otherDate Date) bool {
	if otherDate == date {
		return true
	}
	if date.Year < otherDate.Year {
		return true
	}
	if date.Year > otherDate.Year {
		return false
	}
	if date.Month < otherDate.Month {
		return true
	}
	if date.Month > otherDate.Month {
		return false
	}
	return date.Day < otherDate.Day
}

func (date Date) IsLower(otherDate Date) bool {
	if date.Year < otherDate.Year {
		return true
	}
	if date.Year > otherDate.Year {
		return false
	}
	if date.Month < otherDate.Month {
		return true
	}
	if date.Month > otherDate.Month {
		return false
	}
	return date.Day < otherDate.Day
}

func GetDateFieldType(fieldName string) DateFieldType {
	fieldName = strings.ToLower(fieldName)
	switch fieldName {
	case "day":
		return DATE_FIELD_DAY
	case "month":
		return DATE_FIELD_MONTH
	case "year":
		return DATE_FIELD_YEAR
	case "quarter":
		return DATE_FIELD_QUARTER
	}
	return -1
}

func (field *Field) FieldTypeToString() string {
	switch field.FieldType {
	case DATE_FIELD_DAY:
		return "day"
	case DATE_FIELD_MONTH:
		return "month"
	case DATE_FIELD_YEAR:
		return "year"
	case DATE_FIELD_QUARTER:
		return "quarter"
	}
	return "unknown"
}

func (rangeP *RangeParameterInt) IsInRange(value interface{}) bool {
	switch typedValue := value.(type) {
	case int:
		return typedValue >= rangeP.MinValue && typedValue <= rangeP.MaxValue
	case float64:
		return int(typedValue) >= rangeP.MinValue && int(typedValue) <= rangeP.MaxValue
	}
	return false
}

func (rangeP *RangeParameterFloat) IsInRange(value interface{}) bool {
	switch typedValue := value.(type) {
	case float64:
		return typedValue >= rangeP.MinValue && typedValue <= rangeP.MaxValue
	case int:
		return float64(typedValue) >= rangeP.MinValue && float64(typedValue) <= rangeP.MaxValue
	}
	return false
}

func (rangeP *RangeParameterDate) IsInRange(value interface{}) bool {
	dateValue := value.(Date)
	return dateValue.IsHigherOrEqual(rangeP.MinValue) && dateValue.IsLowerOrEqual(rangeP.MaxValue)
}

/*
	case STRING_TYPE:
		stringValue := value.(string)
		return stringValue >= rangeP.MinValue.(string) && stringValue <= rangeP.MaxValue.(string)
*/

func (sparseP *SparseParameterInt) IsInRange(value interface{}) bool {
	has := false
	switch typedValue := value.(type) {
	case int:
		_, has = sparseP.Values[typedValue]
	case float64:
		intV := int(typedValue)
		if float64(intV) < typedValue { //This means typedValue had some decimals, thus it can't be in a sparse range of ints
			has = false //Unecessary but it's here for readibility
		} else {
			_, has = sparseP.Values[int(typedValue)]
		}
	}
	return has
}

func (sparseP *SparseParameterFloat) IsInRange(value interface{}) bool {
	has := false
	switch typedValue := value.(type) {
	case float64:
		_, has = sparseP.Values[typedValue]
	case int:
		_, has = sparseP.Values[float64(typedValue)]
	}
	return has
}

func (sparseP *SparseParameterDate) IsInRange(value interface{}) bool {
	_, has := sparseP.Values[value.(Date)]
	return has
}

func (sparseP *SparseParameterString) IsInRange(value interface{}) bool {
	_, has := sparseP.Values[value.(string)]
	return has
}

//Conditions are in the form of:
//PARAMETER op value
func (rangeP *RangeParameterInt) IsInInequalityRange(value interface{}, op CondType) bool {
	//An example:
	//Let's assume DATE.min = 03-03-95 and DATE.max = 24-03-95
	//[DATE] > Orderdate <=> Orderdate < [DATE]: we want this to be true for any value that is < 24-03-95
	//[DATE] < Shipdate <=> Shipdate > [DATE]: we want this to be true for any value that is > 03-03-95
	//The first is true if value < MAX or MAX > value
	//The second is true if value > MIN or MIN < value
	switch typedV := value.(type) {
	case int:
		return intInequalityRange(rangeP.MinValue, rangeP.MaxValue, typedV, op)
	case float64:
		return floatInequalityRange(float64(rangeP.MinValue), float64(rangeP.MaxValue), typedV, op)
	}
	return false
}

func (rangeP *RangeParameterFloat) IsInInequalityRange(value interface{}, op CondType) bool {
	switch typedV := value.(type) {
	case float64:
		return floatInequalityRange(rangeP.MinValue, rangeP.MaxValue, typedV, op)
	case int:
		return floatInequalityRange(rangeP.MinValue, rangeP.MaxValue, float64(typedV), op)
	}
	return false
}

func (rangeP *RangeParameterDate) IsInInequalityRange(value interface{}, op CondType) bool {
	return dateInequalityRange(rangeP.MinValue, rangeP.MaxValue, value.(Date), op)
}

//Inequality range for sparse is the same logic as for range
//E.g.: revenue > [VALUE].
//If e.g. revenue = 1000 and VALUE = {200, 500, 1100}, this will be true for the same situations
//as if VALUE = [200,1100]. In this case revenue is valid for 200 and 500 but not 1100.
//Aka, the same rule of revenue > min applies.
func (sparseP *SparseParameterInt) IsInInequalityRange(value interface{}, op CondType) bool {
	switch typedV := value.(type) {
	case int:
		return intInequalityRange(sparseP.MinValue, sparseP.MaxValue, typedV, op)
	case float64:
		return floatInequalityRange(float64(sparseP.MinValue), float64(sparseP.MaxValue), typedV, op)
	}
	return false
}

func (sparseP *SparseParameterFloat) IsInInequalityRange(value interface{}, op CondType) bool {
	switch typedV := value.(type) {
	case float64:
		return floatInequalityRange(sparseP.MinValue, sparseP.MaxValue, typedV, op)
	case int:
		return floatInequalityRange(sparseP.MinValue, sparseP.MaxValue, float64(typedV), op)
	}
	return false
}

func (sparseP *SparseParameterDate) IsInInequalityRange(value interface{}, op CondType) bool {
	return dateInequalityRange(sparseP.MinValue, sparseP.MaxValue, value.(Date), op)
}

func (sparseP *SparseParameterString) IsInInequalityRange(value interface{}, op CondType) bool {
	//Not supported.
	return false
}

func intInequalityRange(min, max, value int, op CondType) bool {
	switch op {
	case HIGHER:
		return max > value
	case HIGHER_EQ:
		return max >= value
	case LOWER:
		return min < value
	case LOWER_EQ:
		return min <= value
	case NOT_EQ:
		if min != max {
			return true //If this really reflects a range, it will always be different of at least one of the values.
		}
		return min != value //min == max, so it will be different if value != min or != max.
	}
	return false
}

func floatInequalityRange(min, max, value float64, op CondType) bool {
	switch op {
	case HIGHER:
		return max > value
	case HIGHER_EQ:
		return max >= value
	case LOWER:
		return min < value
	case LOWER_EQ:
		return min <= value
	case NOT_EQ:
		if min != max {
			return true //If this really reflects a range, it will always be different of at least one of the values.
		}
		return min != value //min == max, so it will be different if value != min or != max.
	}
	return false
}

func dateInequalityRange(min, max, value Date, op CondType) bool {
	fmt.Printf("[DateIneqRange]. Min: %s, Max: %s, Value: %s. CondType: %d\n", min.ToString(),
		max.ToString(), value.ToString(), op)
	switch op {
	case HIGHER:
		return max.IsHigher(value)
	case HIGHER_EQ:
		return max.IsHigherOrEqual(value)
	case LOWER:
		return min.IsLower(value)
	case LOWER_EQ:
		return min.IsLowerOrEqual(value)
	case NOT_EQ:
		if min != max {
			return true
		}
		return min != value
	}
	return false
}

func (sparseP *SparseParameterInt) CalculateMinMax() {
	min, max := math.MaxInt64, math.MinInt64
	for value := range sparseP.Values {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	sparseP.MinValue, sparseP.MaxValue = min, max
}

func (sparseP *SparseParameterFloat) CalculateMinMax() {
	min, max := math.MaxFloat64, -math.MaxFloat64
	for value := range sparseP.Values {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	sparseP.MinValue, sparseP.MaxValue = min, max
}

func (sparseP *SparseParameterDate) CalculateMinMax() {
	min, max := Date{Day: 1, Month: 1, Year: 1}, Date{Day: 30, Month: 1, Year: math.MaxInt64}
	for value := range sparseP.Values {
		if value.IsLower(min) {
			min = value
		}
		if value.IsHigher(max) {
			max = value
		}
	}
	sparseP.MinValue, sparseP.MaxValue = min, max
}

func (rangeP *RangeParameterInt) GetIneqKeys(conds []Condition, objValues map[string]interface{}) (keys []string) {
	min, max := GetMaxMinIntFromConds(rangeP.MinValue, rangeP.MaxValue, conds, objValues)
	if max < min { //Not possible to satisfact all conditions
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

func (sparseP *SparseParameterInt) GetIneqKeys(conds []Condition, objValues map[string]interface{}) (keys []string) {
	min, max := GetMaxMinIntFromConds(sparseP.MinValue, sparseP.MaxValue, conds, objValues)
	if max < min { //Not possible to satisfy all conditions
		keys = make([]string, 0)
	} else {
		keys = make([]string, IntMin(max-min+1, len(sparseP.Values)))
		currI := 0
		for paramV, _ := range sparseP.Values {
			if paramV >= min && paramV <= max {
				keys[currI] = strconv.Itoa(paramV)
				currI++
			}
		}
		keys = keys[:currI]
	}
	return keys
}

func GetMaxMinIntFromConds(initialMin, initialMax int, conds []Condition, objValues map[string]interface{}) (newMin, newMax int) {
	newMin, newMax = initialMin, initialMax
	fmt.Printf("Min %d, max %d\n", newMin, newMax)
	for _, cond := range conds {
		value := cond.Math.GetCalculable().CalculateValue(objValues)
		var intV int
		switch typedV := value.(type) {
		case float64:
			intV = int(typedV)
		case int:
			intV = typedV
		}
		switch cond.Op {
		case HIGHER:
			newMin = IntMax(newMin, intV+1)
		case HIGHER_EQ:
			newMin = IntMax(newMin, intV)
		case LOWER:
			newMax = IntMin(newMax, intV-1)
		case LOWER_EQ:
			newMax = IntMin(newMax, intV)
		}
		fmt.Printf("Min %d, max %d\n", newMin, newMax)
	}
	return
}

func (rangeP *RangeParameterFloat) GetIneqKeys(conds []Condition, objValues map[string]interface{}) (keys []string) {
	step, precision := rangeP.GetFloatStep(), rangeP.FloatPrecision
	min, max := GetMaxMinFloatFromConds(rangeP.MinValue, rangeP.MaxValue, step, conds, objValues)

	if max < min { //Not possible to satisfy all conditions
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

func (sparseP *SparseParameterFloat) GetIneqKeys(conds []Condition, objValues map[string]interface{}) (keys []string) {
	min, max := GetMaxMinFloatFromConds(sparseP.MinValue, sparseP.MaxValue, 0.00000000001, conds, objValues)

	if max < min { //Not possible to satisfy all conditions
		keys = make([]string, 0)
	} else {
		keys = make([]string, len(sparseP.Values))
		currI := 0
		for paramV, _ := range sparseP.Values {
			if paramV >= min && paramV <= max {
				keys[currI] = strconv.FormatFloat(paramV, 'f', -1, 64)
				currI++
			}
		}
		keys = keys[:currI]
	}
	return
}

func GetMaxMinFloatFromConds(initialMin, initialMax, step float64, conds []Condition, objValues map[string]interface{}) (newMin, newMax float64) {
	newMin, newMax = initialMin, initialMax
	for _, cond := range conds {
		value := cond.Math.GetCalculable().CalculateValue(objValues)
		var floatV float64
		switch typedV := value.(type) {
		case float64:
			floatV = typedV
		case int:
			floatV = float64(typedV)
		}
		switch cond.Op {
		case HIGHER:
			newMin = math.Max(newMin, floatV+step)
		case HIGHER_EQ:
			newMin = math.Max(newMin, floatV)
		case LOWER:
			newMax = math.Min(newMax, floatV-step)
		case LOWER_EQ:
			newMax = math.Min(newMax, floatV)
		}
	}
	return
}

func (rangeP *RangeParameterDate) GetIneqKeys(conds []Condition, objValues map[string]interface{}) (keys []string) {
	min, max := GetMaxMinDateFromConds(rangeP.MinValue, rangeP.MaxValue, conds, objValues)

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

func (sparseP *SparseParameterDate) GetIneqKeys(conds []Condition, objValues map[string]interface{}) (keys []string) {
	min, max := GetMaxMinDateFromConds(sparseP.MinValue, sparseP.MaxValue, conds, objValues)

	minD := time.Date(min.Year, time.Month(min.Month), min.Day, 0, 0, 0, 0, time.UTC)
	maxD := time.Date(max.Year, time.Month(max.Month), max.Day, 0, 0, 0, 0, time.UTC)
	diffDays := int(maxD.Sub(minD).Hours()/24) + 1
	//TODO: Consider storing these Sparse Date as time.Date
	if diffDays < 0 { //Not possible to satisfy all conditions
		keys = make([]string, 0)
	} else {
		keys = make([]string, IntMin(diffDays, len(sparseP.Values)))
		currI := 0
		for paramV, _ := range sparseP.Values {
			if paramV.IsHigherOrEqual(min) && paramV.IsLowerOrEqual(max) {
				keys[currI] = ValueToString(paramV)
				currI++
			}
		}
		keys = keys[:currI]
	}
	return
}

func GetMaxMinDateFromConds(initialMin, initialMax Date, conds []Condition, objValues map[string]interface{}) (newMin, newMax Date) {
	newMin, newMax = initialMin, initialMax
	for _, cond := range conds {
		value := cond.Math.GetCalculable().CalculateValue(objValues).(Date)
		switch cond.Op {
		case HIGHER:
			if value.IsHigher(newMin) {
				tmp := time.Date(value.Year, time.Month(value.Month), value.Day+1, 0, 0, 0, 0, time.UTC)
				newMin = Date{Year: tmp.Year(), Month: int(tmp.Month()), Day: int(tmp.Day())}
			}
		case HIGHER_EQ:
			if value.IsHigher(newMin) {
				newMin = value
			}
		case LOWER:
			if value.IsLower(newMax) {
				tmp := time.Date(value.Year, time.Month(value.Month), value.Day-1, 0, 0, 0, 0, time.UTC)
				newMax = Date{Year: tmp.Year(), Month: int(tmp.Month()), Day: int(tmp.Day())}
			}
		case LOWER_EQ:
			if value.IsLower(newMax) {
				newMax = value
			}
		}
	}
	return
}

func (sparseP *SparseParameterString) GetIneqKeys(conds []Condition, objValues map[string]interface{}) (keys []string) {
	//Not supported
	return make([]string, 0)
}

/*
func (ineqP *InequalParameter) IsInRange(value interface{}) bool {
	switch ineqP.Type {
	case INT_TYPE:
		minI, maxI := ineqP.MinValue.(int), ineqP.MaxValue.(int)
		switch typedValue := value.(type) {
		case int:
			return typedValue >= minI && typedValue <= maxI
		case float64:
			return int(typedValue) >= minI && int(typedValue) <= maxI
		}
	case FLOAT_TYPE:
		minF, maxF := ineqP.MinValue.(float64), ineqP.MaxValue.(float64)
		switch typedValue := value.(type) {
		case float64:
			return typedValue >= minF && typedValue <= maxF
		case int:
			return float64(typedValue) >= minF && float64(typedValue) <= maxF
		}
	case DATE_TYPE:
		dateValue := value.(Date)
		return dateValue.IsHigherOrEqual(ineqP.MinValue.(Date)) && dateValue.IsLowerOrEqual(ineqP.MaxValue.(Date))
	case STRING_TYPE:
		stringValue := value.(string)
		return stringValue >= ineqP.MinValue.(string) && stringValue <= ineqP.MaxValue.(string)
	}
	return false
}

func (ineqP *InequalParameter) GetFloatStep() float64 {
	return math.Pow(0.1, float64(ineqP.FloatPrecision))
}
*/

//Forces a string representation of any of the data types
func ValueToString(value interface{}) string {
	switch typedV := value.(type) {
	case int:
		return strconv.Itoa(typedV)
	case float64:
		return strconv.FormatFloat(typedV, 'f', -1, 64)
	case string:
		return typedV
	case Date:
		return typedV.ToString()
	}
	return ""
}

/*****SQL -> Math conversion methods*****/

func ValueHelper(ctx *parser.ValueContext, currentMath Math) {

}

func MinusHelper(ctx *parser.MinusContext, currentMath Math) {
	currentMath.Stack.Push(&Minus{ToCalculate: currentMath.Stack.Pop()})
}

func AddOrSubHelper(ctx *parser.AddOrSubContext, currentMath Math) {
	right, left := currentMath.Stack.Pop(), currentMath.Stack.Pop()
	if ctx.GetOpType().GetText() == "+" {
		currentMath.Stack.Push(&Add{Left: left, Right: right})
	} else {
		currentMath.Stack.Push(&Sub{Left: left, Right: right})
	}
}

func MultOrDivHelper(ctx *parser.MultOrDivContext, currentMath Math) {
	right, left := currentMath.Stack.Pop(), currentMath.Stack.Pop()
	if ctx.GetOpType().GetText() == "*" {
		currentMath.Stack.Push(&Mult{Left: left, Right: right})
	} else {
		currentMath.Stack.Push(&Div{Left: left, Right: right})
	}
}
