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
		"'KEY'",
	}
	staticData.symbolicNames = []string{
		"", "ADD", "SUB", "MULT", "DIV", "SUM", "AVG", "MAX", "MIN", "COUNT",
		"AND", "ASC", "DESC", "EQUAL", "HIGHER", "HIGHER_EQUAL", "LOWER", "LOWER_EQUAL",
		"NOT_EQUAL", "DOT", "SEPARATOR", "RANGE_SEP", "LEFT_P", "RIGHT_P", "LEFT_SP",
		"RIGHT_SP", "INV_COMMA", "LEFT_CURLY", "RIGHT_CURLY", "AS", "CREATE",
		"VIEW", "SELECT", "FROM", "WHERE", "GROUP", "BY", "ORDER", "LIMIT",
		"WITH", "IN", "PRIMARY", "KEY", "DATE", "STRING", "INT", "FLOAT", "WHITESPACE",
	}
	staticData.ruleNames = []string{
		"name", "constant", "aggrFunc", "field", "parameter", "nameable", "key",
		"math", "asClause", "aggregation", "count", "calc", "comp", "condition",
		"sortOrder", "continuousRange", "sparseRange", "range", "create", "with",
		"select", "from", "where", "groupby", "orderby", "limit", "view", "start",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 47, 275, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5,
		1, 65, 8, 1, 10, 1, 12, 1, 68, 9, 1, 1, 1, 3, 1, 71, 8, 1, 1, 2, 1, 2,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 3, 5,
		86, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 3, 7, 97,
		8, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 105, 8, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 113, 8, 7, 10, 7, 12, 7, 116, 9, 7, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10,
		1, 10, 1, 10, 3, 10, 133, 8, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1,
		11, 1, 11, 1, 11, 1, 11, 3, 11, 144, 8, 11, 1, 12, 1, 12, 3, 12, 148, 8,
		12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 5, 16, 167, 8, 16, 10,
		16, 12, 16, 170, 9, 16, 1, 16, 1, 16, 1, 17, 1, 17, 3, 17, 176, 8, 17,
		1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1,
		19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 5, 19, 194, 8, 19, 10, 19, 12, 19,
		197, 9, 19, 1, 20, 1, 20, 1, 20, 1, 20, 5, 20, 203, 8, 20, 10, 20, 12,
		20, 206, 9, 20, 1, 21, 1, 21, 1, 21, 1, 21, 5, 21, 212, 8, 21, 10, 21,
		12, 21, 215, 9, 21, 1, 22, 1, 22, 1, 22, 1, 22, 5, 22, 221, 8, 22, 10,
		22, 12, 22, 224, 9, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 5, 23, 231,
		8, 23, 10, 23, 12, 23, 234, 9, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1,
		24, 1, 24, 1, 24, 5, 24, 244, 8, 24, 10, 24, 12, 24, 247, 9, 24, 1, 25,
		1, 25, 1, 25, 1, 26, 1, 26, 3, 26, 254, 8, 26, 1, 26, 1, 26, 1, 26, 1,
		26, 3, 26, 260, 8, 26, 1, 26, 3, 26, 263, 8, 26, 1, 26, 1, 26, 3, 26, 267,
		8, 26, 1, 26, 3, 26, 270, 8, 26, 1, 27, 1, 27, 1, 27, 1, 27, 0, 1, 14,
		28, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34,
		36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 0, 7, 2, 0, 44, 44, 47, 47, 1,
		0, 5, 8, 1, 0, 3, 4, 1, 0, 1, 2, 1, 0, 13, 18, 1, 0, 11, 12, 1, 0, 24,
		25, 276, 0, 56, 1, 0, 0, 0, 2, 70, 1, 0, 0, 0, 4, 72, 1, 0, 0, 0, 6, 74,
		1, 0, 0, 0, 8, 78, 1, 0, 0, 0, 10, 85, 1, 0, 0, 0, 12, 87, 1, 0, 0, 0,
		14, 104, 1, 0, 0, 0, 16, 117, 1, 0, 0, 0, 18, 121, 1, 0, 0, 0, 20, 128,
		1, 0, 0, 0, 22, 143, 1, 0, 0, 0, 24, 147, 1, 0, 0, 0, 26, 149, 1, 0, 0,
		0, 28, 153, 1, 0, 0, 0, 30, 155, 1, 0, 0, 0, 32, 162, 1, 0, 0, 0, 34, 175,
		1, 0, 0, 0, 36, 177, 1, 0, 0, 0, 38, 185, 1, 0, 0, 0, 40, 198, 1, 0, 0,
		0, 42, 207, 1, 0, 0, 0, 44, 216, 1, 0, 0, 0, 46, 225, 1, 0, 0, 0, 48, 235,
		1, 0, 0, 0, 50, 248, 1, 0, 0, 0, 52, 251, 1, 0, 0, 0, 54, 271, 1, 0, 0,
		0, 56, 57, 5, 44, 0, 0, 57, 1, 1, 0, 0, 0, 58, 71, 5, 43, 0, 0, 59, 71,
		5, 45, 0, 0, 60, 71, 5, 46, 0, 0, 61, 62, 5, 26, 0, 0, 62, 66, 5, 44, 0,
		0, 63, 65, 7, 0, 0, 0, 64, 63, 1, 0, 0, 0, 65, 68, 1, 0, 0, 0, 66, 64,
		1, 0, 0, 0, 66, 67, 1, 0, 0, 0, 67, 69, 1, 0, 0, 0, 68, 66, 1, 0, 0, 0,
		69, 71, 5, 26, 0, 0, 70, 58, 1, 0, 0, 0, 70, 59, 1, 0, 0, 0, 70, 60, 1,
		0, 0, 0, 70, 61, 1, 0, 0, 0, 71, 3, 1, 0, 0, 0, 72, 73, 7, 1, 0, 0, 73,
		5, 1, 0, 0, 0, 74, 75, 3, 0, 0, 0, 75, 76, 5, 19, 0, 0, 76, 77, 3, 0, 0,
		0, 77, 7, 1, 0, 0, 0, 78, 79, 5, 24, 0, 0, 79, 80, 3, 0, 0, 0, 80, 81,
		5, 25, 0, 0, 81, 9, 1, 0, 0, 0, 82, 86, 3, 0, 0, 0, 83, 86, 3, 6, 3, 0,
		84, 86, 3, 8, 4, 0, 85, 82, 1, 0, 0, 0, 85, 83, 1, 0, 0, 0, 85, 84, 1,
		0, 0, 0, 86, 11, 1, 0, 0, 0, 87, 88, 5, 41, 0, 0, 88, 89, 5, 42, 0, 0,
		89, 90, 5, 22, 0, 0, 90, 91, 3, 10, 5, 0, 91, 92, 5, 23, 0, 0, 92, 13,
		1, 0, 0, 0, 93, 96, 6, 7, -1, 0, 94, 97, 3, 10, 5, 0, 95, 97, 3, 2, 1,
		0, 96, 94, 1, 0, 0, 0, 96, 95, 1, 0, 0, 0, 97, 105, 1, 0, 0, 0, 98, 99,
		5, 2, 0, 0, 99, 105, 3, 14, 7, 4, 100, 101, 5, 22, 0, 0, 101, 102, 3, 14,
		7, 0, 102, 103, 5, 23, 0, 0, 103, 105, 1, 0, 0, 0, 104, 93, 1, 0, 0, 0,
		104, 98, 1, 0, 0, 0, 104, 100, 1, 0, 0, 0, 105, 114, 1, 0, 0, 0, 106, 107,
		10, 2, 0, 0, 107, 108, 7, 2, 0, 0, 108, 113, 3, 14, 7, 3, 109, 110, 10,
		1, 0, 0, 110, 111, 7, 3, 0, 0, 111, 113, 3, 14, 7, 2, 112, 106, 1, 0, 0,
		0, 112, 109, 1, 0, 0, 0, 113, 116, 1, 0, 0, 0, 114, 112, 1, 0, 0, 0, 114,
		115, 1, 0, 0, 0, 115, 15, 1, 0, 0, 0, 116, 114, 1, 0, 0, 0, 117, 118, 3,
		14, 7, 0, 118, 119, 5, 29, 0, 0, 119, 120, 3, 0, 0, 0, 120, 17, 1, 0, 0,
		0, 121, 122, 3, 4, 2, 0, 122, 123, 5, 22, 0, 0, 123, 124, 3, 14, 7, 0,
		124, 125, 5, 23, 0, 0, 125, 126, 5, 29, 0, 0, 126, 127, 3, 0, 0, 0, 127,
		19, 1, 0, 0, 0, 128, 129, 5, 9, 0, 0, 129, 132, 5, 22, 0, 0, 130, 133,
		3, 10, 5, 0, 131, 133, 5, 3, 0, 0, 132, 130, 1, 0, 0, 0, 132, 131, 1, 0,
		0, 0, 133, 134, 1, 0, 0, 0, 134, 135, 5, 23, 0, 0, 135, 136, 5, 29, 0,
		0, 136, 137, 3, 0, 0, 0, 137, 21, 1, 0, 0, 0, 138, 144, 3, 12, 6, 0, 139,
		144, 3, 10, 5, 0, 140, 144, 3, 16, 8, 0, 141, 144, 3, 18, 9, 0, 142, 144,
		3, 20, 10, 0, 143, 138, 1, 0, 0, 0, 143, 139, 1, 0, 0, 0, 143, 140, 1,
		0, 0, 0, 143, 141, 1, 0, 0, 0, 143, 142, 1, 0, 0, 0, 144, 23, 1, 0, 0,
		0, 145, 148, 7, 4, 0, 0, 146, 148, 3, 0, 0, 0, 147, 145, 1, 0, 0, 0, 147,
		146, 1, 0, 0, 0, 148, 25, 1, 0, 0, 0, 149, 150, 3, 10, 5, 0, 150, 151,
		3, 24, 12, 0, 151, 152, 3, 14, 7, 0, 152, 27, 1, 0, 0, 0, 153, 154, 7,
		5, 0, 0, 154, 29, 1, 0, 0, 0, 155, 156, 7, 6, 0, 0, 156, 157, 3, 2, 1,
		0, 157, 158, 5, 21, 0, 0, 158, 159, 3, 2, 1, 0, 159, 160, 1, 0, 0, 0, 160,
		161, 7, 6, 0, 0, 161, 31, 1, 0, 0, 0, 162, 163, 5, 27, 0, 0, 163, 168,
		3, 2, 1, 0, 164, 165, 5, 20, 0, 0, 165, 167, 3, 2, 1, 0, 166, 164, 1, 0,
		0, 0, 167, 170, 1, 0, 0, 0, 168, 166, 1, 0, 0, 0, 168, 169, 1, 0, 0, 0,
		169, 171, 1, 0, 0, 0, 170, 168, 1, 0, 0, 0, 171, 172, 5, 28, 0, 0, 172,
		33, 1, 0, 0, 0, 173, 176, 3, 30, 15, 0, 174, 176, 3, 32, 16, 0, 175, 173,
		1, 0, 0, 0, 175, 174, 1, 0, 0, 0, 176, 35, 1, 0, 0, 0, 177, 178, 5, 30,
		0, 0, 178, 179, 5, 31, 0, 0, 179, 180, 5, 22, 0, 0, 180, 181, 5, 44, 0,
		0, 181, 182, 5, 20, 0, 0, 182, 183, 5, 44, 0, 0, 183, 184, 5, 23, 0, 0,
		184, 37, 1, 0, 0, 0, 185, 186, 5, 39, 0, 0, 186, 187, 5, 44, 0, 0, 187,
		188, 5, 40, 0, 0, 188, 195, 3, 34, 17, 0, 189, 190, 5, 20, 0, 0, 190, 191,
		5, 44, 0, 0, 191, 192, 5, 40, 0, 0, 192, 194, 3, 34, 17, 0, 193, 189, 1,
		0, 0, 0, 194, 197, 1, 0, 0, 0, 195, 193, 1, 0, 0, 0, 195, 196, 1, 0, 0,
		0, 196, 39, 1, 0, 0, 0, 197, 195, 1, 0, 0, 0, 198, 199, 5, 32, 0, 0, 199,
		204, 3, 22, 11, 0, 200, 201, 5, 20, 0, 0, 201, 203, 3, 22, 11, 0, 202,
		200, 1, 0, 0, 0, 203, 206, 1, 0, 0, 0, 204, 202, 1, 0, 0, 0, 204, 205,
		1, 0, 0, 0, 205, 41, 1, 0, 0, 0, 206, 204, 1, 0, 0, 0, 207, 208, 5, 33,
		0, 0, 208, 213, 3, 0, 0, 0, 209, 210, 5, 20, 0, 0, 210, 212, 3, 0, 0, 0,
		211, 209, 1, 0, 0, 0, 212, 215, 1, 0, 0, 0, 213, 211, 1, 0, 0, 0, 213,
		214, 1, 0, 0, 0, 214, 43, 1, 0, 0, 0, 215, 213, 1, 0, 0, 0, 216, 217, 5,
		34, 0, 0, 217, 222, 3, 26, 13, 0, 218, 219, 5, 10, 0, 0, 219, 221, 3, 26,
		13, 0, 220, 218, 1, 0, 0, 0, 221, 224, 1, 0, 0, 0, 222, 220, 1, 0, 0, 0,
		222, 223, 1, 0, 0, 0, 223, 45, 1, 0, 0, 0, 224, 222, 1, 0, 0, 0, 225, 226,
		5, 35, 0, 0, 226, 227, 5, 36, 0, 0, 227, 232, 3, 10, 5, 0, 228, 229, 5,
		20, 0, 0, 229, 231, 3, 10, 5, 0, 230, 228, 1, 0, 0, 0, 231, 234, 1, 0,
		0, 0, 232, 230, 1, 0, 0, 0, 232, 233, 1, 0, 0, 0, 233, 47, 1, 0, 0, 0,
		234, 232, 1, 0, 0, 0, 235, 236, 5, 37, 0, 0, 236, 237, 5, 36, 0, 0, 237,
		238, 3, 10, 5, 0, 238, 245, 3, 28, 14, 0, 239, 240, 5, 20, 0, 0, 240, 241,
		3, 10, 5, 0, 241, 242, 3, 28, 14, 0, 242, 244, 1, 0, 0, 0, 243, 239, 1,
		0, 0, 0, 244, 247, 1, 0, 0, 0, 245, 243, 1, 0, 0, 0, 245, 246, 1, 0, 0,
		0, 246, 49, 1, 0, 0, 0, 247, 245, 1, 0, 0, 0, 248, 249, 5, 38, 0, 0, 249,
		250, 5, 45, 0, 0, 250, 51, 1, 0, 0, 0, 251, 253, 3, 36, 18, 0, 252, 254,
		3, 38, 19, 0, 253, 252, 1, 0, 0, 0, 253, 254, 1, 0, 0, 0, 254, 255, 1,
		0, 0, 0, 255, 256, 5, 29, 0, 0, 256, 257, 3, 40, 20, 0, 257, 259, 3, 42,
		21, 0, 258, 260, 3, 44, 22, 0, 259, 258, 1, 0, 0, 0, 259, 260, 1, 0, 0,
		0, 260, 262, 1, 0, 0, 0, 261, 263, 3, 46, 23, 0, 262, 261, 1, 0, 0, 0,
		262, 263, 1, 0, 0, 0, 263, 264, 1, 0, 0, 0, 264, 266, 3, 46, 23, 0, 265,
		267, 3, 48, 24, 0, 266, 265, 1, 0, 0, 0, 266, 267, 1, 0, 0, 0, 267, 269,
		1, 0, 0, 0, 268, 270, 3, 50, 25, 0, 269, 268, 1, 0, 0, 0, 269, 270, 1,
		0, 0, 0, 270, 53, 1, 0, 0, 0, 271, 272, 3, 52, 26, 0, 272, 273, 5, 0, 0,
		1, 273, 55, 1, 0, 0, 0, 23, 66, 70, 85, 96, 104, 112, 114, 132, 143, 147,
		168, 175, 195, 204, 213, 222, 232, 245, 253, 259, 262, 266, 269,
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
	ViewSQLParserDATE         = 43
	ViewSQLParserSTRING       = 44
	ViewSQLParserINT          = 45
	ViewSQLParserFLOAT        = 46
	ViewSQLParserWHITESPACE   = 47
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
	ViewSQLParserRULE_start           = 27
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
		p.SetState(56)
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

	p.SetState(70)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ViewSQLParserDATE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(58)
			p.Match(ViewSQLParserDATE)
		}

	case ViewSQLParserINT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(59)
			p.Match(ViewSQLParserINT)
		}

	case ViewSQLParserFLOAT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(60)
			p.Match(ViewSQLParserFLOAT)
		}

	case ViewSQLParserINV_COMMA:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(61)
			p.Match(ViewSQLParserINV_COMMA)
		}
		{
			p.SetState(62)
			p.Match(ViewSQLParserSTRING)
		}
		p.SetState(66)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ViewSQLParserSTRING || _la == ViewSQLParserWHITESPACE {
			{
				p.SetState(63)
				_la = p.GetTokenStream().LA(1)

				if !(_la == ViewSQLParserSTRING || _la == ViewSQLParserWHITESPACE) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

			p.SetState(68)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(69)
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
		p.SetState(72)
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
		p.SetState(74)
		p.Name()
	}
	{
		p.SetState(75)
		p.Match(ViewSQLParserDOT)
	}
	{
		p.SetState(76)
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
		p.SetState(78)
		p.Match(ViewSQLParserLEFT_SP)
	}
	{
		p.SetState(79)
		p.Name()
	}
	{
		p.SetState(80)
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

	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(82)
			p.Name()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(83)
			p.Field()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(84)
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
		p.SetState(87)
		p.Match(ViewSQLParserPRIMARY)
	}
	{
		p.SetState(88)
		p.Match(ViewSQLParserKEY)
	}
	{
		p.SetState(89)
		p.Match(ViewSQLParserLEFT_P)
	}
	{
		p.SetState(90)
		p.Nameable()
	}
	{
		p.SetState(91)
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
	p.SetState(104)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ViewSQLParserLEFT_SP, ViewSQLParserINV_COMMA, ViewSQLParserDATE, ViewSQLParserSTRING, ViewSQLParserINT, ViewSQLParserFLOAT:
		localctx = NewValueContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		p.SetState(96)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case ViewSQLParserLEFT_SP, ViewSQLParserSTRING:
			{
				p.SetState(94)
				p.Nameable()
			}

		case ViewSQLParserINV_COMMA, ViewSQLParserDATE, ViewSQLParserINT, ViewSQLParserFLOAT:
			{
				p.SetState(95)
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
			p.SetState(98)
			p.Match(ViewSQLParserSUB)
		}
		{
			p.SetState(99)
			p.math(4)
		}

	case ViewSQLParserLEFT_P:
		localctx = NewParenthesesContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(100)
			p.Match(ViewSQLParserLEFT_P)
		}
		{
			p.SetState(101)
			p.math(0)
		}
		{
			p.SetState(102)
			p.Match(ViewSQLParserRIGHT_P)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(112)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMultOrDivContext(p, NewMathContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ViewSQLParserRULE_math)
				p.SetState(106)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(107)

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
					p.SetState(108)
					p.math(3)
				}

			case 2:
				localctx = NewAddOrSubContext(p, NewMathContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ViewSQLParserRULE_math)
				p.SetState(109)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(110)

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
					p.SetState(111)
					p.math(2)
				}

			}

		}
		p.SetState(116)
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
		p.SetState(117)
		p.math(0)
	}
	{
		p.SetState(118)
		p.Match(ViewSQLParserAS)
	}
	{
		p.SetState(119)
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
		p.SetState(121)
		p.AggrFunc()
	}
	{
		p.SetState(122)
		p.Match(ViewSQLParserLEFT_P)
	}
	{
		p.SetState(123)
		p.math(0)
	}
	{
		p.SetState(124)
		p.Match(ViewSQLParserRIGHT_P)
	}
	{
		p.SetState(125)
		p.Match(ViewSQLParserAS)
	}
	{
		p.SetState(126)
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
		p.SetState(128)
		p.Match(ViewSQLParserCOUNT)
	}
	{
		p.SetState(129)
		p.Match(ViewSQLParserLEFT_P)
	}
	p.SetState(132)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ViewSQLParserLEFT_SP, ViewSQLParserSTRING:
		{
			p.SetState(130)
			p.Nameable()
		}

	case ViewSQLParserMULT:
		{
			p.SetState(131)
			p.Match(ViewSQLParserMULT)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(134)
		p.Match(ViewSQLParserRIGHT_P)
	}
	{
		p.SetState(135)
		p.Match(ViewSQLParserAS)
	}
	{
		p.SetState(136)
		p.Name()
	}

	return localctx
}

// ICalcContext is an interface to support dynamic dispatch.
type ICalcContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Key() IKeyContext
	Nameable() INameableContext
	AsClause() IAsClauseContext
	Aggregation() IAggregationContext
	Count() ICountContext

	// IsCalcContext differentiates from other interfaces.
	IsCalcContext()
}

type CalcContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
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

	p.SetState(143)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(138)
			p.Key()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(139)
			p.Nameable()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(140)
			p.AsClause()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(141)
			p.Aggregation()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(142)
			p.Count()
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

	p.SetState(147)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ViewSQLParserEQUAL, ViewSQLParserHIGHER, ViewSQLParserHIGHER_EQUAL, ViewSQLParserLOWER, ViewSQLParserLOWER_EQUAL, ViewSQLParserNOT_EQUAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(145)

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
			p.SetState(146)
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
		p.SetState(149)
		p.Nameable()
	}
	{
		p.SetState(150)
		p.Comp()
	}
	{
		p.SetState(151)
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
		p.SetState(153)
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
		p.SetState(155)

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
		p.SetState(156)
		p.Constant()
	}
	{
		p.SetState(157)
		p.Match(ViewSQLParserRANGE_SEP)
	}
	{
		p.SetState(158)
		p.Constant()
	}

	{
		p.SetState(160)

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
		p.SetState(162)
		p.Match(ViewSQLParserLEFT_CURLY)
	}
	{
		p.SetState(163)
		p.Constant()
	}
	p.SetState(168)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ViewSQLParserSEPARATOR {
		{
			p.SetState(164)
			p.Match(ViewSQLParserSEPARATOR)
		}
		{
			p.SetState(165)
			p.Constant()
		}

		p.SetState(170)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(171)
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

	p.SetState(175)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ViewSQLParserLEFT_SP, ViewSQLParserRIGHT_SP:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(173)
			p.ContinuousRange()
		}

	case ViewSQLParserLEFT_CURLY:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(174)
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
		p.SetState(177)
		p.Match(ViewSQLParserCREATE)
	}
	{
		p.SetState(178)
		p.Match(ViewSQLParserVIEW)
	}
	{
		p.SetState(179)
		p.Match(ViewSQLParserLEFT_P)
	}
	{
		p.SetState(180)
		p.Match(ViewSQLParserSTRING)
	}
	{
		p.SetState(181)
		p.Match(ViewSQLParserSEPARATOR)
	}
	{
		p.SetState(182)
		p.Match(ViewSQLParserSTRING)
	}
	{
		p.SetState(183)
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
		p.SetState(185)
		p.Match(ViewSQLParserWITH)
	}
	{
		p.SetState(186)
		p.Match(ViewSQLParserSTRING)
	}
	{
		p.SetState(187)
		p.Match(ViewSQLParserIN)
	}
	{
		p.SetState(188)
		p.Range_()
	}
	p.SetState(195)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ViewSQLParserSEPARATOR {
		{
			p.SetState(189)
			p.Match(ViewSQLParserSEPARATOR)
		}
		{
			p.SetState(190)
			p.Match(ViewSQLParserSTRING)
		}
		{
			p.SetState(191)
			p.Match(ViewSQLParserIN)
		}
		{
			p.SetState(192)
			p.Range_()
		}

		p.SetState(197)
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
		p.SetState(198)
		p.Match(ViewSQLParserSELECT)
	}
	{
		p.SetState(199)
		p.Calc()
	}
	p.SetState(204)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ViewSQLParserSEPARATOR {
		{
			p.SetState(200)
			p.Match(ViewSQLParserSEPARATOR)
		}
		{
			p.SetState(201)
			p.Calc()
		}

		p.SetState(206)
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
		p.SetState(207)
		p.Match(ViewSQLParserFROM)
	}
	{
		p.SetState(208)
		p.Name()
	}
	p.SetState(213)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ViewSQLParserSEPARATOR {
		{
			p.SetState(209)
			p.Match(ViewSQLParserSEPARATOR)
		}
		{
			p.SetState(210)
			p.Name()
		}

		p.SetState(215)
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
		p.SetState(216)
		p.Match(ViewSQLParserWHERE)
	}
	{
		p.SetState(217)
		p.Condition()
	}
	p.SetState(222)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ViewSQLParserAND {
		{
			p.SetState(218)
			p.Match(ViewSQLParserAND)
		}
		{
			p.SetState(219)
			p.Condition()
		}

		p.SetState(224)
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
		p.SetState(225)
		p.Match(ViewSQLParserGROUP)
	}
	{
		p.SetState(226)
		p.Match(ViewSQLParserBY)
	}
	{
		p.SetState(227)
		p.Nameable()
	}
	p.SetState(232)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ViewSQLParserSEPARATOR {
		{
			p.SetState(228)
			p.Match(ViewSQLParserSEPARATOR)
		}
		{
			p.SetState(229)
			p.Nameable()
		}

		p.SetState(234)
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
		p.SetState(235)
		p.Match(ViewSQLParserORDER)
	}
	{
		p.SetState(236)
		p.Match(ViewSQLParserBY)
	}
	{
		p.SetState(237)
		p.Nameable()
	}
	{
		p.SetState(238)
		p.SortOrder()
	}
	p.SetState(245)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ViewSQLParserSEPARATOR {
		{
			p.SetState(239)
			p.Match(ViewSQLParserSEPARATOR)
		}
		{
			p.SetState(240)
			p.Nameable()
		}
		{
			p.SetState(241)
			p.SortOrder()
		}

		p.SetState(247)
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
		p.SetState(248)
		p.Match(ViewSQLParserLIMIT)
	}
	{
		p.SetState(249)
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
		p.SetState(251)
		p.Create()
	}
	p.SetState(253)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ViewSQLParserWITH {
		{
			p.SetState(252)
			p.With()
		}

	}
	{
		p.SetState(255)
		p.Match(ViewSQLParserAS)
	}
	{
		p.SetState(256)
		p.Select_()
	}
	{
		p.SetState(257)
		p.From()
	}
	p.SetState(259)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ViewSQLParserWHERE {
		{
			p.SetState(258)
			p.Where()
		}

	}
	p.SetState(262)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(261)

			var _x = p.Groupby()

			localctx.(*ViewContext).firstGroupBy = _x
		}

	}
	{
		p.SetState(264)

		var _x = p.Groupby()

		localctx.(*ViewContext).secondGroupBy = _x
	}
	p.SetState(266)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ViewSQLParserORDER {
		{
			p.SetState(265)
			p.Orderby()
		}

	}
	p.SetState(269)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ViewSQLParserLIMIT {
		{
			p.SetState(268)
			p.Limit()
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
	View() IViewContext
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

func (s *StartContext) View() IViewContext {
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
	p.EnterRule(localctx, 54, ViewSQLParserRULE_start)

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
		p.SetState(271)
		p.View()
	}
	{
		p.SetState(272)
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
