package sql

import (
	"sqlToKeyValue/src/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type ListenerDropTable struct {
	Parse *parser.ViewSQLParser

	TableName string
}

func MakeDropTableListener(parse *parser.ViewSQLParser) *ListenerDropTable {
	return &ListenerDropTable{Parse: parse}
}

// VisitTerminal is called when a terminal node is visited.
func (listen *ListenerDropTable) VisitTerminal(node antlr.TerminalNode) {
	//fmt.Println("[DropTableL]Terminal", node.GetText())
}

// VisitErrorNode is called when an error node is visited.
func (listen *ListenerDropTable) VisitErrorNode(node antlr.ErrorNode) {
	//fmt.Println("[DropTableL]Error")
}

// EnterEveryRule is called when any rule is entered.
func (listen *ListenerDropTable) EnterEveryRule(ctx antlr.ParserRuleContext) {
	//fmt.Println("[DropTableL]Entered rule", ctx.GetText())
}

// ExitEveryRule is called when any rule is exited.
func (listen *ListenerDropTable) ExitEveryRule(ctx antlr.ParserRuleContext) {
	//fmt.Println("[DropTableL]Exit rule", ctx.GetText())
}

// EnterName is called when entering the constant production.
func (listen *ListenerDropTable) EnterName(ctx *parser.NameContext) {
	listen.TableName = ctx.GetText()
}

// EnterConstant is called when entering the constant production.
func (listen *ListenerDropTable) EnterConstant(ctx *parser.ConstantContext) {}

// EnterAggrFunc is called when entering the aggrFunc production.
func (listen *ListenerDropTable) EnterAggrFunc(ctx *parser.AggrFuncContext) {}

// EnterField is called when entering the field production.
func (listen *ListenerDropTable) EnterField(ctx *parser.FieldContext) {}

// EnterParameter is called when entering the parameter production.
func (listen *ListenerDropTable) EnterParameter(ctx *parser.ParameterContext) {}

// EnterNameable is called when entering the nameable production.
func (listen *ListenerDropTable) EnterNameable(ctx *parser.NameableContext) {}

// EnterKey is called when entering the key production.
func (listen *ListenerDropTable) EnterKey(ctx *parser.KeyContext) {}

// EnterAddOrSub is called when entering the AddOrSub production.
func (listen *ListenerDropTable) EnterAddOrSub(ctx *parser.AddOrSubContext) {}

// EnterValue is called when entering the Value production.
func (listen *ListenerDropTable) EnterValue(ctx *parser.ValueContext) {}

// EnterMinus is called when entering the Minus production.
func (listen *ListenerDropTable) EnterMinus(ctx *parser.MinusContext) {}

// EnterParentheses is called when entering the Parentheses production.
func (listen *ListenerDropTable) EnterParentheses(ctx *parser.ParenthesesContext) {}

// EnterMultOrDiv is called when entering the MultOrDiv production.
func (listen *ListenerDropTable) EnterMultOrDiv(ctx *parser.MultOrDivContext) {}

// EnterAsClause is called when entering the asClause production.
func (listen *ListenerDropTable) EnterAsClause(ctx *parser.AsClauseContext) {}

// EnterAggregation is called when entering the aggregation production.
func (listen *ListenerDropTable) EnterAggregation(ctx *parser.AggregationContext) {}

// EnterCount is called when entering the count production.
func (listen *ListenerDropTable) EnterCount(ctx *parser.CountContext) {}

// EnterCalc is called when entering the calc production.
func (listen *ListenerDropTable) EnterCalc(ctx *parser.CalcContext) {}

// EnterComp is called when entering the comp production.
func (listen *ListenerDropTable) EnterComp(ctx *parser.CompContext) {}

// EnterCondition is called when entering the condition production.
func (listen *ListenerDropTable) EnterCondition(ctx *parser.ConditionContext) {}

// EnterSortOrder is called when entering the sortOrder production.
func (listen *ListenerDropTable) EnterSortOrder(ctx *parser.SortOrderContext) {}

// EnterContinuousRange is called when entering the continuousRange production.
func (listen *ListenerDropTable) EnterContinuousRange(ctx *parser.ContinuousRangeContext) {}

// EnterSparseRange is called when entering the sparseRange production.
func (listen *ListenerDropTable) EnterSparseRange(ctx *parser.SparseRangeContext) {}

// EnterRange is called when entering the range production.
func (listen *ListenerDropTable) EnterRange(ctx *parser.RangeContext) {}

// EnterCreate is called when entering the create production.
func (listen *ListenerDropTable) EnterCreate(ctx *parser.CreateContext) {}

// EnterWith is called when entering the with production.
func (listen *ListenerDropTable) EnterWith(ctx *parser.WithContext) {}

// EnterSelect is called when entering the select production.
func (listen *ListenerDropTable) EnterSelect(ctx *parser.SelectContext) {}

// EnterFrom is called when entering the from production.
func (listen *ListenerDropTable) EnterFrom(ctx *parser.FromContext) {}

// EnterWhere is called when entering the where production.
func (listen *ListenerDropTable) EnterWhere(ctx *parser.WhereContext) {}

// EnterGroupby is called when entering the groupby production.
func (listen *ListenerDropTable) EnterGroupby(ctx *parser.GroupbyContext) {}

// EnterOrderby is called when entering the orderby production.
func (listen *ListenerDropTable) EnterOrderby(ctx *parser.OrderbyContext) {}

// EnterLimit is called when entering the limit production.
func (listen *ListenerDropTable) EnterLimit(ctx *parser.LimitContext) {}

// EnterView is called when entering the view production.
func (listen *ListenerDropTable) EnterView(ctx *parser.ViewContext) {}

// EnterCheck is called when entering the check production.
func (listen *ListenerDropTable) EnterCheck(ctx *parser.CheckContext) {}

// EnterForeignkey is called when entering the foreignkey production.
func (listen *ListenerDropTable) EnterForeignkey(ctx *parser.ForeignkeyContext) {}

// EnterConstraint is called when entering the constraint production.
func (listen *ListenerDropTable) EnterConstraint(ctx *parser.ConstraintContext) {}

// EnterPrimarykey is called when entering the primarykey production.
func (listen *ListenerDropTable) EnterPrimarykey(ctx *parser.PrimarykeyContext) {}

// EnterColumns is called when entering the columns production.
func (listen *ListenerDropTable) EnterColumns(ctx *parser.ColumnsContext) {}

// EnterCreatetable is called when entering the table production.
func (listen *ListenerDropTable) EnterCreatetable(ctx *parser.CreatetableContext) {}

// EnterCreateindex is called when entering the createindex production.
func (listen *ListenerDropTable) EnterCreateindex(ctx *parser.CreateindexContext) {}

// EnterDrop is called when entering the drop production.
func (listen *ListenerDropTable) EnterDrop(ctx *parser.DropContext) {}

// EnterDelete is called when entering the delete production.
func (listen *ListenerDropTable) EnterDelete(ctx *parser.DeleteContext) {}

// EnterSet is called when entering the set production.
func (listen *ListenerDropTable) EnterSet(ctx *parser.SetContext) {}

// EnterUpdate is called when entering the update production.
func (listen *ListenerDropTable) EnterUpdate(ctx *parser.UpdateContext) {}

// EnterValues is called when entering the values production.
func (listen *ListenerDropTable) EnterValues(ctx *parser.ValuesContext) {}

// EnterColumnNames is called when entering the columnNames production.
func (listen *ListenerDropTable) EnterColumnNames(ctx *parser.ColumnNamesContext) {}

// EnterInsert is called when entering the insert production.
func (listen *ListenerDropTable) EnterInsert(ctx *parser.InsertContext) {}

// EnterQuery is called when entering the query production.
func (listen *ListenerDropTable) EnterQuery(ctx *parser.QueryContext) {}

// EnterStatement is called when entering the statement production.
func (listen *ListenerDropTable) EnterStatement(ctx *parser.StatementContext) {}

// EnterStart is called when entering the start production.
func (listen *ListenerDropTable) EnterStart(ctx *parser.StartContext) {}

// ExitName is called when exiting the name production.
func (listen *ListenerDropTable) ExitName(ctx *parser.NameContext) {}

// ExitConstant is called when exiting the constant production.
func (listen *ListenerDropTable) ExitConstant(ctx *parser.ConstantContext) {}

// ExitAggrFunc is called when exiting the aggrFunc production.
func (listen *ListenerDropTable) ExitAggrFunc(ctx *parser.AggrFuncContext) {}

// ExitField is called when exiting the field production.
func (listen *ListenerDropTable) ExitField(ctx *parser.FieldContext) {}

// ExitParameter is called when exiting the parameter production.
func (listen *ListenerDropTable) ExitParameter(ctx *parser.ParameterContext) {}

// ExitNameable is called when exiting the nameable production.
func (listen *ListenerDropTable) ExitNameable(ctx *parser.NameableContext) {}

// ExitKey is called when exiting the key production.
func (listen *ListenerDropTable) ExitKey(ctx *parser.KeyContext) {}

// ExitAddOrSub is called when exiting the AddOrSub production.
func (listen *ListenerDropTable) ExitAddOrSub(ctx *parser.AddOrSubContext) {}

// ExitValue is called when exiting the Value production.
func (listen *ListenerDropTable) ExitValue(ctx *parser.ValueContext) {}

// ExitMinus is called when exiting the Minus production.
func (listen *ListenerDropTable) ExitMinus(ctx *parser.MinusContext) {}

// ExitParentheses is called when exiting the Parentheses production.
func (listen *ListenerDropTable) ExitParentheses(ctx *parser.ParenthesesContext) {}

// ExitMultOrDiv is called when exiting the MultOrDiv production.
func (listen *ListenerDropTable) ExitMultOrDiv(ctx *parser.MultOrDivContext) {}

// ExitAsClause is called when exiting the asClause production.
func (listen *ListenerDropTable) ExitAsClause(ctx *parser.AsClauseContext) {}

// ExitAggregation is called when exiting the aggregation production.
func (listen *ListenerDropTable) ExitAggregation(ctx *parser.AggregationContext) {}

// ExitCount is called when exiting the count production.
func (listen *ListenerDropTable) ExitCount(ctx *parser.CountContext) {}

// ExitCalc is called when exiting the calc production.
func (listen *ListenerDropTable) ExitCalc(ctx *parser.CalcContext) {}

// ExitComp is called when exiting the comp production.
func (listen *ListenerDropTable) ExitComp(ctx *parser.CompContext) {}

// ExitCondition is called when exiting the condition production.
func (listen *ListenerDropTable) ExitCondition(ctx *parser.ConditionContext) {}

// ExitSortOrder is called when exiting the sortOrder production.
func (listen *ListenerDropTable) ExitSortOrder(ctx *parser.SortOrderContext) {}

// ExitContinuousRange is called when exiting the continuousRange production.
func (listen *ListenerDropTable) ExitContinuousRange(ctx *parser.ContinuousRangeContext) {}

// ExitSparseRange is called when exiting the sparseRange production.
func (listen *ListenerDropTable) ExitSparseRange(ctx *parser.SparseRangeContext) {}

// ExitRange is called when exiting the range production.
func (listen *ListenerDropTable) ExitRange(ctx *parser.RangeContext) {}

// ExitCreate is called when exiting the create production.
func (listen *ListenerDropTable) ExitCreate(ctx *parser.CreateContext) {}

// ExitCreate is called when exiting the create production.
func (listen *ListenerDropTable) ExitCreatetable(ctx *parser.CreatetableContext) {}

// ExitCreateindex is called when exiting the createindex production.
func (listen *ListenerDropTable) ExitCreateindex(ctx *parser.CreateindexContext) {}

// ExitWith is called when exiting the with production.
func (listen *ListenerDropTable) ExitWith(ctx *parser.WithContext) {}

// ExitSelect is called when exiting the select production.
func (listen *ListenerDropTable) ExitSelect(ctx *parser.SelectContext) {}

// ExitFrom is called when exiting the from production.
func (listen *ListenerDropTable) ExitFrom(ctx *parser.FromContext) {}

// ExitWhere is called when exiting the where production.
func (listen *ListenerDropTable) ExitWhere(ctx *parser.WhereContext) {}

// ExitGroupby is called when exiting the groupby production.
func (listen *ListenerDropTable) ExitGroupby(ctx *parser.GroupbyContext) {}

// ExitOrderby is called when exiting the orderby production.
func (listen *ListenerDropTable) ExitOrderby(ctx *parser.OrderbyContext) {}

// ExitLimit is called when exiting the limit production.
func (listen *ListenerDropTable) ExitLimit(ctx *parser.LimitContext) {}

// ExitView is called when exiting the view production.
func (listen *ListenerDropTable) ExitView(ctx *parser.ViewContext) {}

// ExitCheck is called when exiting the check production.
func (listen *ListenerDropTable) ExitCheck(ctx *parser.CheckContext) {}

// ExitForeignkey is called when exiting the foreignkey production.
func (listen *ListenerDropTable) ExitForeignkey(ctx *parser.ForeignkeyContext) {}

// ExitPrimarykey is called when exiting the primarykey production.
func (listen *ListenerDropTable) ExitPrimarykey(ctx *parser.PrimarykeyContext) {}

// ExitConstraint is called when exiting the constraint production.
func (listen *ListenerDropTable) ExitConstraint(ctx *parser.ConstraintContext) {}

// ExitColumns is called when exiting the columns production.
func (listen *ListenerDropTable) ExitColumns(ctx *parser.ColumnsContext) {}

// ExitTable is called when exiting the table production.
func (listen *ListenerDropTable) ExitTable(ctx *parser.CreatetableContext) {}

// ExitDrop is called when exiting the drop production.
func (listen *ListenerDropTable) ExitDrop(ctx *parser.DropContext) {}

// ExitDelete is called when exiting the delete production.
func (listen *ListenerDropTable) ExitDelete(ctx *parser.DeleteContext) {}

// ExitSet is called when exiting the set production.
func (listen *ListenerDropTable) ExitSet(ctx *parser.SetContext) {}

// ExitUpdate is called when exiting the update production.
func (listen *ListenerDropTable) ExitUpdate(ctx *parser.UpdateContext) {}

// ExitValues is called when exiting the values production.
func (listen *ListenerDropTable) ExitValues(ctx *parser.ValuesContext) {}

// ExitColumnNames is called when exiting the columnNames production.
func (listen *ListenerDropTable) ExitColumnNames(ctx *parser.ColumnNamesContext) {}

// ExitInsert is called when exiting the insert production.
func (listen *ListenerDropTable) ExitInsert(ctx *parser.InsertContext) {}

// ExitQuery is called when exiting the query production.
func (listen *ListenerDropTable) ExitQuery(ctx *parser.QueryContext) {}

// ExitStatement is called when exiting the statement production.
func (listen *ListenerDropTable) ExitStatement(ctx *parser.StatementContext) {}

// ExitStart is called when exiting the start production.
func (listen *ListenerDropTable) ExitStart(ctx *parser.StartContext) {}
