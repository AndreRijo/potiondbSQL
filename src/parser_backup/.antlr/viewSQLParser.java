// Generated from /Users/a.rijo/Documents/University_6th_year/goProjects/sqlToKeyValue/src/parser/ViewSQL.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue"})
public class ViewSQLParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

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
		REFERENCES=52, VALUES=53, CHECK=54, ON=55, AW=56, RW=57, LWW=58, MW=59, 
		EW=60, DW=61, DEFAULT=62, INTEGER=63, COUNTER=64, BOOLEAN=65, VARCHAR=66, 
		DATE_TYPE=67, DATE=68, BOOL=69, STRING=70, INT=71, FLOAT=72, WHITESPACE=73;
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
	private static String[] makeRuleNames() {
		return new String[] {
			"name", "constant", "aggrFunc", "field", "parameter", "nameable", "key", 
			"math", "asClause", "aggregation", "count", "calc", "comp", "condition", 
			"sortOrder", "continuousRange", "sparseRange", "range", "create", "with", 
			"select", "from", "where", "groupby", "orderby", "limit", "view", "check", 
			"foreignkey", "primarykey", "constraint", "columns", "createtable", "createindex", 
			"drop", "delete", "set", "update", "values", "columnNames", "insert", 
			"query", "statement", "start"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'+'", "'-'", "'*'", "'/'", "'SUM'", "'AVG'", "'MAX'", "'MIN'", 
			"'COUNT'", "'AND'", "'ASC'", "'DESC'", "'='", "'>'", "'>='", "'<'", "'<='", 
			"'!='", "'.'", "','", "':'", "'('", "')'", "'['", "']'", "'\"'", "'{'", 
			"'}'", "'AS'", "'CREATE'", "'VIEW'", "'SELECT'", "'FROM'", "'WHERE'", 
			"'GROUP'", "'BY'", "'ORDER'", "'LIMIT'", "'WITH'", "'IN'", "'PRIMARY'", 
			"'KEY'", "'TABLE'", "'INDEX'", "'UPDATE'", "'DELETE'", "'INSERT'", "'INTO'", 
			"'DROP'", "'SET'", "'FOREIGN'", "'REFERENCES'", "'VALUES'", "'CHECK'", 
			"'ON'", "'AW'", "'RW'", "'LWW'", "'MW'", "'EW'", "'DW'", "'DEFAULT'", 
			"'integer'", "'counter'", "'boolean'", "'varchar'", "'date'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "ADD", "SUB", "MULT", "DIV", "SUM", "AVG", "MAX", "MIN", "COUNT", 
			"AND", "ASC", "DESC", "EQUAL", "HIGHER", "HIGHER_EQUAL", "LOWER", "LOWER_EQUAL", 
			"NOT_EQUAL", "DOT", "SEPARATOR", "RANGE_SEP", "LEFT_P", "RIGHT_P", "LEFT_SP", 
			"RIGHT_SP", "INV_COMMA", "LEFT_CURLY", "RIGHT_CURLY", "AS", "CREATE", 
			"VIEW", "SELECT", "FROM", "WHERE", "GROUP", "BY", "ORDER", "LIMIT", "WITH", 
			"IN", "PRIMARY", "KEY", "TABLE", "INDEX", "UPDATE", "DELETE", "INSERT", 
			"INTO", "DROP", "SET", "FOREIGN", "REFERENCES", "VALUES", "CHECK", "ON", 
			"AW", "RW", "LWW", "MW", "EW", "DW", "DEFAULT", "INTEGER", "COUNTER", 
			"BOOLEAN", "VARCHAR", "DATE_TYPE", "DATE", "BOOL", "STRING", "INT", "FLOAT", 
			"WHITESPACE"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
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

	@SuppressWarnings("CheckReturnValue")
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

	@SuppressWarnings("CheckReturnValue")
	public static class ConstantContext extends ParserRuleContext {
		public TerminalNode DATE() { return getToken(ViewSQLParser.DATE, 0); }
		public TerminalNode BOOL() { return getToken(ViewSQLParser.BOOL, 0); }
		public TerminalNode INT() { return getToken(ViewSQLParser.INT, 0); }
		public TerminalNode FLOAT() { return getToken(ViewSQLParser.FLOAT, 0); }
		public List<TerminalNode> INV_COMMA() { return getTokens(ViewSQLParser.INV_COMMA); }
		public TerminalNode INV_COMMA(int i) {
			return getToken(ViewSQLParser.INV_COMMA, i);
		}
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

	@SuppressWarnings("CheckReturnValue")
	public static class AggrFuncContext extends ParserRuleContext {
		public TerminalNode SUM() { return getToken(ViewSQLParser.SUM, 0); }
		public TerminalNode AVG() { return getToken(ViewSQLParser.AVG, 0); }
		public TerminalNode MAX() { return getToken(ViewSQLParser.MAX, 0); }
		public TerminalNode MIN() { return getToken(ViewSQLParser.MIN, 0); }
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
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 480L) != 0)) ) {
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

	@SuppressWarnings("CheckReturnValue")
	public static class FieldContext extends ParserRuleContext {
		public List<NameContext> name() {
			return getRuleContexts(NameContext.class);
		}
		public NameContext name(int i) {
			return getRuleContext(NameContext.class,i);
		}
		public TerminalNode DOT() { return getToken(ViewSQLParser.DOT, 0); }
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

	@SuppressWarnings("CheckReturnValue")
	public static class ParameterContext extends ParserRuleContext {
		public TerminalNode LEFT_SP() { return getToken(ViewSQLParser.LEFT_SP, 0); }
		public NameContext name() {
			return getRuleContext(NameContext.class,0);
		}
		public TerminalNode RIGHT_SP() { return getToken(ViewSQLParser.RIGHT_SP, 0); }
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

	@SuppressWarnings("CheckReturnValue")
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

	@SuppressWarnings("CheckReturnValue")
	public static class KeyContext extends ParserRuleContext {
		public TerminalNode PRIMARY() { return getToken(ViewSQLParser.PRIMARY, 0); }
		public TerminalNode KEY() { return getToken(ViewSQLParser.KEY, 0); }
		public TerminalNode LEFT_P() { return getToken(ViewSQLParser.LEFT_P, 0); }
		public NameableContext nameable() {
			return getRuleContext(NameableContext.class,0);
		}
		public TerminalNode RIGHT_P() { return getToken(ViewSQLParser.RIGHT_P, 0); }
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

	@SuppressWarnings("CheckReturnValue")
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
	@SuppressWarnings("CheckReturnValue")
	public static class AddOrSubContext extends MathContext {
		public Token opType;
		public List<MathContext> math() {
			return getRuleContexts(MathContext.class);
		}
		public MathContext math(int i) {
			return getRuleContext(MathContext.class,i);
		}
		public TerminalNode ADD() { return getToken(ViewSQLParser.ADD, 0); }
		public TerminalNode SUB() { return getToken(ViewSQLParser.SUB, 0); }
		public AddOrSubContext(MathContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ValueContext extends MathContext {
		public NameableContext nameable() {
			return getRuleContext(NameableContext.class,0);
		}
		public ConstantContext constant() {
			return getRuleContext(ConstantContext.class,0);
		}
		public ValueContext(MathContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class MinusContext extends MathContext {
		public TerminalNode SUB() { return getToken(ViewSQLParser.SUB, 0); }
		public MathContext math() {
			return getRuleContext(MathContext.class,0);
		}
		public MinusContext(MathContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ParenthesesContext extends MathContext {
		public TerminalNode LEFT_P() { return getToken(ViewSQLParser.LEFT_P, 0); }
		public MathContext math() {
			return getRuleContext(MathContext.class,0);
		}
		public TerminalNode RIGHT_P() { return getToken(ViewSQLParser.RIGHT_P, 0); }
		public ParenthesesContext(MathContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class MultOrDivContext extends MathContext {
		public Token opType;
		public List<MathContext> math() {
			return getRuleContexts(MathContext.class);
		}
		public MathContext math(int i) {
			return getRuleContext(MathContext.class,i);
		}
		public TerminalNode MULT() { return getToken(ViewSQLParser.MULT, 0); }
		public TerminalNode DIV() { return getToken(ViewSQLParser.DIV, 0); }
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

	@SuppressWarnings("CheckReturnValue")
	public static class AsClauseContext extends ParserRuleContext {
		public MathContext math() {
			return getRuleContext(MathContext.class,0);
		}
		public TerminalNode AS() { return getToken(ViewSQLParser.AS, 0); }
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

	@SuppressWarnings("CheckReturnValue")
	public static class AggregationContext extends ParserRuleContext {
		public AggrFuncContext aggrFunc() {
			return getRuleContext(AggrFuncContext.class,0);
		}
		public TerminalNode LEFT_P() { return getToken(ViewSQLParser.LEFT_P, 0); }
		public MathContext math() {
			return getRuleContext(MathContext.class,0);
		}
		public TerminalNode RIGHT_P() { return getToken(ViewSQLParser.RIGHT_P, 0); }
		public TerminalNode AS() { return getToken(ViewSQLParser.AS, 0); }
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

	@SuppressWarnings("CheckReturnValue")
	public static class CountContext extends ParserRuleContext {
		public TerminalNode COUNT() { return getToken(ViewSQLParser.COUNT, 0); }
		public TerminalNode LEFT_P() { return getToken(ViewSQLParser.LEFT_P, 0); }
		public TerminalNode RIGHT_P() { return getToken(ViewSQLParser.RIGHT_P, 0); }
		public TerminalNode AS() { return getToken(ViewSQLParser.AS, 0); }
		public NameContext name() {
			return getRuleContext(NameContext.class,0);
		}
		public NameableContext nameable() {
			return getRuleContext(NameableContext.class,0);
		}
		public TerminalNode MULT() { return getToken(ViewSQLParser.MULT, 0); }
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

	@SuppressWarnings("CheckReturnValue")
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
		public TerminalNode MULT() { return getToken(ViewSQLParser.MULT, 0); }
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

	@SuppressWarnings("CheckReturnValue")
	public static class CompContext extends ParserRuleContext {
		public Token opType;
		public TerminalNode EQUAL() { return getToken(ViewSQLParser.EQUAL, 0); }
		public TerminalNode HIGHER() { return getToken(ViewSQLParser.HIGHER, 0); }
		public TerminalNode LOWER() { return getToken(ViewSQLParser.LOWER, 0); }
		public TerminalNode HIGHER_EQUAL() { return getToken(ViewSQLParser.HIGHER_EQUAL, 0); }
		public TerminalNode LOWER_EQUAL() { return getToken(ViewSQLParser.LOWER_EQUAL, 0); }
		public TerminalNode NOT_EQUAL() { return getToken(ViewSQLParser.NOT_EQUAL, 0); }
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
				if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 516096L) != 0)) ) {
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

	@SuppressWarnings("CheckReturnValue")
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

	@SuppressWarnings("CheckReturnValue")
	public static class SortOrderContext extends ParserRuleContext {
		public TerminalNode DESC() { return getToken(ViewSQLParser.DESC, 0); }
		public TerminalNode ASC() { return getToken(ViewSQLParser.ASC, 0); }
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

	@SuppressWarnings("CheckReturnValue")
	public static class ContinuousRangeContext extends ParserRuleContext {
		public Token left;
		public Token right;
		public List<TerminalNode> LEFT_SP() { return getTokens(ViewSQLParser.LEFT_SP); }
		public TerminalNode LEFT_SP(int i) {
			return getToken(ViewSQLParser.LEFT_SP, i);
		}
		public List<TerminalNode> RIGHT_SP() { return getTokens(ViewSQLParser.RIGHT_SP); }
		public TerminalNode RIGHT_SP(int i) {
			return getToken(ViewSQLParser.RIGHT_SP, i);
		}
		public List<ConstantContext> constant() {
			return getRuleContexts(ConstantContext.class);
		}
		public ConstantContext constant(int i) {
			return getRuleContext(ConstantContext.class,i);
		}
		public TerminalNode RANGE_SEP() { return getToken(ViewSQLParser.RANGE_SEP, 0); }
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

	@SuppressWarnings("CheckReturnValue")
	public static class SparseRangeContext extends ParserRuleContext {
		public TerminalNode LEFT_CURLY() { return getToken(ViewSQLParser.LEFT_CURLY, 0); }
		public List<ConstantContext> constant() {
			return getRuleContexts(ConstantContext.class);
		}
		public ConstantContext constant(int i) {
			return getRuleContext(ConstantContext.class,i);
		}
		public TerminalNode RIGHT_CURLY() { return getToken(ViewSQLParser.RIGHT_CURLY, 0); }
		public List<TerminalNode> SEPARATOR() { return getTokens(ViewSQLParser.SEPARATOR); }
		public TerminalNode SEPARATOR(int i) {
			return getToken(ViewSQLParser.SEPARATOR, i);
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

	@SuppressWarnings("CheckReturnValue")
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

	@SuppressWarnings("CheckReturnValue")
	public static class CreateContext extends ParserRuleContext {
		public TerminalNode CREATE() { return getToken(ViewSQLParser.CREATE, 0); }
		public TerminalNode VIEW() { return getToken(ViewSQLParser.VIEW, 0); }
		public TerminalNode LEFT_P() { return getToken(ViewSQLParser.LEFT_P, 0); }
		public List<TerminalNode> STRING() { return getTokens(ViewSQLParser.STRING); }
		public TerminalNode STRING(int i) {
			return getToken(ViewSQLParser.STRING, i);
		}
		public TerminalNode SEPARATOR() { return getToken(ViewSQLParser.SEPARATOR, 0); }
		public TerminalNode RIGHT_P() { return getToken(ViewSQLParser.RIGHT_P, 0); }
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

	@SuppressWarnings("CheckReturnValue")
	public static class WithContext extends ParserRuleContext {
		public TerminalNode WITH() { return getToken(ViewSQLParser.WITH, 0); }
		public List<TerminalNode> STRING() { return getTokens(ViewSQLParser.STRING); }
		public TerminalNode STRING(int i) {
			return getToken(ViewSQLParser.STRING, i);
		}
		public List<TerminalNode> IN() { return getTokens(ViewSQLParser.IN); }
		public TerminalNode IN(int i) {
			return getToken(ViewSQLParser.IN, i);
		}
		public List<RangeContext> range() {
			return getRuleContexts(RangeContext.class);
		}
		public RangeContext range(int i) {
			return getRuleContext(RangeContext.class,i);
		}
		public List<TerminalNode> SEPARATOR() { return getTokens(ViewSQLParser.SEPARATOR); }
		public TerminalNode SEPARATOR(int i) {
			return getToken(ViewSQLParser.SEPARATOR, i);
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

	@SuppressWarnings("CheckReturnValue")
	public static class SelectContext extends ParserRuleContext {
		public TerminalNode SELECT() { return getToken(ViewSQLParser.SELECT, 0); }
		public List<CalcContext> calc() {
			return getRuleContexts(CalcContext.class);
		}
		public CalcContext calc(int i) {
			return getRuleContext(CalcContext.class,i);
		}
		public List<TerminalNode> SEPARATOR() { return getTokens(ViewSQLParser.SEPARATOR); }
		public TerminalNode SEPARATOR(int i) {
			return getToken(ViewSQLParser.SEPARATOR, i);
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

	@SuppressWarnings("CheckReturnValue")
	public static class FromContext extends ParserRuleContext {
		public TerminalNode FROM() { return getToken(ViewSQLParser.FROM, 0); }
		public List<NameContext> name() {
			return getRuleContexts(NameContext.class);
		}
		public NameContext name(int i) {
			return getRuleContext(NameContext.class,i);
		}
		public List<TerminalNode> SEPARATOR() { return getTokens(ViewSQLParser.SEPARATOR); }
		public TerminalNode SEPARATOR(int i) {
			return getToken(ViewSQLParser.SEPARATOR, i);
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

	@SuppressWarnings("CheckReturnValue")
	public static class WhereContext extends ParserRuleContext {
		public TerminalNode WHERE() { return getToken(ViewSQLParser.WHERE, 0); }
		public List<ConditionContext> condition() {
			return getRuleContexts(ConditionContext.class);
		}
		public ConditionContext condition(int i) {
			return getRuleContext(ConditionContext.class,i);
		}
		public List<TerminalNode> AND() { return getTokens(ViewSQLParser.AND); }
		public TerminalNode AND(int i) {
			return getToken(ViewSQLParser.AND, i);
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

	@SuppressWarnings("CheckReturnValue")
	public static class GroupbyContext extends ParserRuleContext {
		public TerminalNode GROUP() { return getToken(ViewSQLParser.GROUP, 0); }
		public TerminalNode BY() { return getToken(ViewSQLParser.BY, 0); }
		public List<NameableContext> nameable() {
			return getRuleContexts(NameableContext.class);
		}
		public NameableContext nameable(int i) {
			return getRuleContext(NameableContext.class,i);
		}
		public List<TerminalNode> SEPARATOR() { return getTokens(ViewSQLParser.SEPARATOR); }
		public TerminalNode SEPARATOR(int i) {
			return getToken(ViewSQLParser.SEPARATOR, i);
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

	@SuppressWarnings("CheckReturnValue")
	public static class OrderbyContext extends ParserRuleContext {
		public TerminalNode ORDER() { return getToken(ViewSQLParser.ORDER, 0); }
		public TerminalNode BY() { return getToken(ViewSQLParser.BY, 0); }
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
		public List<TerminalNode> SEPARATOR() { return getTokens(ViewSQLParser.SEPARATOR); }
		public TerminalNode SEPARATOR(int i) {
			return getToken(ViewSQLParser.SEPARATOR, i);
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

	@SuppressWarnings("CheckReturnValue")
	public static class LimitContext extends ParserRuleContext {
		public TerminalNode LIMIT() { return getToken(ViewSQLParser.LIMIT, 0); }
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

	@SuppressWarnings("CheckReturnValue")
	public static class ViewContext extends ParserRuleContext {
		public GroupbyContext firstGroupBy;
		public GroupbyContext secondGroupBy;
		public CreateContext create() {
			return getRuleContext(CreateContext.class,0);
		}
		public TerminalNode AS() { return getToken(ViewSQLParser.AS, 0); }
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

	@SuppressWarnings("CheckReturnValue")
	public static class CheckContext extends ParserRuleContext {
		public Token type;
		public TerminalNode CHECK() { return getToken(ViewSQLParser.CHECK, 0); }
		public ConstantContext constant() {
			return getRuleContext(ConstantContext.class,0);
		}
		public TerminalNode HIGHER() { return getToken(ViewSQLParser.HIGHER, 0); }
		public TerminalNode LOWER() { return getToken(ViewSQLParser.LOWER, 0); }
		public TerminalNode HIGHER_EQUAL() { return getToken(ViewSQLParser.HIGHER_EQUAL, 0); }
		public TerminalNode LOWER_EQUAL() { return getToken(ViewSQLParser.LOWER_EQUAL, 0); }
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
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 245760L) != 0)) ) {
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

	@SuppressWarnings("CheckReturnValue")
	public static class ForeignkeyContext extends ParserRuleContext {
		public NameContext tableName;
		public NameContext columnName;
		public TerminalNode FOREIGN() { return getToken(ViewSQLParser.FOREIGN, 0); }
		public TerminalNode KEY() { return getToken(ViewSQLParser.KEY, 0); }
		public TerminalNode REFERENCES() { return getToken(ViewSQLParser.REFERENCES, 0); }
		public TerminalNode LEFT_P() { return getToken(ViewSQLParser.LEFT_P, 0); }
		public TerminalNode RIGHT_P() { return getToken(ViewSQLParser.RIGHT_P, 0); }
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

	@SuppressWarnings("CheckReturnValue")
	public static class PrimarykeyContext extends ParserRuleContext {
		public TerminalNode PRIMARY() { return getToken(ViewSQLParser.PRIMARY, 0); }
		public TerminalNode KEY() { return getToken(ViewSQLParser.KEY, 0); }
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

	@SuppressWarnings("CheckReturnValue")
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

	@SuppressWarnings("CheckReturnValue")
	public static class ColumnsContext extends ParserRuleContext {
		public Token type;
		public Token policy;
		public TerminalNode STRING() { return getToken(ViewSQLParser.STRING, 0); }
		public TerminalNode COUNTER() { return getToken(ViewSQLParser.COUNTER, 0); }
		public TerminalNode INTEGER() { return getToken(ViewSQLParser.INTEGER, 0); }
		public TerminalNode BOOLEAN() { return getToken(ViewSQLParser.BOOLEAN, 0); }
		public TerminalNode VARCHAR() { return getToken(ViewSQLParser.VARCHAR, 0); }
		public TerminalNode DATE_TYPE() { return getToken(ViewSQLParser.DATE_TYPE, 0); }
		public TerminalNode LWW() { return getToken(ViewSQLParser.LWW, 0); }
		public TerminalNode MW() { return getToken(ViewSQLParser.MW, 0); }
		public TerminalNode EW() { return getToken(ViewSQLParser.EW, 0); }
		public TerminalNode DW() { return getToken(ViewSQLParser.DW, 0); }
		public TerminalNode DEFAULT() { return getToken(ViewSQLParser.DEFAULT, 0); }
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
			if ( !(((((_la - 63)) & ~0x3f) == 0 && ((1L << (_la - 63)) & 31L) != 0)) ) {
				((ColumnsContext)_localctx).type = (Token)_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(327);
			((ColumnsContext)_localctx).policy = _input.LT(1);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 4323455642275676160L) != 0)) ) {
				((ColumnsContext)_localctx).policy = (Token)_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(330);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==DEFAULT) {
				{
				setState(328);
				match(DEFAULT);
				setState(329);
				constant();
				}
			}

			setState(333);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 20268397346422784L) != 0)) {
				{
				setState(332);
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

	@SuppressWarnings("CheckReturnValue")
	public static class CreatetableContext extends ParserRuleContext {
		public Token policy;
		public TerminalNode CREATE() { return getToken(ViewSQLParser.CREATE, 0); }
		public TerminalNode TABLE() { return getToken(ViewSQLParser.TABLE, 0); }
		public NameContext name() {
			return getRuleContext(NameContext.class,0);
		}
		public TerminalNode LEFT_P() { return getToken(ViewSQLParser.LEFT_P, 0); }
		public List<ColumnsContext> columns() {
			return getRuleContexts(ColumnsContext.class);
		}
		public ColumnsContext columns(int i) {
			return getRuleContext(ColumnsContext.class,i);
		}
		public TerminalNode RIGHT_P() { return getToken(ViewSQLParser.RIGHT_P, 0); }
		public TerminalNode AW() { return getToken(ViewSQLParser.AW, 0); }
		public TerminalNode RW() { return getToken(ViewSQLParser.RW, 0); }
		public TerminalNode LWW() { return getToken(ViewSQLParser.LWW, 0); }
		public List<TerminalNode> SEPARATOR() { return getTokens(ViewSQLParser.SEPARATOR); }
		public TerminalNode SEPARATOR(int i) {
			return getToken(ViewSQLParser.SEPARATOR, i);
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
			setState(335);
			match(CREATE);
			setState(336);
			((CreatetableContext)_localctx).policy = _input.LT(1);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 504403158265495552L) != 0)) ) {
				((CreatetableContext)_localctx).policy = (Token)_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(337);
			match(TABLE);
			setState(338);
			name();
			setState(339);
			match(LEFT_P);
			setState(345);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,26,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(340);
					columns();
					setState(341);
					match(SEPARATOR);
					}
					} 
				}
				setState(347);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,26,_ctx);
			}
			setState(348);
			columns();
			setState(349);
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

	@SuppressWarnings("CheckReturnValue")
	public static class CreateindexContext extends ParserRuleContext {
		public NameContext indexName;
		public NameContext tableName;
		public NameContext columnName;
		public TerminalNode CREATE() { return getToken(ViewSQLParser.CREATE, 0); }
		public TerminalNode INDEX() { return getToken(ViewSQLParser.INDEX, 0); }
		public TerminalNode ON() { return getToken(ViewSQLParser.ON, 0); }
		public TerminalNode LEFT_P() { return getToken(ViewSQLParser.LEFT_P, 0); }
		public TerminalNode RIGHT_P() { return getToken(ViewSQLParser.RIGHT_P, 0); }
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
			setState(351);
			match(CREATE);
			setState(352);
			match(INDEX);
			setState(353);
			((CreateindexContext)_localctx).indexName = name();
			setState(354);
			match(ON);
			setState(355);
			((CreateindexContext)_localctx).tableName = name();
			setState(356);
			match(LEFT_P);
			setState(357);
			((CreateindexContext)_localctx).columnName = name();
			setState(358);
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

	@SuppressWarnings("CheckReturnValue")
	public static class DropContext extends ParserRuleContext {
		public TerminalNode DROP() { return getToken(ViewSQLParser.DROP, 0); }
		public TerminalNode TABLE() { return getToken(ViewSQLParser.TABLE, 0); }
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
			setState(360);
			match(DROP);
			setState(361);
			match(TABLE);
			setState(362);
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

	@SuppressWarnings("CheckReturnValue")
	public static class DeleteContext extends ParserRuleContext {
		public TerminalNode DELETE() { return getToken(ViewSQLParser.DELETE, 0); }
		public TerminalNode FROM() { return getToken(ViewSQLParser.FROM, 0); }
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
			setState(364);
			match(DELETE);
			setState(365);
			match(FROM);
			setState(366);
			name();
			setState(367);
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

	@SuppressWarnings("CheckReturnValue")
	public static class SetContext extends ParserRuleContext {
		public TerminalNode SET() { return getToken(ViewSQLParser.SET, 0); }
		public List<NameContext> name() {
			return getRuleContexts(NameContext.class);
		}
		public NameContext name(int i) {
			return getRuleContext(NameContext.class,i);
		}
		public List<TerminalNode> EQUAL() { return getTokens(ViewSQLParser.EQUAL); }
		public TerminalNode EQUAL(int i) {
			return getToken(ViewSQLParser.EQUAL, i);
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
		public List<TerminalNode> SEPARATOR() { return getTokens(ViewSQLParser.SEPARATOR); }
		public TerminalNode SEPARATOR(int i) {
			return getToken(ViewSQLParser.SEPARATOR, i);
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
			setState(369);
			match(SET);
			setState(377);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,27,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(370);
					name();
					setState(371);
					match(EQUAL);
					setState(372);
					constant();
					setState(373);
					match(SEPARATOR);
					}
					} 
				}
				setState(379);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,27,_ctx);
			}
			setState(380);
			name();
			setState(381);
			match(EQUAL);
			setState(382);
			constant();
			setState(383);
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

	@SuppressWarnings("CheckReturnValue")
	public static class UpdateContext extends ParserRuleContext {
		public TerminalNode UPDATE() { return getToken(ViewSQLParser.UPDATE, 0); }
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
			setState(385);
			match(UPDATE);
			setState(386);
			name();
			setState(387);
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

	@SuppressWarnings("CheckReturnValue")
	public static class ValuesContext extends ParserRuleContext {
		public TerminalNode VALUES() { return getToken(ViewSQLParser.VALUES, 0); }
		public TerminalNode LEFT_P() { return getToken(ViewSQLParser.LEFT_P, 0); }
		public List<ConstantContext> constant() {
			return getRuleContexts(ConstantContext.class);
		}
		public ConstantContext constant(int i) {
			return getRuleContext(ConstantContext.class,i);
		}
		public TerminalNode RIGHT_P() { return getToken(ViewSQLParser.RIGHT_P, 0); }
		public List<TerminalNode> SEPARATOR() { return getTokens(ViewSQLParser.SEPARATOR); }
		public TerminalNode SEPARATOR(int i) {
			return getToken(ViewSQLParser.SEPARATOR, i);
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
			setState(389);
			match(VALUES);
			setState(390);
			match(LEFT_P);
			setState(396);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,28,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(391);
					constant();
					setState(392);
					match(SEPARATOR);
					}
					} 
				}
				setState(398);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,28,_ctx);
			}
			setState(399);
			constant();
			setState(400);
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

	@SuppressWarnings("CheckReturnValue")
	public static class ColumnNamesContext extends ParserRuleContext {
		public TerminalNode LEFT_P() { return getToken(ViewSQLParser.LEFT_P, 0); }
		public List<NameContext> name() {
			return getRuleContexts(NameContext.class);
		}
		public NameContext name(int i) {
			return getRuleContext(NameContext.class,i);
		}
		public TerminalNode RIGHT_P() { return getToken(ViewSQLParser.RIGHT_P, 0); }
		public List<TerminalNode> SEPARATOR() { return getTokens(ViewSQLParser.SEPARATOR); }
		public TerminalNode SEPARATOR(int i) {
			return getToken(ViewSQLParser.SEPARATOR, i);
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
			setState(402);
			match(LEFT_P);
			setState(408);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,29,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(403);
					name();
					setState(404);
					match(SEPARATOR);
					}
					} 
				}
				setState(410);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,29,_ctx);
			}
			setState(411);
			name();
			setState(412);
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

	@SuppressWarnings("CheckReturnValue")
	public static class InsertContext extends ParserRuleContext {
		public TerminalNode INSERT() { return getToken(ViewSQLParser.INSERT, 0); }
		public TerminalNode INTO() { return getToken(ViewSQLParser.INTO, 0); }
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
			setState(414);
			match(INSERT);
			setState(415);
			match(INTO);
			setState(416);
			name();
			setState(418);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==LEFT_P) {
				{
				setState(417);
				columnNames();
				}
			}

			setState(420);
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

	@SuppressWarnings("CheckReturnValue")
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
			setState(422);
			select();
			setState(423);
			from();
			setState(425);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==WHERE) {
				{
				setState(424);
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

	@SuppressWarnings("CheckReturnValue")
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
			setState(435);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,32,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(427);
				view();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(428);
				createtable();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(429);
				createindex();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(430);
				insert();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(431);
				update();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(432);
				delete();
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(433);
				drop();
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(434);
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

	@SuppressWarnings("CheckReturnValue")
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
			setState(437);
			statement();
			setState(438);
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
		"\u0004\u0001I\u01b9\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b\u0007\u000b\u0002"+
		"\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e\u0002\u000f\u0007\u000f"+
		"\u0002\u0010\u0007\u0010\u0002\u0011\u0007\u0011\u0002\u0012\u0007\u0012"+
		"\u0002\u0013\u0007\u0013\u0002\u0014\u0007\u0014\u0002\u0015\u0007\u0015"+
		"\u0002\u0016\u0007\u0016\u0002\u0017\u0007\u0017\u0002\u0018\u0007\u0018"+
		"\u0002\u0019\u0007\u0019\u0002\u001a\u0007\u001a\u0002\u001b\u0007\u001b"+
		"\u0002\u001c\u0007\u001c\u0002\u001d\u0007\u001d\u0002\u001e\u0007\u001e"+
		"\u0002\u001f\u0007\u001f\u0002 \u0007 \u0002!\u0007!\u0002\"\u0007\"\u0002"+
		"#\u0007#\u0002$\u0007$\u0002%\u0007%\u0002&\u0007&\u0002\'\u0007\'\u0002"+
		"(\u0007(\u0002)\u0007)\u0002*\u0007*\u0002+\u0007+\u0001\u0000\u0001\u0000"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0005\u0001b\b\u0001\n\u0001\f\u0001e\t\u0001\u0001\u0001"+
		"\u0003\u0001h\b\u0001\u0001\u0002\u0001\u0002\u0001\u0003\u0001\u0003"+
		"\u0001\u0003\u0001\u0003\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004"+
		"\u0001\u0005\u0001\u0005\u0001\u0005\u0003\u0005w\b\u0005\u0001\u0006"+
		"\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0007"+
		"\u0001\u0007\u0001\u0007\u0003\u0007\u0082\b\u0007\u0001\u0007\u0001\u0007"+
		"\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0003\u0007\u008a\b\u0007"+
		"\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007"+
		"\u0005\u0007\u0092\b\u0007\n\u0007\f\u0007\u0095\t\u0007\u0001\b\u0001"+
		"\b\u0001\b\u0001\b\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001"+
		"\t\u0001\n\u0001\n\u0001\n\u0001\n\u0003\n\u00a6\b\n\u0001\n\u0001\n\u0001"+
		"\n\u0001\n\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b"+
		"\u0001\u000b\u0003\u000b\u00b2\b\u000b\u0001\f\u0001\f\u0003\f\u00b6\b"+
		"\f\u0001\r\u0001\r\u0001\r\u0001\r\u0001\u000e\u0001\u000e\u0001\u000f"+
		"\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f"+
		"\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0005\u0010\u00c9\b\u0010"+
		"\n\u0010\f\u0010\u00cc\t\u0010\u0001\u0010\u0001\u0010\u0001\u0011\u0001"+
		"\u0011\u0003\u0011\u00d2\b\u0011\u0001\u0012\u0001\u0012\u0001\u0012\u0001"+
		"\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0013\u0001"+
		"\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001"+
		"\u0013\u0005\u0013\u00e4\b\u0013\n\u0013\f\u0013\u00e7\t\u0013\u0001\u0014"+
		"\u0001\u0014\u0001\u0014\u0001\u0014\u0005\u0014\u00ed\b\u0014\n\u0014"+
		"\f\u0014\u00f0\t\u0014\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015"+
		"\u0005\u0015\u00f6\b\u0015\n\u0015\f\u0015\u00f9\t\u0015\u0001\u0016\u0001"+
		"\u0016\u0001\u0016\u0001\u0016\u0005\u0016\u00ff\b\u0016\n\u0016\f\u0016"+
		"\u0102\t\u0016\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017"+
		"\u0005\u0017\u0109\b\u0017\n\u0017\f\u0017\u010c\t\u0017\u0001\u0018\u0001"+
		"\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001"+
		"\u0018\u0005\u0018\u0116\b\u0018\n\u0018\f\u0018\u0119\t\u0018\u0001\u0019"+
		"\u0001\u0019\u0001\u0019\u0001\u001a\u0001\u001a\u0003\u001a\u0120\b\u001a"+
		"\u0001\u001a\u0001\u001a\u0001\u001a\u0001\u001a\u0003\u001a\u0126\b\u001a"+
		"\u0001\u001a\u0003\u001a\u0129\b\u001a\u0001\u001a\u0001\u001a\u0003\u001a"+
		"\u012d\b\u001a\u0001\u001a\u0003\u001a\u0130\b\u001a\u0001\u001b\u0001"+
		"\u001b\u0001\u001b\u0001\u001b\u0001\u001c\u0001\u001c\u0001\u001c\u0001"+
		"\u001c\u0001\u001c\u0001\u001c\u0001\u001c\u0001\u001c\u0001\u001d\u0001"+
		"\u001d\u0001\u001d\u0001\u001e\u0001\u001e\u0001\u001e\u0003\u001e\u0144"+
		"\b\u001e\u0001\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0003"+
		"\u001f\u014b\b\u001f\u0001\u001f\u0003\u001f\u014e\b\u001f\u0001 \u0001"+
		" \u0001 \u0001 \u0001 \u0001 \u0001 \u0001 \u0005 \u0158\b \n \f \u015b"+
		"\t \u0001 \u0001 \u0001 \u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001"+
		"!\u0001!\u0001!\u0001\"\u0001\"\u0001\"\u0001\"\u0001#\u0001#\u0001#\u0001"+
		"#\u0001#\u0001$\u0001$\u0001$\u0001$\u0001$\u0001$\u0005$\u0178\b$\n$"+
		"\f$\u017b\t$\u0001$\u0001$\u0001$\u0001$\u0001$\u0001%\u0001%\u0001%\u0001"+
		"%\u0001&\u0001&\u0001&\u0001&\u0001&\u0005&\u018b\b&\n&\f&\u018e\t&\u0001"+
		"&\u0001&\u0001&\u0001\'\u0001\'\u0001\'\u0001\'\u0005\'\u0197\b\'\n\'"+
		"\f\'\u019a\t\'\u0001\'\u0001\'\u0001\'\u0001(\u0001(\u0001(\u0001(\u0003"+
		"(\u01a3\b(\u0001(\u0001(\u0001)\u0001)\u0001)\u0003)\u01aa\b)\u0001*\u0001"+
		"*\u0001*\u0001*\u0001*\u0001*\u0001*\u0001*\u0003*\u01b4\b*\u0001+\u0001"+
		"+\u0001+\u0001+\u0000\u0001\u000e,\u0000\u0002\u0004\u0006\b\n\f\u000e"+
		"\u0010\u0012\u0014\u0016\u0018\u001a\u001c\u001e \"$&(*,.02468:<>@BDF"+
		"HJLNPRTV\u0000\u000b\u0002\u0000FFII\u0001\u0000\u0005\b\u0001\u0000\u0003"+
		"\u0004\u0001\u0000\u0001\u0002\u0001\u0000\r\u0012\u0001\u0000\u000b\f"+
		"\u0001\u0000\u0018\u0019\u0001\u0000\u000e\u0011\u0001\u0000?C\u0001\u0000"+
		":=\u0001\u00008:\u01bd\u0000X\u0001\u0000\u0000\u0000\u0002g\u0001\u0000"+
		"\u0000\u0000\u0004i\u0001\u0000\u0000\u0000\u0006k\u0001\u0000\u0000\u0000"+
		"\bo\u0001\u0000\u0000\u0000\nv\u0001\u0000\u0000\u0000\fx\u0001\u0000"+
		"\u0000\u0000\u000e\u0089\u0001\u0000\u0000\u0000\u0010\u0096\u0001\u0000"+
		"\u0000\u0000\u0012\u009a\u0001\u0000\u0000\u0000\u0014\u00a1\u0001\u0000"+
		"\u0000\u0000\u0016\u00b1\u0001\u0000\u0000\u0000\u0018\u00b5\u0001\u0000"+
		"\u0000\u0000\u001a\u00b7\u0001\u0000\u0000\u0000\u001c\u00bb\u0001\u0000"+
		"\u0000\u0000\u001e\u00bd\u0001\u0000\u0000\u0000 \u00c4\u0001\u0000\u0000"+
		"\u0000\"\u00d1\u0001\u0000\u0000\u0000$\u00d3\u0001\u0000\u0000\u0000"+
		"&\u00db\u0001\u0000\u0000\u0000(\u00e8\u0001\u0000\u0000\u0000*\u00f1"+
		"\u0001\u0000\u0000\u0000,\u00fa\u0001\u0000\u0000\u0000.\u0103\u0001\u0000"+
		"\u0000\u00000\u010d\u0001\u0000\u0000\u00002\u011a\u0001\u0000\u0000\u0000"+
		"4\u011d\u0001\u0000\u0000\u00006\u0131\u0001\u0000\u0000\u00008\u0135"+
		"\u0001\u0000\u0000\u0000:\u013d\u0001\u0000\u0000\u0000<\u0143\u0001\u0000"+
		"\u0000\u0000>\u0145\u0001\u0000\u0000\u0000@\u014f\u0001\u0000\u0000\u0000"+
		"B\u015f\u0001\u0000\u0000\u0000D\u0168\u0001\u0000\u0000\u0000F\u016c"+
		"\u0001\u0000\u0000\u0000H\u0171\u0001\u0000\u0000\u0000J\u0181\u0001\u0000"+
		"\u0000\u0000L\u0185\u0001\u0000\u0000\u0000N\u0192\u0001\u0000\u0000\u0000"+
		"P\u019e\u0001\u0000\u0000\u0000R\u01a6\u0001\u0000\u0000\u0000T\u01b3"+
		"\u0001\u0000\u0000\u0000V\u01b5\u0001\u0000\u0000\u0000XY\u0005F\u0000"+
		"\u0000Y\u0001\u0001\u0000\u0000\u0000Zh\u0005D\u0000\u0000[h\u0005E\u0000"+
		"\u0000\\h\u0005G\u0000\u0000]h\u0005H\u0000\u0000^_\u0005\u001a\u0000"+
		"\u0000_c\u0005F\u0000\u0000`b\u0007\u0000\u0000\u0000a`\u0001\u0000\u0000"+
		"\u0000be\u0001\u0000\u0000\u0000ca\u0001\u0000\u0000\u0000cd\u0001\u0000"+
		"\u0000\u0000df\u0001\u0000\u0000\u0000ec\u0001\u0000\u0000\u0000fh\u0005"+
		"\u001a\u0000\u0000gZ\u0001\u0000\u0000\u0000g[\u0001\u0000\u0000\u0000"+
		"g\\\u0001\u0000\u0000\u0000g]\u0001\u0000\u0000\u0000g^\u0001\u0000\u0000"+
		"\u0000h\u0003\u0001\u0000\u0000\u0000ij\u0007\u0001\u0000\u0000j\u0005"+
		"\u0001\u0000\u0000\u0000kl\u0003\u0000\u0000\u0000lm\u0005\u0013\u0000"+
		"\u0000mn\u0003\u0000\u0000\u0000n\u0007\u0001\u0000\u0000\u0000op\u0005"+
		"\u0018\u0000\u0000pq\u0003\u0000\u0000\u0000qr\u0005\u0019\u0000\u0000"+
		"r\t\u0001\u0000\u0000\u0000sw\u0003\u0000\u0000\u0000tw\u0003\u0006\u0003"+
		"\u0000uw\u0003\b\u0004\u0000vs\u0001\u0000\u0000\u0000vt\u0001\u0000\u0000"+
		"\u0000vu\u0001\u0000\u0000\u0000w\u000b\u0001\u0000\u0000\u0000xy\u0005"+
		")\u0000\u0000yz\u0005*\u0000\u0000z{\u0005\u0016\u0000\u0000{|\u0003\n"+
		"\u0005\u0000|}\u0005\u0017\u0000\u0000}\r\u0001\u0000\u0000\u0000~\u0081"+
		"\u0006\u0007\uffff\uffff\u0000\u007f\u0082\u0003\n\u0005\u0000\u0080\u0082"+
		"\u0003\u0002\u0001\u0000\u0081\u007f\u0001\u0000\u0000\u0000\u0081\u0080"+
		"\u0001\u0000\u0000\u0000\u0082\u008a\u0001\u0000\u0000\u0000\u0083\u0084"+
		"\u0005\u0002\u0000\u0000\u0084\u008a\u0003\u000e\u0007\u0004\u0085\u0086"+
		"\u0005\u0016\u0000\u0000\u0086\u0087\u0003\u000e\u0007\u0000\u0087\u0088"+
		"\u0005\u0017\u0000\u0000\u0088\u008a\u0001\u0000\u0000\u0000\u0089~\u0001"+
		"\u0000\u0000\u0000\u0089\u0083\u0001\u0000\u0000\u0000\u0089\u0085\u0001"+
		"\u0000\u0000\u0000\u008a\u0093\u0001\u0000\u0000\u0000\u008b\u008c\n\u0002"+
		"\u0000\u0000\u008c\u008d\u0007\u0002\u0000\u0000\u008d\u0092\u0003\u000e"+
		"\u0007\u0003\u008e\u008f\n\u0001\u0000\u0000\u008f\u0090\u0007\u0003\u0000"+
		"\u0000\u0090\u0092\u0003\u000e\u0007\u0002\u0091\u008b\u0001\u0000\u0000"+
		"\u0000\u0091\u008e\u0001\u0000\u0000\u0000\u0092\u0095\u0001\u0000\u0000"+
		"\u0000\u0093\u0091\u0001\u0000\u0000\u0000\u0093\u0094\u0001\u0000\u0000"+
		"\u0000\u0094\u000f\u0001\u0000\u0000\u0000\u0095\u0093\u0001\u0000\u0000"+
		"\u0000\u0096\u0097\u0003\u000e\u0007\u0000\u0097\u0098\u0005\u001d\u0000"+
		"\u0000\u0098\u0099\u0003\u0000\u0000\u0000\u0099\u0011\u0001\u0000\u0000"+
		"\u0000\u009a\u009b\u0003\u0004\u0002\u0000\u009b\u009c\u0005\u0016\u0000"+
		"\u0000\u009c\u009d\u0003\u000e\u0007\u0000\u009d\u009e\u0005\u0017\u0000"+
		"\u0000\u009e\u009f\u0005\u001d\u0000\u0000\u009f\u00a0\u0003\u0000\u0000"+
		"\u0000\u00a0\u0013\u0001\u0000\u0000\u0000\u00a1\u00a2\u0005\t\u0000\u0000"+
		"\u00a2\u00a5\u0005\u0016\u0000\u0000\u00a3\u00a6\u0003\n\u0005\u0000\u00a4"+
		"\u00a6\u0005\u0003\u0000\u0000\u00a5\u00a3\u0001\u0000\u0000\u0000\u00a5"+
		"\u00a4\u0001\u0000\u0000\u0000\u00a6\u00a7\u0001\u0000\u0000\u0000\u00a7"+
		"\u00a8\u0005\u0017\u0000\u0000\u00a8\u00a9\u0005\u001d\u0000\u0000\u00a9"+
		"\u00aa\u0003\u0000\u0000\u0000\u00aa\u0015\u0001\u0000\u0000\u0000\u00ab"+
		"\u00b2\u0003\f\u0006\u0000\u00ac\u00b2\u0003\n\u0005\u0000\u00ad\u00b2"+
		"\u0003\u0010\b\u0000\u00ae\u00b2\u0003\u0012\t\u0000\u00af\u00b2\u0003"+
		"\u0014\n\u0000\u00b0\u00b2\u0005\u0003\u0000\u0000\u00b1\u00ab\u0001\u0000"+
		"\u0000\u0000\u00b1\u00ac\u0001\u0000\u0000\u0000\u00b1\u00ad\u0001\u0000"+
		"\u0000\u0000\u00b1\u00ae\u0001\u0000\u0000\u0000\u00b1\u00af\u0001\u0000"+
		"\u0000\u0000\u00b1\u00b0\u0001\u0000\u0000\u0000\u00b2\u0017\u0001\u0000"+
		"\u0000\u0000\u00b3\u00b6\u0007\u0004\u0000\u0000\u00b4\u00b6\u0003\u0000"+
		"\u0000\u0000\u00b5\u00b3\u0001\u0000\u0000\u0000\u00b5\u00b4\u0001\u0000"+
		"\u0000\u0000\u00b6\u0019\u0001\u0000\u0000\u0000\u00b7\u00b8\u0003\n\u0005"+
		"\u0000\u00b8\u00b9\u0003\u0018\f\u0000\u00b9\u00ba\u0003\u000e\u0007\u0000"+
		"\u00ba\u001b\u0001\u0000\u0000\u0000\u00bb\u00bc\u0007\u0005\u0000\u0000"+
		"\u00bc\u001d\u0001\u0000\u0000\u0000\u00bd\u00be\u0007\u0006\u0000\u0000"+
		"\u00be\u00bf\u0003\u0002\u0001\u0000\u00bf\u00c0\u0005\u0015\u0000\u0000"+
		"\u00c0\u00c1\u0003\u0002\u0001\u0000\u00c1\u00c2\u0001\u0000\u0000\u0000"+
		"\u00c2\u00c3\u0007\u0006\u0000\u0000\u00c3\u001f\u0001\u0000\u0000\u0000"+
		"\u00c4\u00c5\u0005\u001b\u0000\u0000\u00c5\u00ca\u0003\u0002\u0001\u0000"+
		"\u00c6\u00c7\u0005\u0014\u0000\u0000\u00c7\u00c9\u0003\u0002\u0001\u0000"+
		"\u00c8\u00c6\u0001\u0000\u0000\u0000\u00c9\u00cc\u0001\u0000\u0000\u0000"+
		"\u00ca\u00c8\u0001\u0000\u0000\u0000\u00ca\u00cb\u0001\u0000\u0000\u0000"+
		"\u00cb\u00cd\u0001\u0000\u0000\u0000\u00cc\u00ca\u0001\u0000\u0000\u0000"+
		"\u00cd\u00ce\u0005\u001c\u0000\u0000\u00ce!\u0001\u0000\u0000\u0000\u00cf"+
		"\u00d2\u0003\u001e\u000f\u0000\u00d0\u00d2\u0003 \u0010\u0000\u00d1\u00cf"+
		"\u0001\u0000\u0000\u0000\u00d1\u00d0\u0001\u0000\u0000\u0000\u00d2#\u0001"+
		"\u0000\u0000\u0000\u00d3\u00d4\u0005\u001e\u0000\u0000\u00d4\u00d5\u0005"+
		"\u001f\u0000\u0000\u00d5\u00d6\u0005\u0016\u0000\u0000\u00d6\u00d7\u0005"+
		"F\u0000\u0000\u00d7\u00d8\u0005\u0014\u0000\u0000\u00d8\u00d9\u0005F\u0000"+
		"\u0000\u00d9\u00da\u0005\u0017\u0000\u0000\u00da%\u0001\u0000\u0000\u0000"+
		"\u00db\u00dc\u0005\'\u0000\u0000\u00dc\u00dd\u0005F\u0000\u0000\u00dd"+
		"\u00de\u0005(\u0000\u0000\u00de\u00e5\u0003\"\u0011\u0000\u00df\u00e0"+
		"\u0005\u0014\u0000\u0000\u00e0\u00e1\u0005F\u0000\u0000\u00e1\u00e2\u0005"+
		"(\u0000\u0000\u00e2\u00e4\u0003\"\u0011\u0000\u00e3\u00df\u0001\u0000"+
		"\u0000\u0000\u00e4\u00e7\u0001\u0000\u0000\u0000\u00e5\u00e3\u0001\u0000"+
		"\u0000\u0000\u00e5\u00e6\u0001\u0000\u0000\u0000\u00e6\'\u0001\u0000\u0000"+
		"\u0000\u00e7\u00e5\u0001\u0000\u0000\u0000\u00e8\u00e9\u0005 \u0000\u0000"+
		"\u00e9\u00ee\u0003\u0016\u000b\u0000\u00ea\u00eb\u0005\u0014\u0000\u0000"+
		"\u00eb\u00ed\u0003\u0016\u000b\u0000\u00ec\u00ea\u0001\u0000\u0000\u0000"+
		"\u00ed\u00f0\u0001\u0000\u0000\u0000\u00ee\u00ec\u0001\u0000\u0000\u0000"+
		"\u00ee\u00ef\u0001\u0000\u0000\u0000\u00ef)\u0001\u0000\u0000\u0000\u00f0"+
		"\u00ee\u0001\u0000\u0000\u0000\u00f1\u00f2\u0005!\u0000\u0000\u00f2\u00f7"+
		"\u0003\u0000\u0000\u0000\u00f3\u00f4\u0005\u0014\u0000\u0000\u00f4\u00f6"+
		"\u0003\u0000\u0000\u0000\u00f5\u00f3\u0001\u0000\u0000\u0000\u00f6\u00f9"+
		"\u0001\u0000\u0000\u0000\u00f7\u00f5\u0001\u0000\u0000\u0000\u00f7\u00f8"+
		"\u0001\u0000\u0000\u0000\u00f8+\u0001\u0000\u0000\u0000\u00f9\u00f7\u0001"+
		"\u0000\u0000\u0000\u00fa\u00fb\u0005\"\u0000\u0000\u00fb\u0100\u0003\u001a"+
		"\r\u0000\u00fc\u00fd\u0005\n\u0000\u0000\u00fd\u00ff\u0003\u001a\r\u0000"+
		"\u00fe\u00fc\u0001\u0000\u0000\u0000\u00ff\u0102\u0001\u0000\u0000\u0000"+
		"\u0100\u00fe\u0001\u0000\u0000\u0000\u0100\u0101\u0001\u0000\u0000\u0000"+
		"\u0101-\u0001\u0000\u0000\u0000\u0102\u0100\u0001\u0000\u0000\u0000\u0103"+
		"\u0104\u0005#\u0000\u0000\u0104\u0105\u0005$\u0000\u0000\u0105\u010a\u0003"+
		"\n\u0005\u0000\u0106\u0107\u0005\u0014\u0000\u0000\u0107\u0109\u0003\n"+
		"\u0005\u0000\u0108\u0106\u0001\u0000\u0000\u0000\u0109\u010c\u0001\u0000"+
		"\u0000\u0000\u010a\u0108\u0001\u0000\u0000\u0000\u010a\u010b\u0001\u0000"+
		"\u0000\u0000\u010b/\u0001\u0000\u0000\u0000\u010c\u010a\u0001\u0000\u0000"+
		"\u0000\u010d\u010e\u0005%\u0000\u0000\u010e\u010f\u0005$\u0000\u0000\u010f"+
		"\u0110\u0003\n\u0005\u0000\u0110\u0117\u0003\u001c\u000e\u0000\u0111\u0112"+
		"\u0005\u0014\u0000\u0000\u0112\u0113\u0003\n\u0005\u0000\u0113\u0114\u0003"+
		"\u001c\u000e\u0000\u0114\u0116\u0001\u0000\u0000\u0000\u0115\u0111\u0001"+
		"\u0000\u0000\u0000\u0116\u0119\u0001\u0000\u0000\u0000\u0117\u0115\u0001"+
		"\u0000\u0000\u0000\u0117\u0118\u0001\u0000\u0000\u0000\u01181\u0001\u0000"+
		"\u0000\u0000\u0119\u0117\u0001\u0000\u0000\u0000\u011a\u011b\u0005&\u0000"+
		"\u0000\u011b\u011c\u0005G\u0000\u0000\u011c3\u0001\u0000\u0000\u0000\u011d"+
		"\u011f\u0003$\u0012\u0000\u011e\u0120\u0003&\u0013\u0000\u011f\u011e\u0001"+
		"\u0000\u0000\u0000\u011f\u0120\u0001\u0000\u0000\u0000\u0120\u0121\u0001"+
		"\u0000\u0000\u0000\u0121\u0122\u0005\u001d\u0000\u0000\u0122\u0123\u0003"+
		"(\u0014\u0000\u0123\u0125\u0003*\u0015\u0000\u0124\u0126\u0003,\u0016"+
		"\u0000\u0125\u0124\u0001\u0000\u0000\u0000\u0125\u0126\u0001\u0000\u0000"+
		"\u0000\u0126\u0128\u0001\u0000\u0000\u0000\u0127\u0129\u0003.\u0017\u0000"+
		"\u0128\u0127\u0001\u0000\u0000\u0000\u0128\u0129\u0001\u0000\u0000\u0000"+
		"\u0129\u012a\u0001\u0000\u0000\u0000\u012a\u012c\u0003.\u0017\u0000\u012b"+
		"\u012d\u00030\u0018\u0000\u012c\u012b\u0001\u0000\u0000\u0000\u012c\u012d"+
		"\u0001\u0000\u0000\u0000\u012d\u012f\u0001\u0000\u0000\u0000\u012e\u0130"+
		"\u00032\u0019\u0000\u012f\u012e\u0001\u0000\u0000\u0000\u012f\u0130\u0001"+
		"\u0000\u0000\u0000\u01305\u0001\u0000\u0000\u0000\u0131\u0132\u00056\u0000"+
		"\u0000\u0132\u0133\u0007\u0007\u0000\u0000\u0133\u0134\u0003\u0002\u0001"+
		"\u0000\u01347\u0001\u0000\u0000\u0000\u0135\u0136\u00053\u0000\u0000\u0136"+
		"\u0137\u0005*\u0000\u0000\u0137\u0138\u00054\u0000\u0000\u0138\u0139\u0003"+
		"\u0000\u0000\u0000\u0139\u013a\u0005\u0016\u0000\u0000\u013a\u013b\u0003"+
		"\u0000\u0000\u0000\u013b\u013c\u0005\u0017\u0000\u0000\u013c9\u0001\u0000"+
		"\u0000\u0000\u013d\u013e\u0005)\u0000\u0000\u013e\u013f\u0005*\u0000\u0000"+
		"\u013f;\u0001\u0000\u0000\u0000\u0140\u0144\u0003:\u001d\u0000\u0141\u0144"+
		"\u00038\u001c\u0000\u0142\u0144\u00036\u001b\u0000\u0143\u0140\u0001\u0000"+
		"\u0000\u0000\u0143\u0141\u0001\u0000\u0000\u0000\u0143\u0142\u0001\u0000"+
		"\u0000\u0000\u0144=\u0001\u0000\u0000\u0000\u0145\u0146\u0005F\u0000\u0000"+
		"\u0146\u0147\u0007\b\u0000\u0000\u0147\u014a\u0007\t\u0000\u0000\u0148"+
		"\u0149\u0005>\u0000\u0000\u0149\u014b\u0003\u0002\u0001\u0000\u014a\u0148"+
		"\u0001\u0000\u0000\u0000\u014a\u014b\u0001\u0000\u0000\u0000\u014b\u014d"+
		"\u0001\u0000\u0000\u0000\u014c\u014e\u0003<\u001e\u0000\u014d\u014c\u0001"+
		"\u0000\u0000\u0000\u014d\u014e\u0001\u0000\u0000\u0000\u014e?\u0001\u0000"+
		"\u0000\u0000\u014f\u0150\u0005\u001e\u0000\u0000\u0150\u0151\u0007\n\u0000"+
		"\u0000\u0151\u0152\u0005+\u0000\u0000\u0152\u0153\u0003\u0000\u0000\u0000"+
		"\u0153\u0159\u0005\u0016\u0000\u0000\u0154\u0155\u0003>\u001f\u0000\u0155"+
		"\u0156\u0005\u0014\u0000\u0000\u0156\u0158\u0001\u0000\u0000\u0000\u0157"+
		"\u0154\u0001\u0000\u0000\u0000\u0158\u015b\u0001\u0000\u0000\u0000\u0159"+
		"\u0157\u0001\u0000\u0000\u0000\u0159\u015a\u0001\u0000\u0000\u0000\u015a"+
		"\u015c\u0001\u0000\u0000\u0000\u015b\u0159\u0001\u0000\u0000\u0000\u015c"+
		"\u015d\u0003>\u001f\u0000\u015d\u015e\u0005\u0017\u0000\u0000\u015eA\u0001"+
		"\u0000\u0000\u0000\u015f\u0160\u0005\u001e\u0000\u0000\u0160\u0161\u0005"+
		",\u0000\u0000\u0161\u0162\u0003\u0000\u0000\u0000\u0162\u0163\u00057\u0000"+
		"\u0000\u0163\u0164\u0003\u0000\u0000\u0000\u0164\u0165\u0005\u0016\u0000"+
		"\u0000\u0165\u0166\u0003\u0000\u0000\u0000\u0166\u0167\u0005\u0017\u0000"+
		"\u0000\u0167C\u0001\u0000\u0000\u0000\u0168\u0169\u00051\u0000\u0000\u0169"+
		"\u016a\u0005+\u0000\u0000\u016a\u016b\u0003\u0000\u0000\u0000\u016bE\u0001"+
		"\u0000\u0000\u0000\u016c\u016d\u0005.\u0000\u0000\u016d\u016e\u0005!\u0000"+
		"\u0000\u016e\u016f\u0003\u0000\u0000\u0000\u016f\u0170\u0003,\u0016\u0000"+
		"\u0170G\u0001\u0000\u0000\u0000\u0171\u0179\u00052\u0000\u0000\u0172\u0173"+
		"\u0003\u0000\u0000\u0000\u0173\u0174\u0005\r\u0000\u0000\u0174\u0175\u0003"+
		"\u0002\u0001\u0000\u0175\u0176\u0005\u0014\u0000\u0000\u0176\u0178\u0001"+
		"\u0000\u0000\u0000\u0177\u0172\u0001\u0000\u0000\u0000\u0178\u017b\u0001"+
		"\u0000\u0000\u0000\u0179\u0177\u0001\u0000\u0000\u0000\u0179\u017a\u0001"+
		"\u0000\u0000\u0000\u017a\u017c\u0001\u0000\u0000\u0000\u017b\u0179\u0001"+
		"\u0000\u0000\u0000\u017c\u017d\u0003\u0000\u0000\u0000\u017d\u017e\u0005"+
		"\r\u0000\u0000\u017e\u017f\u0003\u0002\u0001\u0000\u017f\u0180\u0003,"+
		"\u0016\u0000\u0180I\u0001\u0000\u0000\u0000\u0181\u0182\u0005-\u0000\u0000"+
		"\u0182\u0183\u0003\u0000\u0000\u0000\u0183\u0184\u0003H$\u0000\u0184K"+
		"\u0001\u0000\u0000\u0000\u0185\u0186\u00055\u0000\u0000\u0186\u018c\u0005"+
		"\u0016\u0000\u0000\u0187\u0188\u0003\u0002\u0001\u0000\u0188\u0189\u0005"+
		"\u0014\u0000\u0000\u0189\u018b\u0001\u0000\u0000\u0000\u018a\u0187\u0001"+
		"\u0000\u0000\u0000\u018b\u018e\u0001\u0000\u0000\u0000\u018c\u018a\u0001"+
		"\u0000\u0000\u0000\u018c\u018d\u0001\u0000\u0000\u0000\u018d\u018f\u0001"+
		"\u0000\u0000\u0000\u018e\u018c\u0001\u0000\u0000\u0000\u018f\u0190\u0003"+
		"\u0002\u0001\u0000\u0190\u0191\u0005\u0017\u0000\u0000\u0191M\u0001\u0000"+
		"\u0000\u0000\u0192\u0198\u0005\u0016\u0000\u0000\u0193\u0194\u0003\u0000"+
		"\u0000\u0000\u0194\u0195\u0005\u0014\u0000\u0000\u0195\u0197\u0001\u0000"+
		"\u0000\u0000\u0196\u0193\u0001\u0000\u0000\u0000\u0197\u019a\u0001\u0000"+
		"\u0000\u0000\u0198\u0196\u0001\u0000\u0000\u0000\u0198\u0199\u0001\u0000"+
		"\u0000\u0000\u0199\u019b\u0001\u0000\u0000\u0000\u019a\u0198\u0001\u0000"+
		"\u0000\u0000\u019b\u019c\u0003\u0000\u0000\u0000\u019c\u019d\u0005\u0017"+
		"\u0000\u0000\u019dO\u0001\u0000\u0000\u0000\u019e\u019f\u0005/\u0000\u0000"+
		"\u019f\u01a0\u00050\u0000\u0000\u01a0\u01a2\u0003\u0000\u0000\u0000\u01a1"+
		"\u01a3\u0003N\'\u0000\u01a2\u01a1\u0001\u0000\u0000\u0000\u01a2\u01a3"+
		"\u0001\u0000\u0000\u0000\u01a3\u01a4\u0001\u0000\u0000\u0000\u01a4\u01a5"+
		"\u0003L&\u0000\u01a5Q\u0001\u0000\u0000\u0000\u01a6\u01a7\u0003(\u0014"+
		"\u0000\u01a7\u01a9\u0003*\u0015\u0000\u01a8\u01aa\u0003,\u0016\u0000\u01a9"+
		"\u01a8\u0001\u0000\u0000\u0000\u01a9\u01aa\u0001\u0000\u0000\u0000\u01aa"+
		"S\u0001\u0000\u0000\u0000\u01ab\u01b4\u00034\u001a\u0000\u01ac\u01b4\u0003"+
		"@ \u0000\u01ad\u01b4\u0003B!\u0000\u01ae\u01b4\u0003P(\u0000\u01af\u01b4"+
		"\u0003J%\u0000\u01b0\u01b4\u0003F#\u0000\u01b1\u01b4\u0003D\"\u0000\u01b2"+
		"\u01b4\u0003R)\u0000\u01b3\u01ab\u0001\u0000\u0000\u0000\u01b3\u01ac\u0001"+
		"\u0000\u0000\u0000\u01b3\u01ad\u0001\u0000\u0000\u0000\u01b3\u01ae\u0001"+
		"\u0000\u0000\u0000\u01b3\u01af\u0001\u0000\u0000\u0000\u01b3\u01b0\u0001"+
		"\u0000\u0000\u0000\u01b3\u01b1\u0001\u0000\u0000\u0000\u01b3\u01b2\u0001"+
		"\u0000\u0000\u0000\u01b4U\u0001\u0000\u0000\u0000\u01b5\u01b6\u0003T*"+
		"\u0000\u01b6\u01b7\u0005\u0000\u0000\u0001\u01b7W\u0001\u0000\u0000\u0000"+
		"!cgv\u0081\u0089\u0091\u0093\u00a5\u00b1\u00b5\u00ca\u00d1\u00e5\u00ee"+
		"\u00f7\u0100\u010a\u0117\u011f\u0125\u0128\u012c\u012f\u0143\u014a\u014d"+
		"\u0159\u0179\u018c\u0198\u01a2\u01a9\u01b3";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}