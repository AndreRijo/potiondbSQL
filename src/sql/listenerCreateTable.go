package sql

import (
	"fmt"
	"strconv"
	"strings"

	"sqlToKeyValue/src/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type CTPolicy int
type SQLDatatype int

type Invariant interface {
	IsInvariant() bool //This method is only for type purposes
}

type PrimaryKey bool

type ForeignKey struct {
	ForeignTable  string
	ForeignColumn string
}

type CheckConstraint struct {
	ConditionType CondType
	Value         int
}

func (pk PrimaryKey) IsInvariant() bool      { return true }
func (fk ForeignKey) IsInvariant() bool      { return true }
func (cc CheckConstraint) IsInvariant() bool { return true }

const (
	AW, RW, LWW                                           CTPolicy    = 1, 2, 3
	COUNTER, INTEGER, BOOLEAN, VARCHAR, DATE, UNSUPPORTED SQLDatatype = 1, 2, 3, 4, 5, 0
)

type ListenerCreateTable struct {
	Parse *parser.ViewSQLParser

	ConcurrencyPolicy CTPolicy
	TableName         string
	//Types             map[string]proto.CRDTType
	//Defaults          map[string]crdt.UpdateArguments
	Types      map[string]SQLDatatype
	Defaults   map[string]string
	Invariants map[string]Invariant
	Columns    []string

	//Temporary/helper variables
	currentVarName string
}

func MakeCreateTableListener(parse *parser.ViewSQLParser) *ListenerCreateTable {
	return &ListenerCreateTable{Parse: parse}
}

// VisitTerminal is called when a terminal node is visited.
func (listen *ListenerCreateTable) VisitTerminal(node antlr.TerminalNode) {
}

// VisitErrorNode is called when an error node is visited.
func (listen *ListenerCreateTable) VisitErrorNode(node antlr.ErrorNode) {
}

// EnterEveryRule is called when any rule is entered.
func (listen *ListenerCreateTable) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println("[CreateTableL]Entered rule")
}

// ExitEveryRule is called when any rule is exited.
func (listen *ListenerCreateTable) ExitEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println("[CreateTableL]Exit rule")
}

// EnterName is called when entering the constant production.
func (listen *ListenerCreateTable) EnterName(ctx *parser.NameContext) {}

// EnterConstant is called when entering the constant production.
func (listen *ListenerCreateTable) EnterConstant(ctx *parser.ConstantContext) {}

// EnterAggrFunc is called when entering the aggrFunc production.
func (listen *ListenerCreateTable) EnterAggrFunc(ctx *parser.AggrFuncContext) {}

// EnterField is called when entering the field production.
func (listen *ListenerCreateTable) EnterField(ctx *parser.FieldContext) {}

// EnterParameter is called when entering the parameter production.
func (listen *ListenerCreateTable) EnterParameter(ctx *parser.ParameterContext) {}

// EnterNameable is called when entering the nameable production.
func (listen *ListenerCreateTable) EnterNameable(ctx *parser.NameableContext) {}

// EnterKey is called when entering the key production.
func (listen *ListenerCreateTable) EnterKey(ctx *parser.KeyContext) {}

// EnterAddOrSub is called when entering the AddOrSub production.
func (listen *ListenerCreateTable) EnterAddOrSub(ctx *parser.AddOrSubContext) {}

// EnterValue is called when entering the Value production.
func (listen *ListenerCreateTable) EnterValue(ctx *parser.ValueContext) {}

// EnterMinus is called when entering the Minus production.
func (listen *ListenerCreateTable) EnterMinus(ctx *parser.MinusContext) {}

// EnterParentheses is called when entering the Parentheses production.
func (listen *ListenerCreateTable) EnterParentheses(ctx *parser.ParenthesesContext) {}

// EnterMultOrDiv is called when entering the MultOrDiv production.
func (listen *ListenerCreateTable) EnterMultOrDiv(ctx *parser.MultOrDivContext) {}

// EnterAsClause is called when entering the asClause production.
func (listen *ListenerCreateTable) EnterAsClause(ctx *parser.AsClauseContext) {}

// EnterAggregation is called when entering the aggregation production.
func (listen *ListenerCreateTable) EnterAggregation(ctx *parser.AggregationContext) {}

// EnterCount is called when entering the count production.
func (listen *ListenerCreateTable) EnterCount(ctx *parser.CountContext) {}

// EnterCalc is called when entering the calc production.
func (listen *ListenerCreateTable) EnterCalc(ctx *parser.CalcContext) {}

// EnterComp is called when entering the comp production.
func (listen *ListenerCreateTable) EnterComp(ctx *parser.CompContext) {}

// EnterCondition is called when entering the condition production.
func (listen *ListenerCreateTable) EnterCondition(ctx *parser.ConditionContext) {}

// EnterSortOrder is called when entering the sortOrder production.
func (listen *ListenerCreateTable) EnterSortOrder(ctx *parser.SortOrderContext) {}

// EnterContinuousRange is called when entering the continuousRange production.
func (listen *ListenerCreateTable) EnterContinuousRange(ctx *parser.ContinuousRangeContext) {}

// EnterSparseRange is called when entering the sparseRange production.
func (listen *ListenerCreateTable) EnterSparseRange(ctx *parser.SparseRangeContext) {}

// EnterRange is called when entering the range production.
func (listen *ListenerCreateTable) EnterRange(ctx *parser.RangeContext) {}

// EnterCreate is called when entering the create production.
func (listen *ListenerCreateTable) EnterCreate(ctx *parser.CreateContext) {
	fmt.Println("[CreateTableL]***Enter create***")
}

// EnterWith is called when entering the with production.
func (listen *ListenerCreateTable) EnterWith(ctx *parser.WithContext) {}

// EnterSelect is called when entering the select production.
func (listen *ListenerCreateTable) EnterSelect(ctx *parser.SelectContext) {}

// EnterFrom is called when entering the from production.
func (listen *ListenerCreateTable) EnterFrom(ctx *parser.FromContext) {}

// EnterWhere is called when entering the where production.
func (listen *ListenerCreateTable) EnterWhere(ctx *parser.WhereContext) {}

// EnterGroupby is called when entering the groupby production.
func (listen *ListenerCreateTable) EnterGroupby(ctx *parser.GroupbyContext) {}

// EnterOrderby is called when entering the orderby production.
func (listen *ListenerCreateTable) EnterOrderby(ctx *parser.OrderbyContext) {}

// EnterLimit is called when entering the limit production.
func (listen *ListenerCreateTable) EnterLimit(ctx *parser.LimitContext) {}

// EnterView is called when entering the view production.
func (listen *ListenerCreateTable) EnterView(ctx *parser.ViewContext) {}

// EnterCheck is called when entering the check production.
func (listen *ListenerCreateTable) EnterCheck(ctx *parser.CheckContext) {
	//Constant must be int as check only supports ints
	value, _ := strconv.Atoi(ctx.Constant().INT().GetText())
	compTypeS, check := ctx.GetType_().GetText(), CheckConstraint{Value: value}
	switch compTypeS {
	case ">":
		check.ConditionType = HIGHER
	case "<":
		check.ConditionType = LOWER
	case ">=":
		check.ConditionType = HIGHER_EQ
	case "<=":
		check.ConditionType = LOWER_EQ
	default:
		fmt.Println("[CreateTableL]Check: unsupported comparing type, ", compTypeS)
	}
	listen.Invariants[listen.currentVarName] = check
	//listen.Types[listen.currentVarName] = proto.CRDTType_FATCOUNTER //Update to bounded counter
}

// EnterForeignkey is called when entering the foreignkey production.
func (listen *ListenerCreateTable) EnterForeignkey(ctx *parser.ForeignkeyContext) {
	listen.Invariants[listen.currentVarName] = ForeignKey{ForeignTable: ctx.GetTableName().GetText(), ForeignColumn: ctx.GetColumnName().GetText()}
}

// EnterPrimarykey is called when entering the primarykey production.
func (listen *ListenerCreateTable) EnterPrimarykey(ctx *parser.PrimarykeyContext) {
	listen.Invariants[listen.currentVarName] = PrimaryKey(true)
}

// EnterConstraint is called when entering the constraint production.
func (listen *ListenerCreateTable) EnterConstraint(ctx *parser.ConstraintContext) {}

// EnterColumns is called when entering the columns production.
func (listen *ListenerCreateTable) EnterColumns(ctx *parser.ColumnsContext) {
	listen.currentVarName = ctx.STRING().GetText()
	fmt.Println("[CreateTableL]VarName:", listen.currentVarName)
	listen.Columns = append(listen.Columns, listen.currentVarName)

	dataTypeString := strings.ToLower(ctx.GetType_().GetText())
	//crdtType := DataTypeSQLToCRDTType(dataTypeString)
	//listen.Types[listen.currentVarName] = crdtType
	//fmt.Println("[CreateTableL]CRDTType:", crdtType)
	sqlType := ParsedTypeToSQLDatatype(dataTypeString)
	listen.Types[listen.currentVarName] = sqlType
	fmt.Println("[CreateTableL]SQLType:", sqlType)

	fmt.Println("[CreateTableL]Constant:", ctx.Constant())
	if defaultValue := ctx.Constant(); defaultValue != nil {
		//listen.Defaults[listen.currentVarName] = SQLUpdateToCRDTUpdate(dataTypeString, defaultValue)
		listen.Defaults[listen.currentVarName] = SQLUpdateToString(defaultValue)
	}
}

// EnterCreatetable is called when entering the table production.
func (listen *ListenerCreateTable) EnterCreatetable(ctx *parser.CreatetableContext) {
	fmt.Println("[CreateTableL]***Enter createtable***")
	//listen.Types, listen.Defaults, listen.Invariants = make(map[string]proto.CRDTType),
	//make(map[string]crdt.UpdateArguments), make(map[string]Invariant)
	listen.Types, listen.Defaults, listen.Invariants = make(map[string]SQLDatatype),
		make(map[string]string), make(map[string]Invariant)
	policyString, tableName := ctx.GetPolicy().GetText(), ctx.Name()
	listen.TableName = tableName.GetText()
	switch policyString {
	case "AW":
		listen.ConcurrencyPolicy = AW
	case "RW":
		listen.ConcurrencyPolicy = RW
	case "LWW":
		listen.ConcurrencyPolicy = LWW
	}
}

// EnterCreateindex is called when entering the createindex production.
func (listen *ListenerCreateTable) EnterCreateindex(ctx *parser.CreateindexContext) {}

// EnterDrop is called when entering the drop production.
func (listen *ListenerCreateTable) EnterDrop(ctx *parser.DropContext) {}

// EnterDelete is called when entering the delete production.
func (listen *ListenerCreateTable) EnterDelete(ctx *parser.DeleteContext) {}

// EnterSet is called when entering the set production.
func (listen *ListenerCreateTable) EnterSet(ctx *parser.SetContext) {}

// EnterUpdate is called when entering the update production.
func (listen *ListenerCreateTable) EnterUpdate(ctx *parser.UpdateContext) {}

// EnterValues is called when entering the values production.
func (listen *ListenerCreateTable) EnterValues(ctx *parser.ValuesContext) {}

// EnterColumnNames is called when entering the columnNames production.
func (listen *ListenerCreateTable) EnterColumnNames(ctx *parser.ColumnNamesContext) {}

// EnterInsert is called when entering the insert production.
func (listen *ListenerCreateTable) EnterInsert(ctx *parser.InsertContext) {}

// EnterQuery is called when entering the query production.
func (listen *ListenerCreateTable) EnterQuery(ctx *parser.QueryContext) {}

// EnterStatement is called when entering the statement production.
func (listen *ListenerCreateTable) EnterStatement(ctx *parser.StatementContext) {}

// EnterStart is called when entering the start production.
func (listen *ListenerCreateTable) EnterStart(ctx *parser.StartContext) {}

// ExitName is called when exiting the name production.
func (listen *ListenerCreateTable) ExitName(ctx *parser.NameContext) {}

// ExitConstant is called when exiting the constant production.
func (listen *ListenerCreateTable) ExitConstant(ctx *parser.ConstantContext) {}

// ExitAggrFunc is called when exiting the aggrFunc production.
func (listen *ListenerCreateTable) ExitAggrFunc(ctx *parser.AggrFuncContext) {}

// ExitField is called when exiting the field production.
func (listen *ListenerCreateTable) ExitField(ctx *parser.FieldContext) {}

// ExitParameter is called when exiting the parameter production.
func (listen *ListenerCreateTable) ExitParameter(ctx *parser.ParameterContext) {}

// ExitNameable is called when exiting the nameable production.
func (listen *ListenerCreateTable) ExitNameable(ctx *parser.NameableContext) {}

// ExitKey is called when exiting the key production.
func (listen *ListenerCreateTable) ExitKey(ctx *parser.KeyContext) {}

// ExitAddOrSub is called when exiting the AddOrSub production.
func (listen *ListenerCreateTable) ExitAddOrSub(ctx *parser.AddOrSubContext) {}

// ExitValue is called when exiting the Value production.
func (listen *ListenerCreateTable) ExitValue(ctx *parser.ValueContext) {}

// ExitMinus is called when exiting the Minus production.
func (listen *ListenerCreateTable) ExitMinus(ctx *parser.MinusContext) {}

// ExitParentheses is called when exiting the Parentheses production.
func (listen *ListenerCreateTable) ExitParentheses(ctx *parser.ParenthesesContext) {}

// ExitMultOrDiv is called when exiting the MultOrDiv production.
func (listen *ListenerCreateTable) ExitMultOrDiv(ctx *parser.MultOrDivContext) {}

// ExitAsClause is called when exiting the asClause production.
func (listen *ListenerCreateTable) ExitAsClause(ctx *parser.AsClauseContext) {}

// ExitAggregation is called when exiting the aggregation production.
func (listen *ListenerCreateTable) ExitAggregation(ctx *parser.AggregationContext) {}

// ExitCount is called when exiting the count production.
func (listen *ListenerCreateTable) ExitCount(ctx *parser.CountContext) {}

// ExitCalc is called when exiting the calc production.
func (listen *ListenerCreateTable) ExitCalc(ctx *parser.CalcContext) {}

// ExitComp is called when exiting the comp production.
func (listen *ListenerCreateTable) ExitComp(ctx *parser.CompContext) {}

// ExitCondition is called when exiting the condition production.
func (listen *ListenerCreateTable) ExitCondition(ctx *parser.ConditionContext) {}

// ExitSortOrder is called when exiting the sortOrder production.
func (listen *ListenerCreateTable) ExitSortOrder(ctx *parser.SortOrderContext) {}

// ExitContinuousRange is called when exiting the continuousRange production.
func (listen *ListenerCreateTable) ExitContinuousRange(ctx *parser.ContinuousRangeContext) {}

// ExitSparseRange is called when exiting the sparseRange production.
func (listen *ListenerCreateTable) ExitSparseRange(ctx *parser.SparseRangeContext) {}

// ExitRange is called when exiting the range production.
func (listen *ListenerCreateTable) ExitRange(ctx *parser.RangeContext) {}

// ExitCreate is called when exiting the create production.
func (listen *ListenerCreateTable) ExitCreate(ctx *parser.CreateContext) {}

// ExitCreate is called when exiting the create production.
func (listen *ListenerCreateTable) ExitCreatetable(ctx *parser.CreatetableContext) {}

// ExitCreateindex is called when exiting the createindex production.
func (listen *ListenerCreateTable) ExitCreateindex(ctx *parser.CreateindexContext) {}

// ExitWith is called when exiting the with production.
func (listen *ListenerCreateTable) ExitWith(ctx *parser.WithContext) {}

// ExitSelect is called when exiting the select production.
func (listen *ListenerCreateTable) ExitSelect(ctx *parser.SelectContext) {}

// ExitFrom is called when exiting the from production.
func (listen *ListenerCreateTable) ExitFrom(ctx *parser.FromContext) {}

// ExitWhere is called when exiting the where production.
func (listen *ListenerCreateTable) ExitWhere(ctx *parser.WhereContext) {}

// ExitGroupby is called when exiting the groupby production.
func (listen *ListenerCreateTable) ExitGroupby(ctx *parser.GroupbyContext) {}

// ExitOrderby is called when exiting the orderby production.
func (listen *ListenerCreateTable) ExitOrderby(ctx *parser.OrderbyContext) {}

// ExitLimit is called when exiting the limit production.
func (listen *ListenerCreateTable) ExitLimit(ctx *parser.LimitContext) {}

// ExitView is called when exiting the view production.
func (listen *ListenerCreateTable) ExitView(ctx *parser.ViewContext) {}

// ExitCheck is called when exiting the check production.
func (listen *ListenerCreateTable) ExitCheck(ctx *parser.CheckContext) {}

// ExitForeignkey is called when exiting the foreignkey production.
func (listen *ListenerCreateTable) ExitForeignkey(ctx *parser.ForeignkeyContext) {}

// ExitPrimarykey is called when exiting the primarykey production.
func (listen *ListenerCreateTable) ExitPrimarykey(ctx *parser.PrimarykeyContext) {}

// ExitConstraint is called when exiting the constraint production.
func (listen *ListenerCreateTable) ExitConstraint(ctx *parser.ConstraintContext) {}

// ExitColumns is called when exiting the columns production.
func (listen *ListenerCreateTable) ExitColumns(ctx *parser.ColumnsContext) {}

// ExitTable is called when exiting the table production.
func (listen *ListenerCreateTable) ExitTable(ctx *parser.CreatetableContext) {}

// ExitDrop is called when exiting the drop production.
func (listen *ListenerCreateTable) ExitDrop(ctx *parser.DropContext) {}

// ExitDelete is called when exiting the delete production.
func (listen *ListenerCreateTable) ExitDelete(ctx *parser.DeleteContext) {}

// ExitSet is called when exiting the set production.
func (listen *ListenerCreateTable) ExitSet(ctx *parser.SetContext) {}

// ExitUpdate is called when exiting the update production.
func (listen *ListenerCreateTable) ExitUpdate(ctx *parser.UpdateContext) {}

// ExitValues is called when exiting the values production.
func (listen *ListenerCreateTable) ExitValues(ctx *parser.ValuesContext) {}

// ExitColumnNames is called when exiting the columnNames production.
func (listen *ListenerCreateTable) ExitColumnNames(ctx *parser.ColumnNamesContext) {}

// ExitInsert is called when exiting the insert production.
func (listen *ListenerCreateTable) ExitInsert(ctx *parser.InsertContext) {}

// ExitQuery is called when exiting the query production.
func (listen *ListenerCreateTable) ExitQuery(ctx *parser.QueryContext) {}

// ExitStatement is called when exiting the statement production.
func (listen *ListenerCreateTable) ExitStatement(ctx *parser.StatementContext) {}

// ExitStart is called when exiting the start production.
func (listen *ListenerCreateTable) ExitStart(ctx *parser.StartContext) {
	fmt.Println("[CreateTableL]Finished processing text (ExitStart)")
	fmt.Println("Concurrency policy:", listen.ConcurrencyPolicy)
	fmt.Println("TableName:", listen.TableName)
	fmt.Println("Types:", listen.Types)
	fmt.Println("Defaults:", listen.Defaults)
	fmt.Println("Invariants:", listen.Invariants)
}
