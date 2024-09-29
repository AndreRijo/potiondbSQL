package sql

import (
	"fmt"

	"sqlToKeyValue/src/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type ListenerCreateIndex struct {
	Parse                            *parser.ViewSQLParser
	IndexName, TableName, ColumnName string
}

func MakeCreateIndexListener(parse *parser.ViewSQLParser) *ListenerCreateIndex {
	return &ListenerCreateIndex{Parse: parse}
}

// VisitTerminal is called when a terminal node is visited.
func (listen *ListenerCreateIndex) VisitTerminal(node antlr.TerminalNode) {
}

// VisitErrorNode is called when an error node is visited.
func (listen *ListenerCreateIndex) VisitErrorNode(node antlr.ErrorNode) {
}

// EnterEveryRule is called when any rule is entered.
func (listen *ListenerCreateIndex) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println("[CreateIndexL]Entered rule")
}

// ExitEveryRule is called when any rule is exited.
func (listen *ListenerCreateIndex) ExitEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println("[CreateIndexL]Exit rule")
}

// EnterName is called when entering the constant production.
func (listen *ListenerCreateIndex) EnterName(ctx *parser.NameContext) {}

// EnterConstant is called when entering the constant production.
func (listen *ListenerCreateIndex) EnterConstant(ctx *parser.ConstantContext) {}

// EnterAggrFunc is called when entering the aggrFunc production.
func (listen *ListenerCreateIndex) EnterAggrFunc(ctx *parser.AggrFuncContext) {}

// EnterField is called when entering the field production.
func (listen *ListenerCreateIndex) EnterField(ctx *parser.FieldContext) {}

// EnterParameter is called when entering the parameter production.
func (listen *ListenerCreateIndex) EnterParameter(ctx *parser.ParameterContext) {}

// EnterNameable is called when entering the nameable production.
func (listen *ListenerCreateIndex) EnterNameable(ctx *parser.NameableContext) {}

// EnterKey is called when entering the key production.
func (listen *ListenerCreateIndex) EnterKey(ctx *parser.KeyContext) {}

// EnterAddOrSub is called when entering the AddOrSub production.
func (listen *ListenerCreateIndex) EnterAddOrSub(ctx *parser.AddOrSubContext) {}

// EnterValue is called when entering the Value production.
func (listen *ListenerCreateIndex) EnterValue(ctx *parser.ValueContext) {}

// EnterMinus is called when entering the Minus production.
func (listen *ListenerCreateIndex) EnterMinus(ctx *parser.MinusContext) {}

// EnterParentheses is called when entering the Parentheses production.
func (listen *ListenerCreateIndex) EnterParentheses(ctx *parser.ParenthesesContext) {}

// EnterMultOrDiv is called when entering the MultOrDiv production.
func (listen *ListenerCreateIndex) EnterMultOrDiv(ctx *parser.MultOrDivContext) {}

// EnterAsClause is called when entering the asClause production.
func (listen *ListenerCreateIndex) EnterAsClause(ctx *parser.AsClauseContext) {}

// EnterAggregation is called when entering the aggregation production.
func (listen *ListenerCreateIndex) EnterAggregation(ctx *parser.AggregationContext) {}

// EnterCount is called when entering the count production.
func (listen *ListenerCreateIndex) EnterCount(ctx *parser.CountContext) {}

// EnterCalc is called when entering the calc production.
func (listen *ListenerCreateIndex) EnterCalc(ctx *parser.CalcContext) {}

// EnterComp is called when entering the comp production.
func (listen *ListenerCreateIndex) EnterComp(ctx *parser.CompContext) {}

// EnterCondition is called when entering the condition production.
func (listen *ListenerCreateIndex) EnterCondition(ctx *parser.ConditionContext) {}

// EnterSortOrder is called when entering the sortOrder production.
func (listen *ListenerCreateIndex) EnterSortOrder(ctx *parser.SortOrderContext) {}

// EnterContinuousRange is called when entering the continuousRange production.
func (listen *ListenerCreateIndex) EnterContinuousRange(ctx *parser.ContinuousRangeContext) {}

// EnterSparseRange is called when entering the sparseRange production.
func (listen *ListenerCreateIndex) EnterSparseRange(ctx *parser.SparseRangeContext) {}

// EnterRange is called when entering the range production.
func (listen *ListenerCreateIndex) EnterRange(ctx *parser.RangeContext) {}

// EnterCreate is called when entering the create production.
func (listen *ListenerCreateIndex) EnterCreate(ctx *parser.CreateContext) {}

// EnterWith is called when entering the with production.
func (listen *ListenerCreateIndex) EnterWith(ctx *parser.WithContext) {}

// EnterSelect is called when entering the select production.
func (listen *ListenerCreateIndex) EnterSelect(ctx *parser.SelectContext) {}

// EnterFrom is called when entering the from production.
func (listen *ListenerCreateIndex) EnterFrom(ctx *parser.FromContext) {}

// EnterWhere is called when entering the where production.
func (listen *ListenerCreateIndex) EnterWhere(ctx *parser.WhereContext) {}

// EnterGroupby is called when entering the groupby production.
func (listen *ListenerCreateIndex) EnterGroupby(ctx *parser.GroupbyContext) {}

// EnterOrderby is called when entering the orderby production.
func (listen *ListenerCreateIndex) EnterOrderby(ctx *parser.OrderbyContext) {}

// EnterLimit is called when entering the limit production.
func (listen *ListenerCreateIndex) EnterLimit(ctx *parser.LimitContext) {}

// EnterView is called when entering the view production.
func (listen *ListenerCreateIndex) EnterView(ctx *parser.ViewContext) {}

// EnterCheck is called when entering the check production.
func (listen *ListenerCreateIndex) EnterCheck(ctx *parser.CheckContext) {}

// EnterForeignkey is called when entering the foreignkey production.
func (listen *ListenerCreateIndex) EnterForeignkey(ctx *parser.ForeignkeyContext) {}

// EnterPrimarykey is called when entering the primarykey production.
func (listen *ListenerCreateIndex) EnterPrimarykey(ctx *parser.PrimarykeyContext) {}

// EnterConstraint is called when entering the constraint production.
func (listen *ListenerCreateIndex) EnterConstraint(ctx *parser.ConstraintContext) {}

// EnterColumns is called when entering the columns production.
func (listen *ListenerCreateIndex) EnterColumns(ctx *parser.ColumnsContext) {}

// EnterCreatetable is called when entering the table production.
func (listen *ListenerCreateIndex) EnterCreatetable(ctx *parser.CreatetableContext) {}

// EnterCreateindex is called when entering the createindex production.
func (listen *ListenerCreateIndex) EnterCreateindex(ctx *parser.CreateindexContext) {
	listen.IndexName, listen.TableName, listen.ColumnName = ctx.GetIndexName().GetText(),
		ctx.GetTableName().GetText(), ctx.GetColumnName().GetText()
}

// EnterDrop is called when entering the drop production.
func (listen *ListenerCreateIndex) EnterDrop(ctx *parser.DropContext) {}

// EnterDelete is called when entering the delete production.
func (listen *ListenerCreateIndex) EnterDelete(ctx *parser.DeleteContext) {}

// EnterSet is called when entering the set production.
func (listen *ListenerCreateIndex) EnterSet(ctx *parser.SetContext) {}

// EnterUpdate is called when entering the update production.
func (listen *ListenerCreateIndex) EnterUpdate(ctx *parser.UpdateContext) {}

// EnterValues is called when entering the values production.
func (listen *ListenerCreateIndex) EnterValues(ctx *parser.ValuesContext) {}

// EnterColumnNames is called when entering the columnNames production.
func (listen *ListenerCreateIndex) EnterColumnNames(ctx *parser.ColumnNamesContext) {}

// EnterInsert is called when entering the insert production.
func (listen *ListenerCreateIndex) EnterInsert(ctx *parser.InsertContext) {}

// EnterQuery is called when entering the query production.
func (listen *ListenerCreateIndex) EnterQuery(ctx *parser.QueryContext) {}

// EnterStatement is called when entering the statement production.
func (listen *ListenerCreateIndex) EnterStatement(ctx *parser.StatementContext) {}

// EnterStart is called when entering the start production.
func (listen *ListenerCreateIndex) EnterStart(ctx *parser.StartContext) {}

// ExitName is called when exiting the name production.
func (listen *ListenerCreateIndex) ExitName(ctx *parser.NameContext) {}

// ExitConstant is called when exiting the constant production.
func (listen *ListenerCreateIndex) ExitConstant(ctx *parser.ConstantContext) {}

// ExitAggrFunc is called when exiting the aggrFunc production.
func (listen *ListenerCreateIndex) ExitAggrFunc(ctx *parser.AggrFuncContext) {}

// ExitField is called when exiting the field production.
func (listen *ListenerCreateIndex) ExitField(ctx *parser.FieldContext) {}

// ExitParameter is called when exiting the parameter production.
func (listen *ListenerCreateIndex) ExitParameter(ctx *parser.ParameterContext) {}

// ExitNameable is called when exiting the nameable production.
func (listen *ListenerCreateIndex) ExitNameable(ctx *parser.NameableContext) {}

// ExitKey is called when exiting the key production.
func (listen *ListenerCreateIndex) ExitKey(ctx *parser.KeyContext) {}

// ExitAddOrSub is called when exiting the AddOrSub production.
func (listen *ListenerCreateIndex) ExitAddOrSub(ctx *parser.AddOrSubContext) {}

// ExitValue is called when exiting the Value production.
func (listen *ListenerCreateIndex) ExitValue(ctx *parser.ValueContext) {}

// ExitMinus is called when exiting the Minus production.
func (listen *ListenerCreateIndex) ExitMinus(ctx *parser.MinusContext) {}

// ExitParentheses is called when exiting the Parentheses production.
func (listen *ListenerCreateIndex) ExitParentheses(ctx *parser.ParenthesesContext) {}

// ExitMultOrDiv is called when exiting the MultOrDiv production.
func (listen *ListenerCreateIndex) ExitMultOrDiv(ctx *parser.MultOrDivContext) {}

// ExitAsClause is called when exiting the asClause production.
func (listen *ListenerCreateIndex) ExitAsClause(ctx *parser.AsClauseContext) {}

// ExitAggregation is called when exiting the aggregation production.
func (listen *ListenerCreateIndex) ExitAggregation(ctx *parser.AggregationContext) {}

// ExitCount is called when exiting the count production.
func (listen *ListenerCreateIndex) ExitCount(ctx *parser.CountContext) {}

// ExitCalc is called when exiting the calc production.
func (listen *ListenerCreateIndex) ExitCalc(ctx *parser.CalcContext) {}

// ExitComp is called when exiting the comp production.
func (listen *ListenerCreateIndex) ExitComp(ctx *parser.CompContext) {}

// ExitCondition is called when exiting the condition production.
func (listen *ListenerCreateIndex) ExitCondition(ctx *parser.ConditionContext) {}

// ExitSortOrder is called when exiting the sortOrder production.
func (listen *ListenerCreateIndex) ExitSortOrder(ctx *parser.SortOrderContext) {}

// ExitContinuousRange is called when exiting the continuousRange production.
func (listen *ListenerCreateIndex) ExitContinuousRange(ctx *parser.ContinuousRangeContext) {}

// ExitSparseRange is called when exiting the sparseRange production.
func (listen *ListenerCreateIndex) ExitSparseRange(ctx *parser.SparseRangeContext) {}

// ExitRange is called when exiting the range production.
func (listen *ListenerCreateIndex) ExitRange(ctx *parser.RangeContext) {}

// ExitCreate is called when exiting the create production.
func (listen *ListenerCreateIndex) ExitCreate(ctx *parser.CreateContext) {}

// ExitCreate is called when exiting the create production.
func (listen *ListenerCreateIndex) ExitCreatetable(ctx *parser.CreatetableContext) {}

// ExitCreateindex is called when exiting the createindex production.
func (listen *ListenerCreateIndex) ExitCreateindex(ctx *parser.CreateindexContext) {
	fmt.Println("[CreatendexL]***Exit createindex***")
}

// ExitWith is called when exiting the with production.
func (listen *ListenerCreateIndex) ExitWith(ctx *parser.WithContext) {}

// ExitSelect is called when exiting the select production.
func (listen *ListenerCreateIndex) ExitSelect(ctx *parser.SelectContext) {}

// ExitFrom is called when exiting the from production.
func (listen *ListenerCreateIndex) ExitFrom(ctx *parser.FromContext) {}

// ExitWhere is called when exiting the where production.
func (listen *ListenerCreateIndex) ExitWhere(ctx *parser.WhereContext) {}

// ExitGroupby is called when exiting the groupby production.
func (listen *ListenerCreateIndex) ExitGroupby(ctx *parser.GroupbyContext) {}

// ExitOrderby is called when exiting the orderby production.
func (listen *ListenerCreateIndex) ExitOrderby(ctx *parser.OrderbyContext) {}

// ExitLimit is called when exiting the limit production.
func (listen *ListenerCreateIndex) ExitLimit(ctx *parser.LimitContext) {}

// ExitView is called when exiting the view production.
func (listen *ListenerCreateIndex) ExitView(ctx *parser.ViewContext) {}

// ExitCheck is called when exiting the check production.
func (listen *ListenerCreateIndex) ExitCheck(ctx *parser.CheckContext) {}

// ExitForeignkey is called when exiting the foreignkey production.
func (listen *ListenerCreateIndex) ExitForeignkey(ctx *parser.ForeignkeyContext) {}

// ExitPrimarykey is called when exiting the primarykey production.
func (listen *ListenerCreateIndex) ExitPrimarykey(ctx *parser.PrimarykeyContext) {}

// ExitConstraint is called when exiting the constraint production.
func (listen *ListenerCreateIndex) ExitConstraint(ctx *parser.ConstraintContext) {}

// ExitColumns is called when exiting the columns production.
func (listen *ListenerCreateIndex) ExitColumns(ctx *parser.ColumnsContext) {}

// ExitTable is called when exiting the table production.
func (listen *ListenerCreateIndex) ExitTable(ctx *parser.CreatetableContext) {}

// ExitDrop is called when exiting the drop production.
func (listen *ListenerCreateIndex) ExitDrop(ctx *parser.DropContext) {}

// ExitDelete is called when exiting the delete production.
func (listen *ListenerCreateIndex) ExitDelete(ctx *parser.DeleteContext) {}

// ExitSet is called when exiting the set production.
func (listen *ListenerCreateIndex) ExitSet(ctx *parser.SetContext) {}

// ExitUpdate is called when exiting the update production.
func (listen *ListenerCreateIndex) ExitUpdate(ctx *parser.UpdateContext) {}

// ExitValues is called when exiting the values production.
func (listen *ListenerCreateIndex) ExitValues(ctx *parser.ValuesContext) {}

// ExitColumnNames is called when exiting the columnNames production.
func (listen *ListenerCreateIndex) ExitColumnNames(ctx *parser.ColumnNamesContext) {}

// ExitInsert is called when exiting the insert production.
func (listen *ListenerCreateIndex) ExitInsert(ctx *parser.InsertContext) {}

// ExitQuery is called when exiting the query production.
func (listen *ListenerCreateIndex) ExitQuery(ctx *parser.QueryContext) {}

// ExitStatement is called when exiting the statement production.
func (listen *ListenerCreateIndex) ExitStatement(ctx *parser.StatementContext) {}

// ExitStart is called when exiting the start production.
func (listen *ListenerCreateIndex) ExitStart(ctx *parser.StartContext) {
	fmt.Println("[CreateIndexL]Finished processing text.")
	fmt.Println("IndexName:", listen.IndexName)
	fmt.Println("TableName:", listen.TableName)
	fmt.Println("ColumnName:", listen.ColumnName)
}
