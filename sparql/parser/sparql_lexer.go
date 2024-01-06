// Code generated from c://git//src//github.com//lercher//rdf//sparql//Sparql.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type SparqlLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var SparqlLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func sparqllexerLexerInit() {
	staticData := &SparqlLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'BASE'", "'PREFIX'", "'SELECT'", "'DISTINCT'", "'REDUCED'", "'*'",
		"'CONSTRUCT'", "'DESCRIBE'", "'ASK'", "'FROM'", "'NAMED'", "'WHERE'",
		"'ORDER'", "'BY'", "'ASC'", "'DESC'", "'LIMIT'", "'OFFSET'", "'{'",
		"'.'", "'}'", "'OPTIONAL'", "'GRAPH'", "'UNION'", "'FILTER'", "'('",
		"','", "')'", "';'", "'A'", "'['", "']'", "'||'", "'&&'", "'='", "'!='",
		"'<'", "'>'", "'<='", "'>='", "'+'", "'-'", "'/'", "'!'", "'STR'", "'LANG'",
		"'LANGMATCHES'", "'DATATYPE'", "'BOUND'", "'SAMETERM'", "'ISIRI'", "'ISURI'",
		"'ISBLANK'", "'ISLITERAL'", "'REGEX'", "'^^'", "'TRUE'", "'FALSE'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "IRI_REF", "PNAME_NS", "PNAME_LN", "BLANK_NODE_LABEL",
		"VAR1", "VAR2", "LANGTAG", "INTEGER", "DECIMAL", "DOUBLE", "INTEGER_POSITIVE",
		"DECIMAL_POSITIVE", "DOUBLE_POSITIVE", "INTEGER_NEGATIVE", "DECIMAL_NEGATIVE",
		"DOUBLE_NEGATIVE", "EXPONENT", "STRING_LITERAL1", "STRING_LITERAL2",
		"STRING_LITERAL_LONG1", "STRING_LITERAL_LONG2", "ECHAR", "NIL", "ANON",
		"PN_CHARS_U", "VARNAME", "PN_PREFIX", "PN_LOCAL", "WS",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
		"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24",
		"T__25", "T__26", "T__27", "T__28", "T__29", "T__30", "T__31", "T__32",
		"T__33", "T__34", "T__35", "T__36", "T__37", "T__38", "T__39", "T__40",
		"T__41", "T__42", "T__43", "T__44", "T__45", "T__46", "T__47", "T__48",
		"T__49", "T__50", "T__51", "T__52", "T__53", "T__54", "T__55", "T__56",
		"T__57", "IRI_REF", "PNAME_NS", "PNAME_LN", "BLANK_NODE_LABEL", "VAR1",
		"VAR2", "LANGTAG", "INTEGER", "DECIMAL", "DOUBLE", "INTEGER_POSITIVE",
		"DECIMAL_POSITIVE", "DOUBLE_POSITIVE", "INTEGER_NEGATIVE", "DECIMAL_NEGATIVE",
		"DOUBLE_NEGATIVE", "EXPONENT", "STRING_LITERAL1", "STRING_LITERAL2",
		"STRING_LITERAL_LONG1", "STRING_LITERAL_LONG2", "ECHAR", "NIL", "ANON",
		"PN_CHARS_U", "VARNAME", "PN_CHARS", "PN_PREFIX", "PN_LOCAL", "PN_CHARS_BASE",
		"DIGIT", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 87, 735, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
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
		2, 68, 7, 68, 2, 69, 7, 69, 2, 70, 7, 70, 2, 71, 7, 71, 2, 72, 7, 72, 2,
		73, 7, 73, 2, 74, 7, 74, 2, 75, 7, 75, 2, 76, 7, 76, 2, 77, 7, 77, 2, 78,
		7, 78, 2, 79, 7, 79, 2, 80, 7, 80, 2, 81, 7, 81, 2, 82, 7, 82, 2, 83, 7,
		83, 2, 84, 7, 84, 2, 85, 7, 85, 2, 86, 7, 86, 2, 87, 7, 87, 2, 88, 7, 88,
		2, 89, 7, 89, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1,
		16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18,
		1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1,
		21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23,
		1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1,
		24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 28, 1, 28, 1, 29, 1, 29,
		1, 30, 1, 30, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1,
		34, 1, 34, 1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 37, 1, 37, 1, 38, 1, 38,
		1, 38, 1, 39, 1, 39, 1, 39, 1, 40, 1, 40, 1, 41, 1, 41, 1, 42, 1, 42, 1,
		43, 1, 43, 1, 44, 1, 44, 1, 44, 1, 44, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45,
		1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1,
		46, 1, 46, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47,
		1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 48, 1, 49, 1, 49, 1, 49, 1, 49, 1,
		49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50,
		1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 1, 52, 1, 52, 1, 52, 1, 52, 1,
		52, 1, 52, 1, 52, 1, 52, 1, 53, 1, 53, 1, 53, 1, 53, 1, 53, 1, 53, 1, 53,
		1, 53, 1, 53, 1, 53, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1, 55, 1,
		55, 1, 55, 1, 56, 1, 56, 1, 56, 1, 56, 1, 56, 1, 57, 1, 57, 1, 57, 1, 57,
		1, 57, 1, 57, 1, 58, 1, 58, 1, 58, 5, 58, 466, 8, 58, 10, 58, 12, 58, 469,
		9, 58, 1, 58, 1, 58, 1, 59, 3, 59, 474, 8, 59, 1, 59, 1, 59, 1, 60, 1,
		60, 1, 60, 1, 61, 1, 61, 1, 61, 1, 61, 1, 61, 1, 62, 1, 62, 1, 62, 1, 63,
		1, 63, 1, 63, 1, 64, 1, 64, 4, 64, 494, 8, 64, 11, 64, 12, 64, 495, 1,
		64, 1, 64, 1, 64, 1, 64, 4, 64, 502, 8, 64, 11, 64, 12, 64, 503, 5, 64,
		506, 8, 64, 10, 64, 12, 64, 509, 9, 64, 1, 65, 4, 65, 512, 8, 65, 11, 65,
		12, 65, 513, 1, 66, 4, 66, 517, 8, 66, 11, 66, 12, 66, 518, 1, 66, 1, 66,
		5, 66, 523, 8, 66, 10, 66, 12, 66, 526, 9, 66, 1, 66, 1, 66, 4, 66, 530,
		8, 66, 11, 66, 12, 66, 531, 3, 66, 534, 8, 66, 1, 67, 4, 67, 537, 8, 67,
		11, 67, 12, 67, 538, 1, 67, 1, 67, 5, 67, 543, 8, 67, 10, 67, 12, 67, 546,
		9, 67, 1, 67, 1, 67, 1, 67, 1, 67, 4, 67, 552, 8, 67, 11, 67, 12, 67, 553,
		1, 67, 1, 67, 1, 67, 4, 67, 559, 8, 67, 11, 67, 12, 67, 560, 1, 67, 1,
		67, 3, 67, 565, 8, 67, 1, 68, 1, 68, 1, 68, 1, 69, 1, 69, 1, 69, 1, 70,
		1, 70, 1, 70, 1, 71, 1, 71, 1, 71, 1, 72, 1, 72, 1, 72, 1, 73, 1, 73, 1,
		73, 1, 74, 1, 74, 3, 74, 587, 8, 74, 1, 74, 4, 74, 590, 8, 74, 11, 74,
		12, 74, 591, 1, 75, 1, 75, 1, 75, 5, 75, 597, 8, 75, 10, 75, 12, 75, 600,
		9, 75, 1, 75, 1, 75, 1, 76, 1, 76, 1, 76, 5, 76, 607, 8, 76, 10, 76, 12,
		76, 610, 9, 76, 1, 76, 1, 76, 1, 77, 1, 77, 1, 77, 1, 77, 1, 77, 1, 77,
		1, 77, 3, 77, 621, 8, 77, 1, 77, 1, 77, 3, 77, 625, 8, 77, 5, 77, 627,
		8, 77, 10, 77, 12, 77, 630, 9, 77, 1, 77, 1, 77, 1, 77, 1, 77, 1, 78, 1,
		78, 1, 78, 1, 78, 1, 78, 1, 78, 1, 78, 3, 78, 643, 8, 78, 1, 78, 1, 78,
		3, 78, 647, 8, 78, 5, 78, 649, 8, 78, 10, 78, 12, 78, 652, 9, 78, 1, 78,
		1, 78, 1, 78, 1, 78, 1, 79, 1, 79, 1, 79, 1, 80, 1, 80, 5, 80, 663, 8,
		80, 10, 80, 12, 80, 666, 9, 80, 1, 80, 1, 80, 1, 81, 1, 81, 5, 81, 672,
		8, 81, 10, 81, 12, 81, 675, 9, 81, 1, 81, 1, 81, 1, 82, 1, 82, 3, 82, 681,
		8, 82, 1, 83, 1, 83, 3, 83, 685, 8, 83, 1, 83, 1, 83, 1, 83, 5, 83, 690,
		8, 83, 10, 83, 12, 83, 693, 9, 83, 1, 84, 1, 84, 1, 84, 3, 84, 698, 8,
		84, 1, 85, 1, 85, 1, 85, 5, 85, 703, 8, 85, 10, 85, 12, 85, 706, 9, 85,
		1, 85, 3, 85, 709, 8, 85, 1, 86, 1, 86, 3, 86, 713, 8, 86, 1, 86, 1, 86,
		5, 86, 717, 8, 86, 10, 86, 12, 86, 720, 9, 86, 1, 86, 3, 86, 723, 8, 86,
		1, 87, 1, 87, 1, 88, 1, 88, 1, 89, 4, 89, 730, 8, 89, 11, 89, 12, 89, 731,
		1, 89, 1, 89, 0, 0, 90, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15,
		8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17,
		35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26,
		53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35,
		71, 36, 73, 37, 75, 38, 77, 39, 79, 40, 81, 41, 83, 42, 85, 43, 87, 44,
		89, 45, 91, 46, 93, 47, 95, 48, 97, 49, 99, 50, 101, 51, 103, 52, 105,
		53, 107, 54, 109, 55, 111, 56, 113, 57, 115, 58, 117, 59, 119, 60, 121,
		61, 123, 62, 125, 63, 127, 64, 129, 65, 131, 66, 133, 67, 135, 68, 137,
		69, 139, 70, 141, 71, 143, 72, 145, 73, 147, 74, 149, 75, 151, 76, 153,
		77, 155, 78, 157, 79, 159, 80, 161, 81, 163, 82, 165, 83, 167, 84, 169,
		0, 171, 85, 173, 86, 175, 0, 177, 0, 179, 87, 1, 0, 10, 7, 0, 34, 34, 60,
		60, 62, 62, 92, 92, 94, 94, 96, 96, 123, 125, 2, 0, 69, 69, 101, 101, 2,
		0, 43, 43, 45, 45, 4, 0, 10, 10, 13, 13, 39, 39, 92, 92, 4, 0, 10, 10,
		13, 13, 34, 34, 92, 92, 2, 0, 39, 39, 92, 92, 7, 0, 34, 34, 39, 39, 98,
		98, 102, 102, 110, 110, 114, 114, 116, 116, 3, 0, 183, 183, 768, 879, 8255,
		8256, 13, 0, 65, 90, 97, 122, 192, 214, 216, 246, 248, 767, 880, 893, 895,
		8191, 8204, 8205, 8304, 8591, 11264, 12271, 12289, 55295, 63744, 64975,
		65008, 65533, 3, 0, 9, 10, 13, 13, 32, 32, 779, 0, 1, 1, 0, 0, 0, 0, 3,
		1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11,
		1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0,
		19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0,
		0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0,
		0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0,
		0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1,
		0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57,
		1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0,
		65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0,
		0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0,
		0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 0, 85, 1, 0, 0, 0, 0, 87, 1, 0,
		0, 0, 0, 89, 1, 0, 0, 0, 0, 91, 1, 0, 0, 0, 0, 93, 1, 0, 0, 0, 0, 95, 1,
		0, 0, 0, 0, 97, 1, 0, 0, 0, 0, 99, 1, 0, 0, 0, 0, 101, 1, 0, 0, 0, 0, 103,
		1, 0, 0, 0, 0, 105, 1, 0, 0, 0, 0, 107, 1, 0, 0, 0, 0, 109, 1, 0, 0, 0,
		0, 111, 1, 0, 0, 0, 0, 113, 1, 0, 0, 0, 0, 115, 1, 0, 0, 0, 0, 117, 1,
		0, 0, 0, 0, 119, 1, 0, 0, 0, 0, 121, 1, 0, 0, 0, 0, 123, 1, 0, 0, 0, 0,
		125, 1, 0, 0, 0, 0, 127, 1, 0, 0, 0, 0, 129, 1, 0, 0, 0, 0, 131, 1, 0,
		0, 0, 0, 133, 1, 0, 0, 0, 0, 135, 1, 0, 0, 0, 0, 137, 1, 0, 0, 0, 0, 139,
		1, 0, 0, 0, 0, 141, 1, 0, 0, 0, 0, 143, 1, 0, 0, 0, 0, 145, 1, 0, 0, 0,
		0, 147, 1, 0, 0, 0, 0, 149, 1, 0, 0, 0, 0, 151, 1, 0, 0, 0, 0, 153, 1,
		0, 0, 0, 0, 155, 1, 0, 0, 0, 0, 157, 1, 0, 0, 0, 0, 159, 1, 0, 0, 0, 0,
		161, 1, 0, 0, 0, 0, 163, 1, 0, 0, 0, 0, 165, 1, 0, 0, 0, 0, 167, 1, 0,
		0, 0, 0, 171, 1, 0, 0, 0, 0, 173, 1, 0, 0, 0, 0, 179, 1, 0, 0, 0, 1, 181,
		1, 0, 0, 0, 3, 186, 1, 0, 0, 0, 5, 193, 1, 0, 0, 0, 7, 200, 1, 0, 0, 0,
		9, 209, 1, 0, 0, 0, 11, 217, 1, 0, 0, 0, 13, 219, 1, 0, 0, 0, 15, 229,
		1, 0, 0, 0, 17, 238, 1, 0, 0, 0, 19, 242, 1, 0, 0, 0, 21, 247, 1, 0, 0,
		0, 23, 253, 1, 0, 0, 0, 25, 259, 1, 0, 0, 0, 27, 265, 1, 0, 0, 0, 29, 268,
		1, 0, 0, 0, 31, 272, 1, 0, 0, 0, 33, 277, 1, 0, 0, 0, 35, 283, 1, 0, 0,
		0, 37, 290, 1, 0, 0, 0, 39, 292, 1, 0, 0, 0, 41, 294, 1, 0, 0, 0, 43, 296,
		1, 0, 0, 0, 45, 305, 1, 0, 0, 0, 47, 311, 1, 0, 0, 0, 49, 317, 1, 0, 0,
		0, 51, 324, 1, 0, 0, 0, 53, 326, 1, 0, 0, 0, 55, 328, 1, 0, 0, 0, 57, 330,
		1, 0, 0, 0, 59, 332, 1, 0, 0, 0, 61, 334, 1, 0, 0, 0, 63, 336, 1, 0, 0,
		0, 65, 338, 1, 0, 0, 0, 67, 341, 1, 0, 0, 0, 69, 344, 1, 0, 0, 0, 71, 346,
		1, 0, 0, 0, 73, 349, 1, 0, 0, 0, 75, 351, 1, 0, 0, 0, 77, 353, 1, 0, 0,
		0, 79, 356, 1, 0, 0, 0, 81, 359, 1, 0, 0, 0, 83, 361, 1, 0, 0, 0, 85, 363,
		1, 0, 0, 0, 87, 365, 1, 0, 0, 0, 89, 367, 1, 0, 0, 0, 91, 371, 1, 0, 0,
		0, 93, 376, 1, 0, 0, 0, 95, 388, 1, 0, 0, 0, 97, 397, 1, 0, 0, 0, 99, 403,
		1, 0, 0, 0, 101, 412, 1, 0, 0, 0, 103, 418, 1, 0, 0, 0, 105, 424, 1, 0,
		0, 0, 107, 432, 1, 0, 0, 0, 109, 442, 1, 0, 0, 0, 111, 448, 1, 0, 0, 0,
		113, 451, 1, 0, 0, 0, 115, 456, 1, 0, 0, 0, 117, 462, 1, 0, 0, 0, 119,
		473, 1, 0, 0, 0, 121, 477, 1, 0, 0, 0, 123, 480, 1, 0, 0, 0, 125, 485,
		1, 0, 0, 0, 127, 488, 1, 0, 0, 0, 129, 491, 1, 0, 0, 0, 131, 511, 1, 0,
		0, 0, 133, 533, 1, 0, 0, 0, 135, 564, 1, 0, 0, 0, 137, 566, 1, 0, 0, 0,
		139, 569, 1, 0, 0, 0, 141, 572, 1, 0, 0, 0, 143, 575, 1, 0, 0, 0, 145,
		578, 1, 0, 0, 0, 147, 581, 1, 0, 0, 0, 149, 584, 1, 0, 0, 0, 151, 593,
		1, 0, 0, 0, 153, 603, 1, 0, 0, 0, 155, 613, 1, 0, 0, 0, 157, 635, 1, 0,
		0, 0, 159, 657, 1, 0, 0, 0, 161, 660, 1, 0, 0, 0, 163, 669, 1, 0, 0, 0,
		165, 680, 1, 0, 0, 0, 167, 684, 1, 0, 0, 0, 169, 697, 1, 0, 0, 0, 171,
		699, 1, 0, 0, 0, 173, 712, 1, 0, 0, 0, 175, 724, 1, 0, 0, 0, 177, 726,
		1, 0, 0, 0, 179, 729, 1, 0, 0, 0, 181, 182, 5, 66, 0, 0, 182, 183, 5, 65,
		0, 0, 183, 184, 5, 83, 0, 0, 184, 185, 5, 69, 0, 0, 185, 2, 1, 0, 0, 0,
		186, 187, 5, 80, 0, 0, 187, 188, 5, 82, 0, 0, 188, 189, 5, 69, 0, 0, 189,
		190, 5, 70, 0, 0, 190, 191, 5, 73, 0, 0, 191, 192, 5, 88, 0, 0, 192, 4,
		1, 0, 0, 0, 193, 194, 5, 83, 0, 0, 194, 195, 5, 69, 0, 0, 195, 196, 5,
		76, 0, 0, 196, 197, 5, 69, 0, 0, 197, 198, 5, 67, 0, 0, 198, 199, 5, 84,
		0, 0, 199, 6, 1, 0, 0, 0, 200, 201, 5, 68, 0, 0, 201, 202, 5, 73, 0, 0,
		202, 203, 5, 83, 0, 0, 203, 204, 5, 84, 0, 0, 204, 205, 5, 73, 0, 0, 205,
		206, 5, 78, 0, 0, 206, 207, 5, 67, 0, 0, 207, 208, 5, 84, 0, 0, 208, 8,
		1, 0, 0, 0, 209, 210, 5, 82, 0, 0, 210, 211, 5, 69, 0, 0, 211, 212, 5,
		68, 0, 0, 212, 213, 5, 85, 0, 0, 213, 214, 5, 67, 0, 0, 214, 215, 5, 69,
		0, 0, 215, 216, 5, 68, 0, 0, 216, 10, 1, 0, 0, 0, 217, 218, 5, 42, 0, 0,
		218, 12, 1, 0, 0, 0, 219, 220, 5, 67, 0, 0, 220, 221, 5, 79, 0, 0, 221,
		222, 5, 78, 0, 0, 222, 223, 5, 83, 0, 0, 223, 224, 5, 84, 0, 0, 224, 225,
		5, 82, 0, 0, 225, 226, 5, 85, 0, 0, 226, 227, 5, 67, 0, 0, 227, 228, 5,
		84, 0, 0, 228, 14, 1, 0, 0, 0, 229, 230, 5, 68, 0, 0, 230, 231, 5, 69,
		0, 0, 231, 232, 5, 83, 0, 0, 232, 233, 5, 67, 0, 0, 233, 234, 5, 82, 0,
		0, 234, 235, 5, 73, 0, 0, 235, 236, 5, 66, 0, 0, 236, 237, 5, 69, 0, 0,
		237, 16, 1, 0, 0, 0, 238, 239, 5, 65, 0, 0, 239, 240, 5, 83, 0, 0, 240,
		241, 5, 75, 0, 0, 241, 18, 1, 0, 0, 0, 242, 243, 5, 70, 0, 0, 243, 244,
		5, 82, 0, 0, 244, 245, 5, 79, 0, 0, 245, 246, 5, 77, 0, 0, 246, 20, 1,
		0, 0, 0, 247, 248, 5, 78, 0, 0, 248, 249, 5, 65, 0, 0, 249, 250, 5, 77,
		0, 0, 250, 251, 5, 69, 0, 0, 251, 252, 5, 68, 0, 0, 252, 22, 1, 0, 0, 0,
		253, 254, 5, 87, 0, 0, 254, 255, 5, 72, 0, 0, 255, 256, 5, 69, 0, 0, 256,
		257, 5, 82, 0, 0, 257, 258, 5, 69, 0, 0, 258, 24, 1, 0, 0, 0, 259, 260,
		5, 79, 0, 0, 260, 261, 5, 82, 0, 0, 261, 262, 5, 68, 0, 0, 262, 263, 5,
		69, 0, 0, 263, 264, 5, 82, 0, 0, 264, 26, 1, 0, 0, 0, 265, 266, 5, 66,
		0, 0, 266, 267, 5, 89, 0, 0, 267, 28, 1, 0, 0, 0, 268, 269, 5, 65, 0, 0,
		269, 270, 5, 83, 0, 0, 270, 271, 5, 67, 0, 0, 271, 30, 1, 0, 0, 0, 272,
		273, 5, 68, 0, 0, 273, 274, 5, 69, 0, 0, 274, 275, 5, 83, 0, 0, 275, 276,
		5, 67, 0, 0, 276, 32, 1, 0, 0, 0, 277, 278, 5, 76, 0, 0, 278, 279, 5, 73,
		0, 0, 279, 280, 5, 77, 0, 0, 280, 281, 5, 73, 0, 0, 281, 282, 5, 84, 0,
		0, 282, 34, 1, 0, 0, 0, 283, 284, 5, 79, 0, 0, 284, 285, 5, 70, 0, 0, 285,
		286, 5, 70, 0, 0, 286, 287, 5, 83, 0, 0, 287, 288, 5, 69, 0, 0, 288, 289,
		5, 84, 0, 0, 289, 36, 1, 0, 0, 0, 290, 291, 5, 123, 0, 0, 291, 38, 1, 0,
		0, 0, 292, 293, 5, 46, 0, 0, 293, 40, 1, 0, 0, 0, 294, 295, 5, 125, 0,
		0, 295, 42, 1, 0, 0, 0, 296, 297, 5, 79, 0, 0, 297, 298, 5, 80, 0, 0, 298,
		299, 5, 84, 0, 0, 299, 300, 5, 73, 0, 0, 300, 301, 5, 79, 0, 0, 301, 302,
		5, 78, 0, 0, 302, 303, 5, 65, 0, 0, 303, 304, 5, 76, 0, 0, 304, 44, 1,
		0, 0, 0, 305, 306, 5, 71, 0, 0, 306, 307, 5, 82, 0, 0, 307, 308, 5, 65,
		0, 0, 308, 309, 5, 80, 0, 0, 309, 310, 5, 72, 0, 0, 310, 46, 1, 0, 0, 0,
		311, 312, 5, 85, 0, 0, 312, 313, 5, 78, 0, 0, 313, 314, 5, 73, 0, 0, 314,
		315, 5, 79, 0, 0, 315, 316, 5, 78, 0, 0, 316, 48, 1, 0, 0, 0, 317, 318,
		5, 70, 0, 0, 318, 319, 5, 73, 0, 0, 319, 320, 5, 76, 0, 0, 320, 321, 5,
		84, 0, 0, 321, 322, 5, 69, 0, 0, 322, 323, 5, 82, 0, 0, 323, 50, 1, 0,
		0, 0, 324, 325, 5, 40, 0, 0, 325, 52, 1, 0, 0, 0, 326, 327, 5, 44, 0, 0,
		327, 54, 1, 0, 0, 0, 328, 329, 5, 41, 0, 0, 329, 56, 1, 0, 0, 0, 330, 331,
		5, 59, 0, 0, 331, 58, 1, 0, 0, 0, 332, 333, 5, 65, 0, 0, 333, 60, 1, 0,
		0, 0, 334, 335, 5, 91, 0, 0, 335, 62, 1, 0, 0, 0, 336, 337, 5, 93, 0, 0,
		337, 64, 1, 0, 0, 0, 338, 339, 5, 124, 0, 0, 339, 340, 5, 124, 0, 0, 340,
		66, 1, 0, 0, 0, 341, 342, 5, 38, 0, 0, 342, 343, 5, 38, 0, 0, 343, 68,
		1, 0, 0, 0, 344, 345, 5, 61, 0, 0, 345, 70, 1, 0, 0, 0, 346, 347, 5, 33,
		0, 0, 347, 348, 5, 61, 0, 0, 348, 72, 1, 0, 0, 0, 349, 350, 5, 60, 0, 0,
		350, 74, 1, 0, 0, 0, 351, 352, 5, 62, 0, 0, 352, 76, 1, 0, 0, 0, 353, 354,
		5, 60, 0, 0, 354, 355, 5, 61, 0, 0, 355, 78, 1, 0, 0, 0, 356, 357, 5, 62,
		0, 0, 357, 358, 5, 61, 0, 0, 358, 80, 1, 0, 0, 0, 359, 360, 5, 43, 0, 0,
		360, 82, 1, 0, 0, 0, 361, 362, 5, 45, 0, 0, 362, 84, 1, 0, 0, 0, 363, 364,
		5, 47, 0, 0, 364, 86, 1, 0, 0, 0, 365, 366, 5, 33, 0, 0, 366, 88, 1, 0,
		0, 0, 367, 368, 5, 83, 0, 0, 368, 369, 5, 84, 0, 0, 369, 370, 5, 82, 0,
		0, 370, 90, 1, 0, 0, 0, 371, 372, 5, 76, 0, 0, 372, 373, 5, 65, 0, 0, 373,
		374, 5, 78, 0, 0, 374, 375, 5, 71, 0, 0, 375, 92, 1, 0, 0, 0, 376, 377,
		5, 76, 0, 0, 377, 378, 5, 65, 0, 0, 378, 379, 5, 78, 0, 0, 379, 380, 5,
		71, 0, 0, 380, 381, 5, 77, 0, 0, 381, 382, 5, 65, 0, 0, 382, 383, 5, 84,
		0, 0, 383, 384, 5, 67, 0, 0, 384, 385, 5, 72, 0, 0, 385, 386, 5, 69, 0,
		0, 386, 387, 5, 83, 0, 0, 387, 94, 1, 0, 0, 0, 388, 389, 5, 68, 0, 0, 389,
		390, 5, 65, 0, 0, 390, 391, 5, 84, 0, 0, 391, 392, 5, 65, 0, 0, 392, 393,
		5, 84, 0, 0, 393, 394, 5, 89, 0, 0, 394, 395, 5, 80, 0, 0, 395, 396, 5,
		69, 0, 0, 396, 96, 1, 0, 0, 0, 397, 398, 5, 66, 0, 0, 398, 399, 5, 79,
		0, 0, 399, 400, 5, 85, 0, 0, 400, 401, 5, 78, 0, 0, 401, 402, 5, 68, 0,
		0, 402, 98, 1, 0, 0, 0, 403, 404, 5, 83, 0, 0, 404, 405, 5, 65, 0, 0, 405,
		406, 5, 77, 0, 0, 406, 407, 5, 69, 0, 0, 407, 408, 5, 84, 0, 0, 408, 409,
		5, 69, 0, 0, 409, 410, 5, 82, 0, 0, 410, 411, 5, 77, 0, 0, 411, 100, 1,
		0, 0, 0, 412, 413, 5, 73, 0, 0, 413, 414, 5, 83, 0, 0, 414, 415, 5, 73,
		0, 0, 415, 416, 5, 82, 0, 0, 416, 417, 5, 73, 0, 0, 417, 102, 1, 0, 0,
		0, 418, 419, 5, 73, 0, 0, 419, 420, 5, 83, 0, 0, 420, 421, 5, 85, 0, 0,
		421, 422, 5, 82, 0, 0, 422, 423, 5, 73, 0, 0, 423, 104, 1, 0, 0, 0, 424,
		425, 5, 73, 0, 0, 425, 426, 5, 83, 0, 0, 426, 427, 5, 66, 0, 0, 427, 428,
		5, 76, 0, 0, 428, 429, 5, 65, 0, 0, 429, 430, 5, 78, 0, 0, 430, 431, 5,
		75, 0, 0, 431, 106, 1, 0, 0, 0, 432, 433, 5, 73, 0, 0, 433, 434, 5, 83,
		0, 0, 434, 435, 5, 76, 0, 0, 435, 436, 5, 73, 0, 0, 436, 437, 5, 84, 0,
		0, 437, 438, 5, 69, 0, 0, 438, 439, 5, 82, 0, 0, 439, 440, 5, 65, 0, 0,
		440, 441, 5, 76, 0, 0, 441, 108, 1, 0, 0, 0, 442, 443, 5, 82, 0, 0, 443,
		444, 5, 69, 0, 0, 444, 445, 5, 71, 0, 0, 445, 446, 5, 69, 0, 0, 446, 447,
		5, 88, 0, 0, 447, 110, 1, 0, 0, 0, 448, 449, 5, 94, 0, 0, 449, 450, 5,
		94, 0, 0, 450, 112, 1, 0, 0, 0, 451, 452, 5, 84, 0, 0, 452, 453, 5, 82,
		0, 0, 453, 454, 5, 85, 0, 0, 454, 455, 5, 69, 0, 0, 455, 114, 1, 0, 0,
		0, 456, 457, 5, 70, 0, 0, 457, 458, 5, 65, 0, 0, 458, 459, 5, 76, 0, 0,
		459, 460, 5, 83, 0, 0, 460, 461, 5, 69, 0, 0, 461, 116, 1, 0, 0, 0, 462,
		467, 5, 60, 0, 0, 463, 466, 8, 0, 0, 0, 464, 466, 3, 169, 84, 0, 465, 463,
		1, 0, 0, 0, 465, 464, 1, 0, 0, 0, 466, 469, 1, 0, 0, 0, 467, 465, 1, 0,
		0, 0, 467, 468, 1, 0, 0, 0, 468, 470, 1, 0, 0, 0, 469, 467, 1, 0, 0, 0,
		470, 471, 5, 62, 0, 0, 471, 118, 1, 0, 0, 0, 472, 474, 3, 171, 85, 0, 473,
		472, 1, 0, 0, 0, 473, 474, 1, 0, 0, 0, 474, 475, 1, 0, 0, 0, 475, 476,
		5, 58, 0, 0, 476, 120, 1, 0, 0, 0, 477, 478, 3, 119, 59, 0, 478, 479, 3,
		173, 86, 0, 479, 122, 1, 0, 0, 0, 480, 481, 5, 95, 0, 0, 481, 482, 5, 58,
		0, 0, 482, 483, 1, 0, 0, 0, 483, 484, 3, 173, 86, 0, 484, 124, 1, 0, 0,
		0, 485, 486, 5, 63, 0, 0, 486, 487, 3, 167, 83, 0, 487, 126, 1, 0, 0, 0,
		488, 489, 5, 36, 0, 0, 489, 490, 3, 167, 83, 0, 490, 128, 1, 0, 0, 0, 491,
		493, 5, 64, 0, 0, 492, 494, 3, 175, 87, 0, 493, 492, 1, 0, 0, 0, 494, 495,
		1, 0, 0, 0, 495, 493, 1, 0, 0, 0, 495, 496, 1, 0, 0, 0, 496, 507, 1, 0,
		0, 0, 497, 501, 5, 45, 0, 0, 498, 499, 3, 175, 87, 0, 499, 500, 3, 177,
		88, 0, 500, 502, 1, 0, 0, 0, 501, 498, 1, 0, 0, 0, 502, 503, 1, 0, 0, 0,
		503, 501, 1, 0, 0, 0, 503, 504, 1, 0, 0, 0, 504, 506, 1, 0, 0, 0, 505,
		497, 1, 0, 0, 0, 506, 509, 1, 0, 0, 0, 507, 505, 1, 0, 0, 0, 507, 508,
		1, 0, 0, 0, 508, 130, 1, 0, 0, 0, 509, 507, 1, 0, 0, 0, 510, 512, 3, 177,
		88, 0, 511, 510, 1, 0, 0, 0, 512, 513, 1, 0, 0, 0, 513, 511, 1, 0, 0, 0,
		513, 514, 1, 0, 0, 0, 514, 132, 1, 0, 0, 0, 515, 517, 3, 177, 88, 0, 516,
		515, 1, 0, 0, 0, 517, 518, 1, 0, 0, 0, 518, 516, 1, 0, 0, 0, 518, 519,
		1, 0, 0, 0, 519, 520, 1, 0, 0, 0, 520, 524, 5, 46, 0, 0, 521, 523, 3, 177,
		88, 0, 522, 521, 1, 0, 0, 0, 523, 526, 1, 0, 0, 0, 524, 522, 1, 0, 0, 0,
		524, 525, 1, 0, 0, 0, 525, 534, 1, 0, 0, 0, 526, 524, 1, 0, 0, 0, 527,
		529, 5, 46, 0, 0, 528, 530, 3, 177, 88, 0, 529, 528, 1, 0, 0, 0, 530, 531,
		1, 0, 0, 0, 531, 529, 1, 0, 0, 0, 531, 532, 1, 0, 0, 0, 532, 534, 1, 0,
		0, 0, 533, 516, 1, 0, 0, 0, 533, 527, 1, 0, 0, 0, 534, 134, 1, 0, 0, 0,
		535, 537, 3, 177, 88, 0, 536, 535, 1, 0, 0, 0, 537, 538, 1, 0, 0, 0, 538,
		536, 1, 0, 0, 0, 538, 539, 1, 0, 0, 0, 539, 540, 1, 0, 0, 0, 540, 544,
		5, 46, 0, 0, 541, 543, 3, 177, 88, 0, 542, 541, 1, 0, 0, 0, 543, 546, 1,
		0, 0, 0, 544, 542, 1, 0, 0, 0, 544, 545, 1, 0, 0, 0, 545, 547, 1, 0, 0,
		0, 546, 544, 1, 0, 0, 0, 547, 548, 3, 149, 74, 0, 548, 565, 1, 0, 0, 0,
		549, 551, 5, 46, 0, 0, 550, 552, 3, 177, 88, 0, 551, 550, 1, 0, 0, 0, 552,
		553, 1, 0, 0, 0, 553, 551, 1, 0, 0, 0, 553, 554, 1, 0, 0, 0, 554, 555,
		1, 0, 0, 0, 555, 556, 3, 149, 74, 0, 556, 565, 1, 0, 0, 0, 557, 559, 3,
		177, 88, 0, 558, 557, 1, 0, 0, 0, 559, 560, 1, 0, 0, 0, 560, 558, 1, 0,
		0, 0, 560, 561, 1, 0, 0, 0, 561, 562, 1, 0, 0, 0, 562, 563, 3, 149, 74,
		0, 563, 565, 1, 0, 0, 0, 564, 536, 1, 0, 0, 0, 564, 549, 1, 0, 0, 0, 564,
		558, 1, 0, 0, 0, 565, 136, 1, 0, 0, 0, 566, 567, 5, 43, 0, 0, 567, 568,
		3, 131, 65, 0, 568, 138, 1, 0, 0, 0, 569, 570, 5, 43, 0, 0, 570, 571, 3,
		133, 66, 0, 571, 140, 1, 0, 0, 0, 572, 573, 5, 43, 0, 0, 573, 574, 3, 135,
		67, 0, 574, 142, 1, 0, 0, 0, 575, 576, 5, 45, 0, 0, 576, 577, 3, 131, 65,
		0, 577, 144, 1, 0, 0, 0, 578, 579, 5, 45, 0, 0, 579, 580, 3, 133, 66, 0,
		580, 146, 1, 0, 0, 0, 581, 582, 5, 45, 0, 0, 582, 583, 3, 135, 67, 0, 583,
		148, 1, 0, 0, 0, 584, 586, 7, 1, 0, 0, 585, 587, 7, 2, 0, 0, 586, 585,
		1, 0, 0, 0, 586, 587, 1, 0, 0, 0, 587, 589, 1, 0, 0, 0, 588, 590, 3, 177,
		88, 0, 589, 588, 1, 0, 0, 0, 590, 591, 1, 0, 0, 0, 591, 589, 1, 0, 0, 0,
		591, 592, 1, 0, 0, 0, 592, 150, 1, 0, 0, 0, 593, 598, 5, 39, 0, 0, 594,
		597, 8, 3, 0, 0, 595, 597, 3, 159, 79, 0, 596, 594, 1, 0, 0, 0, 596, 595,
		1, 0, 0, 0, 597, 600, 1, 0, 0, 0, 598, 596, 1, 0, 0, 0, 598, 599, 1, 0,
		0, 0, 599, 601, 1, 0, 0, 0, 600, 598, 1, 0, 0, 0, 601, 602, 5, 39, 0, 0,
		602, 152, 1, 0, 0, 0, 603, 608, 5, 34, 0, 0, 604, 607, 8, 4, 0, 0, 605,
		607, 3, 159, 79, 0, 606, 604, 1, 0, 0, 0, 606, 605, 1, 0, 0, 0, 607, 610,
		1, 0, 0, 0, 608, 606, 1, 0, 0, 0, 608, 609, 1, 0, 0, 0, 609, 611, 1, 0,
		0, 0, 610, 608, 1, 0, 0, 0, 611, 612, 5, 34, 0, 0, 612, 154, 1, 0, 0, 0,
		613, 614, 5, 39, 0, 0, 614, 615, 5, 39, 0, 0, 615, 616, 5, 39, 0, 0, 616,
		628, 1, 0, 0, 0, 617, 621, 5, 39, 0, 0, 618, 619, 5, 39, 0, 0, 619, 621,
		5, 39, 0, 0, 620, 617, 1, 0, 0, 0, 620, 618, 1, 0, 0, 0, 620, 621, 1, 0,
		0, 0, 621, 624, 1, 0, 0, 0, 622, 625, 8, 5, 0, 0, 623, 625, 3, 159, 79,
		0, 624, 622, 1, 0, 0, 0, 624, 623, 1, 0, 0, 0, 625, 627, 1, 0, 0, 0, 626,
		620, 1, 0, 0, 0, 627, 630, 1, 0, 0, 0, 628, 626, 1, 0, 0, 0, 628, 629,
		1, 0, 0, 0, 629, 631, 1, 0, 0, 0, 630, 628, 1, 0, 0, 0, 631, 632, 5, 39,
		0, 0, 632, 633, 5, 39, 0, 0, 633, 634, 5, 39, 0, 0, 634, 156, 1, 0, 0,
		0, 635, 636, 5, 34, 0, 0, 636, 637, 5, 34, 0, 0, 637, 638, 5, 34, 0, 0,
		638, 650, 1, 0, 0, 0, 639, 643, 5, 34, 0, 0, 640, 641, 5, 34, 0, 0, 641,
		643, 5, 34, 0, 0, 642, 639, 1, 0, 0, 0, 642, 640, 1, 0, 0, 0, 642, 643,
		1, 0, 0, 0, 643, 646, 1, 0, 0, 0, 644, 647, 8, 5, 0, 0, 645, 647, 3, 159,
		79, 0, 646, 644, 1, 0, 0, 0, 646, 645, 1, 0, 0, 0, 647, 649, 1, 0, 0, 0,
		648, 642, 1, 0, 0, 0, 649, 652, 1, 0, 0, 0, 650, 648, 1, 0, 0, 0, 650,
		651, 1, 0, 0, 0, 651, 653, 1, 0, 0, 0, 652, 650, 1, 0, 0, 0, 653, 654,
		5, 34, 0, 0, 654, 655, 5, 34, 0, 0, 655, 656, 5, 34, 0, 0, 656, 158, 1,
		0, 0, 0, 657, 658, 5, 92, 0, 0, 658, 659, 7, 6, 0, 0, 659, 160, 1, 0, 0,
		0, 660, 664, 5, 40, 0, 0, 661, 663, 3, 179, 89, 0, 662, 661, 1, 0, 0, 0,
		663, 666, 1, 0, 0, 0, 664, 662, 1, 0, 0, 0, 664, 665, 1, 0, 0, 0, 665,
		667, 1, 0, 0, 0, 666, 664, 1, 0, 0, 0, 667, 668, 5, 41, 0, 0, 668, 162,
		1, 0, 0, 0, 669, 673, 5, 91, 0, 0, 670, 672, 3, 179, 89, 0, 671, 670, 1,
		0, 0, 0, 672, 675, 1, 0, 0, 0, 673, 671, 1, 0, 0, 0, 673, 674, 1, 0, 0,
		0, 674, 676, 1, 0, 0, 0, 675, 673, 1, 0, 0, 0, 676, 677, 5, 93, 0, 0, 677,
		164, 1, 0, 0, 0, 678, 681, 3, 175, 87, 0, 679, 681, 5, 95, 0, 0, 680, 678,
		1, 0, 0, 0, 680, 679, 1, 0, 0, 0, 681, 166, 1, 0, 0, 0, 682, 685, 3, 165,
		82, 0, 683, 685, 3, 177, 88, 0, 684, 682, 1, 0, 0, 0, 684, 683, 1, 0, 0,
		0, 685, 691, 1, 0, 0, 0, 686, 690, 3, 165, 82, 0, 687, 690, 3, 177, 88,
		0, 688, 690, 7, 7, 0, 0, 689, 686, 1, 0, 0, 0, 689, 687, 1, 0, 0, 0, 689,
		688, 1, 0, 0, 0, 690, 693, 1, 0, 0, 0, 691, 689, 1, 0, 0, 0, 691, 692,
		1, 0, 0, 0, 692, 168, 1, 0, 0, 0, 693, 691, 1, 0, 0, 0, 694, 698, 3, 165,
		82, 0, 695, 698, 5, 45, 0, 0, 696, 698, 3, 177, 88, 0, 697, 694, 1, 0,
		0, 0, 697, 695, 1, 0, 0, 0, 697, 696, 1, 0, 0, 0, 698, 170, 1, 0, 0, 0,
		699, 708, 3, 175, 87, 0, 700, 703, 3, 169, 84, 0, 701, 703, 5, 46, 0, 0,
		702, 700, 1, 0, 0, 0, 702, 701, 1, 0, 0, 0, 703, 706, 1, 0, 0, 0, 704,
		702, 1, 0, 0, 0, 704, 705, 1, 0, 0, 0, 705, 707, 1, 0, 0, 0, 706, 704,
		1, 0, 0, 0, 707, 709, 3, 169, 84, 0, 708, 704, 1, 0, 0, 0, 708, 709, 1,
		0, 0, 0, 709, 172, 1, 0, 0, 0, 710, 713, 3, 165, 82, 0, 711, 713, 3, 177,
		88, 0, 712, 710, 1, 0, 0, 0, 712, 711, 1, 0, 0, 0, 713, 722, 1, 0, 0, 0,
		714, 717, 3, 169, 84, 0, 715, 717, 5, 46, 0, 0, 716, 714, 1, 0, 0, 0, 716,
		715, 1, 0, 0, 0, 717, 720, 1, 0, 0, 0, 718, 716, 1, 0, 0, 0, 718, 719,
		1, 0, 0, 0, 719, 721, 1, 0, 0, 0, 720, 718, 1, 0, 0, 0, 721, 723, 3, 169,
		84, 0, 722, 718, 1, 0, 0, 0, 722, 723, 1, 0, 0, 0, 723, 174, 1, 0, 0, 0,
		724, 725, 7, 8, 0, 0, 725, 176, 1, 0, 0, 0, 726, 727, 2, 48, 57, 0, 727,
		178, 1, 0, 0, 0, 728, 730, 7, 9, 0, 0, 729, 728, 1, 0, 0, 0, 730, 731,
		1, 0, 0, 0, 731, 729, 1, 0, 0, 0, 731, 732, 1, 0, 0, 0, 732, 733, 1, 0,
		0, 0, 733, 734, 6, 89, 0, 0, 734, 180, 1, 0, 0, 0, 44, 0, 465, 467, 473,
		495, 503, 507, 513, 518, 524, 531, 533, 538, 544, 553, 560, 564, 586, 591,
		596, 598, 606, 608, 620, 624, 628, 642, 646, 650, 664, 673, 680, 684, 689,
		691, 697, 702, 704, 708, 712, 716, 718, 722, 731, 1, 6, 0, 0,
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

// SparqlLexerInit initializes any static state used to implement SparqlLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewSparqlLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func SparqlLexerInit() {
	staticData := &SparqlLexerLexerStaticData
	staticData.once.Do(sparqllexerLexerInit)
}

// NewSparqlLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewSparqlLexer(input antlr.CharStream) *SparqlLexer {
	SparqlLexerInit()
	l := new(SparqlLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &SparqlLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Sparql.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// SparqlLexer tokens.
const (
	SparqlLexerT__0                 = 1
	SparqlLexerT__1                 = 2
	SparqlLexerT__2                 = 3
	SparqlLexerT__3                 = 4
	SparqlLexerT__4                 = 5
	SparqlLexerT__5                 = 6
	SparqlLexerT__6                 = 7
	SparqlLexerT__7                 = 8
	SparqlLexerT__8                 = 9
	SparqlLexerT__9                 = 10
	SparqlLexerT__10                = 11
	SparqlLexerT__11                = 12
	SparqlLexerT__12                = 13
	SparqlLexerT__13                = 14
	SparqlLexerT__14                = 15
	SparqlLexerT__15                = 16
	SparqlLexerT__16                = 17
	SparqlLexerT__17                = 18
	SparqlLexerT__18                = 19
	SparqlLexerT__19                = 20
	SparqlLexerT__20                = 21
	SparqlLexerT__21                = 22
	SparqlLexerT__22                = 23
	SparqlLexerT__23                = 24
	SparqlLexerT__24                = 25
	SparqlLexerT__25                = 26
	SparqlLexerT__26                = 27
	SparqlLexerT__27                = 28
	SparqlLexerT__28                = 29
	SparqlLexerT__29                = 30
	SparqlLexerT__30                = 31
	SparqlLexerT__31                = 32
	SparqlLexerT__32                = 33
	SparqlLexerT__33                = 34
	SparqlLexerT__34                = 35
	SparqlLexerT__35                = 36
	SparqlLexerT__36                = 37
	SparqlLexerT__37                = 38
	SparqlLexerT__38                = 39
	SparqlLexerT__39                = 40
	SparqlLexerT__40                = 41
	SparqlLexerT__41                = 42
	SparqlLexerT__42                = 43
	SparqlLexerT__43                = 44
	SparqlLexerT__44                = 45
	SparqlLexerT__45                = 46
	SparqlLexerT__46                = 47
	SparqlLexerT__47                = 48
	SparqlLexerT__48                = 49
	SparqlLexerT__49                = 50
	SparqlLexerT__50                = 51
	SparqlLexerT__51                = 52
	SparqlLexerT__52                = 53
	SparqlLexerT__53                = 54
	SparqlLexerT__54                = 55
	SparqlLexerT__55                = 56
	SparqlLexerT__56                = 57
	SparqlLexerT__57                = 58
	SparqlLexerIRI_REF              = 59
	SparqlLexerPNAME_NS             = 60
	SparqlLexerPNAME_LN             = 61
	SparqlLexerBLANK_NODE_LABEL     = 62
	SparqlLexerVAR1                 = 63
	SparqlLexerVAR2                 = 64
	SparqlLexerLANGTAG              = 65
	SparqlLexerINTEGER              = 66
	SparqlLexerDECIMAL              = 67
	SparqlLexerDOUBLE               = 68
	SparqlLexerINTEGER_POSITIVE     = 69
	SparqlLexerDECIMAL_POSITIVE     = 70
	SparqlLexerDOUBLE_POSITIVE      = 71
	SparqlLexerINTEGER_NEGATIVE     = 72
	SparqlLexerDECIMAL_NEGATIVE     = 73
	SparqlLexerDOUBLE_NEGATIVE      = 74
	SparqlLexerEXPONENT             = 75
	SparqlLexerSTRING_LITERAL1      = 76
	SparqlLexerSTRING_LITERAL2      = 77
	SparqlLexerSTRING_LITERAL_LONG1 = 78
	SparqlLexerSTRING_LITERAL_LONG2 = 79
	SparqlLexerECHAR                = 80
	SparqlLexerNIL                  = 81
	SparqlLexerANON                 = 82
	SparqlLexerPN_CHARS_U           = 83
	SparqlLexerVARNAME              = 84
	SparqlLexerPN_PREFIX            = 85
	SparqlLexerPN_LOCAL             = 86
	SparqlLexerWS                   = 87
)
