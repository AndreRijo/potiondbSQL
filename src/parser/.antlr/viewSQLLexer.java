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
		ORDER=37, LIMIT=38, WITH=39, IN=40, PRIMARY=41, KEY=42, DATE=43, STRING=44, 
		INT=45, FLOAT=46, WHITESPACE=47;
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
		"IN", "PRIMARY", "KEY", "DATE", "DAY", "MONTH", "YEAR", "STRING", "INT", 
		"FLOAT", "WHITESPACE"
	};

	private static final String[] _LITERAL_NAMES = {
		null, "'+'", "'-'", "'*'", "'/'", "'SUM'", "'AVG'", "'MAX'", "'MIN'", 
		"'COUNT'", "'AND'", "'ASC'", "'DESC'", "'='", "'>'", "'>='", "'<'", "'<='", 
		"'!='", "'.'", "','", "':'", "'('", "')'", "'['", "']'", "'\"'", "'{'", 
		"'}'", "'AS'", "'CREATE'", "'VIEW'", "'SELECT'", "'FROM'", "'WHERE'", 
		"'GROUP'", "'BY'", "'ORDER'", "'LIMIT'", "'WITH'", "'IN'", "'PRIMARY'", 
		"'KEY'"
	};
	private static final String[] _SYMBOLIC_NAMES = {
		null, "ADD", "SUB", "MULT", "DIV", "SUM", "AVG", "MAX", "MIN", "COUNT", 
		"AND", "ASC", "DESC", "EQUAL", "HIGHER", "HIGHER_EQUAL", "LOWER", "LOWER_EQUAL", 
		"NOT_EQUAL", "DOT", "SEPARATOR", "RANGE_SEP", "LEFT_P", "RIGHT_P", "LEFT_SP", 
		"RIGHT_SP", "INV_COMMA", "LEFT_CURLY", "RIGHT_CURLY", "AS", "CREATE", 
		"VIEW", "SELECT", "FROM", "WHERE", "GROUP", "BY", "ORDER", "LIMIT", "WITH", 
		"IN", "PRIMARY", "KEY", "DATE", "STRING", "INT", "FLOAT", "WHITESPACE"
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\61\u0143\b\1\4\2"+
		"\t\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4"+
		"\13\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22"+
		"\t\22\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31"+
		"\t\31\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t"+
		" \4!\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t"+
		"+\4,\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\3\2"+
		"\3\2\3\3\3\3\3\4\3\4\3\5\3\5\3\6\3\6\3\6\3\6\3\7\3\7\3\7\3\7\3\b\3\b\3"+
		"\b\3\b\3\t\3\t\3\t\3\t\3\n\3\n\3\n\3\n\3\n\3\n\3\13\3\13\3\13\3\13\3\f"+
		"\3\f\3\f\3\f\3\r\3\r\3\r\3\r\3\r\3\16\3\16\3\17\3\17\3\20\3\20\3\20\3"+
		"\21\3\21\3\22\3\22\3\22\3\23\3\23\3\23\3\24\3\24\3\25\3\25\3\26\3\26\3"+
		"\27\3\27\3\30\3\30\3\31\3\31\3\32\3\32\3\33\3\33\3\34\3\34\3\35\3\35\3"+
		"\36\3\36\3\36\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3 \3 \3 \3 \3 \3!\3!"+
		"\3!\3!\3!\3!\3!\3\"\3\"\3\"\3\"\3\"\3#\3#\3#\3#\3#\3#\3$\3$\3$\3$\3$\3"+
		"$\3%\3%\3%\3&\3&\3&\3&\3&\3&\3\'\3\'\3\'\3\'\3\'\3\'\3(\3(\3(\3(\3(\3"+
		")\3)\3)\3*\3*\3*\3*\3*\3*\3*\3*\3+\3+\3+\3+\3,\3,\3,\3,\3,\3,\3-\3-\3"+
		"-\3-\3-\3-\5-\u010c\n-\3.\3.\3.\3.\5.\u0112\n.\3/\3/\3/\3/\3/\3\60\3\60"+
		"\7\60\u011b\n\60\f\60\16\60\u011e\13\60\3\61\3\61\5\61\u0122\n\61\3\61"+
		"\3\61\7\61\u0126\n\61\f\61\16\61\u0129\13\61\5\61\u012b\n\61\3\62\5\62"+
		"\u012e\n\62\3\62\3\62\3\62\7\62\u0133\n\62\f\62\16\62\u0136\13\62\3\62"+
		"\6\62\u0139\n\62\r\62\16\62\u013a\3\63\6\63\u013e\n\63\r\63\16\63\u013f"+
		"\3\63\3\63\2\2\64\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13\25\f\27\r\31"+
		"\16\33\17\35\20\37\21!\22#\23%\24\'\25)\26+\27-\30/\31\61\32\63\33\65"+
		"\34\67\359\36;\37= ?!A\"C#E$G%I&K\'M(O)Q*S+U,W-Y\2[\2]\2_.a/c\60e\61\3"+
		"\2\t\3\2\63;\3\2\63\64\3\2\62;\3\2\62\63\5\2C\\aac|\6\2\62;C\\aac|\5\2"+
		"\13\f\17\17\"\"\2\u014a\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2"+
		"\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25"+
		"\3\2\2\2\2\27\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3\2"+
		"\2\2\2!\3\2\2\2\2#\3\2\2\2\2%\3\2\2\2\2\'\3\2\2\2\2)\3\2\2\2\2+\3\2\2"+
		"\2\2-\3\2\2\2\2/\3\2\2\2\2\61\3\2\2\2\2\63\3\2\2\2\2\65\3\2\2\2\2\67\3"+
		"\2\2\2\29\3\2\2\2\2;\3\2\2\2\2=\3\2\2\2\2?\3\2\2\2\2A\3\2\2\2\2C\3\2\2"+
		"\2\2E\3\2\2\2\2G\3\2\2\2\2I\3\2\2\2\2K\3\2\2\2\2M\3\2\2\2\2O\3\2\2\2\2"+
		"Q\3\2\2\2\2S\3\2\2\2\2U\3\2\2\2\2W\3\2\2\2\2_\3\2\2\2\2a\3\2\2\2\2c\3"+
		"\2\2\2\2e\3\2\2\2\3g\3\2\2\2\5i\3\2\2\2\7k\3\2\2\2\tm\3\2\2\2\13o\3\2"+
		"\2\2\rs\3\2\2\2\17w\3\2\2\2\21{\3\2\2\2\23\177\3\2\2\2\25\u0085\3\2\2"+
		"\2\27\u0089\3\2\2\2\31\u008d\3\2\2\2\33\u0092\3\2\2\2\35\u0094\3\2\2\2"+
		"\37\u0096\3\2\2\2!\u0099\3\2\2\2#\u009b\3\2\2\2%\u009e\3\2\2\2\'\u00a1"+
		"\3\2\2\2)\u00a3\3\2\2\2+\u00a5\3\2\2\2-\u00a7\3\2\2\2/\u00a9\3\2\2\2\61"+
		"\u00ab\3\2\2\2\63\u00ad\3\2\2\2\65\u00af\3\2\2\2\67\u00b1\3\2\2\29\u00b3"+
		"\3\2\2\2;\u00b5\3\2\2\2=\u00b8\3\2\2\2?\u00bf\3\2\2\2A\u00c4\3\2\2\2C"+
		"\u00cb\3\2\2\2E\u00d0\3\2\2\2G\u00d6\3\2\2\2I\u00dc\3\2\2\2K\u00df\3\2"+
		"\2\2M\u00e5\3\2\2\2O\u00eb\3\2\2\2Q\u00f0\3\2\2\2S\u00f3\3\2\2\2U\u00fb"+
		"\3\2\2\2W\u00ff\3\2\2\2Y\u010b\3\2\2\2[\u0111\3\2\2\2]\u0113\3\2\2\2_"+
		"\u0118\3\2\2\2a\u012a\3\2\2\2c\u012d\3\2\2\2e\u013d\3\2\2\2gh\7-\2\2h"+
		"\4\3\2\2\2ij\7/\2\2j\6\3\2\2\2kl\7,\2\2l\b\3\2\2\2mn\7\61\2\2n\n\3\2\2"+
		"\2op\7U\2\2pq\7W\2\2qr\7O\2\2r\f\3\2\2\2st\7C\2\2tu\7X\2\2uv\7I\2\2v\16"+
		"\3\2\2\2wx\7O\2\2xy\7C\2\2yz\7Z\2\2z\20\3\2\2\2{|\7O\2\2|}\7K\2\2}~\7"+
		"P\2\2~\22\3\2\2\2\177\u0080\7E\2\2\u0080\u0081\7Q\2\2\u0081\u0082\7W\2"+
		"\2\u0082\u0083\7P\2\2\u0083\u0084\7V\2\2\u0084\24\3\2\2\2\u0085\u0086"+
		"\7C\2\2\u0086\u0087\7P\2\2\u0087\u0088\7F\2\2\u0088\26\3\2\2\2\u0089\u008a"+
		"\7C\2\2\u008a\u008b\7U\2\2\u008b\u008c\7E\2\2\u008c\30\3\2\2\2\u008d\u008e"+
		"\7F\2\2\u008e\u008f\7G\2\2\u008f\u0090\7U\2\2\u0090\u0091\7E\2\2\u0091"+
		"\32\3\2\2\2\u0092\u0093\7?\2\2\u0093\34\3\2\2\2\u0094\u0095\7@\2\2\u0095"+
		"\36\3\2\2\2\u0096\u0097\7@\2\2\u0097\u0098\7?\2\2\u0098 \3\2\2\2\u0099"+
		"\u009a\7>\2\2\u009a\"\3\2\2\2\u009b\u009c\7>\2\2\u009c\u009d\7?\2\2\u009d"+
		"$\3\2\2\2\u009e\u009f\7#\2\2\u009f\u00a0\7?\2\2\u00a0&\3\2\2\2\u00a1\u00a2"+
		"\7\60\2\2\u00a2(\3\2\2\2\u00a3\u00a4\7.\2\2\u00a4*\3\2\2\2\u00a5\u00a6"+
		"\7<\2\2\u00a6,\3\2\2\2\u00a7\u00a8\7*\2\2\u00a8.\3\2\2\2\u00a9\u00aa\7"+
		"+\2\2\u00aa\60\3\2\2\2\u00ab\u00ac\7]\2\2\u00ac\62\3\2\2\2\u00ad\u00ae"+
		"\7_\2\2\u00ae\64\3\2\2\2\u00af\u00b0\7$\2\2\u00b0\66\3\2\2\2\u00b1\u00b2"+
		"\7}\2\2\u00b28\3\2\2\2\u00b3\u00b4\7\177\2\2\u00b4:\3\2\2\2\u00b5\u00b6"+
		"\7C\2\2\u00b6\u00b7\7U\2\2\u00b7<\3\2\2\2\u00b8\u00b9\7E\2\2\u00b9\u00ba"+
		"\7T\2\2\u00ba\u00bb\7G\2\2\u00bb\u00bc\7C\2\2\u00bc\u00bd\7V\2\2\u00bd"+
		"\u00be\7G\2\2\u00be>\3\2\2\2\u00bf\u00c0\7X\2\2\u00c0\u00c1\7K\2\2\u00c1"+
		"\u00c2\7G\2\2\u00c2\u00c3\7Y\2\2\u00c3@\3\2\2\2\u00c4\u00c5\7U\2\2\u00c5"+
		"\u00c6\7G\2\2\u00c6\u00c7\7N\2\2\u00c7\u00c8\7G\2\2\u00c8\u00c9\7E\2\2"+
		"\u00c9\u00ca\7V\2\2\u00caB\3\2\2\2\u00cb\u00cc\7H\2\2\u00cc\u00cd\7T\2"+
		"\2\u00cd\u00ce\7Q\2\2\u00ce\u00cf\7O\2\2\u00cfD\3\2\2\2\u00d0\u00d1\7"+
		"Y\2\2\u00d1\u00d2\7J\2\2\u00d2\u00d3\7G\2\2\u00d3\u00d4\7T\2\2\u00d4\u00d5"+
		"\7G\2\2\u00d5F\3\2\2\2\u00d6\u00d7\7I\2\2\u00d7\u00d8\7T\2\2\u00d8\u00d9"+
		"\7Q\2\2\u00d9\u00da\7W\2\2\u00da\u00db\7R\2\2\u00dbH\3\2\2\2\u00dc\u00dd"+
		"\7D\2\2\u00dd\u00de\7[\2\2\u00deJ\3\2\2\2\u00df\u00e0\7Q\2\2\u00e0\u00e1"+
		"\7T\2\2\u00e1\u00e2\7F\2\2\u00e2\u00e3\7G\2\2\u00e3\u00e4\7T\2\2\u00e4"+
		"L\3\2\2\2\u00e5\u00e6\7N\2\2\u00e6\u00e7\7K\2\2\u00e7\u00e8\7O\2\2\u00e8"+
		"\u00e9\7K\2\2\u00e9\u00ea\7V\2\2\u00eaN\3\2\2\2\u00eb\u00ec\7Y\2\2\u00ec"+
		"\u00ed\7K\2\2\u00ed\u00ee\7V\2\2\u00ee\u00ef\7J\2\2\u00efP\3\2\2\2\u00f0"+
		"\u00f1\7K\2\2\u00f1\u00f2\7P\2\2\u00f2R\3\2\2\2\u00f3\u00f4\7R\2\2\u00f4"+
		"\u00f5\7T\2\2\u00f5\u00f6\7K\2\2\u00f6\u00f7\7O\2\2\u00f7\u00f8\7C\2\2"+
		"\u00f8\u00f9\7T\2\2\u00f9\u00fa\7[\2\2\u00faT\3\2\2\2\u00fb\u00fc\7M\2"+
		"\2\u00fc\u00fd\7G\2\2\u00fd\u00fe\7[\2\2\u00feV\3\2\2\2\u00ff\u0100\5"+
		"]/\2\u0100\u0101\7/\2\2\u0101\u0102\5[.\2\u0102\u0103\7/\2\2\u0103\u0104"+
		"\5Y-\2\u0104X\3\2\2\2\u0105\u0106\7\62\2\2\u0106\u010c\t\2\2\2\u0107\u0108"+
		"\t\3\2\2\u0108\u010c\t\4\2\2\u0109\u010a\7\65\2\2\u010a\u010c\t\5\2\2"+
		"\u010b\u0105\3\2\2\2\u010b\u0107\3\2\2\2\u010b\u0109\3\2\2\2\u010cZ\3"+
		"\2\2\2\u010d\u010e\7\62\2\2\u010e\u0112\t\2\2\2\u010f\u0110\7\63\2\2\u0110"+
		"\u0112\t\3\2\2\u0111\u010d\3\2\2\2\u0111\u010f\3\2\2\2\u0112\\\3\2\2\2"+
		"\u0113\u0114\t\4\2\2\u0114\u0115\t\4\2\2\u0115\u0116\t\4\2\2\u0116\u0117"+
		"\t\4\2\2\u0117^\3\2\2\2\u0118\u011c\t\6\2\2\u0119\u011b\t\7\2\2\u011a"+
		"\u0119\3\2\2\2\u011b\u011e\3\2\2\2\u011c\u011a\3\2\2\2\u011c\u011d\3\2"+
		"\2\2\u011d`\3\2\2\2\u011e\u011c\3\2\2\2\u011f\u012b\7\62\2\2\u0120\u0122"+
		"\7/\2\2\u0121\u0120\3\2\2\2\u0121\u0122\3\2\2\2\u0122\u0123\3\2\2\2\u0123"+
		"\u0127\t\2\2\2\u0124\u0126\t\4\2\2\u0125\u0124\3\2\2\2\u0126\u0129\3\2"+
		"\2\2\u0127\u0125\3\2\2\2\u0127\u0128\3\2\2\2\u0128\u012b\3\2\2\2\u0129"+
		"\u0127\3\2\2\2\u012a\u011f\3\2\2\2\u012a\u0121\3\2\2\2\u012bb\3\2\2\2"+
		"\u012c\u012e\7/\2\2\u012d\u012c\3\2\2\2\u012d\u012e\3\2\2\2\u012e\u012f"+
		"\3\2\2\2\u012f\u0130\t\4\2\2\u0130\u0134\7\60\2\2\u0131\u0133\t\4\2\2"+
		"\u0132\u0131\3\2\2\2\u0133\u0136\3\2\2\2\u0134\u0132\3\2\2\2\u0134\u0135"+
		"\3\2\2\2\u0135\u0138\3\2\2\2\u0136\u0134\3\2\2\2\u0137\u0139\t\2\2\2\u0138"+
		"\u0137\3\2\2\2\u0139\u013a\3\2\2\2\u013a\u0138\3\2\2\2\u013a\u013b\3\2"+
		"\2\2\u013bd\3\2\2\2\u013c\u013e\t\b\2\2\u013d\u013c\3\2\2\2\u013e\u013f"+
		"\3\2\2\2\u013f\u013d\3\2\2\2\u013f\u0140\3\2\2\2\u0140\u0141\3\2\2\2\u0141"+
		"\u0142\b\63\2\2\u0142f\3\2\2\2\r\2\u010b\u0111\u011c\u0121\u0127\u012a"+
		"\u012d\u0134\u013a\u013f\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}