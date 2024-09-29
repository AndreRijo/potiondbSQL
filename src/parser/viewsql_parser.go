// Code generated from ViewSQL.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // ViewSQL
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type ViewSQLParser struct {
	*antlr.BaseParser
}

var viewsqlParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func viewsqlParserInit() {
	staticData := &viewsqlParserStaticData
	staticData.literalNames = []string{
		"", "'+'", "'-'", "'*'", "'/'", "'SUM'", "'AVG'", "'MAX'", "'MIN'",
		"'COUNT'", "'AND'", "'ASC'", "'DESC'", "'='", "'>'", "'>='", "'<'",
		"'<='", "'!='", "'.'", "','", "':'", "'('", "')'", "'['", "']'", "'\"'",
		"'{'", "'}'", "'AS'", "'CREATE'", "'VIEW'", "'SELECT'", "'FROM'", "'WHERE'",
		"'GROUP'", "'BY'", "'ORDER'", "'LIMIT'", "'WITH'", "'IN'", "'PRIMARY'",
		"'KEY'", "'TABLE'", "'INDEX'", "'UPDATE'", "'DELETE'", "'INSERT'", "'INTO'",
		"'DROP'", "'SET'", "'FOREIGN'", "'REFERENCES'", "'VALUES'", "'CHECK'",
		"'ON'", "'AW'", "'RW'", "'LWW'", "'DEFAULT'", "'integer'", "'counter'",
		"'boolean'", "'varchar'", "'date'",
	}
	staticData.symbolicNames = []string{
		"", "ADD", "SUB", "MULT", "DIV", "SUM", "AVG", "MAX", "MIN", "COUNT",
		"AND", "ASC", "DESC", "EQUAL", "HIGHER", "HIGHER_EQUAL", "LOWER", "LOWER_EQUAL",
		"NOT_EQUAL", "DOT", "SEPARATOR", "RANGE_SEP", "LEFT_P", "RIGHT_P", "LEFT_SP",
		"RIGHT_SP", "INV_COMMA", "LEFT_CURLY", "RIGHT_CURLY", "AS", "CREATE",
		"VIEW", "SELECT", "FROM", "WHERE", "GROUP", "BY", "ORDER", "LIMIT",
		"WITH", "IN", "PRIMARY", "KEY", "TABLE", "INDEX", "UPDATE", "DELETE",
		"INSERT", "INTO", "DROP", "SET", "FOREIGN", "REFERENCES", "VALUES",
		"CHECK", "ON", "AW", "RW", "LWW", "DEFAULT", "INTEGER", "COUNTER", "BOOLEAN",
		"VARCHAR", "DATE_TYPE", "DATE", "BOOL", "STRING", "INT", "FLOAT", "WHITESPACE",
	}
	staticData.ruleNames = []string{
		"name", "constant", "aggrFunc", "field", "parameter", "nameable", "key",
		"math", "asClause", "aggregation", "count", "calc", "comp", "condition",
		"sortOrder", "continuousRange", "sparseRange", "range", "create", "with",
		"select", "from", "where", "groupby", "orderby", "limit", "view", "check",
		"foreignkey", "primarykey", "constraint", "columns", "createtable",
		"createindex", "drop", "delete", "set", "update", "values", "columnNames",
		"insert", "query", "statement", "start",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 70, 440, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 5, 1, 98, 8, 1, 10, 1, 12, 1, 101, 9, 1, 1, 1, 3, 1, 104, 8, 1, 1,
		2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1,
		5, 3, 5, 119, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1,
		7, 3, 7, 130, 8, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 138, 8, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 146, 8, 7, 10, 7, 12, 7, 149,
		9, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 166, 8, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 178, 8, 11, 1, 12,
		1, 12, 3, 12, 182, 8, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16,
		5, 16, 201, 8, 16, 10, 16, 12, 16, 204, 9, 16, 1, 16, 1, 16, 1, 17, 1,
		17, 3, 17, 210, 8, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18,
		1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 5, 19, 228,
		8, 19, 10, 19, 12, 19, 231, 9, 19, 1, 20, 1, 20, 1, 20, 1, 20, 5, 20, 237,
		8, 20, 10, 20, 12, 20, 240, 9, 20, 1, 21, 1, 21, 1, 21, 1, 21, 5, 21, 246,
		8, 21, 10, 21, 12, 21, 249, 9, 21, 1, 22, 1, 22, 1, 22, 1, 22, 5, 22, 255,
		8, 22, 10, 22, 12, 22, 258, 9, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 5,
		23, 265, 8, 23, 10, 23, 12, 23, 268, 9, 23, 1, 24, 1, 24, 1, 24, 1, 24,
		1, 24, 1, 24, 1, 24, 1, 24, 5, 24, 278, 8, 24, 10, 24, 12, 24, 281, 9,
		24, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 3, 26, 288, 8, 26, 1, 26, 1, 26,
		1, 26, 1, 26, 3, 26, 294, 8, 26, 1, 26, 3, 26, 297, 8, 26, 1, 26, 1, 26,
		3, 26, 301, 8, 26, 1, 26, 3, 26, 304, 8, 26, 1, 27, 1, 27, 1, 27, 1, 27,
		1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1,
		29, 1, 30, 1, 30, 1, 30, 3, 30, 324, 8, 30, 1, 31, 1, 31, 1, 31, 1, 31,
		3, 31, 330, 8, 31, 1, 31, 3, 31, 333, 8, 31, 1, 32, 1, 32, 1, 32, 1, 32,
		1, 32, 1, 32, 1, 32, 1, 32, 5, 32, 343, 8, 32, 10, 32, 12, 32, 346, 9,
		32, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33,
		1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1, 35, 1, 35, 1, 35, 1, 35, 1,
		35, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 5, 36, 375, 8, 36, 10, 36,
		12, 36, 378, 9, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 37, 1, 37, 1,
		37, 1, 37, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 5, 38, 394, 8, 38, 10, 38,
		12, 38, 397, 9, 38, 1, 38, 1, 38, 1, 38, 1, 39, 1, 39, 1, 39, 1, 39, 5,
		39, 406, 8, 39, 10, 39, 12, 39, 409, 9, 39, 1, 39, 1, 39, 1, 39, 1, 40,
		1, 40, 1, 40, 1, 40, 3, 40, 418, 8, 40, 1, 40, 1, 40, 1, 41, 1, 41, 1,
		41, 3, 41, 425, 8, 41, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42,
		1, 42, 3, 42, 435, 8, 42, 1, 43, 1, 43, 1, 43, 1, 43, 0, 1, 14, 44, 0,
		2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38,
		40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74,
		76, 78, 80, 82, 84, 86, 0, 10, 2, 0, 67, 67, 70, 70, 1, 0, 5, 8, 1, 0,
		3, 4, 1, 0, 1, 2, 1, 0, 13, 18, 1, 0, 11, 12, 1, 0, 24, 25, 1, 0, 14, 17,
		1, 0, 60, 64, 1, 0, 56, 58, 444, 0, 88, 1, 0, 0, 0, 2, 103, 1, 0, 0, 0,
		4, 105, 1, 0, 0, 0, 6, 107, 1, 0, 0, 0, 8, 111, 1, 0, 0, 0, 10, 118, 1,
		0, 0, 0, 12, 120, 1, 0, 0, 0, 14, 137, 1, 0, 0, 0, 16, 150, 1, 0, 0, 0,
		18, 154, 1, 0, 0, 0, 20, 161, 1, 0, 0, 0, 22, 177, 1, 0, 0, 0, 24, 181,
		1, 0, 0, 0, 26, 183, 1, 0, 0, 0, 28, 187, 1, 0, 0, 0, 30, 189, 1, 0, 0,
		0, 32, 196, 1, 0, 0, 0, 34, 209, 1, 0, 0, 0, 36, 211, 1, 0, 0, 0, 38, 219,
		1, 0, 0, 0, 40, 232, 1, 0, 0, 0, 42, 241, 1, 0, 0, 0, 44, 250, 1, 0, 0,
		0, 46, 259, 1, 0, 0, 0, 48, 269, 1, 0, 0, 0, 50, 282, 1, 0, 0, 0, 52, 285,
		1, 0, 0, 0, 54, 305, 1, 0, 0, 0, 56, 309, 1, 0, 0, 0, 58, 317, 1, 0, 0,
		0, 60, 323, 1, 0, 0, 0, 62, 325, 1, 0, 0, 0, 64, 334, 1, 0, 0, 0, 66, 350,
		1, 0, 0, 0, 68, 359, 1, 0, 0, 0, 70, 363, 1, 0, 0, 0, 72, 368, 1, 0, 0,
		0, 74, 384, 1, 0, 0, 0, 76, 388, 1, 0, 0, 0, 78, 401, 1, 0, 0, 0, 80, 413,
		1, 0, 0, 0, 82, 421, 1, 0, 0, 0, 84, 434, 1, 0, 0, 0, 86, 436, 1, 0, 0,
		0, 88, 89, 5, 67, 0, 0, 89, 1, 1, 0, 0, 0, 90, 104, 5, 65, 0, 0, 91, 104,
		5, 66, 0, 0, 92, 104, 5, 68, 0, 0, 93, 104, 5, 69, 0, 0, 94, 95, 5, 26,
		0, 0, 95, 99, 5, 67, 0, 0, 96, 98, 7, 0, 0, 0, 97, 96, 1, 0, 0, 0, 98,
		101, 1, 0, 0, 0, 99, 97, 1, 0, 0, 0, 99, 100, 1, 0, 0, 0, 100, 102, 1,
		0, 0, 0, 101, 99, 1, 0, 0, 0, 102, 104, 5, 26, 0, 0, 103, 90, 1, 0, 0,
		0, 103, 91, 1, 0, 0, 0, 103, 92, 1, 0, 0, 0, 103, 93, 1, 0, 0, 0, 103,
		94, 1, 0, 0, 0, 104, 3, 1, 0, 0, 0, 105, 106, 7, 1, 0, 0, 106, 5, 1, 0,
		0, 0, 107, 108, 3, 0, 0, 0, 108, 109, 5, 19, 0, 0, 109, 110, 3, 0, 0, 0,
		110, 7, 1, 0, 0, 0, 111, 112, 5, 24, 0, 0, 112, 113, 3, 0, 0, 0, 113, 114,
		5, 25, 0, 0, 114, 9, 1, 0, 0, 0, 115, 119, 3, 0, 0, 0, 116, 119, 3, 6,
		3, 0, 117, 119, 3, 8, 4, 0, 118, 115, 1, 0, 0, 0, 118, 116, 1, 0, 0, 0,
		118, 117, 1, 0, 0, 0, 119, 11, 1, 0, 0, 0, 120, 121, 5, 41, 0, 0, 121,
		122, 5, 42, 0, 0, 122, 123, 5, 22, 0, 0, 123, 124, 3, 10, 5, 0, 124, 125,
		5, 23, 0, 0, 125, 13, 1, 0, 0, 0, 126, 129, 6, 7, -1, 0, 127, 130, 3, 10,
		5, 0, 128, 130, 3, 2, 1, 0, 129, 127, 1, 0, 0, 0, 129, 128, 1, 0, 0, 0,
		130, 138, 1, 0, 0, 0, 131, 132, 5, 2, 0, 0, 132, 138, 3, 14, 7, 4, 133,
		134, 5, 22, 0, 0, 134, 135, 3, 14, 7, 0, 135, 136, 5, 23, 0, 0, 136, 138,
		1, 0, 0, 0, 137, 126, 1, 0, 0, 0, 137, 131, 1, 0, 0, 0, 137, 133, 1, 0,
		0, 0, 138, 147, 1, 0, 0, 0, 139, 140, 10, 2, 0, 0, 140, 141, 7, 2, 0, 0,
		141, 146, 3, 14, 7, 3, 142, 143, 10, 1, 0, 0, 143, 144, 7, 3, 0, 0, 144,
		146, 3, 14, 7, 2, 145, 139, 1, 0, 0, 0, 145, 142, 1, 0, 0, 0, 146, 149,
		1, 0, 0, 0, 147, 145, 1, 0, 0, 0, 147, 148, 1, 0, 0, 0, 148, 15, 1, 0,
		0, 0, 149, 147, 1, 0, 0, 0, 150, 151, 3, 14, 7, 0, 151, 152, 5, 29, 0,
		0, 152, 153, 3, 0, 0, 0, 153, 17, 1, 0, 0, 0, 154, 155, 3, 4, 2, 0, 155,
		156, 5, 22, 0, 0, 156, 157, 3, 14, 7, 0, 157, 158, 5, 23, 0, 0, 158, 159,
		5, 29, 0, 0, 159, 160, 3, 0, 0, 0, 160, 19, 1, 0, 0, 0, 161, 162, 5, 9,
		0, 0, 162, 165, 5, 22, 0, 0, 163, 166, 3, 10, 5, 0, 164, 166, 5, 3, 0,
		0, 165, 163, 1, 0, 0, 0, 165, 164, 1, 0, 0, 0, 166, 167, 1, 0, 0, 0, 167,
		168, 5, 23, 0, 0, 168, 169, 5, 29, 0, 0, 169, 170, 3, 0, 0, 0, 170, 21,
		1, 0, 0, 0, 171, 178, 3, 12, 6, 0, 172, 178, 3, 10, 5, 0, 173, 178, 3,
		16, 8, 0, 174, 178, 3, 18, 9, 0, 175, 178, 3, 20, 10, 0, 176, 178, 5, 3,
		0, 0, 177, 171, 1, 0, 0, 0, 177, 172, 1, 0, 0, 0, 177, 173, 1, 0, 0, 0,
		177, 174, 1, 0, 0, 0, 177, 175, 1, 0, 0, 0, 177, 176, 1, 0, 0, 0, 178,
		23, 1, 0, 0, 0, 179, 182, 7, 4, 0, 0, 180, 182, 3, 0, 0, 0, 181, 179, 1,
		0, 0, 0, 181, 180, 1, 0, 0, 0, 182, 25, 1, 0, 0, 0, 183, 184, 3, 10, 5,
		0, 184, 185, 3, 24, 12, 0, 185, 186, 3, 14, 7, 0, 186, 27, 1, 0, 0, 0,
		187, 188, 7, 5, 0, 0, 188, 29, 1, 0, 0, 0, 189, 190, 7, 6, 0, 0, 190, 191,
		3, 2, 1, 0, 191, 192, 5, 21, 0, 0, 192, 193, 3, 2, 1, 0, 193, 194, 1, 0,
		0, 0, 194, 195, 7, 6, 0, 0, 195, 31, 1, 0, 0, 0, 196, 197, 5, 27, 0, 0,
		197, 202, 3, 2, 1, 0, 198, 199, 5, 20, 0, 0, 199, 201, 3, 2, 1, 0, 200,
		198, 1, 0, 0, 0, 201, 204, 1, 0, 0, 0, 202, 200, 1, 0, 0, 0, 202, 203,
		1, 0, 0, 0, 203, 205, 1, 0, 0, 0, 204, 202, 1, 0, 0, 0, 205, 206, 5, 28,
		0, 0, 206, 33, 1, 0, 0, 0, 207, 210, 3, 30, 15, 0, 208, 210, 3, 32, 16,
		0, 209, 207, 1, 0, 0, 0, 209, 208, 1, 0, 0, 0, 210, 35, 1, 0, 0, 0, 211,
		212, 5, 30, 0, 0, 212, 213, 5, 31, 0, 0, 213, 214, 5, 22, 0, 0, 214, 215,
		5, 67, 0, 0, 215, 216, 5, 20, 0, 0, 216, 217, 5, 67, 0, 0, 217, 218, 5,
		23, 0, 0, 218, 37, 1, 0, 0, 0, 219, 220, 5, 39, 0, 0, 220, 221, 5, 67,
		0, 0, 221, 222, 5, 40, 0, 0, 222, 229, 3, 34, 17, 0, 223, 224, 5, 20, 0,
		0, 224, 225, 5, 67, 0, 0, 225, 226, 5, 40, 0, 0, 226, 228, 3, 34, 17, 0,
		227, 223, 1, 0, 0, 0, 228, 231, 1, 0, 0, 0, 229, 227, 1, 0, 0, 0, 229,
		230, 1, 0, 0, 0, 230, 39, 1, 0, 0, 0, 231, 229, 1, 0, 0, 0, 232, 233, 5,
		32, 0, 0, 233, 238, 3, 22, 11, 0, 234, 235, 5, 20, 0, 0, 235, 237, 3, 22,
		11, 0, 236, 234, 1, 0, 0, 0, 237, 240, 1, 0, 0, 0, 238, 236, 1, 0, 0, 0,
		238, 239, 1, 0, 0, 0, 239, 41, 1, 0, 0, 0, 240, 238, 1, 0, 0, 0, 241, 242,
		5, 33, 0, 0, 242, 247, 3, 0, 0, 0, 243, 244, 5, 20, 0, 0, 244, 246, 3,
		0, 0, 0, 245, 243, 1, 0, 0, 0, 246, 249, 1, 0, 0, 0, 247, 245, 1, 0, 0,
		0, 247, 248, 1, 0, 0, 0, 248, 43, 1, 0, 0, 0, 249, 247, 1, 0, 0, 0, 250,
		251, 5, 34, 0, 0, 251, 256, 3, 26, 13, 0, 252, 253, 5, 10, 0, 0, 253, 255,
		3, 26, 13, 0, 254, 252, 1, 0, 0, 0, 255, 258, 1, 0, 0, 0, 256, 254, 1,
		0, 0, 0, 256, 257, 1, 0, 0, 0, 257, 45, 1, 0, 0, 0, 258, 256, 1, 0, 0,
		0, 259, 260, 5, 35, 0, 0, 260, 261, 5, 36, 0, 0, 261, 266, 3, 10, 5, 0,
		262, 263, 5, 20, 0, 0, 263, 265, 3, 10, 5, 0, 264, 262, 1, 0, 0, 0, 265,
		268, 1, 0, 0, 0, 266, 264, 1, 0, 0, 0, 266, 267, 1, 0, 0, 0, 267, 47, 1,
		0, 0, 0, 268, 266, 1, 0, 0, 0, 269, 270, 5, 37, 0, 0, 270, 271, 5, 36,
		0, 0, 271, 272, 3, 10, 5, 0, 272, 279, 3, 28, 14, 0, 273, 274, 5, 20, 0,
		0, 274, 275, 3, 10, 5, 0, 275, 276, 3, 28, 14, 0, 276, 278, 1, 0, 0, 0,
		277, 273, 1, 0, 0, 0, 278, 281, 1, 0, 0, 0, 279, 277, 1, 0, 0, 0, 279,
		280, 1, 0, 0, 0, 280, 49, 1, 0, 0, 0, 281, 279, 1, 0, 0, 0, 282, 283, 5,
		38, 0, 0, 283, 284, 5, 68, 0, 0, 284, 51, 1, 0, 0, 0, 285, 287, 3, 36,
		18, 0, 286, 288, 3, 38, 19, 0, 287, 286, 1, 0, 0, 0, 287, 288, 1, 0, 0,
		0, 288, 289, 1, 0, 0, 0, 289, 290, 5, 29, 0, 0, 290, 291, 3, 40, 20, 0,
		291, 293, 3, 42, 21, 0, 292, 294, 3, 44, 22, 0, 293, 292, 1, 0, 0, 0, 293,
		294, 1, 0, 0, 0, 294, 296, 1, 0, 0, 0, 295, 297, 3, 46, 23, 0, 296, 295,
		1, 0, 0, 0, 296, 297, 1, 0, 0, 0, 297, 298, 1, 0, 0, 0, 298, 300, 3, 46,
		23, 0, 299, 301, 3, 48, 24, 0, 300, 299, 1, 0, 0, 0, 300, 301, 1, 0, 0,
		0, 301, 303, 1, 0, 0, 0, 302, 304, 3, 50, 25, 0, 303, 302, 1, 0, 0, 0,
		303, 304, 1, 0, 0, 0, 304, 53, 1, 0, 0, 0, 305, 306, 5, 54, 0, 0, 306,
		307, 7, 7, 0, 0, 307, 308, 3, 2, 1, 0, 308, 55, 1, 0, 0, 0, 309, 310, 5,
		51, 0, 0, 310, 311, 5, 42, 0, 0, 311, 312, 5, 52, 0, 0, 312, 313, 3, 0,
		0, 0, 313, 314, 5, 22, 0, 0, 314, 315, 3, 0, 0, 0, 315, 316, 5, 23, 0,
		0, 316, 57, 1, 0, 0, 0, 317, 318, 5, 41, 0, 0, 318, 319, 5, 42, 0, 0, 319,
		59, 1, 0, 0, 0, 320, 324, 3, 58, 29, 0, 321, 324, 3, 56, 28, 0, 322, 324,
		3, 54, 27, 0, 323, 320, 1, 0, 0, 0, 323, 321, 1, 0, 0, 0, 323, 322, 1,
		0, 0, 0, 324, 61, 1, 0, 0, 0, 325, 326, 5, 67, 0, 0, 326, 329, 7, 8, 0,
		0, 327, 328, 5, 59, 0, 0, 328, 330, 3, 2, 1, 0, 329, 327, 1, 0, 0, 0, 329,
		330, 1, 0, 0, 0, 330, 332, 1, 0, 0, 0, 331, 333, 3, 60, 30, 0, 332, 331,
		1, 0, 0, 0, 332, 333, 1, 0, 0, 0, 333, 63, 1, 0, 0, 0, 334, 335, 5, 30,
		0, 0, 335, 336, 7, 9, 0, 0, 336, 337, 5, 43, 0, 0, 337, 338, 3, 0, 0, 0,
		338, 344, 5, 22, 0, 0, 339, 340, 3, 62, 31, 0, 340, 341, 5, 20, 0, 0, 341,
		343, 1, 0, 0, 0, 342, 339, 1, 0, 0, 0, 343, 346, 1, 0, 0, 0, 344, 342,
		1, 0, 0, 0, 344, 345, 1, 0, 0, 0, 345, 347, 1, 0, 0, 0, 346, 344, 1, 0,
		0, 0, 347, 348, 3, 62, 31, 0, 348, 349, 5, 23, 0, 0, 349, 65, 1, 0, 0,
		0, 350, 351, 5, 30, 0, 0, 351, 352, 5, 44, 0, 0, 352, 353, 3, 0, 0, 0,
		353, 354, 5, 55, 0, 0, 354, 355, 3, 0, 0, 0, 355, 356, 5, 22, 0, 0, 356,
		357, 3, 0, 0, 0, 357, 358, 5, 23, 0, 0, 358, 67, 1, 0, 0, 0, 359, 360,
		5, 49, 0, 0, 360, 361, 5, 43, 0, 0, 361, 362, 3, 0, 0, 0, 362, 69, 1, 0,
		0, 0, 363, 364, 5, 46, 0, 0, 364, 365, 5, 33, 0, 0, 365, 366, 3, 0, 0,
		0, 366, 367, 3, 44, 22, 0, 367, 71, 1, 0, 0, 0, 368, 376, 5, 50, 0, 0,
		369, 370, 3, 0, 0, 0, 370, 371, 5, 13, 0, 0, 371, 372, 3, 2, 1, 0, 372,
		373, 5, 20, 0, 0, 373, 375, 1, 0, 0, 0, 374, 369, 1, 0, 0, 0, 375, 378,
		1, 0, 0, 0, 376, 374, 1, 0, 0, 0, 376, 377, 1, 0, 0, 0, 377, 379, 1, 0,
		0, 0, 378, 376, 1, 0, 0, 0, 379, 380, 3, 0, 0, 0, 380, 381, 5, 13, 0, 0,
		381, 382, 3, 2, 1, 0, 382, 383, 3, 44, 22, 0, 383, 73, 1, 0, 0, 0, 384,
		385, 5, 45, 0, 0, 385, 386, 3, 0, 0, 0, 386, 387, 3, 72, 36, 0, 387, 75,
		1, 0, 0, 0, 388, 389, 5, 53, 0, 0, 389, 395, 5, 22, 0, 0, 390, 391, 3,
		2, 1, 0, 391, 392, 5, 20, 0, 0, 392, 394, 1, 0, 0, 0, 393, 390, 1, 0, 0,
		0, 394, 397, 1, 0, 0, 0, 395, 393, 1, 0, 0, 0, 395, 396, 1, 0, 0, 0, 396,
		398, 1, 0, 0, 0, 397, 395, 1, 0, 0, 0, 398, 399, 3, 2, 1, 0, 399, 400,
		5, 23, 0, 0, 400, 77, 1, 0, 0, 0, 401, 407, 5, 22, 0, 0, 402, 403, 3, 0,
		0, 0, 403, 404, 5, 20, 0, 0, 404, 406, 1, 0, 0, 0, 405, 402, 1, 0, 0, 0,
		406, 409, 1, 0, 0, 0, 407, 405, 1, 0, 0, 0, 407, 408, 1, 0, 0, 0, 408,
		410, 1, 0, 0, 0, 409, 407, 1, 0, 0, 0, 410, 411, 3, 0, 0, 0, 411, 412,
		5, 23, 0, 0, 412, 79, 1, 0, 0, 0, 413, 414, 5, 47, 0, 0, 414, 415, 5, 48,
		0, 0, 415, 417, 3, 0, 0, 0, 416, 418, 3, 78, 39, 0, 417, 416, 1, 0, 0,
		0, 417, 418, 1, 0, 0, 0, 418, 419, 1, 0, 0, 0, 419, 420, 3, 76, 38, 0,
		420, 81, 1, 0, 0, 0, 421, 422, 3, 40, 20, 0, 422, 424, 3, 42, 21, 0, 423,
		425, 3, 44, 22, 0, 424, 423, 1, 0, 0, 0, 424, 425, 1, 0, 0, 0, 425, 83,
		1, 0, 0, 0, 426, 435, 3, 52, 26, 0, 427, 435, 3, 64, 32, 0, 428, 435, 3,
		66, 33, 0, 429, 435, 3, 80, 40, 0, 430, 435, 3, 74, 37, 0, 431, 435, 3,
		70, 35, 0, 432, 435, 3, 68, 34, 0, 433, 435, 3, 82, 41, 0, 434, 426, 1,
		0, 0, 0, 434, 427, 1, 0, 0, 0, 434, 428, 1, 0, 0, 0, 434, 429, 1, 0, 0,
		0, 434, 430, 1, 0, 0, 0, 434, 431, 1, 0, 0, 0, 434, 432, 1, 0, 0, 0, 434,
		433, 1, 0, 0, 0, 435, 85, 1, 0, 0, 0, 436, 437, 3, 84, 42, 0, 437, 438,
		5, 0, 0, 1, 438, 87, 1, 0, 0, 0, 33, 99, 103, 118, 129, 137, 145, 147,
		165, 177, 181, 202, 209, 229, 238, 247, 256, 266, 279, 287, 293, 296, 300,
		303, 323, 329, 332, 344, 376, 395, 407, 417, 424, 434,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// ViewSQLParserInit initializes any static state used to implement ViewSQLParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewViewSQLParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ViewSQLParserInit() {
	staticData := &viewsqlParserStaticData
	staticData.once.Do(viewsqlParserInit)
}

// NewViewSQLParser produces a new parser instance for the optional input antlr.TokenStream.
func NewViewSQLParser(input antlr.TokenStream) *ViewSQLParser {
	ViewSQLParserInit()
	this := new(ViewSQLParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &viewsqlParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "ViewSQL.g4"

	return this
}

// ViewSQLParser tokens.
const (
	ViewSQLParserEOF          = antlr.TokenEOF
	ViewSQLParserADD          = 1
	ViewSQLParserSUB          = 2
	ViewSQLParserMULT         = 3
	ViewSQLParserDIV          = 4
	ViewSQLParserSUM          = 5
	ViewSQLParserAVG          = 6
	ViewSQLParserMAX          = 7
	ViewSQLParserMIN          = 8
	ViewSQLParserCOUNT        = 9
	ViewSQLParserAND          = 10
	ViewSQLParserASC          = 11
	ViewSQLParserDESC         = 12
	ViewSQLParserEQUAL        = 13
	ViewSQLParserHIGHER       = 14
	ViewSQLParserHIGHER_EQUAL = 15
	ViewSQLParserLOWER        = 16
	ViewSQLParserLOWER_EQUAL  = 17
	ViewSQLParserNOT_EQUAL    = 18
	ViewSQLParserDOT          = 19
	ViewSQLParserSEPARATOR    = 20
	ViewSQLParserRANGE_SEP    = 21
	ViewSQLParserLEFT_P       = 22
	ViewSQLParserRIGHT_P      = 23
	ViewSQLParserLEFT_SP      = 24
	ViewSQLParserRIGHT_SP     = 25
	ViewSQLParserINV_COMMA    = 26
	ViewSQLParserLEFT_CURLY   = 27
	ViewSQLParserRIGHT_CURLY  = 28
	ViewSQLParserAS           = 29
	ViewSQLParserCREATE       = 30
	ViewSQLParserVIEW         = 31
	ViewSQLParserSELECT       = 32
	ViewSQLParserFROM         = 33
	ViewSQLParserWHERE        = 34
	ViewSQLParserGROUP        = 35
	ViewSQLParserBY           = 36
	ViewSQLParserORDER        = 37
	ViewSQLParserLIMIT        = 38
	ViewSQLParserWITH         = 39
	ViewSQLParserIN           = 40
	ViewSQLParserPRIMARY      = 41
	ViewSQLParserKEY          = 42
	ViewSQLParserTABLE        = 43
	ViewSQLParserINDEX        = 44
	ViewSQLParserUPDATE       = 45
	ViewSQLParserDELETE       = 46
	ViewSQLParserINSERT       = 47
	ViewSQLParserINTO         = 48
	ViewSQLParserDROP         = 49
	ViewSQLParserSET          = 50
	ViewSQLParserFOREIGN      = 51
	ViewSQLParserREFERENCES   = 52
	ViewSQLParserVALUES       = 53
	ViewSQLParserCHECK        = 54
	ViewSQLParserON           = 55
	ViewSQLParserAW           = 56
	ViewSQLParserRW           = 57
	ViewSQLParserLWW          = 58
	ViewSQLParserDEFAULT      = 59
	ViewSQLParserINTEGER      = 60
	ViewSQLParserCOUNTER      = 61
	ViewSQLParserBOOLEAN      = 62
	ViewSQLParserVARCHAR      = 63
	ViewSQLParserDATE_TYPE    = 64
	ViewSQLParserDATE         = 65
	ViewSQLParserBOOL         = 66
	ViewSQLParserSTRING       = 67
	ViewSQLParserINT          = 68
	ViewSQLParserFLOAT        = 69
	ViewSQLParserWHITESPACE   = 70
)

// ViewSQLParser rules.
const (
	ViewSQLParserRULE_name            = 0
	ViewSQLParserRULE_constant        = 1
	ViewSQLParserRULE_aggrFunc        = 2
	ViewSQLParserRULE_field           = 3
	ViewSQLParserRULE_parameter       = 4
	ViewSQLParserRULE_nameable        = 5
	ViewSQLParserRULE_key             = 6
	ViewSQLParserRULE_math            = 7
	ViewSQLParserRULE_asClause        = 8
	ViewSQLParserRULE_aggregation     = 9
	ViewSQLParserRULE_count           = 10
	ViewSQLParserRULE_calc            = 11
	ViewSQLParserRULE_comp            = 12
	ViewSQLParserRULE_condition       = 13
	ViewSQLParserRULE_sortOrder       = 14
	ViewSQLParserRULE_continuousRange = 15
	ViewSQLParserRULE_sparseRange     = 16
	ViewSQLParserRULE_range           = 17
	ViewSQLParserRULE_create          = 18
	ViewSQLParserRULE_with            = 19
	ViewSQLParserRULE_select          = 20
	ViewSQLParserRULE_from            = 21
	ViewSQLParserRULE_where           = 22
	ViewSQLParserRULE_groupby         = 23
	ViewSQLParserRULE_orderby         = 24
	ViewSQLParserRULE_limit           = 25
	ViewSQLParserRULE_view            = 26
	ViewSQLParserRULE_check           = 27
	ViewSQLParserRULE_foreignkey      = 28
	ViewSQLParserRULE_primarykey      = 29
	ViewSQLParserRULE_constraint      = 30
	ViewSQLParserRULE_columns         = 31
	ViewSQLParserRULE_createtable     = 32
	ViewSQLParserRULE_createindex     = 33
	ViewSQLParserRULE_drop            = 34
	ViewSQLParserRULE_delete          = 35
	ViewSQLParserRULE_set             = 36
	ViewSQLParserRULE_update          = 37
	ViewSQLParserRULE_values          = 38
	ViewSQLParserRULE_columnNames     = 39
	ViewSQLParserRULE_insert          = 40
	ViewSQLParserRULE_query           = 41
	ViewSQLParserRULE_statement       = 42
	ViewSQLParserRULE_start           = 43
)

// INameContext is an interface to support dynamic dispatch.
type INameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode

	// IsNameContext differentiates from other interfaces.
	IsNameContext()
}

type NameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNameContext() *NameContext {
	var p = new(NameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ViewSQLParserRULE_name
	return p
}

func (*NameContext) IsNameContext() {}

func NewNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameContext {
	var p = new(NameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ViewSQLParserRULE_name

	return p
}

func (s *NameContext) GetParser() antlr.Parser { return s.parser }

func (s *NameContext) STRING() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserSTRING, 0)
}

func (s *NameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.EnterName(s)
	}
}

func (s *NameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.ExitName(s)
	}
}

func (p *ViewSQLParser) Name() (localctx INameContext) {
	this := p
	_ = this

	localctx = NewNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ViewSQLParserRULE_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(88)
		p.Match(ViewSQLParserSTRING)
	}

	return localctx
}

// IConstantContext is an interface to support dynamic dispatch.
type IConstantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DATE() antlr.TerminalNode
	BOOL() antlr.TerminalNode
	INT() antlr.TerminalNode
	FLOAT() antlr.TerminalNode
	AllINV_COMMA() []antlr.TerminalNode
	INV_COMMA(i int) antlr.TerminalNode
	AllSTRING() []antlr.TerminalNode
	STRING(i int) antlr.TerminalNode
	AllWHITESPACE() []antlr.TerminalNode
	WHITESPACE(i int) antlr.TerminalNode

	// IsConstantContext differentiates from other interfaces.
	IsConstantContext()
}

type ConstantContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantContext() *ConstantContext {
	var p = new(ConstantContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ViewSQLParserRULE_constant
	return p
}

func (*ConstantContext) IsConstantContext() {}

func NewConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantContext {
	var p = new(ConstantContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ViewSQLParserRULE_constant

	return p
}

func (s *ConstantContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantContext) DATE() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserDATE, 0)
}

func (s *ConstantContext) BOOL() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserBOOL, 0)
}

func (s *ConstantContext) INT() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserINT, 0)
}

func (s *ConstantContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserFLOAT, 0)
}

func (s *ConstantContext) AllINV_COMMA() []antlr.TerminalNode {
	return s.GetTokens(ViewSQLParserINV_COMMA)
}

func (s *ConstantContext) INV_COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ViewSQLParserINV_COMMA, i)
}

func (s *ConstantContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(ViewSQLParserSTRING)
}

func (s *ConstantContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(ViewSQLParserSTRING, i)
}

func (s *ConstantContext) AllWHITESPACE() []antlr.TerminalNode {
	return s.GetTokens(ViewSQLParserWHITESPACE)
}

func (s *ConstantContext) WHITESPACE(i int) antlr.TerminalNode {
	return s.GetToken(ViewSQLParserWHITESPACE, i)
}

func (s *ConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.EnterConstant(s)
	}
}

func (s *ConstantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.ExitConstant(s)
	}
}

func (p *ViewSQLParser) Constant() (localctx IConstantContext) {
	this := p
	_ = this

	localctx = NewConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ViewSQLParserRULE_constant)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(103)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ViewSQLParserDATE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(90)
			p.Match(ViewSQLParserDATE)
		}

	case ViewSQLParserBOOL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(91)
			p.Match(ViewSQLParserBOOL)
		}

	case ViewSQLParserINT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(92)
			p.Match(ViewSQLParserINT)
		}

	case ViewSQLParserFLOAT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(93)
			p.Match(ViewSQLParserFLOAT)
		}

	case ViewSQLParserINV_COMMA:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(94)
			p.Match(ViewSQLParserINV_COMMA)
		}
		{
			p.SetState(95)
			p.Match(ViewSQLParserSTRING)
		}
		p.SetState(99)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ViewSQLParserSTRING || _la == ViewSQLParserWHITESPACE {
			{
				p.SetState(96)
				_la = p.GetTokenStream().LA(1)

				if !(_la == ViewSQLParserSTRING || _la == ViewSQLParserWHITESPACE) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

			p.SetState(101)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(102)
			p.Match(ViewSQLParserINV_COMMA)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAggrFuncContext is an interface to support dynamic dispatch.
type IAggrFuncContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SUM() antlr.TerminalNode
	AVG() antlr.TerminalNode
	MAX() antlr.TerminalNode
	MIN() antlr.TerminalNode

	// IsAggrFuncContext differentiates from other interfaces.
	IsAggrFuncContext()
}

type AggrFuncContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAggrFuncContext() *AggrFuncContext {
	var p = new(AggrFuncContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ViewSQLParserRULE_aggrFunc
	return p
}

func (*AggrFuncContext) IsAggrFuncContext() {}

func NewAggrFuncContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AggrFuncContext {
	var p = new(AggrFuncContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ViewSQLParserRULE_aggrFunc

	return p
}

func (s *AggrFuncContext) GetParser() antlr.Parser { return s.parser }

func (s *AggrFuncContext) SUM() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserSUM, 0)
}

func (s *AggrFuncContext) AVG() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserAVG, 0)
}

func (s *AggrFuncContext) MAX() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserMAX, 0)
}

func (s *AggrFuncContext) MIN() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserMIN, 0)
}

func (s *AggrFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AggrFuncContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AggrFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.EnterAggrFunc(s)
	}
}

func (s *AggrFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.ExitAggrFunc(s)
	}
}

func (p *ViewSQLParser) AggrFunc() (localctx IAggrFuncContext) {
	this := p
	_ = this

	localctx = NewAggrFuncContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ViewSQLParserRULE_aggrFunc)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(105)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&480) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllName() []INameContext
	Name(i int) INameContext
	DOT() antlr.TerminalNode

	// IsFieldContext differentiates from other interfaces.
	IsFieldContext()
}

type FieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldContext() *FieldContext {
	var p = new(FieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ViewSQLParserRULE_field
	return p
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ViewSQLParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) AllName() []INameContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INameContext); ok {
			len++
		}
	}

	tst := make([]INameContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INameContext); ok {
			tst[i] = t.(INameContext)
			i++
		}
	}

	return tst
}

func (s *FieldContext) Name(i int) INameContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *FieldContext) DOT() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserDOT, 0)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.EnterField(s)
	}
}

func (s *FieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.ExitField(s)
	}
}

func (p *ViewSQLParser) Field() (localctx IFieldContext) {
	this := p
	_ = this

	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ViewSQLParserRULE_field)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(107)
		p.Name()
	}
	{
		p.SetState(108)
		p.Match(ViewSQLParserDOT)
	}
	{
		p.SetState(109)
		p.Name()
	}

	return localctx
}

// IParameterContext is an interface to support dynamic dispatch.
type IParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LEFT_SP() antlr.TerminalNode
	Name() INameContext
	RIGHT_SP() antlr.TerminalNode

	// IsParameterContext differentiates from other interfaces.
	IsParameterContext()
}

type ParameterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterContext() *ParameterContext {
	var p = new(ParameterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ViewSQLParserRULE_parameter
	return p
}

func (*ParameterContext) IsParameterContext() {}

func NewParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterContext {
	var p = new(ParameterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ViewSQLParserRULE_parameter

	return p
}

func (s *ParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterContext) LEFT_SP() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserLEFT_SP, 0)
}

func (s *ParameterContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *ParameterContext) RIGHT_SP() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserRIGHT_SP, 0)
}

func (s *ParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.EnterParameter(s)
	}
}

func (s *ParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.ExitParameter(s)
	}
}

func (p *ViewSQLParser) Parameter() (localctx IParameterContext) {
	this := p
	_ = this

	localctx = NewParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ViewSQLParserRULE_parameter)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(111)
		p.Match(ViewSQLParserLEFT_SP)
	}
	{
		p.SetState(112)
		p.Name()
	}
	{
		p.SetState(113)
		p.Match(ViewSQLParserRIGHT_SP)
	}

	return localctx
}

// INameableContext is an interface to support dynamic dispatch.
type INameableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Name() INameContext
	Field() IFieldContext
	Parameter() IParameterContext

	// IsNameableContext differentiates from other interfaces.
	IsNameableContext()
}

type NameableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNameableContext() *NameableContext {
	var p = new(NameableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ViewSQLParserRULE_nameable
	return p
}

func (*NameableContext) IsNameableContext() {}

func NewNameableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameableContext {
	var p = new(NameableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ViewSQLParserRULE_nameable

	return p
}

func (s *NameableContext) GetParser() antlr.Parser { return s.parser }

func (s *NameableContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *NameableContext) Field() IFieldContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *NameableContext) Parameter() IParameterContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParameterContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParameterContext)
}

func (s *NameableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NameableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.EnterNameable(s)
	}
}

func (s *NameableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.ExitNameable(s)
	}
}

func (p *ViewSQLParser) Nameable() (localctx INameableContext) {
	this := p
	_ = this

	localctx = NewNameableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ViewSQLParserRULE_nameable)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(118)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(115)
			p.Name()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(116)
			p.Field()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(117)
			p.Parameter()
		}

	}

	return localctx
}

// IKeyContext is an interface to support dynamic dispatch.
type IKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PRIMARY() antlr.TerminalNode
	KEY() antlr.TerminalNode
	LEFT_P() antlr.TerminalNode
	Nameable() INameableContext
	RIGHT_P() antlr.TerminalNode

	// IsKeyContext differentiates from other interfaces.
	IsKeyContext()
}

type KeyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeyContext() *KeyContext {
	var p = new(KeyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ViewSQLParserRULE_key
	return p
}

func (*KeyContext) IsKeyContext() {}

func NewKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyContext {
	var p = new(KeyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ViewSQLParserRULE_key

	return p
}

func (s *KeyContext) GetParser() antlr.Parser { return s.parser }

func (s *KeyContext) PRIMARY() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserPRIMARY, 0)
}

func (s *KeyContext) KEY() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserKEY, 0)
}

func (s *KeyContext) LEFT_P() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserLEFT_P, 0)
}

func (s *KeyContext) Nameable() INameableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameableContext)
}

func (s *KeyContext) RIGHT_P() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserRIGHT_P, 0)
}

func (s *KeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.EnterKey(s)
	}
}

func (s *KeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.ExitKey(s)
	}
}

func (p *ViewSQLParser) Key() (localctx IKeyContext) {
	this := p
	_ = this

	localctx = NewKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ViewSQLParserRULE_key)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(120)
		p.Match(ViewSQLParserPRIMARY)
	}
	{
		p.SetState(121)
		p.Match(ViewSQLParserKEY)
	}
	{
		p.SetState(122)
		p.Match(ViewSQLParserLEFT_P)
	}
	{
		p.SetState(123)
		p.Nameable()
	}
	{
		p.SetState(124)
		p.Match(ViewSQLParserRIGHT_P)
	}

	return localctx
}

// IMathContext is an interface to support dynamic dispatch.
type IMathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsMathContext differentiates from other interfaces.
	IsMathContext()
}

type MathContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMathContext() *MathContext {
	var p = new(MathContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ViewSQLParserRULE_math
	return p
}

func (*MathContext) IsMathContext() {}

func NewMathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MathContext {
	var p = new(MathContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ViewSQLParserRULE_math

	return p
}

func (s *MathContext) GetParser() antlr.Parser { return s.parser }

func (s *MathContext) CopyFrom(ctx *MathContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *MathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AddOrSubContext struct {
	*MathContext
	opType antlr.Token
}

func NewAddOrSubContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddOrSubContext {
	var p = new(AddOrSubContext)

	p.MathContext = NewEmptyMathContext()
	p.parser = parser
	p.CopyFrom(ctx.(*MathContext))

	return p
}

func (s *AddOrSubContext) GetOpType() antlr.Token { return s.opType }

func (s *AddOrSubContext) SetOpType(v antlr.Token) { s.opType = v }

func (s *AddOrSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddOrSubContext) AllMath() []IMathContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMathContext); ok {
			len++
		}
	}

	tst := make([]IMathContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMathContext); ok {
			tst[i] = t.(IMathContext)
			i++
		}
	}

	return tst
}

func (s *AddOrSubContext) Math(i int) IMathContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMathContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMathContext)
}

func (s *AddOrSubContext) ADD() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserADD, 0)
}

func (s *AddOrSubContext) SUB() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserSUB, 0)
}

func (s *AddOrSubContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.EnterAddOrSub(s)
	}
}

func (s *AddOrSubContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.ExitAddOrSub(s)
	}
}

type ValueContext struct {
	*MathContext
}

func NewValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValueContext {
	var p = new(ValueContext)

	p.MathContext = NewEmptyMathContext()
	p.parser = parser
	p.CopyFrom(ctx.(*MathContext))

	return p
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) Nameable() INameableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameableContext)
}

func (s *ValueContext) Constant() IConstantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.ExitValue(s)
	}
}

type MinusContext struct {
	*MathContext
}

func NewMinusContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MinusContext {
	var p = new(MinusContext)

	p.MathContext = NewEmptyMathContext()
	p.parser = parser
	p.CopyFrom(ctx.(*MathContext))

	return p
}

func (s *MinusContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MinusContext) SUB() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserSUB, 0)
}

func (s *MinusContext) Math() IMathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMathContext)
}

func (s *MinusContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.EnterMinus(s)
	}
}

func (s *MinusContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.ExitMinus(s)
	}
}

type ParenthesesContext struct {
	*MathContext
}

func NewParenthesesContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenthesesContext {
	var p = new(ParenthesesContext)

	p.MathContext = NewEmptyMathContext()
	p.parser = parser
	p.CopyFrom(ctx.(*MathContext))

	return p
}

func (s *ParenthesesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenthesesContext) LEFT_P() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserLEFT_P, 0)
}

func (s *ParenthesesContext) Math() IMathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMathContext)
}

func (s *ParenthesesContext) RIGHT_P() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserRIGHT_P, 0)
}

func (s *ParenthesesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.EnterParentheses(s)
	}
}

func (s *ParenthesesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.ExitParentheses(s)
	}
}

type MultOrDivContext struct {
	*MathContext
	opType antlr.Token
}

func NewMultOrDivContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MultOrDivContext {
	var p = new(MultOrDivContext)

	p.MathContext = NewEmptyMathContext()
	p.parser = parser
	p.CopyFrom(ctx.(*MathContext))

	return p
}

func (s *MultOrDivContext) GetOpType() antlr.Token { return s.opType }

func (s *MultOrDivContext) SetOpType(v antlr.Token) { s.opType = v }

func (s *MultOrDivContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultOrDivContext) AllMath() []IMathContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMathContext); ok {
			len++
		}
	}

	tst := make([]IMathContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMathContext); ok {
			tst[i] = t.(IMathContext)
			i++
		}
	}

	return tst
}

func (s *MultOrDivContext) Math(i int) IMathContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMathContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMathContext)
}

func (s *MultOrDivContext) MULT() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserMULT, 0)
}

func (s *MultOrDivContext) DIV() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserDIV, 0)
}

func (s *MultOrDivContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.EnterMultOrDiv(s)
	}
}

func (s *MultOrDivContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.ExitMultOrDiv(s)
	}
}

func (p *ViewSQLParser) Math() (localctx IMathContext) {
	return p.math(0)
}

func (p *ViewSQLParser) math(_p int) (localctx IMathContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewMathContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IMathContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 14
	p.EnterRecursionRule(localctx, 14, ViewSQLParserRULE_math, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(137)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ViewSQLParserLEFT_SP, ViewSQLParserINV_COMMA, ViewSQLParserDATE, ViewSQLParserBOOL, ViewSQLParserSTRING, ViewSQLParserINT, ViewSQLParserFLOAT:
		localctx = NewValueContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		p.SetState(129)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case ViewSQLParserLEFT_SP, ViewSQLParserSTRING:
			{
				p.SetState(127)
				p.Nameable()
			}

		case ViewSQLParserINV_COMMA, ViewSQLParserDATE, ViewSQLParserBOOL, ViewSQLParserINT, ViewSQLParserFLOAT:
			{
				p.SetState(128)
				p.Constant()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

	case ViewSQLParserSUB:
		localctx = NewMinusContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(131)
			p.Match(ViewSQLParserSUB)
		}
		{
			p.SetState(132)
			p.math(4)
		}

	case ViewSQLParserLEFT_P:
		localctx = NewParenthesesContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(133)
			p.Match(ViewSQLParserLEFT_P)
		}
		{
			p.SetState(134)
			p.math(0)
		}
		{
			p.SetState(135)
			p.Match(ViewSQLParserRIGHT_P)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(147)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(145)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMultOrDivContext(p, NewMathContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ViewSQLParserRULE_math)
				p.SetState(139)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(140)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MultOrDivContext).opType = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == ViewSQLParserMULT || _la == ViewSQLParserDIV) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MultOrDivContext).opType = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(141)
					p.math(3)
				}

			case 2:
				localctx = NewAddOrSubContext(p, NewMathContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ViewSQLParserRULE_math)
				p.SetState(142)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(143)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AddOrSubContext).opType = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == ViewSQLParserADD || _la == ViewSQLParserSUB) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AddOrSubContext).opType = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(144)
					p.math(2)
				}

			}

		}
		p.SetState(149)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
	}

	return localctx
}

// IAsClauseContext is an interface to support dynamic dispatch.
type IAsClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Math() IMathContext
	AS() antlr.TerminalNode
	Name() INameContext

	// IsAsClauseContext differentiates from other interfaces.
	IsAsClauseContext()
}

type AsClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAsClauseContext() *AsClauseContext {
	var p = new(AsClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ViewSQLParserRULE_asClause
	return p
}

func (*AsClauseContext) IsAsClauseContext() {}

func NewAsClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AsClauseContext {
	var p = new(AsClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ViewSQLParserRULE_asClause

	return p
}

func (s *AsClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *AsClauseContext) Math() IMathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMathContext)
}

func (s *AsClauseContext) AS() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserAS, 0)
}

func (s *AsClauseContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *AsClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AsClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.EnterAsClause(s)
	}
}

func (s *AsClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.ExitAsClause(s)
	}
}

func (p *ViewSQLParser) AsClause() (localctx IAsClauseContext) {
	this := p
	_ = this

	localctx = NewAsClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ViewSQLParserRULE_asClause)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(150)
		p.math(0)
	}
	{
		p.SetState(151)
		p.Match(ViewSQLParserAS)
	}
	{
		p.SetState(152)
		p.Name()
	}

	return localctx
}

// IAggregationContext is an interface to support dynamic dispatch.
type IAggregationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AggrFunc() IAggrFuncContext
	LEFT_P() antlr.TerminalNode
	Math() IMathContext
	RIGHT_P() antlr.TerminalNode
	AS() antlr.TerminalNode
	Name() INameContext

	// IsAggregationContext differentiates from other interfaces.
	IsAggregationContext()
}

type AggregationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAggregationContext() *AggregationContext {
	var p = new(AggregationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ViewSQLParserRULE_aggregation
	return p
}

func (*AggregationContext) IsAggregationContext() {}

func NewAggregationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AggregationContext {
	var p = new(AggregationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ViewSQLParserRULE_aggregation

	return p
}

func (s *AggregationContext) GetParser() antlr.Parser { return s.parser }

func (s *AggregationContext) AggrFunc() IAggrFuncContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAggrFuncContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAggrFuncContext)
}

func (s *AggregationContext) LEFT_P() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserLEFT_P, 0)
}

func (s *AggregationContext) Math() IMathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMathContext)
}

func (s *AggregationContext) RIGHT_P() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserRIGHT_P, 0)
}

func (s *AggregationContext) AS() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserAS, 0)
}

func (s *AggregationContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *AggregationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AggregationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AggregationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.EnterAggregation(s)
	}
}

func (s *AggregationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.ExitAggregation(s)
	}
}

func (p *ViewSQLParser) Aggregation() (localctx IAggregationContext) {
	this := p
	_ = this

	localctx = NewAggregationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ViewSQLParserRULE_aggregation)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(154)
		p.AggrFunc()
	}
	{
		p.SetState(155)
		p.Match(ViewSQLParserLEFT_P)
	}
	{
		p.SetState(156)
		p.math(0)
	}
	{
		p.SetState(157)
		p.Match(ViewSQLParserRIGHT_P)
	}
	{
		p.SetState(158)
		p.Match(ViewSQLParserAS)
	}
	{
		p.SetState(159)
		p.Name()
	}

	return localctx
}

// ICountContext is an interface to support dynamic dispatch.
type ICountContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COUNT() antlr.TerminalNode
	LEFT_P() antlr.TerminalNode
	RIGHT_P() antlr.TerminalNode
	AS() antlr.TerminalNode
	Name() INameContext
	Nameable() INameableContext
	MULT() antlr.TerminalNode

	// IsCountContext differentiates from other interfaces.
	IsCountContext()
}

type CountContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCountContext() *CountContext {
	var p = new(CountContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ViewSQLParserRULE_count
	return p
}

func (*CountContext) IsCountContext() {}

func NewCountContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CountContext {
	var p = new(CountContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ViewSQLParserRULE_count

	return p
}

func (s *CountContext) GetParser() antlr.Parser { return s.parser }

func (s *CountContext) COUNT() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserCOUNT, 0)
}

func (s *CountContext) LEFT_P() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserLEFT_P, 0)
}

func (s *CountContext) RIGHT_P() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserRIGHT_P, 0)
}

func (s *CountContext) AS() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserAS, 0)
}

func (s *CountContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *CountContext) Nameable() INameableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameableContext)
}

func (s *CountContext) MULT() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserMULT, 0)
}

func (s *CountContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CountContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CountContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.EnterCount(s)
	}
}

func (s *CountContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.ExitCount(s)
	}
}

func (p *ViewSQLParser) Count() (localctx ICountContext) {
	this := p
	_ = this

	localctx = NewCountContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ViewSQLParserRULE_count)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(161)
		p.Match(ViewSQLParserCOUNT)
	}
	{
		p.SetState(162)
		p.Match(ViewSQLParserLEFT_P)
	}
	p.SetState(165)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ViewSQLParserLEFT_SP, ViewSQLParserSTRING:
		{
			p.SetState(163)
			p.Nameable()
		}

	case ViewSQLParserMULT:
		{
			p.SetState(164)
			p.Match(ViewSQLParserMULT)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(167)
		p.Match(ViewSQLParserRIGHT_P)
	}
	{
		p.SetState(168)
		p.Match(ViewSQLParserAS)
	}
	{
		p.SetState(169)
		p.Name()
	}

	return localctx
}

// ICalcContext is an interface to support dynamic dispatch.
type ICalcContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetAll returns the all token.
	GetAll() antlr.Token

	// SetAll sets the all token.
	SetAll(antlr.Token)

	// Getter signatures
	Key() IKeyContext
	Nameable() INameableContext
	AsClause() IAsClauseContext
	Aggregation() IAggregationContext
	Count() ICountContext
	MULT() antlr.TerminalNode

	// IsCalcContext differentiates from other interfaces.
	IsCalcContext()
}

type CalcContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	all    antlr.Token
}

func NewEmptyCalcContext() *CalcContext {
	var p = new(CalcContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ViewSQLParserRULE_calc
	return p
}

func (*CalcContext) IsCalcContext() {}

func NewCalcContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CalcContext {
	var p = new(CalcContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ViewSQLParserRULE_calc

	return p
}

func (s *CalcContext) GetParser() antlr.Parser { return s.parser }

func (s *CalcContext) GetAll() antlr.Token { return s.all }

func (s *CalcContext) SetAll(v antlr.Token) { s.all = v }

func (s *CalcContext) Key() IKeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKeyContext)
}

func (s *CalcContext) Nameable() INameableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameableContext)
}

func (s *CalcContext) AsClause() IAsClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAsClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAsClauseContext)
}

func (s *CalcContext) Aggregation() IAggregationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAggregationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAggregationContext)
}

func (s *CalcContext) Count() ICountContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICountContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICountContext)
}

func (s *CalcContext) MULT() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserMULT, 0)
}

func (s *CalcContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CalcContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CalcContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.EnterCalc(s)
	}
}

func (s *CalcContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.ExitCalc(s)
	}
}

func (p *ViewSQLParser) Calc() (localctx ICalcContext) {
	this := p
	_ = this

	localctx = NewCalcContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ViewSQLParserRULE_calc)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(177)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(171)
			p.Key()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(172)
			p.Nameable()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(173)
			p.AsClause()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(174)
			p.Aggregation()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(175)
			p.Count()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(176)

			var _m = p.Match(ViewSQLParserMULT)

			localctx.(*CalcContext).all = _m
		}

	}

	return localctx
}

// ICompContext is an interface to support dynamic dispatch.
type ICompContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOpType returns the opType token.
	GetOpType() antlr.Token

	// SetOpType sets the opType token.
	SetOpType(antlr.Token)

	// Getter signatures
	EQUAL() antlr.TerminalNode
	HIGHER() antlr.TerminalNode
	LOWER() antlr.TerminalNode
	HIGHER_EQUAL() antlr.TerminalNode
	LOWER_EQUAL() antlr.TerminalNode
	NOT_EQUAL() antlr.TerminalNode
	Name() INameContext

	// IsCompContext differentiates from other interfaces.
	IsCompContext()
}

type CompContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	opType antlr.Token
}

func NewEmptyCompContext() *CompContext {
	var p = new(CompContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ViewSQLParserRULE_comp
	return p
}

func (*CompContext) IsCompContext() {}

func NewCompContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompContext {
	var p = new(CompContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ViewSQLParserRULE_comp

	return p
}

func (s *CompContext) GetParser() antlr.Parser { return s.parser }

func (s *CompContext) GetOpType() antlr.Token { return s.opType }

func (s *CompContext) SetOpType(v antlr.Token) { s.opType = v }

func (s *CompContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserEQUAL, 0)
}

func (s *CompContext) HIGHER() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserHIGHER, 0)
}

func (s *CompContext) LOWER() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserLOWER, 0)
}

func (s *CompContext) HIGHER_EQUAL() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserHIGHER_EQUAL, 0)
}

func (s *CompContext) LOWER_EQUAL() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserLOWER_EQUAL, 0)
}

func (s *CompContext) NOT_EQUAL() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserNOT_EQUAL, 0)
}

func (s *CompContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *CompContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CompContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.EnterComp(s)
	}
}

func (s *CompContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.ExitComp(s)
	}
}

func (p *ViewSQLParser) Comp() (localctx ICompContext) {
	this := p
	_ = this

	localctx = NewCompContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, ViewSQLParserRULE_comp)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(181)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ViewSQLParserEQUAL, ViewSQLParserHIGHER, ViewSQLParserHIGHER_EQUAL, ViewSQLParserLOWER, ViewSQLParserLOWER_EQUAL, ViewSQLParserNOT_EQUAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(179)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*CompContext).opType = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&516096) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*CompContext).opType = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case ViewSQLParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(180)
			p.Name()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IConditionContext is an interface to support dynamic dispatch.
type IConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Nameable() INameableContext
	Comp() ICompContext
	Math() IMathContext

	// IsConditionContext differentiates from other interfaces.
	IsConditionContext()
}

type ConditionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionContext() *ConditionContext {
	var p = new(ConditionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ViewSQLParserRULE_condition
	return p
}

func (*ConditionContext) IsConditionContext() {}

func NewConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionContext {
	var p = new(ConditionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ViewSQLParserRULE_condition

	return p
}

func (s *ConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionContext) Nameable() INameableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameableContext)
}

func (s *ConditionContext) Comp() ICompContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICompContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICompContext)
}

func (s *ConditionContext) Math() IMathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMathContext)
}

func (s *ConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.EnterCondition(s)
	}
}

func (s *ConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.ExitCondition(s)
	}
}

func (p *ViewSQLParser) Condition() (localctx IConditionContext) {
	this := p
	_ = this

	localctx = NewConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, ViewSQLParserRULE_condition)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(183)
		p.Nameable()
	}
	{
		p.SetState(184)
		p.Comp()
	}
	{
		p.SetState(185)
		p.math(0)
	}

	return localctx
}

// ISortOrderContext is an interface to support dynamic dispatch.
type ISortOrderContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DESC() antlr.TerminalNode
	ASC() antlr.TerminalNode

	// IsSortOrderContext differentiates from other interfaces.
	IsSortOrderContext()
}

type SortOrderContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySortOrderContext() *SortOrderContext {
	var p = new(SortOrderContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ViewSQLParserRULE_sortOrder
	return p
}

func (*SortOrderContext) IsSortOrderContext() {}

func NewSortOrderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SortOrderContext {
	var p = new(SortOrderContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ViewSQLParserRULE_sortOrder

	return p
}

func (s *SortOrderContext) GetParser() antlr.Parser { return s.parser }

func (s *SortOrderContext) DESC() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserDESC, 0)
}

func (s *SortOrderContext) ASC() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserASC, 0)
}

func (s *SortOrderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SortOrderContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SortOrderContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.EnterSortOrder(s)
	}
}

func (s *SortOrderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.ExitSortOrder(s)
	}
}

func (p *ViewSQLParser) SortOrder() (localctx ISortOrderContext) {
	this := p
	_ = this

	localctx = NewSortOrderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, ViewSQLParserRULE_sortOrder)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(187)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ViewSQLParserASC || _la == ViewSQLParserDESC) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IContinuousRangeContext is an interface to support dynamic dispatch.
type IContinuousRangeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLeft returns the left token.
	GetLeft() antlr.Token

	// GetRight returns the right token.
	GetRight() antlr.Token

	// SetLeft sets the left token.
	SetLeft(antlr.Token)

	// SetRight sets the right token.
	SetRight(antlr.Token)

	// Getter signatures
	AllLEFT_SP() []antlr.TerminalNode
	LEFT_SP(i int) antlr.TerminalNode
	AllRIGHT_SP() []antlr.TerminalNode
	RIGHT_SP(i int) antlr.TerminalNode
	AllConstant() []IConstantContext
	Constant(i int) IConstantContext
	RANGE_SEP() antlr.TerminalNode

	// IsContinuousRangeContext differentiates from other interfaces.
	IsContinuousRangeContext()
}

type ContinuousRangeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	left   antlr.Token
	right  antlr.Token
}

func NewEmptyContinuousRangeContext() *ContinuousRangeContext {
	var p = new(ContinuousRangeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ViewSQLParserRULE_continuousRange
	return p
}

func (*ContinuousRangeContext) IsContinuousRangeContext() {}

func NewContinuousRangeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContinuousRangeContext {
	var p = new(ContinuousRangeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ViewSQLParserRULE_continuousRange

	return p
}

func (s *ContinuousRangeContext) GetParser() antlr.Parser { return s.parser }

func (s *ContinuousRangeContext) GetLeft() antlr.Token { return s.left }

func (s *ContinuousRangeContext) GetRight() antlr.Token { return s.right }

func (s *ContinuousRangeContext) SetLeft(v antlr.Token) { s.left = v }

func (s *ContinuousRangeContext) SetRight(v antlr.Token) { s.right = v }

func (s *ContinuousRangeContext) AllLEFT_SP() []antlr.TerminalNode {
	return s.GetTokens(ViewSQLParserLEFT_SP)
}

func (s *ContinuousRangeContext) LEFT_SP(i int) antlr.TerminalNode {
	return s.GetToken(ViewSQLParserLEFT_SP, i)
}

func (s *ContinuousRangeContext) AllRIGHT_SP() []antlr.TerminalNode {
	return s.GetTokens(ViewSQLParserRIGHT_SP)
}

func (s *ContinuousRangeContext) RIGHT_SP(i int) antlr.TerminalNode {
	return s.GetToken(ViewSQLParserRIGHT_SP, i)
}

func (s *ContinuousRangeContext) AllConstant() []IConstantContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConstantContext); ok {
			len++
		}
	}

	tst := make([]IConstantContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConstantContext); ok {
			tst[i] = t.(IConstantContext)
			i++
		}
	}

	return tst
}

func (s *ContinuousRangeContext) Constant(i int) IConstantContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *ContinuousRangeContext) RANGE_SEP() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserRANGE_SEP, 0)
}

func (s *ContinuousRangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContinuousRangeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContinuousRangeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.EnterContinuousRange(s)
	}
}

func (s *ContinuousRangeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.ExitContinuousRange(s)
	}
}

func (p *ViewSQLParser) ContinuousRange() (localctx IContinuousRangeContext) {
	this := p
	_ = this

	localctx = NewContinuousRangeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, ViewSQLParserRULE_continuousRange)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(189)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*ContinuousRangeContext).left = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == ViewSQLParserLEFT_SP || _la == ViewSQLParserRIGHT_SP) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*ContinuousRangeContext).left = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	{
		p.SetState(190)
		p.Constant()
	}
	{
		p.SetState(191)
		p.Match(ViewSQLParserRANGE_SEP)
	}
	{
		p.SetState(192)
		p.Constant()
	}

	{
		p.SetState(194)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*ContinuousRangeContext).right = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == ViewSQLParserLEFT_SP || _la == ViewSQLParserRIGHT_SP) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*ContinuousRangeContext).right = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ISparseRangeContext is an interface to support dynamic dispatch.
type ISparseRangeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LEFT_CURLY() antlr.TerminalNode
	AllConstant() []IConstantContext
	Constant(i int) IConstantContext
	RIGHT_CURLY() antlr.TerminalNode
	AllSEPARATOR() []antlr.TerminalNode
	SEPARATOR(i int) antlr.TerminalNode

	// IsSparseRangeContext differentiates from other interfaces.
	IsSparseRangeContext()
}

type SparseRangeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySparseRangeContext() *SparseRangeContext {
	var p = new(SparseRangeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ViewSQLParserRULE_sparseRange
	return p
}

func (*SparseRangeContext) IsSparseRangeContext() {}

func NewSparseRangeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SparseRangeContext {
	var p = new(SparseRangeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ViewSQLParserRULE_sparseRange

	return p
}

func (s *SparseRangeContext) GetParser() antlr.Parser { return s.parser }

func (s *SparseRangeContext) LEFT_CURLY() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserLEFT_CURLY, 0)
}

func (s *SparseRangeContext) AllConstant() []IConstantContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConstantContext); ok {
			len++
		}
	}

	tst := make([]IConstantContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConstantContext); ok {
			tst[i] = t.(IConstantContext)
			i++
		}
	}

	return tst
}

func (s *SparseRangeContext) Constant(i int) IConstantContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *SparseRangeContext) RIGHT_CURLY() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserRIGHT_CURLY, 0)
}

func (s *SparseRangeContext) AllSEPARATOR() []antlr.TerminalNode {
	return s.GetTokens(ViewSQLParserSEPARATOR)
}

func (s *SparseRangeContext) SEPARATOR(i int) antlr.TerminalNode {
	return s.GetToken(ViewSQLParserSEPARATOR, i)
}

func (s *SparseRangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SparseRangeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SparseRangeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.EnterSparseRange(s)
	}
}

func (s *SparseRangeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.ExitSparseRange(s)
	}
}

func (p *ViewSQLParser) SparseRange() (localctx ISparseRangeContext) {
	this := p
	_ = this

	localctx = NewSparseRangeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, ViewSQLParserRULE_sparseRange)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(196)
		p.Match(ViewSQLParserLEFT_CURLY)
	}
	{
		p.SetState(197)
		p.Constant()
	}
	p.SetState(202)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ViewSQLParserSEPARATOR {
		{
			p.SetState(198)
			p.Match(ViewSQLParserSEPARATOR)
		}
		{
			p.SetState(199)
			p.Constant()
		}

		p.SetState(204)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(205)
		p.Match(ViewSQLParserRIGHT_CURLY)
	}

	return localctx
}

// IRangeContext is an interface to support dynamic dispatch.
type IRangeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ContinuousRange() IContinuousRangeContext
	SparseRange() ISparseRangeContext

	// IsRangeContext differentiates from other interfaces.
	IsRangeContext()
}

type RangeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRangeContext() *RangeContext {
	var p = new(RangeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ViewSQLParserRULE_range
	return p
}

func (*RangeContext) IsRangeContext() {}

func NewRangeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RangeContext {
	var p = new(RangeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ViewSQLParserRULE_range

	return p
}

func (s *RangeContext) GetParser() antlr.Parser { return s.parser }

func (s *RangeContext) ContinuousRange() IContinuousRangeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContinuousRangeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContinuousRangeContext)
}

func (s *RangeContext) SparseRange() ISparseRangeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISparseRangeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISparseRangeContext)
}

func (s *RangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RangeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.EnterRange(s)
	}
}

func (s *RangeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.ExitRange(s)
	}
}

func (p *ViewSQLParser) Range_() (localctx IRangeContext) {
	this := p
	_ = this

	localctx = NewRangeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, ViewSQLParserRULE_range)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(209)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ViewSQLParserLEFT_SP, ViewSQLParserRIGHT_SP:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(207)
			p.ContinuousRange()
		}

	case ViewSQLParserLEFT_CURLY:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(208)
			p.SparseRange()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ICreateContext is an interface to support dynamic dispatch.
type ICreateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CREATE() antlr.TerminalNode
	VIEW() antlr.TerminalNode
	LEFT_P() antlr.TerminalNode
	AllSTRING() []antlr.TerminalNode
	STRING(i int) antlr.TerminalNode
	SEPARATOR() antlr.TerminalNode
	RIGHT_P() antlr.TerminalNode

	// IsCreateContext differentiates from other interfaces.
	IsCreateContext()
}

type CreateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCreateContext() *CreateContext {
	var p = new(CreateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ViewSQLParserRULE_create
	return p
}

func (*CreateContext) IsCreateContext() {}

func NewCreateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CreateContext {
	var p = new(CreateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ViewSQLParserRULE_create

	return p
}

func (s *CreateContext) GetParser() antlr.Parser { return s.parser }

func (s *CreateContext) CREATE() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserCREATE, 0)
}

func (s *CreateContext) VIEW() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserVIEW, 0)
}

func (s *CreateContext) LEFT_P() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserLEFT_P, 0)
}

func (s *CreateContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(ViewSQLParserSTRING)
}

func (s *CreateContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(ViewSQLParserSTRING, i)
}

func (s *CreateContext) SEPARATOR() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserSEPARATOR, 0)
}

func (s *CreateContext) RIGHT_P() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserRIGHT_P, 0)
}

func (s *CreateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CreateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CreateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.EnterCreate(s)
	}
}

func (s *CreateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.ExitCreate(s)
	}
}

func (p *ViewSQLParser) Create() (localctx ICreateContext) {
	this := p
	_ = this

	localctx = NewCreateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, ViewSQLParserRULE_create)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(211)
		p.Match(ViewSQLParserCREATE)
	}
	{
		p.SetState(212)
		p.Match(ViewSQLParserVIEW)
	}
	{
		p.SetState(213)
		p.Match(ViewSQLParserLEFT_P)
	}
	{
		p.SetState(214)
		p.Match(ViewSQLParserSTRING)
	}
	{
		p.SetState(215)
		p.Match(ViewSQLParserSEPARATOR)
	}
	{
		p.SetState(216)
		p.Match(ViewSQLParserSTRING)
	}
	{
		p.SetState(217)
		p.Match(ViewSQLParserRIGHT_P)
	}

	return localctx
}

// IWithContext is an interface to support dynamic dispatch.
type IWithContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WITH() antlr.TerminalNode
	AllSTRING() []antlr.TerminalNode
	STRING(i int) antlr.TerminalNode
	AllIN() []antlr.TerminalNode
	IN(i int) antlr.TerminalNode
	AllRange_() []IRangeContext
	Range_(i int) IRangeContext
	AllSEPARATOR() []antlr.TerminalNode
	SEPARATOR(i int) antlr.TerminalNode

	// IsWithContext differentiates from other interfaces.
	IsWithContext()
}

type WithContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWithContext() *WithContext {
	var p = new(WithContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ViewSQLParserRULE_with
	return p
}

func (*WithContext) IsWithContext() {}

func NewWithContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WithContext {
	var p = new(WithContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ViewSQLParserRULE_with

	return p
}

func (s *WithContext) GetParser() antlr.Parser { return s.parser }

func (s *WithContext) WITH() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserWITH, 0)
}

func (s *WithContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(ViewSQLParserSTRING)
}

func (s *WithContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(ViewSQLParserSTRING, i)
}

func (s *WithContext) AllIN() []antlr.TerminalNode {
	return s.GetTokens(ViewSQLParserIN)
}

func (s *WithContext) IN(i int) antlr.TerminalNode {
	return s.GetToken(ViewSQLParserIN, i)
}

func (s *WithContext) AllRange_() []IRangeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRangeContext); ok {
			len++
		}
	}

	tst := make([]IRangeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRangeContext); ok {
			tst[i] = t.(IRangeContext)
			i++
		}
	}

	return tst
}

func (s *WithContext) Range_(i int) IRangeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRangeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRangeContext)
}

func (s *WithContext) AllSEPARATOR() []antlr.TerminalNode {
	return s.GetTokens(ViewSQLParserSEPARATOR)
}

func (s *WithContext) SEPARATOR(i int) antlr.TerminalNode {
	return s.GetToken(ViewSQLParserSEPARATOR, i)
}

func (s *WithContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WithContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WithContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.EnterWith(s)
	}
}

func (s *WithContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.ExitWith(s)
	}
}

func (p *ViewSQLParser) With() (localctx IWithContext) {
	this := p
	_ = this

	localctx = NewWithContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, ViewSQLParserRULE_with)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(219)
		p.Match(ViewSQLParserWITH)
	}
	{
		p.SetState(220)
		p.Match(ViewSQLParserSTRING)
	}
	{
		p.SetState(221)
		p.Match(ViewSQLParserIN)
	}
	{
		p.SetState(222)
		p.Range_()
	}
	p.SetState(229)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ViewSQLParserSEPARATOR {
		{
			p.SetState(223)
			p.Match(ViewSQLParserSEPARATOR)
		}
		{
			p.SetState(224)
			p.Match(ViewSQLParserSTRING)
		}
		{
			p.SetState(225)
			p.Match(ViewSQLParserIN)
		}
		{
			p.SetState(226)
			p.Range_()
		}

		p.SetState(231)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ISelectContext is an interface to support dynamic dispatch.
type ISelectContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SELECT() antlr.TerminalNode
	AllCalc() []ICalcContext
	Calc(i int) ICalcContext
	AllSEPARATOR() []antlr.TerminalNode
	SEPARATOR(i int) antlr.TerminalNode

	// IsSelectContext differentiates from other interfaces.
	IsSelectContext()
}

type SelectContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectContext() *SelectContext {
	var p = new(SelectContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ViewSQLParserRULE_select
	return p
}

func (*SelectContext) IsSelectContext() {}

func NewSelectContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectContext {
	var p = new(SelectContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ViewSQLParserRULE_select

	return p
}

func (s *SelectContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectContext) SELECT() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserSELECT, 0)
}

func (s *SelectContext) AllCalc() []ICalcContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICalcContext); ok {
			len++
		}
	}

	tst := make([]ICalcContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICalcContext); ok {
			tst[i] = t.(ICalcContext)
			i++
		}
	}

	return tst
}

func (s *SelectContext) Calc(i int) ICalcContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICalcContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICalcContext)
}

func (s *SelectContext) AllSEPARATOR() []antlr.TerminalNode {
	return s.GetTokens(ViewSQLParserSEPARATOR)
}

func (s *SelectContext) SEPARATOR(i int) antlr.TerminalNode {
	return s.GetToken(ViewSQLParserSEPARATOR, i)
}

func (s *SelectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.EnterSelect(s)
	}
}

func (s *SelectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.ExitSelect(s)
	}
}

func (p *ViewSQLParser) Select_() (localctx ISelectContext) {
	this := p
	_ = this

	localctx = NewSelectContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, ViewSQLParserRULE_select)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(232)
		p.Match(ViewSQLParserSELECT)
	}
	{
		p.SetState(233)
		p.Calc()
	}
	p.SetState(238)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ViewSQLParserSEPARATOR {
		{
			p.SetState(234)
			p.Match(ViewSQLParserSEPARATOR)
		}
		{
			p.SetState(235)
			p.Calc()
		}

		p.SetState(240)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFromContext is an interface to support dynamic dispatch.
type IFromContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FROM() antlr.TerminalNode
	AllName() []INameContext
	Name(i int) INameContext
	AllSEPARATOR() []antlr.TerminalNode
	SEPARATOR(i int) antlr.TerminalNode

	// IsFromContext differentiates from other interfaces.
	IsFromContext()
}

type FromContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFromContext() *FromContext {
	var p = new(FromContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ViewSQLParserRULE_from
	return p
}

func (*FromContext) IsFromContext() {}

func NewFromContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FromContext {
	var p = new(FromContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ViewSQLParserRULE_from

	return p
}

func (s *FromContext) GetParser() antlr.Parser { return s.parser }

func (s *FromContext) FROM() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserFROM, 0)
}

func (s *FromContext) AllName() []INameContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INameContext); ok {
			len++
		}
	}

	tst := make([]INameContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INameContext); ok {
			tst[i] = t.(INameContext)
			i++
		}
	}

	return tst
}

func (s *FromContext) Name(i int) INameContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *FromContext) AllSEPARATOR() []antlr.TerminalNode {
	return s.GetTokens(ViewSQLParserSEPARATOR)
}

func (s *FromContext) SEPARATOR(i int) antlr.TerminalNode {
	return s.GetToken(ViewSQLParserSEPARATOR, i)
}

func (s *FromContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FromContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FromContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.EnterFrom(s)
	}
}

func (s *FromContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.ExitFrom(s)
	}
}

func (p *ViewSQLParser) From() (localctx IFromContext) {
	this := p
	_ = this

	localctx = NewFromContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, ViewSQLParserRULE_from)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(241)
		p.Match(ViewSQLParserFROM)
	}
	{
		p.SetState(242)
		p.Name()
	}
	p.SetState(247)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ViewSQLParserSEPARATOR {
		{
			p.SetState(243)
			p.Match(ViewSQLParserSEPARATOR)
		}
		{
			p.SetState(244)
			p.Name()
		}

		p.SetState(249)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IWhereContext is an interface to support dynamic dispatch.
type IWhereContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WHERE() antlr.TerminalNode
	AllCondition() []IConditionContext
	Condition(i int) IConditionContext
	AllAND() []antlr.TerminalNode
	AND(i int) antlr.TerminalNode

	// IsWhereContext differentiates from other interfaces.
	IsWhereContext()
}

type WhereContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhereContext() *WhereContext {
	var p = new(WhereContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ViewSQLParserRULE_where
	return p
}

func (*WhereContext) IsWhereContext() {}

func NewWhereContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhereContext {
	var p = new(WhereContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ViewSQLParserRULE_where

	return p
}

func (s *WhereContext) GetParser() antlr.Parser { return s.parser }

func (s *WhereContext) WHERE() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserWHERE, 0)
}

func (s *WhereContext) AllCondition() []IConditionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConditionContext); ok {
			len++
		}
	}

	tst := make([]IConditionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConditionContext); ok {
			tst[i] = t.(IConditionContext)
			i++
		}
	}

	return tst
}

func (s *WhereContext) Condition(i int) IConditionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *WhereContext) AllAND() []antlr.TerminalNode {
	return s.GetTokens(ViewSQLParserAND)
}

func (s *WhereContext) AND(i int) antlr.TerminalNode {
	return s.GetToken(ViewSQLParserAND, i)
}

func (s *WhereContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhereContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhereContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.EnterWhere(s)
	}
}

func (s *WhereContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.ExitWhere(s)
	}
}

func (p *ViewSQLParser) Where() (localctx IWhereContext) {
	this := p
	_ = this

	localctx = NewWhereContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, ViewSQLParserRULE_where)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(250)
		p.Match(ViewSQLParserWHERE)
	}
	{
		p.SetState(251)
		p.Condition()
	}
	p.SetState(256)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ViewSQLParserAND {
		{
			p.SetState(252)
			p.Match(ViewSQLParserAND)
		}
		{
			p.SetState(253)
			p.Condition()
		}

		p.SetState(258)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IGroupbyContext is an interface to support dynamic dispatch.
type IGroupbyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	GROUP() antlr.TerminalNode
	BY() antlr.TerminalNode
	AllNameable() []INameableContext
	Nameable(i int) INameableContext
	AllSEPARATOR() []antlr.TerminalNode
	SEPARATOR(i int) antlr.TerminalNode

	// IsGroupbyContext differentiates from other interfaces.
	IsGroupbyContext()
}

type GroupbyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGroupbyContext() *GroupbyContext {
	var p = new(GroupbyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ViewSQLParserRULE_groupby
	return p
}

func (*GroupbyContext) IsGroupbyContext() {}

func NewGroupbyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GroupbyContext {
	var p = new(GroupbyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ViewSQLParserRULE_groupby

	return p
}

func (s *GroupbyContext) GetParser() antlr.Parser { return s.parser }

func (s *GroupbyContext) GROUP() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserGROUP, 0)
}

func (s *GroupbyContext) BY() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserBY, 0)
}

func (s *GroupbyContext) AllNameable() []INameableContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INameableContext); ok {
			len++
		}
	}

	tst := make([]INameableContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INameableContext); ok {
			tst[i] = t.(INameableContext)
			i++
		}
	}

	return tst
}

func (s *GroupbyContext) Nameable(i int) INameableContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameableContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameableContext)
}

func (s *GroupbyContext) AllSEPARATOR() []antlr.TerminalNode {
	return s.GetTokens(ViewSQLParserSEPARATOR)
}

func (s *GroupbyContext) SEPARATOR(i int) antlr.TerminalNode {
	return s.GetToken(ViewSQLParserSEPARATOR, i)
}

func (s *GroupbyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupbyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GroupbyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.EnterGroupby(s)
	}
}

func (s *GroupbyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.ExitGroupby(s)
	}
}

func (p *ViewSQLParser) Groupby() (localctx IGroupbyContext) {
	this := p
	_ = this

	localctx = NewGroupbyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, ViewSQLParserRULE_groupby)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(259)
		p.Match(ViewSQLParserGROUP)
	}
	{
		p.SetState(260)
		p.Match(ViewSQLParserBY)
	}
	{
		p.SetState(261)
		p.Nameable()
	}
	p.SetState(266)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ViewSQLParserSEPARATOR {
		{
			p.SetState(262)
			p.Match(ViewSQLParserSEPARATOR)
		}
		{
			p.SetState(263)
			p.Nameable()
		}

		p.SetState(268)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IOrderbyContext is an interface to support dynamic dispatch.
type IOrderbyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ORDER() antlr.TerminalNode
	BY() antlr.TerminalNode
	AllNameable() []INameableContext
	Nameable(i int) INameableContext
	AllSortOrder() []ISortOrderContext
	SortOrder(i int) ISortOrderContext
	AllSEPARATOR() []antlr.TerminalNode
	SEPARATOR(i int) antlr.TerminalNode

	// IsOrderbyContext differentiates from other interfaces.
	IsOrderbyContext()
}

type OrderbyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrderbyContext() *OrderbyContext {
	var p = new(OrderbyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ViewSQLParserRULE_orderby
	return p
}

func (*OrderbyContext) IsOrderbyContext() {}

func NewOrderbyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrderbyContext {
	var p = new(OrderbyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ViewSQLParserRULE_orderby

	return p
}

func (s *OrderbyContext) GetParser() antlr.Parser { return s.parser }

func (s *OrderbyContext) ORDER() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserORDER, 0)
}

func (s *OrderbyContext) BY() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserBY, 0)
}

func (s *OrderbyContext) AllNameable() []INameableContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INameableContext); ok {
			len++
		}
	}

	tst := make([]INameableContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INameableContext); ok {
			tst[i] = t.(INameableContext)
			i++
		}
	}

	return tst
}

func (s *OrderbyContext) Nameable(i int) INameableContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameableContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameableContext)
}

func (s *OrderbyContext) AllSortOrder() []ISortOrderContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISortOrderContext); ok {
			len++
		}
	}

	tst := make([]ISortOrderContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISortOrderContext); ok {
			tst[i] = t.(ISortOrderContext)
			i++
		}
	}

	return tst
}

func (s *OrderbyContext) SortOrder(i int) ISortOrderContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISortOrderContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISortOrderContext)
}

func (s *OrderbyContext) AllSEPARATOR() []antlr.TerminalNode {
	return s.GetTokens(ViewSQLParserSEPARATOR)
}

func (s *OrderbyContext) SEPARATOR(i int) antlr.TerminalNode {
	return s.GetToken(ViewSQLParserSEPARATOR, i)
}

func (s *OrderbyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrderbyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrderbyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.EnterOrderby(s)
	}
}

func (s *OrderbyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.ExitOrderby(s)
	}
}

func (p *ViewSQLParser) Orderby() (localctx IOrderbyContext) {
	this := p
	_ = this

	localctx = NewOrderbyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, ViewSQLParserRULE_orderby)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(269)
		p.Match(ViewSQLParserORDER)
	}
	{
		p.SetState(270)
		p.Match(ViewSQLParserBY)
	}
	{
		p.SetState(271)
		p.Nameable()
	}
	{
		p.SetState(272)
		p.SortOrder()
	}
	p.SetState(279)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ViewSQLParserSEPARATOR {
		{
			p.SetState(273)
			p.Match(ViewSQLParserSEPARATOR)
		}
		{
			p.SetState(274)
			p.Nameable()
		}
		{
			p.SetState(275)
			p.SortOrder()
		}

		p.SetState(281)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ILimitContext is an interface to support dynamic dispatch.
type ILimitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LIMIT() antlr.TerminalNode
	INT() antlr.TerminalNode

	// IsLimitContext differentiates from other interfaces.
	IsLimitContext()
}

type LimitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLimitContext() *LimitContext {
	var p = new(LimitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ViewSQLParserRULE_limit
	return p
}

func (*LimitContext) IsLimitContext() {}

func NewLimitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LimitContext {
	var p = new(LimitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ViewSQLParserRULE_limit

	return p
}

func (s *LimitContext) GetParser() antlr.Parser { return s.parser }

func (s *LimitContext) LIMIT() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserLIMIT, 0)
}

func (s *LimitContext) INT() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserINT, 0)
}

func (s *LimitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LimitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LimitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.EnterLimit(s)
	}
}

func (s *LimitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.ExitLimit(s)
	}
}

func (p *ViewSQLParser) Limit() (localctx ILimitContext) {
	this := p
	_ = this

	localctx = NewLimitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, ViewSQLParserRULE_limit)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(282)
		p.Match(ViewSQLParserLIMIT)
	}
	{
		p.SetState(283)
		p.Match(ViewSQLParserINT)
	}

	return localctx
}

// IViewContext is an interface to support dynamic dispatch.
type IViewContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetFirstGroupBy returns the firstGroupBy rule contexts.
	GetFirstGroupBy() IGroupbyContext

	// GetSecondGroupBy returns the secondGroupBy rule contexts.
	GetSecondGroupBy() IGroupbyContext

	// SetFirstGroupBy sets the firstGroupBy rule contexts.
	SetFirstGroupBy(IGroupbyContext)

	// SetSecondGroupBy sets the secondGroupBy rule contexts.
	SetSecondGroupBy(IGroupbyContext)

	// Getter signatures
	Create() ICreateContext
	AS() antlr.TerminalNode
	Select_() ISelectContext
	From() IFromContext
	AllGroupby() []IGroupbyContext
	Groupby(i int) IGroupbyContext
	With() IWithContext
	Where() IWhereContext
	Orderby() IOrderbyContext
	Limit() ILimitContext

	// IsViewContext differentiates from other interfaces.
	IsViewContext()
}

type ViewContext struct {
	*antlr.BaseParserRuleContext
	parser        antlr.Parser
	firstGroupBy  IGroupbyContext
	secondGroupBy IGroupbyContext
}

func NewEmptyViewContext() *ViewContext {
	var p = new(ViewContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ViewSQLParserRULE_view
	return p
}

func (*ViewContext) IsViewContext() {}

func NewViewContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ViewContext {
	var p = new(ViewContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ViewSQLParserRULE_view

	return p
}

func (s *ViewContext) GetParser() antlr.Parser { return s.parser }

func (s *ViewContext) GetFirstGroupBy() IGroupbyContext { return s.firstGroupBy }

func (s *ViewContext) GetSecondGroupBy() IGroupbyContext { return s.secondGroupBy }

func (s *ViewContext) SetFirstGroupBy(v IGroupbyContext) { s.firstGroupBy = v }

func (s *ViewContext) SetSecondGroupBy(v IGroupbyContext) { s.secondGroupBy = v }

func (s *ViewContext) Create() ICreateContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICreateContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICreateContext)
}

func (s *ViewContext) AS() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserAS, 0)
}

func (s *ViewContext) Select_() ISelectContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectContext)
}

func (s *ViewContext) From() IFromContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFromContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFromContext)
}

func (s *ViewContext) AllGroupby() []IGroupbyContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IGroupbyContext); ok {
			len++
		}
	}

	tst := make([]IGroupbyContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IGroupbyContext); ok {
			tst[i] = t.(IGroupbyContext)
			i++
		}
	}

	return tst
}

func (s *ViewContext) Groupby(i int) IGroupbyContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGroupbyContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGroupbyContext)
}

func (s *ViewContext) With() IWithContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWithContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWithContext)
}

func (s *ViewContext) Where() IWhereContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhereContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhereContext)
}

func (s *ViewContext) Orderby() IOrderbyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOrderbyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOrderbyContext)
}

func (s *ViewContext) Limit() ILimitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILimitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILimitContext)
}

func (s *ViewContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ViewContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ViewContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.EnterView(s)
	}
}

func (s *ViewContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.ExitView(s)
	}
}

func (p *ViewSQLParser) View() (localctx IViewContext) {
	this := p
	_ = this

	localctx = NewViewContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, ViewSQLParserRULE_view)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(285)
		p.Create()
	}
	p.SetState(287)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ViewSQLParserWITH {
		{
			p.SetState(286)
			p.With()
		}

	}
	{
		p.SetState(289)
		p.Match(ViewSQLParserAS)
	}
	{
		p.SetState(290)
		p.Select_()
	}
	{
		p.SetState(291)
		p.From()
	}
	p.SetState(293)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ViewSQLParserWHERE {
		{
			p.SetState(292)
			p.Where()
		}

	}
	p.SetState(296)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(295)

			var _x = p.Groupby()

			localctx.(*ViewContext).firstGroupBy = _x
		}

	}
	{
		p.SetState(298)

		var _x = p.Groupby()

		localctx.(*ViewContext).secondGroupBy = _x
	}
	p.SetState(300)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ViewSQLParserORDER {
		{
			p.SetState(299)
			p.Orderby()
		}

	}
	p.SetState(303)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ViewSQLParserLIMIT {
		{
			p.SetState(302)
			p.Limit()
		}

	}

	return localctx
}

// ICheckContext is an interface to support dynamic dispatch.
type ICheckContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetType_ returns the type_ token.
	GetType_() antlr.Token

	// SetType_ sets the type_ token.
	SetType_(antlr.Token)

	// Getter signatures
	CHECK() antlr.TerminalNode
	Constant() IConstantContext
	HIGHER() antlr.TerminalNode
	LOWER() antlr.TerminalNode
	HIGHER_EQUAL() antlr.TerminalNode
	LOWER_EQUAL() antlr.TerminalNode

	// IsCheckContext differentiates from other interfaces.
	IsCheckContext()
}

type CheckContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	type_  antlr.Token
}

func NewEmptyCheckContext() *CheckContext {
	var p = new(CheckContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ViewSQLParserRULE_check
	return p
}

func (*CheckContext) IsCheckContext() {}

func NewCheckContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CheckContext {
	var p = new(CheckContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ViewSQLParserRULE_check

	return p
}

func (s *CheckContext) GetParser() antlr.Parser { return s.parser }

func (s *CheckContext) GetType_() antlr.Token { return s.type_ }

func (s *CheckContext) SetType_(v antlr.Token) { s.type_ = v }

func (s *CheckContext) CHECK() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserCHECK, 0)
}

func (s *CheckContext) Constant() IConstantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *CheckContext) HIGHER() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserHIGHER, 0)
}

func (s *CheckContext) LOWER() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserLOWER, 0)
}

func (s *CheckContext) HIGHER_EQUAL() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserHIGHER_EQUAL, 0)
}

func (s *CheckContext) LOWER_EQUAL() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserLOWER_EQUAL, 0)
}

func (s *CheckContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CheckContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CheckContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.EnterCheck(s)
	}
}

func (s *CheckContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.ExitCheck(s)
	}
}

func (p *ViewSQLParser) Check() (localctx ICheckContext) {
	this := p
	_ = this

	localctx = NewCheckContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, ViewSQLParserRULE_check)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(305)
		p.Match(ViewSQLParserCHECK)
	}
	{
		p.SetState(306)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*CheckContext).type_ = _lt

		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&245760) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*CheckContext).type_ = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(307)
		p.Constant()
	}

	return localctx
}

// IForeignkeyContext is an interface to support dynamic dispatch.
type IForeignkeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTableName returns the tableName rule contexts.
	GetTableName() INameContext

	// GetColumnName returns the columnName rule contexts.
	GetColumnName() INameContext

	// SetTableName sets the tableName rule contexts.
	SetTableName(INameContext)

	// SetColumnName sets the columnName rule contexts.
	SetColumnName(INameContext)

	// Getter signatures
	FOREIGN() antlr.TerminalNode
	KEY() antlr.TerminalNode
	REFERENCES() antlr.TerminalNode
	LEFT_P() antlr.TerminalNode
	RIGHT_P() antlr.TerminalNode
	AllName() []INameContext
	Name(i int) INameContext

	// IsForeignkeyContext differentiates from other interfaces.
	IsForeignkeyContext()
}

type ForeignkeyContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	tableName  INameContext
	columnName INameContext
}

func NewEmptyForeignkeyContext() *ForeignkeyContext {
	var p = new(ForeignkeyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ViewSQLParserRULE_foreignkey
	return p
}

func (*ForeignkeyContext) IsForeignkeyContext() {}

func NewForeignkeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForeignkeyContext {
	var p = new(ForeignkeyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ViewSQLParserRULE_foreignkey

	return p
}

func (s *ForeignkeyContext) GetParser() antlr.Parser { return s.parser }

func (s *ForeignkeyContext) GetTableName() INameContext { return s.tableName }

func (s *ForeignkeyContext) GetColumnName() INameContext { return s.columnName }

func (s *ForeignkeyContext) SetTableName(v INameContext) { s.tableName = v }

func (s *ForeignkeyContext) SetColumnName(v INameContext) { s.columnName = v }

func (s *ForeignkeyContext) FOREIGN() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserFOREIGN, 0)
}

func (s *ForeignkeyContext) KEY() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserKEY, 0)
}

func (s *ForeignkeyContext) REFERENCES() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserREFERENCES, 0)
}

func (s *ForeignkeyContext) LEFT_P() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserLEFT_P, 0)
}

func (s *ForeignkeyContext) RIGHT_P() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserRIGHT_P, 0)
}

func (s *ForeignkeyContext) AllName() []INameContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INameContext); ok {
			len++
		}
	}

	tst := make([]INameContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INameContext); ok {
			tst[i] = t.(INameContext)
			i++
		}
	}

	return tst
}

func (s *ForeignkeyContext) Name(i int) INameContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *ForeignkeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForeignkeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForeignkeyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.EnterForeignkey(s)
	}
}

func (s *ForeignkeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.ExitForeignkey(s)
	}
}

func (p *ViewSQLParser) Foreignkey() (localctx IForeignkeyContext) {
	this := p
	_ = this

	localctx = NewForeignkeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, ViewSQLParserRULE_foreignkey)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(309)
		p.Match(ViewSQLParserFOREIGN)
	}
	{
		p.SetState(310)
		p.Match(ViewSQLParserKEY)
	}
	{
		p.SetState(311)
		p.Match(ViewSQLParserREFERENCES)
	}
	{
		p.SetState(312)

		var _x = p.Name()

		localctx.(*ForeignkeyContext).tableName = _x
	}
	{
		p.SetState(313)
		p.Match(ViewSQLParserLEFT_P)
	}
	{
		p.SetState(314)

		var _x = p.Name()

		localctx.(*ForeignkeyContext).columnName = _x
	}
	{
		p.SetState(315)
		p.Match(ViewSQLParserRIGHT_P)
	}

	return localctx
}

// IPrimarykeyContext is an interface to support dynamic dispatch.
type IPrimarykeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PRIMARY() antlr.TerminalNode
	KEY() antlr.TerminalNode

	// IsPrimarykeyContext differentiates from other interfaces.
	IsPrimarykeyContext()
}

type PrimarykeyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimarykeyContext() *PrimarykeyContext {
	var p = new(PrimarykeyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ViewSQLParserRULE_primarykey
	return p
}

func (*PrimarykeyContext) IsPrimarykeyContext() {}

func NewPrimarykeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimarykeyContext {
	var p = new(PrimarykeyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ViewSQLParserRULE_primarykey

	return p
}

func (s *PrimarykeyContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimarykeyContext) PRIMARY() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserPRIMARY, 0)
}

func (s *PrimarykeyContext) KEY() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserKEY, 0)
}

func (s *PrimarykeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimarykeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimarykeyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.EnterPrimarykey(s)
	}
}

func (s *PrimarykeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.ExitPrimarykey(s)
	}
}

func (p *ViewSQLParser) Primarykey() (localctx IPrimarykeyContext) {
	this := p
	_ = this

	localctx = NewPrimarykeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, ViewSQLParserRULE_primarykey)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(317)
		p.Match(ViewSQLParserPRIMARY)
	}
	{
		p.SetState(318)
		p.Match(ViewSQLParserKEY)
	}

	return localctx
}

// IConstraintContext is an interface to support dynamic dispatch.
type IConstraintContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Primarykey() IPrimarykeyContext
	Foreignkey() IForeignkeyContext
	Check() ICheckContext

	// IsConstraintContext differentiates from other interfaces.
	IsConstraintContext()
}

type ConstraintContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstraintContext() *ConstraintContext {
	var p = new(ConstraintContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ViewSQLParserRULE_constraint
	return p
}

func (*ConstraintContext) IsConstraintContext() {}

func NewConstraintContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstraintContext {
	var p = new(ConstraintContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ViewSQLParserRULE_constraint

	return p
}

func (s *ConstraintContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstraintContext) Primarykey() IPrimarykeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimarykeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimarykeyContext)
}

func (s *ConstraintContext) Foreignkey() IForeignkeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForeignkeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForeignkeyContext)
}

func (s *ConstraintContext) Check() ICheckContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICheckContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICheckContext)
}

func (s *ConstraintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstraintContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstraintContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.EnterConstraint(s)
	}
}

func (s *ConstraintContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.ExitConstraint(s)
	}
}

func (p *ViewSQLParser) Constraint() (localctx IConstraintContext) {
	this := p
	_ = this

	localctx = NewConstraintContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, ViewSQLParserRULE_constraint)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(323)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ViewSQLParserPRIMARY:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(320)
			p.Primarykey()
		}

	case ViewSQLParserFOREIGN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(321)
			p.Foreignkey()
		}

	case ViewSQLParserCHECK:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(322)
			p.Check()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IColumnsContext is an interface to support dynamic dispatch.
type IColumnsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetType_ returns the type_ token.
	GetType_() antlr.Token

	// SetType_ sets the type_ token.
	SetType_(antlr.Token)

	// Getter signatures
	STRING() antlr.TerminalNode
	COUNTER() antlr.TerminalNode
	INTEGER() antlr.TerminalNode
	BOOLEAN() antlr.TerminalNode
	VARCHAR() antlr.TerminalNode
	DATE_TYPE() antlr.TerminalNode
	DEFAULT() antlr.TerminalNode
	Constant() IConstantContext
	Constraint() IConstraintContext

	// IsColumnsContext differentiates from other interfaces.
	IsColumnsContext()
}

type ColumnsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	type_  antlr.Token
}

func NewEmptyColumnsContext() *ColumnsContext {
	var p = new(ColumnsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ViewSQLParserRULE_columns
	return p
}

func (*ColumnsContext) IsColumnsContext() {}

func NewColumnsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColumnsContext {
	var p = new(ColumnsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ViewSQLParserRULE_columns

	return p
}

func (s *ColumnsContext) GetParser() antlr.Parser { return s.parser }

func (s *ColumnsContext) GetType_() antlr.Token { return s.type_ }

func (s *ColumnsContext) SetType_(v antlr.Token) { s.type_ = v }

func (s *ColumnsContext) STRING() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserSTRING, 0)
}

func (s *ColumnsContext) COUNTER() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserCOUNTER, 0)
}

func (s *ColumnsContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserINTEGER, 0)
}

func (s *ColumnsContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserBOOLEAN, 0)
}

func (s *ColumnsContext) VARCHAR() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserVARCHAR, 0)
}

func (s *ColumnsContext) DATE_TYPE() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserDATE_TYPE, 0)
}

func (s *ColumnsContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserDEFAULT, 0)
}

func (s *ColumnsContext) Constant() IConstantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *ColumnsContext) Constraint() IConstraintContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstraintContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstraintContext)
}

func (s *ColumnsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ColumnsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.EnterColumns(s)
	}
}

func (s *ColumnsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.ExitColumns(s)
	}
}

func (p *ViewSQLParser) Columns() (localctx IColumnsContext) {
	this := p
	_ = this

	localctx = NewColumnsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, ViewSQLParserRULE_columns)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(325)
		p.Match(ViewSQLParserSTRING)
	}
	{
		p.SetState(326)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*ColumnsContext).type_ = _lt

		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-60)) & ^0x3f) == 0 && ((int64(1)<<(_la-60))&31) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*ColumnsContext).type_ = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(329)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ViewSQLParserDEFAULT {
		{
			p.SetState(327)
			p.Match(ViewSQLParserDEFAULT)
		}
		{
			p.SetState(328)
			p.Constant()
		}

	}
	p.SetState(332)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&20268397346422784) != 0 {
		{
			p.SetState(331)
			p.Constraint()
		}

	}

	return localctx
}

// ICreatetableContext is an interface to support dynamic dispatch.
type ICreatetableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetPolicy returns the policy token.
	GetPolicy() antlr.Token

	// SetPolicy sets the policy token.
	SetPolicy(antlr.Token)

	// Getter signatures
	CREATE() antlr.TerminalNode
	TABLE() antlr.TerminalNode
	Name() INameContext
	LEFT_P() antlr.TerminalNode
	AllColumns() []IColumnsContext
	Columns(i int) IColumnsContext
	RIGHT_P() antlr.TerminalNode
	AW() antlr.TerminalNode
	RW() antlr.TerminalNode
	LWW() antlr.TerminalNode
	AllSEPARATOR() []antlr.TerminalNode
	SEPARATOR(i int) antlr.TerminalNode

	// IsCreatetableContext differentiates from other interfaces.
	IsCreatetableContext()
}

type CreatetableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	policy antlr.Token
}

func NewEmptyCreatetableContext() *CreatetableContext {
	var p = new(CreatetableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ViewSQLParserRULE_createtable
	return p
}

func (*CreatetableContext) IsCreatetableContext() {}

func NewCreatetableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CreatetableContext {
	var p = new(CreatetableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ViewSQLParserRULE_createtable

	return p
}

func (s *CreatetableContext) GetParser() antlr.Parser { return s.parser }

func (s *CreatetableContext) GetPolicy() antlr.Token { return s.policy }

func (s *CreatetableContext) SetPolicy(v antlr.Token) { s.policy = v }

func (s *CreatetableContext) CREATE() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserCREATE, 0)
}

func (s *CreatetableContext) TABLE() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserTABLE, 0)
}

func (s *CreatetableContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *CreatetableContext) LEFT_P() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserLEFT_P, 0)
}

func (s *CreatetableContext) AllColumns() []IColumnsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IColumnsContext); ok {
			len++
		}
	}

	tst := make([]IColumnsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IColumnsContext); ok {
			tst[i] = t.(IColumnsContext)
			i++
		}
	}

	return tst
}

func (s *CreatetableContext) Columns(i int) IColumnsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumnsContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumnsContext)
}

func (s *CreatetableContext) RIGHT_P() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserRIGHT_P, 0)
}

func (s *CreatetableContext) AW() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserAW, 0)
}

func (s *CreatetableContext) RW() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserRW, 0)
}

func (s *CreatetableContext) LWW() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserLWW, 0)
}

func (s *CreatetableContext) AllSEPARATOR() []antlr.TerminalNode {
	return s.GetTokens(ViewSQLParserSEPARATOR)
}

func (s *CreatetableContext) SEPARATOR(i int) antlr.TerminalNode {
	return s.GetToken(ViewSQLParserSEPARATOR, i)
}

func (s *CreatetableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CreatetableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CreatetableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.EnterCreatetable(s)
	}
}

func (s *CreatetableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.ExitCreatetable(s)
	}
}

func (p *ViewSQLParser) Createtable() (localctx ICreatetableContext) {
	this := p
	_ = this

	localctx = NewCreatetableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, ViewSQLParserRULE_createtable)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(334)
		p.Match(ViewSQLParserCREATE)
	}
	{
		p.SetState(335)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*CreatetableContext).policy = _lt

		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&504403158265495552) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*CreatetableContext).policy = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(336)
		p.Match(ViewSQLParserTABLE)
	}
	{
		p.SetState(337)
		p.Name()
	}
	{
		p.SetState(338)
		p.Match(ViewSQLParserLEFT_P)
	}
	p.SetState(344)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(339)
				p.Columns()
			}
			{
				p.SetState(340)
				p.Match(ViewSQLParserSEPARATOR)
			}

		}
		p.SetState(346)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext())
	}
	{
		p.SetState(347)
		p.Columns()
	}
	{
		p.SetState(348)
		p.Match(ViewSQLParserRIGHT_P)
	}

	return localctx
}

// ICreateindexContext is an interface to support dynamic dispatch.
type ICreateindexContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetIndexName returns the indexName rule contexts.
	GetIndexName() INameContext

	// GetTableName returns the tableName rule contexts.
	GetTableName() INameContext

	// GetColumnName returns the columnName rule contexts.
	GetColumnName() INameContext

	// SetIndexName sets the indexName rule contexts.
	SetIndexName(INameContext)

	// SetTableName sets the tableName rule contexts.
	SetTableName(INameContext)

	// SetColumnName sets the columnName rule contexts.
	SetColumnName(INameContext)

	// Getter signatures
	CREATE() antlr.TerminalNode
	INDEX() antlr.TerminalNode
	ON() antlr.TerminalNode
	LEFT_P() antlr.TerminalNode
	RIGHT_P() antlr.TerminalNode
	AllName() []INameContext
	Name(i int) INameContext

	// IsCreateindexContext differentiates from other interfaces.
	IsCreateindexContext()
}

type CreateindexContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	indexName  INameContext
	tableName  INameContext
	columnName INameContext
}

func NewEmptyCreateindexContext() *CreateindexContext {
	var p = new(CreateindexContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ViewSQLParserRULE_createindex
	return p
}

func (*CreateindexContext) IsCreateindexContext() {}

func NewCreateindexContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CreateindexContext {
	var p = new(CreateindexContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ViewSQLParserRULE_createindex

	return p
}

func (s *CreateindexContext) GetParser() antlr.Parser { return s.parser }

func (s *CreateindexContext) GetIndexName() INameContext { return s.indexName }

func (s *CreateindexContext) GetTableName() INameContext { return s.tableName }

func (s *CreateindexContext) GetColumnName() INameContext { return s.columnName }

func (s *CreateindexContext) SetIndexName(v INameContext) { s.indexName = v }

func (s *CreateindexContext) SetTableName(v INameContext) { s.tableName = v }

func (s *CreateindexContext) SetColumnName(v INameContext) { s.columnName = v }

func (s *CreateindexContext) CREATE() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserCREATE, 0)
}

func (s *CreateindexContext) INDEX() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserINDEX, 0)
}

func (s *CreateindexContext) ON() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserON, 0)
}

func (s *CreateindexContext) LEFT_P() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserLEFT_P, 0)
}

func (s *CreateindexContext) RIGHT_P() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserRIGHT_P, 0)
}

func (s *CreateindexContext) AllName() []INameContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INameContext); ok {
			len++
		}
	}

	tst := make([]INameContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INameContext); ok {
			tst[i] = t.(INameContext)
			i++
		}
	}

	return tst
}

func (s *CreateindexContext) Name(i int) INameContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *CreateindexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CreateindexContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CreateindexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.EnterCreateindex(s)
	}
}

func (s *CreateindexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.ExitCreateindex(s)
	}
}

func (p *ViewSQLParser) Createindex() (localctx ICreateindexContext) {
	this := p
	_ = this

	localctx = NewCreateindexContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, ViewSQLParserRULE_createindex)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(350)
		p.Match(ViewSQLParserCREATE)
	}
	{
		p.SetState(351)
		p.Match(ViewSQLParserINDEX)
	}
	{
		p.SetState(352)

		var _x = p.Name()

		localctx.(*CreateindexContext).indexName = _x
	}
	{
		p.SetState(353)
		p.Match(ViewSQLParserON)
	}
	{
		p.SetState(354)

		var _x = p.Name()

		localctx.(*CreateindexContext).tableName = _x
	}
	{
		p.SetState(355)
		p.Match(ViewSQLParserLEFT_P)
	}
	{
		p.SetState(356)

		var _x = p.Name()

		localctx.(*CreateindexContext).columnName = _x
	}
	{
		p.SetState(357)
		p.Match(ViewSQLParserRIGHT_P)
	}

	return localctx
}

// IDropContext is an interface to support dynamic dispatch.
type IDropContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DROP() antlr.TerminalNode
	TABLE() antlr.TerminalNode
	Name() INameContext

	// IsDropContext differentiates from other interfaces.
	IsDropContext()
}

type DropContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDropContext() *DropContext {
	var p = new(DropContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ViewSQLParserRULE_drop
	return p
}

func (*DropContext) IsDropContext() {}

func NewDropContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DropContext {
	var p = new(DropContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ViewSQLParserRULE_drop

	return p
}

func (s *DropContext) GetParser() antlr.Parser { return s.parser }

func (s *DropContext) DROP() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserDROP, 0)
}

func (s *DropContext) TABLE() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserTABLE, 0)
}

func (s *DropContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *DropContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DropContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DropContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.EnterDrop(s)
	}
}

func (s *DropContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.ExitDrop(s)
	}
}

func (p *ViewSQLParser) Drop() (localctx IDropContext) {
	this := p
	_ = this

	localctx = NewDropContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, ViewSQLParserRULE_drop)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(359)
		p.Match(ViewSQLParserDROP)
	}
	{
		p.SetState(360)
		p.Match(ViewSQLParserTABLE)
	}
	{
		p.SetState(361)
		p.Name()
	}

	return localctx
}

// IDeleteContext is an interface to support dynamic dispatch.
type IDeleteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DELETE() antlr.TerminalNode
	FROM() antlr.TerminalNode
	Name() INameContext
	Where() IWhereContext

	// IsDeleteContext differentiates from other interfaces.
	IsDeleteContext()
}

type DeleteContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeleteContext() *DeleteContext {
	var p = new(DeleteContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ViewSQLParserRULE_delete
	return p
}

func (*DeleteContext) IsDeleteContext() {}

func NewDeleteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeleteContext {
	var p = new(DeleteContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ViewSQLParserRULE_delete

	return p
}

func (s *DeleteContext) GetParser() antlr.Parser { return s.parser }

func (s *DeleteContext) DELETE() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserDELETE, 0)
}

func (s *DeleteContext) FROM() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserFROM, 0)
}

func (s *DeleteContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *DeleteContext) Where() IWhereContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhereContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhereContext)
}

func (s *DeleteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeleteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeleteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.EnterDelete(s)
	}
}

func (s *DeleteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.ExitDelete(s)
	}
}

func (p *ViewSQLParser) Delete_() (localctx IDeleteContext) {
	this := p
	_ = this

	localctx = NewDeleteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, ViewSQLParserRULE_delete)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(363)
		p.Match(ViewSQLParserDELETE)
	}
	{
		p.SetState(364)
		p.Match(ViewSQLParserFROM)
	}
	{
		p.SetState(365)
		p.Name()
	}
	{
		p.SetState(366)
		p.Where()
	}

	return localctx
}

// ISetContext is an interface to support dynamic dispatch.
type ISetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SET() antlr.TerminalNode
	AllName() []INameContext
	Name(i int) INameContext
	AllEQUAL() []antlr.TerminalNode
	EQUAL(i int) antlr.TerminalNode
	AllConstant() []IConstantContext
	Constant(i int) IConstantContext
	Where() IWhereContext
	AllSEPARATOR() []antlr.TerminalNode
	SEPARATOR(i int) antlr.TerminalNode

	// IsSetContext differentiates from other interfaces.
	IsSetContext()
}

type SetContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySetContext() *SetContext {
	var p = new(SetContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ViewSQLParserRULE_set
	return p
}

func (*SetContext) IsSetContext() {}

func NewSetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SetContext {
	var p = new(SetContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ViewSQLParserRULE_set

	return p
}

func (s *SetContext) GetParser() antlr.Parser { return s.parser }

func (s *SetContext) SET() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserSET, 0)
}

func (s *SetContext) AllName() []INameContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INameContext); ok {
			len++
		}
	}

	tst := make([]INameContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INameContext); ok {
			tst[i] = t.(INameContext)
			i++
		}
	}

	return tst
}

func (s *SetContext) Name(i int) INameContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *SetContext) AllEQUAL() []antlr.TerminalNode {
	return s.GetTokens(ViewSQLParserEQUAL)
}

func (s *SetContext) EQUAL(i int) antlr.TerminalNode {
	return s.GetToken(ViewSQLParserEQUAL, i)
}

func (s *SetContext) AllConstant() []IConstantContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConstantContext); ok {
			len++
		}
	}

	tst := make([]IConstantContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConstantContext); ok {
			tst[i] = t.(IConstantContext)
			i++
		}
	}

	return tst
}

func (s *SetContext) Constant(i int) IConstantContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *SetContext) Where() IWhereContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhereContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhereContext)
}

func (s *SetContext) AllSEPARATOR() []antlr.TerminalNode {
	return s.GetTokens(ViewSQLParserSEPARATOR)
}

func (s *SetContext) SEPARATOR(i int) antlr.TerminalNode {
	return s.GetToken(ViewSQLParserSEPARATOR, i)
}

func (s *SetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.EnterSet(s)
	}
}

func (s *SetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.ExitSet(s)
	}
}

func (p *ViewSQLParser) Set() (localctx ISetContext) {
	this := p
	_ = this

	localctx = NewSetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, ViewSQLParserRULE_set)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(368)
		p.Match(ViewSQLParserSET)
	}
	p.SetState(376)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(369)
				p.Name()
			}
			{
				p.SetState(370)
				p.Match(ViewSQLParserEQUAL)
			}
			{
				p.SetState(371)
				p.Constant()
			}
			{
				p.SetState(372)
				p.Match(ViewSQLParserSEPARATOR)
			}

		}
		p.SetState(378)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext())
	}
	{
		p.SetState(379)
		p.Name()
	}
	{
		p.SetState(380)
		p.Match(ViewSQLParserEQUAL)
	}
	{
		p.SetState(381)
		p.Constant()
	}
	{
		p.SetState(382)
		p.Where()
	}

	return localctx
}

// IUpdateContext is an interface to support dynamic dispatch.
type IUpdateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	UPDATE() antlr.TerminalNode
	Name() INameContext
	Set() ISetContext

	// IsUpdateContext differentiates from other interfaces.
	IsUpdateContext()
}

type UpdateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUpdateContext() *UpdateContext {
	var p = new(UpdateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ViewSQLParserRULE_update
	return p
}

func (*UpdateContext) IsUpdateContext() {}

func NewUpdateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UpdateContext {
	var p = new(UpdateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ViewSQLParserRULE_update

	return p
}

func (s *UpdateContext) GetParser() antlr.Parser { return s.parser }

func (s *UpdateContext) UPDATE() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserUPDATE, 0)
}

func (s *UpdateContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *UpdateContext) Set() ISetContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISetContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISetContext)
}

func (s *UpdateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UpdateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UpdateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.EnterUpdate(s)
	}
}

func (s *UpdateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.ExitUpdate(s)
	}
}

func (p *ViewSQLParser) Update() (localctx IUpdateContext) {
	this := p
	_ = this

	localctx = NewUpdateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, ViewSQLParserRULE_update)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(384)
		p.Match(ViewSQLParserUPDATE)
	}
	{
		p.SetState(385)
		p.Name()
	}
	{
		p.SetState(386)
		p.Set()
	}

	return localctx
}

// IValuesContext is an interface to support dynamic dispatch.
type IValuesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VALUES() antlr.TerminalNode
	LEFT_P() antlr.TerminalNode
	AllConstant() []IConstantContext
	Constant(i int) IConstantContext
	RIGHT_P() antlr.TerminalNode
	AllSEPARATOR() []antlr.TerminalNode
	SEPARATOR(i int) antlr.TerminalNode

	// IsValuesContext differentiates from other interfaces.
	IsValuesContext()
}

type ValuesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValuesContext() *ValuesContext {
	var p = new(ValuesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ViewSQLParserRULE_values
	return p
}

func (*ValuesContext) IsValuesContext() {}

func NewValuesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValuesContext {
	var p = new(ValuesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ViewSQLParserRULE_values

	return p
}

func (s *ValuesContext) GetParser() antlr.Parser { return s.parser }

func (s *ValuesContext) VALUES() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserVALUES, 0)
}

func (s *ValuesContext) LEFT_P() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserLEFT_P, 0)
}

func (s *ValuesContext) AllConstant() []IConstantContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConstantContext); ok {
			len++
		}
	}

	tst := make([]IConstantContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConstantContext); ok {
			tst[i] = t.(IConstantContext)
			i++
		}
	}

	return tst
}

func (s *ValuesContext) Constant(i int) IConstantContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *ValuesContext) RIGHT_P() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserRIGHT_P, 0)
}

func (s *ValuesContext) AllSEPARATOR() []antlr.TerminalNode {
	return s.GetTokens(ViewSQLParserSEPARATOR)
}

func (s *ValuesContext) SEPARATOR(i int) antlr.TerminalNode {
	return s.GetToken(ViewSQLParserSEPARATOR, i)
}

func (s *ValuesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValuesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValuesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.EnterValues(s)
	}
}

func (s *ValuesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.ExitValues(s)
	}
}

func (p *ViewSQLParser) Values() (localctx IValuesContext) {
	this := p
	_ = this

	localctx = NewValuesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, ViewSQLParserRULE_values)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(388)
		p.Match(ViewSQLParserVALUES)
	}
	{
		p.SetState(389)
		p.Match(ViewSQLParserLEFT_P)
	}
	p.SetState(395)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(390)
				p.Constant()
			}
			{
				p.SetState(391)
				p.Match(ViewSQLParserSEPARATOR)
			}

		}
		p.SetState(397)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext())
	}
	{
		p.SetState(398)
		p.Constant()
	}
	{
		p.SetState(399)
		p.Match(ViewSQLParserRIGHT_P)
	}

	return localctx
}

// IColumnNamesContext is an interface to support dynamic dispatch.
type IColumnNamesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LEFT_P() antlr.TerminalNode
	AllName() []INameContext
	Name(i int) INameContext
	RIGHT_P() antlr.TerminalNode
	AllSEPARATOR() []antlr.TerminalNode
	SEPARATOR(i int) antlr.TerminalNode

	// IsColumnNamesContext differentiates from other interfaces.
	IsColumnNamesContext()
}

type ColumnNamesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumnNamesContext() *ColumnNamesContext {
	var p = new(ColumnNamesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ViewSQLParserRULE_columnNames
	return p
}

func (*ColumnNamesContext) IsColumnNamesContext() {}

func NewColumnNamesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColumnNamesContext {
	var p = new(ColumnNamesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ViewSQLParserRULE_columnNames

	return p
}

func (s *ColumnNamesContext) GetParser() antlr.Parser { return s.parser }

func (s *ColumnNamesContext) LEFT_P() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserLEFT_P, 0)
}

func (s *ColumnNamesContext) AllName() []INameContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INameContext); ok {
			len++
		}
	}

	tst := make([]INameContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INameContext); ok {
			tst[i] = t.(INameContext)
			i++
		}
	}

	return tst
}

func (s *ColumnNamesContext) Name(i int) INameContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *ColumnNamesContext) RIGHT_P() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserRIGHT_P, 0)
}

func (s *ColumnNamesContext) AllSEPARATOR() []antlr.TerminalNode {
	return s.GetTokens(ViewSQLParserSEPARATOR)
}

func (s *ColumnNamesContext) SEPARATOR(i int) antlr.TerminalNode {
	return s.GetToken(ViewSQLParserSEPARATOR, i)
}

func (s *ColumnNamesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnNamesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ColumnNamesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.EnterColumnNames(s)
	}
}

func (s *ColumnNamesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.ExitColumnNames(s)
	}
}

func (p *ViewSQLParser) ColumnNames() (localctx IColumnNamesContext) {
	this := p
	_ = this

	localctx = NewColumnNamesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, ViewSQLParserRULE_columnNames)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(401)
		p.Match(ViewSQLParserLEFT_P)
	}
	p.SetState(407)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(402)
				p.Name()
			}
			{
				p.SetState(403)
				p.Match(ViewSQLParserSEPARATOR)
			}

		}
		p.SetState(409)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext())
	}
	{
		p.SetState(410)
		p.Name()
	}
	{
		p.SetState(411)
		p.Match(ViewSQLParserRIGHT_P)
	}

	return localctx
}

// IInsertContext is an interface to support dynamic dispatch.
type IInsertContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INSERT() antlr.TerminalNode
	INTO() antlr.TerminalNode
	Name() INameContext
	Values() IValuesContext
	ColumnNames() IColumnNamesContext

	// IsInsertContext differentiates from other interfaces.
	IsInsertContext()
}

type InsertContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInsertContext() *InsertContext {
	var p = new(InsertContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ViewSQLParserRULE_insert
	return p
}

func (*InsertContext) IsInsertContext() {}

func NewInsertContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InsertContext {
	var p = new(InsertContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ViewSQLParserRULE_insert

	return p
}

func (s *InsertContext) GetParser() antlr.Parser { return s.parser }

func (s *InsertContext) INSERT() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserINSERT, 0)
}

func (s *InsertContext) INTO() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserINTO, 0)
}

func (s *InsertContext) Name() INameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *InsertContext) Values() IValuesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValuesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValuesContext)
}

func (s *InsertContext) ColumnNames() IColumnNamesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumnNamesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumnNamesContext)
}

func (s *InsertContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InsertContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InsertContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.EnterInsert(s)
	}
}

func (s *InsertContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.ExitInsert(s)
	}
}

func (p *ViewSQLParser) Insert() (localctx IInsertContext) {
	this := p
	_ = this

	localctx = NewInsertContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, ViewSQLParserRULE_insert)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(413)
		p.Match(ViewSQLParserINSERT)
	}
	{
		p.SetState(414)
		p.Match(ViewSQLParserINTO)
	}
	{
		p.SetState(415)
		p.Name()
	}
	p.SetState(417)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ViewSQLParserLEFT_P {
		{
			p.SetState(416)
			p.ColumnNames()
		}

	}
	{
		p.SetState(419)
		p.Values()
	}

	return localctx
}

// IQueryContext is an interface to support dynamic dispatch.
type IQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Select_() ISelectContext
	From() IFromContext
	Where() IWhereContext

	// IsQueryContext differentiates from other interfaces.
	IsQueryContext()
}

type QueryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQueryContext() *QueryContext {
	var p = new(QueryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ViewSQLParserRULE_query
	return p
}

func (*QueryContext) IsQueryContext() {}

func NewQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QueryContext {
	var p = new(QueryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ViewSQLParserRULE_query

	return p
}

func (s *QueryContext) GetParser() antlr.Parser { return s.parser }

func (s *QueryContext) Select_() ISelectContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectContext)
}

func (s *QueryContext) From() IFromContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFromContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFromContext)
}

func (s *QueryContext) Where() IWhereContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhereContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhereContext)
}

func (s *QueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QueryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.EnterQuery(s)
	}
}

func (s *QueryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.ExitQuery(s)
	}
}

func (p *ViewSQLParser) Query() (localctx IQueryContext) {
	this := p
	_ = this

	localctx = NewQueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, ViewSQLParserRULE_query)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(421)
		p.Select_()
	}
	{
		p.SetState(422)
		p.From()
	}
	p.SetState(424)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ViewSQLParserWHERE {
		{
			p.SetState(423)
			p.Where()
		}

	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	View() IViewContext
	Createtable() ICreatetableContext
	Createindex() ICreateindexContext
	Insert() IInsertContext
	Update() IUpdateContext
	Delete_() IDeleteContext
	Drop() IDropContext
	Query() IQueryContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ViewSQLParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ViewSQLParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) View() IViewContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IViewContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IViewContext)
}

func (s *StatementContext) Createtable() ICreatetableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICreatetableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICreatetableContext)
}

func (s *StatementContext) Createindex() ICreateindexContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICreateindexContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICreateindexContext)
}

func (s *StatementContext) Insert() IInsertContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInsertContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInsertContext)
}

func (s *StatementContext) Update() IUpdateContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUpdateContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUpdateContext)
}

func (s *StatementContext) Delete_() IDeleteContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeleteContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeleteContext)
}

func (s *StatementContext) Drop() IDropContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDropContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDropContext)
}

func (s *StatementContext) Query() IQueryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQueryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQueryContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *ViewSQLParser) Statement() (localctx IStatementContext) {
	this := p
	_ = this

	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, ViewSQLParserRULE_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(434)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(426)
			p.View()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(427)
			p.Createtable()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(428)
			p.Createindex()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(429)
			p.Insert()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(430)
			p.Update()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(431)
			p.Delete_()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(432)
			p.Drop()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(433)
			p.Query()
		}

	}

	return localctx
}

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Statement() IStatementContext
	EOF() antlr.TerminalNode

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ViewSQLParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ViewSQLParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) Statement() IStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(ViewSQLParserEOF, 0)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ViewSQLListener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *ViewSQLParser) Start() (localctx IStartContext) {
	this := p
	_ = this

	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, ViewSQLParserRULE_start)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(436)
		p.Statement()
	}
	{
		p.SetState(437)
		p.Match(ViewSQLParserEOF)
	}

	return localctx
}

func (p *ViewSQLParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 7:
		var t *MathContext = nil
		if localctx != nil {
			t = localctx.(*MathContext)
		}
		return p.Math_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *ViewSQLParser) Math_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
