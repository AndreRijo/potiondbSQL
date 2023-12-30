// Code generated from ViewSQL.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // ViewSQL
import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// BaseViewSQLListener is a complete listener for a parse tree produced by ViewSQLParser.
type BaseViewSQLListener struct{}

var _ ViewSQLListener = &BaseViewSQLListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseViewSQLListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseViewSQLListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseViewSQLListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseViewSQLListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterName is called when production name is entered.
func (s *BaseViewSQLListener) EnterName(ctx *NameContext) {}

// ExitName is called when production name is exited.
func (s *BaseViewSQLListener) ExitName(ctx *NameContext) {}

// EnterConstant is called when production constant is entered.
func (s *BaseViewSQLListener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *BaseViewSQLListener) ExitConstant(ctx *ConstantContext) {}

// EnterAggrFunc is called when production aggrFunc is entered.
func (s *BaseViewSQLListener) EnterAggrFunc(ctx *AggrFuncContext) {}

// ExitAggrFunc is called when production aggrFunc is exited.
func (s *BaseViewSQLListener) ExitAggrFunc(ctx *AggrFuncContext) {}

// EnterField is called when production field is entered.
func (s *BaseViewSQLListener) EnterField(ctx *FieldContext) {}

// ExitField is called when production field is exited.
func (s *BaseViewSQLListener) ExitField(ctx *FieldContext) {}

// EnterParameter is called when production parameter is entered.
func (s *BaseViewSQLListener) EnterParameter(ctx *ParameterContext) {}

// ExitParameter is called when production parameter is exited.
func (s *BaseViewSQLListener) ExitParameter(ctx *ParameterContext) {}

// EnterNameable is called when production nameable is entered.
func (s *BaseViewSQLListener) EnterNameable(ctx *NameableContext) {}

// ExitNameable is called when production nameable is exited.
func (s *BaseViewSQLListener) ExitNameable(ctx *NameableContext) {}

// EnterKey is called when production key is entered.
func (s *BaseViewSQLListener) EnterKey(ctx *KeyContext) {}

// ExitKey is called when production key is exited.
func (s *BaseViewSQLListener) ExitKey(ctx *KeyContext) {}

// EnterAddOrSub is called when production AddOrSub is entered.
func (s *BaseViewSQLListener) EnterAddOrSub(ctx *AddOrSubContext) {}

// ExitAddOrSub is called when production AddOrSub is exited.
func (s *BaseViewSQLListener) ExitAddOrSub(ctx *AddOrSubContext) {}

// EnterValue is called when production Value is entered.
func (s *BaseViewSQLListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production Value is exited.
func (s *BaseViewSQLListener) ExitValue(ctx *ValueContext) {}

// EnterMinus is called when production Minus is entered.
func (s *BaseViewSQLListener) EnterMinus(ctx *MinusContext) {}

// ExitMinus is called when production Minus is exited.
func (s *BaseViewSQLListener) ExitMinus(ctx *MinusContext) {}

// EnterParentheses is called when production Parentheses is entered.
func (s *BaseViewSQLListener) EnterParentheses(ctx *ParenthesesContext) {}

// ExitParentheses is called when production Parentheses is exited.
func (s *BaseViewSQLListener) ExitParentheses(ctx *ParenthesesContext) {}

// EnterMultOrDiv is called when production MultOrDiv is entered.
func (s *BaseViewSQLListener) EnterMultOrDiv(ctx *MultOrDivContext) {}

// ExitMultOrDiv is called when production MultOrDiv is exited.
func (s *BaseViewSQLListener) ExitMultOrDiv(ctx *MultOrDivContext) {}

// EnterAsClause is called when production asClause is entered.
func (s *BaseViewSQLListener) EnterAsClause(ctx *AsClauseContext) {}

// ExitAsClause is called when production asClause is exited.
func (s *BaseViewSQLListener) ExitAsClause(ctx *AsClauseContext) {}

// EnterAggregation is called when production aggregation is entered.
func (s *BaseViewSQLListener) EnterAggregation(ctx *AggregationContext) {}

// ExitAggregation is called when production aggregation is exited.
func (s *BaseViewSQLListener) ExitAggregation(ctx *AggregationContext) {}

// EnterCount is called when production count is entered.
func (s *BaseViewSQLListener) EnterCount(ctx *CountContext) {}

// ExitCount is called when production count is exited.
func (s *BaseViewSQLListener) ExitCount(ctx *CountContext) {}

// EnterCalc is called when production calc is entered.
func (s *BaseViewSQLListener) EnterCalc(ctx *CalcContext) {}

// ExitCalc is called when production calc is exited.
func (s *BaseViewSQLListener) ExitCalc(ctx *CalcContext) {}

// EnterComp is called when production comp is entered.
func (s *BaseViewSQLListener) EnterComp(ctx *CompContext) {}

// ExitComp is called when production comp is exited.
func (s *BaseViewSQLListener) ExitComp(ctx *CompContext) {}

// EnterCondition is called when production condition is entered.
func (s *BaseViewSQLListener) EnterCondition(ctx *ConditionContext) {}

// ExitCondition is called when production condition is exited.
func (s *BaseViewSQLListener) ExitCondition(ctx *ConditionContext) {}

// EnterSortOrder is called when production sortOrder is entered.
func (s *BaseViewSQLListener) EnterSortOrder(ctx *SortOrderContext) {}

// ExitSortOrder is called when production sortOrder is exited.
func (s *BaseViewSQLListener) ExitSortOrder(ctx *SortOrderContext) {}

// EnterContinuousRange is called when production continuousRange is entered.
func (s *BaseViewSQLListener) EnterContinuousRange(ctx *ContinuousRangeContext) {}

// ExitContinuousRange is called when production continuousRange is exited.
func (s *BaseViewSQLListener) ExitContinuousRange(ctx *ContinuousRangeContext) {}

// EnterSparseRange is called when production sparseRange is entered.
func (s *BaseViewSQLListener) EnterSparseRange(ctx *SparseRangeContext) {}

// ExitSparseRange is called when production sparseRange is exited.
func (s *BaseViewSQLListener) ExitSparseRange(ctx *SparseRangeContext) {}

// EnterRange is called when production range is entered.
func (s *BaseViewSQLListener) EnterRange(ctx *RangeContext) {}

// ExitRange is called when production range is exited.
func (s *BaseViewSQLListener) ExitRange(ctx *RangeContext) {}

// EnterCreate is called when production create is entered.
func (s *BaseViewSQLListener) EnterCreate(ctx *CreateContext) {}

// ExitCreate is called when production create is exited.
func (s *BaseViewSQLListener) ExitCreate(ctx *CreateContext) {}

// EnterWith is called when production with is entered.
func (s *BaseViewSQLListener) EnterWith(ctx *WithContext) {}

// ExitWith is called when production with is exited.
func (s *BaseViewSQLListener) ExitWith(ctx *WithContext) {}

// EnterSelect is called when production select is entered.
func (s *BaseViewSQLListener) EnterSelect(ctx *SelectContext) {}

// ExitSelect is called when production select is exited.
func (s *BaseViewSQLListener) ExitSelect(ctx *SelectContext) {}

// EnterFrom is called when production from is entered.
func (s *BaseViewSQLListener) EnterFrom(ctx *FromContext) {}

// ExitFrom is called when production from is exited.
func (s *BaseViewSQLListener) ExitFrom(ctx *FromContext) {}

// EnterWhere is called when production where is entered.
func (s *BaseViewSQLListener) EnterWhere(ctx *WhereContext) {}

// ExitWhere is called when production where is exited.
func (s *BaseViewSQLListener) ExitWhere(ctx *WhereContext) {}

// EnterGroupby is called when production groupby is entered.
func (s *BaseViewSQLListener) EnterGroupby(ctx *GroupbyContext) {}

// ExitGroupby is called when production groupby is exited.
func (s *BaseViewSQLListener) ExitGroupby(ctx *GroupbyContext) {}

// EnterOrderby is called when production orderby is entered.
func (s *BaseViewSQLListener) EnterOrderby(ctx *OrderbyContext) {}

// ExitOrderby is called when production orderby is exited.
func (s *BaseViewSQLListener) ExitOrderby(ctx *OrderbyContext) {}

// EnterLimit is called when production limit is entered.
func (s *BaseViewSQLListener) EnterLimit(ctx *LimitContext) {}

// ExitLimit is called when production limit is exited.
func (s *BaseViewSQLListener) ExitLimit(ctx *LimitContext) {}

// EnterView is called when production view is entered.
func (s *BaseViewSQLListener) EnterView(ctx *ViewContext) {}

// ExitView is called when production view is exited.
func (s *BaseViewSQLListener) ExitView(ctx *ViewContext) {}

// EnterStart is called when production start is entered.
func (s *BaseViewSQLListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseViewSQLListener) ExitStart(ctx *StartContext) {}
