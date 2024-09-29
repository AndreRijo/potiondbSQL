package sql

import (
	"fmt"

	"sqlToKeyValue/src/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type ListenerDelete struct {
	Parse       *parser.ViewSQLParser
	TableName   string
	Conditions  []SimpleCondition
	currentMath Math
	currIndex   int
}

func MakeDeleteListener(parse *parser.ViewSQLParser) *ListenerDelete {
	return &ListenerDelete{Parse: parse, currIndex: 0}
}

// VisitTerminal is called when a terminal node is visited.
func (listen *ListenerDelete) VisitTerminal(node antlr.TerminalNode) {
}

// VisitErrorNode is called when an error node is visited.
func (listen *ListenerDelete) VisitErrorNode(node antlr.ErrorNode) {
}

// EnterEveryRule is called when any rule is entered.
func (listen *ListenerDelete) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println("[DeleteL]Entered rule")
}

// ExitEveryRule is called when any rule is exited.
func (listen *ListenerDelete) ExitEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println("[DeleteL]Exit rule")
}

// EnterName is called when entering the constant production.
func (listen *ListenerDelete) EnterName(ctx *parser.NameContext) {}

// EnterConstant is called when entering the constant production.
func (listen *ListenerDelete) EnterConstant(ctx *parser.ConstantContext) {}

// EnterAggrFunc is called when entering the aggrFunc production.
func (listen *ListenerDelete) EnterAggrFunc(ctx *parser.AggrFuncContext) {}

// EnterField is called when entering the field production.
func (listen *ListenerDelete) EnterField(ctx *parser.FieldContext) {}

// EnterParameter is called when entering the parameter production.
func (listen *ListenerDelete) EnterParameter(ctx *parser.ParameterContext) {}

// EnterNameable is called when entering the nameable production.
func (listen *ListenerDelete) EnterNameable(ctx *parser.NameableContext) {}

// EnterKey is called when entering the key production.
func (listen *ListenerDelete) EnterKey(ctx *parser.KeyContext) {}

// EnterAddOrSub is called when entering the AddOrSub production.
func (listen *ListenerDelete) EnterAddOrSub(ctx *parser.AddOrSubContext) {}

// EnterValue is called when entering the Value production.
func (listen *ListenerDelete) EnterValue(ctx *parser.ValueContext) {}

// EnterMinus is called when entering the Minus production.
func (listen *ListenerDelete) EnterMinus(ctx *parser.MinusContext) {}

// EnterParentheses is called when entering the Parentheses production.
func (listen *ListenerDelete) EnterParentheses(ctx *parser.ParenthesesContext) {}

// EnterMultOrDiv is called when entering the MultOrDiv production.
func (listen *ListenerDelete) EnterMultOrDiv(ctx *parser.MultOrDivContext) {}

// EnterAsClause is called when entering the asClause production.
func (listen *ListenerDelete) EnterAsClause(ctx *parser.AsClauseContext) {}

// EnterAggregation is called when entering the aggregation production.
func (listen *ListenerDelete) EnterAggregation(ctx *parser.AggregationContext) {}

// EnterCount is called when entering the count production.
func (listen *ListenerDelete) EnterCount(ctx *parser.CountContext) {}

// EnterCalc is called when entering the calc production.
func (listen *ListenerDelete) EnterCalc(ctx *parser.CalcContext) {}

// EnterComp is called when entering the comp production.
func (listen *ListenerDelete) EnterComp(ctx *parser.CompContext) {}

// EnterCondition is called when entering the condition production.
func (listen *ListenerDelete) EnterCondition(ctx *parser.ConditionContext) {
	listen.currentMath = CreateMath()
}

// EnterSortOrder is called when entering the sortOrder production.
func (listen *ListenerDelete) EnterSortOrder(ctx *parser.SortOrderContext) {}

// EnterContinuousRange is called when entering the continuousRange production.
func (listen *ListenerDelete) EnterContinuousRange(ctx *parser.ContinuousRangeContext) {}

// EnterSparseRange is called when entering the sparseRange production.
func (listen *ListenerDelete) EnterSparseRange(ctx *parser.SparseRangeContext) {}

// EnterRange is called when entering the range production.
func (listen *ListenerDelete) EnterRange(ctx *parser.RangeContext) {}

// EnterCreate is called when entering the create production.
func (listen *ListenerDelete) EnterCreate(ctx *parser.CreateContext) {}

// EnterWith is called when entering the with production.
func (listen *ListenerDelete) EnterWith(ctx *parser.WithContext) {}

// EnterSelect is called when entering the select production.
func (listen *ListenerDelete) EnterSelect(ctx *parser.SelectContext) {}

// EnterFrom is called when entering the from production.
func (listen *ListenerDelete) EnterFrom(ctx *parser.FromContext) {}

// EnterWhere is called when entering the where production.
func (listen *ListenerDelete) EnterWhere(ctx *parser.WhereContext) {
	//TODO: Probably some special processing for primary key
	listen.Conditions = make([]SimpleCondition, len(ctx.AllCondition()))
}

// EnterGroupby is called when entering the groupby production.
func (listen *ListenerDelete) EnterGroupby(ctx *parser.GroupbyContext) {}

// EnterOrderby is called when entering the orderby production.
func (listen *ListenerDelete) EnterOrderby(ctx *parser.OrderbyContext) {}

// EnterLimit is called when entering the limit production.
func (listen *ListenerDelete) EnterLimit(ctx *parser.LimitContext) {}

// EnterView is called when entering the view production.
func (listen *ListenerDelete) EnterView(ctx *parser.ViewContext) {}

// EnterCheck is called when entering the check production.
func (listen *ListenerDelete) EnterCheck(ctx *parser.CheckContext) {}

// EnterForeignkey is called when entering the foreignkey production.
func (listen *ListenerDelete) EnterForeignkey(ctx *parser.ForeignkeyContext) {}

// EnterPrimarykey is called when entering the primarykey production.
func (listen *ListenerDelete) EnterPrimarykey(ctx *parser.PrimarykeyContext) {}

// EnterConstraint is called when entering the constraint production.
func (listen *ListenerDelete) EnterConstraint(ctx *parser.ConstraintContext) {}

// EnterColumns is called when entering the columns production.
func (listen *ListenerDelete) EnterColumns(ctx *parser.ColumnsContext) {}

// EnterCreatetable is called when entering the table production.
func (listen *ListenerDelete) EnterCreatetable(ctx *parser.CreatetableContext) {}

// EnterCreateindex is called when entering the createindex production.
func (listen *ListenerDelete) EnterCreateindex(ctx *parser.CreateindexContext) {}

// EnterDrop is called when entering the drop production.
func (listen *ListenerDelete) EnterDrop(ctx *parser.DropContext) {}

// EnterDelete is called when entering the delete production.
func (listen *ListenerDelete) EnterDelete(ctx *parser.DeleteContext) {
	listen.TableName = ctx.Name().GetText()
}

// EnterSet is called when entering the set production.
func (listen *ListenerDelete) EnterSet(ctx *parser.SetContext) {}

// EnterUpdate is called when entering the update production.
func (listen *ListenerDelete) EnterUpdate(ctx *parser.UpdateContext) {}

// EnterValues is called when entering the values production.
func (listen *ListenerDelete) EnterValues(ctx *parser.ValuesContext) {}

// EnterColumnNames is called when entering the columnNames production.
func (listen *ListenerDelete) EnterColumnNames(ctx *parser.ColumnNamesContext) {}

// EnterInsert is called when entering the insert production.
func (listen *ListenerDelete) EnterInsert(ctx *parser.InsertContext) {}

// EnterQuery is called when entering the query production.
func (listen *ListenerDelete) EnterQuery(ctx *parser.QueryContext) {}

// EnterStatement is called when entering the statement production.
func (listen *ListenerDelete) EnterStatement(ctx *parser.StatementContext) {}

// EnterStart is called when entering the start production.
func (listen *ListenerDelete) EnterStart(ctx *parser.StartContext) {}

// ExitName is called when exiting the name production.
func (listen *ListenerDelete) ExitName(ctx *parser.NameContext) {}

// ExitConstant is called when exiting the constant production.
func (listen *ListenerDelete) ExitConstant(ctx *parser.ConstantContext) {}

// ExitAggrFunc is called when exiting the aggrFunc production.
func (listen *ListenerDelete) ExitAggrFunc(ctx *parser.AggrFuncContext) {}

// ExitField is called when exiting the field production.
func (listen *ListenerDelete) ExitField(ctx *parser.FieldContext) {}

// ExitParameter is called when exiting the parameter production.
func (listen *ListenerDelete) ExitParameter(ctx *parser.ParameterContext) {}

// ExitNameable is called when exiting the nameable production.
func (listen *ListenerDelete) ExitNameable(ctx *parser.NameableContext) {}

// ExitKey is called when exiting the key production.
func (listen *ListenerDelete) ExitKey(ctx *parser.KeyContext) {}

// ExitAddOrSub is called when exiting the AddOrSub production.
func (listen *ListenerDelete) ExitAddOrSub(ctx *parser.AddOrSubContext) {
	AddOrSubHelper(ctx, listen.currentMath)
}

// ExitValue is called when exiting the Value production.
func (listen *ListenerDelete) ExitValue(ctx *parser.ValueContext) {
	if nameable := ctx.Nameable(); nameable != nil {
		listen.currentMath.Stack.Push(MakeNameableFromContext(nameable))
	} else { //Constant
		constant := MakeConstantFromContext(ctx.Constant())
		listen.currentMath.Stack.Push(constant)
	}
}

// ExitMinus is called when exiting the Minus production.
func (listen *ListenerDelete) ExitMinus(ctx *parser.MinusContext) {
	MinusHelper(ctx, listen.currentMath)
}

// ExitParentheses is called when exiting the Parentheses production.
func (listen *ListenerDelete) ExitParentheses(ctx *parser.ParenthesesContext) {}

// ExitMultOrDiv is called when exiting the MultOrDiv production.
func (listen *ListenerDelete) ExitMultOrDiv(ctx *parser.MultOrDivContext) {
	MultOrDivHelper(ctx, listen.currentMath)
}

// ExitAsClause is called when exiting the asClause production.
func (listen *ListenerDelete) ExitAsClause(ctx *parser.AsClauseContext) {}

// ExitAggregation is called when exiting the aggregation production.
func (listen *ListenerDelete) ExitAggregation(ctx *parser.AggregationContext) {}

// ExitCount is called when exiting the count production.
func (listen *ListenerDelete) ExitCount(ctx *parser.CountContext) {}

// ExitCalc is called when exiting the calc production.
func (listen *ListenerDelete) ExitCalc(ctx *parser.CalcContext) {}

// ExitComp is called when exiting the comp production.
func (listen *ListenerDelete) ExitComp(ctx *parser.CompContext) {}

// ExitCondition is called when exiting the condition production.
func (listen *ListenerDelete) ExitCondition(ctx *parser.ConditionContext) {
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
func (listen *ListenerDelete) ExitSortOrder(ctx *parser.SortOrderContext) {}

// ExitContinuousRange is called when exiting the continuousRange production.
func (listen *ListenerDelete) ExitContinuousRange(ctx *parser.ContinuousRangeContext) {}

// ExitSparseRange is called when exiting the sparseRange production.
func (listen *ListenerDelete) ExitSparseRange(ctx *parser.SparseRangeContext) {}

// ExitRange is called when exiting the range production.
func (listen *ListenerDelete) ExitRange(ctx *parser.RangeContext) {}

// ExitCreate is called when exiting the create production.
func (listen *ListenerDelete) ExitCreate(ctx *parser.CreateContext) {}

// ExitCreate is called when exiting the create production.
func (listen *ListenerDelete) ExitCreatetable(ctx *parser.CreatetableContext) {}

// ExitCreateindex is called when exiting the createindex production.
func (listen *ListenerDelete) ExitCreateindex(ctx *parser.CreateindexContext) {}

// ExitWith is called when exiting the with production.
func (listen *ListenerDelete) ExitWith(ctx *parser.WithContext) {}

// ExitSelect is called when exiting the select production.
func (listen *ListenerDelete) ExitSelect(ctx *parser.SelectContext) {}

// ExitFrom is called when exiting the from production.
func (listen *ListenerDelete) ExitFrom(ctx *parser.FromContext) {}

// ExitWhere is called when exiting the where production.
func (listen *ListenerDelete) ExitWhere(ctx *parser.WhereContext) {}

// ExitGroupby is called when exiting the groupby production.
func (listen *ListenerDelete) ExitGroupby(ctx *parser.GroupbyContext) {}

// ExitOrderby is called when exiting the orderby production.
func (listen *ListenerDelete) ExitOrderby(ctx *parser.OrderbyContext) {}

// ExitLimit is called when exiting the limit production.
func (listen *ListenerDelete) ExitLimit(ctx *parser.LimitContext) {}

// ExitView is called when exiting the view production.
func (listen *ListenerDelete) ExitView(ctx *parser.ViewContext) {}

// ExitCheck is called when exiting the check production.
func (listen *ListenerDelete) ExitCheck(ctx *parser.CheckContext) {}

// ExitForeignkey is called when exiting the foreignkey production.
func (listen *ListenerDelete) ExitForeignkey(ctx *parser.ForeignkeyContext) {}

// ExitPrimarykey is called when exiting the primarykey production.
func (listen *ListenerDelete) ExitPrimarykey(ctx *parser.PrimarykeyContext) {}

// ExitConstraint is called when exiting the constraint production.
func (listen *ListenerDelete) ExitConstraint(ctx *parser.ConstraintContext) {}

// ExitColumns is called when exiting the columns production.
func (listen *ListenerDelete) ExitColumns(ctx *parser.ColumnsContext) {}

// ExitTable is called when exiting the table production.
func (listen *ListenerDelete) ExitTable(ctx *parser.CreatetableContext) {}

// ExitDrop is called when exiting the drop production.
func (listen *ListenerDelete) ExitDrop(ctx *parser.DropContext) {}

// ExitDelete is called when exiting the delete production.
func (listen *ListenerDelete) ExitDelete(ctx *parser.DeleteContext) {}

// ExitSet is called when exiting the set production.
func (listen *ListenerDelete) ExitSet(ctx *parser.SetContext) {}

// ExitUpdate is called when exiting the update production.
func (listen *ListenerDelete) ExitUpdate(ctx *parser.UpdateContext) {}

// ExitValues is called when exiting the values production.
func (listen *ListenerDelete) ExitValues(ctx *parser.ValuesContext) {}

// ExitColumnNames is called when exiting the columnNames production.
func (listen *ListenerDelete) ExitColumnNames(ctx *parser.ColumnNamesContext) {}

// ExitInsert is called when exiting the insert production.
func (listen *ListenerDelete) ExitInsert(ctx *parser.InsertContext) {}

// ExitQuery is called when exiting the query production.
func (listen *ListenerDelete) ExitQuery(ctx *parser.QueryContext) {}

// ExitStatement is called when exiting the statement production.
func (listen *ListenerDelete) ExitStatement(ctx *parser.StatementContext) {}

// ExitStart is called when exiting the start production.
func (listen *ListenerDelete) ExitStart(ctx *parser.StartContext) {
	fmt.Println("[DeleteL]Finished processing text.")
	fmt.Println("TableName:", listen.TableName)
	fmt.Print("Conditions: [")
	for _, cond := range listen.Conditions {
		fmt.Print(cond.ToString(), ",")
	}
	fmt.Println("]")
}
