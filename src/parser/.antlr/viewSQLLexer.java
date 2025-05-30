// Generated from /Users/a.rijo/Documents/University_6th_year/goProjects/sqlToKeyValue/src/parser/ViewSQL.g4 by ANTLR 4.7.1
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class ViewSQLLexer extends Lexer {
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
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	public static final String[] ruleNames = {
		"ADD", "SUB", "MULT", "DIV", "SUM", "AVG", "MAX", "MIN", "COUNT", "AND", 
		"ASC", "DESC", "EQUAL", "HIGHER", "HIGHER_EQUAL", "LOWER", "LOWER_EQUAL", 
		"NOT_EQUAL", "DOT", "SEPARATOR", "RANGE_SEP", "LEFT_P", "RIGHT_P", "LEFT_SP", 
		"RIGHT_SP", "INV_COMMA", "LEFT_CURLY", "RIGHT_CURLY", "AS", "CREATE", 
		"VIEW", "SELECT", "FROM", "WHERE", "GROUP", "BY", "ORDER", "LIMIT", "WITH", 
		"IN", "PRIMARY", "KEY", "TABLE", "INDEX", "UPDATE", "DELETE", "INSERT", 
		"INTO", "DROP", "SET", "FOREIGN", "REFERENCES", "VALUES", "CHECK", "ON", 
		"AW", "RW", "LWW", "DEFAULT", "INTEGER", "COUNTER", "BOOLEAN", "VARCHAR", 
		"DATE_TYPE", "DATE", "DAY", "MONTH", "YEAR", "BOOL", "STRING", "INT", 
		"FLOAT", "WHITESPACE"
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


	public ViewSQLLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "ViewSQL.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2H\u0205\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64\t"+
		"\64\4\65\t\65\4\66\t\66\4\67\t\67\48\t8\49\t9\4:\t:\4;\t;\4<\t<\4=\t="+
		"\4>\t>\4?\t?\4@\t@\4A\tA\4B\tB\4C\tC\4D\tD\4E\tE\4F\tF\4G\tG\4H\tH\4I"+
		"\tI\4J\tJ\3\2\3\2\3\3\3\3\3\4\3\4\3\5\3\5\3\6\3\6\3\6\3\6\3\7\3\7\3\7"+
		"\3\7\3\b\3\b\3\b\3\b\3\t\3\t\3\t\3\t\3\n\3\n\3\n\3\n\3\n\3\n\3\13\3\13"+
		"\3\13\3\13\3\f\3\f\3\f\3\f\3\r\3\r\3\r\3\r\3\r\3\16\3\16\3\17\3\17\3\20"+
		"\3\20\3\20\3\21\3\21\3\22\3\22\3\22\3\23\3\23\3\23\3\24\3\24\3\25\3\25"+
		"\3\26\3\26\3\27\3\27\3\30\3\30\3\31\3\31\3\32\3\32\3\33\3\33\3\34\3\34"+
		"\3\35\3\35\3\36\3\36\3\36\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3 \3 \3 "+
		"\3 \3 \3!\3!\3!\3!\3!\3!\3!\3\"\3\"\3\"\3\"\3\"\3#\3#\3#\3#\3#\3#\3$\3"+
		"$\3$\3$\3$\3$\3%\3%\3%\3&\3&\3&\3&\3&\3&\3\'\3\'\3\'\3\'\3\'\3\'\3(\3"+
		"(\3(\3(\3(\3)\3)\3)\3*\3*\3*\3*\3*\3*\3*\3*\3+\3+\3+\3+\3,\3,\3,\3,\3"+
		",\3,\3-\3-\3-\3-\3-\3-\3.\3.\3.\3.\3.\3.\3.\3/\3/\3/\3/\3/\3/\3/\3\60"+
		"\3\60\3\60\3\60\3\60\3\60\3\60\3\61\3\61\3\61\3\61\3\61\3\62\3\62\3\62"+
		"\3\62\3\62\3\63\3\63\3\63\3\63\3\64\3\64\3\64\3\64\3\64\3\64\3\64\3\64"+
		"\3\65\3\65\3\65\3\65\3\65\3\65\3\65\3\65\3\65\3\65\3\65\3\66\3\66\3\66"+
		"\3\66\3\66\3\66\3\66\3\67\3\67\3\67\3\67\3\67\3\67\38\38\38\39\39\39\3"+
		":\3:\3:\3;\3;\3;\3;\3<\3<\3<\3<\3<\3<\3<\3<\3=\3=\3=\3=\3=\3=\3=\3=\3"+
		">\3>\3>\3>\3>\3>\3>\3>\3?\3?\3?\3?\3?\3?\3?\3?\3@\3@\3@\3@\3@\3@\3@\3"+
		"@\3A\3A\3A\3A\3A\3B\3B\3B\3B\3B\3B\3C\3C\3C\3C\3C\3C\5C\u01c3\nC\3D\3"+
		"D\3D\3D\5D\u01c9\nD\3E\3E\3E\3E\3E\3F\3F\3F\3F\3F\3F\3F\3F\3F\5F\u01d9"+
		"\nF\3G\3G\7G\u01dd\nG\fG\16G\u01e0\13G\3H\3H\5H\u01e4\nH\3H\3H\7H\u01e8"+
		"\nH\fH\16H\u01eb\13H\5H\u01ed\nH\3I\5I\u01f0\nI\3I\3I\3I\7I\u01f5\nI\f"+
		"I\16I\u01f8\13I\3I\6I\u01fb\nI\rI\16I\u01fc\3J\6J\u0200\nJ\rJ\16J\u0201"+
		"\3J\3J\2\2K\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13\25\f\27\r\31\16\33"+
		"\17\35\20\37\21!\22#\23%\24\'\25)\26+\27-\30/\31\61\32\63\33\65\34\67"+
		"\359\36;\37= ?!A\"C#E$G%I&K\'M(O)Q*S+U,W-Y.[/]\60_\61a\62c\63e\64g\65"+
		"i\66k\67m8o9q:s;u<w=y>{?}@\177A\u0081B\u0083C\u0085\2\u0087\2\u0089\2"+
		"\u008bD\u008dE\u008fF\u0091G\u0093H\3\2\t\3\2\63;\3\2\63\64\3\2\62;\3"+
		"\2\62\63\5\2C\\aac|\6\2\62;C\\aac|\5\2\13\f\17\17\"\"\2\u020d\2\3\3\2"+
		"\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17"+
		"\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2\2\31\3\2"+
		"\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2\2\2\2%\3"+
		"\2\2\2\2\'\3\2\2\2\2)\3\2\2\2\2+\3\2\2\2\2-\3\2\2\2\2/\3\2\2\2\2\61\3"+
		"\2\2\2\2\63\3\2\2\2\2\65\3\2\2\2\2\67\3\2\2\2\29\3\2\2\2\2;\3\2\2\2\2"+
		"=\3\2\2\2\2?\3\2\2\2\2A\3\2\2\2\2C\3\2\2\2\2E\3\2\2\2\2G\3\2\2\2\2I\3"+
		"\2\2\2\2K\3\2\2\2\2M\3\2\2\2\2O\3\2\2\2\2Q\3\2\2\2\2S\3\2\2\2\2U\3\2\2"+
		"\2\2W\3\2\2\2\2Y\3\2\2\2\2[\3\2\2\2\2]\3\2\2\2\2_\3\2\2\2\2a\3\2\2\2\2"+
		"c\3\2\2\2\2e\3\2\2\2\2g\3\2\2\2\2i\3\2\2\2\2k\3\2\2\2\2m\3\2\2\2\2o\3"+
		"\2\2\2\2q\3\2\2\2\2s\3\2\2\2\2u\3\2\2\2\2w\3\2\2\2\2y\3\2\2\2\2{\3\2\2"+
		"\2\2}\3\2\2\2\2\177\3\2\2\2\2\u0081\3\2\2\2\2\u0083\3\2\2\2\2\u008b\3"+
		"\2\2\2\2\u008d\3\2\2\2\2\u008f\3\2\2\2\2\u0091\3\2\2\2\2\u0093\3\2\2\2"+
		"\3\u0095\3\2\2\2\5\u0097\3\2\2\2\7\u0099\3\2\2\2\t\u009b\3\2\2\2\13\u009d"+
		"\3\2\2\2\r\u00a1\3\2\2\2\17\u00a5\3\2\2\2\21\u00a9\3\2\2\2\23\u00ad\3"+
		"\2\2\2\25\u00b3\3\2\2\2\27\u00b7\3\2\2\2\31\u00bb\3\2\2\2\33\u00c0\3\2"+
		"\2\2\35\u00c2\3\2\2\2\37\u00c4\3\2\2\2!\u00c7\3\2\2\2#\u00c9\3\2\2\2%"+
		"\u00cc\3\2\2\2\'\u00cf\3\2\2\2)\u00d1\3\2\2\2+\u00d3\3\2\2\2-\u00d5\3"+
		"\2\2\2/\u00d7\3\2\2\2\61\u00d9\3\2\2\2\63\u00db\3\2\2\2\65\u00dd\3\2\2"+
		"\2\67\u00df\3\2\2\29\u00e1\3\2\2\2;\u00e3\3\2\2\2=\u00e6\3\2\2\2?\u00ed"+
		"\3\2\2\2A\u00f2\3\2\2\2C\u00f9\3\2\2\2E\u00fe\3\2\2\2G\u0104\3\2\2\2I"+
		"\u010a\3\2\2\2K\u010d\3\2\2\2M\u0113\3\2\2\2O\u0119\3\2\2\2Q\u011e\3\2"+
		"\2\2S\u0121\3\2\2\2U\u0129\3\2\2\2W\u012d\3\2\2\2Y\u0133\3\2\2\2[\u0139"+
		"\3\2\2\2]\u0140\3\2\2\2_\u0147\3\2\2\2a\u014e\3\2\2\2c\u0153\3\2\2\2e"+
		"\u0158\3\2\2\2g\u015c\3\2\2\2i\u0164\3\2\2\2k\u016f\3\2\2\2m\u0176\3\2"+
		"\2\2o\u017c\3\2\2\2q\u017f\3\2\2\2s\u0182\3\2\2\2u\u0185\3\2\2\2w\u0189"+
		"\3\2\2\2y\u0191\3\2\2\2{\u0199\3\2\2\2}\u01a1\3\2\2\2\177\u01a9\3\2\2"+
		"\2\u0081\u01b1\3\2\2\2\u0083\u01b6\3\2\2\2\u0085\u01c2\3\2\2\2\u0087\u01c8"+
		"\3\2\2\2\u0089\u01ca\3\2\2\2\u008b\u01d8\3\2\2\2\u008d\u01da\3\2\2\2\u008f"+
		"\u01ec\3\2\2\2\u0091\u01ef\3\2\2\2\u0093\u01ff\3\2\2\2\u0095\u0096\7-"+
		"\2\2\u0096\4\3\2\2\2\u0097\u0098\7/\2\2\u0098\6\3\2\2\2\u0099\u009a\7"+
		",\2\2\u009a\b\3\2\2\2\u009b\u009c\7\61\2\2\u009c\n\3\2\2\2\u009d\u009e"+
		"\7U\2\2\u009e\u009f\7W\2\2\u009f\u00a0\7O\2\2\u00a0\f\3\2\2\2\u00a1\u00a2"+
		"\7C\2\2\u00a2\u00a3\7X\2\2\u00a3\u00a4\7I\2\2\u00a4\16\3\2\2\2\u00a5\u00a6"+
		"\7O\2\2\u00a6\u00a7\7C\2\2\u00a7\u00a8\7Z\2\2\u00a8\20\3\2\2\2\u00a9\u00aa"+
		"\7O\2\2\u00aa\u00ab\7K\2\2\u00ab\u00ac\7P\2\2\u00ac\22\3\2\2\2\u00ad\u00ae"+
		"\7E\2\2\u00ae\u00af\7Q\2\2\u00af\u00b0\7W\2\2\u00b0\u00b1\7P\2\2\u00b1"+
		"\u00b2\7V\2\2\u00b2\24\3\2\2\2\u00b3\u00b4\7C\2\2\u00b4\u00b5\7P\2\2\u00b5"+
		"\u00b6\7F\2\2\u00b6\26\3\2\2\2\u00b7\u00b8\7C\2\2\u00b8\u00b9\7U\2\2\u00b9"+
		"\u00ba\7E\2\2\u00ba\30\3\2\2\2\u00bb\u00bc\7F\2\2\u00bc\u00bd\7G\2\2\u00bd"+
		"\u00be\7U\2\2\u00be\u00bf\7E\2\2\u00bf\32\3\2\2\2\u00c0\u00c1\7?\2\2\u00c1"+
		"\34\3\2\2\2\u00c2\u00c3\7@\2\2\u00c3\36\3\2\2\2\u00c4\u00c5\7@\2\2\u00c5"+
		"\u00c6\7?\2\2\u00c6 \3\2\2\2\u00c7\u00c8\7>\2\2\u00c8\"\3\2\2\2\u00c9"+
		"\u00ca\7>\2\2\u00ca\u00cb\7?\2\2\u00cb$\3\2\2\2\u00cc\u00cd\7#\2\2\u00cd"+
		"\u00ce\7?\2\2\u00ce&\3\2\2\2\u00cf\u00d0\7\60\2\2\u00d0(\3\2\2\2\u00d1"+
		"\u00d2\7.\2\2\u00d2*\3\2\2\2\u00d3\u00d4\7<\2\2\u00d4,\3\2\2\2\u00d5\u00d6"+
		"\7*\2\2\u00d6.\3\2\2\2\u00d7\u00d8\7+\2\2\u00d8\60\3\2\2\2\u00d9\u00da"+
		"\7]\2\2\u00da\62\3\2\2\2\u00db\u00dc\7_\2\2\u00dc\64\3\2\2\2\u00dd\u00de"+
		"\7$\2\2\u00de\66\3\2\2\2\u00df\u00e0\7}\2\2\u00e08\3\2\2\2\u00e1\u00e2"+
		"\7\177\2\2\u00e2:\3\2\2\2\u00e3\u00e4\7C\2\2\u00e4\u00e5\7U\2\2\u00e5"+
		"<\3\2\2\2\u00e6\u00e7\7E\2\2\u00e7\u00e8\7T\2\2\u00e8\u00e9\7G\2\2\u00e9"+
		"\u00ea\7C\2\2\u00ea\u00eb\7V\2\2\u00eb\u00ec\7G\2\2\u00ec>\3\2\2\2\u00ed"+
		"\u00ee\7X\2\2\u00ee\u00ef\7K\2\2\u00ef\u00f0\7G\2\2\u00f0\u00f1\7Y\2\2"+
		"\u00f1@\3\2\2\2\u00f2\u00f3\7U\2\2\u00f3\u00f4\7G\2\2\u00f4\u00f5\7N\2"+
		"\2\u00f5\u00f6\7G\2\2\u00f6\u00f7\7E\2\2\u00f7\u00f8\7V\2\2\u00f8B\3\2"+
		"\2\2\u00f9\u00fa\7H\2\2\u00fa\u00fb\7T\2\2\u00fb\u00fc\7Q\2\2\u00fc\u00fd"+
		"\7O\2\2\u00fdD\3\2\2\2\u00fe\u00ff\7Y\2\2\u00ff\u0100\7J\2\2\u0100\u0101"+
		"\7G\2\2\u0101\u0102\7T\2\2\u0102\u0103\7G\2\2\u0103F\3\2\2\2\u0104\u0105"+
		"\7I\2\2\u0105\u0106\7T\2\2\u0106\u0107\7Q\2\2\u0107\u0108\7W\2\2\u0108"+
		"\u0109\7R\2\2\u0109H\3\2\2\2\u010a\u010b\7D\2\2\u010b\u010c\7[\2\2\u010c"+
		"J\3\2\2\2\u010d\u010e\7Q\2\2\u010e\u010f\7T\2\2\u010f\u0110\7F\2\2\u0110"+
		"\u0111\7G\2\2\u0111\u0112\7T\2\2\u0112L\3\2\2\2\u0113\u0114\7N\2\2\u0114"+
		"\u0115\7K\2\2\u0115\u0116\7O\2\2\u0116\u0117\7K\2\2\u0117\u0118\7V\2\2"+
		"\u0118N\3\2\2\2\u0119\u011a\7Y\2\2\u011a\u011b\7K\2\2\u011b\u011c\7V\2"+
		"\2\u011c\u011d\7J\2\2\u011dP\3\2\2\2\u011e\u011f\7K\2\2\u011f\u0120\7"+
		"P\2\2\u0120R\3\2\2\2\u0121\u0122\7R\2\2\u0122\u0123\7T\2\2\u0123\u0124"+
		"\7K\2\2\u0124\u0125\7O\2\2\u0125\u0126\7C\2\2\u0126\u0127\7T\2\2\u0127"+
		"\u0128\7[\2\2\u0128T\3\2\2\2\u0129\u012a\7M\2\2\u012a\u012b\7G\2\2\u012b"+
		"\u012c\7[\2\2\u012cV\3\2\2\2\u012d\u012e\7V\2\2\u012e\u012f\7C\2\2\u012f"+
		"\u0130\7D\2\2\u0130\u0131\7N\2\2\u0131\u0132\7G\2\2\u0132X\3\2\2\2\u0133"+
		"\u0134\7K\2\2\u0134\u0135\7P\2\2\u0135\u0136\7F\2\2\u0136\u0137\7G\2\2"+
		"\u0137\u0138\7Z\2\2\u0138Z\3\2\2\2\u0139\u013a\7W\2\2\u013a\u013b\7R\2"+
		"\2\u013b\u013c\7F\2\2\u013c\u013d\7C\2\2\u013d\u013e\7V\2\2\u013e\u013f"+
		"\7G\2\2\u013f\\\3\2\2\2\u0140\u0141\7F\2\2\u0141\u0142\7G\2\2\u0142\u0143"+
		"\7N\2\2\u0143\u0144\7G\2\2\u0144\u0145\7V\2\2\u0145\u0146\7G\2\2\u0146"+
		"^\3\2\2\2\u0147\u0148\7K\2\2\u0148\u0149\7P\2\2\u0149\u014a\7U\2\2\u014a"+
		"\u014b\7G\2\2\u014b\u014c\7T\2\2\u014c\u014d\7V\2\2\u014d`\3\2\2\2\u014e"+
		"\u014f\7K\2\2\u014f\u0150\7P\2\2\u0150\u0151\7V\2\2\u0151\u0152\7Q\2\2"+
		"\u0152b\3\2\2\2\u0153\u0154\7F\2\2\u0154\u0155\7T\2\2\u0155\u0156\7Q\2"+
		"\2\u0156\u0157\7R\2\2\u0157d\3\2\2\2\u0158\u0159\7U\2\2\u0159\u015a\7"+
		"G\2\2\u015a\u015b\7V\2\2\u015bf\3\2\2\2\u015c\u015d\7H\2\2\u015d\u015e"+
		"\7Q\2\2\u015e\u015f\7T\2\2\u015f\u0160\7G\2\2\u0160\u0161\7K\2\2\u0161"+
		"\u0162\7I\2\2\u0162\u0163\7P\2\2\u0163h\3\2\2\2\u0164\u0165\7T\2\2\u0165"+
		"\u0166\7G\2\2\u0166\u0167\7H\2\2\u0167\u0168\7G\2\2\u0168\u0169\7T\2\2"+
		"\u0169\u016a\7G\2\2\u016a\u016b\7P\2\2\u016b\u016c\7E\2\2\u016c\u016d"+
		"\7G\2\2\u016d\u016e\7U\2\2\u016ej\3\2\2\2\u016f\u0170\7X\2\2\u0170\u0171"+
		"\7C\2\2\u0171\u0172\7N\2\2\u0172\u0173\7W\2\2\u0173\u0174\7G\2\2\u0174"+
		"\u0175\7U\2\2\u0175l\3\2\2\2\u0176\u0177\7E\2\2\u0177\u0178\7J\2\2\u0178"+
		"\u0179\7G\2\2\u0179\u017a\7E\2\2\u017a\u017b\7M\2\2\u017bn\3\2\2\2\u017c"+
		"\u017d\7Q\2\2\u017d\u017e\7P\2\2\u017ep\3\2\2\2\u017f\u0180\7C\2\2\u0180"+
		"\u0181\7Y\2\2\u0181r\3\2\2\2\u0182\u0183\7T\2\2\u0183\u0184\7Y\2\2\u0184"+
		"t\3\2\2\2\u0185\u0186\7N\2\2\u0186\u0187\7Y\2\2\u0187\u0188\7Y\2\2\u0188"+
		"v\3\2\2\2\u0189\u018a\7F\2\2\u018a\u018b\7G\2\2\u018b\u018c\7H\2\2\u018c"+
		"\u018d\7C\2\2\u018d\u018e\7W\2\2\u018e\u018f\7N\2\2\u018f\u0190\7V\2\2"+
		"\u0190x\3\2\2\2\u0191\u0192\7k\2\2\u0192\u0193\7p\2\2\u0193\u0194\7v\2"+
		"\2\u0194\u0195\7g\2\2\u0195\u0196\7i\2\2\u0196\u0197\7g\2\2\u0197\u0198"+
		"\7t\2\2\u0198z\3\2\2\2\u0199\u019a\7e\2\2\u019a\u019b\7q\2\2\u019b\u019c"+
		"\7w\2\2\u019c\u019d\7p\2\2\u019d\u019e\7v\2\2\u019e\u019f\7g\2\2\u019f"+
		"\u01a0\7t\2\2\u01a0|\3\2\2\2\u01a1\u01a2\7d\2\2\u01a2\u01a3\7q\2\2\u01a3"+
		"\u01a4\7q\2\2\u01a4\u01a5\7n\2\2\u01a5\u01a6\7g\2\2\u01a6\u01a7\7c\2\2"+
		"\u01a7\u01a8\7p\2\2\u01a8~\3\2\2\2\u01a9\u01aa\7x\2\2\u01aa\u01ab\7c\2"+
		"\2\u01ab\u01ac\7t\2\2\u01ac\u01ad\7e\2\2\u01ad\u01ae\7j\2\2\u01ae\u01af"+
		"\7c\2\2\u01af\u01b0\7t\2\2\u01b0\u0080\3\2\2\2\u01b1\u01b2\7f\2\2\u01b2"+
		"\u01b3\7c\2\2\u01b3\u01b4\7v\2\2\u01b4\u01b5\7g\2\2\u01b5\u0082\3\2\2"+
		"\2\u01b6\u01b7\5\u0089E\2\u01b7\u01b8\7/\2\2\u01b8\u01b9\5\u0087D\2\u01b9"+
		"\u01ba\7/\2\2\u01ba\u01bb\5\u0085C\2\u01bb\u0084\3\2\2\2\u01bc\u01bd\7"+
		"\62\2\2\u01bd\u01c3\t\2\2\2\u01be\u01bf\t\3\2\2\u01bf\u01c3\t\4\2\2\u01c0"+
		"\u01c1\7\65\2\2\u01c1\u01c3\t\5\2\2\u01c2\u01bc\3\2\2\2\u01c2\u01be\3"+
		"\2\2\2\u01c2\u01c0\3\2\2\2\u01c3\u0086\3\2\2\2\u01c4\u01c5\7\62\2\2\u01c5"+
		"\u01c9\t\2\2\2\u01c6\u01c7\7\63\2\2\u01c7\u01c9\t\3\2\2\u01c8\u01c4\3"+
		"\2\2\2\u01c8\u01c6\3\2\2\2\u01c9\u0088\3\2\2\2\u01ca\u01cb\t\4\2\2\u01cb"+
		"\u01cc\t\4\2\2\u01cc\u01cd\t\4\2\2\u01cd\u01ce\t\4\2\2\u01ce\u008a\3\2"+
		"\2\2\u01cf\u01d0\7v\2\2\u01d0\u01d1\7t\2\2\u01d1\u01d2\7w\2\2\u01d2\u01d9"+
		"\7g\2\2\u01d3\u01d4\7h\2\2\u01d4\u01d5\7c\2\2\u01d5\u01d6\7n\2\2\u01d6"+
		"\u01d7\7u\2\2\u01d7\u01d9\7g\2\2\u01d8\u01cf\3\2\2\2\u01d8\u01d3\3\2\2"+
		"\2\u01d9\u008c\3\2\2\2\u01da\u01de\t\6\2\2\u01db\u01dd\t\7\2\2\u01dc\u01db"+
		"\3\2\2\2\u01dd\u01e0\3\2\2\2\u01de\u01dc\3\2\2\2\u01de\u01df\3\2\2\2\u01df"+
		"\u008e\3\2\2\2\u01e0\u01de\3\2\2\2\u01e1\u01ed\7\62\2\2\u01e2\u01e4\7"+
		"/\2\2\u01e3\u01e2\3\2\2\2\u01e3\u01e4\3\2\2\2\u01e4\u01e5\3\2\2\2\u01e5"+
		"\u01e9\t\2\2\2\u01e6\u01e8\t\4\2\2\u01e7\u01e6\3\2\2\2\u01e8\u01eb\3\2"+
		"\2\2\u01e9\u01e7\3\2\2\2\u01e9\u01ea\3\2\2\2\u01ea\u01ed\3\2\2\2\u01eb"+
		"\u01e9\3\2\2\2\u01ec\u01e1\3\2\2\2\u01ec\u01e3\3\2\2\2\u01ed\u0090\3\2"+
		"\2\2\u01ee\u01f0\7/\2\2\u01ef\u01ee\3\2\2\2\u01ef\u01f0\3\2\2\2\u01f0"+
		"\u01f1\3\2\2\2\u01f1\u01f2\t\4\2\2\u01f2\u01f6\7\60\2\2\u01f3\u01f5\t"+
		"\4\2\2\u01f4\u01f3\3\2\2\2\u01f5\u01f8\3\2\2\2\u01f6\u01f4\3\2\2\2\u01f6"+
		"\u01f7\3\2\2\2\u01f7\u01fa\3\2\2\2\u01f8\u01f6\3\2\2\2\u01f9\u01fb\t\2"+
		"\2\2\u01fa\u01f9\3\2\2\2\u01fb\u01fc\3\2\2\2\u01fc\u01fa\3\2\2\2\u01fc"+
		"\u01fd\3\2\2\2\u01fd\u0092\3\2\2\2\u01fe\u0200\t\b\2\2\u01ff\u01fe\3\2"+
		"\2\2\u0200\u0201\3\2\2\2\u0201\u01ff\3\2\2\2\u0201\u0202\3\2\2\2\u0202"+
		"\u0203\3\2\2\2\u0203\u0204\bJ\2\2\u0204\u0094\3\2\2\2\16\2\u01c2\u01c8"+
		"\u01d8\u01de\u01e3\u01e9\u01ec\u01ef\u01f6\u01fc\u0201\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}