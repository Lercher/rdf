// Code generated from c://git//src//github.com//lercher//rdf//sparql//Sparql.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Sparql

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type SparqlParser struct {
	*antlr.BaseParser
}

var SparqlParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func sparqlParserInit() {
	staticData := &SparqlParserStaticData
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
		"query", "prologue", "baseDecl", "prefixDecl", "selectQuery", "constructQuery",
		"describeQuery", "askQuery", "datasetClause", "defaultGraphClause",
		"namedGraphClause", "sourceSelector", "whereClause", "solutionModifier",
		"limitOffsetClauses", "orderClause", "orderCondition", "limitClause",
		"offsetClause", "groupGraphPattern", "triplesBlock", "graphPatternNotTriples",
		"optionalGraphPattern", "graphGraphPattern", "groupOrUnionGraphPattern",
		"filter", "constraint", "functionCall", "argList", "constructTemplate",
		"constructTriples", "triplesSameSubject", "propertyListNotEmpty", "propertyList",
		"objectList", "object", "verb", "triplesNode", "blankNodePropertyList",
		"collection", "graphNode", "varOrTerm", "varOrIRIref", "varx", "graphTerm",
		"expression", "conditionalOrExpression", "conditionalAndExpression",
		"valueLogical", "relationalExpression", "numericExpression", "additiveExpression",
		"multiplicativeExpression", "unaryExpression", "primaryExpression",
		"brackettedExpression", "builtInCall", "regexExpression", "iriRefOrFunction",
		"rdfLiteral", "numericLiteral", "numericLiteralUnsigned", "numericLiteralPositive",
		"numericLiteralNegative", "booleanLiteral", "rdfstring", "iriRef", "prefixedName",
		"blankNode",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 87, 623, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 2, 47,
		7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2, 52, 7,
		52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57, 7, 57,
		2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 2, 61, 7, 61, 2, 62, 7, 62, 2,
		63, 7, 63, 2, 64, 7, 64, 2, 65, 7, 65, 2, 66, 7, 66, 2, 67, 7, 67, 2, 68,
		7, 68, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 3, 0, 144, 8, 0, 1, 0, 1, 0, 1, 1,
		3, 1, 149, 8, 1, 1, 1, 5, 1, 152, 8, 1, 10, 1, 12, 1, 155, 9, 1, 1, 2,
		1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 3, 4, 166, 8, 4, 1, 4,
		4, 4, 169, 8, 4, 11, 4, 12, 4, 170, 1, 4, 3, 4, 174, 8, 4, 1, 4, 5, 4,
		177, 8, 4, 10, 4, 12, 4, 180, 9, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5,
		5, 5, 188, 8, 5, 10, 5, 12, 5, 191, 9, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6,
		4, 6, 198, 8, 6, 11, 6, 12, 6, 199, 1, 6, 3, 6, 203, 8, 6, 1, 6, 5, 6,
		206, 8, 6, 10, 6, 12, 6, 209, 9, 6, 1, 6, 3, 6, 212, 8, 6, 1, 6, 1, 6,
		1, 7, 1, 7, 5, 7, 218, 8, 7, 10, 7, 12, 7, 221, 9, 7, 1, 7, 1, 7, 1, 8,
		1, 8, 1, 8, 3, 8, 228, 8, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 11, 1,
		11, 1, 12, 3, 12, 238, 8, 12, 1, 12, 1, 12, 1, 13, 3, 13, 243, 8, 13, 1,
		13, 3, 13, 246, 8, 13, 1, 14, 1, 14, 3, 14, 250, 8, 14, 1, 14, 1, 14, 3,
		14, 254, 8, 14, 3, 14, 256, 8, 14, 1, 15, 1, 15, 1, 15, 4, 15, 261, 8,
		15, 11, 15, 12, 15, 262, 1, 16, 1, 16, 1, 16, 1, 16, 3, 16, 269, 8, 16,
		3, 16, 271, 8, 16, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 19, 1,
		19, 3, 19, 281, 8, 19, 1, 19, 1, 19, 3, 19, 285, 8, 19, 1, 19, 3, 19, 288,
		8, 19, 1, 19, 3, 19, 291, 8, 19, 5, 19, 293, 8, 19, 10, 19, 12, 19, 296,
		9, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 3, 20, 303, 8, 20, 3, 20, 305,
		8, 20, 1, 21, 1, 21, 1, 21, 3, 21, 310, 8, 21, 1, 22, 1, 22, 1, 22, 1,
		23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 5, 24, 322, 8, 24, 10, 24,
		12, 24, 325, 9, 24, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 3, 26, 333,
		8, 26, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 5, 28, 343,
		8, 28, 10, 28, 12, 28, 346, 9, 28, 1, 28, 1, 28, 3, 28, 350, 8, 28, 1,
		29, 1, 29, 3, 29, 354, 8, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 3, 30,
		361, 8, 30, 3, 30, 363, 8, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31,
		3, 31, 371, 8, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 3, 32, 379,
		8, 32, 5, 32, 381, 8, 32, 10, 32, 12, 32, 384, 9, 32, 1, 33, 3, 33, 387,
		8, 33, 1, 34, 1, 34, 1, 34, 5, 34, 392, 8, 34, 10, 34, 12, 34, 395, 9,
		34, 1, 35, 1, 35, 1, 36, 1, 36, 3, 36, 401, 8, 36, 1, 37, 1, 37, 3, 37,
		405, 8, 37, 1, 38, 1, 38, 1, 38, 1, 38, 1, 39, 1, 39, 4, 39, 413, 8, 39,
		11, 39, 12, 39, 414, 1, 39, 1, 39, 1, 40, 1, 40, 3, 40, 421, 8, 40, 1,
		41, 1, 41, 3, 41, 425, 8, 41, 1, 42, 1, 42, 3, 42, 429, 8, 42, 1, 43, 1,
		43, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 1, 44, 3, 44, 439, 8, 44, 1, 45,
		1, 45, 1, 46, 1, 46, 1, 46, 5, 46, 446, 8, 46, 10, 46, 12, 46, 449, 9,
		46, 1, 47, 1, 47, 1, 47, 5, 47, 454, 8, 47, 10, 47, 12, 47, 457, 9, 47,
		1, 48, 1, 48, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1,
		49, 1, 49, 1, 49, 1, 49, 1, 49, 3, 49, 474, 8, 49, 1, 50, 1, 50, 1, 51,
		1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 5, 51, 485, 8, 51, 10, 51, 12,
		51, 488, 9, 51, 1, 52, 1, 52, 1, 52, 1, 52, 1, 52, 5, 52, 495, 8, 52, 10,
		52, 12, 52, 498, 9, 52, 1, 53, 1, 53, 1, 53, 1, 53, 1, 53, 1, 53, 1, 53,
		3, 53, 507, 8, 53, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 3,
		54, 516, 8, 54, 1, 55, 1, 55, 1, 55, 1, 55, 1, 56, 1, 56, 1, 56, 1, 56,
		1, 56, 1, 56, 1, 56, 1, 56, 1, 56, 1, 56, 1, 56, 1, 56, 1, 56, 1, 56, 1,
		56, 1, 56, 1, 56, 1, 56, 1, 56, 1, 56, 1, 56, 1, 56, 1, 56, 1, 56, 1, 56,
		1, 56, 1, 56, 1, 56, 1, 56, 1, 56, 1, 56, 1, 56, 1, 56, 1, 56, 1, 56, 1,
		56, 1, 56, 1, 56, 1, 56, 1, 56, 1, 56, 1, 56, 1, 56, 1, 56, 1, 56, 1, 56,
		1, 56, 1, 56, 1, 56, 1, 56, 1, 56, 1, 56, 1, 56, 1, 56, 1, 56, 3, 56, 577,
		8, 56, 1, 57, 1, 57, 1, 57, 1, 57, 1, 57, 1, 57, 1, 57, 3, 57, 586, 8,
		57, 1, 57, 1, 57, 1, 58, 1, 58, 3, 58, 592, 8, 58, 1, 59, 1, 59, 1, 59,
		1, 59, 3, 59, 598, 8, 59, 1, 60, 1, 60, 1, 60, 3, 60, 603, 8, 60, 1, 61,
		1, 61, 1, 62, 1, 62, 1, 63, 1, 63, 1, 64, 1, 64, 1, 65, 1, 65, 1, 66, 1,
		66, 3, 66, 617, 8, 66, 1, 67, 1, 67, 1, 68, 1, 68, 1, 68, 0, 0, 69, 0,
		2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38,
		40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74,
		76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 100, 102, 104, 106, 108,
		110, 112, 114, 116, 118, 120, 122, 124, 126, 128, 130, 132, 134, 136, 0,
		10, 1, 0, 4, 5, 1, 0, 15, 16, 1, 0, 63, 64, 1, 0, 66, 68, 1, 0, 69, 71,
		1, 0, 72, 74, 1, 0, 57, 58, 1, 0, 76, 77, 1, 0, 60, 61, 2, 0, 62, 62, 82,
		82, 651, 0, 138, 1, 0, 0, 0, 2, 148, 1, 0, 0, 0, 4, 156, 1, 0, 0, 0, 6,
		159, 1, 0, 0, 0, 8, 163, 1, 0, 0, 0, 10, 184, 1, 0, 0, 0, 12, 195, 1, 0,
		0, 0, 14, 215, 1, 0, 0, 0, 16, 224, 1, 0, 0, 0, 18, 229, 1, 0, 0, 0, 20,
		231, 1, 0, 0, 0, 22, 234, 1, 0, 0, 0, 24, 237, 1, 0, 0, 0, 26, 242, 1,
		0, 0, 0, 28, 255, 1, 0, 0, 0, 30, 257, 1, 0, 0, 0, 32, 270, 1, 0, 0, 0,
		34, 272, 1, 0, 0, 0, 36, 275, 1, 0, 0, 0, 38, 278, 1, 0, 0, 0, 40, 299,
		1, 0, 0, 0, 42, 309, 1, 0, 0, 0, 44, 311, 1, 0, 0, 0, 46, 314, 1, 0, 0,
		0, 48, 318, 1, 0, 0, 0, 50, 326, 1, 0, 0, 0, 52, 332, 1, 0, 0, 0, 54, 334,
		1, 0, 0, 0, 56, 349, 1, 0, 0, 0, 58, 351, 1, 0, 0, 0, 60, 357, 1, 0, 0,
		0, 62, 370, 1, 0, 0, 0, 64, 372, 1, 0, 0, 0, 66, 386, 1, 0, 0, 0, 68, 388,
		1, 0, 0, 0, 70, 396, 1, 0, 0, 0, 72, 400, 1, 0, 0, 0, 74, 404, 1, 0, 0,
		0, 76, 406, 1, 0, 0, 0, 78, 410, 1, 0, 0, 0, 80, 420, 1, 0, 0, 0, 82, 424,
		1, 0, 0, 0, 84, 428, 1, 0, 0, 0, 86, 430, 1, 0, 0, 0, 88, 438, 1, 0, 0,
		0, 90, 440, 1, 0, 0, 0, 92, 442, 1, 0, 0, 0, 94, 450, 1, 0, 0, 0, 96, 458,
		1, 0, 0, 0, 98, 460, 1, 0, 0, 0, 100, 475, 1, 0, 0, 0, 102, 477, 1, 0,
		0, 0, 104, 489, 1, 0, 0, 0, 106, 506, 1, 0, 0, 0, 108, 515, 1, 0, 0, 0,
		110, 517, 1, 0, 0, 0, 112, 576, 1, 0, 0, 0, 114, 578, 1, 0, 0, 0, 116,
		589, 1, 0, 0, 0, 118, 593, 1, 0, 0, 0, 120, 602, 1, 0, 0, 0, 122, 604,
		1, 0, 0, 0, 124, 606, 1, 0, 0, 0, 126, 608, 1, 0, 0, 0, 128, 610, 1, 0,
		0, 0, 130, 612, 1, 0, 0, 0, 132, 616, 1, 0, 0, 0, 134, 618, 1, 0, 0, 0,
		136, 620, 1, 0, 0, 0, 138, 143, 3, 2, 1, 0, 139, 144, 3, 8, 4, 0, 140,
		144, 3, 10, 5, 0, 141, 144, 3, 12, 6, 0, 142, 144, 3, 14, 7, 0, 143, 139,
		1, 0, 0, 0, 143, 140, 1, 0, 0, 0, 143, 141, 1, 0, 0, 0, 143, 142, 1, 0,
		0, 0, 144, 145, 1, 0, 0, 0, 145, 146, 5, 0, 0, 1, 146, 1, 1, 0, 0, 0, 147,
		149, 3, 4, 2, 0, 148, 147, 1, 0, 0, 0, 148, 149, 1, 0, 0, 0, 149, 153,
		1, 0, 0, 0, 150, 152, 3, 6, 3, 0, 151, 150, 1, 0, 0, 0, 152, 155, 1, 0,
		0, 0, 153, 151, 1, 0, 0, 0, 153, 154, 1, 0, 0, 0, 154, 3, 1, 0, 0, 0, 155,
		153, 1, 0, 0, 0, 156, 157, 5, 1, 0, 0, 157, 158, 5, 59, 0, 0, 158, 5, 1,
		0, 0, 0, 159, 160, 5, 2, 0, 0, 160, 161, 5, 60, 0, 0, 161, 162, 5, 59,
		0, 0, 162, 7, 1, 0, 0, 0, 163, 165, 5, 3, 0, 0, 164, 166, 7, 0, 0, 0, 165,
		164, 1, 0, 0, 0, 165, 166, 1, 0, 0, 0, 166, 173, 1, 0, 0, 0, 167, 169,
		3, 86, 43, 0, 168, 167, 1, 0, 0, 0, 169, 170, 1, 0, 0, 0, 170, 168, 1,
		0, 0, 0, 170, 171, 1, 0, 0, 0, 171, 174, 1, 0, 0, 0, 172, 174, 5, 6, 0,
		0, 173, 168, 1, 0, 0, 0, 173, 172, 1, 0, 0, 0, 174, 178, 1, 0, 0, 0, 175,
		177, 3, 16, 8, 0, 176, 175, 1, 0, 0, 0, 177, 180, 1, 0, 0, 0, 178, 176,
		1, 0, 0, 0, 178, 179, 1, 0, 0, 0, 179, 181, 1, 0, 0, 0, 180, 178, 1, 0,
		0, 0, 181, 182, 3, 24, 12, 0, 182, 183, 3, 26, 13, 0, 183, 9, 1, 0, 0,
		0, 184, 185, 5, 7, 0, 0, 185, 189, 3, 58, 29, 0, 186, 188, 3, 16, 8, 0,
		187, 186, 1, 0, 0, 0, 188, 191, 1, 0, 0, 0, 189, 187, 1, 0, 0, 0, 189,
		190, 1, 0, 0, 0, 190, 192, 1, 0, 0, 0, 191, 189, 1, 0, 0, 0, 192, 193,
		3, 24, 12, 0, 193, 194, 3, 26, 13, 0, 194, 11, 1, 0, 0, 0, 195, 202, 5,
		8, 0, 0, 196, 198, 3, 84, 42, 0, 197, 196, 1, 0, 0, 0, 198, 199, 1, 0,
		0, 0, 199, 197, 1, 0, 0, 0, 199, 200, 1, 0, 0, 0, 200, 203, 1, 0, 0, 0,
		201, 203, 5, 6, 0, 0, 202, 197, 1, 0, 0, 0, 202, 201, 1, 0, 0, 0, 203,
		207, 1, 0, 0, 0, 204, 206, 3, 16, 8, 0, 205, 204, 1, 0, 0, 0, 206, 209,
		1, 0, 0, 0, 207, 205, 1, 0, 0, 0, 207, 208, 1, 0, 0, 0, 208, 211, 1, 0,
		0, 0, 209, 207, 1, 0, 0, 0, 210, 212, 3, 24, 12, 0, 211, 210, 1, 0, 0,
		0, 211, 212, 1, 0, 0, 0, 212, 213, 1, 0, 0, 0, 213, 214, 3, 26, 13, 0,
		214, 13, 1, 0, 0, 0, 215, 219, 5, 9, 0, 0, 216, 218, 3, 16, 8, 0, 217,
		216, 1, 0, 0, 0, 218, 221, 1, 0, 0, 0, 219, 217, 1, 0, 0, 0, 219, 220,
		1, 0, 0, 0, 220, 222, 1, 0, 0, 0, 221, 219, 1, 0, 0, 0, 222, 223, 3, 24,
		12, 0, 223, 15, 1, 0, 0, 0, 224, 227, 5, 10, 0, 0, 225, 228, 3, 18, 9,
		0, 226, 228, 3, 20, 10, 0, 227, 225, 1, 0, 0, 0, 227, 226, 1, 0, 0, 0,
		228, 17, 1, 0, 0, 0, 229, 230, 3, 22, 11, 0, 230, 19, 1, 0, 0, 0, 231,
		232, 5, 11, 0, 0, 232, 233, 3, 22, 11, 0, 233, 21, 1, 0, 0, 0, 234, 235,
		3, 132, 66, 0, 235, 23, 1, 0, 0, 0, 236, 238, 5, 12, 0, 0, 237, 236, 1,
		0, 0, 0, 237, 238, 1, 0, 0, 0, 238, 239, 1, 0, 0, 0, 239, 240, 3, 38, 19,
		0, 240, 25, 1, 0, 0, 0, 241, 243, 3, 30, 15, 0, 242, 241, 1, 0, 0, 0, 242,
		243, 1, 0, 0, 0, 243, 245, 1, 0, 0, 0, 244, 246, 3, 28, 14, 0, 245, 244,
		1, 0, 0, 0, 245, 246, 1, 0, 0, 0, 246, 27, 1, 0, 0, 0, 247, 249, 3, 34,
		17, 0, 248, 250, 3, 36, 18, 0, 249, 248, 1, 0, 0, 0, 249, 250, 1, 0, 0,
		0, 250, 256, 1, 0, 0, 0, 251, 253, 3, 36, 18, 0, 252, 254, 3, 34, 17, 0,
		253, 252, 1, 0, 0, 0, 253, 254, 1, 0, 0, 0, 254, 256, 1, 0, 0, 0, 255,
		247, 1, 0, 0, 0, 255, 251, 1, 0, 0, 0, 256, 29, 1, 0, 0, 0, 257, 258, 5,
		13, 0, 0, 258, 260, 5, 14, 0, 0, 259, 261, 3, 32, 16, 0, 260, 259, 1, 0,
		0, 0, 261, 262, 1, 0, 0, 0, 262, 260, 1, 0, 0, 0, 262, 263, 1, 0, 0, 0,
		263, 31, 1, 0, 0, 0, 264, 265, 7, 1, 0, 0, 265, 271, 3, 110, 55, 0, 266,
		269, 3, 52, 26, 0, 267, 269, 3, 86, 43, 0, 268, 266, 1, 0, 0, 0, 268, 267,
		1, 0, 0, 0, 269, 271, 1, 0, 0, 0, 270, 264, 1, 0, 0, 0, 270, 268, 1, 0,
		0, 0, 271, 33, 1, 0, 0, 0, 272, 273, 5, 17, 0, 0, 273, 274, 5, 66, 0, 0,
		274, 35, 1, 0, 0, 0, 275, 276, 5, 18, 0, 0, 276, 277, 5, 66, 0, 0, 277,
		37, 1, 0, 0, 0, 278, 280, 5, 19, 0, 0, 279, 281, 3, 40, 20, 0, 280, 279,
		1, 0, 0, 0, 280, 281, 1, 0, 0, 0, 281, 294, 1, 0, 0, 0, 282, 285, 3, 42,
		21, 0, 283, 285, 3, 50, 25, 0, 284, 282, 1, 0, 0, 0, 284, 283, 1, 0, 0,
		0, 285, 287, 1, 0, 0, 0, 286, 288, 5, 20, 0, 0, 287, 286, 1, 0, 0, 0, 287,
		288, 1, 0, 0, 0, 288, 290, 1, 0, 0, 0, 289, 291, 3, 40, 20, 0, 290, 289,
		1, 0, 0, 0, 290, 291, 1, 0, 0, 0, 291, 293, 1, 0, 0, 0, 292, 284, 1, 0,
		0, 0, 293, 296, 1, 0, 0, 0, 294, 292, 1, 0, 0, 0, 294, 295, 1, 0, 0, 0,
		295, 297, 1, 0, 0, 0, 296, 294, 1, 0, 0, 0, 297, 298, 5, 21, 0, 0, 298,
		39, 1, 0, 0, 0, 299, 304, 3, 62, 31, 0, 300, 302, 5, 20, 0, 0, 301, 303,
		3, 40, 20, 0, 302, 301, 1, 0, 0, 0, 302, 303, 1, 0, 0, 0, 303, 305, 1,
		0, 0, 0, 304, 300, 1, 0, 0, 0, 304, 305, 1, 0, 0, 0, 305, 41, 1, 0, 0,
		0, 306, 310, 3, 44, 22, 0, 307, 310, 3, 48, 24, 0, 308, 310, 3, 46, 23,
		0, 309, 306, 1, 0, 0, 0, 309, 307, 1, 0, 0, 0, 309, 308, 1, 0, 0, 0, 310,
		43, 1, 0, 0, 0, 311, 312, 5, 22, 0, 0, 312, 313, 3, 38, 19, 0, 313, 45,
		1, 0, 0, 0, 314, 315, 5, 23, 0, 0, 315, 316, 3, 84, 42, 0, 316, 317, 3,
		38, 19, 0, 317, 47, 1, 0, 0, 0, 318, 323, 3, 38, 19, 0, 319, 320, 5, 24,
		0, 0, 320, 322, 3, 38, 19, 0, 321, 319, 1, 0, 0, 0, 322, 325, 1, 0, 0,
		0, 323, 321, 1, 0, 0, 0, 323, 324, 1, 0, 0, 0, 324, 49, 1, 0, 0, 0, 325,
		323, 1, 0, 0, 0, 326, 327, 5, 25, 0, 0, 327, 328, 3, 52, 26, 0, 328, 51,
		1, 0, 0, 0, 329, 333, 3, 110, 55, 0, 330, 333, 3, 112, 56, 0, 331, 333,
		3, 54, 27, 0, 332, 329, 1, 0, 0, 0, 332, 330, 1, 0, 0, 0, 332, 331, 1,
		0, 0, 0, 333, 53, 1, 0, 0, 0, 334, 335, 3, 132, 66, 0, 335, 336, 3, 56,
		28, 0, 336, 55, 1, 0, 0, 0, 337, 350, 5, 81, 0, 0, 338, 339, 5, 26, 0,
		0, 339, 344, 3, 90, 45, 0, 340, 341, 5, 27, 0, 0, 341, 343, 3, 90, 45,
		0, 342, 340, 1, 0, 0, 0, 343, 346, 1, 0, 0, 0, 344, 342, 1, 0, 0, 0, 344,
		345, 1, 0, 0, 0, 345, 347, 1, 0, 0, 0, 346, 344, 1, 0, 0, 0, 347, 348,
		5, 28, 0, 0, 348, 350, 1, 0, 0, 0, 349, 337, 1, 0, 0, 0, 349, 338, 1, 0,
		0, 0, 350, 57, 1, 0, 0, 0, 351, 353, 5, 19, 0, 0, 352, 354, 3, 60, 30,
		0, 353, 352, 1, 0, 0, 0, 353, 354, 1, 0, 0, 0, 354, 355, 1, 0, 0, 0, 355,
		356, 5, 21, 0, 0, 356, 59, 1, 0, 0, 0, 357, 362, 3, 62, 31, 0, 358, 360,
		5, 20, 0, 0, 359, 361, 3, 60, 30, 0, 360, 359, 1, 0, 0, 0, 360, 361, 1,
		0, 0, 0, 361, 363, 1, 0, 0, 0, 362, 358, 1, 0, 0, 0, 362, 363, 1, 0, 0,
		0, 363, 61, 1, 0, 0, 0, 364, 365, 3, 82, 41, 0, 365, 366, 3, 64, 32, 0,
		366, 371, 1, 0, 0, 0, 367, 368, 3, 74, 37, 0, 368, 369, 3, 66, 33, 0, 369,
		371, 1, 0, 0, 0, 370, 364, 1, 0, 0, 0, 370, 367, 1, 0, 0, 0, 371, 63, 1,
		0, 0, 0, 372, 373, 3, 72, 36, 0, 373, 382, 3, 68, 34, 0, 374, 378, 5, 29,
		0, 0, 375, 376, 3, 72, 36, 0, 376, 377, 3, 68, 34, 0, 377, 379, 1, 0, 0,
		0, 378, 375, 1, 0, 0, 0, 378, 379, 1, 0, 0, 0, 379, 381, 1, 0, 0, 0, 380,
		374, 1, 0, 0, 0, 381, 384, 1, 0, 0, 0, 382, 380, 1, 0, 0, 0, 382, 383,
		1, 0, 0, 0, 383, 65, 1, 0, 0, 0, 384, 382, 1, 0, 0, 0, 385, 387, 3, 64,
		32, 0, 386, 385, 1, 0, 0, 0, 386, 387, 1, 0, 0, 0, 387, 67, 1, 0, 0, 0,
		388, 393, 3, 70, 35, 0, 389, 390, 5, 27, 0, 0, 390, 392, 3, 70, 35, 0,
		391, 389, 1, 0, 0, 0, 392, 395, 1, 0, 0, 0, 393, 391, 1, 0, 0, 0, 393,
		394, 1, 0, 0, 0, 394, 69, 1, 0, 0, 0, 395, 393, 1, 0, 0, 0, 396, 397, 3,
		80, 40, 0, 397, 71, 1, 0, 0, 0, 398, 401, 3, 84, 42, 0, 399, 401, 5, 30,
		0, 0, 400, 398, 1, 0, 0, 0, 400, 399, 1, 0, 0, 0, 401, 73, 1, 0, 0, 0,
		402, 405, 3, 78, 39, 0, 403, 405, 3, 76, 38, 0, 404, 402, 1, 0, 0, 0, 404,
		403, 1, 0, 0, 0, 405, 75, 1, 0, 0, 0, 406, 407, 5, 31, 0, 0, 407, 408,
		3, 64, 32, 0, 408, 409, 5, 32, 0, 0, 409, 77, 1, 0, 0, 0, 410, 412, 5,
		26, 0, 0, 411, 413, 3, 80, 40, 0, 412, 411, 1, 0, 0, 0, 413, 414, 1, 0,
		0, 0, 414, 412, 1, 0, 0, 0, 414, 415, 1, 0, 0, 0, 415, 416, 1, 0, 0, 0,
		416, 417, 5, 28, 0, 0, 417, 79, 1, 0, 0, 0, 418, 421, 3, 82, 41, 0, 419,
		421, 3, 74, 37, 0, 420, 418, 1, 0, 0, 0, 420, 419, 1, 0, 0, 0, 421, 81,
		1, 0, 0, 0, 422, 425, 3, 86, 43, 0, 423, 425, 3, 88, 44, 0, 424, 422, 1,
		0, 0, 0, 424, 423, 1, 0, 0, 0, 425, 83, 1, 0, 0, 0, 426, 429, 3, 86, 43,
		0, 427, 429, 3, 132, 66, 0, 428, 426, 1, 0, 0, 0, 428, 427, 1, 0, 0, 0,
		429, 85, 1, 0, 0, 0, 430, 431, 7, 2, 0, 0, 431, 87, 1, 0, 0, 0, 432, 439,
		3, 132, 66, 0, 433, 439, 3, 118, 59, 0, 434, 439, 3, 120, 60, 0, 435, 439,
		3, 128, 64, 0, 436, 439, 3, 136, 68, 0, 437, 439, 5, 81, 0, 0, 438, 432,
		1, 0, 0, 0, 438, 433, 1, 0, 0, 0, 438, 434, 1, 0, 0, 0, 438, 435, 1, 0,
		0, 0, 438, 436, 1, 0, 0, 0, 438, 437, 1, 0, 0, 0, 439, 89, 1, 0, 0, 0,
		440, 441, 3, 92, 46, 0, 441, 91, 1, 0, 0, 0, 442, 447, 3, 94, 47, 0, 443,
		444, 5, 33, 0, 0, 444, 446, 3, 94, 47, 0, 445, 443, 1, 0, 0, 0, 446, 449,
		1, 0, 0, 0, 447, 445, 1, 0, 0, 0, 447, 448, 1, 0, 0, 0, 448, 93, 1, 0,
		0, 0, 449, 447, 1, 0, 0, 0, 450, 455, 3, 96, 48, 0, 451, 452, 5, 34, 0,
		0, 452, 454, 3, 96, 48, 0, 453, 451, 1, 0, 0, 0, 454, 457, 1, 0, 0, 0,
		455, 453, 1, 0, 0, 0, 455, 456, 1, 0, 0, 0, 456, 95, 1, 0, 0, 0, 457, 455,
		1, 0, 0, 0, 458, 459, 3, 98, 49, 0, 459, 97, 1, 0, 0, 0, 460, 473, 3, 100,
		50, 0, 461, 462, 5, 35, 0, 0, 462, 474, 3, 100, 50, 0, 463, 464, 5, 36,
		0, 0, 464, 474, 3, 100, 50, 0, 465, 466, 5, 37, 0, 0, 466, 474, 3, 100,
		50, 0, 467, 468, 5, 38, 0, 0, 468, 474, 3, 100, 50, 0, 469, 470, 5, 39,
		0, 0, 470, 474, 3, 100, 50, 0, 471, 472, 5, 40, 0, 0, 472, 474, 3, 100,
		50, 0, 473, 461, 1, 0, 0, 0, 473, 463, 1, 0, 0, 0, 473, 465, 1, 0, 0, 0,
		473, 467, 1, 0, 0, 0, 473, 469, 1, 0, 0, 0, 473, 471, 1, 0, 0, 0, 473,
		474, 1, 0, 0, 0, 474, 99, 1, 0, 0, 0, 475, 476, 3, 102, 51, 0, 476, 101,
		1, 0, 0, 0, 477, 486, 3, 104, 52, 0, 478, 479, 5, 41, 0, 0, 479, 485, 3,
		104, 52, 0, 480, 481, 5, 42, 0, 0, 481, 485, 3, 104, 52, 0, 482, 485, 3,
		124, 62, 0, 483, 485, 3, 126, 63, 0, 484, 478, 1, 0, 0, 0, 484, 480, 1,
		0, 0, 0, 484, 482, 1, 0, 0, 0, 484, 483, 1, 0, 0, 0, 485, 488, 1, 0, 0,
		0, 486, 484, 1, 0, 0, 0, 486, 487, 1, 0, 0, 0, 487, 103, 1, 0, 0, 0, 488,
		486, 1, 0, 0, 0, 489, 496, 3, 106, 53, 0, 490, 491, 5, 6, 0, 0, 491, 495,
		3, 106, 53, 0, 492, 493, 5, 43, 0, 0, 493, 495, 3, 106, 53, 0, 494, 490,
		1, 0, 0, 0, 494, 492, 1, 0, 0, 0, 495, 498, 1, 0, 0, 0, 496, 494, 1, 0,
		0, 0, 496, 497, 1, 0, 0, 0, 497, 105, 1, 0, 0, 0, 498, 496, 1, 0, 0, 0,
		499, 500, 5, 44, 0, 0, 500, 507, 3, 108, 54, 0, 501, 502, 5, 41, 0, 0,
		502, 507, 3, 108, 54, 0, 503, 504, 5, 42, 0, 0, 504, 507, 3, 108, 54, 0,
		505, 507, 3, 108, 54, 0, 506, 499, 1, 0, 0, 0, 506, 501, 1, 0, 0, 0, 506,
		503, 1, 0, 0, 0, 506, 505, 1, 0, 0, 0, 507, 107, 1, 0, 0, 0, 508, 516,
		3, 110, 55, 0, 509, 516, 3, 112, 56, 0, 510, 516, 3, 116, 58, 0, 511, 516,
		3, 118, 59, 0, 512, 516, 3, 120, 60, 0, 513, 516, 3, 128, 64, 0, 514, 516,
		3, 86, 43, 0, 515, 508, 1, 0, 0, 0, 515, 509, 1, 0, 0, 0, 515, 510, 1,
		0, 0, 0, 515, 511, 1, 0, 0, 0, 515, 512, 1, 0, 0, 0, 515, 513, 1, 0, 0,
		0, 515, 514, 1, 0, 0, 0, 516, 109, 1, 0, 0, 0, 517, 518, 5, 26, 0, 0, 518,
		519, 3, 90, 45, 0, 519, 520, 5, 28, 0, 0, 520, 111, 1, 0, 0, 0, 521, 522,
		5, 45, 0, 0, 522, 523, 5, 26, 0, 0, 523, 524, 3, 90, 45, 0, 524, 525, 5,
		28, 0, 0, 525, 577, 1, 0, 0, 0, 526, 527, 5, 46, 0, 0, 527, 528, 5, 26,
		0, 0, 528, 529, 3, 90, 45, 0, 529, 530, 5, 28, 0, 0, 530, 577, 1, 0, 0,
		0, 531, 532, 5, 47, 0, 0, 532, 533, 5, 26, 0, 0, 533, 534, 3, 90, 45, 0,
		534, 535, 5, 27, 0, 0, 535, 536, 3, 90, 45, 0, 536, 537, 5, 28, 0, 0, 537,
		577, 1, 0, 0, 0, 538, 539, 5, 48, 0, 0, 539, 540, 5, 26, 0, 0, 540, 541,
		3, 90, 45, 0, 541, 542, 5, 28, 0, 0, 542, 577, 1, 0, 0, 0, 543, 544, 5,
		49, 0, 0, 544, 545, 5, 26, 0, 0, 545, 546, 3, 86, 43, 0, 546, 547, 5, 28,
		0, 0, 547, 577, 1, 0, 0, 0, 548, 549, 5, 50, 0, 0, 549, 550, 5, 26, 0,
		0, 550, 551, 3, 90, 45, 0, 551, 552, 5, 27, 0, 0, 552, 553, 3, 90, 45,
		0, 553, 554, 5, 28, 0, 0, 554, 577, 1, 0, 0, 0, 555, 556, 5, 51, 0, 0,
		556, 557, 5, 26, 0, 0, 557, 558, 3, 90, 45, 0, 558, 559, 5, 28, 0, 0, 559,
		577, 1, 0, 0, 0, 560, 561, 5, 52, 0, 0, 561, 562, 5, 26, 0, 0, 562, 563,
		3, 90, 45, 0, 563, 564, 5, 28, 0, 0, 564, 577, 1, 0, 0, 0, 565, 566, 5,
		53, 0, 0, 566, 567, 5, 26, 0, 0, 567, 568, 3, 90, 45, 0, 568, 569, 5, 28,
		0, 0, 569, 577, 1, 0, 0, 0, 570, 571, 5, 54, 0, 0, 571, 572, 5, 26, 0,
		0, 572, 573, 3, 90, 45, 0, 573, 574, 5, 28, 0, 0, 574, 577, 1, 0, 0, 0,
		575, 577, 3, 114, 57, 0, 576, 521, 1, 0, 0, 0, 576, 526, 1, 0, 0, 0, 576,
		531, 1, 0, 0, 0, 576, 538, 1, 0, 0, 0, 576, 543, 1, 0, 0, 0, 576, 548,
		1, 0, 0, 0, 576, 555, 1, 0, 0, 0, 576, 560, 1, 0, 0, 0, 576, 565, 1, 0,
		0, 0, 576, 570, 1, 0, 0, 0, 576, 575, 1, 0, 0, 0, 577, 113, 1, 0, 0, 0,
		578, 579, 5, 55, 0, 0, 579, 580, 5, 26, 0, 0, 580, 581, 3, 90, 45, 0, 581,
		582, 5, 27, 0, 0, 582, 585, 3, 90, 45, 0, 583, 584, 5, 27, 0, 0, 584, 586,
		3, 90, 45, 0, 585, 583, 1, 0, 0, 0, 585, 586, 1, 0, 0, 0, 586, 587, 1,
		0, 0, 0, 587, 588, 5, 28, 0, 0, 588, 115, 1, 0, 0, 0, 589, 591, 3, 132,
		66, 0, 590, 592, 3, 56, 28, 0, 591, 590, 1, 0, 0, 0, 591, 592, 1, 0, 0,
		0, 592, 117, 1, 0, 0, 0, 593, 597, 3, 130, 65, 0, 594, 598, 5, 65, 0, 0,
		595, 596, 5, 56, 0, 0, 596, 598, 3, 132, 66, 0, 597, 594, 1, 0, 0, 0, 597,
		595, 1, 0, 0, 0, 597, 598, 1, 0, 0, 0, 598, 119, 1, 0, 0, 0, 599, 603,
		3, 122, 61, 0, 600, 603, 3, 124, 62, 0, 601, 603, 3, 126, 63, 0, 602, 599,
		1, 0, 0, 0, 602, 600, 1, 0, 0, 0, 602, 601, 1, 0, 0, 0, 603, 121, 1, 0,
		0, 0, 604, 605, 7, 3, 0, 0, 605, 123, 1, 0, 0, 0, 606, 607, 7, 4, 0, 0,
		607, 125, 1, 0, 0, 0, 608, 609, 7, 5, 0, 0, 609, 127, 1, 0, 0, 0, 610,
		611, 7, 6, 0, 0, 611, 129, 1, 0, 0, 0, 612, 613, 7, 7, 0, 0, 613, 131,
		1, 0, 0, 0, 614, 617, 5, 59, 0, 0, 615, 617, 3, 134, 67, 0, 616, 614, 1,
		0, 0, 0, 616, 615, 1, 0, 0, 0, 617, 133, 1, 0, 0, 0, 618, 619, 7, 8, 0,
		0, 619, 135, 1, 0, 0, 0, 620, 621, 7, 9, 0, 0, 621, 137, 1, 0, 0, 0, 65,
		143, 148, 153, 165, 170, 173, 178, 189, 199, 202, 207, 211, 219, 227, 237,
		242, 245, 249, 253, 255, 262, 268, 270, 280, 284, 287, 290, 294, 302, 304,
		309, 323, 332, 344, 349, 353, 360, 362, 370, 378, 382, 386, 393, 400, 404,
		414, 420, 424, 428, 438, 447, 455, 473, 484, 486, 494, 496, 506, 515, 576,
		585, 591, 597, 602, 616,
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

// SparqlParserInit initializes any static state used to implement SparqlParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSparqlParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SparqlParserInit() {
	staticData := &SparqlParserStaticData
	staticData.once.Do(sparqlParserInit)
}

// NewSparqlParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSparqlParser(input antlr.TokenStream) *SparqlParser {
	SparqlParserInit()
	this := new(SparqlParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &SparqlParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Sparql.g4"

	return this
}

// SparqlParser tokens.
const (
	SparqlParserEOF                  = antlr.TokenEOF
	SparqlParserT__0                 = 1
	SparqlParserT__1                 = 2
	SparqlParserT__2                 = 3
	SparqlParserT__3                 = 4
	SparqlParserT__4                 = 5
	SparqlParserT__5                 = 6
	SparqlParserT__6                 = 7
	SparqlParserT__7                 = 8
	SparqlParserT__8                 = 9
	SparqlParserT__9                 = 10
	SparqlParserT__10                = 11
	SparqlParserT__11                = 12
	SparqlParserT__12                = 13
	SparqlParserT__13                = 14
	SparqlParserT__14                = 15
	SparqlParserT__15                = 16
	SparqlParserT__16                = 17
	SparqlParserT__17                = 18
	SparqlParserT__18                = 19
	SparqlParserT__19                = 20
	SparqlParserT__20                = 21
	SparqlParserT__21                = 22
	SparqlParserT__22                = 23
	SparqlParserT__23                = 24
	SparqlParserT__24                = 25
	SparqlParserT__25                = 26
	SparqlParserT__26                = 27
	SparqlParserT__27                = 28
	SparqlParserT__28                = 29
	SparqlParserT__29                = 30
	SparqlParserT__30                = 31
	SparqlParserT__31                = 32
	SparqlParserT__32                = 33
	SparqlParserT__33                = 34
	SparqlParserT__34                = 35
	SparqlParserT__35                = 36
	SparqlParserT__36                = 37
	SparqlParserT__37                = 38
	SparqlParserT__38                = 39
	SparqlParserT__39                = 40
	SparqlParserT__40                = 41
	SparqlParserT__41                = 42
	SparqlParserT__42                = 43
	SparqlParserT__43                = 44
	SparqlParserT__44                = 45
	SparqlParserT__45                = 46
	SparqlParserT__46                = 47
	SparqlParserT__47                = 48
	SparqlParserT__48                = 49
	SparqlParserT__49                = 50
	SparqlParserT__50                = 51
	SparqlParserT__51                = 52
	SparqlParserT__52                = 53
	SparqlParserT__53                = 54
	SparqlParserT__54                = 55
	SparqlParserT__55                = 56
	SparqlParserT__56                = 57
	SparqlParserT__57                = 58
	SparqlParserIRI_REF              = 59
	SparqlParserPNAME_NS             = 60
	SparqlParserPNAME_LN             = 61
	SparqlParserBLANK_NODE_LABEL     = 62
	SparqlParserVAR1                 = 63
	SparqlParserVAR2                 = 64
	SparqlParserLANGTAG              = 65
	SparqlParserINTEGER              = 66
	SparqlParserDECIMAL              = 67
	SparqlParserDOUBLE               = 68
	SparqlParserINTEGER_POSITIVE     = 69
	SparqlParserDECIMAL_POSITIVE     = 70
	SparqlParserDOUBLE_POSITIVE      = 71
	SparqlParserINTEGER_NEGATIVE     = 72
	SparqlParserDECIMAL_NEGATIVE     = 73
	SparqlParserDOUBLE_NEGATIVE      = 74
	SparqlParserEXPONENT             = 75
	SparqlParserSTRING_LITERAL1      = 76
	SparqlParserSTRING_LITERAL2      = 77
	SparqlParserSTRING_LITERAL_LONG1 = 78
	SparqlParserSTRING_LITERAL_LONG2 = 79
	SparqlParserECHAR                = 80
	SparqlParserNIL                  = 81
	SparqlParserANON                 = 82
	SparqlParserPN_CHARS_U           = 83
	SparqlParserVARNAME              = 84
	SparqlParserPN_PREFIX            = 85
	SparqlParserPN_LOCAL             = 86
	SparqlParserWS                   = 87
)

// SparqlParser rules.
const (
	SparqlParserRULE_query                    = 0
	SparqlParserRULE_prologue                 = 1
	SparqlParserRULE_baseDecl                 = 2
	SparqlParserRULE_prefixDecl               = 3
	SparqlParserRULE_selectQuery              = 4
	SparqlParserRULE_constructQuery           = 5
	SparqlParserRULE_describeQuery            = 6
	SparqlParserRULE_askQuery                 = 7
	SparqlParserRULE_datasetClause            = 8
	SparqlParserRULE_defaultGraphClause       = 9
	SparqlParserRULE_namedGraphClause         = 10
	SparqlParserRULE_sourceSelector           = 11
	SparqlParserRULE_whereClause              = 12
	SparqlParserRULE_solutionModifier         = 13
	SparqlParserRULE_limitOffsetClauses       = 14
	SparqlParserRULE_orderClause              = 15
	SparqlParserRULE_orderCondition           = 16
	SparqlParserRULE_limitClause              = 17
	SparqlParserRULE_offsetClause             = 18
	SparqlParserRULE_groupGraphPattern        = 19
	SparqlParserRULE_triplesBlock             = 20
	SparqlParserRULE_graphPatternNotTriples   = 21
	SparqlParserRULE_optionalGraphPattern     = 22
	SparqlParserRULE_graphGraphPattern        = 23
	SparqlParserRULE_groupOrUnionGraphPattern = 24
	SparqlParserRULE_filter                   = 25
	SparqlParserRULE_constraint               = 26
	SparqlParserRULE_functionCall             = 27
	SparqlParserRULE_argList                  = 28
	SparqlParserRULE_constructTemplate        = 29
	SparqlParserRULE_constructTriples         = 30
	SparqlParserRULE_triplesSameSubject       = 31
	SparqlParserRULE_propertyListNotEmpty     = 32
	SparqlParserRULE_propertyList             = 33
	SparqlParserRULE_objectList               = 34
	SparqlParserRULE_object                   = 35
	SparqlParserRULE_verb                     = 36
	SparqlParserRULE_triplesNode              = 37
	SparqlParserRULE_blankNodePropertyList    = 38
	SparqlParserRULE_collection               = 39
	SparqlParserRULE_graphNode                = 40
	SparqlParserRULE_varOrTerm                = 41
	SparqlParserRULE_varOrIRIref              = 42
	SparqlParserRULE_varx                     = 43
	SparqlParserRULE_graphTerm                = 44
	SparqlParserRULE_expression               = 45
	SparqlParserRULE_conditionalOrExpression  = 46
	SparqlParserRULE_conditionalAndExpression = 47
	SparqlParserRULE_valueLogical             = 48
	SparqlParserRULE_relationalExpression     = 49
	SparqlParserRULE_numericExpression        = 50
	SparqlParserRULE_additiveExpression       = 51
	SparqlParserRULE_multiplicativeExpression = 52
	SparqlParserRULE_unaryExpression          = 53
	SparqlParserRULE_primaryExpression        = 54
	SparqlParserRULE_brackettedExpression     = 55
	SparqlParserRULE_builtInCall              = 56
	SparqlParserRULE_regexExpression          = 57
	SparqlParserRULE_iriRefOrFunction         = 58
	SparqlParserRULE_rdfLiteral               = 59
	SparqlParserRULE_numericLiteral           = 60
	SparqlParserRULE_numericLiteralUnsigned   = 61
	SparqlParserRULE_numericLiteralPositive   = 62
	SparqlParserRULE_numericLiteralNegative   = 63
	SparqlParserRULE_booleanLiteral           = 64
	SparqlParserRULE_rdfstring                = 65
	SparqlParserRULE_iriRef                   = 66
	SparqlParserRULE_prefixedName             = 67
	SparqlParserRULE_blankNode                = 68
)

// IQueryContext is an interface to support dynamic dispatch.
type IQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Prologue() IPrologueContext
	EOF() antlr.TerminalNode
	SelectQuery() ISelectQueryContext
	ConstructQuery() IConstructQueryContext
	DescribeQuery() IDescribeQueryContext
	AskQuery() IAskQueryContext

	// IsQueryContext differentiates from other interfaces.
	IsQueryContext()
}

type QueryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQueryContext() *QueryContext {
	var p = new(QueryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_query
	return p
}

func InitEmptyQueryContext(p *QueryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_query
}

func (*QueryContext) IsQueryContext() {}

func NewQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QueryContext {
	var p = new(QueryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_query

	return p
}

func (s *QueryContext) GetParser() antlr.Parser { return s.parser }

func (s *QueryContext) Prologue() IPrologueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrologueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrologueContext)
}

func (s *QueryContext) EOF() antlr.TerminalNode {
	return s.GetToken(SparqlParserEOF, 0)
}

func (s *QueryContext) SelectQuery() ISelectQueryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectQueryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectQueryContext)
}

func (s *QueryContext) ConstructQuery() IConstructQueryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstructQueryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstructQueryContext)
}

func (s *QueryContext) DescribeQuery() IDescribeQueryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDescribeQueryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDescribeQueryContext)
}

func (s *QueryContext) AskQuery() IAskQueryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAskQueryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAskQueryContext)
}

func (s *QueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QueryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.EnterQuery(s)
	}
}

func (s *QueryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.ExitQuery(s)
	}
}

func (p *SparqlParser) Query() (localctx IQueryContext) {
	localctx = NewQueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SparqlParserRULE_query)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(138)
		p.Prologue()
	}
	p.SetState(143)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SparqlParserT__2:
		{
			p.SetState(139)
			p.SelectQuery()
		}

	case SparqlParserT__6:
		{
			p.SetState(140)
			p.ConstructQuery()
		}

	case SparqlParserT__7:
		{
			p.SetState(141)
			p.DescribeQuery()
		}

	case SparqlParserT__8:
		{
			p.SetState(142)
			p.AskQuery()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(145)
		p.Match(SparqlParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrologueContext is an interface to support dynamic dispatch.
type IPrologueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BaseDecl() IBaseDeclContext
	AllPrefixDecl() []IPrefixDeclContext
	PrefixDecl(i int) IPrefixDeclContext

	// IsPrologueContext differentiates from other interfaces.
	IsPrologueContext()
}

type PrologueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrologueContext() *PrologueContext {
	var p = new(PrologueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_prologue
	return p
}

func InitEmptyPrologueContext(p *PrologueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_prologue
}

func (*PrologueContext) IsPrologueContext() {}

func NewPrologueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrologueContext {
	var p = new(PrologueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_prologue

	return p
}

func (s *PrologueContext) GetParser() antlr.Parser { return s.parser }

func (s *PrologueContext) BaseDecl() IBaseDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBaseDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBaseDeclContext)
}

func (s *PrologueContext) AllPrefixDecl() []IPrefixDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPrefixDeclContext); ok {
			len++
		}
	}

	tst := make([]IPrefixDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPrefixDeclContext); ok {
			tst[i] = t.(IPrefixDeclContext)
			i++
		}
	}

	return tst
}

func (s *PrologueContext) PrefixDecl(i int) IPrefixDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrefixDeclContext); ok {
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

	return t.(IPrefixDeclContext)
}

func (s *PrologueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrologueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrologueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.EnterPrologue(s)
	}
}

func (s *PrologueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.ExitPrologue(s)
	}
}

func (p *SparqlParser) Prologue() (localctx IPrologueContext) {
	localctx = NewPrologueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SparqlParserRULE_prologue)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(148)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SparqlParserT__0 {
		{
			p.SetState(147)
			p.BaseDecl()
		}

	}
	p.SetState(153)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SparqlParserT__1 {
		{
			p.SetState(150)
			p.PrefixDecl()
		}

		p.SetState(155)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBaseDeclContext is an interface to support dynamic dispatch.
type IBaseDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetIri returns the iri token.
	GetIri() antlr.Token

	// SetIri sets the iri token.
	SetIri(antlr.Token)

	// Getter signatures
	IRI_REF() antlr.TerminalNode

	// IsBaseDeclContext differentiates from other interfaces.
	IsBaseDeclContext()
}

type BaseDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	iri    antlr.Token
}

func NewEmptyBaseDeclContext() *BaseDeclContext {
	var p = new(BaseDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_baseDecl
	return p
}

func InitEmptyBaseDeclContext(p *BaseDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_baseDecl
}

func (*BaseDeclContext) IsBaseDeclContext() {}

func NewBaseDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BaseDeclContext {
	var p = new(BaseDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_baseDecl

	return p
}

func (s *BaseDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *BaseDeclContext) GetIri() antlr.Token { return s.iri }

func (s *BaseDeclContext) SetIri(v antlr.Token) { s.iri = v }

func (s *BaseDeclContext) IRI_REF() antlr.TerminalNode {
	return s.GetToken(SparqlParserIRI_REF, 0)
}

func (s *BaseDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BaseDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BaseDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.EnterBaseDecl(s)
	}
}

func (s *BaseDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.ExitBaseDecl(s)
	}
}

func (p *SparqlParser) BaseDecl() (localctx IBaseDeclContext) {
	localctx = NewBaseDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SparqlParserRULE_baseDecl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(156)
		p.Match(SparqlParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(157)

		var _m = p.Match(SparqlParserIRI_REF)

		localctx.(*BaseDeclContext).iri = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrefixDeclContext is an interface to support dynamic dispatch.
type IPrefixDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetPrefix returns the prefix token.
	GetPrefix() antlr.Token

	// GetIri returns the iri token.
	GetIri() antlr.Token

	// SetPrefix sets the prefix token.
	SetPrefix(antlr.Token)

	// SetIri sets the iri token.
	SetIri(antlr.Token)

	// Getter signatures
	PNAME_NS() antlr.TerminalNode
	IRI_REF() antlr.TerminalNode

	// IsPrefixDeclContext differentiates from other interfaces.
	IsPrefixDeclContext()
}

type PrefixDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	prefix antlr.Token
	iri    antlr.Token
}

func NewEmptyPrefixDeclContext() *PrefixDeclContext {
	var p = new(PrefixDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_prefixDecl
	return p
}

func InitEmptyPrefixDeclContext(p *PrefixDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_prefixDecl
}

func (*PrefixDeclContext) IsPrefixDeclContext() {}

func NewPrefixDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrefixDeclContext {
	var p = new(PrefixDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_prefixDecl

	return p
}

func (s *PrefixDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *PrefixDeclContext) GetPrefix() antlr.Token { return s.prefix }

func (s *PrefixDeclContext) GetIri() antlr.Token { return s.iri }

func (s *PrefixDeclContext) SetPrefix(v antlr.Token) { s.prefix = v }

func (s *PrefixDeclContext) SetIri(v antlr.Token) { s.iri = v }

func (s *PrefixDeclContext) PNAME_NS() antlr.TerminalNode {
	return s.GetToken(SparqlParserPNAME_NS, 0)
}

func (s *PrefixDeclContext) IRI_REF() antlr.TerminalNode {
	return s.GetToken(SparqlParserIRI_REF, 0)
}

func (s *PrefixDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrefixDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrefixDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.EnterPrefixDecl(s)
	}
}

func (s *PrefixDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.ExitPrefixDecl(s)
	}
}

func (p *SparqlParser) PrefixDecl() (localctx IPrefixDeclContext) {
	localctx = NewPrefixDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SparqlParserRULE_prefixDecl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(159)
		p.Match(SparqlParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(160)

		var _m = p.Match(SparqlParserPNAME_NS)

		localctx.(*PrefixDeclContext).prefix = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(161)

		var _m = p.Match(SparqlParserIRI_REF)

		localctx.(*PrefixDeclContext).iri = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISelectQueryContext is an interface to support dynamic dispatch.
type ISelectQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetMod returns the mod token.
	GetMod() antlr.Token

	// GetVarstar returns the varstar token.
	GetVarstar() antlr.Token

	// SetMod sets the mod token.
	SetMod(antlr.Token)

	// SetVarstar sets the varstar token.
	SetVarstar(antlr.Token)

	// Get_varx returns the _varx rule contexts.
	Get_varx() IVarxContext

	// Set_varx sets the _varx rule contexts.
	Set_varx(IVarxContext)

	// GetVars returns the vars rule context list.
	GetVars() []IVarxContext

	// SetVars sets the vars rule context list.
	SetVars([]IVarxContext)

	// Getter signatures
	WhereClause() IWhereClauseContext
	SolutionModifier() ISolutionModifierContext
	AllDatasetClause() []IDatasetClauseContext
	DatasetClause(i int) IDatasetClauseContext
	AllVarx() []IVarxContext
	Varx(i int) IVarxContext

	// IsSelectQueryContext differentiates from other interfaces.
	IsSelectQueryContext()
}

type SelectQueryContext struct {
	antlr.BaseParserRuleContext
	parser  antlr.Parser
	mod     antlr.Token
	_varx   IVarxContext
	vars    []IVarxContext
	varstar antlr.Token
}

func NewEmptySelectQueryContext() *SelectQueryContext {
	var p = new(SelectQueryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_selectQuery
	return p
}

func InitEmptySelectQueryContext(p *SelectQueryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_selectQuery
}

func (*SelectQueryContext) IsSelectQueryContext() {}

func NewSelectQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectQueryContext {
	var p = new(SelectQueryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_selectQuery

	return p
}

func (s *SelectQueryContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectQueryContext) GetMod() antlr.Token { return s.mod }

func (s *SelectQueryContext) GetVarstar() antlr.Token { return s.varstar }

func (s *SelectQueryContext) SetMod(v antlr.Token) { s.mod = v }

func (s *SelectQueryContext) SetVarstar(v antlr.Token) { s.varstar = v }

func (s *SelectQueryContext) Get_varx() IVarxContext { return s._varx }

func (s *SelectQueryContext) Set_varx(v IVarxContext) { s._varx = v }

func (s *SelectQueryContext) GetVars() []IVarxContext { return s.vars }

func (s *SelectQueryContext) SetVars(v []IVarxContext) { s.vars = v }

func (s *SelectQueryContext) WhereClause() IWhereClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhereClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhereClauseContext)
}

func (s *SelectQueryContext) SolutionModifier() ISolutionModifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISolutionModifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISolutionModifierContext)
}

func (s *SelectQueryContext) AllDatasetClause() []IDatasetClauseContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDatasetClauseContext); ok {
			len++
		}
	}

	tst := make([]IDatasetClauseContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDatasetClauseContext); ok {
			tst[i] = t.(IDatasetClauseContext)
			i++
		}
	}

	return tst
}

func (s *SelectQueryContext) DatasetClause(i int) IDatasetClauseContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDatasetClauseContext); ok {
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

	return t.(IDatasetClauseContext)
}

func (s *SelectQueryContext) AllVarx() []IVarxContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVarxContext); ok {
			len++
		}
	}

	tst := make([]IVarxContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVarxContext); ok {
			tst[i] = t.(IVarxContext)
			i++
		}
	}

	return tst
}

func (s *SelectQueryContext) Varx(i int) IVarxContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarxContext); ok {
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

	return t.(IVarxContext)
}

func (s *SelectQueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectQueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectQueryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.EnterSelectQuery(s)
	}
}

func (s *SelectQueryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.ExitSelectQuery(s)
	}
}

func (p *SparqlParser) SelectQuery() (localctx ISelectQueryContext) {
	localctx = NewSelectQueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SparqlParserRULE_selectQuery)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(163)
		p.Match(SparqlParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(165)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SparqlParserT__3 || _la == SparqlParserT__4 {
		{
			p.SetState(164)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*SelectQueryContext).mod = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == SparqlParserT__3 || _la == SparqlParserT__4) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*SelectQueryContext).mod = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	p.SetState(173)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SparqlParserVAR1, SparqlParserVAR2:
		p.SetState(168)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == SparqlParserVAR1 || _la == SparqlParserVAR2 {
			{
				p.SetState(167)

				var _x = p.Varx()

				localctx.(*SelectQueryContext)._varx = _x
			}
			localctx.(*SelectQueryContext).vars = append(localctx.(*SelectQueryContext).vars, localctx.(*SelectQueryContext)._varx)

			p.SetState(170)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case SparqlParserT__5:
		{
			p.SetState(172)

			var _m = p.Match(SparqlParserT__5)

			localctx.(*SelectQueryContext).varstar = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.SetState(178)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SparqlParserT__9 {
		{
			p.SetState(175)
			p.DatasetClause()
		}

		p.SetState(180)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(181)
		p.WhereClause()
	}
	{
		p.SetState(182)
		p.SolutionModifier()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConstructQueryContext is an interface to support dynamic dispatch.
type IConstructQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ConstructTemplate() IConstructTemplateContext
	WhereClause() IWhereClauseContext
	SolutionModifier() ISolutionModifierContext
	AllDatasetClause() []IDatasetClauseContext
	DatasetClause(i int) IDatasetClauseContext

	// IsConstructQueryContext differentiates from other interfaces.
	IsConstructQueryContext()
}

type ConstructQueryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstructQueryContext() *ConstructQueryContext {
	var p = new(ConstructQueryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_constructQuery
	return p
}

func InitEmptyConstructQueryContext(p *ConstructQueryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_constructQuery
}

func (*ConstructQueryContext) IsConstructQueryContext() {}

func NewConstructQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstructQueryContext {
	var p = new(ConstructQueryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_constructQuery

	return p
}

func (s *ConstructQueryContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstructQueryContext) ConstructTemplate() IConstructTemplateContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstructTemplateContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstructTemplateContext)
}

func (s *ConstructQueryContext) WhereClause() IWhereClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhereClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhereClauseContext)
}

func (s *ConstructQueryContext) SolutionModifier() ISolutionModifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISolutionModifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISolutionModifierContext)
}

func (s *ConstructQueryContext) AllDatasetClause() []IDatasetClauseContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDatasetClauseContext); ok {
			len++
		}
	}

	tst := make([]IDatasetClauseContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDatasetClauseContext); ok {
			tst[i] = t.(IDatasetClauseContext)
			i++
		}
	}

	return tst
}

func (s *ConstructQueryContext) DatasetClause(i int) IDatasetClauseContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDatasetClauseContext); ok {
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

	return t.(IDatasetClauseContext)
}

func (s *ConstructQueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstructQueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstructQueryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.EnterConstructQuery(s)
	}
}

func (s *ConstructQueryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.ExitConstructQuery(s)
	}
}

func (p *SparqlParser) ConstructQuery() (localctx IConstructQueryContext) {
	localctx = NewConstructQueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SparqlParserRULE_constructQuery)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(184)
		p.Match(SparqlParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(185)
		p.ConstructTemplate()
	}
	p.SetState(189)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SparqlParserT__9 {
		{
			p.SetState(186)
			p.DatasetClause()
		}

		p.SetState(191)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(192)
		p.WhereClause()
	}
	{
		p.SetState(193)
		p.SolutionModifier()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDescribeQueryContext is an interface to support dynamic dispatch.
type IDescribeQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SolutionModifier() ISolutionModifierContext
	AllDatasetClause() []IDatasetClauseContext
	DatasetClause(i int) IDatasetClauseContext
	WhereClause() IWhereClauseContext
	AllVarOrIRIref() []IVarOrIRIrefContext
	VarOrIRIref(i int) IVarOrIRIrefContext

	// IsDescribeQueryContext differentiates from other interfaces.
	IsDescribeQueryContext()
}

type DescribeQueryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDescribeQueryContext() *DescribeQueryContext {
	var p = new(DescribeQueryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_describeQuery
	return p
}

func InitEmptyDescribeQueryContext(p *DescribeQueryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_describeQuery
}

func (*DescribeQueryContext) IsDescribeQueryContext() {}

func NewDescribeQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DescribeQueryContext {
	var p = new(DescribeQueryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_describeQuery

	return p
}

func (s *DescribeQueryContext) GetParser() antlr.Parser { return s.parser }

func (s *DescribeQueryContext) SolutionModifier() ISolutionModifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISolutionModifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISolutionModifierContext)
}

func (s *DescribeQueryContext) AllDatasetClause() []IDatasetClauseContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDatasetClauseContext); ok {
			len++
		}
	}

	tst := make([]IDatasetClauseContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDatasetClauseContext); ok {
			tst[i] = t.(IDatasetClauseContext)
			i++
		}
	}

	return tst
}

func (s *DescribeQueryContext) DatasetClause(i int) IDatasetClauseContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDatasetClauseContext); ok {
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

	return t.(IDatasetClauseContext)
}

func (s *DescribeQueryContext) WhereClause() IWhereClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhereClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhereClauseContext)
}

func (s *DescribeQueryContext) AllVarOrIRIref() []IVarOrIRIrefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVarOrIRIrefContext); ok {
			len++
		}
	}

	tst := make([]IVarOrIRIrefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVarOrIRIrefContext); ok {
			tst[i] = t.(IVarOrIRIrefContext)
			i++
		}
	}

	return tst
}

func (s *DescribeQueryContext) VarOrIRIref(i int) IVarOrIRIrefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarOrIRIrefContext); ok {
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

	return t.(IVarOrIRIrefContext)
}

func (s *DescribeQueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DescribeQueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DescribeQueryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.EnterDescribeQuery(s)
	}
}

func (s *DescribeQueryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.ExitDescribeQuery(s)
	}
}

func (p *SparqlParser) DescribeQuery() (localctx IDescribeQueryContext) {
	localctx = NewDescribeQueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SparqlParserRULE_describeQuery)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(195)
		p.Match(SparqlParserT__7)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(202)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SparqlParserIRI_REF, SparqlParserPNAME_NS, SparqlParserPNAME_LN, SparqlParserVAR1, SparqlParserVAR2:
		p.SetState(197)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = ((int64((_la-59)) & ^0x3f) == 0 && ((int64(1)<<(_la-59))&55) != 0) {
			{
				p.SetState(196)
				p.VarOrIRIref()
			}

			p.SetState(199)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case SparqlParserT__5:
		{
			p.SetState(201)
			p.Match(SparqlParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.SetState(207)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SparqlParserT__9 {
		{
			p.SetState(204)
			p.DatasetClause()
		}

		p.SetState(209)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(211)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SparqlParserT__11 || _la == SparqlParserT__18 {
		{
			p.SetState(210)
			p.WhereClause()
		}

	}
	{
		p.SetState(213)
		p.SolutionModifier()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAskQueryContext is an interface to support dynamic dispatch.
type IAskQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WhereClause() IWhereClauseContext
	AllDatasetClause() []IDatasetClauseContext
	DatasetClause(i int) IDatasetClauseContext

	// IsAskQueryContext differentiates from other interfaces.
	IsAskQueryContext()
}

type AskQueryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAskQueryContext() *AskQueryContext {
	var p = new(AskQueryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_askQuery
	return p
}

func InitEmptyAskQueryContext(p *AskQueryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_askQuery
}

func (*AskQueryContext) IsAskQueryContext() {}

func NewAskQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AskQueryContext {
	var p = new(AskQueryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_askQuery

	return p
}

func (s *AskQueryContext) GetParser() antlr.Parser { return s.parser }

func (s *AskQueryContext) WhereClause() IWhereClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhereClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhereClauseContext)
}

func (s *AskQueryContext) AllDatasetClause() []IDatasetClauseContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDatasetClauseContext); ok {
			len++
		}
	}

	tst := make([]IDatasetClauseContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDatasetClauseContext); ok {
			tst[i] = t.(IDatasetClauseContext)
			i++
		}
	}

	return tst
}

func (s *AskQueryContext) DatasetClause(i int) IDatasetClauseContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDatasetClauseContext); ok {
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

	return t.(IDatasetClauseContext)
}

func (s *AskQueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AskQueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AskQueryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.EnterAskQuery(s)
	}
}

func (s *AskQueryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.ExitAskQuery(s)
	}
}

func (p *SparqlParser) AskQuery() (localctx IAskQueryContext) {
	localctx = NewAskQueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SparqlParserRULE_askQuery)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(215)
		p.Match(SparqlParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(219)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SparqlParserT__9 {
		{
			p.SetState(216)
			p.DatasetClause()
		}

		p.SetState(221)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(222)
		p.WhereClause()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDatasetClauseContext is an interface to support dynamic dispatch.
type IDatasetClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DefaultGraphClause() IDefaultGraphClauseContext
	NamedGraphClause() INamedGraphClauseContext

	// IsDatasetClauseContext differentiates from other interfaces.
	IsDatasetClauseContext()
}

type DatasetClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDatasetClauseContext() *DatasetClauseContext {
	var p = new(DatasetClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_datasetClause
	return p
}

func InitEmptyDatasetClauseContext(p *DatasetClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_datasetClause
}

func (*DatasetClauseContext) IsDatasetClauseContext() {}

func NewDatasetClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DatasetClauseContext {
	var p = new(DatasetClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_datasetClause

	return p
}

func (s *DatasetClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *DatasetClauseContext) DefaultGraphClause() IDefaultGraphClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefaultGraphClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefaultGraphClauseContext)
}

func (s *DatasetClauseContext) NamedGraphClause() INamedGraphClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INamedGraphClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INamedGraphClauseContext)
}

func (s *DatasetClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DatasetClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DatasetClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.EnterDatasetClause(s)
	}
}

func (s *DatasetClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.ExitDatasetClause(s)
	}
}

func (p *SparqlParser) DatasetClause() (localctx IDatasetClauseContext) {
	localctx = NewDatasetClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SparqlParserRULE_datasetClause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(224)
		p.Match(SparqlParserT__9)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(227)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SparqlParserIRI_REF, SparqlParserPNAME_NS, SparqlParserPNAME_LN:
		{
			p.SetState(225)
			p.DefaultGraphClause()
		}

	case SparqlParserT__10:
		{
			p.SetState(226)
			p.NamedGraphClause()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDefaultGraphClauseContext is an interface to support dynamic dispatch.
type IDefaultGraphClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SourceSelector() ISourceSelectorContext

	// IsDefaultGraphClauseContext differentiates from other interfaces.
	IsDefaultGraphClauseContext()
}

type DefaultGraphClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefaultGraphClauseContext() *DefaultGraphClauseContext {
	var p = new(DefaultGraphClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_defaultGraphClause
	return p
}

func InitEmptyDefaultGraphClauseContext(p *DefaultGraphClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_defaultGraphClause
}

func (*DefaultGraphClauseContext) IsDefaultGraphClauseContext() {}

func NewDefaultGraphClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefaultGraphClauseContext {
	var p = new(DefaultGraphClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_defaultGraphClause

	return p
}

func (s *DefaultGraphClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *DefaultGraphClauseContext) SourceSelector() ISourceSelectorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISourceSelectorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISourceSelectorContext)
}

func (s *DefaultGraphClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefaultGraphClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefaultGraphClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.EnterDefaultGraphClause(s)
	}
}

func (s *DefaultGraphClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.ExitDefaultGraphClause(s)
	}
}

func (p *SparqlParser) DefaultGraphClause() (localctx IDefaultGraphClauseContext) {
	localctx = NewDefaultGraphClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, SparqlParserRULE_defaultGraphClause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(229)
		p.SourceSelector()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INamedGraphClauseContext is an interface to support dynamic dispatch.
type INamedGraphClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SourceSelector() ISourceSelectorContext

	// IsNamedGraphClauseContext differentiates from other interfaces.
	IsNamedGraphClauseContext()
}

type NamedGraphClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNamedGraphClauseContext() *NamedGraphClauseContext {
	var p = new(NamedGraphClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_namedGraphClause
	return p
}

func InitEmptyNamedGraphClauseContext(p *NamedGraphClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_namedGraphClause
}

func (*NamedGraphClauseContext) IsNamedGraphClauseContext() {}

func NewNamedGraphClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NamedGraphClauseContext {
	var p = new(NamedGraphClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_namedGraphClause

	return p
}

func (s *NamedGraphClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *NamedGraphClauseContext) SourceSelector() ISourceSelectorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISourceSelectorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISourceSelectorContext)
}

func (s *NamedGraphClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamedGraphClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NamedGraphClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.EnterNamedGraphClause(s)
	}
}

func (s *NamedGraphClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.ExitNamedGraphClause(s)
	}
}

func (p *SparqlParser) NamedGraphClause() (localctx INamedGraphClauseContext) {
	localctx = NewNamedGraphClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, SparqlParserRULE_namedGraphClause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(231)
		p.Match(SparqlParserT__10)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(232)
		p.SourceSelector()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISourceSelectorContext is an interface to support dynamic dispatch.
type ISourceSelectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IriRef() IIriRefContext

	// IsSourceSelectorContext differentiates from other interfaces.
	IsSourceSelectorContext()
}

type SourceSelectorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySourceSelectorContext() *SourceSelectorContext {
	var p = new(SourceSelectorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_sourceSelector
	return p
}

func InitEmptySourceSelectorContext(p *SourceSelectorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_sourceSelector
}

func (*SourceSelectorContext) IsSourceSelectorContext() {}

func NewSourceSelectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SourceSelectorContext {
	var p = new(SourceSelectorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_sourceSelector

	return p
}

func (s *SourceSelectorContext) GetParser() antlr.Parser { return s.parser }

func (s *SourceSelectorContext) IriRef() IIriRefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIriRefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIriRefContext)
}

func (s *SourceSelectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SourceSelectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SourceSelectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.EnterSourceSelector(s)
	}
}

func (s *SourceSelectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.ExitSourceSelector(s)
	}
}

func (p *SparqlParser) SourceSelector() (localctx ISourceSelectorContext) {
	localctx = NewSourceSelectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, SparqlParserRULE_sourceSelector)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(234)
		p.IriRef()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IWhereClauseContext is an interface to support dynamic dispatch.
type IWhereClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetGgp returns the ggp rule contexts.
	GetGgp() IGroupGraphPatternContext

	// SetGgp sets the ggp rule contexts.
	SetGgp(IGroupGraphPatternContext)

	// Getter signatures
	GroupGraphPattern() IGroupGraphPatternContext

	// IsWhereClauseContext differentiates from other interfaces.
	IsWhereClauseContext()
}

type WhereClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	ggp    IGroupGraphPatternContext
}

func NewEmptyWhereClauseContext() *WhereClauseContext {
	var p = new(WhereClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_whereClause
	return p
}

func InitEmptyWhereClauseContext(p *WhereClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_whereClause
}

func (*WhereClauseContext) IsWhereClauseContext() {}

func NewWhereClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhereClauseContext {
	var p = new(WhereClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_whereClause

	return p
}

func (s *WhereClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *WhereClauseContext) GetGgp() IGroupGraphPatternContext { return s.ggp }

func (s *WhereClauseContext) SetGgp(v IGroupGraphPatternContext) { s.ggp = v }

func (s *WhereClauseContext) GroupGraphPattern() IGroupGraphPatternContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGroupGraphPatternContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGroupGraphPatternContext)
}

func (s *WhereClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhereClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhereClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.EnterWhereClause(s)
	}
}

func (s *WhereClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.ExitWhereClause(s)
	}
}

func (p *SparqlParser) WhereClause() (localctx IWhereClauseContext) {
	localctx = NewWhereClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, SparqlParserRULE_whereClause)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(237)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SparqlParserT__11 {
		{
			p.SetState(236)
			p.Match(SparqlParserT__11)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(239)

		var _x = p.GroupGraphPattern()

		localctx.(*WhereClauseContext).ggp = _x
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISolutionModifierContext is an interface to support dynamic dispatch.
type ISolutionModifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OrderClause() IOrderClauseContext
	LimitOffsetClauses() ILimitOffsetClausesContext

	// IsSolutionModifierContext differentiates from other interfaces.
	IsSolutionModifierContext()
}

type SolutionModifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySolutionModifierContext() *SolutionModifierContext {
	var p = new(SolutionModifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_solutionModifier
	return p
}

func InitEmptySolutionModifierContext(p *SolutionModifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_solutionModifier
}

func (*SolutionModifierContext) IsSolutionModifierContext() {}

func NewSolutionModifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SolutionModifierContext {
	var p = new(SolutionModifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_solutionModifier

	return p
}

func (s *SolutionModifierContext) GetParser() antlr.Parser { return s.parser }

func (s *SolutionModifierContext) OrderClause() IOrderClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOrderClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOrderClauseContext)
}

func (s *SolutionModifierContext) LimitOffsetClauses() ILimitOffsetClausesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILimitOffsetClausesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILimitOffsetClausesContext)
}

func (s *SolutionModifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SolutionModifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SolutionModifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.EnterSolutionModifier(s)
	}
}

func (s *SolutionModifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.ExitSolutionModifier(s)
	}
}

func (p *SparqlParser) SolutionModifier() (localctx ISolutionModifierContext) {
	localctx = NewSolutionModifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, SparqlParserRULE_solutionModifier)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(242)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SparqlParserT__12 {
		{
			p.SetState(241)
			p.OrderClause()
		}

	}
	p.SetState(245)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SparqlParserT__16 || _la == SparqlParserT__17 {
		{
			p.SetState(244)
			p.LimitOffsetClauses()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILimitOffsetClausesContext is an interface to support dynamic dispatch.
type ILimitOffsetClausesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LimitClause() ILimitClauseContext
	OffsetClause() IOffsetClauseContext

	// IsLimitOffsetClausesContext differentiates from other interfaces.
	IsLimitOffsetClausesContext()
}

type LimitOffsetClausesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLimitOffsetClausesContext() *LimitOffsetClausesContext {
	var p = new(LimitOffsetClausesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_limitOffsetClauses
	return p
}

func InitEmptyLimitOffsetClausesContext(p *LimitOffsetClausesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_limitOffsetClauses
}

func (*LimitOffsetClausesContext) IsLimitOffsetClausesContext() {}

func NewLimitOffsetClausesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LimitOffsetClausesContext {
	var p = new(LimitOffsetClausesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_limitOffsetClauses

	return p
}

func (s *LimitOffsetClausesContext) GetParser() antlr.Parser { return s.parser }

func (s *LimitOffsetClausesContext) LimitClause() ILimitClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILimitClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILimitClauseContext)
}

func (s *LimitOffsetClausesContext) OffsetClause() IOffsetClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOffsetClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOffsetClauseContext)
}

func (s *LimitOffsetClausesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LimitOffsetClausesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LimitOffsetClausesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.EnterLimitOffsetClauses(s)
	}
}

func (s *LimitOffsetClausesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.ExitLimitOffsetClauses(s)
	}
}

func (p *SparqlParser) LimitOffsetClauses() (localctx ILimitOffsetClausesContext) {
	localctx = NewLimitOffsetClausesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, SparqlParserRULE_limitOffsetClauses)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(255)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SparqlParserT__16:
		{
			p.SetState(247)
			p.LimitClause()
		}
		p.SetState(249)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SparqlParserT__17 {
			{
				p.SetState(248)
				p.OffsetClause()
			}

		}

	case SparqlParserT__17:
		{
			p.SetState(251)
			p.OffsetClause()
		}
		p.SetState(253)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SparqlParserT__16 {
			{
				p.SetState(252)
				p.LimitClause()
			}

		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOrderClauseContext is an interface to support dynamic dispatch.
type IOrderClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllOrderCondition() []IOrderConditionContext
	OrderCondition(i int) IOrderConditionContext

	// IsOrderClauseContext differentiates from other interfaces.
	IsOrderClauseContext()
}

type OrderClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrderClauseContext() *OrderClauseContext {
	var p = new(OrderClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_orderClause
	return p
}

func InitEmptyOrderClauseContext(p *OrderClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_orderClause
}

func (*OrderClauseContext) IsOrderClauseContext() {}

func NewOrderClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrderClauseContext {
	var p = new(OrderClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_orderClause

	return p
}

func (s *OrderClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *OrderClauseContext) AllOrderCondition() []IOrderConditionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOrderConditionContext); ok {
			len++
		}
	}

	tst := make([]IOrderConditionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOrderConditionContext); ok {
			tst[i] = t.(IOrderConditionContext)
			i++
		}
	}

	return tst
}

func (s *OrderClauseContext) OrderCondition(i int) IOrderConditionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOrderConditionContext); ok {
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

	return t.(IOrderConditionContext)
}

func (s *OrderClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrderClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrderClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.EnterOrderClause(s)
	}
}

func (s *OrderClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.ExitOrderClause(s)
	}
}

func (p *SparqlParser) OrderClause() (localctx IOrderClauseContext) {
	localctx = NewOrderClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, SparqlParserRULE_orderClause)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(257)
		p.Match(SparqlParserT__12)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(258)
		p.Match(SparqlParserT__13)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(260)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64((_la-15)) & ^0x3f) == 0 && ((int64(1)<<(_la-15))&969768181958659) != 0) {
		{
			p.SetState(259)
			p.OrderCondition()
		}

		p.SetState(262)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOrderConditionContext is an interface to support dynamic dispatch.
type IOrderConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BrackettedExpression() IBrackettedExpressionContext
	Constraint() IConstraintContext
	Varx() IVarxContext

	// IsOrderConditionContext differentiates from other interfaces.
	IsOrderConditionContext()
}

type OrderConditionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrderConditionContext() *OrderConditionContext {
	var p = new(OrderConditionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_orderCondition
	return p
}

func InitEmptyOrderConditionContext(p *OrderConditionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_orderCondition
}

func (*OrderConditionContext) IsOrderConditionContext() {}

func NewOrderConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrderConditionContext {
	var p = new(OrderConditionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_orderCondition

	return p
}

func (s *OrderConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *OrderConditionContext) BrackettedExpression() IBrackettedExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBrackettedExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBrackettedExpressionContext)
}

func (s *OrderConditionContext) Constraint() IConstraintContext {
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

func (s *OrderConditionContext) Varx() IVarxContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarxContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarxContext)
}

func (s *OrderConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrderConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrderConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.EnterOrderCondition(s)
	}
}

func (s *OrderConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.ExitOrderCondition(s)
	}
}

func (p *SparqlParser) OrderCondition() (localctx IOrderConditionContext) {
	localctx = NewOrderConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, SparqlParserRULE_orderCondition)
	var _la int

	p.SetState(270)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SparqlParserT__14, SparqlParserT__15:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(264)
			_la = p.GetTokenStream().LA(1)

			if !(_la == SparqlParserT__14 || _la == SparqlParserT__15) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(265)
			p.BrackettedExpression()
		}

	case SparqlParserT__25, SparqlParserT__44, SparqlParserT__45, SparqlParserT__46, SparqlParserT__47, SparqlParserT__48, SparqlParserT__49, SparqlParserT__50, SparqlParserT__51, SparqlParserT__52, SparqlParserT__53, SparqlParserT__54, SparqlParserIRI_REF, SparqlParserPNAME_NS, SparqlParserPNAME_LN, SparqlParserVAR1, SparqlParserVAR2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(268)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case SparqlParserT__25, SparqlParserT__44, SparqlParserT__45, SparqlParserT__46, SparqlParserT__47, SparqlParserT__48, SparqlParserT__49, SparqlParserT__50, SparqlParserT__51, SparqlParserT__52, SparqlParserT__53, SparqlParserT__54, SparqlParserIRI_REF, SparqlParserPNAME_NS, SparqlParserPNAME_LN:
			{
				p.SetState(266)
				p.Constraint()
			}

		case SparqlParserVAR1, SparqlParserVAR2:
			{
				p.SetState(267)
				p.Varx()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILimitClauseContext is an interface to support dynamic dispatch.
type ILimitClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INTEGER() antlr.TerminalNode

	// IsLimitClauseContext differentiates from other interfaces.
	IsLimitClauseContext()
}

type LimitClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLimitClauseContext() *LimitClauseContext {
	var p = new(LimitClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_limitClause
	return p
}

func InitEmptyLimitClauseContext(p *LimitClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_limitClause
}

func (*LimitClauseContext) IsLimitClauseContext() {}

func NewLimitClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LimitClauseContext {
	var p = new(LimitClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_limitClause

	return p
}

func (s *LimitClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *LimitClauseContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(SparqlParserINTEGER, 0)
}

func (s *LimitClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LimitClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LimitClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.EnterLimitClause(s)
	}
}

func (s *LimitClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.ExitLimitClause(s)
	}
}

func (p *SparqlParser) LimitClause() (localctx ILimitClauseContext) {
	localctx = NewLimitClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, SparqlParserRULE_limitClause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(272)
		p.Match(SparqlParserT__16)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(273)
		p.Match(SparqlParserINTEGER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOffsetClauseContext is an interface to support dynamic dispatch.
type IOffsetClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INTEGER() antlr.TerminalNode

	// IsOffsetClauseContext differentiates from other interfaces.
	IsOffsetClauseContext()
}

type OffsetClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOffsetClauseContext() *OffsetClauseContext {
	var p = new(OffsetClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_offsetClause
	return p
}

func InitEmptyOffsetClauseContext(p *OffsetClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_offsetClause
}

func (*OffsetClauseContext) IsOffsetClauseContext() {}

func NewOffsetClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OffsetClauseContext {
	var p = new(OffsetClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_offsetClause

	return p
}

func (s *OffsetClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *OffsetClauseContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(SparqlParserINTEGER, 0)
}

func (s *OffsetClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OffsetClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OffsetClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.EnterOffsetClause(s)
	}
}

func (s *OffsetClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.ExitOffsetClause(s)
	}
}

func (p *SparqlParser) OffsetClause() (localctx IOffsetClauseContext) {
	localctx = NewOffsetClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, SparqlParserRULE_offsetClause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(275)
		p.Match(SparqlParserT__17)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(276)
		p.Match(SparqlParserINTEGER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGroupGraphPatternContext is an interface to support dynamic dispatch.
type IGroupGraphPatternContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_triplesBlock returns the _triplesBlock rule contexts.
	Get_triplesBlock() ITriplesBlockContext

	// Get_graphPatternNotTriples returns the _graphPatternNotTriples rule contexts.
	Get_graphPatternNotTriples() IGraphPatternNotTriplesContext

	// Get_filter returns the _filter rule contexts.
	Get_filter() IFilterContext

	// Set_triplesBlock sets the _triplesBlock rule contexts.
	Set_triplesBlock(ITriplesBlockContext)

	// Set_graphPatternNotTriples sets the _graphPatternNotTriples rule contexts.
	Set_graphPatternNotTriples(IGraphPatternNotTriplesContext)

	// Set_filter sets the _filter rule contexts.
	Set_filter(IFilterContext)

	// GetTb returns the tb rule context list.
	GetTb() []ITriplesBlockContext

	// GetGpnt returns the gpnt rule context list.
	GetGpnt() []IGraphPatternNotTriplesContext

	// GetFlt returns the flt rule context list.
	GetFlt() []IFilterContext

	// SetTb sets the tb rule context list.
	SetTb([]ITriplesBlockContext)

	// SetGpnt sets the gpnt rule context list.
	SetGpnt([]IGraphPatternNotTriplesContext)

	// SetFlt sets the flt rule context list.
	SetFlt([]IFilterContext)

	// Getter signatures
	AllTriplesBlock() []ITriplesBlockContext
	TriplesBlock(i int) ITriplesBlockContext
	AllGraphPatternNotTriples() []IGraphPatternNotTriplesContext
	GraphPatternNotTriples(i int) IGraphPatternNotTriplesContext
	AllFilter() []IFilterContext
	Filter(i int) IFilterContext

	// IsGroupGraphPatternContext differentiates from other interfaces.
	IsGroupGraphPatternContext()
}

type GroupGraphPatternContext struct {
	antlr.BaseParserRuleContext
	parser                  antlr.Parser
	_triplesBlock           ITriplesBlockContext
	tb                      []ITriplesBlockContext
	_graphPatternNotTriples IGraphPatternNotTriplesContext
	gpnt                    []IGraphPatternNotTriplesContext
	_filter                 IFilterContext
	flt                     []IFilterContext
}

func NewEmptyGroupGraphPatternContext() *GroupGraphPatternContext {
	var p = new(GroupGraphPatternContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_groupGraphPattern
	return p
}

func InitEmptyGroupGraphPatternContext(p *GroupGraphPatternContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_groupGraphPattern
}

func (*GroupGraphPatternContext) IsGroupGraphPatternContext() {}

func NewGroupGraphPatternContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GroupGraphPatternContext {
	var p = new(GroupGraphPatternContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_groupGraphPattern

	return p
}

func (s *GroupGraphPatternContext) GetParser() antlr.Parser { return s.parser }

func (s *GroupGraphPatternContext) Get_triplesBlock() ITriplesBlockContext { return s._triplesBlock }

func (s *GroupGraphPatternContext) Get_graphPatternNotTriples() IGraphPatternNotTriplesContext {
	return s._graphPatternNotTriples
}

func (s *GroupGraphPatternContext) Get_filter() IFilterContext { return s._filter }

func (s *GroupGraphPatternContext) Set_triplesBlock(v ITriplesBlockContext) { s._triplesBlock = v }

func (s *GroupGraphPatternContext) Set_graphPatternNotTriples(v IGraphPatternNotTriplesContext) {
	s._graphPatternNotTriples = v
}

func (s *GroupGraphPatternContext) Set_filter(v IFilterContext) { s._filter = v }

func (s *GroupGraphPatternContext) GetTb() []ITriplesBlockContext { return s.tb }

func (s *GroupGraphPatternContext) GetGpnt() []IGraphPatternNotTriplesContext { return s.gpnt }

func (s *GroupGraphPatternContext) GetFlt() []IFilterContext { return s.flt }

func (s *GroupGraphPatternContext) SetTb(v []ITriplesBlockContext) { s.tb = v }

func (s *GroupGraphPatternContext) SetGpnt(v []IGraphPatternNotTriplesContext) { s.gpnt = v }

func (s *GroupGraphPatternContext) SetFlt(v []IFilterContext) { s.flt = v }

func (s *GroupGraphPatternContext) AllTriplesBlock() []ITriplesBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITriplesBlockContext); ok {
			len++
		}
	}

	tst := make([]ITriplesBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITriplesBlockContext); ok {
			tst[i] = t.(ITriplesBlockContext)
			i++
		}
	}

	return tst
}

func (s *GroupGraphPatternContext) TriplesBlock(i int) ITriplesBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITriplesBlockContext); ok {
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

	return t.(ITriplesBlockContext)
}

func (s *GroupGraphPatternContext) AllGraphPatternNotTriples() []IGraphPatternNotTriplesContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IGraphPatternNotTriplesContext); ok {
			len++
		}
	}

	tst := make([]IGraphPatternNotTriplesContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IGraphPatternNotTriplesContext); ok {
			tst[i] = t.(IGraphPatternNotTriplesContext)
			i++
		}
	}

	return tst
}

func (s *GroupGraphPatternContext) GraphPatternNotTriples(i int) IGraphPatternNotTriplesContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGraphPatternNotTriplesContext); ok {
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

	return t.(IGraphPatternNotTriplesContext)
}

func (s *GroupGraphPatternContext) AllFilter() []IFilterContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFilterContext); ok {
			len++
		}
	}

	tst := make([]IFilterContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFilterContext); ok {
			tst[i] = t.(IFilterContext)
			i++
		}
	}

	return tst
}

func (s *GroupGraphPatternContext) Filter(i int) IFilterContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFilterContext); ok {
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

	return t.(IFilterContext)
}

func (s *GroupGraphPatternContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupGraphPatternContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GroupGraphPatternContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.EnterGroupGraphPattern(s)
	}
}

func (s *GroupGraphPatternContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.ExitGroupGraphPattern(s)
	}
}

func (p *SparqlParser) GroupGraphPattern() (localctx IGroupGraphPatternContext) {
	localctx = NewGroupGraphPatternContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, SparqlParserRULE_groupGraphPattern)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(278)
		p.Match(SparqlParserT__18)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(280)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-26)) & ^0x3f) == 0 && ((int64(1)<<(_la-26))&112026488827543585) != 0 {
		{
			p.SetState(279)

			var _x = p.TriplesBlock()

			localctx.(*GroupGraphPatternContext)._triplesBlock = _x
		}
		localctx.(*GroupGraphPatternContext).tb = append(localctx.(*GroupGraphPatternContext).tb, localctx.(*GroupGraphPatternContext)._triplesBlock)

	}
	p.SetState(294)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&46661632) != 0 {
		p.SetState(284)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case SparqlParserT__18, SparqlParserT__21, SparqlParserT__22:
			{
				p.SetState(282)

				var _x = p.GraphPatternNotTriples()

				localctx.(*GroupGraphPatternContext)._graphPatternNotTriples = _x
			}
			localctx.(*GroupGraphPatternContext).gpnt = append(localctx.(*GroupGraphPatternContext).gpnt, localctx.(*GroupGraphPatternContext)._graphPatternNotTriples)

		case SparqlParserT__24:
			{
				p.SetState(283)

				var _x = p.Filter()

				localctx.(*GroupGraphPatternContext)._filter = _x
			}
			localctx.(*GroupGraphPatternContext).flt = append(localctx.(*GroupGraphPatternContext).flt, localctx.(*GroupGraphPatternContext)._filter)

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}
		p.SetState(287)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SparqlParserT__19 {
			{
				p.SetState(286)
				p.Match(SparqlParserT__19)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(290)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64((_la-26)) & ^0x3f) == 0 && ((int64(1)<<(_la-26))&112026488827543585) != 0 {
			{
				p.SetState(289)

				var _x = p.TriplesBlock()

				localctx.(*GroupGraphPatternContext)._triplesBlock = _x
			}
			localctx.(*GroupGraphPatternContext).tb = append(localctx.(*GroupGraphPatternContext).tb, localctx.(*GroupGraphPatternContext)._triplesBlock)

		}

		p.SetState(296)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(297)
		p.Match(SparqlParserT__20)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITriplesBlockContext is an interface to support dynamic dispatch.
type ITriplesBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTss returns the tss rule contexts.
	GetTss() ITriplesSameSubjectContext

	// GetTb returns the tb rule contexts.
	GetTb() ITriplesBlockContext

	// SetTss sets the tss rule contexts.
	SetTss(ITriplesSameSubjectContext)

	// SetTb sets the tb rule contexts.
	SetTb(ITriplesBlockContext)

	// Getter signatures
	TriplesSameSubject() ITriplesSameSubjectContext
	TriplesBlock() ITriplesBlockContext

	// IsTriplesBlockContext differentiates from other interfaces.
	IsTriplesBlockContext()
}

type TriplesBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	tss    ITriplesSameSubjectContext
	tb     ITriplesBlockContext
}

func NewEmptyTriplesBlockContext() *TriplesBlockContext {
	var p = new(TriplesBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_triplesBlock
	return p
}

func InitEmptyTriplesBlockContext(p *TriplesBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_triplesBlock
}

func (*TriplesBlockContext) IsTriplesBlockContext() {}

func NewTriplesBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TriplesBlockContext {
	var p = new(TriplesBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_triplesBlock

	return p
}

func (s *TriplesBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *TriplesBlockContext) GetTss() ITriplesSameSubjectContext { return s.tss }

func (s *TriplesBlockContext) GetTb() ITriplesBlockContext { return s.tb }

func (s *TriplesBlockContext) SetTss(v ITriplesSameSubjectContext) { s.tss = v }

func (s *TriplesBlockContext) SetTb(v ITriplesBlockContext) { s.tb = v }

func (s *TriplesBlockContext) TriplesSameSubject() ITriplesSameSubjectContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITriplesSameSubjectContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITriplesSameSubjectContext)
}

func (s *TriplesBlockContext) TriplesBlock() ITriplesBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITriplesBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITriplesBlockContext)
}

func (s *TriplesBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TriplesBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TriplesBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.EnterTriplesBlock(s)
	}
}

func (s *TriplesBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.ExitTriplesBlock(s)
	}
}

func (p *SparqlParser) TriplesBlock() (localctx ITriplesBlockContext) {
	localctx = NewTriplesBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, SparqlParserRULE_triplesBlock)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(299)

		var _x = p.TriplesSameSubject()

		localctx.(*TriplesBlockContext).tss = _x
	}
	p.SetState(304)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SparqlParserT__19 {
		{
			p.SetState(300)
			p.Match(SparqlParserT__19)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(302)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64((_la-26)) & ^0x3f) == 0 && ((int64(1)<<(_la-26))&112026488827543585) != 0 {
			{
				p.SetState(301)

				var _x = p.TriplesBlock()

				localctx.(*TriplesBlockContext).tb = _x
			}

		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGraphPatternNotTriplesContext is an interface to support dynamic dispatch.
type IGraphPatternNotTriplesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOgp returns the ogp rule contexts.
	GetOgp() IOptionalGraphPatternContext

	// GetGup returns the gup rule contexts.
	GetGup() IGroupOrUnionGraphPatternContext

	// GetGgp returns the ggp rule contexts.
	GetGgp() IGraphGraphPatternContext

	// SetOgp sets the ogp rule contexts.
	SetOgp(IOptionalGraphPatternContext)

	// SetGup sets the gup rule contexts.
	SetGup(IGroupOrUnionGraphPatternContext)

	// SetGgp sets the ggp rule contexts.
	SetGgp(IGraphGraphPatternContext)

	// Getter signatures
	OptionalGraphPattern() IOptionalGraphPatternContext
	GroupOrUnionGraphPattern() IGroupOrUnionGraphPatternContext
	GraphGraphPattern() IGraphGraphPatternContext

	// IsGraphPatternNotTriplesContext differentiates from other interfaces.
	IsGraphPatternNotTriplesContext()
}

type GraphPatternNotTriplesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	ogp    IOptionalGraphPatternContext
	gup    IGroupOrUnionGraphPatternContext
	ggp    IGraphGraphPatternContext
}

func NewEmptyGraphPatternNotTriplesContext() *GraphPatternNotTriplesContext {
	var p = new(GraphPatternNotTriplesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_graphPatternNotTriples
	return p
}

func InitEmptyGraphPatternNotTriplesContext(p *GraphPatternNotTriplesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_graphPatternNotTriples
}

func (*GraphPatternNotTriplesContext) IsGraphPatternNotTriplesContext() {}

func NewGraphPatternNotTriplesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GraphPatternNotTriplesContext {
	var p = new(GraphPatternNotTriplesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_graphPatternNotTriples

	return p
}

func (s *GraphPatternNotTriplesContext) GetParser() antlr.Parser { return s.parser }

func (s *GraphPatternNotTriplesContext) GetOgp() IOptionalGraphPatternContext { return s.ogp }

func (s *GraphPatternNotTriplesContext) GetGup() IGroupOrUnionGraphPatternContext { return s.gup }

func (s *GraphPatternNotTriplesContext) GetGgp() IGraphGraphPatternContext { return s.ggp }

func (s *GraphPatternNotTriplesContext) SetOgp(v IOptionalGraphPatternContext) { s.ogp = v }

func (s *GraphPatternNotTriplesContext) SetGup(v IGroupOrUnionGraphPatternContext) { s.gup = v }

func (s *GraphPatternNotTriplesContext) SetGgp(v IGraphGraphPatternContext) { s.ggp = v }

func (s *GraphPatternNotTriplesContext) OptionalGraphPattern() IOptionalGraphPatternContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOptionalGraphPatternContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOptionalGraphPatternContext)
}

func (s *GraphPatternNotTriplesContext) GroupOrUnionGraphPattern() IGroupOrUnionGraphPatternContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGroupOrUnionGraphPatternContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGroupOrUnionGraphPatternContext)
}

func (s *GraphPatternNotTriplesContext) GraphGraphPattern() IGraphGraphPatternContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGraphGraphPatternContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGraphGraphPatternContext)
}

func (s *GraphPatternNotTriplesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GraphPatternNotTriplesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GraphPatternNotTriplesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.EnterGraphPatternNotTriples(s)
	}
}

func (s *GraphPatternNotTriplesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.ExitGraphPatternNotTriples(s)
	}
}

func (p *SparqlParser) GraphPatternNotTriples() (localctx IGraphPatternNotTriplesContext) {
	localctx = NewGraphPatternNotTriplesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, SparqlParserRULE_graphPatternNotTriples)
	p.SetState(309)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SparqlParserT__21:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(306)

			var _x = p.OptionalGraphPattern()

			localctx.(*GraphPatternNotTriplesContext).ogp = _x
		}

	case SparqlParserT__18:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(307)

			var _x = p.GroupOrUnionGraphPattern()

			localctx.(*GraphPatternNotTriplesContext).gup = _x
		}

	case SparqlParserT__22:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(308)

			var _x = p.GraphGraphPattern()

			localctx.(*GraphPatternNotTriplesContext).ggp = _x
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOptionalGraphPatternContext is an interface to support dynamic dispatch.
type IOptionalGraphPatternContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetGgp returns the ggp rule contexts.
	GetGgp() IGroupGraphPatternContext

	// SetGgp sets the ggp rule contexts.
	SetGgp(IGroupGraphPatternContext)

	// Getter signatures
	GroupGraphPattern() IGroupGraphPatternContext

	// IsOptionalGraphPatternContext differentiates from other interfaces.
	IsOptionalGraphPatternContext()
}

type OptionalGraphPatternContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	ggp    IGroupGraphPatternContext
}

func NewEmptyOptionalGraphPatternContext() *OptionalGraphPatternContext {
	var p = new(OptionalGraphPatternContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_optionalGraphPattern
	return p
}

func InitEmptyOptionalGraphPatternContext(p *OptionalGraphPatternContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_optionalGraphPattern
}

func (*OptionalGraphPatternContext) IsOptionalGraphPatternContext() {}

func NewOptionalGraphPatternContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionalGraphPatternContext {
	var p = new(OptionalGraphPatternContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_optionalGraphPattern

	return p
}

func (s *OptionalGraphPatternContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionalGraphPatternContext) GetGgp() IGroupGraphPatternContext { return s.ggp }

func (s *OptionalGraphPatternContext) SetGgp(v IGroupGraphPatternContext) { s.ggp = v }

func (s *OptionalGraphPatternContext) GroupGraphPattern() IGroupGraphPatternContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGroupGraphPatternContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGroupGraphPatternContext)
}

func (s *OptionalGraphPatternContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionalGraphPatternContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionalGraphPatternContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.EnterOptionalGraphPattern(s)
	}
}

func (s *OptionalGraphPatternContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.ExitOptionalGraphPattern(s)
	}
}

func (p *SparqlParser) OptionalGraphPattern() (localctx IOptionalGraphPatternContext) {
	localctx = NewOptionalGraphPatternContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, SparqlParserRULE_optionalGraphPattern)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(311)
		p.Match(SparqlParserT__21)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(312)

		var _x = p.GroupGraphPattern()

		localctx.(*OptionalGraphPatternContext).ggp = _x
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGraphGraphPatternContext is an interface to support dynamic dispatch.
type IGraphGraphPatternContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetVi returns the vi rule contexts.
	GetVi() IVarOrIRIrefContext

	// GetGgp returns the ggp rule contexts.
	GetGgp() IGroupGraphPatternContext

	// SetVi sets the vi rule contexts.
	SetVi(IVarOrIRIrefContext)

	// SetGgp sets the ggp rule contexts.
	SetGgp(IGroupGraphPatternContext)

	// Getter signatures
	VarOrIRIref() IVarOrIRIrefContext
	GroupGraphPattern() IGroupGraphPatternContext

	// IsGraphGraphPatternContext differentiates from other interfaces.
	IsGraphGraphPatternContext()
}

type GraphGraphPatternContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	vi     IVarOrIRIrefContext
	ggp    IGroupGraphPatternContext
}

func NewEmptyGraphGraphPatternContext() *GraphGraphPatternContext {
	var p = new(GraphGraphPatternContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_graphGraphPattern
	return p
}

func InitEmptyGraphGraphPatternContext(p *GraphGraphPatternContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_graphGraphPattern
}

func (*GraphGraphPatternContext) IsGraphGraphPatternContext() {}

func NewGraphGraphPatternContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GraphGraphPatternContext {
	var p = new(GraphGraphPatternContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_graphGraphPattern

	return p
}

func (s *GraphGraphPatternContext) GetParser() antlr.Parser { return s.parser }

func (s *GraphGraphPatternContext) GetVi() IVarOrIRIrefContext { return s.vi }

func (s *GraphGraphPatternContext) GetGgp() IGroupGraphPatternContext { return s.ggp }

func (s *GraphGraphPatternContext) SetVi(v IVarOrIRIrefContext) { s.vi = v }

func (s *GraphGraphPatternContext) SetGgp(v IGroupGraphPatternContext) { s.ggp = v }

func (s *GraphGraphPatternContext) VarOrIRIref() IVarOrIRIrefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarOrIRIrefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarOrIRIrefContext)
}

func (s *GraphGraphPatternContext) GroupGraphPattern() IGroupGraphPatternContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGroupGraphPatternContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGroupGraphPatternContext)
}

func (s *GraphGraphPatternContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GraphGraphPatternContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GraphGraphPatternContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.EnterGraphGraphPattern(s)
	}
}

func (s *GraphGraphPatternContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.ExitGraphGraphPattern(s)
	}
}

func (p *SparqlParser) GraphGraphPattern() (localctx IGraphGraphPatternContext) {
	localctx = NewGraphGraphPatternContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, SparqlParserRULE_graphGraphPattern)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(314)
		p.Match(SparqlParserT__22)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(315)

		var _x = p.VarOrIRIref()

		localctx.(*GraphGraphPatternContext).vi = _x
	}
	{
		p.SetState(316)

		var _x = p.GroupGraphPattern()

		localctx.(*GraphGraphPatternContext).ggp = _x
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGroupOrUnionGraphPatternContext is an interface to support dynamic dispatch.
type IGroupOrUnionGraphPatternContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_groupGraphPattern returns the _groupGraphPattern rule contexts.
	Get_groupGraphPattern() IGroupGraphPatternContext

	// Set_groupGraphPattern sets the _groupGraphPattern rule contexts.
	Set_groupGraphPattern(IGroupGraphPatternContext)

	// GetUnion returns the union rule context list.
	GetUnion() []IGroupGraphPatternContext

	// SetUnion sets the union rule context list.
	SetUnion([]IGroupGraphPatternContext)

	// Getter signatures
	AllGroupGraphPattern() []IGroupGraphPatternContext
	GroupGraphPattern(i int) IGroupGraphPatternContext

	// IsGroupOrUnionGraphPatternContext differentiates from other interfaces.
	IsGroupOrUnionGraphPatternContext()
}

type GroupOrUnionGraphPatternContext struct {
	antlr.BaseParserRuleContext
	parser             antlr.Parser
	_groupGraphPattern IGroupGraphPatternContext
	union              []IGroupGraphPatternContext
}

func NewEmptyGroupOrUnionGraphPatternContext() *GroupOrUnionGraphPatternContext {
	var p = new(GroupOrUnionGraphPatternContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_groupOrUnionGraphPattern
	return p
}

func InitEmptyGroupOrUnionGraphPatternContext(p *GroupOrUnionGraphPatternContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_groupOrUnionGraphPattern
}

func (*GroupOrUnionGraphPatternContext) IsGroupOrUnionGraphPatternContext() {}

func NewGroupOrUnionGraphPatternContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GroupOrUnionGraphPatternContext {
	var p = new(GroupOrUnionGraphPatternContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_groupOrUnionGraphPattern

	return p
}

func (s *GroupOrUnionGraphPatternContext) GetParser() antlr.Parser { return s.parser }

func (s *GroupOrUnionGraphPatternContext) Get_groupGraphPattern() IGroupGraphPatternContext {
	return s._groupGraphPattern
}

func (s *GroupOrUnionGraphPatternContext) Set_groupGraphPattern(v IGroupGraphPatternContext) {
	s._groupGraphPattern = v
}

func (s *GroupOrUnionGraphPatternContext) GetUnion() []IGroupGraphPatternContext { return s.union }

func (s *GroupOrUnionGraphPatternContext) SetUnion(v []IGroupGraphPatternContext) { s.union = v }

func (s *GroupOrUnionGraphPatternContext) AllGroupGraphPattern() []IGroupGraphPatternContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IGroupGraphPatternContext); ok {
			len++
		}
	}

	tst := make([]IGroupGraphPatternContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IGroupGraphPatternContext); ok {
			tst[i] = t.(IGroupGraphPatternContext)
			i++
		}
	}

	return tst
}

func (s *GroupOrUnionGraphPatternContext) GroupGraphPattern(i int) IGroupGraphPatternContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGroupGraphPatternContext); ok {
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

	return t.(IGroupGraphPatternContext)
}

func (s *GroupOrUnionGraphPatternContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupOrUnionGraphPatternContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GroupOrUnionGraphPatternContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.EnterGroupOrUnionGraphPattern(s)
	}
}

func (s *GroupOrUnionGraphPatternContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.ExitGroupOrUnionGraphPattern(s)
	}
}

func (p *SparqlParser) GroupOrUnionGraphPattern() (localctx IGroupOrUnionGraphPatternContext) {
	localctx = NewGroupOrUnionGraphPatternContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, SparqlParserRULE_groupOrUnionGraphPattern)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(318)

		var _x = p.GroupGraphPattern()

		localctx.(*GroupOrUnionGraphPatternContext)._groupGraphPattern = _x
	}
	localctx.(*GroupOrUnionGraphPatternContext).union = append(localctx.(*GroupOrUnionGraphPatternContext).union, localctx.(*GroupOrUnionGraphPatternContext)._groupGraphPattern)
	p.SetState(323)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SparqlParserT__23 {
		{
			p.SetState(319)
			p.Match(SparqlParserT__23)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(320)

			var _x = p.GroupGraphPattern()

			localctx.(*GroupOrUnionGraphPatternContext)._groupGraphPattern = _x
		}
		localctx.(*GroupOrUnionGraphPatternContext).union = append(localctx.(*GroupOrUnionGraphPatternContext).union, localctx.(*GroupOrUnionGraphPatternContext)._groupGraphPattern)

		p.SetState(325)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFilterContext is an interface to support dynamic dispatch.
type IFilterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Constraint() IConstraintContext

	// IsFilterContext differentiates from other interfaces.
	IsFilterContext()
}

type FilterContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFilterContext() *FilterContext {
	var p = new(FilterContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_filter
	return p
}

func InitEmptyFilterContext(p *FilterContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_filter
}

func (*FilterContext) IsFilterContext() {}

func NewFilterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FilterContext {
	var p = new(FilterContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_filter

	return p
}

func (s *FilterContext) GetParser() antlr.Parser { return s.parser }

func (s *FilterContext) Constraint() IConstraintContext {
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

func (s *FilterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FilterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FilterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.EnterFilter(s)
	}
}

func (s *FilterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.ExitFilter(s)
	}
}

func (p *SparqlParser) Filter() (localctx IFilterContext) {
	localctx = NewFilterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, SparqlParserRULE_filter)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(326)
		p.Match(SparqlParserT__24)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(327)
		p.Constraint()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConstraintContext is an interface to support dynamic dispatch.
type IConstraintContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BrackettedExpression() IBrackettedExpressionContext
	BuiltInCall() IBuiltInCallContext
	FunctionCall() IFunctionCallContext

	// IsConstraintContext differentiates from other interfaces.
	IsConstraintContext()
}

type ConstraintContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstraintContext() *ConstraintContext {
	var p = new(ConstraintContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_constraint
	return p
}

func InitEmptyConstraintContext(p *ConstraintContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_constraint
}

func (*ConstraintContext) IsConstraintContext() {}

func NewConstraintContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstraintContext {
	var p = new(ConstraintContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_constraint

	return p
}

func (s *ConstraintContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstraintContext) BrackettedExpression() IBrackettedExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBrackettedExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBrackettedExpressionContext)
}

func (s *ConstraintContext) BuiltInCall() IBuiltInCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBuiltInCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBuiltInCallContext)
}

func (s *ConstraintContext) FunctionCall() IFunctionCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *ConstraintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstraintContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstraintContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.EnterConstraint(s)
	}
}

func (s *ConstraintContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.ExitConstraint(s)
	}
}

func (p *SparqlParser) Constraint() (localctx IConstraintContext) {
	localctx = NewConstraintContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, SparqlParserRULE_constraint)
	p.SetState(332)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SparqlParserT__25:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(329)
			p.BrackettedExpression()
		}

	case SparqlParserT__44, SparqlParserT__45, SparqlParserT__46, SparqlParserT__47, SparqlParserT__48, SparqlParserT__49, SparqlParserT__50, SparqlParserT__51, SparqlParserT__52, SparqlParserT__53, SparqlParserT__54:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(330)
			p.BuiltInCall()
		}

	case SparqlParserIRI_REF, SparqlParserPNAME_NS, SparqlParserPNAME_LN:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(331)
			p.FunctionCall()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunctionCallContext is an interface to support dynamic dispatch.
type IFunctionCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IriRef() IIriRefContext
	ArgList() IArgListContext

	// IsFunctionCallContext differentiates from other interfaces.
	IsFunctionCallContext()
}

type FunctionCallContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionCallContext() *FunctionCallContext {
	var p = new(FunctionCallContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_functionCall
	return p
}

func InitEmptyFunctionCallContext(p *FunctionCallContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_functionCall
}

func (*FunctionCallContext) IsFunctionCallContext() {}

func NewFunctionCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionCallContext {
	var p = new(FunctionCallContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_functionCall

	return p
}

func (s *FunctionCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionCallContext) IriRef() IIriRefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIriRefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIriRefContext)
}

func (s *FunctionCallContext) ArgList() IArgListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgListContext)
}

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.EnterFunctionCall(s)
	}
}

func (s *FunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.ExitFunctionCall(s)
	}
}

func (p *SparqlParser) FunctionCall() (localctx IFunctionCallContext) {
	localctx = NewFunctionCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, SparqlParserRULE_functionCall)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(334)
		p.IriRef()
	}
	{
		p.SetState(335)
		p.ArgList()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArgListContext is an interface to support dynamic dispatch.
type IArgListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NIL() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext

	// IsArgListContext differentiates from other interfaces.
	IsArgListContext()
}

type ArgListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgListContext() *ArgListContext {
	var p = new(ArgListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_argList
	return p
}

func InitEmptyArgListContext(p *ArgListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_argList
}

func (*ArgListContext) IsArgListContext() {}

func NewArgListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgListContext {
	var p = new(ArgListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_argList

	return p
}

func (s *ArgListContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgListContext) NIL() antlr.TerminalNode {
	return s.GetToken(SparqlParserNIL, 0)
}

func (s *ArgListContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ArgListContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *ArgListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.EnterArgList(s)
	}
}

func (s *ArgListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.ExitArgList(s)
	}
}

func (p *SparqlParser) ArgList() (localctx IArgListContext) {
	localctx = NewArgListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, SparqlParserRULE_argList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(349)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SparqlParserNIL:
		{
			p.SetState(337)
			p.Match(SparqlParserNIL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SparqlParserT__25:
		{
			p.SetState(338)
			p.Match(SparqlParserT__25)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(339)
			p.Expression()
		}
		p.SetState(344)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == SparqlParserT__26 {
			{
				p.SetState(340)
				p.Match(SparqlParserT__26)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(341)
				p.Expression()
			}

			p.SetState(346)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(347)
			p.Match(SparqlParserT__27)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConstructTemplateContext is an interface to support dynamic dispatch.
type IConstructTemplateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ConstructTriples() IConstructTriplesContext

	// IsConstructTemplateContext differentiates from other interfaces.
	IsConstructTemplateContext()
}

type ConstructTemplateContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstructTemplateContext() *ConstructTemplateContext {
	var p = new(ConstructTemplateContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_constructTemplate
	return p
}

func InitEmptyConstructTemplateContext(p *ConstructTemplateContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_constructTemplate
}

func (*ConstructTemplateContext) IsConstructTemplateContext() {}

func NewConstructTemplateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstructTemplateContext {
	var p = new(ConstructTemplateContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_constructTemplate

	return p
}

func (s *ConstructTemplateContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstructTemplateContext) ConstructTriples() IConstructTriplesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstructTriplesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstructTriplesContext)
}

func (s *ConstructTemplateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstructTemplateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstructTemplateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.EnterConstructTemplate(s)
	}
}

func (s *ConstructTemplateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.ExitConstructTemplate(s)
	}
}

func (p *SparqlParser) ConstructTemplate() (localctx IConstructTemplateContext) {
	localctx = NewConstructTemplateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, SparqlParserRULE_constructTemplate)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(351)
		p.Match(SparqlParserT__18)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(353)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-26)) & ^0x3f) == 0 && ((int64(1)<<(_la-26))&112026488827543585) != 0 {
		{
			p.SetState(352)
			p.ConstructTriples()
		}

	}
	{
		p.SetState(355)
		p.Match(SparqlParserT__20)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConstructTriplesContext is an interface to support dynamic dispatch.
type IConstructTriplesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TriplesSameSubject() ITriplesSameSubjectContext
	ConstructTriples() IConstructTriplesContext

	// IsConstructTriplesContext differentiates from other interfaces.
	IsConstructTriplesContext()
}

type ConstructTriplesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstructTriplesContext() *ConstructTriplesContext {
	var p = new(ConstructTriplesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_constructTriples
	return p
}

func InitEmptyConstructTriplesContext(p *ConstructTriplesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_constructTriples
}

func (*ConstructTriplesContext) IsConstructTriplesContext() {}

func NewConstructTriplesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstructTriplesContext {
	var p = new(ConstructTriplesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_constructTriples

	return p
}

func (s *ConstructTriplesContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstructTriplesContext) TriplesSameSubject() ITriplesSameSubjectContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITriplesSameSubjectContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITriplesSameSubjectContext)
}

func (s *ConstructTriplesContext) ConstructTriples() IConstructTriplesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstructTriplesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstructTriplesContext)
}

func (s *ConstructTriplesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstructTriplesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstructTriplesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.EnterConstructTriples(s)
	}
}

func (s *ConstructTriplesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.ExitConstructTriples(s)
	}
}

func (p *SparqlParser) ConstructTriples() (localctx IConstructTriplesContext) {
	localctx = NewConstructTriplesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, SparqlParserRULE_constructTriples)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(357)
		p.TriplesSameSubject()
	}
	p.SetState(362)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SparqlParserT__19 {
		{
			p.SetState(358)
			p.Match(SparqlParserT__19)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(360)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64((_la-26)) & ^0x3f) == 0 && ((int64(1)<<(_la-26))&112026488827543585) != 0 {
			{
				p.SetState(359)
				p.ConstructTriples()
			}

		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITriplesSameSubjectContext is an interface to support dynamic dispatch.
type ITriplesSameSubjectContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetSubject returns the subject rule contexts.
	GetSubject() IVarOrTermContext

	// GetProperties returns the properties rule contexts.
	GetProperties() IPropertyListNotEmptyContext

	// SetSubject sets the subject rule contexts.
	SetSubject(IVarOrTermContext)

	// SetProperties sets the properties rule contexts.
	SetProperties(IPropertyListNotEmptyContext)

	// Getter signatures
	VarOrTerm() IVarOrTermContext
	PropertyListNotEmpty() IPropertyListNotEmptyContext
	TriplesNode() ITriplesNodeContext
	PropertyList() IPropertyListContext

	// IsTriplesSameSubjectContext differentiates from other interfaces.
	IsTriplesSameSubjectContext()
}

type TriplesSameSubjectContext struct {
	antlr.BaseParserRuleContext
	parser     antlr.Parser
	subject    IVarOrTermContext
	properties IPropertyListNotEmptyContext
}

func NewEmptyTriplesSameSubjectContext() *TriplesSameSubjectContext {
	var p = new(TriplesSameSubjectContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_triplesSameSubject
	return p
}

func InitEmptyTriplesSameSubjectContext(p *TriplesSameSubjectContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_triplesSameSubject
}

func (*TriplesSameSubjectContext) IsTriplesSameSubjectContext() {}

func NewTriplesSameSubjectContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TriplesSameSubjectContext {
	var p = new(TriplesSameSubjectContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_triplesSameSubject

	return p
}

func (s *TriplesSameSubjectContext) GetParser() antlr.Parser { return s.parser }

func (s *TriplesSameSubjectContext) GetSubject() IVarOrTermContext { return s.subject }

func (s *TriplesSameSubjectContext) GetProperties() IPropertyListNotEmptyContext { return s.properties }

func (s *TriplesSameSubjectContext) SetSubject(v IVarOrTermContext) { s.subject = v }

func (s *TriplesSameSubjectContext) SetProperties(v IPropertyListNotEmptyContext) { s.properties = v }

func (s *TriplesSameSubjectContext) VarOrTerm() IVarOrTermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarOrTermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarOrTermContext)
}

func (s *TriplesSameSubjectContext) PropertyListNotEmpty() IPropertyListNotEmptyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPropertyListNotEmptyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPropertyListNotEmptyContext)
}

func (s *TriplesSameSubjectContext) TriplesNode() ITriplesNodeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITriplesNodeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITriplesNodeContext)
}

func (s *TriplesSameSubjectContext) PropertyList() IPropertyListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPropertyListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPropertyListContext)
}

func (s *TriplesSameSubjectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TriplesSameSubjectContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TriplesSameSubjectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.EnterTriplesSameSubject(s)
	}
}

func (s *TriplesSameSubjectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.ExitTriplesSameSubject(s)
	}
}

func (p *SparqlParser) TriplesSameSubject() (localctx ITriplesSameSubjectContext) {
	localctx = NewTriplesSameSubjectContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, SparqlParserRULE_triplesSameSubject)
	p.SetState(370)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SparqlParserT__56, SparqlParserT__57, SparqlParserIRI_REF, SparqlParserPNAME_NS, SparqlParserPNAME_LN, SparqlParserBLANK_NODE_LABEL, SparqlParserVAR1, SparqlParserVAR2, SparqlParserINTEGER, SparqlParserDECIMAL, SparqlParserDOUBLE, SparqlParserINTEGER_POSITIVE, SparqlParserDECIMAL_POSITIVE, SparqlParserDOUBLE_POSITIVE, SparqlParserINTEGER_NEGATIVE, SparqlParserDECIMAL_NEGATIVE, SparqlParserDOUBLE_NEGATIVE, SparqlParserSTRING_LITERAL1, SparqlParserSTRING_LITERAL2, SparqlParserNIL, SparqlParserANON:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(364)

			var _x = p.VarOrTerm()

			localctx.(*TriplesSameSubjectContext).subject = _x
		}
		{
			p.SetState(365)

			var _x = p.PropertyListNotEmpty()

			localctx.(*TriplesSameSubjectContext).properties = _x
		}

	case SparqlParserT__25, SparqlParserT__30:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(367)
			p.TriplesNode()
		}
		{
			p.SetState(368)
			p.PropertyList()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPropertyListNotEmptyContext is an interface to support dynamic dispatch.
type IPropertyListNotEmptyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_verb returns the _verb rule contexts.
	Get_verb() IVerbContext

	// Get_objectList returns the _objectList rule contexts.
	Get_objectList() IObjectListContext

	// Set_verb sets the _verb rule contexts.
	Set_verb(IVerbContext)

	// Set_objectList sets the _objectList rule contexts.
	Set_objectList(IObjectListContext)

	// GetVerbs returns the verbs rule context list.
	GetVerbs() []IVerbContext

	// GetOl returns the ol rule context list.
	GetOl() []IObjectListContext

	// SetVerbs sets the verbs rule context list.
	SetVerbs([]IVerbContext)

	// SetOl sets the ol rule context list.
	SetOl([]IObjectListContext)

	// Getter signatures
	AllVerb() []IVerbContext
	Verb(i int) IVerbContext
	AllObjectList() []IObjectListContext
	ObjectList(i int) IObjectListContext

	// IsPropertyListNotEmptyContext differentiates from other interfaces.
	IsPropertyListNotEmptyContext()
}

type PropertyListNotEmptyContext struct {
	antlr.BaseParserRuleContext
	parser      antlr.Parser
	_verb       IVerbContext
	verbs       []IVerbContext
	_objectList IObjectListContext
	ol          []IObjectListContext
}

func NewEmptyPropertyListNotEmptyContext() *PropertyListNotEmptyContext {
	var p = new(PropertyListNotEmptyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_propertyListNotEmpty
	return p
}

func InitEmptyPropertyListNotEmptyContext(p *PropertyListNotEmptyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_propertyListNotEmpty
}

func (*PropertyListNotEmptyContext) IsPropertyListNotEmptyContext() {}

func NewPropertyListNotEmptyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropertyListNotEmptyContext {
	var p = new(PropertyListNotEmptyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_propertyListNotEmpty

	return p
}

func (s *PropertyListNotEmptyContext) GetParser() antlr.Parser { return s.parser }

func (s *PropertyListNotEmptyContext) Get_verb() IVerbContext { return s._verb }

func (s *PropertyListNotEmptyContext) Get_objectList() IObjectListContext { return s._objectList }

func (s *PropertyListNotEmptyContext) Set_verb(v IVerbContext) { s._verb = v }

func (s *PropertyListNotEmptyContext) Set_objectList(v IObjectListContext) { s._objectList = v }

func (s *PropertyListNotEmptyContext) GetVerbs() []IVerbContext { return s.verbs }

func (s *PropertyListNotEmptyContext) GetOl() []IObjectListContext { return s.ol }

func (s *PropertyListNotEmptyContext) SetVerbs(v []IVerbContext) { s.verbs = v }

func (s *PropertyListNotEmptyContext) SetOl(v []IObjectListContext) { s.ol = v }

func (s *PropertyListNotEmptyContext) AllVerb() []IVerbContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVerbContext); ok {
			len++
		}
	}

	tst := make([]IVerbContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVerbContext); ok {
			tst[i] = t.(IVerbContext)
			i++
		}
	}

	return tst
}

func (s *PropertyListNotEmptyContext) Verb(i int) IVerbContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVerbContext); ok {
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

	return t.(IVerbContext)
}

func (s *PropertyListNotEmptyContext) AllObjectList() []IObjectListContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IObjectListContext); ok {
			len++
		}
	}

	tst := make([]IObjectListContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IObjectListContext); ok {
			tst[i] = t.(IObjectListContext)
			i++
		}
	}

	return tst
}

func (s *PropertyListNotEmptyContext) ObjectList(i int) IObjectListContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObjectListContext); ok {
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

	return t.(IObjectListContext)
}

func (s *PropertyListNotEmptyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyListNotEmptyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PropertyListNotEmptyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.EnterPropertyListNotEmpty(s)
	}
}

func (s *PropertyListNotEmptyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.ExitPropertyListNotEmpty(s)
	}
}

func (p *SparqlParser) PropertyListNotEmpty() (localctx IPropertyListNotEmptyContext) {
	localctx = NewPropertyListNotEmptyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, SparqlParserRULE_propertyListNotEmpty)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(372)

		var _x = p.Verb()

		localctx.(*PropertyListNotEmptyContext)._verb = _x
	}
	localctx.(*PropertyListNotEmptyContext).verbs = append(localctx.(*PropertyListNotEmptyContext).verbs, localctx.(*PropertyListNotEmptyContext)._verb)
	{
		p.SetState(373)

		var _x = p.ObjectList()

		localctx.(*PropertyListNotEmptyContext)._objectList = _x
	}
	localctx.(*PropertyListNotEmptyContext).ol = append(localctx.(*PropertyListNotEmptyContext).ol, localctx.(*PropertyListNotEmptyContext)._objectList)
	p.SetState(382)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SparqlParserT__28 {
		{
			p.SetState(374)
			p.Match(SparqlParserT__28)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(378)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64((_la-30)) & ^0x3f) == 0 && ((int64(1)<<(_la-30))&29527900161) != 0 {
			{
				p.SetState(375)

				var _x = p.Verb()

				localctx.(*PropertyListNotEmptyContext)._verb = _x
			}
			localctx.(*PropertyListNotEmptyContext).verbs = append(localctx.(*PropertyListNotEmptyContext).verbs, localctx.(*PropertyListNotEmptyContext)._verb)
			{
				p.SetState(376)

				var _x = p.ObjectList()

				localctx.(*PropertyListNotEmptyContext)._objectList = _x
			}
			localctx.(*PropertyListNotEmptyContext).ol = append(localctx.(*PropertyListNotEmptyContext).ol, localctx.(*PropertyListNotEmptyContext)._objectList)

		}

		p.SetState(384)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPropertyListContext is an interface to support dynamic dispatch.
type IPropertyListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PropertyListNotEmpty() IPropertyListNotEmptyContext

	// IsPropertyListContext differentiates from other interfaces.
	IsPropertyListContext()
}

type PropertyListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPropertyListContext() *PropertyListContext {
	var p = new(PropertyListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_propertyList
	return p
}

func InitEmptyPropertyListContext(p *PropertyListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_propertyList
}

func (*PropertyListContext) IsPropertyListContext() {}

func NewPropertyListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropertyListContext {
	var p = new(PropertyListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_propertyList

	return p
}

func (s *PropertyListContext) GetParser() antlr.Parser { return s.parser }

func (s *PropertyListContext) PropertyListNotEmpty() IPropertyListNotEmptyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPropertyListNotEmptyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPropertyListNotEmptyContext)
}

func (s *PropertyListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PropertyListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.EnterPropertyList(s)
	}
}

func (s *PropertyListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.ExitPropertyList(s)
	}
}

func (p *SparqlParser) PropertyList() (localctx IPropertyListContext) {
	localctx = NewPropertyListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, SparqlParserRULE_propertyList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(386)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-30)) & ^0x3f) == 0 && ((int64(1)<<(_la-30))&29527900161) != 0 {
		{
			p.SetState(385)
			p.PropertyListNotEmpty()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IObjectListContext is an interface to support dynamic dispatch.
type IObjectListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_object returns the _object rule contexts.
	Get_object() IObjectContext

	// Set_object sets the _object rule contexts.
	Set_object(IObjectContext)

	// GetOb returns the ob rule context list.
	GetOb() []IObjectContext

	// SetOb sets the ob rule context list.
	SetOb([]IObjectContext)

	// Getter signatures
	AllObject() []IObjectContext
	Object(i int) IObjectContext

	// IsObjectListContext differentiates from other interfaces.
	IsObjectListContext()
}

type ObjectListContext struct {
	antlr.BaseParserRuleContext
	parser  antlr.Parser
	_object IObjectContext
	ob      []IObjectContext
}

func NewEmptyObjectListContext() *ObjectListContext {
	var p = new(ObjectListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_objectList
	return p
}

func InitEmptyObjectListContext(p *ObjectListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_objectList
}

func (*ObjectListContext) IsObjectListContext() {}

func NewObjectListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectListContext {
	var p = new(ObjectListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_objectList

	return p
}

func (s *ObjectListContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectListContext) Get_object() IObjectContext { return s._object }

func (s *ObjectListContext) Set_object(v IObjectContext) { s._object = v }

func (s *ObjectListContext) GetOb() []IObjectContext { return s.ob }

func (s *ObjectListContext) SetOb(v []IObjectContext) { s.ob = v }

func (s *ObjectListContext) AllObject() []IObjectContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IObjectContext); ok {
			len++
		}
	}

	tst := make([]IObjectContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IObjectContext); ok {
			tst[i] = t.(IObjectContext)
			i++
		}
	}

	return tst
}

func (s *ObjectListContext) Object(i int) IObjectContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObjectContext); ok {
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

	return t.(IObjectContext)
}

func (s *ObjectListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.EnterObjectList(s)
	}
}

func (s *ObjectListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.ExitObjectList(s)
	}
}

func (p *SparqlParser) ObjectList() (localctx IObjectListContext) {
	localctx = NewObjectListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, SparqlParserRULE_objectList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(388)

		var _x = p.Object()

		localctx.(*ObjectListContext)._object = _x
	}
	localctx.(*ObjectListContext).ob = append(localctx.(*ObjectListContext).ob, localctx.(*ObjectListContext)._object)
	p.SetState(393)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SparqlParserT__26 {
		{
			p.SetState(389)
			p.Match(SparqlParserT__26)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(390)

			var _x = p.Object()

			localctx.(*ObjectListContext)._object = _x
		}
		localctx.(*ObjectListContext).ob = append(localctx.(*ObjectListContext).ob, localctx.(*ObjectListContext)._object)

		p.SetState(395)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IObjectContext is an interface to support dynamic dispatch.
type IObjectContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetGn returns the gn rule contexts.
	GetGn() IGraphNodeContext

	// SetGn sets the gn rule contexts.
	SetGn(IGraphNodeContext)

	// Getter signatures
	GraphNode() IGraphNodeContext

	// IsObjectContext differentiates from other interfaces.
	IsObjectContext()
}

type ObjectContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	gn     IGraphNodeContext
}

func NewEmptyObjectContext() *ObjectContext {
	var p = new(ObjectContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_object
	return p
}

func InitEmptyObjectContext(p *ObjectContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_object
}

func (*ObjectContext) IsObjectContext() {}

func NewObjectContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectContext {
	var p = new(ObjectContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_object

	return p
}

func (s *ObjectContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectContext) GetGn() IGraphNodeContext { return s.gn }

func (s *ObjectContext) SetGn(v IGraphNodeContext) { s.gn = v }

func (s *ObjectContext) GraphNode() IGraphNodeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGraphNodeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGraphNodeContext)
}

func (s *ObjectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.EnterObject(s)
	}
}

func (s *ObjectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.ExitObject(s)
	}
}

func (p *SparqlParser) Object() (localctx IObjectContext) {
	localctx = NewObjectContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, SparqlParserRULE_object)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(396)

		var _x = p.GraphNode()

		localctx.(*ObjectContext).gn = _x
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVerbContext is an interface to support dynamic dispatch.
type IVerbContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetA returns the a token.
	GetA() antlr.Token

	// SetA sets the a token.
	SetA(antlr.Token)

	// GetVi returns the vi rule contexts.
	GetVi() IVarOrIRIrefContext

	// SetVi sets the vi rule contexts.
	SetVi(IVarOrIRIrefContext)

	// Getter signatures
	VarOrIRIref() IVarOrIRIrefContext

	// IsVerbContext differentiates from other interfaces.
	IsVerbContext()
}

type VerbContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	vi     IVarOrIRIrefContext
	a      antlr.Token
}

func NewEmptyVerbContext() *VerbContext {
	var p = new(VerbContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_verb
	return p
}

func InitEmptyVerbContext(p *VerbContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_verb
}

func (*VerbContext) IsVerbContext() {}

func NewVerbContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VerbContext {
	var p = new(VerbContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_verb

	return p
}

func (s *VerbContext) GetParser() antlr.Parser { return s.parser }

func (s *VerbContext) GetA() antlr.Token { return s.a }

func (s *VerbContext) SetA(v antlr.Token) { s.a = v }

func (s *VerbContext) GetVi() IVarOrIRIrefContext { return s.vi }

func (s *VerbContext) SetVi(v IVarOrIRIrefContext) { s.vi = v }

func (s *VerbContext) VarOrIRIref() IVarOrIRIrefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarOrIRIrefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarOrIRIrefContext)
}

func (s *VerbContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VerbContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VerbContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.EnterVerb(s)
	}
}

func (s *VerbContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.ExitVerb(s)
	}
}

func (p *SparqlParser) Verb() (localctx IVerbContext) {
	localctx = NewVerbContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, SparqlParserRULE_verb)
	p.SetState(400)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SparqlParserIRI_REF, SparqlParserPNAME_NS, SparqlParserPNAME_LN, SparqlParserVAR1, SparqlParserVAR2:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(398)

			var _x = p.VarOrIRIref()

			localctx.(*VerbContext).vi = _x
		}

	case SparqlParserT__29:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(399)

			var _m = p.Match(SparqlParserT__29)

			localctx.(*VerbContext).a = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITriplesNodeContext is an interface to support dynamic dispatch.
type ITriplesNodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Collection() ICollectionContext
	BlankNodePropertyList() IBlankNodePropertyListContext

	// IsTriplesNodeContext differentiates from other interfaces.
	IsTriplesNodeContext()
}

type TriplesNodeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTriplesNodeContext() *TriplesNodeContext {
	var p = new(TriplesNodeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_triplesNode
	return p
}

func InitEmptyTriplesNodeContext(p *TriplesNodeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_triplesNode
}

func (*TriplesNodeContext) IsTriplesNodeContext() {}

func NewTriplesNodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TriplesNodeContext {
	var p = new(TriplesNodeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_triplesNode

	return p
}

func (s *TriplesNodeContext) GetParser() antlr.Parser { return s.parser }

func (s *TriplesNodeContext) Collection() ICollectionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICollectionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICollectionContext)
}

func (s *TriplesNodeContext) BlankNodePropertyList() IBlankNodePropertyListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlankNodePropertyListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlankNodePropertyListContext)
}

func (s *TriplesNodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TriplesNodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TriplesNodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.EnterTriplesNode(s)
	}
}

func (s *TriplesNodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.ExitTriplesNode(s)
	}
}

func (p *SparqlParser) TriplesNode() (localctx ITriplesNodeContext) {
	localctx = NewTriplesNodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, SparqlParserRULE_triplesNode)
	p.SetState(404)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SparqlParserT__25:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(402)
			p.Collection()
		}

	case SparqlParserT__30:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(403)
			p.BlankNodePropertyList()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlankNodePropertyListContext is an interface to support dynamic dispatch.
type IBlankNodePropertyListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PropertyListNotEmpty() IPropertyListNotEmptyContext

	// IsBlankNodePropertyListContext differentiates from other interfaces.
	IsBlankNodePropertyListContext()
}

type BlankNodePropertyListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlankNodePropertyListContext() *BlankNodePropertyListContext {
	var p = new(BlankNodePropertyListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_blankNodePropertyList
	return p
}

func InitEmptyBlankNodePropertyListContext(p *BlankNodePropertyListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_blankNodePropertyList
}

func (*BlankNodePropertyListContext) IsBlankNodePropertyListContext() {}

func NewBlankNodePropertyListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlankNodePropertyListContext {
	var p = new(BlankNodePropertyListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_blankNodePropertyList

	return p
}

func (s *BlankNodePropertyListContext) GetParser() antlr.Parser { return s.parser }

func (s *BlankNodePropertyListContext) PropertyListNotEmpty() IPropertyListNotEmptyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPropertyListNotEmptyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPropertyListNotEmptyContext)
}

func (s *BlankNodePropertyListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlankNodePropertyListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlankNodePropertyListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.EnterBlankNodePropertyList(s)
	}
}

func (s *BlankNodePropertyListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.ExitBlankNodePropertyList(s)
	}
}

func (p *SparqlParser) BlankNodePropertyList() (localctx IBlankNodePropertyListContext) {
	localctx = NewBlankNodePropertyListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, SparqlParserRULE_blankNodePropertyList)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(406)
		p.Match(SparqlParserT__30)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(407)
		p.PropertyListNotEmpty()
	}
	{
		p.SetState(408)
		p.Match(SparqlParserT__31)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICollectionContext is an interface to support dynamic dispatch.
type ICollectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllGraphNode() []IGraphNodeContext
	GraphNode(i int) IGraphNodeContext

	// IsCollectionContext differentiates from other interfaces.
	IsCollectionContext()
}

type CollectionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCollectionContext() *CollectionContext {
	var p = new(CollectionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_collection
	return p
}

func InitEmptyCollectionContext(p *CollectionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_collection
}

func (*CollectionContext) IsCollectionContext() {}

func NewCollectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CollectionContext {
	var p = new(CollectionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_collection

	return p
}

func (s *CollectionContext) GetParser() antlr.Parser { return s.parser }

func (s *CollectionContext) AllGraphNode() []IGraphNodeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IGraphNodeContext); ok {
			len++
		}
	}

	tst := make([]IGraphNodeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IGraphNodeContext); ok {
			tst[i] = t.(IGraphNodeContext)
			i++
		}
	}

	return tst
}

func (s *CollectionContext) GraphNode(i int) IGraphNodeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGraphNodeContext); ok {
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

	return t.(IGraphNodeContext)
}

func (s *CollectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CollectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CollectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.EnterCollection(s)
	}
}

func (s *CollectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.ExitCollection(s)
	}
}

func (p *SparqlParser) Collection() (localctx ICollectionContext) {
	localctx = NewCollectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, SparqlParserRULE_collection)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(410)
		p.Match(SparqlParserT__25)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(412)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64((_la-26)) & ^0x3f) == 0 && ((int64(1)<<(_la-26))&112026488827543585) != 0) {
		{
			p.SetState(411)
			p.GraphNode()
		}

		p.SetState(414)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(416)
		p.Match(SparqlParserT__27)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGraphNodeContext is an interface to support dynamic dispatch.
type IGraphNodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetVt returns the vt rule contexts.
	GetVt() IVarOrTermContext

	// GetTn returns the tn rule contexts.
	GetTn() ITriplesNodeContext

	// SetVt sets the vt rule contexts.
	SetVt(IVarOrTermContext)

	// SetTn sets the tn rule contexts.
	SetTn(ITriplesNodeContext)

	// Getter signatures
	VarOrTerm() IVarOrTermContext
	TriplesNode() ITriplesNodeContext

	// IsGraphNodeContext differentiates from other interfaces.
	IsGraphNodeContext()
}

type GraphNodeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	vt     IVarOrTermContext
	tn     ITriplesNodeContext
}

func NewEmptyGraphNodeContext() *GraphNodeContext {
	var p = new(GraphNodeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_graphNode
	return p
}

func InitEmptyGraphNodeContext(p *GraphNodeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_graphNode
}

func (*GraphNodeContext) IsGraphNodeContext() {}

func NewGraphNodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GraphNodeContext {
	var p = new(GraphNodeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_graphNode

	return p
}

func (s *GraphNodeContext) GetParser() antlr.Parser { return s.parser }

func (s *GraphNodeContext) GetVt() IVarOrTermContext { return s.vt }

func (s *GraphNodeContext) GetTn() ITriplesNodeContext { return s.tn }

func (s *GraphNodeContext) SetVt(v IVarOrTermContext) { s.vt = v }

func (s *GraphNodeContext) SetTn(v ITriplesNodeContext) { s.tn = v }

func (s *GraphNodeContext) VarOrTerm() IVarOrTermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarOrTermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarOrTermContext)
}

func (s *GraphNodeContext) TriplesNode() ITriplesNodeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITriplesNodeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITriplesNodeContext)
}

func (s *GraphNodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GraphNodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GraphNodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.EnterGraphNode(s)
	}
}

func (s *GraphNodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.ExitGraphNode(s)
	}
}

func (p *SparqlParser) GraphNode() (localctx IGraphNodeContext) {
	localctx = NewGraphNodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, SparqlParserRULE_graphNode)
	p.SetState(420)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SparqlParserT__56, SparqlParserT__57, SparqlParserIRI_REF, SparqlParserPNAME_NS, SparqlParserPNAME_LN, SparqlParserBLANK_NODE_LABEL, SparqlParserVAR1, SparqlParserVAR2, SparqlParserINTEGER, SparqlParserDECIMAL, SparqlParserDOUBLE, SparqlParserINTEGER_POSITIVE, SparqlParserDECIMAL_POSITIVE, SparqlParserDOUBLE_POSITIVE, SparqlParserINTEGER_NEGATIVE, SparqlParserDECIMAL_NEGATIVE, SparqlParserDOUBLE_NEGATIVE, SparqlParserSTRING_LITERAL1, SparqlParserSTRING_LITERAL2, SparqlParserNIL, SparqlParserANON:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(418)

			var _x = p.VarOrTerm()

			localctx.(*GraphNodeContext).vt = _x
		}

	case SparqlParserT__25, SparqlParserT__30:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(419)

			var _x = p.TriplesNode()

			localctx.(*GraphNodeContext).tn = _x
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVarOrTermContext is an interface to support dynamic dispatch.
type IVarOrTermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetVariable returns the variable rule contexts.
	GetVariable() IVarxContext

	// GetGt returns the gt rule contexts.
	GetGt() IGraphTermContext

	// SetVariable sets the variable rule contexts.
	SetVariable(IVarxContext)

	// SetGt sets the gt rule contexts.
	SetGt(IGraphTermContext)

	// Getter signatures
	Varx() IVarxContext
	GraphTerm() IGraphTermContext

	// IsVarOrTermContext differentiates from other interfaces.
	IsVarOrTermContext()
}

type VarOrTermContext struct {
	antlr.BaseParserRuleContext
	parser   antlr.Parser
	variable IVarxContext
	gt       IGraphTermContext
}

func NewEmptyVarOrTermContext() *VarOrTermContext {
	var p = new(VarOrTermContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_varOrTerm
	return p
}

func InitEmptyVarOrTermContext(p *VarOrTermContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_varOrTerm
}

func (*VarOrTermContext) IsVarOrTermContext() {}

func NewVarOrTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarOrTermContext {
	var p = new(VarOrTermContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_varOrTerm

	return p
}

func (s *VarOrTermContext) GetParser() antlr.Parser { return s.parser }

func (s *VarOrTermContext) GetVariable() IVarxContext { return s.variable }

func (s *VarOrTermContext) GetGt() IGraphTermContext { return s.gt }

func (s *VarOrTermContext) SetVariable(v IVarxContext) { s.variable = v }

func (s *VarOrTermContext) SetGt(v IGraphTermContext) { s.gt = v }

func (s *VarOrTermContext) Varx() IVarxContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarxContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarxContext)
}

func (s *VarOrTermContext) GraphTerm() IGraphTermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGraphTermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGraphTermContext)
}

func (s *VarOrTermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarOrTermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarOrTermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.EnterVarOrTerm(s)
	}
}

func (s *VarOrTermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.ExitVarOrTerm(s)
	}
}

func (p *SparqlParser) VarOrTerm() (localctx IVarOrTermContext) {
	localctx = NewVarOrTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, SparqlParserRULE_varOrTerm)
	p.SetState(424)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SparqlParserVAR1, SparqlParserVAR2:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(422)

			var _x = p.Varx()

			localctx.(*VarOrTermContext).variable = _x
		}

	case SparqlParserT__56, SparqlParserT__57, SparqlParserIRI_REF, SparqlParserPNAME_NS, SparqlParserPNAME_LN, SparqlParserBLANK_NODE_LABEL, SparqlParserINTEGER, SparqlParserDECIMAL, SparqlParserDOUBLE, SparqlParserINTEGER_POSITIVE, SparqlParserDECIMAL_POSITIVE, SparqlParserDOUBLE_POSITIVE, SparqlParserINTEGER_NEGATIVE, SparqlParserDECIMAL_NEGATIVE, SparqlParserDOUBLE_NEGATIVE, SparqlParserSTRING_LITERAL1, SparqlParserSTRING_LITERAL2, SparqlParserNIL, SparqlParserANON:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(423)

			var _x = p.GraphTerm()

			localctx.(*VarOrTermContext).gt = _x
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVarOrIRIrefContext is an interface to support dynamic dispatch.
type IVarOrIRIrefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetVariable returns the variable rule contexts.
	GetVariable() IVarxContext

	// GetIri returns the iri rule contexts.
	GetIri() IIriRefContext

	// SetVariable sets the variable rule contexts.
	SetVariable(IVarxContext)

	// SetIri sets the iri rule contexts.
	SetIri(IIriRefContext)

	// Getter signatures
	Varx() IVarxContext
	IriRef() IIriRefContext

	// IsVarOrIRIrefContext differentiates from other interfaces.
	IsVarOrIRIrefContext()
}

type VarOrIRIrefContext struct {
	antlr.BaseParserRuleContext
	parser   antlr.Parser
	variable IVarxContext
	iri      IIriRefContext
}

func NewEmptyVarOrIRIrefContext() *VarOrIRIrefContext {
	var p = new(VarOrIRIrefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_varOrIRIref
	return p
}

func InitEmptyVarOrIRIrefContext(p *VarOrIRIrefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_varOrIRIref
}

func (*VarOrIRIrefContext) IsVarOrIRIrefContext() {}

func NewVarOrIRIrefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarOrIRIrefContext {
	var p = new(VarOrIRIrefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_varOrIRIref

	return p
}

func (s *VarOrIRIrefContext) GetParser() antlr.Parser { return s.parser }

func (s *VarOrIRIrefContext) GetVariable() IVarxContext { return s.variable }

func (s *VarOrIRIrefContext) GetIri() IIriRefContext { return s.iri }

func (s *VarOrIRIrefContext) SetVariable(v IVarxContext) { s.variable = v }

func (s *VarOrIRIrefContext) SetIri(v IIriRefContext) { s.iri = v }

func (s *VarOrIRIrefContext) Varx() IVarxContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarxContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarxContext)
}

func (s *VarOrIRIrefContext) IriRef() IIriRefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIriRefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIriRefContext)
}

func (s *VarOrIRIrefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarOrIRIrefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarOrIRIrefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.EnterVarOrIRIref(s)
	}
}

func (s *VarOrIRIrefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.ExitVarOrIRIref(s)
	}
}

func (p *SparqlParser) VarOrIRIref() (localctx IVarOrIRIrefContext) {
	localctx = NewVarOrIRIrefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, SparqlParserRULE_varOrIRIref)
	p.SetState(428)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SparqlParserVAR1, SparqlParserVAR2:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(426)

			var _x = p.Varx()

			localctx.(*VarOrIRIrefContext).variable = _x
		}

	case SparqlParserIRI_REF, SparqlParserPNAME_NS, SparqlParserPNAME_LN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(427)

			var _x = p.IriRef()

			localctx.(*VarOrIRIrefContext).iri = _x
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVarxContext is an interface to support dynamic dispatch.
type IVarxContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VAR1() antlr.TerminalNode
	VAR2() antlr.TerminalNode

	// IsVarxContext differentiates from other interfaces.
	IsVarxContext()
}

type VarxContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarxContext() *VarxContext {
	var p = new(VarxContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_varx
	return p
}

func InitEmptyVarxContext(p *VarxContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_varx
}

func (*VarxContext) IsVarxContext() {}

func NewVarxContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarxContext {
	var p = new(VarxContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_varx

	return p
}

func (s *VarxContext) GetParser() antlr.Parser { return s.parser }

func (s *VarxContext) VAR1() antlr.TerminalNode {
	return s.GetToken(SparqlParserVAR1, 0)
}

func (s *VarxContext) VAR2() antlr.TerminalNode {
	return s.GetToken(SparqlParserVAR2, 0)
}

func (s *VarxContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarxContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarxContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.EnterVarx(s)
	}
}

func (s *VarxContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.ExitVarx(s)
	}
}

func (p *SparqlParser) Varx() (localctx IVarxContext) {
	localctx = NewVarxContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, SparqlParserRULE_varx)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(430)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SparqlParserVAR1 || _la == SparqlParserVAR2) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGraphTermContext is an interface to support dynamic dispatch.
type IGraphTermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetEmp returns the emp token.
	GetEmp() antlr.Token

	// SetEmp sets the emp token.
	SetEmp(antlr.Token)

	// GetIri returns the iri rule contexts.
	GetIri() IIriRefContext

	// GetLit returns the lit rule contexts.
	GetLit() IRdfLiteralContext

	// GetNum returns the num rule contexts.
	GetNum() INumericLiteralContext

	// GetBol returns the bol rule contexts.
	GetBol() IBooleanLiteralContext

	// GetBlk returns the blk rule contexts.
	GetBlk() IBlankNodeContext

	// SetIri sets the iri rule contexts.
	SetIri(IIriRefContext)

	// SetLit sets the lit rule contexts.
	SetLit(IRdfLiteralContext)

	// SetNum sets the num rule contexts.
	SetNum(INumericLiteralContext)

	// SetBol sets the bol rule contexts.
	SetBol(IBooleanLiteralContext)

	// SetBlk sets the blk rule contexts.
	SetBlk(IBlankNodeContext)

	// Getter signatures
	IriRef() IIriRefContext
	RdfLiteral() IRdfLiteralContext
	NumericLiteral() INumericLiteralContext
	BooleanLiteral() IBooleanLiteralContext
	BlankNode() IBlankNodeContext
	NIL() antlr.TerminalNode

	// IsGraphTermContext differentiates from other interfaces.
	IsGraphTermContext()
}

type GraphTermContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	iri    IIriRefContext
	lit    IRdfLiteralContext
	num    INumericLiteralContext
	bol    IBooleanLiteralContext
	blk    IBlankNodeContext
	emp    antlr.Token
}

func NewEmptyGraphTermContext() *GraphTermContext {
	var p = new(GraphTermContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_graphTerm
	return p
}

func InitEmptyGraphTermContext(p *GraphTermContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_graphTerm
}

func (*GraphTermContext) IsGraphTermContext() {}

func NewGraphTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GraphTermContext {
	var p = new(GraphTermContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_graphTerm

	return p
}

func (s *GraphTermContext) GetParser() antlr.Parser { return s.parser }

func (s *GraphTermContext) GetEmp() antlr.Token { return s.emp }

func (s *GraphTermContext) SetEmp(v antlr.Token) { s.emp = v }

func (s *GraphTermContext) GetIri() IIriRefContext { return s.iri }

func (s *GraphTermContext) GetLit() IRdfLiteralContext { return s.lit }

func (s *GraphTermContext) GetNum() INumericLiteralContext { return s.num }

func (s *GraphTermContext) GetBol() IBooleanLiteralContext { return s.bol }

func (s *GraphTermContext) GetBlk() IBlankNodeContext { return s.blk }

func (s *GraphTermContext) SetIri(v IIriRefContext) { s.iri = v }

func (s *GraphTermContext) SetLit(v IRdfLiteralContext) { s.lit = v }

func (s *GraphTermContext) SetNum(v INumericLiteralContext) { s.num = v }

func (s *GraphTermContext) SetBol(v IBooleanLiteralContext) { s.bol = v }

func (s *GraphTermContext) SetBlk(v IBlankNodeContext) { s.blk = v }

func (s *GraphTermContext) IriRef() IIriRefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIriRefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIriRefContext)
}

func (s *GraphTermContext) RdfLiteral() IRdfLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRdfLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRdfLiteralContext)
}

func (s *GraphTermContext) NumericLiteral() INumericLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumericLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumericLiteralContext)
}

func (s *GraphTermContext) BooleanLiteral() IBooleanLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBooleanLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBooleanLiteralContext)
}

func (s *GraphTermContext) BlankNode() IBlankNodeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlankNodeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlankNodeContext)
}

func (s *GraphTermContext) NIL() antlr.TerminalNode {
	return s.GetToken(SparqlParserNIL, 0)
}

func (s *GraphTermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GraphTermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GraphTermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.EnterGraphTerm(s)
	}
}

func (s *GraphTermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.ExitGraphTerm(s)
	}
}

func (p *SparqlParser) GraphTerm() (localctx IGraphTermContext) {
	localctx = NewGraphTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, SparqlParserRULE_graphTerm)
	p.SetState(438)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SparqlParserIRI_REF, SparqlParserPNAME_NS, SparqlParserPNAME_LN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(432)

			var _x = p.IriRef()

			localctx.(*GraphTermContext).iri = _x
		}

	case SparqlParserSTRING_LITERAL1, SparqlParserSTRING_LITERAL2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(433)

			var _x = p.RdfLiteral()

			localctx.(*GraphTermContext).lit = _x
		}

	case SparqlParserINTEGER, SparqlParserDECIMAL, SparqlParserDOUBLE, SparqlParserINTEGER_POSITIVE, SparqlParserDECIMAL_POSITIVE, SparqlParserDOUBLE_POSITIVE, SparqlParserINTEGER_NEGATIVE, SparqlParserDECIMAL_NEGATIVE, SparqlParserDOUBLE_NEGATIVE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(434)

			var _x = p.NumericLiteral()

			localctx.(*GraphTermContext).num = _x
		}

	case SparqlParserT__56, SparqlParserT__57:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(435)

			var _x = p.BooleanLiteral()

			localctx.(*GraphTermContext).bol = _x
		}

	case SparqlParserBLANK_NODE_LABEL, SparqlParserANON:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(436)

			var _x = p.BlankNode()

			localctx.(*GraphTermContext).blk = _x
		}

	case SparqlParserNIL:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(437)

			var _m = p.Match(SparqlParserNIL)

			localctx.(*GraphTermContext).emp = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ConditionalOrExpression() IConditionalOrExpressionContext

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) ConditionalOrExpression() IConditionalOrExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionalOrExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionalOrExpressionContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *SparqlParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, SparqlParserRULE_expression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(440)
		p.ConditionalOrExpression()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConditionalOrExpressionContext is an interface to support dynamic dispatch.
type IConditionalOrExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllConditionalAndExpression() []IConditionalAndExpressionContext
	ConditionalAndExpression(i int) IConditionalAndExpressionContext

	// IsConditionalOrExpressionContext differentiates from other interfaces.
	IsConditionalOrExpressionContext()
}

type ConditionalOrExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionalOrExpressionContext() *ConditionalOrExpressionContext {
	var p = new(ConditionalOrExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_conditionalOrExpression
	return p
}

func InitEmptyConditionalOrExpressionContext(p *ConditionalOrExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_conditionalOrExpression
}

func (*ConditionalOrExpressionContext) IsConditionalOrExpressionContext() {}

func NewConditionalOrExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionalOrExpressionContext {
	var p = new(ConditionalOrExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_conditionalOrExpression

	return p
}

func (s *ConditionalOrExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionalOrExpressionContext) AllConditionalAndExpression() []IConditionalAndExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConditionalAndExpressionContext); ok {
			len++
		}
	}

	tst := make([]IConditionalAndExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConditionalAndExpressionContext); ok {
			tst[i] = t.(IConditionalAndExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ConditionalOrExpressionContext) ConditionalAndExpression(i int) IConditionalAndExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionalAndExpressionContext); ok {
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

	return t.(IConditionalAndExpressionContext)
}

func (s *ConditionalOrExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionalOrExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionalOrExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.EnterConditionalOrExpression(s)
	}
}

func (s *ConditionalOrExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.ExitConditionalOrExpression(s)
	}
}

func (p *SparqlParser) ConditionalOrExpression() (localctx IConditionalOrExpressionContext) {
	localctx = NewConditionalOrExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, SparqlParserRULE_conditionalOrExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(442)
		p.ConditionalAndExpression()
	}
	p.SetState(447)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SparqlParserT__32 {
		{
			p.SetState(443)
			p.Match(SparqlParserT__32)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(444)
			p.ConditionalAndExpression()
		}

		p.SetState(449)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConditionalAndExpressionContext is an interface to support dynamic dispatch.
type IConditionalAndExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllValueLogical() []IValueLogicalContext
	ValueLogical(i int) IValueLogicalContext

	// IsConditionalAndExpressionContext differentiates from other interfaces.
	IsConditionalAndExpressionContext()
}

type ConditionalAndExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionalAndExpressionContext() *ConditionalAndExpressionContext {
	var p = new(ConditionalAndExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_conditionalAndExpression
	return p
}

func InitEmptyConditionalAndExpressionContext(p *ConditionalAndExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_conditionalAndExpression
}

func (*ConditionalAndExpressionContext) IsConditionalAndExpressionContext() {}

func NewConditionalAndExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionalAndExpressionContext {
	var p = new(ConditionalAndExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_conditionalAndExpression

	return p
}

func (s *ConditionalAndExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionalAndExpressionContext) AllValueLogical() []IValueLogicalContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IValueLogicalContext); ok {
			len++
		}
	}

	tst := make([]IValueLogicalContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IValueLogicalContext); ok {
			tst[i] = t.(IValueLogicalContext)
			i++
		}
	}

	return tst
}

func (s *ConditionalAndExpressionContext) ValueLogical(i int) IValueLogicalContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueLogicalContext); ok {
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

	return t.(IValueLogicalContext)
}

func (s *ConditionalAndExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionalAndExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionalAndExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.EnterConditionalAndExpression(s)
	}
}

func (s *ConditionalAndExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.ExitConditionalAndExpression(s)
	}
}

func (p *SparqlParser) ConditionalAndExpression() (localctx IConditionalAndExpressionContext) {
	localctx = NewConditionalAndExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, SparqlParserRULE_conditionalAndExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(450)
		p.ValueLogical()
	}
	p.SetState(455)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SparqlParserT__33 {
		{
			p.SetState(451)
			p.Match(SparqlParserT__33)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(452)
			p.ValueLogical()
		}

		p.SetState(457)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IValueLogicalContext is an interface to support dynamic dispatch.
type IValueLogicalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RelationalExpression() IRelationalExpressionContext

	// IsValueLogicalContext differentiates from other interfaces.
	IsValueLogicalContext()
}

type ValueLogicalContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueLogicalContext() *ValueLogicalContext {
	var p = new(ValueLogicalContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_valueLogical
	return p
}

func InitEmptyValueLogicalContext(p *ValueLogicalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_valueLogical
}

func (*ValueLogicalContext) IsValueLogicalContext() {}

func NewValueLogicalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueLogicalContext {
	var p = new(ValueLogicalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_valueLogical

	return p
}

func (s *ValueLogicalContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueLogicalContext) RelationalExpression() IRelationalExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationalExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationalExpressionContext)
}

func (s *ValueLogicalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueLogicalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueLogicalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.EnterValueLogical(s)
	}
}

func (s *ValueLogicalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.ExitValueLogical(s)
	}
}

func (p *SparqlParser) ValueLogical() (localctx IValueLogicalContext) {
	localctx = NewValueLogicalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, SparqlParserRULE_valueLogical)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(458)
		p.RelationalExpression()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRelationalExpressionContext is an interface to support dynamic dispatch.
type IRelationalExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllNumericExpression() []INumericExpressionContext
	NumericExpression(i int) INumericExpressionContext

	// IsRelationalExpressionContext differentiates from other interfaces.
	IsRelationalExpressionContext()
}

type RelationalExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelationalExpressionContext() *RelationalExpressionContext {
	var p = new(RelationalExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_relationalExpression
	return p
}

func InitEmptyRelationalExpressionContext(p *RelationalExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_relationalExpression
}

func (*RelationalExpressionContext) IsRelationalExpressionContext() {}

func NewRelationalExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationalExpressionContext {
	var p = new(RelationalExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_relationalExpression

	return p
}

func (s *RelationalExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationalExpressionContext) AllNumericExpression() []INumericExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INumericExpressionContext); ok {
			len++
		}
	}

	tst := make([]INumericExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INumericExpressionContext); ok {
			tst[i] = t.(INumericExpressionContext)
			i++
		}
	}

	return tst
}

func (s *RelationalExpressionContext) NumericExpression(i int) INumericExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumericExpressionContext); ok {
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

	return t.(INumericExpressionContext)
}

func (s *RelationalExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationalExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelationalExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.EnterRelationalExpression(s)
	}
}

func (s *RelationalExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.ExitRelationalExpression(s)
	}
}

func (p *SparqlParser) RelationalExpression() (localctx IRelationalExpressionContext) {
	localctx = NewRelationalExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, SparqlParserRULE_relationalExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(460)
		p.NumericExpression()
	}
	p.SetState(473)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	switch p.GetTokenStream().LA(1) {
	case SparqlParserT__34:
		{
			p.SetState(461)
			p.Match(SparqlParserT__34)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(462)
			p.NumericExpression()
		}

	case SparqlParserT__35:
		{
			p.SetState(463)
			p.Match(SparqlParserT__35)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(464)
			p.NumericExpression()
		}

	case SparqlParserT__36:
		{
			p.SetState(465)
			p.Match(SparqlParserT__36)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(466)
			p.NumericExpression()
		}

	case SparqlParserT__37:
		{
			p.SetState(467)
			p.Match(SparqlParserT__37)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(468)
			p.NumericExpression()
		}

	case SparqlParserT__38:
		{
			p.SetState(469)
			p.Match(SparqlParserT__38)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(470)
			p.NumericExpression()
		}

	case SparqlParserT__39:
		{
			p.SetState(471)
			p.Match(SparqlParserT__39)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(472)
			p.NumericExpression()
		}

	case SparqlParserT__26, SparqlParserT__27, SparqlParserT__32, SparqlParserT__33:

	default:
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INumericExpressionContext is an interface to support dynamic dispatch.
type INumericExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AdditiveExpression() IAdditiveExpressionContext

	// IsNumericExpressionContext differentiates from other interfaces.
	IsNumericExpressionContext()
}

type NumericExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumericExpressionContext() *NumericExpressionContext {
	var p = new(NumericExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_numericExpression
	return p
}

func InitEmptyNumericExpressionContext(p *NumericExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_numericExpression
}

func (*NumericExpressionContext) IsNumericExpressionContext() {}

func NewNumericExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumericExpressionContext {
	var p = new(NumericExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_numericExpression

	return p
}

func (s *NumericExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *NumericExpressionContext) AdditiveExpression() IAdditiveExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAdditiveExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAdditiveExpressionContext)
}

func (s *NumericExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumericExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumericExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.EnterNumericExpression(s)
	}
}

func (s *NumericExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.ExitNumericExpression(s)
	}
}

func (p *SparqlParser) NumericExpression() (localctx INumericExpressionContext) {
	localctx = NewNumericExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, SparqlParserRULE_numericExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(475)
		p.AdditiveExpression()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAdditiveExpressionContext is an interface to support dynamic dispatch.
type IAdditiveExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllMultiplicativeExpression() []IMultiplicativeExpressionContext
	MultiplicativeExpression(i int) IMultiplicativeExpressionContext
	AllNumericLiteralPositive() []INumericLiteralPositiveContext
	NumericLiteralPositive(i int) INumericLiteralPositiveContext
	AllNumericLiteralNegative() []INumericLiteralNegativeContext
	NumericLiteralNegative(i int) INumericLiteralNegativeContext

	// IsAdditiveExpressionContext differentiates from other interfaces.
	IsAdditiveExpressionContext()
}

type AdditiveExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAdditiveExpressionContext() *AdditiveExpressionContext {
	var p = new(AdditiveExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_additiveExpression
	return p
}

func InitEmptyAdditiveExpressionContext(p *AdditiveExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_additiveExpression
}

func (*AdditiveExpressionContext) IsAdditiveExpressionContext() {}

func NewAdditiveExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AdditiveExpressionContext {
	var p = new(AdditiveExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_additiveExpression

	return p
}

func (s *AdditiveExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *AdditiveExpressionContext) AllMultiplicativeExpression() []IMultiplicativeExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMultiplicativeExpressionContext); ok {
			len++
		}
	}

	tst := make([]IMultiplicativeExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMultiplicativeExpressionContext); ok {
			tst[i] = t.(IMultiplicativeExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AdditiveExpressionContext) MultiplicativeExpression(i int) IMultiplicativeExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultiplicativeExpressionContext); ok {
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

	return t.(IMultiplicativeExpressionContext)
}

func (s *AdditiveExpressionContext) AllNumericLiteralPositive() []INumericLiteralPositiveContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INumericLiteralPositiveContext); ok {
			len++
		}
	}

	tst := make([]INumericLiteralPositiveContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INumericLiteralPositiveContext); ok {
			tst[i] = t.(INumericLiteralPositiveContext)
			i++
		}
	}

	return tst
}

func (s *AdditiveExpressionContext) NumericLiteralPositive(i int) INumericLiteralPositiveContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumericLiteralPositiveContext); ok {
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

	return t.(INumericLiteralPositiveContext)
}

func (s *AdditiveExpressionContext) AllNumericLiteralNegative() []INumericLiteralNegativeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INumericLiteralNegativeContext); ok {
			len++
		}
	}

	tst := make([]INumericLiteralNegativeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INumericLiteralNegativeContext); ok {
			tst[i] = t.(INumericLiteralNegativeContext)
			i++
		}
	}

	return tst
}

func (s *AdditiveExpressionContext) NumericLiteralNegative(i int) INumericLiteralNegativeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumericLiteralNegativeContext); ok {
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

	return t.(INumericLiteralNegativeContext)
}

func (s *AdditiveExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdditiveExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AdditiveExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.EnterAdditiveExpression(s)
	}
}

func (s *AdditiveExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.ExitAdditiveExpression(s)
	}
}

func (p *SparqlParser) AdditiveExpression() (localctx IAdditiveExpressionContext) {
	localctx = NewAdditiveExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, SparqlParserRULE_additiveExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(477)
		p.MultiplicativeExpression()
	}
	p.SetState(486)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-41)) & ^0x3f) == 0 && ((int64(1)<<(_la-41))&16911433731) != 0 {
		p.SetState(484)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case SparqlParserT__40:
			{
				p.SetState(478)
				p.Match(SparqlParserT__40)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(479)
				p.MultiplicativeExpression()
			}

		case SparqlParserT__41:
			{
				p.SetState(480)
				p.Match(SparqlParserT__41)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(481)
				p.MultiplicativeExpression()
			}

		case SparqlParserINTEGER_POSITIVE, SparqlParserDECIMAL_POSITIVE, SparqlParserDOUBLE_POSITIVE:
			{
				p.SetState(482)
				p.NumericLiteralPositive()
			}

		case SparqlParserINTEGER_NEGATIVE, SparqlParserDECIMAL_NEGATIVE, SparqlParserDOUBLE_NEGATIVE:
			{
				p.SetState(483)
				p.NumericLiteralNegative()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(488)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMultiplicativeExpressionContext is an interface to support dynamic dispatch.
type IMultiplicativeExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllUnaryExpression() []IUnaryExpressionContext
	UnaryExpression(i int) IUnaryExpressionContext

	// IsMultiplicativeExpressionContext differentiates from other interfaces.
	IsMultiplicativeExpressionContext()
}

type MultiplicativeExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiplicativeExpressionContext() *MultiplicativeExpressionContext {
	var p = new(MultiplicativeExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_multiplicativeExpression
	return p
}

func InitEmptyMultiplicativeExpressionContext(p *MultiplicativeExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_multiplicativeExpression
}

func (*MultiplicativeExpressionContext) IsMultiplicativeExpressionContext() {}

func NewMultiplicativeExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiplicativeExpressionContext {
	var p = new(MultiplicativeExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_multiplicativeExpression

	return p
}

func (s *MultiplicativeExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiplicativeExpressionContext) AllUnaryExpression() []IUnaryExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IUnaryExpressionContext); ok {
			len++
		}
	}

	tst := make([]IUnaryExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IUnaryExpressionContext); ok {
			tst[i] = t.(IUnaryExpressionContext)
			i++
		}
	}

	return tst
}

func (s *MultiplicativeExpressionContext) UnaryExpression(i int) IUnaryExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryExpressionContext); ok {
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

	return t.(IUnaryExpressionContext)
}

func (s *MultiplicativeExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplicativeExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultiplicativeExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.EnterMultiplicativeExpression(s)
	}
}

func (s *MultiplicativeExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.ExitMultiplicativeExpression(s)
	}
}

func (p *SparqlParser) MultiplicativeExpression() (localctx IMultiplicativeExpressionContext) {
	localctx = NewMultiplicativeExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 104, SparqlParserRULE_multiplicativeExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(489)
		p.UnaryExpression()
	}
	p.SetState(496)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SparqlParserT__5 || _la == SparqlParserT__42 {
		p.SetState(494)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case SparqlParserT__5:
			{
				p.SetState(490)
				p.Match(SparqlParserT__5)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(491)
				p.UnaryExpression()
			}

		case SparqlParserT__42:
			{
				p.SetState(492)
				p.Match(SparqlParserT__42)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(493)
				p.UnaryExpression()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(498)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUnaryExpressionContext is an interface to support dynamic dispatch.
type IUnaryExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PrimaryExpression() IPrimaryExpressionContext

	// IsUnaryExpressionContext differentiates from other interfaces.
	IsUnaryExpressionContext()
}

type UnaryExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryExpressionContext() *UnaryExpressionContext {
	var p = new(UnaryExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_unaryExpression
	return p
}

func InitEmptyUnaryExpressionContext(p *UnaryExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_unaryExpression
}

func (*UnaryExpressionContext) IsUnaryExpressionContext() {}

func NewUnaryExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryExpressionContext {
	var p = new(UnaryExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_unaryExpression

	return p
}

func (s *UnaryExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryExpressionContext) PrimaryExpression() IPrimaryExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryExpressionContext)
}

func (s *UnaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.EnterUnaryExpression(s)
	}
}

func (s *UnaryExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.ExitUnaryExpression(s)
	}
}

func (p *SparqlParser) UnaryExpression() (localctx IUnaryExpressionContext) {
	localctx = NewUnaryExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 106, SparqlParserRULE_unaryExpression)
	p.SetState(506)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SparqlParserT__43:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(499)
			p.Match(SparqlParserT__43)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(500)
			p.PrimaryExpression()
		}

	case SparqlParserT__40:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(501)
			p.Match(SparqlParserT__40)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(502)
			p.PrimaryExpression()
		}

	case SparqlParserT__41:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(503)
			p.Match(SparqlParserT__41)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(504)
			p.PrimaryExpression()
		}

	case SparqlParserT__25, SparqlParserT__44, SparqlParserT__45, SparqlParserT__46, SparqlParserT__47, SparqlParserT__48, SparqlParserT__49, SparqlParserT__50, SparqlParserT__51, SparqlParserT__52, SparqlParserT__53, SparqlParserT__54, SparqlParserT__56, SparqlParserT__57, SparqlParserIRI_REF, SparqlParserPNAME_NS, SparqlParserPNAME_LN, SparqlParserVAR1, SparqlParserVAR2, SparqlParserINTEGER, SparqlParserDECIMAL, SparqlParserDOUBLE, SparqlParserINTEGER_POSITIVE, SparqlParserDECIMAL_POSITIVE, SparqlParserDOUBLE_POSITIVE, SparqlParserINTEGER_NEGATIVE, SparqlParserDECIMAL_NEGATIVE, SparqlParserDOUBLE_NEGATIVE, SparqlParserSTRING_LITERAL1, SparqlParserSTRING_LITERAL2:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(505)
			p.PrimaryExpression()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrimaryExpressionContext is an interface to support dynamic dispatch.
type IPrimaryExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BrackettedExpression() IBrackettedExpressionContext
	BuiltInCall() IBuiltInCallContext
	IriRefOrFunction() IIriRefOrFunctionContext
	RdfLiteral() IRdfLiteralContext
	NumericLiteral() INumericLiteralContext
	BooleanLiteral() IBooleanLiteralContext
	Varx() IVarxContext

	// IsPrimaryExpressionContext differentiates from other interfaces.
	IsPrimaryExpressionContext()
}

type PrimaryExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryExpressionContext() *PrimaryExpressionContext {
	var p = new(PrimaryExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_primaryExpression
	return p
}

func InitEmptyPrimaryExpressionContext(p *PrimaryExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_primaryExpression
}

func (*PrimaryExpressionContext) IsPrimaryExpressionContext() {}

func NewPrimaryExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryExpressionContext {
	var p = new(PrimaryExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_primaryExpression

	return p
}

func (s *PrimaryExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryExpressionContext) BrackettedExpression() IBrackettedExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBrackettedExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBrackettedExpressionContext)
}

func (s *PrimaryExpressionContext) BuiltInCall() IBuiltInCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBuiltInCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBuiltInCallContext)
}

func (s *PrimaryExpressionContext) IriRefOrFunction() IIriRefOrFunctionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIriRefOrFunctionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIriRefOrFunctionContext)
}

func (s *PrimaryExpressionContext) RdfLiteral() IRdfLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRdfLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRdfLiteralContext)
}

func (s *PrimaryExpressionContext) NumericLiteral() INumericLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumericLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumericLiteralContext)
}

func (s *PrimaryExpressionContext) BooleanLiteral() IBooleanLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBooleanLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBooleanLiteralContext)
}

func (s *PrimaryExpressionContext) Varx() IVarxContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarxContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarxContext)
}

func (s *PrimaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.EnterPrimaryExpression(s)
	}
}

func (s *PrimaryExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.ExitPrimaryExpression(s)
	}
}

func (p *SparqlParser) PrimaryExpression() (localctx IPrimaryExpressionContext) {
	localctx = NewPrimaryExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 108, SparqlParserRULE_primaryExpression)
	p.SetState(515)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SparqlParserT__25:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(508)
			p.BrackettedExpression()
		}

	case SparqlParserT__44, SparqlParserT__45, SparqlParserT__46, SparqlParserT__47, SparqlParserT__48, SparqlParserT__49, SparqlParserT__50, SparqlParserT__51, SparqlParserT__52, SparqlParserT__53, SparqlParserT__54:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(509)
			p.BuiltInCall()
		}

	case SparqlParserIRI_REF, SparqlParserPNAME_NS, SparqlParserPNAME_LN:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(510)
			p.IriRefOrFunction()
		}

	case SparqlParserSTRING_LITERAL1, SparqlParserSTRING_LITERAL2:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(511)
			p.RdfLiteral()
		}

	case SparqlParserINTEGER, SparqlParserDECIMAL, SparqlParserDOUBLE, SparqlParserINTEGER_POSITIVE, SparqlParserDECIMAL_POSITIVE, SparqlParserDOUBLE_POSITIVE, SparqlParserINTEGER_NEGATIVE, SparqlParserDECIMAL_NEGATIVE, SparqlParserDOUBLE_NEGATIVE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(512)
			p.NumericLiteral()
		}

	case SparqlParserT__56, SparqlParserT__57:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(513)
			p.BooleanLiteral()
		}

	case SparqlParserVAR1, SparqlParserVAR2:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(514)
			p.Varx()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBrackettedExpressionContext is an interface to support dynamic dispatch.
type IBrackettedExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext

	// IsBrackettedExpressionContext differentiates from other interfaces.
	IsBrackettedExpressionContext()
}

type BrackettedExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBrackettedExpressionContext() *BrackettedExpressionContext {
	var p = new(BrackettedExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_brackettedExpression
	return p
}

func InitEmptyBrackettedExpressionContext(p *BrackettedExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_brackettedExpression
}

func (*BrackettedExpressionContext) IsBrackettedExpressionContext() {}

func NewBrackettedExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BrackettedExpressionContext {
	var p = new(BrackettedExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_brackettedExpression

	return p
}

func (s *BrackettedExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *BrackettedExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BrackettedExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BrackettedExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BrackettedExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.EnterBrackettedExpression(s)
	}
}

func (s *BrackettedExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.ExitBrackettedExpression(s)
	}
}

func (p *SparqlParser) BrackettedExpression() (localctx IBrackettedExpressionContext) {
	localctx = NewBrackettedExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 110, SparqlParserRULE_brackettedExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(517)
		p.Match(SparqlParserT__25)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(518)
		p.Expression()
	}
	{
		p.SetState(519)
		p.Match(SparqlParserT__27)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBuiltInCallContext is an interface to support dynamic dispatch.
type IBuiltInCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	Varx() IVarxContext
	RegexExpression() IRegexExpressionContext

	// IsBuiltInCallContext differentiates from other interfaces.
	IsBuiltInCallContext()
}

type BuiltInCallContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBuiltInCallContext() *BuiltInCallContext {
	var p = new(BuiltInCallContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_builtInCall
	return p
}

func InitEmptyBuiltInCallContext(p *BuiltInCallContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_builtInCall
}

func (*BuiltInCallContext) IsBuiltInCallContext() {}

func NewBuiltInCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BuiltInCallContext {
	var p = new(BuiltInCallContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_builtInCall

	return p
}

func (s *BuiltInCallContext) GetParser() antlr.Parser { return s.parser }

func (s *BuiltInCallContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *BuiltInCallContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *BuiltInCallContext) Varx() IVarxContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarxContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarxContext)
}

func (s *BuiltInCallContext) RegexExpression() IRegexExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRegexExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRegexExpressionContext)
}

func (s *BuiltInCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BuiltInCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BuiltInCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.EnterBuiltInCall(s)
	}
}

func (s *BuiltInCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.ExitBuiltInCall(s)
	}
}

func (p *SparqlParser) BuiltInCall() (localctx IBuiltInCallContext) {
	localctx = NewBuiltInCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 112, SparqlParserRULE_builtInCall)
	p.SetState(576)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SparqlParserT__44:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(521)
			p.Match(SparqlParserT__44)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(522)
			p.Match(SparqlParserT__25)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(523)
			p.Expression()
		}
		{
			p.SetState(524)
			p.Match(SparqlParserT__27)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SparqlParserT__45:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(526)
			p.Match(SparqlParserT__45)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(527)
			p.Match(SparqlParserT__25)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(528)
			p.Expression()
		}
		{
			p.SetState(529)
			p.Match(SparqlParserT__27)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SparqlParserT__46:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(531)
			p.Match(SparqlParserT__46)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(532)
			p.Match(SparqlParserT__25)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(533)
			p.Expression()
		}
		{
			p.SetState(534)
			p.Match(SparqlParserT__26)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(535)
			p.Expression()
		}
		{
			p.SetState(536)
			p.Match(SparqlParserT__27)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SparqlParserT__47:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(538)
			p.Match(SparqlParserT__47)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(539)
			p.Match(SparqlParserT__25)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(540)
			p.Expression()
		}
		{
			p.SetState(541)
			p.Match(SparqlParserT__27)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SparqlParserT__48:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(543)
			p.Match(SparqlParserT__48)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(544)
			p.Match(SparqlParserT__25)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(545)
			p.Varx()
		}
		{
			p.SetState(546)
			p.Match(SparqlParserT__27)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SparqlParserT__49:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(548)
			p.Match(SparqlParserT__49)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(549)
			p.Match(SparqlParserT__25)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(550)
			p.Expression()
		}
		{
			p.SetState(551)
			p.Match(SparqlParserT__26)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(552)
			p.Expression()
		}
		{
			p.SetState(553)
			p.Match(SparqlParserT__27)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SparqlParserT__50:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(555)
			p.Match(SparqlParserT__50)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(556)
			p.Match(SparqlParserT__25)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(557)
			p.Expression()
		}
		{
			p.SetState(558)
			p.Match(SparqlParserT__27)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SparqlParserT__51:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(560)
			p.Match(SparqlParserT__51)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(561)
			p.Match(SparqlParserT__25)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(562)
			p.Expression()
		}
		{
			p.SetState(563)
			p.Match(SparqlParserT__27)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SparqlParserT__52:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(565)
			p.Match(SparqlParserT__52)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(566)
			p.Match(SparqlParserT__25)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(567)
			p.Expression()
		}
		{
			p.SetState(568)
			p.Match(SparqlParserT__27)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SparqlParserT__53:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(570)
			p.Match(SparqlParserT__53)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(571)
			p.Match(SparqlParserT__25)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(572)
			p.Expression()
		}
		{
			p.SetState(573)
			p.Match(SparqlParserT__27)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SparqlParserT__54:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(575)
			p.RegexExpression()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRegexExpressionContext is an interface to support dynamic dispatch.
type IRegexExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext

	// IsRegexExpressionContext differentiates from other interfaces.
	IsRegexExpressionContext()
}

type RegexExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRegexExpressionContext() *RegexExpressionContext {
	var p = new(RegexExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_regexExpression
	return p
}

func InitEmptyRegexExpressionContext(p *RegexExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_regexExpression
}

func (*RegexExpressionContext) IsRegexExpressionContext() {}

func NewRegexExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RegexExpressionContext {
	var p = new(RegexExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_regexExpression

	return p
}

func (s *RegexExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *RegexExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *RegexExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *RegexExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RegexExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RegexExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.EnterRegexExpression(s)
	}
}

func (s *RegexExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.ExitRegexExpression(s)
	}
}

func (p *SparqlParser) RegexExpression() (localctx IRegexExpressionContext) {
	localctx = NewRegexExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 114, SparqlParserRULE_regexExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(578)
		p.Match(SparqlParserT__54)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(579)
		p.Match(SparqlParserT__25)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(580)
		p.Expression()
	}
	{
		p.SetState(581)
		p.Match(SparqlParserT__26)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(582)
		p.Expression()
	}
	p.SetState(585)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SparqlParserT__26 {
		{
			p.SetState(583)
			p.Match(SparqlParserT__26)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(584)
			p.Expression()
		}

	}
	{
		p.SetState(587)
		p.Match(SparqlParserT__27)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIriRefOrFunctionContext is an interface to support dynamic dispatch.
type IIriRefOrFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IriRef() IIriRefContext
	ArgList() IArgListContext

	// IsIriRefOrFunctionContext differentiates from other interfaces.
	IsIriRefOrFunctionContext()
}

type IriRefOrFunctionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIriRefOrFunctionContext() *IriRefOrFunctionContext {
	var p = new(IriRefOrFunctionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_iriRefOrFunction
	return p
}

func InitEmptyIriRefOrFunctionContext(p *IriRefOrFunctionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_iriRefOrFunction
}

func (*IriRefOrFunctionContext) IsIriRefOrFunctionContext() {}

func NewIriRefOrFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IriRefOrFunctionContext {
	var p = new(IriRefOrFunctionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_iriRefOrFunction

	return p
}

func (s *IriRefOrFunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *IriRefOrFunctionContext) IriRef() IIriRefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIriRefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIriRefContext)
}

func (s *IriRefOrFunctionContext) ArgList() IArgListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgListContext)
}

func (s *IriRefOrFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IriRefOrFunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IriRefOrFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.EnterIriRefOrFunction(s)
	}
}

func (s *IriRefOrFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.ExitIriRefOrFunction(s)
	}
}

func (p *SparqlParser) IriRefOrFunction() (localctx IIriRefOrFunctionContext) {
	localctx = NewIriRefOrFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 116, SparqlParserRULE_iriRefOrFunction)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(589)
		p.IriRef()
	}
	p.SetState(591)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SparqlParserT__25 || _la == SparqlParserNIL {
		{
			p.SetState(590)
			p.ArgList()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRdfLiteralContext is an interface to support dynamic dispatch.
type IRdfLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLang returns the lang token.
	GetLang() antlr.Token

	// SetLang sets the lang token.
	SetLang(antlr.Token)

	// GetStr returns the str rule contexts.
	GetStr() IRdfstringContext

	// GetIri returns the iri rule contexts.
	GetIri() IIriRefContext

	// SetStr sets the str rule contexts.
	SetStr(IRdfstringContext)

	// SetIri sets the iri rule contexts.
	SetIri(IIriRefContext)

	// Getter signatures
	Rdfstring() IRdfstringContext
	LANGTAG() antlr.TerminalNode
	IriRef() IIriRefContext

	// IsRdfLiteralContext differentiates from other interfaces.
	IsRdfLiteralContext()
}

type RdfLiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	str    IRdfstringContext
	lang   antlr.Token
	iri    IIriRefContext
}

func NewEmptyRdfLiteralContext() *RdfLiteralContext {
	var p = new(RdfLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_rdfLiteral
	return p
}

func InitEmptyRdfLiteralContext(p *RdfLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_rdfLiteral
}

func (*RdfLiteralContext) IsRdfLiteralContext() {}

func NewRdfLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RdfLiteralContext {
	var p = new(RdfLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_rdfLiteral

	return p
}

func (s *RdfLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *RdfLiteralContext) GetLang() antlr.Token { return s.lang }

func (s *RdfLiteralContext) SetLang(v antlr.Token) { s.lang = v }

func (s *RdfLiteralContext) GetStr() IRdfstringContext { return s.str }

func (s *RdfLiteralContext) GetIri() IIriRefContext { return s.iri }

func (s *RdfLiteralContext) SetStr(v IRdfstringContext) { s.str = v }

func (s *RdfLiteralContext) SetIri(v IIriRefContext) { s.iri = v }

func (s *RdfLiteralContext) Rdfstring() IRdfstringContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRdfstringContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRdfstringContext)
}

func (s *RdfLiteralContext) LANGTAG() antlr.TerminalNode {
	return s.GetToken(SparqlParserLANGTAG, 0)
}

func (s *RdfLiteralContext) IriRef() IIriRefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIriRefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIriRefContext)
}

func (s *RdfLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RdfLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RdfLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.EnterRdfLiteral(s)
	}
}

func (s *RdfLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.ExitRdfLiteral(s)
	}
}

func (p *SparqlParser) RdfLiteral() (localctx IRdfLiteralContext) {
	localctx = NewRdfLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 118, SparqlParserRULE_rdfLiteral)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(593)

		var _x = p.Rdfstring()

		localctx.(*RdfLiteralContext).str = _x
	}
	p.SetState(597)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	switch p.GetTokenStream().LA(1) {
	case SparqlParserLANGTAG:
		{
			p.SetState(594)

			var _m = p.Match(SparqlParserLANGTAG)

			localctx.(*RdfLiteralContext).lang = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SparqlParserT__55:
		{
			p.SetState(595)
			p.Match(SparqlParserT__55)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(596)

			var _x = p.IriRef()

			localctx.(*RdfLiteralContext).iri = _x
		}

	case SparqlParserT__5, SparqlParserT__18, SparqlParserT__19, SparqlParserT__20, SparqlParserT__21, SparqlParserT__22, SparqlParserT__24, SparqlParserT__25, SparqlParserT__26, SparqlParserT__27, SparqlParserT__28, SparqlParserT__29, SparqlParserT__30, SparqlParserT__31, SparqlParserT__32, SparqlParserT__33, SparqlParserT__34, SparqlParserT__35, SparqlParserT__36, SparqlParserT__37, SparqlParserT__38, SparqlParserT__39, SparqlParserT__40, SparqlParserT__41, SparqlParserT__42, SparqlParserT__56, SparqlParserT__57, SparqlParserIRI_REF, SparqlParserPNAME_NS, SparqlParserPNAME_LN, SparqlParserBLANK_NODE_LABEL, SparqlParserVAR1, SparqlParserVAR2, SparqlParserINTEGER, SparqlParserDECIMAL, SparqlParserDOUBLE, SparqlParserINTEGER_POSITIVE, SparqlParserDECIMAL_POSITIVE, SparqlParserDOUBLE_POSITIVE, SparqlParserINTEGER_NEGATIVE, SparqlParserDECIMAL_NEGATIVE, SparqlParserDOUBLE_NEGATIVE, SparqlParserSTRING_LITERAL1, SparqlParserSTRING_LITERAL2, SparqlParserNIL, SparqlParserANON:

	default:
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INumericLiteralContext is an interface to support dynamic dispatch.
type INumericLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NumericLiteralUnsigned() INumericLiteralUnsignedContext
	NumericLiteralPositive() INumericLiteralPositiveContext
	NumericLiteralNegative() INumericLiteralNegativeContext

	// IsNumericLiteralContext differentiates from other interfaces.
	IsNumericLiteralContext()
}

type NumericLiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumericLiteralContext() *NumericLiteralContext {
	var p = new(NumericLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_numericLiteral
	return p
}

func InitEmptyNumericLiteralContext(p *NumericLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_numericLiteral
}

func (*NumericLiteralContext) IsNumericLiteralContext() {}

func NewNumericLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumericLiteralContext {
	var p = new(NumericLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_numericLiteral

	return p
}

func (s *NumericLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *NumericLiteralContext) NumericLiteralUnsigned() INumericLiteralUnsignedContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumericLiteralUnsignedContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumericLiteralUnsignedContext)
}

func (s *NumericLiteralContext) NumericLiteralPositive() INumericLiteralPositiveContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumericLiteralPositiveContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumericLiteralPositiveContext)
}

func (s *NumericLiteralContext) NumericLiteralNegative() INumericLiteralNegativeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumericLiteralNegativeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumericLiteralNegativeContext)
}

func (s *NumericLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumericLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumericLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.EnterNumericLiteral(s)
	}
}

func (s *NumericLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.ExitNumericLiteral(s)
	}
}

func (p *SparqlParser) NumericLiteral() (localctx INumericLiteralContext) {
	localctx = NewNumericLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 120, SparqlParserRULE_numericLiteral)
	p.SetState(602)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SparqlParserINTEGER, SparqlParserDECIMAL, SparqlParserDOUBLE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(599)
			p.NumericLiteralUnsigned()
		}

	case SparqlParserINTEGER_POSITIVE, SparqlParserDECIMAL_POSITIVE, SparqlParserDOUBLE_POSITIVE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(600)
			p.NumericLiteralPositive()
		}

	case SparqlParserINTEGER_NEGATIVE, SparqlParserDECIMAL_NEGATIVE, SparqlParserDOUBLE_NEGATIVE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(601)
			p.NumericLiteralNegative()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INumericLiteralUnsignedContext is an interface to support dynamic dispatch.
type INumericLiteralUnsignedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INTEGER() antlr.TerminalNode
	DECIMAL() antlr.TerminalNode
	DOUBLE() antlr.TerminalNode

	// IsNumericLiteralUnsignedContext differentiates from other interfaces.
	IsNumericLiteralUnsignedContext()
}

type NumericLiteralUnsignedContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumericLiteralUnsignedContext() *NumericLiteralUnsignedContext {
	var p = new(NumericLiteralUnsignedContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_numericLiteralUnsigned
	return p
}

func InitEmptyNumericLiteralUnsignedContext(p *NumericLiteralUnsignedContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_numericLiteralUnsigned
}

func (*NumericLiteralUnsignedContext) IsNumericLiteralUnsignedContext() {}

func NewNumericLiteralUnsignedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumericLiteralUnsignedContext {
	var p = new(NumericLiteralUnsignedContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_numericLiteralUnsigned

	return p
}

func (s *NumericLiteralUnsignedContext) GetParser() antlr.Parser { return s.parser }

func (s *NumericLiteralUnsignedContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(SparqlParserINTEGER, 0)
}

func (s *NumericLiteralUnsignedContext) DECIMAL() antlr.TerminalNode {
	return s.GetToken(SparqlParserDECIMAL, 0)
}

func (s *NumericLiteralUnsignedContext) DOUBLE() antlr.TerminalNode {
	return s.GetToken(SparqlParserDOUBLE, 0)
}

func (s *NumericLiteralUnsignedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumericLiteralUnsignedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumericLiteralUnsignedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.EnterNumericLiteralUnsigned(s)
	}
}

func (s *NumericLiteralUnsignedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.ExitNumericLiteralUnsigned(s)
	}
}

func (p *SparqlParser) NumericLiteralUnsigned() (localctx INumericLiteralUnsignedContext) {
	localctx = NewNumericLiteralUnsignedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 122, SparqlParserRULE_numericLiteralUnsigned)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(604)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-66)) & ^0x3f) == 0 && ((int64(1)<<(_la-66))&7) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INumericLiteralPositiveContext is an interface to support dynamic dispatch.
type INumericLiteralPositiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INTEGER_POSITIVE() antlr.TerminalNode
	DECIMAL_POSITIVE() antlr.TerminalNode
	DOUBLE_POSITIVE() antlr.TerminalNode

	// IsNumericLiteralPositiveContext differentiates from other interfaces.
	IsNumericLiteralPositiveContext()
}

type NumericLiteralPositiveContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumericLiteralPositiveContext() *NumericLiteralPositiveContext {
	var p = new(NumericLiteralPositiveContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_numericLiteralPositive
	return p
}

func InitEmptyNumericLiteralPositiveContext(p *NumericLiteralPositiveContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_numericLiteralPositive
}

func (*NumericLiteralPositiveContext) IsNumericLiteralPositiveContext() {}

func NewNumericLiteralPositiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumericLiteralPositiveContext {
	var p = new(NumericLiteralPositiveContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_numericLiteralPositive

	return p
}

func (s *NumericLiteralPositiveContext) GetParser() antlr.Parser { return s.parser }

func (s *NumericLiteralPositiveContext) INTEGER_POSITIVE() antlr.TerminalNode {
	return s.GetToken(SparqlParserINTEGER_POSITIVE, 0)
}

func (s *NumericLiteralPositiveContext) DECIMAL_POSITIVE() antlr.TerminalNode {
	return s.GetToken(SparqlParserDECIMAL_POSITIVE, 0)
}

func (s *NumericLiteralPositiveContext) DOUBLE_POSITIVE() antlr.TerminalNode {
	return s.GetToken(SparqlParserDOUBLE_POSITIVE, 0)
}

func (s *NumericLiteralPositiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumericLiteralPositiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumericLiteralPositiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.EnterNumericLiteralPositive(s)
	}
}

func (s *NumericLiteralPositiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.ExitNumericLiteralPositive(s)
	}
}

func (p *SparqlParser) NumericLiteralPositive() (localctx INumericLiteralPositiveContext) {
	localctx = NewNumericLiteralPositiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 124, SparqlParserRULE_numericLiteralPositive)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(606)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-69)) & ^0x3f) == 0 && ((int64(1)<<(_la-69))&7) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INumericLiteralNegativeContext is an interface to support dynamic dispatch.
type INumericLiteralNegativeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INTEGER_NEGATIVE() antlr.TerminalNode
	DECIMAL_NEGATIVE() antlr.TerminalNode
	DOUBLE_NEGATIVE() antlr.TerminalNode

	// IsNumericLiteralNegativeContext differentiates from other interfaces.
	IsNumericLiteralNegativeContext()
}

type NumericLiteralNegativeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumericLiteralNegativeContext() *NumericLiteralNegativeContext {
	var p = new(NumericLiteralNegativeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_numericLiteralNegative
	return p
}

func InitEmptyNumericLiteralNegativeContext(p *NumericLiteralNegativeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_numericLiteralNegative
}

func (*NumericLiteralNegativeContext) IsNumericLiteralNegativeContext() {}

func NewNumericLiteralNegativeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumericLiteralNegativeContext {
	var p = new(NumericLiteralNegativeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_numericLiteralNegative

	return p
}

func (s *NumericLiteralNegativeContext) GetParser() antlr.Parser { return s.parser }

func (s *NumericLiteralNegativeContext) INTEGER_NEGATIVE() antlr.TerminalNode {
	return s.GetToken(SparqlParserINTEGER_NEGATIVE, 0)
}

func (s *NumericLiteralNegativeContext) DECIMAL_NEGATIVE() antlr.TerminalNode {
	return s.GetToken(SparqlParserDECIMAL_NEGATIVE, 0)
}

func (s *NumericLiteralNegativeContext) DOUBLE_NEGATIVE() antlr.TerminalNode {
	return s.GetToken(SparqlParserDOUBLE_NEGATIVE, 0)
}

func (s *NumericLiteralNegativeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumericLiteralNegativeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumericLiteralNegativeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.EnterNumericLiteralNegative(s)
	}
}

func (s *NumericLiteralNegativeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.ExitNumericLiteralNegative(s)
	}
}

func (p *SparqlParser) NumericLiteralNegative() (localctx INumericLiteralNegativeContext) {
	localctx = NewNumericLiteralNegativeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 126, SparqlParserRULE_numericLiteralNegative)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(608)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-72)) & ^0x3f) == 0 && ((int64(1)<<(_la-72))&7) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBooleanLiteralContext is an interface to support dynamic dispatch.
type IBooleanLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsBooleanLiteralContext differentiates from other interfaces.
	IsBooleanLiteralContext()
}

type BooleanLiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBooleanLiteralContext() *BooleanLiteralContext {
	var p = new(BooleanLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_booleanLiteral
	return p
}

func InitEmptyBooleanLiteralContext(p *BooleanLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_booleanLiteral
}

func (*BooleanLiteralContext) IsBooleanLiteralContext() {}

func NewBooleanLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooleanLiteralContext {
	var p = new(BooleanLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_booleanLiteral

	return p
}

func (s *BooleanLiteralContext) GetParser() antlr.Parser { return s.parser }
func (s *BooleanLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BooleanLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.EnterBooleanLiteral(s)
	}
}

func (s *BooleanLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.ExitBooleanLiteral(s)
	}
}

func (p *SparqlParser) BooleanLiteral() (localctx IBooleanLiteralContext) {
	localctx = NewBooleanLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 128, SparqlParserRULE_booleanLiteral)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(610)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SparqlParserT__56 || _la == SparqlParserT__57) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRdfstringContext is an interface to support dynamic dispatch.
type IRdfstringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING_LITERAL1() antlr.TerminalNode
	STRING_LITERAL2() antlr.TerminalNode

	// IsRdfstringContext differentiates from other interfaces.
	IsRdfstringContext()
}

type RdfstringContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRdfstringContext() *RdfstringContext {
	var p = new(RdfstringContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_rdfstring
	return p
}

func InitEmptyRdfstringContext(p *RdfstringContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_rdfstring
}

func (*RdfstringContext) IsRdfstringContext() {}

func NewRdfstringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RdfstringContext {
	var p = new(RdfstringContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_rdfstring

	return p
}

func (s *RdfstringContext) GetParser() antlr.Parser { return s.parser }

func (s *RdfstringContext) STRING_LITERAL1() antlr.TerminalNode {
	return s.GetToken(SparqlParserSTRING_LITERAL1, 0)
}

func (s *RdfstringContext) STRING_LITERAL2() antlr.TerminalNode {
	return s.GetToken(SparqlParserSTRING_LITERAL2, 0)
}

func (s *RdfstringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RdfstringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RdfstringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.EnterRdfstring(s)
	}
}

func (s *RdfstringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.ExitRdfstring(s)
	}
}

func (p *SparqlParser) Rdfstring() (localctx IRdfstringContext) {
	localctx = NewRdfstringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 130, SparqlParserRULE_rdfstring)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(612)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SparqlParserSTRING_LITERAL1 || _la == SparqlParserSTRING_LITERAL2) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIriRefContext is an interface to support dynamic dispatch.
type IIriRefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLiteral returns the literal token.
	GetLiteral() antlr.Token

	// SetLiteral sets the literal token.
	SetLiteral(antlr.Token)

	// GetPrefixed returns the prefixed rule contexts.
	GetPrefixed() IPrefixedNameContext

	// SetPrefixed sets the prefixed rule contexts.
	SetPrefixed(IPrefixedNameContext)

	// Getter signatures
	IRI_REF() antlr.TerminalNode
	PrefixedName() IPrefixedNameContext

	// IsIriRefContext differentiates from other interfaces.
	IsIriRefContext()
}

type IriRefContext struct {
	antlr.BaseParserRuleContext
	parser   antlr.Parser
	literal  antlr.Token
	prefixed IPrefixedNameContext
}

func NewEmptyIriRefContext() *IriRefContext {
	var p = new(IriRefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_iriRef
	return p
}

func InitEmptyIriRefContext(p *IriRefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_iriRef
}

func (*IriRefContext) IsIriRefContext() {}

func NewIriRefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IriRefContext {
	var p = new(IriRefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_iriRef

	return p
}

func (s *IriRefContext) GetParser() antlr.Parser { return s.parser }

func (s *IriRefContext) GetLiteral() antlr.Token { return s.literal }

func (s *IriRefContext) SetLiteral(v antlr.Token) { s.literal = v }

func (s *IriRefContext) GetPrefixed() IPrefixedNameContext { return s.prefixed }

func (s *IriRefContext) SetPrefixed(v IPrefixedNameContext) { s.prefixed = v }

func (s *IriRefContext) IRI_REF() antlr.TerminalNode {
	return s.GetToken(SparqlParserIRI_REF, 0)
}

func (s *IriRefContext) PrefixedName() IPrefixedNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrefixedNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrefixedNameContext)
}

func (s *IriRefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IriRefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IriRefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.EnterIriRef(s)
	}
}

func (s *IriRefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.ExitIriRef(s)
	}
}

func (p *SparqlParser) IriRef() (localctx IIriRefContext) {
	localctx = NewIriRefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 132, SparqlParserRULE_iriRef)
	p.SetState(616)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SparqlParserIRI_REF:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(614)

			var _m = p.Match(SparqlParserIRI_REF)

			localctx.(*IriRefContext).literal = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SparqlParserPNAME_NS, SparqlParserPNAME_LN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(615)

			var _x = p.PrefixedName()

			localctx.(*IriRefContext).prefixed = _x
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrefixedNameContext is an interface to support dynamic dispatch.
type IPrefixedNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PNAME_LN() antlr.TerminalNode
	PNAME_NS() antlr.TerminalNode

	// IsPrefixedNameContext differentiates from other interfaces.
	IsPrefixedNameContext()
}

type PrefixedNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrefixedNameContext() *PrefixedNameContext {
	var p = new(PrefixedNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_prefixedName
	return p
}

func InitEmptyPrefixedNameContext(p *PrefixedNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_prefixedName
}

func (*PrefixedNameContext) IsPrefixedNameContext() {}

func NewPrefixedNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrefixedNameContext {
	var p = new(PrefixedNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_prefixedName

	return p
}

func (s *PrefixedNameContext) GetParser() antlr.Parser { return s.parser }

func (s *PrefixedNameContext) PNAME_LN() antlr.TerminalNode {
	return s.GetToken(SparqlParserPNAME_LN, 0)
}

func (s *PrefixedNameContext) PNAME_NS() antlr.TerminalNode {
	return s.GetToken(SparqlParserPNAME_NS, 0)
}

func (s *PrefixedNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrefixedNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrefixedNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.EnterPrefixedName(s)
	}
}

func (s *PrefixedNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.ExitPrefixedName(s)
	}
}

func (p *SparqlParser) PrefixedName() (localctx IPrefixedNameContext) {
	localctx = NewPrefixedNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 134, SparqlParserRULE_prefixedName)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(618)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SparqlParserPNAME_NS || _la == SparqlParserPNAME_LN) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlankNodeContext is an interface to support dynamic dispatch.
type IBlankNodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BLANK_NODE_LABEL() antlr.TerminalNode
	ANON() antlr.TerminalNode

	// IsBlankNodeContext differentiates from other interfaces.
	IsBlankNodeContext()
}

type BlankNodeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlankNodeContext() *BlankNodeContext {
	var p = new(BlankNodeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_blankNode
	return p
}

func InitEmptyBlankNodeContext(p *BlankNodeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_blankNode
}

func (*BlankNodeContext) IsBlankNodeContext() {}

func NewBlankNodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlankNodeContext {
	var p = new(BlankNodeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_blankNode

	return p
}

func (s *BlankNodeContext) GetParser() antlr.Parser { return s.parser }

func (s *BlankNodeContext) BLANK_NODE_LABEL() antlr.TerminalNode {
	return s.GetToken(SparqlParserBLANK_NODE_LABEL, 0)
}

func (s *BlankNodeContext) ANON() antlr.TerminalNode {
	return s.GetToken(SparqlParserANON, 0)
}

func (s *BlankNodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlankNodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlankNodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.EnterBlankNode(s)
	}
}

func (s *BlankNodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlListener); ok {
		listenerT.ExitBlankNode(s)
	}
}

func (p *SparqlParser) BlankNode() (localctx IBlankNodeContext) {
	localctx = NewBlankNodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 136, SparqlParserRULE_blankNode)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(620)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SparqlParserBLANK_NODE_LABEL || _la == SparqlParserANON) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
