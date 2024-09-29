package sql

import (
	"fmt"

	"sqlToKeyValue/src/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type ListenerQuery struct {
	Parse       *parser.ViewSQLParser
	TableName   string
	ColumnNames []string
	Conditions  []SimpleCondition
	currentMath Math
	currIndex   int
	IsSelectAll bool
}

func MakeQueryListener(parse *parser.ViewSQLParser) *ListenerQuery {
	return &ListenerQuery{Parse: parse, currIndex: 0, IsSelectAll: false}
}

// VisitTerminal is called when a terminal node is visited.
func (listen *ListenerQuery) VisitTerminal(node antlr.TerminalNode) {
	//fmt.Println("[DropTableL]Terminal", node.GetText())
}

// VisitErrorNode is called when an error node is visited.
func (listen *ListenerQuery) VisitErrorNode(node antlr.ErrorNode) {
	//fmt.Println("[DropTableL]Error")
}

// EnterEveryRule is called when any rule is entered.
func (listen *ListenerQuery) EnterEveryRule(ctx antlr.ParserRuleContext) {
	//fmt.Println("[DropTableL]Entered rule", ctx.GetText())
}

// ExitEveryRule is called when any rule is exited.
func (listen *ListenerQuery) ExitEveryRule(ctx antlr.ParserRuleContext) {
	//fmt.Println("[DropTableL]Exit rule", ctx.GetText())
}

// EnterName is called when entering the constant production.
func (listen *ListenerQuery) EnterName(ctx *parser.NameContext) {}

// EnterConstant is called when entering the constant production.
func (listen *ListenerQuery) EnterConstant(ctx *parser.ConstantContext) {}

// EnterAggrFunc is called when entering the aggrFunc production.
func (listen *ListenerQuery) EnterAggrFunc(ctx *parser.AggrFuncContext) {}

// EnterField is called when entering the field production.
func (listen *ListenerQuery) EnterField(ctx *parser.FieldContext) {}

// EnterParameter is called when entering the parameter production.
func (listen *ListenerQuery) EnterParameter(ctx *parser.ParameterContext) {}

// EnterNameable is called when entering the nameable production.
func (listen *ListenerQuery) EnterNameable(ctx *parser.NameableContext) {}

// EnterKey is called when entering the key production.
func (listen *ListenerQuery) EnterKey(ctx *parser.KeyContext) {}

// EnterAddOrSub is called when entering the AddOrSub production.
func (listen *ListenerQuery) EnterAddOrSub(ctx *parser.AddOrSubContext) {}

// EnterValue is called when entering the Value production.
func (listen *ListenerQuery) EnterValue(ctx *parser.ValueContext) {}

// EnterMinus is called when entering the Minus production.
func (listen *ListenerQuery) EnterMinus(ctx *parser.MinusContext) {}

// EnterParentheses is called when entering the Parentheses production.
func (listen *ListenerQuery) EnterParentheses(ctx *parser.ParenthesesContext) {}

// EnterMultOrDiv is called when entering the MultOrDiv production.
func (listen *ListenerQuery) EnterMultOrDiv(ctx *parser.MultOrDivContext) {}

// EnterAsClause is called when entering the asClause production.
func (listen *ListenerQuery) EnterAsClause(ctx *parser.AsClauseContext) {}

// EnterAggregation is called when entering the aggregation production.
func (listen *ListenerQuery) EnterAggregation(ctx *parser.AggregationContext) {}

// EnterCount is called when entering the count production.
func (listen *ListenerQuery) EnterCount(ctx *parser.CountContext) {}

// EnterCalc is called when entering the calc production.
func (listen *ListenerQuery) EnterCalc(ctx *parser.CalcContext) {}

// EnterComp is called when entering the comp production.
func (listen *ListenerQuery) EnterComp(ctx *parser.CompContext) {}

// EnterCondition is called when entering the condition production.
func (listen *ListenerQuery) EnterCondition(ctx *parser.ConditionContext) {
	listen.currentMath = CreateMath()
}

// EnterSortOrder is called when entering the sortOrder production.
func (listen *ListenerQuery) EnterSortOrder(ctx *parser.SortOrderContext) {}

// EnterContinuousRange is called when entering the continuousRange production.
func (listen *ListenerQuery) EnterContinuousRange(ctx *parser.ContinuousRangeContext) {}

// EnterSparseRange is called when entering the sparseRange production.
func (listen *ListenerQuery) EnterSparseRange(ctx *parser.SparseRangeContext) {}

// EnterRange is called when entering the range production.
func (listen *ListenerQuery) EnterRange(ctx *parser.RangeContext) {}

// EnterCreate is called when entering the create production.
func (listen *ListenerQuery) EnterCreate(ctx *parser.CreateContext) {}

// EnterWith is called when entering the with production.
func (listen *ListenerQuery) EnterWith(ctx *parser.WithContext) {}

// EnterSelect is called when entering the select production.
func (listen *ListenerQuery) EnterSelect(ctx *parser.SelectContext) {
	selects := ctx.AllCalc()
	if len(selects) == 1 && selects[0].GetAll() != nil {
		listen.IsSelectAll = true
	} else {
		listen.IsSelectAll = false
		listen.ColumnNames = make([]string, len(selects))
		for i, calc := range selects {
			listen.ColumnNames[i] = calc.Nameable().Name().GetText()
		}
	}
}

// EnterFrom is called when entering the from production.
func (listen *ListenerQuery) EnterFrom(ctx *parser.FromContext) {
	listen.TableName = ctx.Name(0).GetText()
}

// EnterWhere is called when entering the where production.
func (listen *ListenerQuery) EnterWhere(ctx *parser.WhereContext) {
	listen.Conditions = make([]SimpleCondition, len(ctx.AllCondition()))
}

// EnterGroupby is called when entering the groupby production.
func (listen *ListenerQuery) EnterGroupby(ctx *parser.GroupbyContext) {}

// EnterOrderby is called when entering the orderby production.
func (listen *ListenerQuery) EnterOrderby(ctx *parser.OrderbyContext) {}

// EnterLimit is called when entering the limit production.
func (listen *ListenerQuery) EnterLimit(ctx *parser.LimitContext) {}

// EnterView is called when entering the view production.
func (listen *ListenerQuery) EnterView(ctx *parser.ViewContext) {}

// EnterCheck is called when entering the check production.
func (listen *ListenerQuery) EnterCheck(ctx *parser.CheckContext) {}

// EnterForeignkey is called when entering the foreignkey production.
func (listen *ListenerQuery) EnterForeignkey(ctx *parser.ForeignkeyContext) {}

// EnterConstraint is called when entering the constraint production.
func (listen *ListenerQuery) EnterConstraint(ctx *parser.ConstraintContext) {}

// EnterPrimarykey is called when entering the primarykey production.
func (listen *ListenerQuery) EnterPrimarykey(ctx *parser.PrimarykeyContext) {}

// EnterColumns is called when entering the columns production.
func (listen *ListenerQuery) EnterColumns(ctx *parser.ColumnsContext) {}

// EnterCreatetable is called when entering the table production.
func (listen *ListenerQuery) EnterCreatetable(ctx *parser.CreatetableContext) {}

// EnterCreateindex is called when entering the createindex production.
func (listen *ListenerQuery) EnterCreateindex(ctx *parser.CreateindexContext) {}

// EnterDrop is called when entering the drop production.
func (listen *ListenerQuery) EnterDrop(ctx *parser.DropContext) {}

// EnterDelete is called when entering the delete production.
func (listen *ListenerQuery) EnterDelete(ctx *parser.DeleteContext) {}

// EnterSet is called when entering the set production.
func (listen *ListenerQuery) EnterSet(ctx *parser.SetContext) {}

// EnterUpdate is called when entering the update production.
func (listen *ListenerQuery) EnterUpdate(ctx *parser.UpdateContext) {}

// EnterValues is called when entering the values production.
func (listen *ListenerQuery) EnterValues(ctx *parser.ValuesContext) {}

// EnterColumnNames is called when entering the columnNames production.
func (listen *ListenerQuery) EnterColumnNames(ctx *parser.ColumnNamesContext) {}

// EnterInsert is called when entering the insert production.
func (listen *ListenerQuery) EnterInsert(ctx *parser.InsertContext) {}

// EnterQuery is called when entering the query production.
func (listen *ListenerQuery) EnterQuery(ctx *parser.QueryContext) {}

// EnterStatement is called when entering the statement production.
func (listen *ListenerQuery) EnterStatement(ctx *parser.StatementContext) {}

// EnterStart is called when entering the start production.
func (listen *ListenerQuery) EnterStart(ctx *parser.StartContext) {}

// ExitName is called when exiting the name production.
func (listen *ListenerQuery) ExitName(ctx *parser.NameContext) {}

// ExitConstant is called when exiting the constant production.
func (listen *ListenerQuery) ExitConstant(ctx *parser.ConstantContext) {}

// ExitAggrFunc is called when exiting the aggrFunc production.
func (listen *ListenerQuery) ExitAggrFunc(ctx *parser.AggrFuncContext) {}

// ExitField is called when exiting the field production.
func (listen *ListenerQuery) ExitField(ctx *parser.FieldContext) {}

// ExitParameter is called when exiting the parameter production.
func (listen *ListenerQuery) ExitParameter(ctx *parser.ParameterContext) {}

// ExitNameable is called when exiting the nameable production.
func (listen *ListenerQuery) ExitNameable(ctx *parser.NameableContext) {}

// ExitKey is called when exiting the key production.
func (listen *ListenerQuery) ExitKey(ctx *parser.KeyContext) {}

// ExitAddOrSub is called when exiting the AddOrSub production.
func (listen *ListenerQuery) ExitAddOrSub(ctx *parser.AddOrSubContext) {
	AddOrSubHelper(ctx, listen.currentMath)
}

// ExitValue is called when exiting the Value production.
func (listen *ListenerQuery) ExitValue(ctx *parser.ValueContext) {
	if nameable := ctx.Nameable(); nameable != nil {
		listen.currentMath.Stack.Push(MakeNameableFromContext(nameable))
	} else { //Constant
		constant := MakeConstantFromContext(ctx.Constant())
		listen.currentMath.Stack.Push(constant)
	}
}

// ExitMinus is called when exiting the Minus production.
func (listen *ListenerQuery) ExitMinus(ctx *parser.MinusContext) {
	MinusHelper(ctx, listen.currentMath)
}

// ExitParentheses is called when exiting the Parentheses production.
func (listen *ListenerQuery) ExitParentheses(ctx *parser.ParenthesesContext) {}

// ExitMultOrDiv is called when exiting the MultOrDiv production.
func (listen *ListenerQuery) ExitMultOrDiv(ctx *parser.MultOrDivContext) {
	MultOrDivHelper(ctx, listen.currentMath)
}

// ExitAsClause is called when exiting the asClause production.
func (listen *ListenerQuery) ExitAsClause(ctx *parser.AsClauseContext) {}

// ExitAggregation is called when exiting the aggregation production.
func (listen *ListenerQuery) ExitAggregation(ctx *parser.AggregationContext) {}

// ExitCount is called when exiting the count production.
func (listen *ListenerQuery) ExitCount(ctx *parser.CountContext) {}

// ExitCalc is called when exiting the calc production.
func (listen *ListenerQuery) ExitCalc(ctx *parser.CalcContext) {}

// ExitComp is called when exiting the comp production.
func (listen *ListenerQuery) ExitComp(ctx *parser.CompContext) {}

// ExitCondition is called when exiting the condition production.
func (listen *ListenerQuery) ExitCondition(ctx *parser.ConditionContext) {
	cond := SimpleCondition{Math: listen.currentMath, ColumnName: ctx.Nameable().Name().GetText()}
	comp := ctx.Comp()
	if opType := comp.GetOpType(); opType != nil {
		cond.Op = StringOperatorToCondOp(opType)
	} else { //function name
		cond.Fun = MakeFunFromName(comp.Name().GetText())
		cond.Op = FUNCTION
	}
	listen.Conditions[listen.currIndex] = cond
	listen.currIndex++
}

// ExitSortOrder is called when exiting the sortOrder production.
func (listen *ListenerQuery) ExitSortOrder(ctx *parser.SortOrderContext) {}

// ExitContinuousRange is called when exiting the continuousRange production.
func (listen *ListenerQuery) ExitContinuousRange(ctx *parser.ContinuousRangeContext) {}

// ExitSparseRange is called when exiting the sparseRange production.
func (listen *ListenerQuery) ExitSparseRange(ctx *parser.SparseRangeContext) {}

// ExitRange is called when exiting the range production.
func (listen *ListenerQuery) ExitRange(ctx *parser.RangeContext) {}

// ExitCreate is called when exiting the create production.
func (listen *ListenerQuery) ExitCreate(ctx *parser.CreateContext) {}

// ExitCreate is called when exiting the create production.
func (listen *ListenerQuery) ExitCreatetable(ctx *parser.CreatetableContext) {}

// ExitCreateindex is called when exiting the createindex production.
func (listen *ListenerQuery) ExitCreateindex(ctx *parser.CreateindexContext) {}

// ExitWith is called when exiting the with production.
func (listen *ListenerQuery) ExitWith(ctx *parser.WithContext) {}

// ExitSelect is called when exiting the select production.
func (listen *ListenerQuery) ExitSelect(ctx *parser.SelectContext) {}

// ExitFrom is called when exiting the from production.
func (listen *ListenerQuery) ExitFrom(ctx *parser.FromContext) {}

// ExitWhere is called when exiting the where production.
func (listen *ListenerQuery) ExitWhere(ctx *parser.WhereContext) {}

// ExitGroupby is called when exiting the groupby production.
func (listen *ListenerQuery) ExitGroupby(ctx *parser.GroupbyContext) {}

// ExitOrderby is called when exiting the orderby production.
func (listen *ListenerQuery) ExitOrderby(ctx *parser.OrderbyContext) {}

// ExitLimit is called when exiting the limit production.
func (listen *ListenerQuery) ExitLimit(ctx *parser.LimitContext) {}

// ExitView is called when exiting the view production.
func (listen *ListenerQuery) ExitView(ctx *parser.ViewContext) {}

// ExitCheck is called when exiting the check production.
func (listen *ListenerQuery) ExitCheck(ctx *parser.CheckContext) {}

// ExitForeignkey is called when exiting the foreignkey production.
func (listen *ListenerQuery) ExitForeignkey(ctx *parser.ForeignkeyContext) {}

// ExitPrimarykey is called when exiting the primarykey production.
func (listen *ListenerQuery) ExitPrimarykey(ctx *parser.PrimarykeyContext) {}

// ExitConstraint is called when exiting the constraint production.
func (listen *ListenerQuery) ExitConstraint(ctx *parser.ConstraintContext) {}

// ExitColumns is called when exiting the columns production.
func (listen *ListenerQuery) ExitColumns(ctx *parser.ColumnsContext) {}

// ExitTable is called when exiting the table production.
func (listen *ListenerQuery) ExitTable(ctx *parser.CreatetableContext) {}

// ExitDrop is called when exiting the drop production.
func (listen *ListenerQuery) ExitDrop(ctx *parser.DropContext) {}

// ExitDelete is called when exiting the delete production.
func (listen *ListenerQuery) ExitDelete(ctx *parser.DeleteContext) {}

// ExitSet is called when exiting the set production.
func (listen *ListenerQuery) ExitSet(ctx *parser.SetContext) {}

// ExitUpdate is called when exiting the update production.
func (listen *ListenerQuery) ExitUpdate(ctx *parser.UpdateContext) {}

// ExitValues is called when exiting the values production.
func (listen *ListenerQuery) ExitValues(ctx *parser.ValuesContext) {}

// ExitColumnNames is called when exiting the columnNames production.
func (listen *ListenerQuery) ExitColumnNames(ctx *parser.ColumnNamesContext) {}

// ExitInsert is called when exiting the insert production.
func (listen *ListenerQuery) ExitInsert(ctx *parser.InsertContext) {}

// ExitQuery is called when exiting the query production.
func (listen *ListenerQuery) ExitQuery(ctx *parser.QueryContext) {}

// ExitStatement is called when exiting the statement production.
func (listen *ListenerQuery) ExitStatement(ctx *parser.StatementContext) {}

// ExitStart is called when exiting the start production.
func (listen *ListenerQuery) ExitStart(ctx *parser.StartContext) {
	fmt.Println("[QueryL]Finished processing query text.")
	fmt.Println("TableName:", listen.TableName)
	fmt.Println("Columns:", listen.ColumnNames)
	fmt.Println("isSelectAll (*):", listen.IsSelectAll)
	fmt.Print("Conditions: [")
	for _, cond := range listen.Conditions {
		fmt.Print(cond.ToString(), ",")
	}
	fmt.Println("]")
}
