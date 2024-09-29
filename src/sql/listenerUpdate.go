package sql

import (
	"fmt"

	"sqlToKeyValue/src/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

//For now keep info on strings and similar, afterwards update to CRDT updates

type ListenerUpdate struct {
	Parse       *parser.ViewSQLParser
	TableName   string
	Updates     map[string]string //TODO: Maybe do CRDT updates if we know the table info?
	Conditions  []SimpleCondition
	currentMath Math
	currIndex   int
}

//TODO: This should be somewhere else
type SimpleCondition struct {
	ColumnName string //Left side
	Op         CondType
	Fun        Function
	Math       Math
}

func (c SimpleCondition) ToString() string {
	return fmt.Sprintf("{ColumnName: %s; Op: %s, Math: %s, Fun: %T}", c.ColumnName,
		c.Op.ToString(), c.Math.ToString(), c.Fun)
}

func MakeUpdateListener(parse *parser.ViewSQLParser) *ListenerUpdate {
	return &ListenerUpdate{Parse: parse, currIndex: 0}
}

// VisitTerminal is called when a terminal node is visited.
func (listen *ListenerUpdate) VisitTerminal(node antlr.TerminalNode) {
}

// VisitErrorNode is called when an error node is visited.
func (listen *ListenerUpdate) VisitErrorNode(node antlr.ErrorNode) {
}

// EnterEveryRule is called when any rule is entered.
func (listen *ListenerUpdate) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println("[UpdateL]Entered rule")
}

// ExitEveryRule is called when any rule is exited.
func (listen *ListenerUpdate) ExitEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println("[UpdateL]Exit rule")
}

// EnterName is called when entering the constant production.
func (listen *ListenerUpdate) EnterName(ctx *parser.NameContext) {}

// EnterConstant is called when entering the constant production.
func (listen *ListenerUpdate) EnterConstant(ctx *parser.ConstantContext) {}

// EnterAggrFunc is called when entering the aggrFunc production.
func (listen *ListenerUpdate) EnterAggrFunc(ctx *parser.AggrFuncContext) {}

// EnterField is called when entering the field production.
func (listen *ListenerUpdate) EnterField(ctx *parser.FieldContext) {}

// EnterParameter is called when entering the parameter production.
func (listen *ListenerUpdate) EnterParameter(ctx *parser.ParameterContext) {}

// EnterNameable is called when entering the nameable production.
func (listen *ListenerUpdate) EnterNameable(ctx *parser.NameableContext) {}

// EnterKey is called when entering the key production.
func (listen *ListenerUpdate) EnterKey(ctx *parser.KeyContext) {}

// EnterAddOrSub is called when entering the AddOrSub production.
func (listen *ListenerUpdate) EnterAddOrSub(ctx *parser.AddOrSubContext) {}

// EnterValue is called when entering the Value production.
func (listen *ListenerUpdate) EnterValue(ctx *parser.ValueContext) {}

// EnterMinus is called when entering the Minus production.
func (listen *ListenerUpdate) EnterMinus(ctx *parser.MinusContext) {}

// EnterParentheses is called when entering the Parentheses production.
func (listen *ListenerUpdate) EnterParentheses(ctx *parser.ParenthesesContext) {}

// EnterMultOrDiv is called when entering the MultOrDiv production.
func (listen *ListenerUpdate) EnterMultOrDiv(ctx *parser.MultOrDivContext) {}

// EnterAsClause is called when entering the asClause production.
func (listen *ListenerUpdate) EnterAsClause(ctx *parser.AsClauseContext) {}

// EnterAggregation is called when entering the aggregation production.
func (listen *ListenerUpdate) EnterAggregation(ctx *parser.AggregationContext) {}

// EnterCount is called when entering the count production.
func (listen *ListenerUpdate) EnterCount(ctx *parser.CountContext) {}

// EnterCalc is called when entering the calc production.
func (listen *ListenerUpdate) EnterCalc(ctx *parser.CalcContext) {}

// EnterComp is called when entering the comp production.
func (listen *ListenerUpdate) EnterComp(ctx *parser.CompContext) {}

// EnterCondition is called when entering the condition production.
func (listen *ListenerUpdate) EnterCondition(ctx *parser.ConditionContext) {
	listen.currentMath = CreateMath()
}

// EnterSortOrder is called when entering the sortOrder production.
func (listen *ListenerUpdate) EnterSortOrder(ctx *parser.SortOrderContext) {}

// EnterContinuousRange is called when entering the continuousRange production.
func (listen *ListenerUpdate) EnterContinuousRange(ctx *parser.ContinuousRangeContext) {}

// EnterSparseRange is called when entering the sparseRange production.
func (listen *ListenerUpdate) EnterSparseRange(ctx *parser.SparseRangeContext) {}

// EnterRange is called when entering the range production.
func (listen *ListenerUpdate) EnterRange(ctx *parser.RangeContext) {}

// EnterCreate is called when entering the create production.
func (listen *ListenerUpdate) EnterCreate(ctx *parser.CreateContext) {}

// EnterWith is called when entering the with production.
func (listen *ListenerUpdate) EnterWith(ctx *parser.WithContext) {}

// EnterSelect is called when entering the select production.
func (listen *ListenerUpdate) EnterSelect(ctx *parser.SelectContext) {}

// EnterFrom is called when entering the from production.
func (listen *ListenerUpdate) EnterFrom(ctx *parser.FromContext) {}

// EnterWhere is called when entering the where production.
func (listen *ListenerUpdate) EnterWhere(ctx *parser.WhereContext) {
	//TODO: Probably some special processing for primary key
	listen.Conditions = make([]SimpleCondition, len(ctx.AllCondition()))
}

// EnterGroupby is called when entering the groupby production.
func (listen *ListenerUpdate) EnterGroupby(ctx *parser.GroupbyContext) {}

// EnterOrderby is called when entering the orderby production.
func (listen *ListenerUpdate) EnterOrderby(ctx *parser.OrderbyContext) {}

// EnterLimit is called when entering the limit production.
func (listen *ListenerUpdate) EnterLimit(ctx *parser.LimitContext) {}

// EnterView is called when entering the view production.
func (listen *ListenerUpdate) EnterView(ctx *parser.ViewContext) {}

// EnterCheck is called when entering the check production.
func (listen *ListenerUpdate) EnterCheck(ctx *parser.CheckContext) {}

// EnterForeignkey is called when entering the foreignkey production.
func (listen *ListenerUpdate) EnterForeignkey(ctx *parser.ForeignkeyContext) {}

// EnterConstraint is called when entering the constraint production.
func (listen *ListenerUpdate) EnterConstraint(ctx *parser.ConstraintContext) {}

// EnterPrimarykey is called when entering the primarykey production.
func (listen *ListenerUpdate) EnterPrimarykey(ctx *parser.PrimarykeyContext) {}

// EnterColumns is called when entering the columns production.
func (listen *ListenerUpdate) EnterColumns(ctx *parser.ColumnsContext) {}

// EnterCreatetable is called when entering the table production.
func (listen *ListenerUpdate) EnterCreatetable(ctx *parser.CreatetableContext) {}

// EnterCreateindex is called when entering the createindex production.
func (listen *ListenerUpdate) EnterCreateindex(ctx *parser.CreateindexContext) {}

// EnterDrop is called when entering the drop production.
func (listen *ListenerUpdate) EnterDrop(ctx *parser.DropContext) {}

// EnterDelete is called when entering the delete production.
func (listen *ListenerUpdate) EnterDelete(ctx *parser.DeleteContext) {}

// EnterSet is called when entering the set production.
func (listen *ListenerUpdate) EnterSet(ctx *parser.SetContext) {
	names, constants := ctx.AllName(), ctx.AllConstant()
	listen.Updates = make(map[string]string, len(names))
	for i, name := range names {
		listen.Updates[name.GetText()] = constants[i].GetText()
	}
}

// EnterUpdate is called when entering the update production.
func (listen *ListenerUpdate) EnterUpdate(ctx *parser.UpdateContext) {
	listen.TableName = ctx.Name().GetText()
}

// EnterValues is called when entering the values production.
func (listen *ListenerUpdate) EnterValues(ctx *parser.ValuesContext) {}

// EnterColumnNames is called when entering the columnNames production.
func (listen *ListenerUpdate) EnterColumnNames(ctx *parser.ColumnNamesContext) {}

// EnterInsert is called when entering the insert production.
func (listen *ListenerUpdate) EnterInsert(ctx *parser.InsertContext) {}

// EnterQuery is called when entering the query production.
func (listen *ListenerUpdate) EnterQuery(ctx *parser.QueryContext) {}

// EnterStatement is called when entering the statement production.
func (listen *ListenerUpdate) EnterStatement(ctx *parser.StatementContext) {}

// EnterStart is called when entering the start production.
func (listen *ListenerUpdate) EnterStart(ctx *parser.StartContext) {}

// ExitName is called when exiting the name production.
func (listen *ListenerUpdate) ExitName(ctx *parser.NameContext) {}

// ExitConstant is called when exiting the constant production.
func (listen *ListenerUpdate) ExitConstant(ctx *parser.ConstantContext) {}

// ExitAggrFunc is called when exiting the aggrFunc production.
func (listen *ListenerUpdate) ExitAggrFunc(ctx *parser.AggrFuncContext) {}

// ExitField is called when exiting the field production.
func (listen *ListenerUpdate) ExitField(ctx *parser.FieldContext) {}

// ExitParameter is called when exiting the parameter production.
func (listen *ListenerUpdate) ExitParameter(ctx *parser.ParameterContext) {}

// ExitNameable is called when exiting the nameable production.
func (listen *ListenerUpdate) ExitNameable(ctx *parser.NameableContext) {}

// ExitKey is called when exiting the key production.
func (listen *ListenerUpdate) ExitKey(ctx *parser.KeyContext) {}

// ExitAddOrSub is called when exiting the AddOrSub production.
func (listen *ListenerUpdate) ExitAddOrSub(ctx *parser.AddOrSubContext) {
	AddOrSubHelper(ctx, listen.currentMath)
}

// ExitValue is called when exiting the Value production.
func (listen *ListenerUpdate) ExitValue(ctx *parser.ValueContext) {
	if nameable := ctx.Nameable(); nameable != nil {
		listen.currentMath.Stack.Push(MakeNameableFromContext(nameable))
	} else { //Constant
		constant := MakeConstantFromContext(ctx.Constant())
		listen.currentMath.Stack.Push(constant)
	}
}

// ExitMinus is called when exiting the Minus production.
func (listen *ListenerUpdate) ExitMinus(ctx *parser.MinusContext) {
	MinusHelper(ctx, listen.currentMath)
}

// ExitParentheses is called when exiting the Parentheses production.
func (listen *ListenerUpdate) ExitParentheses(ctx *parser.ParenthesesContext) {}

// ExitMultOrDiv is called when exiting the MultOrDiv production.
func (listen *ListenerUpdate) ExitMultOrDiv(ctx *parser.MultOrDivContext) {
	MultOrDivHelper(ctx, listen.currentMath)
}

// ExitAsClause is called when exiting the asClause production.
func (listen *ListenerUpdate) ExitAsClause(ctx *parser.AsClauseContext) {}

// ExitAggregation is called when exiting the aggregation production.
func (listen *ListenerUpdate) ExitAggregation(ctx *parser.AggregationContext) {}

// ExitCount is called when exiting the count production.
func (listen *ListenerUpdate) ExitCount(ctx *parser.CountContext) {}

// ExitCalc is called when exiting the calc production.
func (listen *ListenerUpdate) ExitCalc(ctx *parser.CalcContext) {}

// ExitComp is called when exiting the comp production.
func (listen *ListenerUpdate) ExitComp(ctx *parser.CompContext) {}

// ExitCondition is called when exiting the condition production.
func (listen *ListenerUpdate) ExitCondition(ctx *parser.ConditionContext) {
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
func (listen *ListenerUpdate) ExitSortOrder(ctx *parser.SortOrderContext) {}

// ExitContinuousRange is called when exiting the continuousRange production.
func (listen *ListenerUpdate) ExitContinuousRange(ctx *parser.ContinuousRangeContext) {}

// ExitSparseRange is called when exiting the sparseRange production.
func (listen *ListenerUpdate) ExitSparseRange(ctx *parser.SparseRangeContext) {}

// ExitRange is called when exiting the range production.
func (listen *ListenerUpdate) ExitRange(ctx *parser.RangeContext) {}

// ExitCreate is called when exiting the create production.
func (listen *ListenerUpdate) ExitCreate(ctx *parser.CreateContext) {}

// ExitCreate is called when exiting the create production.
func (listen *ListenerUpdate) ExitCreatetable(ctx *parser.CreatetableContext) {}

// ExitCreateindex is called when exiting the createindex production.
func (listen *ListenerUpdate) ExitCreateindex(ctx *parser.CreateindexContext) {}

// ExitWith is called when exiting the with production.
func (listen *ListenerUpdate) ExitWith(ctx *parser.WithContext) {}

// ExitSelect is called when exiting the select production.
func (listen *ListenerUpdate) ExitSelect(ctx *parser.SelectContext) {}

// ExitFrom is called when exiting the from production.
func (listen *ListenerUpdate) ExitFrom(ctx *parser.FromContext) {}

// ExitWhere is called when exiting the where production.
func (listen *ListenerUpdate) ExitWhere(ctx *parser.WhereContext) {}

// ExitGroupby is called when exiting the groupby production.
func (listen *ListenerUpdate) ExitGroupby(ctx *parser.GroupbyContext) {}

// ExitOrderby is called when exiting the orderby production.
func (listen *ListenerUpdate) ExitOrderby(ctx *parser.OrderbyContext) {}

// ExitLimit is called when exiting the limit production.
func (listen *ListenerUpdate) ExitLimit(ctx *parser.LimitContext) {}

// ExitView is called when exiting the view production.
func (listen *ListenerUpdate) ExitView(ctx *parser.ViewContext) {}

// ExitCheck is called when exiting the check production.
func (listen *ListenerUpdate) ExitCheck(ctx *parser.CheckContext) {}

// ExitForeignkey is called when exiting the foreignkey production.
func (listen *ListenerUpdate) ExitForeignkey(ctx *parser.ForeignkeyContext) {}

// ExitPrimarykey is called when exiting the primarykey production.
func (listen *ListenerUpdate) ExitPrimarykey(ctx *parser.PrimarykeyContext) {}

// ExitConstraint is called when exiting the constraint production.
func (listen *ListenerUpdate) ExitConstraint(ctx *parser.ConstraintContext) {}

// ExitColumns is called when exiting the columns production.
func (listen *ListenerUpdate) ExitColumns(ctx *parser.ColumnsContext) {}

// ExitTable is called when exiting the table production.
func (listen *ListenerUpdate) ExitTable(ctx *parser.CreatetableContext) {}

// ExitDrop is called when exiting the drop production.
func (listen *ListenerUpdate) ExitDrop(ctx *parser.DropContext) {}

// ExitDelete is called when exiting the delete production.
func (listen *ListenerUpdate) ExitDelete(ctx *parser.DeleteContext) {}

// ExitSet is called when exiting the set production.
func (listen *ListenerUpdate) ExitSet(ctx *parser.SetContext) {}

// ExitUpdate is called when exiting the update production.
func (listen *ListenerUpdate) ExitUpdate(ctx *parser.UpdateContext) {}

// ExitValues is called when exiting the values production.
func (listen *ListenerUpdate) ExitValues(ctx *parser.ValuesContext) {}

// ExitColumnNames is called when exiting the columnNames production.
func (listen *ListenerUpdate) ExitColumnNames(ctx *parser.ColumnNamesContext) {}

// ExitInsert is called when exiting the insert production.
func (listen *ListenerUpdate) ExitInsert(ctx *parser.InsertContext) {}

// ExitQuery is called when exiting the query production.
func (listen *ListenerUpdate) ExitQuery(ctx *parser.QueryContext) {}

// ExitStatement is called when exiting the statement production.
func (listen *ListenerUpdate) ExitStatement(ctx *parser.StatementContext) {}

// ExitStart is called when exiting the start production.
func (listen *ListenerUpdate) ExitStart(ctx *parser.StartContext) {
	fmt.Println("[UpdateL]Finished processing text.")
	fmt.Println("TableName:", listen.TableName)
	fmt.Println("Updates:", listen.Updates)
	fmt.Print("Conditions: [")
	for _, cond := range listen.Conditions {
		fmt.Print(cond.ToString(), ",")
	}
	fmt.Println("]")
}
