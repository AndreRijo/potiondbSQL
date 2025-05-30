// Code generated from ViewSQL.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type ViewSQLLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var viewsqllexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func viewsqllexerLexerInit() {
	staticData := &viewsqllexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
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
		"ADD", "SUB", "MULT", "DIV", "SUM", "AVG", "MAX", "MIN", "COUNT", "AND",
		"ASC", "DESC", "EQUAL", "HIGHER", "HIGHER_EQUAL", "LOWER", "LOWER_EQUAL",
		"NOT_EQUAL", "DOT", "SEPARATOR", "RANGE_SEP", "LEFT_P", "RIGHT_P", "LEFT_SP",
		"RIGHT_SP", "INV_COMMA", "LEFT_CURLY", "RIGHT_CURLY", "AS", "CREATE",
		"VIEW", "SELECT", "FROM", "WHERE", "GROUP", "BY", "ORDER", "LIMIT",
		"WITH", "IN", "PRIMARY", "KEY", "TABLE", "INDEX", "UPDATE", "DELETE",
		"INSERT", "INTO", "DROP", "SET", "FOREIGN", "REFERENCES", "VALUES",
		"CHECK", "ON", "AW", "RW", "LWW", "DEFAULT", "INTEGER", "COUNTER", "BOOLEAN",
		"VARCHAR", "DATE_TYPE", "DATE", "DAY", "MONTH", "YEAR", "BOOL", "STRING",
		"INT", "FLOAT", "WHITESPACE",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 70, 515, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2,
		52, 7, 52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57,
		7, 57, 2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 2, 61, 7, 61, 2, 62, 7,
		62, 2, 63, 7, 63, 2, 64, 7, 64, 2, 65, 7, 65, 2, 66, 7, 66, 2, 67, 7, 67,
		2, 68, 7, 68, 2, 69, 7, 69, 2, 70, 7, 70, 2, 71, 7, 71, 2, 72, 7, 72, 1,
		0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1,
		13, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17,
		1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1,
		22, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27,
		1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1,
		30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31,
		1, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1,
		33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 35, 1, 35, 1, 35,
		1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 37, 1, 37, 1, 37, 1, 37, 1,
		37, 1, 37, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 39, 1, 39, 1, 39, 1, 40,
		1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 41, 1, 41, 1, 41, 1,
		41, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 43, 1, 43, 1, 43, 1, 43,
		1, 43, 1, 43, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 45, 1,
		45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46,
		1, 46, 1, 46, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 48, 1, 48, 1, 48, 1,
		48, 1, 48, 1, 49, 1, 49, 1, 49, 1, 49, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50,
		1, 50, 1, 50, 1, 50, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 1,
		51, 1, 51, 1, 51, 1, 51, 1, 52, 1, 52, 1, 52, 1, 52, 1, 52, 1, 52, 1, 52,
		1, 53, 1, 53, 1, 53, 1, 53, 1, 53, 1, 53, 1, 54, 1, 54, 1, 54, 1, 55, 1,
		55, 1, 55, 1, 56, 1, 56, 1, 56, 1, 57, 1, 57, 1, 57, 1, 57, 1, 58, 1, 58,
		1, 58, 1, 58, 1, 58, 1, 58, 1, 58, 1, 58, 1, 59, 1, 59, 1, 59, 1, 59, 1,
		59, 1, 59, 1, 59, 1, 59, 1, 60, 1, 60, 1, 60, 1, 60, 1, 60, 1, 60, 1, 60,
		1, 60, 1, 61, 1, 61, 1, 61, 1, 61, 1, 61, 1, 61, 1, 61, 1, 61, 1, 62, 1,
		62, 1, 62, 1, 62, 1, 62, 1, 62, 1, 62, 1, 62, 1, 63, 1, 63, 1, 63, 1, 63,
		1, 63, 1, 64, 1, 64, 1, 64, 1, 64, 1, 64, 1, 64, 1, 65, 1, 65, 1, 65, 1,
		65, 1, 65, 1, 65, 3, 65, 449, 8, 65, 1, 66, 1, 66, 1, 66, 1, 66, 3, 66,
		455, 8, 66, 1, 67, 1, 67, 1, 67, 1, 67, 1, 67, 1, 68, 1, 68, 1, 68, 1,
		68, 1, 68, 1, 68, 1, 68, 1, 68, 1, 68, 3, 68, 471, 8, 68, 1, 69, 1, 69,
		5, 69, 475, 8, 69, 10, 69, 12, 69, 478, 9, 69, 1, 70, 1, 70, 3, 70, 482,
		8, 70, 1, 70, 1, 70, 5, 70, 486, 8, 70, 10, 70, 12, 70, 489, 9, 70, 3,
		70, 491, 8, 70, 1, 71, 3, 71, 494, 8, 71, 1, 71, 1, 71, 1, 71, 5, 71, 499,
		8, 71, 10, 71, 12, 71, 502, 9, 71, 1, 71, 4, 71, 505, 8, 71, 11, 71, 12,
		71, 506, 1, 72, 4, 72, 510, 8, 72, 11, 72, 12, 72, 511, 1, 72, 1, 72, 0,
		0, 73, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10,
		21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19,
		39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28,
		57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35, 71, 36, 73, 37,
		75, 38, 77, 39, 79, 40, 81, 41, 83, 42, 85, 43, 87, 44, 89, 45, 91, 46,
		93, 47, 95, 48, 97, 49, 99, 50, 101, 51, 103, 52, 105, 53, 107, 54, 109,
		55, 111, 56, 113, 57, 115, 58, 117, 59, 119, 60, 121, 61, 123, 62, 125,
		63, 127, 64, 129, 65, 131, 0, 133, 0, 135, 0, 137, 66, 139, 67, 141, 68,
		143, 69, 145, 70, 1, 0, 7, 1, 0, 49, 57, 1, 0, 49, 50, 1, 0, 48, 57, 1,
		0, 48, 49, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95,
		97, 122, 3, 0, 9, 10, 13, 13, 32, 32, 523, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0,
		0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0,
		0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1,
		0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27,
		1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0,
		35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0,
		0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0,
		0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0,
		0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1,
		0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73,
		1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0,
		81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 0, 85, 1, 0, 0, 0, 0, 87, 1, 0, 0, 0,
		0, 89, 1, 0, 0, 0, 0, 91, 1, 0, 0, 0, 0, 93, 1, 0, 0, 0, 0, 95, 1, 0, 0,
		0, 0, 97, 1, 0, 0, 0, 0, 99, 1, 0, 0, 0, 0, 101, 1, 0, 0, 0, 0, 103, 1,
		0, 0, 0, 0, 105, 1, 0, 0, 0, 0, 107, 1, 0, 0, 0, 0, 109, 1, 0, 0, 0, 0,
		111, 1, 0, 0, 0, 0, 113, 1, 0, 0, 0, 0, 115, 1, 0, 0, 0, 0, 117, 1, 0,
		0, 0, 0, 119, 1, 0, 0, 0, 0, 121, 1, 0, 0, 0, 0, 123, 1, 0, 0, 0, 0, 125,
		1, 0, 0, 0, 0, 127, 1, 0, 0, 0, 0, 129, 1, 0, 0, 0, 0, 137, 1, 0, 0, 0,
		0, 139, 1, 0, 0, 0, 0, 141, 1, 0, 0, 0, 0, 143, 1, 0, 0, 0, 0, 145, 1,
		0, 0, 0, 1, 147, 1, 0, 0, 0, 3, 149, 1, 0, 0, 0, 5, 151, 1, 0, 0, 0, 7,
		153, 1, 0, 0, 0, 9, 155, 1, 0, 0, 0, 11, 159, 1, 0, 0, 0, 13, 163, 1, 0,
		0, 0, 15, 167, 1, 0, 0, 0, 17, 171, 1, 0, 0, 0, 19, 177, 1, 0, 0, 0, 21,
		181, 1, 0, 0, 0, 23, 185, 1, 0, 0, 0, 25, 190, 1, 0, 0, 0, 27, 192, 1,
		0, 0, 0, 29, 194, 1, 0, 0, 0, 31, 197, 1, 0, 0, 0, 33, 199, 1, 0, 0, 0,
		35, 202, 1, 0, 0, 0, 37, 205, 1, 0, 0, 0, 39, 207, 1, 0, 0, 0, 41, 209,
		1, 0, 0, 0, 43, 211, 1, 0, 0, 0, 45, 213, 1, 0, 0, 0, 47, 215, 1, 0, 0,
		0, 49, 217, 1, 0, 0, 0, 51, 219, 1, 0, 0, 0, 53, 221, 1, 0, 0, 0, 55, 223,
		1, 0, 0, 0, 57, 225, 1, 0, 0, 0, 59, 228, 1, 0, 0, 0, 61, 235, 1, 0, 0,
		0, 63, 240, 1, 0, 0, 0, 65, 247, 1, 0, 0, 0, 67, 252, 1, 0, 0, 0, 69, 258,
		1, 0, 0, 0, 71, 264, 1, 0, 0, 0, 73, 267, 1, 0, 0, 0, 75, 273, 1, 0, 0,
		0, 77, 279, 1, 0, 0, 0, 79, 284, 1, 0, 0, 0, 81, 287, 1, 0, 0, 0, 83, 295,
		1, 0, 0, 0, 85, 299, 1, 0, 0, 0, 87, 305, 1, 0, 0, 0, 89, 311, 1, 0, 0,
		0, 91, 318, 1, 0, 0, 0, 93, 325, 1, 0, 0, 0, 95, 332, 1, 0, 0, 0, 97, 337,
		1, 0, 0, 0, 99, 342, 1, 0, 0, 0, 101, 346, 1, 0, 0, 0, 103, 354, 1, 0,
		0, 0, 105, 365, 1, 0, 0, 0, 107, 372, 1, 0, 0, 0, 109, 378, 1, 0, 0, 0,
		111, 381, 1, 0, 0, 0, 113, 384, 1, 0, 0, 0, 115, 387, 1, 0, 0, 0, 117,
		391, 1, 0, 0, 0, 119, 399, 1, 0, 0, 0, 121, 407, 1, 0, 0, 0, 123, 415,
		1, 0, 0, 0, 125, 423, 1, 0, 0, 0, 127, 431, 1, 0, 0, 0, 129, 436, 1, 0,
		0, 0, 131, 448, 1, 0, 0, 0, 133, 454, 1, 0, 0, 0, 135, 456, 1, 0, 0, 0,
		137, 470, 1, 0, 0, 0, 139, 472, 1, 0, 0, 0, 141, 490, 1, 0, 0, 0, 143,
		493, 1, 0, 0, 0, 145, 509, 1, 0, 0, 0, 147, 148, 5, 43, 0, 0, 148, 2, 1,
		0, 0, 0, 149, 150, 5, 45, 0, 0, 150, 4, 1, 0, 0, 0, 151, 152, 5, 42, 0,
		0, 152, 6, 1, 0, 0, 0, 153, 154, 5, 47, 0, 0, 154, 8, 1, 0, 0, 0, 155,
		156, 5, 83, 0, 0, 156, 157, 5, 85, 0, 0, 157, 158, 5, 77, 0, 0, 158, 10,
		1, 0, 0, 0, 159, 160, 5, 65, 0, 0, 160, 161, 5, 86, 0, 0, 161, 162, 5,
		71, 0, 0, 162, 12, 1, 0, 0, 0, 163, 164, 5, 77, 0, 0, 164, 165, 5, 65,
		0, 0, 165, 166, 5, 88, 0, 0, 166, 14, 1, 0, 0, 0, 167, 168, 5, 77, 0, 0,
		168, 169, 5, 73, 0, 0, 169, 170, 5, 78, 0, 0, 170, 16, 1, 0, 0, 0, 171,
		172, 5, 67, 0, 0, 172, 173, 5, 79, 0, 0, 173, 174, 5, 85, 0, 0, 174, 175,
		5, 78, 0, 0, 175, 176, 5, 84, 0, 0, 176, 18, 1, 0, 0, 0, 177, 178, 5, 65,
		0, 0, 178, 179, 5, 78, 0, 0, 179, 180, 5, 68, 0, 0, 180, 20, 1, 0, 0, 0,
		181, 182, 5, 65, 0, 0, 182, 183, 5, 83, 0, 0, 183, 184, 5, 67, 0, 0, 184,
		22, 1, 0, 0, 0, 185, 186, 5, 68, 0, 0, 186, 187, 5, 69, 0, 0, 187, 188,
		5, 83, 0, 0, 188, 189, 5, 67, 0, 0, 189, 24, 1, 0, 0, 0, 190, 191, 5, 61,
		0, 0, 191, 26, 1, 0, 0, 0, 192, 193, 5, 62, 0, 0, 193, 28, 1, 0, 0, 0,
		194, 195, 5, 62, 0, 0, 195, 196, 5, 61, 0, 0, 196, 30, 1, 0, 0, 0, 197,
		198, 5, 60, 0, 0, 198, 32, 1, 0, 0, 0, 199, 200, 5, 60, 0, 0, 200, 201,
		5, 61, 0, 0, 201, 34, 1, 0, 0, 0, 202, 203, 5, 33, 0, 0, 203, 204, 5, 61,
		0, 0, 204, 36, 1, 0, 0, 0, 205, 206, 5, 46, 0, 0, 206, 38, 1, 0, 0, 0,
		207, 208, 5, 44, 0, 0, 208, 40, 1, 0, 0, 0, 209, 210, 5, 58, 0, 0, 210,
		42, 1, 0, 0, 0, 211, 212, 5, 40, 0, 0, 212, 44, 1, 0, 0, 0, 213, 214, 5,
		41, 0, 0, 214, 46, 1, 0, 0, 0, 215, 216, 5, 91, 0, 0, 216, 48, 1, 0, 0,
		0, 217, 218, 5, 93, 0, 0, 218, 50, 1, 0, 0, 0, 219, 220, 5, 34, 0, 0, 220,
		52, 1, 0, 0, 0, 221, 222, 5, 123, 0, 0, 222, 54, 1, 0, 0, 0, 223, 224,
		5, 125, 0, 0, 224, 56, 1, 0, 0, 0, 225, 226, 5, 65, 0, 0, 226, 227, 5,
		83, 0, 0, 227, 58, 1, 0, 0, 0, 228, 229, 5, 67, 0, 0, 229, 230, 5, 82,
		0, 0, 230, 231, 5, 69, 0, 0, 231, 232, 5, 65, 0, 0, 232, 233, 5, 84, 0,
		0, 233, 234, 5, 69, 0, 0, 234, 60, 1, 0, 0, 0, 235, 236, 5, 86, 0, 0, 236,
		237, 5, 73, 0, 0, 237, 238, 5, 69, 0, 0, 238, 239, 5, 87, 0, 0, 239, 62,
		1, 0, 0, 0, 240, 241, 5, 83, 0, 0, 241, 242, 5, 69, 0, 0, 242, 243, 5,
		76, 0, 0, 243, 244, 5, 69, 0, 0, 244, 245, 5, 67, 0, 0, 245, 246, 5, 84,
		0, 0, 246, 64, 1, 0, 0, 0, 247, 248, 5, 70, 0, 0, 248, 249, 5, 82, 0, 0,
		249, 250, 5, 79, 0, 0, 250, 251, 5, 77, 0, 0, 251, 66, 1, 0, 0, 0, 252,
		253, 5, 87, 0, 0, 253, 254, 5, 72, 0, 0, 254, 255, 5, 69, 0, 0, 255, 256,
		5, 82, 0, 0, 256, 257, 5, 69, 0, 0, 257, 68, 1, 0, 0, 0, 258, 259, 5, 71,
		0, 0, 259, 260, 5, 82, 0, 0, 260, 261, 5, 79, 0, 0, 261, 262, 5, 85, 0,
		0, 262, 263, 5, 80, 0, 0, 263, 70, 1, 0, 0, 0, 264, 265, 5, 66, 0, 0, 265,
		266, 5, 89, 0, 0, 266, 72, 1, 0, 0, 0, 267, 268, 5, 79, 0, 0, 268, 269,
		5, 82, 0, 0, 269, 270, 5, 68, 0, 0, 270, 271, 5, 69, 0, 0, 271, 272, 5,
		82, 0, 0, 272, 74, 1, 0, 0, 0, 273, 274, 5, 76, 0, 0, 274, 275, 5, 73,
		0, 0, 275, 276, 5, 77, 0, 0, 276, 277, 5, 73, 0, 0, 277, 278, 5, 84, 0,
		0, 278, 76, 1, 0, 0, 0, 279, 280, 5, 87, 0, 0, 280, 281, 5, 73, 0, 0, 281,
		282, 5, 84, 0, 0, 282, 283, 5, 72, 0, 0, 283, 78, 1, 0, 0, 0, 284, 285,
		5, 73, 0, 0, 285, 286, 5, 78, 0, 0, 286, 80, 1, 0, 0, 0, 287, 288, 5, 80,
		0, 0, 288, 289, 5, 82, 0, 0, 289, 290, 5, 73, 0, 0, 290, 291, 5, 77, 0,
		0, 291, 292, 5, 65, 0, 0, 292, 293, 5, 82, 0, 0, 293, 294, 5, 89, 0, 0,
		294, 82, 1, 0, 0, 0, 295, 296, 5, 75, 0, 0, 296, 297, 5, 69, 0, 0, 297,
		298, 5, 89, 0, 0, 298, 84, 1, 0, 0, 0, 299, 300, 5, 84, 0, 0, 300, 301,
		5, 65, 0, 0, 301, 302, 5, 66, 0, 0, 302, 303, 5, 76, 0, 0, 303, 304, 5,
		69, 0, 0, 304, 86, 1, 0, 0, 0, 305, 306, 5, 73, 0, 0, 306, 307, 5, 78,
		0, 0, 307, 308, 5, 68, 0, 0, 308, 309, 5, 69, 0, 0, 309, 310, 5, 88, 0,
		0, 310, 88, 1, 0, 0, 0, 311, 312, 5, 85, 0, 0, 312, 313, 5, 80, 0, 0, 313,
		314, 5, 68, 0, 0, 314, 315, 5, 65, 0, 0, 315, 316, 5, 84, 0, 0, 316, 317,
		5, 69, 0, 0, 317, 90, 1, 0, 0, 0, 318, 319, 5, 68, 0, 0, 319, 320, 5, 69,
		0, 0, 320, 321, 5, 76, 0, 0, 321, 322, 5, 69, 0, 0, 322, 323, 5, 84, 0,
		0, 323, 324, 5, 69, 0, 0, 324, 92, 1, 0, 0, 0, 325, 326, 5, 73, 0, 0, 326,
		327, 5, 78, 0, 0, 327, 328, 5, 83, 0, 0, 328, 329, 5, 69, 0, 0, 329, 330,
		5, 82, 0, 0, 330, 331, 5, 84, 0, 0, 331, 94, 1, 0, 0, 0, 332, 333, 5, 73,
		0, 0, 333, 334, 5, 78, 0, 0, 334, 335, 5, 84, 0, 0, 335, 336, 5, 79, 0,
		0, 336, 96, 1, 0, 0, 0, 337, 338, 5, 68, 0, 0, 338, 339, 5, 82, 0, 0, 339,
		340, 5, 79, 0, 0, 340, 341, 5, 80, 0, 0, 341, 98, 1, 0, 0, 0, 342, 343,
		5, 83, 0, 0, 343, 344, 5, 69, 0, 0, 344, 345, 5, 84, 0, 0, 345, 100, 1,
		0, 0, 0, 346, 347, 5, 70, 0, 0, 347, 348, 5, 79, 0, 0, 348, 349, 5, 82,
		0, 0, 349, 350, 5, 69, 0, 0, 350, 351, 5, 73, 0, 0, 351, 352, 5, 71, 0,
		0, 352, 353, 5, 78, 0, 0, 353, 102, 1, 0, 0, 0, 354, 355, 5, 82, 0, 0,
		355, 356, 5, 69, 0, 0, 356, 357, 5, 70, 0, 0, 357, 358, 5, 69, 0, 0, 358,
		359, 5, 82, 0, 0, 359, 360, 5, 69, 0, 0, 360, 361, 5, 78, 0, 0, 361, 362,
		5, 67, 0, 0, 362, 363, 5, 69, 0, 0, 363, 364, 5, 83, 0, 0, 364, 104, 1,
		0, 0, 0, 365, 366, 5, 86, 0, 0, 366, 367, 5, 65, 0, 0, 367, 368, 5, 76,
		0, 0, 368, 369, 5, 85, 0, 0, 369, 370, 5, 69, 0, 0, 370, 371, 5, 83, 0,
		0, 371, 106, 1, 0, 0, 0, 372, 373, 5, 67, 0, 0, 373, 374, 5, 72, 0, 0,
		374, 375, 5, 69, 0, 0, 375, 376, 5, 67, 0, 0, 376, 377, 5, 75, 0, 0, 377,
		108, 1, 0, 0, 0, 378, 379, 5, 79, 0, 0, 379, 380, 5, 78, 0, 0, 380, 110,
		1, 0, 0, 0, 381, 382, 5, 65, 0, 0, 382, 383, 5, 87, 0, 0, 383, 112, 1,
		0, 0, 0, 384, 385, 5, 82, 0, 0, 385, 386, 5, 87, 0, 0, 386, 114, 1, 0,
		0, 0, 387, 388, 5, 76, 0, 0, 388, 389, 5, 87, 0, 0, 389, 390, 5, 87, 0,
		0, 390, 116, 1, 0, 0, 0, 391, 392, 5, 68, 0, 0, 392, 393, 5, 69, 0, 0,
		393, 394, 5, 70, 0, 0, 394, 395, 5, 65, 0, 0, 395, 396, 5, 85, 0, 0, 396,
		397, 5, 76, 0, 0, 397, 398, 5, 84, 0, 0, 398, 118, 1, 0, 0, 0, 399, 400,
		5, 105, 0, 0, 400, 401, 5, 110, 0, 0, 401, 402, 5, 116, 0, 0, 402, 403,
		5, 101, 0, 0, 403, 404, 5, 103, 0, 0, 404, 405, 5, 101, 0, 0, 405, 406,
		5, 114, 0, 0, 406, 120, 1, 0, 0, 0, 407, 408, 5, 99, 0, 0, 408, 409, 5,
		111, 0, 0, 409, 410, 5, 117, 0, 0, 410, 411, 5, 110, 0, 0, 411, 412, 5,
		116, 0, 0, 412, 413, 5, 101, 0, 0, 413, 414, 5, 114, 0, 0, 414, 122, 1,
		0, 0, 0, 415, 416, 5, 98, 0, 0, 416, 417, 5, 111, 0, 0, 417, 418, 5, 111,
		0, 0, 418, 419, 5, 108, 0, 0, 419, 420, 5, 101, 0, 0, 420, 421, 5, 97,
		0, 0, 421, 422, 5, 110, 0, 0, 422, 124, 1, 0, 0, 0, 423, 424, 5, 118, 0,
		0, 424, 425, 5, 97, 0, 0, 425, 426, 5, 114, 0, 0, 426, 427, 5, 99, 0, 0,
		427, 428, 5, 104, 0, 0, 428, 429, 5, 97, 0, 0, 429, 430, 5, 114, 0, 0,
		430, 126, 1, 0, 0, 0, 431, 432, 5, 100, 0, 0, 432, 433, 5, 97, 0, 0, 433,
		434, 5, 116, 0, 0, 434, 435, 5, 101, 0, 0, 435, 128, 1, 0, 0, 0, 436, 437,
		3, 135, 67, 0, 437, 438, 5, 45, 0, 0, 438, 439, 3, 133, 66, 0, 439, 440,
		5, 45, 0, 0, 440, 441, 3, 131, 65, 0, 441, 130, 1, 0, 0, 0, 442, 443, 5,
		48, 0, 0, 443, 449, 7, 0, 0, 0, 444, 445, 7, 1, 0, 0, 445, 449, 7, 2, 0,
		0, 446, 447, 5, 51, 0, 0, 447, 449, 7, 3, 0, 0, 448, 442, 1, 0, 0, 0, 448,
		444, 1, 0, 0, 0, 448, 446, 1, 0, 0, 0, 449, 132, 1, 0, 0, 0, 450, 451,
		5, 48, 0, 0, 451, 455, 7, 0, 0, 0, 452, 453, 5, 49, 0, 0, 453, 455, 7,
		1, 0, 0, 454, 450, 1, 0, 0, 0, 454, 452, 1, 0, 0, 0, 455, 134, 1, 0, 0,
		0, 456, 457, 7, 2, 0, 0, 457, 458, 7, 2, 0, 0, 458, 459, 7, 2, 0, 0, 459,
		460, 7, 2, 0, 0, 460, 136, 1, 0, 0, 0, 461, 462, 5, 116, 0, 0, 462, 463,
		5, 114, 0, 0, 463, 464, 5, 117, 0, 0, 464, 471, 5, 101, 0, 0, 465, 466,
		5, 102, 0, 0, 466, 467, 5, 97, 0, 0, 467, 468, 5, 108, 0, 0, 468, 469,
		5, 115, 0, 0, 469, 471, 5, 101, 0, 0, 470, 461, 1, 0, 0, 0, 470, 465, 1,
		0, 0, 0, 471, 138, 1, 0, 0, 0, 472, 476, 7, 4, 0, 0, 473, 475, 7, 5, 0,
		0, 474, 473, 1, 0, 0, 0, 475, 478, 1, 0, 0, 0, 476, 474, 1, 0, 0, 0, 476,
		477, 1, 0, 0, 0, 477, 140, 1, 0, 0, 0, 478, 476, 1, 0, 0, 0, 479, 491,
		5, 48, 0, 0, 480, 482, 5, 45, 0, 0, 481, 480, 1, 0, 0, 0, 481, 482, 1,
		0, 0, 0, 482, 483, 1, 0, 0, 0, 483, 487, 7, 0, 0, 0, 484, 486, 7, 2, 0,
		0, 485, 484, 1, 0, 0, 0, 486, 489, 1, 0, 0, 0, 487, 485, 1, 0, 0, 0, 487,
		488, 1, 0, 0, 0, 488, 491, 1, 0, 0, 0, 489, 487, 1, 0, 0, 0, 490, 479,
		1, 0, 0, 0, 490, 481, 1, 0, 0, 0, 491, 142, 1, 0, 0, 0, 492, 494, 5, 45,
		0, 0, 493, 492, 1, 0, 0, 0, 493, 494, 1, 0, 0, 0, 494, 495, 1, 0, 0, 0,
		495, 496, 7, 2, 0, 0, 496, 500, 5, 46, 0, 0, 497, 499, 7, 2, 0, 0, 498,
		497, 1, 0, 0, 0, 499, 502, 1, 0, 0, 0, 500, 498, 1, 0, 0, 0, 500, 501,
		1, 0, 0, 0, 501, 504, 1, 0, 0, 0, 502, 500, 1, 0, 0, 0, 503, 505, 7, 0,
		0, 0, 504, 503, 1, 0, 0, 0, 505, 506, 1, 0, 0, 0, 506, 504, 1, 0, 0, 0,
		506, 507, 1, 0, 0, 0, 507, 144, 1, 0, 0, 0, 508, 510, 7, 6, 0, 0, 509,
		508, 1, 0, 0, 0, 510, 511, 1, 0, 0, 0, 511, 509, 1, 0, 0, 0, 511, 512,
		1, 0, 0, 0, 512, 513, 1, 0, 0, 0, 513, 514, 6, 72, 0, 0, 514, 146, 1, 0,
		0, 0, 12, 0, 448, 454, 470, 476, 481, 487, 490, 493, 500, 506, 511, 1,
		6, 0, 0,
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

// ViewSQLLexerInit initializes any static state used to implement ViewSQLLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewViewSQLLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func ViewSQLLexerInit() {
	staticData := &viewsqllexerLexerStaticData
	staticData.once.Do(viewsqllexerLexerInit)
}

// NewViewSQLLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewViewSQLLexer(input antlr.CharStream) *ViewSQLLexer {
	ViewSQLLexerInit()
	l := new(ViewSQLLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &viewsqllexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "ViewSQL.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ViewSQLLexer tokens.
const (
	ViewSQLLexerADD          = 1
	ViewSQLLexerSUB          = 2
	ViewSQLLexerMULT         = 3
	ViewSQLLexerDIV          = 4
	ViewSQLLexerSUM          = 5
	ViewSQLLexerAVG          = 6
	ViewSQLLexerMAX          = 7
	ViewSQLLexerMIN          = 8
	ViewSQLLexerCOUNT        = 9
	ViewSQLLexerAND          = 10
	ViewSQLLexerASC          = 11
	ViewSQLLexerDESC         = 12
	ViewSQLLexerEQUAL        = 13
	ViewSQLLexerHIGHER       = 14
	ViewSQLLexerHIGHER_EQUAL = 15
	ViewSQLLexerLOWER        = 16
	ViewSQLLexerLOWER_EQUAL  = 17
	ViewSQLLexerNOT_EQUAL    = 18
	ViewSQLLexerDOT          = 19
	ViewSQLLexerSEPARATOR    = 20
	ViewSQLLexerRANGE_SEP    = 21
	ViewSQLLexerLEFT_P       = 22
	ViewSQLLexerRIGHT_P      = 23
	ViewSQLLexerLEFT_SP      = 24
	ViewSQLLexerRIGHT_SP     = 25
	ViewSQLLexerINV_COMMA    = 26
	ViewSQLLexerLEFT_CURLY   = 27
	ViewSQLLexerRIGHT_CURLY  = 28
	ViewSQLLexerAS           = 29
	ViewSQLLexerCREATE       = 30
	ViewSQLLexerVIEW         = 31
	ViewSQLLexerSELECT       = 32
	ViewSQLLexerFROM         = 33
	ViewSQLLexerWHERE        = 34
	ViewSQLLexerGROUP        = 35
	ViewSQLLexerBY           = 36
	ViewSQLLexerORDER        = 37
	ViewSQLLexerLIMIT        = 38
	ViewSQLLexerWITH         = 39
	ViewSQLLexerIN           = 40
	ViewSQLLexerPRIMARY      = 41
	ViewSQLLexerKEY          = 42
	ViewSQLLexerTABLE        = 43
	ViewSQLLexerINDEX        = 44
	ViewSQLLexerUPDATE       = 45
	ViewSQLLexerDELETE       = 46
	ViewSQLLexerINSERT       = 47
	ViewSQLLexerINTO         = 48
	ViewSQLLexerDROP         = 49
	ViewSQLLexerSET          = 50
	ViewSQLLexerFOREIGN      = 51
	ViewSQLLexerREFERENCES   = 52
	ViewSQLLexerVALUES       = 53
	ViewSQLLexerCHECK        = 54
	ViewSQLLexerON           = 55
	ViewSQLLexerAW           = 56
	ViewSQLLexerRW           = 57
	ViewSQLLexerLWW          = 58
	ViewSQLLexerDEFAULT      = 59
	ViewSQLLexerINTEGER      = 60
	ViewSQLLexerCOUNTER      = 61
	ViewSQLLexerBOOLEAN      = 62
	ViewSQLLexerVARCHAR      = 63
	ViewSQLLexerDATE_TYPE    = 64
	ViewSQLLexerDATE         = 65
	ViewSQLLexerBOOL         = 66
	ViewSQLLexerSTRING       = 67
	ViewSQLLexerINT          = 68
	ViewSQLLexerFLOAT        = 69
	ViewSQLLexerWHITESPACE   = 70
)
