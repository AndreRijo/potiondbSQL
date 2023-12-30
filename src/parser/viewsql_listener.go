// Code generated from ViewSQL.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // ViewSQL
import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// ViewSQLListener is a complete listener for a parse tree produced by ViewSQLParser.
type ViewSQLListener interface {
	antlr.ParseTreeListener

	// EnterName is called when entering the name production.
	EnterName(c *NameContext)

	// EnterConstant is called when entering the constant production.
	EnterConstant(c *ConstantContext)

	// EnterAggrFunc is called when entering the aggrFunc production.
	EnterAggrFunc(c *AggrFuncContext)

	// EnterField is called when entering the field production.
	EnterField(c *FieldContext)

	// EnterParameter is called when entering the parameter production.
	EnterParameter(c *ParameterContext)

	// EnterNameable is called when entering the nameable production.
	EnterNameable(c *NameableContext)

	// EnterKey is called when entering the key production.
	EnterKey(c *KeyContext)

	// EnterAddOrSub is called when entering the AddOrSub production.
	EnterAddOrSub(c *AddOrSubContext)

	// EnterValue is called when entering the Value production.
	EnterValue(c *ValueContext)

	// EnterMinus is called when entering the Minus production.
	EnterMinus(c *MinusContext)

	// EnterParentheses is called when entering the Parentheses production.
	EnterParentheses(c *ParenthesesContext)

	// EnterMultOrDiv is called when entering the MultOrDiv production.
	EnterMultOrDiv(c *MultOrDivContext)

	// EnterAsClause is called when entering the asClause production.
	EnterAsClause(c *AsClauseContext)

	// EnterAggregation is called when entering the aggregation production.
	EnterAggregation(c *AggregationContext)

	// EnterCount is called when entering the count production.
	EnterCount(c *CountContext)

	// EnterCalc is called when entering the calc production.
	EnterCalc(c *CalcContext)

	// EnterComp is called when entering the comp production.
	EnterComp(c *CompContext)

	// EnterCondition is called when entering the condition production.
	EnterCondition(c *ConditionContext)

	// EnterSortOrder is called when entering the sortOrder production.
	EnterSortOrder(c *SortOrderContext)

	// EnterContinuousRange is called when entering the continuousRange production.
	EnterContinuousRange(c *ContinuousRangeContext)

	// EnterSparseRange is called when entering the sparseRange production.
	EnterSparseRange(c *SparseRangeContext)

	// EnterRange is called when entering the range production.
	EnterRange(c *RangeContext)

	// EnterCreate is called when entering the create production.
	EnterCreate(c *CreateContext)

	// EnterWith is called when entering the with production.
	EnterWith(c *WithContext)

	// EnterSelect is called when entering the select production.
	EnterSelect(c *SelectContext)

	// EnterFrom is called when entering the from production.
	EnterFrom(c *FromContext)

	// EnterWhere is called when entering the where production.
	EnterWhere(c *WhereContext)

	// EnterGroupby is called when entering the groupby production.
	EnterGroupby(c *GroupbyContext)

	// EnterOrderby is called when entering the orderby production.
	EnterOrderby(c *OrderbyContext)

	// EnterLimit is called when entering the limit production.
	EnterLimit(c *LimitContext)

	// EnterView is called when entering the view production.
	EnterView(c *ViewContext)

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// ExitName is called when exiting the name production.
	ExitName(c *NameContext)

	// ExitConstant is called when exiting the constant production.
	ExitConstant(c *ConstantContext)

	// ExitAggrFunc is called when exiting the aggrFunc production.
	ExitAggrFunc(c *AggrFuncContext)

	// ExitField is called when exiting the field production.
	ExitField(c *FieldContext)

	// ExitParameter is called when exiting the parameter production.
	ExitParameter(c *ParameterContext)

	// ExitNameable is called when exiting the nameable production.
	ExitNameable(c *NameableContext)

	// ExitKey is called when exiting the key production.
	ExitKey(c *KeyContext)

	// ExitAddOrSub is called when exiting the AddOrSub production.
	ExitAddOrSub(c *AddOrSubContext)

	// ExitValue is called when exiting the Value production.
	ExitValue(c *ValueContext)

	// ExitMinus is called when exiting the Minus production.
	ExitMinus(c *MinusContext)

	// ExitParentheses is called when exiting the Parentheses production.
	ExitParentheses(c *ParenthesesContext)

	// ExitMultOrDiv is called when exiting the MultOrDiv production.
	ExitMultOrDiv(c *MultOrDivContext)

	// ExitAsClause is called when exiting the asClause production.
	ExitAsClause(c *AsClauseContext)

	// ExitAggregation is called when exiting the aggregation production.
	ExitAggregation(c *AggregationContext)

	// ExitCount is called when exiting the count production.
	ExitCount(c *CountContext)

	// ExitCalc is called when exiting the calc production.
	ExitCalc(c *CalcContext)

	// ExitComp is called when exiting the comp production.
	ExitComp(c *CompContext)

	// ExitCondition is called when exiting the condition production.
	ExitCondition(c *ConditionContext)

	// ExitSortOrder is called when exiting the sortOrder production.
	ExitSortOrder(c *SortOrderContext)

	// ExitContinuousRange is called when exiting the continuousRange production.
	ExitContinuousRange(c *ContinuousRangeContext)

	// ExitSparseRange is called when exiting the sparseRange production.
	ExitSparseRange(c *SparseRangeContext)

	// ExitRange is called when exiting the range production.
	ExitRange(c *RangeContext)

	// ExitCreate is called when exiting the create production.
	ExitCreate(c *CreateContext)

	// ExitWith is called when exiting the with production.
	ExitWith(c *WithContext)

	// ExitSelect is called when exiting the select production.
	ExitSelect(c *SelectContext)

	// ExitFrom is called when exiting the from production.
	ExitFrom(c *FromContext)

	// ExitWhere is called when exiting the where production.
	ExitWhere(c *WhereContext)

	// ExitGroupby is called when exiting the groupby production.
	ExitGroupby(c *GroupbyContext)

	// ExitOrderby is called when exiting the orderby production.
	ExitOrderby(c *OrderbyContext)

	// ExitLimit is called when exiting the limit production.
	ExitLimit(c *LimitContext)

	// ExitView is called when exiting the view production.
	ExitView(c *ViewContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)
}
