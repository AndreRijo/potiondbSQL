// Generated from /Users/a.rijo/Documents/University_6th_year/goProjects/sqlToKeyValue/src/parser/ViewSQL.g4 by ANTLR 4.7.1
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class ViewSQLParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.7.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		ADD=1, SUB=2, MULT=3, DIV=4, SUM=5, AVG=6, MAX=7, MIN=8, COUNT=9, AND=10, 
		ASC=11, DESC=12, EQUAL=13, HIGHER=14, HIGHER_EQUAL=15, LOWER=16, LOWER_EQUAL=17, 
		NOT_EQUAL=18, DOT=19, SEPARATOR=20, RANGE_SEP=21, LEFT_P=22, RIGHT_P=23, 
		LEFT_SP=24, RIGHT_SP=25, INV_COMMA=26, LEFT_CURLY=27, RIGHT_CURLY=28, 
		AS=29, CREATE=30, VIEW=31, SELECT=32, FROM=33, WHERE=34, GROUP=35, BY=36, 
		ORDER=37, LIMIT=38, WITH=39, IN=40, PRIMARY=41, KEY=42, TABLE=43, INDEX=44, 
		UPDATE=45, DELETE=46, INSERT=47, INTO=48, DROP=49, SET=50, FOREIGN=51, 
		REFERENCES=52, VALUES=53, CHECK=54, ON=55, AW=56, RW=57, LWW=58, DEFAULT=59, 
		INTEGER=60, COUNTER=61, BOOLEAN=62, VARCHAR=63, DATE_TYPE=64, DATE=65, 
		BOOL=66, STRING=67, INT=68, FLOAT=69, WHITESPACE=70;
	public static final int
		RULE_name = 0, RULE_constant = 1, RULE_aggrFunc = 2, RULE_field = 3, RULE_parameter = 4, 
		RULE_nameable = 5, RULE_key = 6, RULE_math = 7, RULE_asClause = 8, RULE_aggregation = 9, 
		RULE_count = 10, RULE_calc = 11, RULE_comp = 12, RULE_condition = 13, 
		RULE_sortOrder = 14, RULE_continuousRange = 15, RULE_sparseRange = 16, 
		RULE_range = 17, RULE_create = 18, RULE_with = 19, RULE_select = 20, RULE_from = 21, 
		RULE_where = 22, RULE_groupby = 23, RULE_orderby = 24, RULE_limit = 25, 
		RULE_view = 26, RULE_check = 27, RULE_foreignkey = 28, RULE_primarykey = 29, 
		RULE_constraint = 30, RULE_columns = 31, RULE_createtable = 32, RULE_createindex = 33, 
		RULE_drop = 34, RULE_delete = 35, RULE_set = 36, RULE_update = 37, RULE_values = 38, 
		RULE_columnNames = 39, RULE_insert = 40, RULE_query = 41, RULE_statement = 42, 
		RULE_start = 43;
	public static final String[] ruleNames = {
		"name", "constant", "aggrFunc", "field", "parameter", "nameable", "key", 
		"math", "asClause", "aggregation", "count", "calc", "comp", "condition", 
		"sortOrder", "continuousRange", "sparseRange", "range", "create", "with", 
		"select", "from", "where", "groupby", "orderby", "limit", "view", "check", 
		"foreignkey", "primarykey", "constraint", "columns", "createtable", "createindex", 
		"drop", "delete", "set", "update", "values", "columnNames", "insert", 
		"query", "statement", "start"
	};

	private static final String[] _LITERAL_NAMES = {
		null, "'+'", "'-'", "'*'", "'/'", "'SUM'", "'AVG'", "'MAX'", "'MIN'", 
		"'COUNT'", "'AND'", "'ASC'", "'DESC'", "'='", "'>'", "'>='", "'<'", "'<='", 
		"'!='", "'.'", "','", "':'", "'('", "')'", "'['", "']'", "'\"'", "'{'", 
		"'}'", "'AS'", "'CREATE'", "'VIEW'", "'SELECT'", "'FROM'", "'WHERE'", 
		"'GROUP'", "'BY'", "'ORDER'", "'LIMIT'", "'WITH'", "'IN'", "'PRIMARY'", 
		"'KEY'", "'TABLE'", "'INDEX'", "'UPDATE'", "'DELETE'", "'INSERT'", "'INTO'", 
		"'DROP'", "'SET'", "'FOREIGN'", "'REFERENCES'", "'VALUES'", "'CHECK'", 
		"'ON'", "'AW'", "'RW'", "'LWW'", "'DEFAULT'", "'integer'", "'counter'", 
		"'boolean'", "'varchar'", "'date'"
	};
	private static final String[] _SYMBOLIC_NAMES = {
		null, "ADD", "SUB", "MULT", "DIV", "SUM", "AVG", "MAX", "MIN", "COUNT", 
		"AND", "ASC", "DESC", "EQUAL", "HIGHER", "HIGHER_EQUAL", "LOWER", "LOWER_EQUAL", 
		"NOT_EQUAL", "DOT", "SEPARATOR", "RANGE_SEP", "LEFT_P", "RIGHT_P", "LEFT_SP", 
		"RIGHT_SP", "INV_COMMA", "LEFT_CURLY", "RIGHT_CURLY", "AS", "CREATE", 
		"VIEW", "SELECT", "FROM", "WHERE", "GROUP", "BY", "ORDER", "LIMIT", "WITH", 
		"IN", "PRIMARY", "KEY", "TABLE", "INDEX", "UPDATE", "DELETE", "INSERT", 
		"INTO", "DROP", "SET", "FOREIGN", "REFERENCES", "VALUES", "CHECK", "ON", 
		"AW", "RW", "LWW", "DEFAULT", "INTEGER", "COUNTER", "BOOLEAN", "VARCHAR", 
		"DATE_TYPE", "DATE", "BOOL", "STRING", "INT", "FLOAT", "WHITESPACE"
	};
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "ViewSQL.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public ViewSQLParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}
	public static class NameContext extends ParserRuleContext {
		public TerminalNode STRING() { return getToken(ViewSQLParser.STRING, 0); }
		public NameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_name; }
	}

	public final NameContext name() throws RecognitionException {
		NameContext _localctx = new NameContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_name);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(88);
			match(STRING);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ConstantContext extends ParserRuleContext {
		public TerminalNode DATE() { return getToken(ViewSQLParser.DATE, 0); }
		public TerminalNode BOOL() { return getToken(ViewSQLParser.BOOL, 0); }
		public TerminalNode INT() { return getToken(ViewSQLParser.INT, 0); }
		public TerminalNode FLOAT() { return getToken(ViewSQLParser.FLOAT, 0); }
		public List<TerminalNode> STRING() { return getTokens(ViewSQLParser.STRING); }
		public TerminalNode STRING(int i) {
			return getToken(ViewSQLParser.STRING, i);
		}
		public List<TerminalNode> WHITESPACE() { return getTokens(ViewSQLParser.WHITESPACE); }
		public TerminalNode WHITESPACE(int i) {
			return getToken(ViewSQLParser.WHITESPACE, i);
		}
		public ConstantContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_constant; }
	}

	public final ConstantContext constant() throws RecognitionException {
		ConstantContext _localctx = new ConstantContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_constant);
		int _la;
		try {
			setState(103);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case DATE:
				enterOuterAlt(_localctx, 1);
				{
				setState(90);
				match(DATE);
				}
				break;
			case BOOL:
				enterOuterAlt(_localctx, 2);
				{
				setState(91);
				match(BOOL);
				}
				break;
			case INT:
				enterOuterAlt(_localctx, 3);
				{
				setState(92);
				match(INT);
				}
				break;
			case FLOAT:
				enterOuterAlt(_localctx, 4);
				{
				setState(93);
				match(FLOAT);
				}
				break;
			case INV_COMMA:
				enterOuterAlt(_localctx, 5);
				{
				setState(94);
				match(INV_COMMA);
				setState(95);
				match(STRING);
				setState(99);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==STRING || _la==WHITESPACE) {
					{
					{
					setState(96);
					_la = _input.LA(1);
					if ( !(_la==STRING || _la==WHITESPACE) ) {
					_errHandler.recoverInline(this);
					}
					else {
						if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
						_errHandler.reportMatch(this);
						consume();
					}
					}
					}
					setState(101);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(102);
				match(INV_COMMA);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class AggrFuncContext extends ParserRuleContext {
		public AggrFuncContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_aggrFunc; }
	}

	public final AggrFuncContext aggrFunc() throws RecognitionException {
		AggrFuncContext _localctx = new AggrFuncContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_aggrFunc);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(105);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << SUM) | (1L << AVG) | (1L << MAX) | (1L << MIN))) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class FieldContext extends ParserRuleContext {
		public List<NameContext> name() {
			return getRuleContexts(NameContext.class);
		}
		public NameContext name(int i) {
			return getRuleContext(NameContext.class,i);
		}
		public FieldContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_field; }
	}

	public final FieldContext field() throws RecognitionException {
		FieldContext _localctx = new FieldContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_field);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(107);
			name();
			setState(108);
			match(DOT);
			setState(109);
			name();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ParameterContext extends ParserRuleContext {
		public NameContext name() {
			return getRuleContext(NameContext.class,0);
		}
		public ParameterContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parameter; }
	}

	public final ParameterContext parameter() throws RecognitionException {
		ParameterContext _localctx = new ParameterContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_parameter);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(111);
			match(LEFT_SP);
			setState(112);
			name();
			setState(113);
			match(RIGHT_SP);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class NameableContext extends ParserRuleContext {
		public NameContext name() {
			return getRuleContext(NameContext.class,0);
		}
		public FieldContext field() {
			return getRuleContext(FieldContext.class,0);
		}
		public ParameterContext parameter() {
			return getRuleContext(ParameterContext.class,0);
		}
		public NameableContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_nameable; }
	}

	public final NameableContext nameable() throws RecognitionException {
		NameableContext _localctx = new NameableContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_nameable);
		try {
			setState(118);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(115);
				name();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(116);
				field();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(117);
				parameter();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class KeyContext extends ParserRuleContext {
		public NameableContext nameable() {
			return getRuleContext(NameableContext.class,0);
		}
		public KeyContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_key; }
	}

	public final KeyContext key() throws RecognitionException {
		KeyContext _localctx = new KeyContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_key);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(120);
			match(PRIMARY);
			setState(121);
			match(KEY);
			setState(122);
			match(LEFT_P);
			setState(123);
			nameable();
			setState(124);
			match(RIGHT_P);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class MathContext extends ParserRuleContext {
		public MathContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_math; }
	 
		public MathContext() { }
		public void copyFrom(MathContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class AddOrSubContext extends MathContext {
		public Token opType;
		public List<MathContext> math() {
			return getRuleContexts(MathContext.class);
		}
		public MathContext math(int i) {
			return getRuleContext(MathContext.class,i);
		}
		public AddOrSubContext(MathContext ctx) { copyFrom(ctx); }
	}
	public static class ValueContext extends MathContext {
		public NameableContext nameable() {
			return getRuleContext(NameableContext.class,0);
		}
		public ConstantContext constant() {
			return getRuleContext(ConstantContext.class,0);
		}
		public ValueContext(MathContext ctx) { copyFrom(ctx); }
	}
	public static class MinusContext extends MathContext {
		public TerminalNode SUB() { return getToken(ViewSQLParser.SUB, 0); }
		public MathContext math() {
			return getRuleContext(MathContext.class,0);
		}
		public MinusContext(MathContext ctx) { copyFrom(ctx); }
	}
	public static class ParenthesesContext extends MathContext {
		public MathContext math() {
			return getRuleContext(MathContext.class,0);
		}
		public ParenthesesContext(MathContext ctx) { copyFrom(ctx); }
	}
	public static class MultOrDivContext extends MathContext {
		public Token opType;
		public List<MathContext> math() {
			return getRuleContexts(MathContext.class);
		}
		public MathContext math(int i) {
			return getRuleContext(MathContext.class,i);
		}
		public MultOrDivContext(MathContext ctx) { copyFrom(ctx); }
	}

	public final MathContext math() throws RecognitionException {
		return math(0);
	}

	private MathContext math(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		MathContext _localctx = new MathContext(_ctx, _parentState);
		MathContext _prevctx = _localctx;
		int _startState = 14;
		enterRecursionRule(_localctx, 14, RULE_math, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(137);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case LEFT_SP:
			case INV_COMMA:
			case DATE:
			case BOOL:
			case STRING:
			case INT:
			case FLOAT:
				{
				_localctx = new ValueContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(129);
				_errHandler.sync(this);
				switch (_input.LA(1)) {
				case LEFT_SP:
				case STRING:
					{
					setState(127);
					nameable();
					}
					break;
				case INV_COMMA:
				case DATE:
				case BOOL:
				case INT:
				case FLOAT:
					{
					setState(128);
					constant();
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				}
				break;
			case SUB:
				{
				_localctx = new MinusContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(131);
				match(SUB);
				setState(132);
				math(4);
				}
				break;
			case LEFT_P:
				{
				_localctx = new ParenthesesContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(133);
				match(LEFT_P);
				setState(134);
				math(0);
				setState(135);
				match(RIGHT_P);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(147);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,6,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(145);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,5,_ctx) ) {
					case 1:
						{
						_localctx = new MultOrDivContext(new MathContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_math);
						setState(139);
						if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
						setState(140);
						((MultOrDivContext)_localctx).opType = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==MULT || _la==DIV) ) {
							((MultOrDivContext)_localctx).opType = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(141);
						math(3);
						}
						break;
					case 2:
						{
						_localctx = new AddOrSubContext(new MathContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_math);
						setState(142);
						if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
						setState(143);
						((AddOrSubContext)_localctx).opType = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==ADD || _la==SUB) ) {
							((AddOrSubContext)_localctx).opType = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(144);
						math(2);
						}
						break;
					}
					} 
				}
				setState(149);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,6,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class AsClauseContext extends ParserRuleContext {
		public MathContext math() {
			return getRuleContext(MathContext.class,0);
		}
		public NameContext name() {
			return getRuleContext(NameContext.class,0);
		}
		public AsClauseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_asClause; }
	}

	public final AsClauseContext asClause() throws RecognitionException {
		AsClauseContext _localctx = new AsClauseContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_asClause);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(150);
			math(0);
			setState(151);
			match(AS);
			setState(152);
			name();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class AggregationContext extends ParserRuleContext {
		public AggrFuncContext aggrFunc() {
			return getRuleContext(AggrFuncContext.class,0);
		}
		public MathContext math() {
			return getRuleContext(MathContext.class,0);
		}
		public NameContext name() {
			return getRuleContext(NameContext.class,0);
		}
		public AggregationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_aggregation; }
	}

	public final AggregationContext aggregation() throws RecognitionException {
		AggregationContext _localctx = new AggregationContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_aggregation);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(154);
			aggrFunc();
			setState(155);
			match(LEFT_P);
			setState(156);
			math(0);
			setState(157);
			match(RIGHT_P);
			setState(158);
			match(AS);
			setState(159);
			name();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class CountContext extends ParserRuleContext {
		public NameContext name() {
			return getRuleContext(NameContext.class,0);
		}
		public NameableContext nameable() {
			return getRuleContext(NameableContext.class,0);
		}
		public CountContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_count; }
	}

	public final CountContext count() throws RecognitionException {
		CountContext _localctx = new CountContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_count);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(161);
			match(COUNT);
			setState(162);
			match(LEFT_P);
			setState(165);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case LEFT_SP:
			case STRING:
				{
				setState(163);
				nameable();
				}
				break;
			case MULT:
				{
				setState(164);
				match(MULT);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			setState(167);
			match(RIGHT_P);
			setState(168);
			match(AS);
			setState(169);
			name();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class CalcContext extends ParserRuleContext {
		public Token all;
		public KeyContext key() {
			return getRuleContext(KeyContext.class,0);
		}
		public NameableContext nameable() {
			return getRuleContext(NameableContext.class,0);
		}
		public AsClauseContext asClause() {
			return getRuleContext(AsClauseContext.class,0);
		}
		public AggregationContext aggregation() {
			return getRuleContext(AggregationContext.class,0);
		}
		public CountContext count() {
			return getRuleContext(CountContext.class,0);
		}
		public CalcContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_calc; }
	}

	public final CalcContext calc() throws RecognitionException {
		CalcContext _localctx = new CalcContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_calc);
		try {
			setState(177);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,8,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(171);
				key();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(172);
				nameable();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(173);
				asClause();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(174);
				aggregation();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(175);
				count();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(176);
				((CalcContext)_localctx).all = match(MULT);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class CompContext extends ParserRuleContext {
		public Token opType;
		public NameContext name() {
			return getRuleContext(NameContext.class,0);
		}
		public CompContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_comp; }
	}

	public final CompContext comp() throws RecognitionException {
		CompContext _localctx = new CompContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_comp);
		int _la;
		try {
			setState(181);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case EQUAL:
			case HIGHER:
			case HIGHER_EQUAL:
			case LOWER:
			case LOWER_EQUAL:
			case NOT_EQUAL:
				enterOuterAlt(_localctx, 1);
				{
				setState(179);
				((CompContext)_localctx).opType = _input.LT(1);
				_la = _input.LA(1);
				if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << EQUAL) | (1L << HIGHER) | (1L << HIGHER_EQUAL) | (1L << LOWER) | (1L << LOWER_EQUAL) | (1L << NOT_EQUAL))) != 0)) ) {
					((CompContext)_localctx).opType = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				}
				break;
			case STRING:
				enterOuterAlt(_localctx, 2);
				{
				setState(180);
				name();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ConditionContext extends ParserRuleContext {
		public NameableContext nameable() {
			return getRuleContext(NameableContext.class,0);
		}
		public CompContext comp() {
			return getRuleContext(CompContext.class,0);
		}
		public MathContext math() {
			return getRuleContext(MathContext.class,0);
		}
		public ConditionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_condition; }
	}

	public final ConditionContext condition() throws RecognitionException {
		ConditionContext _localctx = new ConditionContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_condition);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(183);
			nameable();
			setState(184);
			comp();
			setState(185);
			math(0);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class SortOrderContext extends ParserRuleContext {
		public SortOrderContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_sortOrder; }
	}

	public final SortOrderContext sortOrder() throws RecognitionException {
		SortOrderContext _localctx = new SortOrderContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_sortOrder);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(187);
			_la = _input.LA(1);
			if ( !(_la==ASC || _la==DESC) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ContinuousRangeContext extends ParserRuleContext {
		public Token left;
		public Token right;
		public List<ConstantContext> constant() {
			return getRuleContexts(ConstantContext.class);
		}
		public ConstantContext constant(int i) {
			return getRuleContext(ConstantContext.class,i);
		}
		public ContinuousRangeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_continuousRange; }
	}

	public final ContinuousRangeContext continuousRange() throws RecognitionException {
		ContinuousRangeContext _localctx = new ContinuousRangeContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_continuousRange);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(189);
			((ContinuousRangeContext)_localctx).left = _input.LT(1);
			_la = _input.LA(1);
			if ( !(_la==LEFT_SP || _la==RIGHT_SP) ) {
				((ContinuousRangeContext)_localctx).left = (Token)_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			{
			setState(190);
			constant();
			setState(191);
			match(RANGE_SEP);
			setState(192);
			constant();
			}
			setState(194);
			((ContinuousRangeContext)_localctx).right = _input.LT(1);
			_la = _input.LA(1);
			if ( !(_la==LEFT_SP || _la==RIGHT_SP) ) {
				((ContinuousRangeContext)_localctx).right = (Token)_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class SparseRangeContext extends ParserRuleContext {
		public List<ConstantContext> constant() {
			return getRuleContexts(ConstantContext.class);
		}
		public ConstantContext constant(int i) {
			return getRuleContext(ConstantContext.class,i);
		}
		public SparseRangeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_sparseRange; }
	}

	public final SparseRangeContext sparseRange() throws RecognitionException {
		SparseRangeContext _localctx = new SparseRangeContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_sparseRange);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(196);
			match(LEFT_CURLY);
			setState(197);
			constant();
			setState(202);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==SEPARATOR) {
				{
				{
				setState(198);
				match(SEPARATOR);
				setState(199);
				constant();
				}
				}
				setState(204);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(205);
			match(RIGHT_CURLY);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class RangeContext extends ParserRuleContext {
		public ContinuousRangeContext continuousRange() {
			return getRuleContext(ContinuousRangeContext.class,0);
		}
		public SparseRangeContext sparseRange() {
			return getRuleContext(SparseRangeContext.class,0);
		}
		public RangeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_range; }
	}

	public final RangeContext range() throws RecognitionException {
		RangeContext _localctx = new RangeContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_range);
		try {
			setState(209);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case LEFT_SP:
			case RIGHT_SP:
				enterOuterAlt(_localctx, 1);
				{
				setState(207);
				continuousRange();
				}
				break;
			case LEFT_CURLY:
				enterOuterAlt(_localctx, 2);
				{
				setState(208);
				sparseRange();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class CreateContext extends ParserRuleContext {
		public List<TerminalNode> STRING() { return getTokens(ViewSQLParser.STRING); }
		public TerminalNode STRING(int i) {
			return getToken(ViewSQLParser.STRING, i);
		}
		public CreateContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_create; }
	}

	public final CreateContext create() throws RecognitionException {
		CreateContext _localctx = new CreateContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_create);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(211);
			match(CREATE);
			setState(212);
			match(VIEW);
			setState(213);
			match(LEFT_P);
			setState(214);
			match(STRING);
			setState(215);
			match(SEPARATOR);
			setState(216);
			match(STRING);
			setState(217);
			match(RIGHT_P);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class WithContext extends ParserRuleContext {
		public List<TerminalNode> STRING() { return getTokens(ViewSQLParser.STRING); }
		public TerminalNode STRING(int i) {
			return getToken(ViewSQLParser.STRING, i);
		}
		public List<RangeContext> range() {
			return getRuleContexts(RangeContext.class);
		}
		public RangeContext range(int i) {
			return getRuleContext(RangeContext.class,i);
		}
		public WithContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_with; }
	}

	public final WithContext with() throws RecognitionException {
		WithContext _localctx = new WithContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_with);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(219);
			match(WITH);
			setState(220);
			match(STRING);
			setState(221);
			match(IN);
			setState(222);
			range();
			setState(229);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==SEPARATOR) {
				{
				{
				setState(223);
				match(SEPARATOR);
				setState(224);
				match(STRING);
				setState(225);
				match(IN);
				setState(226);
				range();
				}
				}
				setState(231);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class SelectContext extends ParserRuleContext {
		public List<CalcContext> calc() {
			return getRuleContexts(CalcContext.class);
		}
		public CalcContext calc(int i) {
			return getRuleContext(CalcContext.class,i);
		}
		public SelectContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_select; }
	}

	public final SelectContext select() throws RecognitionException {
		SelectContext _localctx = new SelectContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_select);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(232);
			match(SELECT);
			setState(233);
			calc();
			setState(238);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==SEPARATOR) {
				{
				{
				setState(234);
				match(SEPARATOR);
				setState(235);
				calc();
				}
				}
				setState(240);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class FromContext extends ParserRuleContext {
		public List<NameContext> name() {
			return getRuleContexts(NameContext.class);
		}
		public NameContext name(int i) {
			return getRuleContext(NameContext.class,i);
		}
		public FromContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_from; }
	}

	public final FromContext from() throws RecognitionException {
		FromContext _localctx = new FromContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_from);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(241);
			match(FROM);
			setState(242);
			name();
			setState(247);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==SEPARATOR) {
				{
				{
				setState(243);
				match(SEPARATOR);
				setState(244);
				name();
				}
				}
				setState(249);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class WhereContext extends ParserRuleContext {
		public List<ConditionContext> condition() {
			return getRuleContexts(ConditionContext.class);
		}
		public ConditionContext condition(int i) {
			return getRuleContext(ConditionContext.class,i);
		}
		public WhereContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_where; }
	}

	public final WhereContext where() throws RecognitionException {
		WhereContext _localctx = new WhereContext(_ctx, getState());
		enterRule(_localctx, 44, RULE_where);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(250);
			match(WHERE);
			setState(251);
			condition();
			setState(256);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==AND) {
				{
				{
				setState(252);
				match(AND);
				setState(253);
				condition();
				}
				}
				setState(258);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class GroupbyContext extends ParserRuleContext {
		public List<NameableContext> nameable() {
			return getRuleContexts(NameableContext.class);
		}
		public NameableContext nameable(int i) {
			return getRuleContext(NameableContext.class,i);
		}
		public GroupbyContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_groupby; }
	}

	public final GroupbyContext groupby() throws RecognitionException {
		GroupbyContext _localctx = new GroupbyContext(_ctx, getState());
		enterRule(_localctx, 46, RULE_groupby);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(259);
			match(GROUP);
			setState(260);
			match(BY);
			setState(261);
			nameable();
			setState(266);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==SEPARATOR) {
				{
				{
				setState(262);
				match(SEPARATOR);
				setState(263);
				nameable();
				}
				}
				setState(268);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class OrderbyContext extends ParserRuleContext {
		public List<NameableContext> nameable() {
			return getRuleContexts(NameableContext.class);
		}
		public NameableContext nameable(int i) {
			return getRuleContext(NameableContext.class,i);
		}
		public List<SortOrderContext> sortOrder() {
			return getRuleContexts(SortOrderContext.class);
		}
		public SortOrderContext sortOrder(int i) {
			return getRuleContext(SortOrderContext.class,i);
		}
		public OrderbyContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_orderby; }
	}

	public final OrderbyContext orderby() throws RecognitionException {
		OrderbyContext _localctx = new OrderbyContext(_ctx, getState());
		enterRule(_localctx, 48, RULE_orderby);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(269);
			match(ORDER);
			setState(270);
			match(BY);
			setState(271);
			nameable();
			setState(272);
			sortOrder();
			setState(279);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==SEPARATOR) {
				{
				{
				setState(273);
				match(SEPARATOR);
				setState(274);
				nameable();
				setState(275);
				sortOrder();
				}
				}
				setState(281);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class LimitContext extends ParserRuleContext {
		public TerminalNode INT() { return getToken(ViewSQLParser.INT, 0); }
		public LimitContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_limit; }
	}

	public final LimitContext limit() throws RecognitionException {
		LimitContext _localctx = new LimitContext(_ctx, getState());
		enterRule(_localctx, 50, RULE_limit);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(282);
			match(LIMIT);
			setState(283);
			match(INT);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ViewContext extends ParserRuleContext {
		public GroupbyContext firstGroupBy;
		public GroupbyContext secondGroupBy;
		public CreateContext create() {
			return getRuleContext(CreateContext.class,0);
		}
		public SelectContext select() {
			return getRuleContext(SelectContext.class,0);
		}
		public FromContext from() {
			return getRuleContext(FromContext.class,0);
		}
		public List<GroupbyContext> groupby() {
			return getRuleContexts(GroupbyContext.class);
		}
		public GroupbyContext groupby(int i) {
			return getRuleContext(GroupbyContext.class,i);
		}
		public WithContext with() {
			return getRuleContext(WithContext.class,0);
		}
		public WhereContext where() {
			return getRuleContext(WhereContext.class,0);
		}
		public OrderbyContext orderby() {
			return getRuleContext(OrderbyContext.class,0);
		}
		public LimitContext limit() {
			return getRuleContext(LimitContext.class,0);
		}
		public ViewContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_view; }
	}

	public final ViewContext view() throws RecognitionException {
		ViewContext _localctx = new ViewContext(_ctx, getState());
		enterRule(_localctx, 52, RULE_view);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(285);
			create();
			setState(287);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==WITH) {
				{
				setState(286);
				with();
				}
			}

			setState(289);
			match(AS);
			setState(290);
			select();
			setState(291);
			from();
			setState(293);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==WHERE) {
				{
				setState(292);
				where();
				}
			}

			setState(296);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,20,_ctx) ) {
			case 1:
				{
				setState(295);
				((ViewContext)_localctx).firstGroupBy = groupby();
				}
				break;
			}
			setState(298);
			((ViewContext)_localctx).secondGroupBy = groupby();
			setState(300);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ORDER) {
				{
				setState(299);
				orderby();
				}
			}

			setState(303);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==LIMIT) {
				{
				setState(302);
				limit();
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class CheckContext extends ParserRuleContext {
		public Token type;
		public ConstantContext constant() {
			return getRuleContext(ConstantContext.class,0);
		}
		public CheckContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_check; }
	}

	public final CheckContext check() throws RecognitionException {
		CheckContext _localctx = new CheckContext(_ctx, getState());
		enterRule(_localctx, 54, RULE_check);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(305);
			match(CHECK);
			setState(306);
			((CheckContext)_localctx).type = _input.LT(1);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << HIGHER) | (1L << HIGHER_EQUAL) | (1L << LOWER) | (1L << LOWER_EQUAL))) != 0)) ) {
				((CheckContext)_localctx).type = (Token)_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(307);
			constant();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ForeignkeyContext extends ParserRuleContext {
		public NameContext tableName;
		public NameContext columnName;
		public List<NameContext> name() {
			return getRuleContexts(NameContext.class);
		}
		public NameContext name(int i) {
			return getRuleContext(NameContext.class,i);
		}
		public ForeignkeyContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_foreignkey; }
	}

	public final ForeignkeyContext foreignkey() throws RecognitionException {
		ForeignkeyContext _localctx = new ForeignkeyContext(_ctx, getState());
		enterRule(_localctx, 56, RULE_foreignkey);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(309);
			match(FOREIGN);
			setState(310);
			match(KEY);
			setState(311);
			match(REFERENCES);
			setState(312);
			((ForeignkeyContext)_localctx).tableName = name();
			setState(313);
			match(LEFT_P);
			setState(314);
			((ForeignkeyContext)_localctx).columnName = name();
			setState(315);
			match(RIGHT_P);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class PrimarykeyContext extends ParserRuleContext {
		public PrimarykeyContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_primarykey; }
	}

	public final PrimarykeyContext primarykey() throws RecognitionException {
		PrimarykeyContext _localctx = new PrimarykeyContext(_ctx, getState());
		enterRule(_localctx, 58, RULE_primarykey);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(317);
			match(PRIMARY);
			setState(318);
			match(KEY);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ConstraintContext extends ParserRuleContext {
		public PrimarykeyContext primarykey() {
			return getRuleContext(PrimarykeyContext.class,0);
		}
		public ForeignkeyContext foreignkey() {
			return getRuleContext(ForeignkeyContext.class,0);
		}
		public CheckContext check() {
			return getRuleContext(CheckContext.class,0);
		}
		public ConstraintContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_constraint; }
	}

	public final ConstraintContext constraint() throws RecognitionException {
		ConstraintContext _localctx = new ConstraintContext(_ctx, getState());
		enterRule(_localctx, 60, RULE_constraint);
		try {
			setState(323);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case PRIMARY:
				enterOuterAlt(_localctx, 1);
				{
				setState(320);
				primarykey();
				}
				break;
			case FOREIGN:
				enterOuterAlt(_localctx, 2);
				{
				setState(321);
				foreignkey();
				}
				break;
			case CHECK:
				enterOuterAlt(_localctx, 3);
				{
				setState(322);
				check();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ColumnsContext extends ParserRuleContext {
		public Token type;
		public TerminalNode STRING() { return getToken(ViewSQLParser.STRING, 0); }
		public TerminalNode COUNTER() { return getToken(ViewSQLParser.COUNTER, 0); }
		public TerminalNode INTEGER() { return getToken(ViewSQLParser.INTEGER, 0); }
		public TerminalNode BOOLEAN() { return getToken(ViewSQLParser.BOOLEAN, 0); }
		public TerminalNode VARCHAR() { return getToken(ViewSQLParser.VARCHAR, 0); }
		public TerminalNode DATE_TYPE() { return getToken(ViewSQLParser.DATE_TYPE, 0); }
		public ConstantContext constant() {
			return getRuleContext(ConstantContext.class,0);
		}
		public ConstraintContext constraint() {
			return getRuleContext(ConstraintContext.class,0);
		}
		public ColumnsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_columns; }
	}

	public final ColumnsContext columns() throws RecognitionException {
		ColumnsContext _localctx = new ColumnsContext(_ctx, getState());
		enterRule(_localctx, 62, RULE_columns);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(325);
			match(STRING);
			setState(326);
			((ColumnsContext)_localctx).type = _input.LT(1);
			_la = _input.LA(1);
			if ( !(((((_la - 60)) & ~0x3f) == 0 && ((1L << (_la - 60)) & ((1L << (INTEGER - 60)) | (1L << (COUNTER - 60)) | (1L << (BOOLEAN - 60)) | (1L << (VARCHAR - 60)) | (1L << (DATE_TYPE - 60)))) != 0)) ) {
				((ColumnsContext)_localctx).type = (Token)_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(329);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==DEFAULT) {
				{
				setState(327);
				match(DEFAULT);
				setState(328);
				constant();
				}
			}

			setState(332);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << PRIMARY) | (1L << FOREIGN) | (1L << CHECK))) != 0)) {
				{
				setState(331);
				constraint();
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class CreatetableContext extends ParserRuleContext {
		public Token policy;
		public NameContext name() {
			return getRuleContext(NameContext.class,0);
		}
		public List<ColumnsContext> columns() {
			return getRuleContexts(ColumnsContext.class);
		}
		public ColumnsContext columns(int i) {
			return getRuleContext(ColumnsContext.class,i);
		}
		public CreatetableContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_createtable; }
	}

	public final CreatetableContext createtable() throws RecognitionException {
		CreatetableContext _localctx = new CreatetableContext(_ctx, getState());
		enterRule(_localctx, 64, RULE_createtable);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(334);
			match(CREATE);
			setState(335);
			((CreatetableContext)_localctx).policy = _input.LT(1);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << AW) | (1L << RW) | (1L << LWW))) != 0)) ) {
				((CreatetableContext)_localctx).policy = (Token)_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(336);
			match(TABLE);
			setState(337);
			name();
			setState(338);
			match(LEFT_P);
			setState(344);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,26,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(339);
					columns();
					setState(340);
					match(SEPARATOR);
					}
					} 
				}
				setState(346);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,26,_ctx);
			}
			setState(347);
			columns();
			setState(348);
			match(RIGHT_P);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class CreateindexContext extends ParserRuleContext {
		public NameContext indexName;
		public NameContext tableName;
		public NameContext columnName;
		public List<NameContext> name() {
			return getRuleContexts(NameContext.class);
		}
		public NameContext name(int i) {
			return getRuleContext(NameContext.class,i);
		}
		public CreateindexContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_createindex; }
	}

	public final CreateindexContext createindex() throws RecognitionException {
		CreateindexContext _localctx = new CreateindexContext(_ctx, getState());
		enterRule(_localctx, 66, RULE_createindex);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(350);
			match(CREATE);
			setState(351);
			match(INDEX);
			setState(352);
			((CreateindexContext)_localctx).indexName = name();
			setState(353);
			match(ON);
			setState(354);
			((CreateindexContext)_localctx).tableName = name();
			setState(355);
			match(LEFT_P);
			setState(356);
			((CreateindexContext)_localctx).columnName = name();
			setState(357);
			match(RIGHT_P);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class DropContext extends ParserRuleContext {
		public NameContext name() {
			return getRuleContext(NameContext.class,0);
		}
		public DropContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_drop; }
	}

	public final DropContext drop() throws RecognitionException {
		DropContext _localctx = new DropContext(_ctx, getState());
		enterRule(_localctx, 68, RULE_drop);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(359);
			match(DROP);
			setState(360);
			match(TABLE);
			setState(361);
			name();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class DeleteContext extends ParserRuleContext {
		public NameContext name() {
			return getRuleContext(NameContext.class,0);
		}
		public WhereContext where() {
			return getRuleContext(WhereContext.class,0);
		}
		public DeleteContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_delete; }
	}

	public final DeleteContext delete() throws RecognitionException {
		DeleteContext _localctx = new DeleteContext(_ctx, getState());
		enterRule(_localctx, 70, RULE_delete);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(363);
			match(DELETE);
			setState(364);
			match(FROM);
			setState(365);
			name();
			setState(366);
			where();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class SetContext extends ParserRuleContext {
		public List<NameContext> name() {
			return getRuleContexts(NameContext.class);
		}
		public NameContext name(int i) {
			return getRuleContext(NameContext.class,i);
		}
		public List<ConstantContext> constant() {
			return getRuleContexts(ConstantContext.class);
		}
		public ConstantContext constant(int i) {
			return getRuleContext(ConstantContext.class,i);
		}
		public WhereContext where() {
			return getRuleContext(WhereContext.class,0);
		}
		public SetContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_set; }
	}

	public final SetContext set() throws RecognitionException {
		SetContext _localctx = new SetContext(_ctx, getState());
		enterRule(_localctx, 72, RULE_set);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(368);
			match(SET);
			setState(376);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,27,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(369);
					name();
					setState(370);
					match(EQUAL);
					setState(371);
					constant();
					setState(372);
					match(SEPARATOR);
					}
					} 
				}
				setState(378);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,27,_ctx);
			}
			setState(379);
			name();
			setState(380);
			match(EQUAL);
			setState(381);
			constant();
			setState(382);
			where();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class UpdateContext extends ParserRuleContext {
		public NameContext name() {
			return getRuleContext(NameContext.class,0);
		}
		public SetContext set() {
			return getRuleContext(SetContext.class,0);
		}
		public UpdateContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_update; }
	}

	public final UpdateContext update() throws RecognitionException {
		UpdateContext _localctx = new UpdateContext(_ctx, getState());
		enterRule(_localctx, 74, RULE_update);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(384);
			match(UPDATE);
			setState(385);
			name();
			setState(386);
			set();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ValuesContext extends ParserRuleContext {
		public List<ConstantContext> constant() {
			return getRuleContexts(ConstantContext.class);
		}
		public ConstantContext constant(int i) {
			return getRuleContext(ConstantContext.class,i);
		}
		public ValuesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_values; }
	}

	public final ValuesContext values() throws RecognitionException {
		ValuesContext _localctx = new ValuesContext(_ctx, getState());
		enterRule(_localctx, 76, RULE_values);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(388);
			match(VALUES);
			setState(389);
			match(LEFT_P);
			setState(395);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,28,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(390);
					constant();
					setState(391);
					match(SEPARATOR);
					}
					} 
				}
				setState(397);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,28,_ctx);
			}
			setState(398);
			constant();
			setState(399);
			match(RIGHT_P);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ColumnNamesContext extends ParserRuleContext {
		public List<NameContext> name() {
			return getRuleContexts(NameContext.class);
		}
		public NameContext name(int i) {
			return getRuleContext(NameContext.class,i);
		}
		public ColumnNamesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_columnNames; }
	}

	public final ColumnNamesContext columnNames() throws RecognitionException {
		ColumnNamesContext _localctx = new ColumnNamesContext(_ctx, getState());
		enterRule(_localctx, 78, RULE_columnNames);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(401);
			match(LEFT_P);
			setState(407);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,29,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(402);
					name();
					setState(403);
					match(SEPARATOR);
					}
					} 
				}
				setState(409);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,29,_ctx);
			}
			setState(410);
			name();
			setState(411);
			match(RIGHT_P);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class InsertContext extends ParserRuleContext {
		public NameContext name() {
			return getRuleContext(NameContext.class,0);
		}
		public ValuesContext values() {
			return getRuleContext(ValuesContext.class,0);
		}
		public ColumnNamesContext columnNames() {
			return getRuleContext(ColumnNamesContext.class,0);
		}
		public InsertContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_insert; }
	}

	public final InsertContext insert() throws RecognitionException {
		InsertContext _localctx = new InsertContext(_ctx, getState());
		enterRule(_localctx, 80, RULE_insert);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(413);
			match(INSERT);
			setState(414);
			match(INTO);
			setState(415);
			name();
			setState(417);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==LEFT_P) {
				{
				setState(416);
				columnNames();
				}
			}

			setState(419);
			values();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class QueryContext extends ParserRuleContext {
		public SelectContext select() {
			return getRuleContext(SelectContext.class,0);
		}
		public FromContext from() {
			return getRuleContext(FromContext.class,0);
		}
		public WhereContext where() {
			return getRuleContext(WhereContext.class,0);
		}
		public QueryContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_query; }
	}

	public final QueryContext query() throws RecognitionException {
		QueryContext _localctx = new QueryContext(_ctx, getState());
		enterRule(_localctx, 82, RULE_query);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(421);
			select();
			setState(422);
			from();
			setState(424);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==WHERE) {
				{
				setState(423);
				where();
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class StatementContext extends ParserRuleContext {
		public ViewContext view() {
			return getRuleContext(ViewContext.class,0);
		}
		public CreatetableContext createtable() {
			return getRuleContext(CreatetableContext.class,0);
		}
		public CreateindexContext createindex() {
			return getRuleContext(CreateindexContext.class,0);
		}
		public InsertContext insert() {
			return getRuleContext(InsertContext.class,0);
		}
		public UpdateContext update() {
			return getRuleContext(UpdateContext.class,0);
		}
		public DeleteContext delete() {
			return getRuleContext(DeleteContext.class,0);
		}
		public DropContext drop() {
			return getRuleContext(DropContext.class,0);
		}
		public QueryContext query() {
			return getRuleContext(QueryContext.class,0);
		}
		public StatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_statement; }
	}

	public final StatementContext statement() throws RecognitionException {
		StatementContext _localctx = new StatementContext(_ctx, getState());
		enterRule(_localctx, 84, RULE_statement);
		try {
			setState(434);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,32,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(426);
				view();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(427);
				createtable();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(428);
				createindex();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(429);
				insert();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(430);
				update();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(431);
				delete();
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(432);
				drop();
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(433);
				query();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class StartContext extends ParserRuleContext {
		public StatementContext statement() {
			return getRuleContext(StatementContext.class,0);
		}
		public TerminalNode EOF() { return getToken(ViewSQLParser.EOF, 0); }
		public StartContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_start; }
	}

	public final StartContext start() throws RecognitionException {
		StartContext _localctx = new StartContext(_ctx, getState());
		enterRule(_localctx, 86, RULE_start);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(436);
			statement();
			setState(437);
			match(EOF);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 7:
			return math_sempred((MathContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean math_sempred(MathContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 2);
		case 1:
			return precpred(_ctx, 1);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3H\u01ba\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\3\2\3\2\3\3\3\3\3\3\3\3\3\3\3\3\3\3\7\3d\n\3\f\3\16\3g\13\3"+
		"\3\3\5\3j\n\3\3\4\3\4\3\5\3\5\3\5\3\5\3\6\3\6\3\6\3\6\3\7\3\7\3\7\5\7"+
		"y\n\7\3\b\3\b\3\b\3\b\3\b\3\b\3\t\3\t\3\t\5\t\u0084\n\t\3\t\3\t\3\t\3"+
		"\t\3\t\3\t\5\t\u008c\n\t\3\t\3\t\3\t\3\t\3\t\3\t\7\t\u0094\n\t\f\t\16"+
		"\t\u0097\13\t\3\n\3\n\3\n\3\n\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\f\3"+
		"\f\3\f\3\f\5\f\u00a8\n\f\3\f\3\f\3\f\3\f\3\r\3\r\3\r\3\r\3\r\3\r\5\r\u00b4"+
		"\n\r\3\16\3\16\5\16\u00b8\n\16\3\17\3\17\3\17\3\17\3\20\3\20\3\21\3\21"+
		"\3\21\3\21\3\21\3\21\3\21\3\22\3\22\3\22\3\22\7\22\u00cb\n\22\f\22\16"+
		"\22\u00ce\13\22\3\22\3\22\3\23\3\23\5\23\u00d4\n\23\3\24\3\24\3\24\3\24"+
		"\3\24\3\24\3\24\3\24\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\25\7\25\u00e6"+
		"\n\25\f\25\16\25\u00e9\13\25\3\26\3\26\3\26\3\26\7\26\u00ef\n\26\f\26"+
		"\16\26\u00f2\13\26\3\27\3\27\3\27\3\27\7\27\u00f8\n\27\f\27\16\27\u00fb"+
		"\13\27\3\30\3\30\3\30\3\30\7\30\u0101\n\30\f\30\16\30\u0104\13\30\3\31"+
		"\3\31\3\31\3\31\3\31\7\31\u010b\n\31\f\31\16\31\u010e\13\31\3\32\3\32"+
		"\3\32\3\32\3\32\3\32\3\32\3\32\7\32\u0118\n\32\f\32\16\32\u011b\13\32"+
		"\3\33\3\33\3\33\3\34\3\34\5\34\u0122\n\34\3\34\3\34\3\34\3\34\5\34\u0128"+
		"\n\34\3\34\5\34\u012b\n\34\3\34\3\34\5\34\u012f\n\34\3\34\5\34\u0132\n"+
		"\34\3\35\3\35\3\35\3\35\3\36\3\36\3\36\3\36\3\36\3\36\3\36\3\36\3\37\3"+
		"\37\3\37\3 \3 \3 \5 \u0146\n \3!\3!\3!\3!\5!\u014c\n!\3!\5!\u014f\n!\3"+
		"\"\3\"\3\"\3\"\3\"\3\"\3\"\3\"\7\"\u0159\n\"\f\"\16\"\u015c\13\"\3\"\3"+
		"\"\3\"\3#\3#\3#\3#\3#\3#\3#\3#\3#\3$\3$\3$\3$\3%\3%\3%\3%\3%\3&\3&\3&"+
		"\3&\3&\3&\7&\u0179\n&\f&\16&\u017c\13&\3&\3&\3&\3&\3&\3\'\3\'\3\'\3\'"+
		"\3(\3(\3(\3(\3(\7(\u018c\n(\f(\16(\u018f\13(\3(\3(\3(\3)\3)\3)\3)\7)\u0198"+
		"\n)\f)\16)\u019b\13)\3)\3)\3)\3*\3*\3*\3*\5*\u01a4\n*\3*\3*\3+\3+\3+\5"+
		"+\u01ab\n+\3,\3,\3,\3,\3,\3,\3,\3,\5,\u01b5\n,\3-\3-\3-\3-\2\3\20.\2\4"+
		"\6\b\n\f\16\20\22\24\26\30\32\34\36 \"$&(*,.\60\62\64\668:<>@BDFHJLNP"+
		"RTVX\2\f\4\2EEHH\3\2\7\n\3\2\5\6\3\2\3\4\3\2\17\24\3\2\r\16\3\2\32\33"+
		"\3\2\20\23\3\2>B\3\2:<\2\u01be\2Z\3\2\2\2\4i\3\2\2\2\6k\3\2\2\2\bm\3\2"+
		"\2\2\nq\3\2\2\2\fx\3\2\2\2\16z\3\2\2\2\20\u008b\3\2\2\2\22\u0098\3\2\2"+
		"\2\24\u009c\3\2\2\2\26\u00a3\3\2\2\2\30\u00b3\3\2\2\2\32\u00b7\3\2\2\2"+
		"\34\u00b9\3\2\2\2\36\u00bd\3\2\2\2 \u00bf\3\2\2\2\"\u00c6\3\2\2\2$\u00d3"+
		"\3\2\2\2&\u00d5\3\2\2\2(\u00dd\3\2\2\2*\u00ea\3\2\2\2,\u00f3\3\2\2\2."+
		"\u00fc\3\2\2\2\60\u0105\3\2\2\2\62\u010f\3\2\2\2\64\u011c\3\2\2\2\66\u011f"+
		"\3\2\2\28\u0133\3\2\2\2:\u0137\3\2\2\2<\u013f\3\2\2\2>\u0145\3\2\2\2@"+
		"\u0147\3\2\2\2B\u0150\3\2\2\2D\u0160\3\2\2\2F\u0169\3\2\2\2H\u016d\3\2"+
		"\2\2J\u0172\3\2\2\2L\u0182\3\2\2\2N\u0186\3\2\2\2P\u0193\3\2\2\2R\u019f"+
		"\3\2\2\2T\u01a7\3\2\2\2V\u01b4\3\2\2\2X\u01b6\3\2\2\2Z[\7E\2\2[\3\3\2"+
		"\2\2\\j\7C\2\2]j\7D\2\2^j\7F\2\2_j\7G\2\2`a\7\34\2\2ae\7E\2\2bd\t\2\2"+
		"\2cb\3\2\2\2dg\3\2\2\2ec\3\2\2\2ef\3\2\2\2fh\3\2\2\2ge\3\2\2\2hj\7\34"+
		"\2\2i\\\3\2\2\2i]\3\2\2\2i^\3\2\2\2i_\3\2\2\2i`\3\2\2\2j\5\3\2\2\2kl\t"+
		"\3\2\2l\7\3\2\2\2mn\5\2\2\2no\7\25\2\2op\5\2\2\2p\t\3\2\2\2qr\7\32\2\2"+
		"rs\5\2\2\2st\7\33\2\2t\13\3\2\2\2uy\5\2\2\2vy\5\b\5\2wy\5\n\6\2xu\3\2"+
		"\2\2xv\3\2\2\2xw\3\2\2\2y\r\3\2\2\2z{\7+\2\2{|\7,\2\2|}\7\30\2\2}~\5\f"+
		"\7\2~\177\7\31\2\2\177\17\3\2\2\2\u0080\u0083\b\t\1\2\u0081\u0084\5\f"+
		"\7\2\u0082\u0084\5\4\3\2\u0083\u0081\3\2\2\2\u0083\u0082\3\2\2\2\u0084"+
		"\u008c\3\2\2\2\u0085\u0086\7\4\2\2\u0086\u008c\5\20\t\6\u0087\u0088\7"+
		"\30\2\2\u0088\u0089\5\20\t\2\u0089\u008a\7\31\2\2\u008a\u008c\3\2\2\2"+
		"\u008b\u0080\3\2\2\2\u008b\u0085\3\2\2\2\u008b\u0087\3\2\2\2\u008c\u0095"+
		"\3\2\2\2\u008d\u008e\f\4\2\2\u008e\u008f\t\4\2\2\u008f\u0094\5\20\t\5"+
		"\u0090\u0091\f\3\2\2\u0091\u0092\t\5\2\2\u0092\u0094\5\20\t\4\u0093\u008d"+
		"\3\2\2\2\u0093\u0090\3\2\2\2\u0094\u0097\3\2\2\2\u0095\u0093\3\2\2\2\u0095"+
		"\u0096\3\2\2\2\u0096\21\3\2\2\2\u0097\u0095\3\2\2\2\u0098\u0099\5\20\t"+
		"\2\u0099\u009a\7\37\2\2\u009a\u009b\5\2\2\2\u009b\23\3\2\2\2\u009c\u009d"+
		"\5\6\4\2\u009d\u009e\7\30\2\2\u009e\u009f\5\20\t\2\u009f\u00a0\7\31\2"+
		"\2\u00a0\u00a1\7\37\2\2\u00a1\u00a2\5\2\2\2\u00a2\25\3\2\2\2\u00a3\u00a4"+
		"\7\13\2\2\u00a4\u00a7\7\30\2\2\u00a5\u00a8\5\f\7\2\u00a6\u00a8\7\5\2\2"+
		"\u00a7\u00a5\3\2\2\2\u00a7\u00a6\3\2\2\2\u00a8\u00a9\3\2\2\2\u00a9\u00aa"+
		"\7\31\2\2\u00aa\u00ab\7\37\2\2\u00ab\u00ac\5\2\2\2\u00ac\27\3\2\2\2\u00ad"+
		"\u00b4\5\16\b\2\u00ae\u00b4\5\f\7\2\u00af\u00b4\5\22\n\2\u00b0\u00b4\5"+
		"\24\13\2\u00b1\u00b4\5\26\f\2\u00b2\u00b4\7\5\2\2\u00b3\u00ad\3\2\2\2"+
		"\u00b3\u00ae\3\2\2\2\u00b3\u00af\3\2\2\2\u00b3\u00b0\3\2\2\2\u00b3\u00b1"+
		"\3\2\2\2\u00b3\u00b2\3\2\2\2\u00b4\31\3\2\2\2\u00b5\u00b8\t\6\2\2\u00b6"+
		"\u00b8\5\2\2\2\u00b7\u00b5\3\2\2\2\u00b7\u00b6\3\2\2\2\u00b8\33\3\2\2"+
		"\2\u00b9\u00ba\5\f\7\2\u00ba\u00bb\5\32\16\2\u00bb\u00bc\5\20\t\2\u00bc"+
		"\35\3\2\2\2\u00bd\u00be\t\7\2\2\u00be\37\3\2\2\2\u00bf\u00c0\t\b\2\2\u00c0"+
		"\u00c1\5\4\3\2\u00c1\u00c2\7\27\2\2\u00c2\u00c3\5\4\3\2\u00c3\u00c4\3"+
		"\2\2\2\u00c4\u00c5\t\b\2\2\u00c5!\3\2\2\2\u00c6\u00c7\7\35\2\2\u00c7\u00cc"+
		"\5\4\3\2\u00c8\u00c9\7\26\2\2\u00c9\u00cb\5\4\3\2\u00ca\u00c8\3\2\2\2"+
		"\u00cb\u00ce\3\2\2\2\u00cc\u00ca\3\2\2\2\u00cc\u00cd\3\2\2\2\u00cd\u00cf"+
		"\3\2\2\2\u00ce\u00cc\3\2\2\2\u00cf\u00d0\7\36\2\2\u00d0#\3\2\2\2\u00d1"+
		"\u00d4\5 \21\2\u00d2\u00d4\5\"\22\2\u00d3\u00d1\3\2\2\2\u00d3\u00d2\3"+
		"\2\2\2\u00d4%\3\2\2\2\u00d5\u00d6\7 \2\2\u00d6\u00d7\7!\2\2\u00d7\u00d8"+
		"\7\30\2\2\u00d8\u00d9\7E\2\2\u00d9\u00da\7\26\2\2\u00da\u00db\7E\2\2\u00db"+
		"\u00dc\7\31\2\2\u00dc\'\3\2\2\2\u00dd\u00de\7)\2\2\u00de\u00df\7E\2\2"+
		"\u00df\u00e0\7*\2\2\u00e0\u00e7\5$\23\2\u00e1\u00e2\7\26\2\2\u00e2\u00e3"+
		"\7E\2\2\u00e3\u00e4\7*\2\2\u00e4\u00e6\5$\23\2\u00e5\u00e1\3\2\2\2\u00e6"+
		"\u00e9\3\2\2\2\u00e7\u00e5\3\2\2\2\u00e7\u00e8\3\2\2\2\u00e8)\3\2\2\2"+
		"\u00e9\u00e7\3\2\2\2\u00ea\u00eb\7\"\2\2\u00eb\u00f0\5\30\r\2\u00ec\u00ed"+
		"\7\26\2\2\u00ed\u00ef\5\30\r\2\u00ee\u00ec\3\2\2\2\u00ef\u00f2\3\2\2\2"+
		"\u00f0\u00ee\3\2\2\2\u00f0\u00f1\3\2\2\2\u00f1+\3\2\2\2\u00f2\u00f0\3"+
		"\2\2\2\u00f3\u00f4\7#\2\2\u00f4\u00f9\5\2\2\2\u00f5\u00f6\7\26\2\2\u00f6"+
		"\u00f8\5\2\2\2\u00f7\u00f5\3\2\2\2\u00f8\u00fb\3\2\2\2\u00f9\u00f7\3\2"+
		"\2\2\u00f9\u00fa\3\2\2\2\u00fa-\3\2\2\2\u00fb\u00f9\3\2\2\2\u00fc\u00fd"+
		"\7$\2\2\u00fd\u0102\5\34\17\2\u00fe\u00ff\7\f\2\2\u00ff\u0101\5\34\17"+
		"\2\u0100\u00fe\3\2\2\2\u0101\u0104\3\2\2\2\u0102\u0100\3\2\2\2\u0102\u0103"+
		"\3\2\2\2\u0103/\3\2\2\2\u0104\u0102\3\2\2\2\u0105\u0106\7%\2\2\u0106\u0107"+
		"\7&\2\2\u0107\u010c\5\f\7\2\u0108\u0109\7\26\2\2\u0109\u010b\5\f\7\2\u010a"+
		"\u0108\3\2\2\2\u010b\u010e\3\2\2\2\u010c\u010a\3\2\2\2\u010c\u010d\3\2"+
		"\2\2\u010d\61\3\2\2\2\u010e\u010c\3\2\2\2\u010f\u0110\7\'\2\2\u0110\u0111"+
		"\7&\2\2\u0111\u0112\5\f\7\2\u0112\u0119\5\36\20\2\u0113\u0114\7\26\2\2"+
		"\u0114\u0115\5\f\7\2\u0115\u0116\5\36\20\2\u0116\u0118\3\2\2\2\u0117\u0113"+
		"\3\2\2\2\u0118\u011b\3\2\2\2\u0119\u0117\3\2\2\2\u0119\u011a\3\2\2\2\u011a"+
		"\63\3\2\2\2\u011b\u0119\3\2\2\2\u011c\u011d\7(\2\2\u011d\u011e\7F\2\2"+
		"\u011e\65\3\2\2\2\u011f\u0121\5&\24\2\u0120\u0122\5(\25\2\u0121\u0120"+
		"\3\2\2\2\u0121\u0122\3\2\2\2\u0122\u0123\3\2\2\2\u0123\u0124\7\37\2\2"+
		"\u0124\u0125\5*\26\2\u0125\u0127\5,\27\2\u0126\u0128\5.\30\2\u0127\u0126"+
		"\3\2\2\2\u0127\u0128\3\2\2\2\u0128\u012a\3\2\2\2\u0129\u012b\5\60\31\2"+
		"\u012a\u0129\3\2\2\2\u012a\u012b\3\2\2\2\u012b\u012c\3\2\2\2\u012c\u012e"+
		"\5\60\31\2\u012d\u012f\5\62\32\2\u012e\u012d\3\2\2\2\u012e\u012f\3\2\2"+
		"\2\u012f\u0131\3\2\2\2\u0130\u0132\5\64\33\2\u0131\u0130\3\2\2\2\u0131"+
		"\u0132\3\2\2\2\u0132\67\3\2\2\2\u0133\u0134\78\2\2\u0134\u0135\t\t\2\2"+
		"\u0135\u0136\5\4\3\2\u01369\3\2\2\2\u0137\u0138\7\65\2\2\u0138\u0139\7"+
		",\2\2\u0139\u013a\7\66\2\2\u013a\u013b\5\2\2\2\u013b\u013c\7\30\2\2\u013c"+
		"\u013d\5\2\2\2\u013d\u013e\7\31\2\2\u013e;\3\2\2\2\u013f\u0140\7+\2\2"+
		"\u0140\u0141\7,\2\2\u0141=\3\2\2\2\u0142\u0146\5<\37\2\u0143\u0146\5:"+
		"\36\2\u0144\u0146\58\35\2\u0145\u0142\3\2\2\2\u0145\u0143\3\2\2\2\u0145"+
		"\u0144\3\2\2\2\u0146?\3\2\2\2\u0147\u0148\7E\2\2\u0148\u014b\t\n\2\2\u0149"+
		"\u014a\7=\2\2\u014a\u014c\5\4\3\2\u014b\u0149\3\2\2\2\u014b\u014c\3\2"+
		"\2\2\u014c\u014e\3\2\2\2\u014d\u014f\5> \2\u014e\u014d\3\2\2\2\u014e\u014f"+
		"\3\2\2\2\u014fA\3\2\2\2\u0150\u0151\7 \2\2\u0151\u0152\t\13\2\2\u0152"+
		"\u0153\7-\2\2\u0153\u0154\5\2\2\2\u0154\u015a\7\30\2\2\u0155\u0156\5@"+
		"!\2\u0156\u0157\7\26\2\2\u0157\u0159\3\2\2\2\u0158\u0155\3\2\2\2\u0159"+
		"\u015c\3\2\2\2\u015a\u0158\3\2\2\2\u015a\u015b\3\2\2\2\u015b\u015d\3\2"+
		"\2\2\u015c\u015a\3\2\2\2\u015d\u015e\5@!\2\u015e\u015f\7\31\2\2\u015f"+
		"C\3\2\2\2\u0160\u0161\7 \2\2\u0161\u0162\7.\2\2\u0162\u0163\5\2\2\2\u0163"+
		"\u0164\79\2\2\u0164\u0165\5\2\2\2\u0165\u0166\7\30\2\2\u0166\u0167\5\2"+
		"\2\2\u0167\u0168\7\31\2\2\u0168E\3\2\2\2\u0169\u016a\7\63\2\2\u016a\u016b"+
		"\7-\2\2\u016b\u016c\5\2\2\2\u016cG\3\2\2\2\u016d\u016e\7\60\2\2\u016e"+
		"\u016f\7#\2\2\u016f\u0170\5\2\2\2\u0170\u0171\5.\30\2\u0171I\3\2\2\2\u0172"+
		"\u017a\7\64\2\2\u0173\u0174\5\2\2\2\u0174\u0175\7\17\2\2\u0175\u0176\5"+
		"\4\3\2\u0176\u0177\7\26\2\2\u0177\u0179\3\2\2\2\u0178\u0173\3\2\2\2\u0179"+
		"\u017c\3\2\2\2\u017a\u0178\3\2\2\2\u017a\u017b\3\2\2\2\u017b\u017d\3\2"+
		"\2\2\u017c\u017a\3\2\2\2\u017d\u017e\5\2\2\2\u017e\u017f\7\17\2\2\u017f"+
		"\u0180\5\4\3\2\u0180\u0181\5.\30\2\u0181K\3\2\2\2\u0182\u0183\7/\2\2\u0183"+
		"\u0184\5\2\2\2\u0184\u0185\5J&\2\u0185M\3\2\2\2\u0186\u0187\7\67\2\2\u0187"+
		"\u018d\7\30\2\2\u0188\u0189\5\4\3\2\u0189\u018a\7\26\2\2\u018a\u018c\3"+
		"\2\2\2\u018b\u0188\3\2\2\2\u018c\u018f\3\2\2\2\u018d\u018b\3\2\2\2\u018d"+
		"\u018e\3\2\2\2\u018e\u0190\3\2\2\2\u018f\u018d\3\2\2\2\u0190\u0191\5\4"+
		"\3\2\u0191\u0192\7\31\2\2\u0192O\3\2\2\2\u0193\u0199\7\30\2\2\u0194\u0195"+
		"\5\2\2\2\u0195\u0196\7\26\2\2\u0196\u0198\3\2\2\2\u0197\u0194\3\2\2\2"+
		"\u0198\u019b\3\2\2\2\u0199\u0197\3\2\2\2\u0199\u019a\3\2\2\2\u019a\u019c"+
		"\3\2\2\2\u019b\u0199\3\2\2\2\u019c\u019d\5\2\2\2\u019d\u019e\7\31\2\2"+
		"\u019eQ\3\2\2\2\u019f\u01a0\7\61\2\2\u01a0\u01a1\7\62\2\2\u01a1\u01a3"+
		"\5\2\2\2\u01a2\u01a4\5P)\2\u01a3\u01a2\3\2\2\2\u01a3\u01a4\3\2\2\2\u01a4"+
		"\u01a5\3\2\2\2\u01a5\u01a6\5N(\2\u01a6S\3\2\2\2\u01a7\u01a8\5*\26\2\u01a8"+
		"\u01aa\5,\27\2\u01a9\u01ab\5.\30\2\u01aa\u01a9\3\2\2\2\u01aa\u01ab\3\2"+
		"\2\2\u01abU\3\2\2\2\u01ac\u01b5\5\66\34\2\u01ad\u01b5\5B\"\2\u01ae\u01b5"+
		"\5D#\2\u01af\u01b5\5R*\2\u01b0\u01b5\5L\'\2\u01b1\u01b5\5H%\2\u01b2\u01b5"+
		"\5F$\2\u01b3\u01b5\5T+\2\u01b4\u01ac\3\2\2\2\u01b4\u01ad\3\2\2\2\u01b4"+
		"\u01ae\3\2\2\2\u01b4\u01af\3\2\2\2\u01b4\u01b0\3\2\2\2\u01b4\u01b1\3\2"+
		"\2\2\u01b4\u01b2\3\2\2\2\u01b4\u01b3\3\2\2\2\u01b5W\3\2\2\2\u01b6\u01b7"+
		"\5V,\2\u01b7\u01b8\7\2\2\3\u01b8Y\3\2\2\2#eix\u0083\u008b\u0093\u0095"+
		"\u00a7\u00b3\u00b7\u00cc\u00d3\u00e7\u00f0\u00f9\u0102\u010c\u0119\u0121"+
		"\u0127\u012a\u012e\u0131\u0145\u014b\u014e\u015a\u017a\u018d\u0199\u01a3"+
		"\u01aa\u01b4";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}