// Code generated from /home/lercher/go/src/github.com/lercher/rdf/sparql/Sparql.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 89, 737,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44,
	9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9,
	49, 4, 50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54,
	4, 55, 9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4, 58, 9, 58, 4, 59, 9, 59, 4,
	60, 9, 60, 4, 61, 9, 61, 4, 62, 9, 62, 4, 63, 9, 63, 4, 64, 9, 64, 4, 65,
	9, 65, 4, 66, 9, 66, 4, 67, 9, 67, 4, 68, 9, 68, 4, 69, 9, 69, 4, 70, 9,
	70, 4, 71, 9, 71, 4, 72, 9, 72, 4, 73, 9, 73, 4, 74, 9, 74, 4, 75, 9, 75,
	4, 76, 9, 76, 4, 77, 9, 77, 4, 78, 9, 78, 4, 79, 9, 79, 4, 80, 9, 80, 4,
	81, 9, 81, 4, 82, 9, 82, 4, 83, 9, 83, 4, 84, 9, 84, 4, 85, 9, 85, 4, 86,
	9, 86, 4, 87, 9, 87, 4, 88, 9, 88, 4, 89, 9, 89, 4, 90, 9, 90, 4, 91, 9,
	91, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3,
	8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3,
	16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18,
	3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3,
	21, 3, 21, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23,
	3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 25, 3, 25, 3,
	25, 3, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26,
	3, 27, 3, 27, 3, 28, 3, 28, 3, 29, 3, 29, 3, 30, 3, 30, 3, 31, 3, 31, 3,
	32, 3, 32, 3, 33, 3, 33, 3, 34, 3, 34, 3, 34, 3, 35, 3, 35, 3, 35, 3, 36,
	3, 36, 3, 37, 3, 37, 3, 37, 3, 38, 3, 38, 3, 39, 3, 39, 3, 40, 3, 40, 3,
	40, 3, 41, 3, 41, 3, 41, 3, 42, 3, 42, 3, 43, 3, 43, 3, 44, 3, 44, 3, 45,
	3, 45, 3, 46, 3, 46, 3, 46, 3, 46, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3,
	48, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48,
	3, 48, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3,
	50, 3, 50, 3, 50, 3, 50, 3, 50, 3, 50, 3, 51, 3, 51, 3, 51, 3, 51, 3, 51,
	3, 51, 3, 51, 3, 51, 3, 51, 3, 52, 3, 52, 3, 52, 3, 52, 3, 52, 3, 52, 3,
	53, 3, 53, 3, 53, 3, 53, 3, 53, 3, 53, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54,
	3, 54, 3, 54, 3, 54, 3, 55, 3, 55, 3, 55, 3, 55, 3, 55, 3, 55, 3, 55, 3,
	55, 3, 55, 3, 55, 3, 56, 3, 56, 3, 56, 3, 56, 3, 56, 3, 56, 3, 57, 3, 57,
	3, 57, 3, 58, 3, 58, 3, 58, 3, 58, 3, 58, 3, 59, 3, 59, 3, 59, 3, 59, 3,
	59, 3, 59, 3, 60, 3, 60, 3, 60, 7, 60, 468, 10, 60, 12, 60, 14, 60, 471,
	11, 60, 3, 60, 3, 60, 3, 61, 5, 61, 476, 10, 61, 3, 61, 3, 61, 3, 62, 3,
	62, 3, 62, 3, 63, 3, 63, 3, 63, 3, 63, 3, 63, 3, 64, 3, 64, 3, 64, 3, 65,
	3, 65, 3, 65, 3, 66, 3, 66, 6, 66, 496, 10, 66, 13, 66, 14, 66, 497, 3,
	66, 3, 66, 3, 66, 3, 66, 6, 66, 504, 10, 66, 13, 66, 14, 66, 505, 7, 66,
	508, 10, 66, 12, 66, 14, 66, 511, 11, 66, 3, 67, 6, 67, 514, 10, 67, 13,
	67, 14, 67, 515, 3, 68, 6, 68, 519, 10, 68, 13, 68, 14, 68, 520, 3, 68,
	3, 68, 7, 68, 525, 10, 68, 12, 68, 14, 68, 528, 11, 68, 3, 68, 3, 68, 6,
	68, 532, 10, 68, 13, 68, 14, 68, 533, 5, 68, 536, 10, 68, 3, 69, 6, 69,
	539, 10, 69, 13, 69, 14, 69, 540, 3, 69, 3, 69, 7, 69, 545, 10, 69, 12,
	69, 14, 69, 548, 11, 69, 3, 69, 3, 69, 3, 69, 3, 69, 6, 69, 554, 10, 69,
	13, 69, 14, 69, 555, 3, 69, 3, 69, 3, 69, 6, 69, 561, 10, 69, 13, 69, 14,
	69, 562, 3, 69, 3, 69, 5, 69, 567, 10, 69, 3, 70, 3, 70, 3, 70, 3, 71,
	3, 71, 3, 71, 3, 72, 3, 72, 3, 72, 3, 73, 3, 73, 3, 73, 3, 74, 3, 74, 3,
	74, 3, 75, 3, 75, 3, 75, 3, 76, 3, 76, 5, 76, 589, 10, 76, 3, 76, 6, 76,
	592, 10, 76, 13, 76, 14, 76, 593, 3, 77, 3, 77, 3, 77, 7, 77, 599, 10,
	77, 12, 77, 14, 77, 602, 11, 77, 3, 77, 3, 77, 3, 78, 3, 78, 3, 78, 7,
	78, 609, 10, 78, 12, 78, 14, 78, 612, 11, 78, 3, 78, 3, 78, 3, 79, 3, 79,
	3, 79, 3, 79, 3, 79, 3, 79, 3, 79, 5, 79, 623, 10, 79, 3, 79, 3, 79, 5,
	79, 627, 10, 79, 7, 79, 629, 10, 79, 12, 79, 14, 79, 632, 11, 79, 3, 79,
	3, 79, 3, 79, 3, 79, 3, 80, 3, 80, 3, 80, 3, 80, 3, 80, 3, 80, 3, 80, 5,
	80, 645, 10, 80, 3, 80, 3, 80, 5, 80, 649, 10, 80, 7, 80, 651, 10, 80,
	12, 80, 14, 80, 654, 11, 80, 3, 80, 3, 80, 3, 80, 3, 80, 3, 81, 3, 81,
	3, 81, 3, 82, 3, 82, 7, 82, 665, 10, 82, 12, 82, 14, 82, 668, 11, 82, 3,
	82, 3, 82, 3, 83, 3, 83, 7, 83, 674, 10, 83, 12, 83, 14, 83, 677, 11, 83,
	3, 83, 3, 83, 3, 84, 3, 84, 5, 84, 683, 10, 84, 3, 85, 3, 85, 5, 85, 687,
	10, 85, 3, 85, 3, 85, 3, 85, 7, 85, 692, 10, 85, 12, 85, 14, 85, 695, 11,
	85, 3, 86, 3, 86, 3, 86, 5, 86, 700, 10, 86, 3, 87, 3, 87, 3, 87, 7, 87,
	705, 10, 87, 12, 87, 14, 87, 708, 11, 87, 3, 87, 5, 87, 711, 10, 87, 3,
	88, 3, 88, 5, 88, 715, 10, 88, 3, 88, 3, 88, 7, 88, 719, 10, 88, 12, 88,
	14, 88, 722, 11, 88, 3, 88, 5, 88, 725, 10, 88, 3, 89, 3, 89, 3, 90, 3,
	90, 3, 91, 6, 91, 732, 10, 91, 13, 91, 14, 91, 733, 3, 91, 3, 91, 2, 2,
	92, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12,
	23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21,
	41, 22, 43, 23, 45, 24, 47, 25, 49, 26, 51, 27, 53, 28, 55, 29, 57, 30,
	59, 31, 61, 32, 63, 33, 65, 34, 67, 35, 69, 36, 71, 37, 73, 38, 75, 39,
	77, 40, 79, 41, 81, 42, 83, 43, 85, 44, 87, 45, 89, 46, 91, 47, 93, 48,
	95, 49, 97, 50, 99, 51, 101, 52, 103, 53, 105, 54, 107, 55, 109, 56, 111,
	57, 113, 58, 115, 59, 117, 60, 119, 61, 121, 62, 123, 63, 125, 64, 127,
	65, 129, 66, 131, 67, 133, 68, 135, 69, 137, 70, 139, 71, 141, 72, 143,
	73, 145, 74, 147, 75, 149, 76, 151, 77, 153, 78, 155, 79, 157, 80, 159,
	81, 161, 82, 163, 83, 165, 84, 167, 85, 169, 86, 171, 2, 173, 87, 175,
	88, 177, 2, 179, 2, 181, 89, 3, 2, 12, 9, 2, 36, 36, 62, 62, 64, 64, 94,
	94, 96, 96, 98, 98, 125, 127, 4, 2, 71, 71, 103, 103, 4, 2, 45, 45, 47,
	47, 6, 2, 12, 12, 15, 15, 41, 41, 94, 94, 6, 2, 12, 12, 15, 15, 36, 36,
	94, 94, 4, 2, 41, 41, 94, 94, 9, 2, 36, 36, 41, 41, 100, 100, 104, 104,
	112, 112, 116, 116, 118, 118, 5, 2, 185, 185, 770, 881, 8257, 8258, 15,
	2, 67, 92, 99, 124, 194, 216, 218, 248, 250, 769, 882, 895, 897, 8193,
	8206, 8207, 8306, 8593, 11266, 12273, 12291, 55297, 63746, 64977, 65010,
	65535, 5, 2, 11, 12, 15, 15, 34, 34, 2, 781, 2, 3, 3, 2, 2, 2, 2, 5, 3,
	2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13,
	3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2,
	21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2,
	2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2,
	2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2,
	2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3,
	2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59,
	3, 2, 2, 2, 2, 61, 3, 2, 2, 2, 2, 63, 3, 2, 2, 2, 2, 65, 3, 2, 2, 2, 2,
	67, 3, 2, 2, 2, 2, 69, 3, 2, 2, 2, 2, 71, 3, 2, 2, 2, 2, 73, 3, 2, 2, 2,
	2, 75, 3, 2, 2, 2, 2, 77, 3, 2, 2, 2, 2, 79, 3, 2, 2, 2, 2, 81, 3, 2, 2,
	2, 2, 83, 3, 2, 2, 2, 2, 85, 3, 2, 2, 2, 2, 87, 3, 2, 2, 2, 2, 89, 3, 2,
	2, 2, 2, 91, 3, 2, 2, 2, 2, 93, 3, 2, 2, 2, 2, 95, 3, 2, 2, 2, 2, 97, 3,
	2, 2, 2, 2, 99, 3, 2, 2, 2, 2, 101, 3, 2, 2, 2, 2, 103, 3, 2, 2, 2, 2,
	105, 3, 2, 2, 2, 2, 107, 3, 2, 2, 2, 2, 109, 3, 2, 2, 2, 2, 111, 3, 2,
	2, 2, 2, 113, 3, 2, 2, 2, 2, 115, 3, 2, 2, 2, 2, 117, 3, 2, 2, 2, 2, 119,
	3, 2, 2, 2, 2, 121, 3, 2, 2, 2, 2, 123, 3, 2, 2, 2, 2, 125, 3, 2, 2, 2,
	2, 127, 3, 2, 2, 2, 2, 129, 3, 2, 2, 2, 2, 131, 3, 2, 2, 2, 2, 133, 3,
	2, 2, 2, 2, 135, 3, 2, 2, 2, 2, 137, 3, 2, 2, 2, 2, 139, 3, 2, 2, 2, 2,
	141, 3, 2, 2, 2, 2, 143, 3, 2, 2, 2, 2, 145, 3, 2, 2, 2, 2, 147, 3, 2,
	2, 2, 2, 149, 3, 2, 2, 2, 2, 151, 3, 2, 2, 2, 2, 153, 3, 2, 2, 2, 2, 155,
	3, 2, 2, 2, 2, 157, 3, 2, 2, 2, 2, 159, 3, 2, 2, 2, 2, 161, 3, 2, 2, 2,
	2, 163, 3, 2, 2, 2, 2, 165, 3, 2, 2, 2, 2, 167, 3, 2, 2, 2, 2, 169, 3,
	2, 2, 2, 2, 173, 3, 2, 2, 2, 2, 175, 3, 2, 2, 2, 2, 181, 3, 2, 2, 2, 3,
	183, 3, 2, 2, 2, 5, 188, 3, 2, 2, 2, 7, 195, 3, 2, 2, 2, 9, 202, 3, 2,
	2, 2, 11, 211, 3, 2, 2, 2, 13, 219, 3, 2, 2, 2, 15, 221, 3, 2, 2, 2, 17,
	231, 3, 2, 2, 2, 19, 240, 3, 2, 2, 2, 21, 244, 3, 2, 2, 2, 23, 249, 3,
	2, 2, 2, 25, 255, 3, 2, 2, 2, 27, 261, 3, 2, 2, 2, 29, 267, 3, 2, 2, 2,
	31, 270, 3, 2, 2, 2, 33, 274, 3, 2, 2, 2, 35, 279, 3, 2, 2, 2, 37, 285,
	3, 2, 2, 2, 39, 292, 3, 2, 2, 2, 41, 294, 3, 2, 2, 2, 43, 296, 3, 2, 2,
	2, 45, 298, 3, 2, 2, 2, 47, 307, 3, 2, 2, 2, 49, 313, 3, 2, 2, 2, 51, 319,
	3, 2, 2, 2, 53, 326, 3, 2, 2, 2, 55, 328, 3, 2, 2, 2, 57, 330, 3, 2, 2,
	2, 59, 332, 3, 2, 2, 2, 61, 334, 3, 2, 2, 2, 63, 336, 3, 2, 2, 2, 65, 338,
	3, 2, 2, 2, 67, 340, 3, 2, 2, 2, 69, 343, 3, 2, 2, 2, 71, 346, 3, 2, 2,
	2, 73, 348, 3, 2, 2, 2, 75, 351, 3, 2, 2, 2, 77, 353, 3, 2, 2, 2, 79, 355,
	3, 2, 2, 2, 81, 358, 3, 2, 2, 2, 83, 361, 3, 2, 2, 2, 85, 363, 3, 2, 2,
	2, 87, 365, 3, 2, 2, 2, 89, 367, 3, 2, 2, 2, 91, 369, 3, 2, 2, 2, 93, 373,
	3, 2, 2, 2, 95, 378, 3, 2, 2, 2, 97, 390, 3, 2, 2, 2, 99, 399, 3, 2, 2,
	2, 101, 405, 3, 2, 2, 2, 103, 414, 3, 2, 2, 2, 105, 420, 3, 2, 2, 2, 107,
	426, 3, 2, 2, 2, 109, 434, 3, 2, 2, 2, 111, 444, 3, 2, 2, 2, 113, 450,
	3, 2, 2, 2, 115, 453, 3, 2, 2, 2, 117, 458, 3, 2, 2, 2, 119, 464, 3, 2,
	2, 2, 121, 475, 3, 2, 2, 2, 123, 479, 3, 2, 2, 2, 125, 482, 3, 2, 2, 2,
	127, 487, 3, 2, 2, 2, 129, 490, 3, 2, 2, 2, 131, 493, 3, 2, 2, 2, 133,
	513, 3, 2, 2, 2, 135, 535, 3, 2, 2, 2, 137, 566, 3, 2, 2, 2, 139, 568,
	3, 2, 2, 2, 141, 571, 3, 2, 2, 2, 143, 574, 3, 2, 2, 2, 145, 577, 3, 2,
	2, 2, 147, 580, 3, 2, 2, 2, 149, 583, 3, 2, 2, 2, 151, 586, 3, 2, 2, 2,
	153, 595, 3, 2, 2, 2, 155, 605, 3, 2, 2, 2, 157, 615, 3, 2, 2, 2, 159,
	637, 3, 2, 2, 2, 161, 659, 3, 2, 2, 2, 163, 662, 3, 2, 2, 2, 165, 671,
	3, 2, 2, 2, 167, 682, 3, 2, 2, 2, 169, 686, 3, 2, 2, 2, 171, 699, 3, 2,
	2, 2, 173, 701, 3, 2, 2, 2, 175, 714, 3, 2, 2, 2, 177, 726, 3, 2, 2, 2,
	179, 728, 3, 2, 2, 2, 181, 731, 3, 2, 2, 2, 183, 184, 7, 68, 2, 2, 184,
	185, 7, 67, 2, 2, 185, 186, 7, 85, 2, 2, 186, 187, 7, 71, 2, 2, 187, 4,
	3, 2, 2, 2, 188, 189, 7, 82, 2, 2, 189, 190, 7, 84, 2, 2, 190, 191, 7,
	71, 2, 2, 191, 192, 7, 72, 2, 2, 192, 193, 7, 75, 2, 2, 193, 194, 7, 90,
	2, 2, 194, 6, 3, 2, 2, 2, 195, 196, 7, 85, 2, 2, 196, 197, 7, 71, 2, 2,
	197, 198, 7, 78, 2, 2, 198, 199, 7, 71, 2, 2, 199, 200, 7, 69, 2, 2, 200,
	201, 7, 86, 2, 2, 201, 8, 3, 2, 2, 2, 202, 203, 7, 70, 2, 2, 203, 204,
	7, 75, 2, 2, 204, 205, 7, 85, 2, 2, 205, 206, 7, 86, 2, 2, 206, 207, 7,
	75, 2, 2, 207, 208, 7, 80, 2, 2, 208, 209, 7, 69, 2, 2, 209, 210, 7, 86,
	2, 2, 210, 10, 3, 2, 2, 2, 211, 212, 7, 84, 2, 2, 212, 213, 7, 71, 2, 2,
	213, 214, 7, 70, 2, 2, 214, 215, 7, 87, 2, 2, 215, 216, 7, 69, 2, 2, 216,
	217, 7, 71, 2, 2, 217, 218, 7, 70, 2, 2, 218, 12, 3, 2, 2, 2, 219, 220,
	7, 44, 2, 2, 220, 14, 3, 2, 2, 2, 221, 222, 7, 69, 2, 2, 222, 223, 7, 81,
	2, 2, 223, 224, 7, 80, 2, 2, 224, 225, 7, 85, 2, 2, 225, 226, 7, 86, 2,
	2, 226, 227, 7, 84, 2, 2, 227, 228, 7, 87, 2, 2, 228, 229, 7, 69, 2, 2,
	229, 230, 7, 86, 2, 2, 230, 16, 3, 2, 2, 2, 231, 232, 7, 70, 2, 2, 232,
	233, 7, 71, 2, 2, 233, 234, 7, 85, 2, 2, 234, 235, 7, 69, 2, 2, 235, 236,
	7, 84, 2, 2, 236, 237, 7, 75, 2, 2, 237, 238, 7, 68, 2, 2, 238, 239, 7,
	71, 2, 2, 239, 18, 3, 2, 2, 2, 240, 241, 7, 67, 2, 2, 241, 242, 7, 85,
	2, 2, 242, 243, 7, 77, 2, 2, 243, 20, 3, 2, 2, 2, 244, 245, 7, 72, 2, 2,
	245, 246, 7, 84, 2, 2, 246, 247, 7, 81, 2, 2, 247, 248, 7, 79, 2, 2, 248,
	22, 3, 2, 2, 2, 249, 250, 7, 80, 2, 2, 250, 251, 7, 67, 2, 2, 251, 252,
	7, 79, 2, 2, 252, 253, 7, 71, 2, 2, 253, 254, 7, 70, 2, 2, 254, 24, 3,
	2, 2, 2, 255, 256, 7, 89, 2, 2, 256, 257, 7, 74, 2, 2, 257, 258, 7, 71,
	2, 2, 258, 259, 7, 84, 2, 2, 259, 260, 7, 71, 2, 2, 260, 26, 3, 2, 2, 2,
	261, 262, 7, 81, 2, 2, 262, 263, 7, 84, 2, 2, 263, 264, 7, 70, 2, 2, 264,
	265, 7, 71, 2, 2, 265, 266, 7, 84, 2, 2, 266, 28, 3, 2, 2, 2, 267, 268,
	7, 68, 2, 2, 268, 269, 7, 91, 2, 2, 269, 30, 3, 2, 2, 2, 270, 271, 7, 67,
	2, 2, 271, 272, 7, 85, 2, 2, 272, 273, 7, 69, 2, 2, 273, 32, 3, 2, 2, 2,
	274, 275, 7, 70, 2, 2, 275, 276, 7, 71, 2, 2, 276, 277, 7, 85, 2, 2, 277,
	278, 7, 69, 2, 2, 278, 34, 3, 2, 2, 2, 279, 280, 7, 78, 2, 2, 280, 281,
	7, 75, 2, 2, 281, 282, 7, 79, 2, 2, 282, 283, 7, 75, 2, 2, 283, 284, 7,
	86, 2, 2, 284, 36, 3, 2, 2, 2, 285, 286, 7, 81, 2, 2, 286, 287, 7, 72,
	2, 2, 287, 288, 7, 72, 2, 2, 288, 289, 7, 85, 2, 2, 289, 290, 7, 71, 2,
	2, 290, 291, 7, 86, 2, 2, 291, 38, 3, 2, 2, 2, 292, 293, 7, 125, 2, 2,
	293, 40, 3, 2, 2, 2, 294, 295, 7, 48, 2, 2, 295, 42, 3, 2, 2, 2, 296, 297,
	7, 127, 2, 2, 297, 44, 3, 2, 2, 2, 298, 299, 7, 81, 2, 2, 299, 300, 7,
	82, 2, 2, 300, 301, 7, 86, 2, 2, 301, 302, 7, 75, 2, 2, 302, 303, 7, 81,
	2, 2, 303, 304, 7, 80, 2, 2, 304, 305, 7, 67, 2, 2, 305, 306, 7, 78, 2,
	2, 306, 46, 3, 2, 2, 2, 307, 308, 7, 73, 2, 2, 308, 309, 7, 84, 2, 2, 309,
	310, 7, 67, 2, 2, 310, 311, 7, 82, 2, 2, 311, 312, 7, 74, 2, 2, 312, 48,
	3, 2, 2, 2, 313, 314, 7, 87, 2, 2, 314, 315, 7, 80, 2, 2, 315, 316, 7,
	75, 2, 2, 316, 317, 7, 81, 2, 2, 317, 318, 7, 80, 2, 2, 318, 50, 3, 2,
	2, 2, 319, 320, 7, 72, 2, 2, 320, 321, 7, 75, 2, 2, 321, 322, 7, 78, 2,
	2, 322, 323, 7, 86, 2, 2, 323, 324, 7, 71, 2, 2, 324, 325, 7, 84, 2, 2,
	325, 52, 3, 2, 2, 2, 326, 327, 7, 42, 2, 2, 327, 54, 3, 2, 2, 2, 328, 329,
	7, 46, 2, 2, 329, 56, 3, 2, 2, 2, 330, 331, 7, 43, 2, 2, 331, 58, 3, 2,
	2, 2, 332, 333, 7, 61, 2, 2, 333, 60, 3, 2, 2, 2, 334, 335, 7, 67, 2, 2,
	335, 62, 3, 2, 2, 2, 336, 337, 7, 93, 2, 2, 337, 64, 3, 2, 2, 2, 338, 339,
	7, 95, 2, 2, 339, 66, 3, 2, 2, 2, 340, 341, 7, 126, 2, 2, 341, 342, 7,
	126, 2, 2, 342, 68, 3, 2, 2, 2, 343, 344, 7, 40, 2, 2, 344, 345, 7, 40,
	2, 2, 345, 70, 3, 2, 2, 2, 346, 347, 7, 63, 2, 2, 347, 72, 3, 2, 2, 2,
	348, 349, 7, 35, 2, 2, 349, 350, 7, 63, 2, 2, 350, 74, 3, 2, 2, 2, 351,
	352, 7, 62, 2, 2, 352, 76, 3, 2, 2, 2, 353, 354, 7, 64, 2, 2, 354, 78,
	3, 2, 2, 2, 355, 356, 7, 62, 2, 2, 356, 357, 7, 63, 2, 2, 357, 80, 3, 2,
	2, 2, 358, 359, 7, 64, 2, 2, 359, 360, 7, 63, 2, 2, 360, 82, 3, 2, 2, 2,
	361, 362, 7, 45, 2, 2, 362, 84, 3, 2, 2, 2, 363, 364, 7, 47, 2, 2, 364,
	86, 3, 2, 2, 2, 365, 366, 7, 49, 2, 2, 366, 88, 3, 2, 2, 2, 367, 368, 7,
	35, 2, 2, 368, 90, 3, 2, 2, 2, 369, 370, 7, 85, 2, 2, 370, 371, 7, 86,
	2, 2, 371, 372, 7, 84, 2, 2, 372, 92, 3, 2, 2, 2, 373, 374, 7, 78, 2, 2,
	374, 375, 7, 67, 2, 2, 375, 376, 7, 80, 2, 2, 376, 377, 7, 73, 2, 2, 377,
	94, 3, 2, 2, 2, 378, 379, 7, 78, 2, 2, 379, 380, 7, 67, 2, 2, 380, 381,
	7, 80, 2, 2, 381, 382, 7, 73, 2, 2, 382, 383, 7, 79, 2, 2, 383, 384, 7,
	67, 2, 2, 384, 385, 7, 86, 2, 2, 385, 386, 7, 69, 2, 2, 386, 387, 7, 74,
	2, 2, 387, 388, 7, 71, 2, 2, 388, 389, 7, 85, 2, 2, 389, 96, 3, 2, 2, 2,
	390, 391, 7, 70, 2, 2, 391, 392, 7, 67, 2, 2, 392, 393, 7, 86, 2, 2, 393,
	394, 7, 67, 2, 2, 394, 395, 7, 86, 2, 2, 395, 396, 7, 91, 2, 2, 396, 397,
	7, 82, 2, 2, 397, 398, 7, 71, 2, 2, 398, 98, 3, 2, 2, 2, 399, 400, 7, 68,
	2, 2, 400, 401, 7, 81, 2, 2, 401, 402, 7, 87, 2, 2, 402, 403, 7, 80, 2,
	2, 403, 404, 7, 70, 2, 2, 404, 100, 3, 2, 2, 2, 405, 406, 7, 85, 2, 2,
	406, 407, 7, 67, 2, 2, 407, 408, 7, 79, 2, 2, 408, 409, 7, 71, 2, 2, 409,
	410, 7, 86, 2, 2, 410, 411, 7, 71, 2, 2, 411, 412, 7, 84, 2, 2, 412, 413,
	7, 79, 2, 2, 413, 102, 3, 2, 2, 2, 414, 415, 7, 75, 2, 2, 415, 416, 7,
	85, 2, 2, 416, 417, 7, 75, 2, 2, 417, 418, 7, 84, 2, 2, 418, 419, 7, 75,
	2, 2, 419, 104, 3, 2, 2, 2, 420, 421, 7, 75, 2, 2, 421, 422, 7, 85, 2,
	2, 422, 423, 7, 87, 2, 2, 423, 424, 7, 84, 2, 2, 424, 425, 7, 75, 2, 2,
	425, 106, 3, 2, 2, 2, 426, 427, 7, 75, 2, 2, 427, 428, 7, 85, 2, 2, 428,
	429, 7, 68, 2, 2, 429, 430, 7, 78, 2, 2, 430, 431, 7, 67, 2, 2, 431, 432,
	7, 80, 2, 2, 432, 433, 7, 77, 2, 2, 433, 108, 3, 2, 2, 2, 434, 435, 7,
	75, 2, 2, 435, 436, 7, 85, 2, 2, 436, 437, 7, 78, 2, 2, 437, 438, 7, 75,
	2, 2, 438, 439, 7, 86, 2, 2, 439, 440, 7, 71, 2, 2, 440, 441, 7, 84, 2,
	2, 441, 442, 7, 67, 2, 2, 442, 443, 7, 78, 2, 2, 443, 110, 3, 2, 2, 2,
	444, 445, 7, 84, 2, 2, 445, 446, 7, 71, 2, 2, 446, 447, 7, 73, 2, 2, 447,
	448, 7, 71, 2, 2, 448, 449, 7, 90, 2, 2, 449, 112, 3, 2, 2, 2, 450, 451,
	7, 96, 2, 2, 451, 452, 7, 96, 2, 2, 452, 114, 3, 2, 2, 2, 453, 454, 7,
	118, 2, 2, 454, 455, 7, 116, 2, 2, 455, 456, 7, 119, 2, 2, 456, 457, 7,
	103, 2, 2, 457, 116, 3, 2, 2, 2, 458, 459, 7, 104, 2, 2, 459, 460, 7, 99,
	2, 2, 460, 461, 7, 110, 2, 2, 461, 462, 7, 117, 2, 2, 462, 463, 7, 103,
	2, 2, 463, 118, 3, 2, 2, 2, 464, 469, 7, 62, 2, 2, 465, 468, 10, 2, 2,
	2, 466, 468, 5, 171, 86, 2, 467, 465, 3, 2, 2, 2, 467, 466, 3, 2, 2, 2,
	468, 471, 3, 2, 2, 2, 469, 467, 3, 2, 2, 2, 469, 470, 3, 2, 2, 2, 470,
	472, 3, 2, 2, 2, 471, 469, 3, 2, 2, 2, 472, 473, 7, 64, 2, 2, 473, 120,
	3, 2, 2, 2, 474, 476, 5, 173, 87, 2, 475, 474, 3, 2, 2, 2, 475, 476, 3,
	2, 2, 2, 476, 477, 3, 2, 2, 2, 477, 478, 7, 60, 2, 2, 478, 122, 3, 2, 2,
	2, 479, 480, 5, 121, 61, 2, 480, 481, 5, 175, 88, 2, 481, 124, 3, 2, 2,
	2, 482, 483, 7, 97, 2, 2, 483, 484, 7, 60, 2, 2, 484, 485, 3, 2, 2, 2,
	485, 486, 5, 175, 88, 2, 486, 126, 3, 2, 2, 2, 487, 488, 7, 65, 2, 2, 488,
	489, 5, 169, 85, 2, 489, 128, 3, 2, 2, 2, 490, 491, 7, 38, 2, 2, 491, 492,
	5, 169, 85, 2, 492, 130, 3, 2, 2, 2, 493, 495, 7, 66, 2, 2, 494, 496, 5,
	177, 89, 2, 495, 494, 3, 2, 2, 2, 496, 497, 3, 2, 2, 2, 497, 495, 3, 2,
	2, 2, 497, 498, 3, 2, 2, 2, 498, 509, 3, 2, 2, 2, 499, 503, 7, 47, 2, 2,
	500, 501, 5, 177, 89, 2, 501, 502, 5, 179, 90, 2, 502, 504, 3, 2, 2, 2,
	503, 500, 3, 2, 2, 2, 504, 505, 3, 2, 2, 2, 505, 503, 3, 2, 2, 2, 505,
	506, 3, 2, 2, 2, 506, 508, 3, 2, 2, 2, 507, 499, 3, 2, 2, 2, 508, 511,
	3, 2, 2, 2, 509, 507, 3, 2, 2, 2, 509, 510, 3, 2, 2, 2, 510, 132, 3, 2,
	2, 2, 511, 509, 3, 2, 2, 2, 512, 514, 5, 179, 90, 2, 513, 512, 3, 2, 2,
	2, 514, 515, 3, 2, 2, 2, 515, 513, 3, 2, 2, 2, 515, 516, 3, 2, 2, 2, 516,
	134, 3, 2, 2, 2, 517, 519, 5, 179, 90, 2, 518, 517, 3, 2, 2, 2, 519, 520,
	3, 2, 2, 2, 520, 518, 3, 2, 2, 2, 520, 521, 3, 2, 2, 2, 521, 522, 3, 2,
	2, 2, 522, 526, 7, 48, 2, 2, 523, 525, 5, 179, 90, 2, 524, 523, 3, 2, 2,
	2, 525, 528, 3, 2, 2, 2, 526, 524, 3, 2, 2, 2, 526, 527, 3, 2, 2, 2, 527,
	536, 3, 2, 2, 2, 528, 526, 3, 2, 2, 2, 529, 531, 7, 48, 2, 2, 530, 532,
	5, 179, 90, 2, 531, 530, 3, 2, 2, 2, 532, 533, 3, 2, 2, 2, 533, 531, 3,
	2, 2, 2, 533, 534, 3, 2, 2, 2, 534, 536, 3, 2, 2, 2, 535, 518, 3, 2, 2,
	2, 535, 529, 3, 2, 2, 2, 536, 136, 3, 2, 2, 2, 537, 539, 5, 179, 90, 2,
	538, 537, 3, 2, 2, 2, 539, 540, 3, 2, 2, 2, 540, 538, 3, 2, 2, 2, 540,
	541, 3, 2, 2, 2, 541, 542, 3, 2, 2, 2, 542, 546, 7, 48, 2, 2, 543, 545,
	5, 179, 90, 2, 544, 543, 3, 2, 2, 2, 545, 548, 3, 2, 2, 2, 546, 544, 3,
	2, 2, 2, 546, 547, 3, 2, 2, 2, 547, 549, 3, 2, 2, 2, 548, 546, 3, 2, 2,
	2, 549, 550, 5, 151, 76, 2, 550, 567, 3, 2, 2, 2, 551, 553, 7, 48, 2, 2,
	552, 554, 5, 179, 90, 2, 553, 552, 3, 2, 2, 2, 554, 555, 3, 2, 2, 2, 555,
	553, 3, 2, 2, 2, 555, 556, 3, 2, 2, 2, 556, 557, 3, 2, 2, 2, 557, 558,
	5, 151, 76, 2, 558, 567, 3, 2, 2, 2, 559, 561, 5, 179, 90, 2, 560, 559,
	3, 2, 2, 2, 561, 562, 3, 2, 2, 2, 562, 560, 3, 2, 2, 2, 562, 563, 3, 2,
	2, 2, 563, 564, 3, 2, 2, 2, 564, 565, 5, 151, 76, 2, 565, 567, 3, 2, 2,
	2, 566, 538, 3, 2, 2, 2, 566, 551, 3, 2, 2, 2, 566, 560, 3, 2, 2, 2, 567,
	138, 3, 2, 2, 2, 568, 569, 7, 45, 2, 2, 569, 570, 5, 133, 67, 2, 570, 140,
	3, 2, 2, 2, 571, 572, 7, 45, 2, 2, 572, 573, 5, 135, 68, 2, 573, 142, 3,
	2, 2, 2, 574, 575, 7, 45, 2, 2, 575, 576, 5, 137, 69, 2, 576, 144, 3, 2,
	2, 2, 577, 578, 7, 47, 2, 2, 578, 579, 5, 133, 67, 2, 579, 146, 3, 2, 2,
	2, 580, 581, 7, 47, 2, 2, 581, 582, 5, 135, 68, 2, 582, 148, 3, 2, 2, 2,
	583, 584, 7, 47, 2, 2, 584, 585, 5, 137, 69, 2, 585, 150, 3, 2, 2, 2, 586,
	588, 9, 3, 2, 2, 587, 589, 9, 4, 2, 2, 588, 587, 3, 2, 2, 2, 588, 589,
	3, 2, 2, 2, 589, 591, 3, 2, 2, 2, 590, 592, 5, 179, 90, 2, 591, 590, 3,
	2, 2, 2, 592, 593, 3, 2, 2, 2, 593, 591, 3, 2, 2, 2, 593, 594, 3, 2, 2,
	2, 594, 152, 3, 2, 2, 2, 595, 600, 7, 41, 2, 2, 596, 599, 10, 5, 2, 2,
	597, 599, 5, 161, 81, 2, 598, 596, 3, 2, 2, 2, 598, 597, 3, 2, 2, 2, 599,
	602, 3, 2, 2, 2, 600, 598, 3, 2, 2, 2, 600, 601, 3, 2, 2, 2, 601, 603,
	3, 2, 2, 2, 602, 600, 3, 2, 2, 2, 603, 604, 7, 41, 2, 2, 604, 154, 3, 2,
	2, 2, 605, 610, 7, 36, 2, 2, 606, 609, 10, 6, 2, 2, 607, 609, 5, 161, 81,
	2, 608, 606, 3, 2, 2, 2, 608, 607, 3, 2, 2, 2, 609, 612, 3, 2, 2, 2, 610,
	608, 3, 2, 2, 2, 610, 611, 3, 2, 2, 2, 611, 613, 3, 2, 2, 2, 612, 610,
	3, 2, 2, 2, 613, 614, 7, 36, 2, 2, 614, 156, 3, 2, 2, 2, 615, 616, 7, 41,
	2, 2, 616, 617, 7, 41, 2, 2, 617, 618, 7, 41, 2, 2, 618, 630, 3, 2, 2,
	2, 619, 623, 7, 41, 2, 2, 620, 621, 7, 41, 2, 2, 621, 623, 7, 41, 2, 2,
	622, 619, 3, 2, 2, 2, 622, 620, 3, 2, 2, 2, 622, 623, 3, 2, 2, 2, 623,
	626, 3, 2, 2, 2, 624, 627, 10, 7, 2, 2, 625, 627, 5, 161, 81, 2, 626, 624,
	3, 2, 2, 2, 626, 625, 3, 2, 2, 2, 627, 629, 3, 2, 2, 2, 628, 622, 3, 2,
	2, 2, 629, 632, 3, 2, 2, 2, 630, 628, 3, 2, 2, 2, 630, 631, 3, 2, 2, 2,
	631, 633, 3, 2, 2, 2, 632, 630, 3, 2, 2, 2, 633, 634, 7, 41, 2, 2, 634,
	635, 7, 41, 2, 2, 635, 636, 7, 41, 2, 2, 636, 158, 3, 2, 2, 2, 637, 638,
	7, 36, 2, 2, 638, 639, 7, 36, 2, 2, 639, 640, 7, 36, 2, 2, 640, 652, 3,
	2, 2, 2, 641, 645, 7, 36, 2, 2, 642, 643, 7, 36, 2, 2, 643, 645, 7, 36,
	2, 2, 644, 641, 3, 2, 2, 2, 644, 642, 3, 2, 2, 2, 644, 645, 3, 2, 2, 2,
	645, 648, 3, 2, 2, 2, 646, 649, 10, 7, 2, 2, 647, 649, 5, 161, 81, 2, 648,
	646, 3, 2, 2, 2, 648, 647, 3, 2, 2, 2, 649, 651, 3, 2, 2, 2, 650, 644,
	3, 2, 2, 2, 651, 654, 3, 2, 2, 2, 652, 650, 3, 2, 2, 2, 652, 653, 3, 2,
	2, 2, 653, 655, 3, 2, 2, 2, 654, 652, 3, 2, 2, 2, 655, 656, 7, 36, 2, 2,
	656, 657, 7, 36, 2, 2, 657, 658, 7, 36, 2, 2, 658, 160, 3, 2, 2, 2, 659,
	660, 7, 94, 2, 2, 660, 661, 9, 8, 2, 2, 661, 162, 3, 2, 2, 2, 662, 666,
	7, 42, 2, 2, 663, 665, 5, 181, 91, 2, 664, 663, 3, 2, 2, 2, 665, 668, 3,
	2, 2, 2, 666, 664, 3, 2, 2, 2, 666, 667, 3, 2, 2, 2, 667, 669, 3, 2, 2,
	2, 668, 666, 3, 2, 2, 2, 669, 670, 7, 43, 2, 2, 670, 164, 3, 2, 2, 2, 671,
	675, 7, 93, 2, 2, 672, 674, 5, 181, 91, 2, 673, 672, 3, 2, 2, 2, 674, 677,
	3, 2, 2, 2, 675, 673, 3, 2, 2, 2, 675, 676, 3, 2, 2, 2, 676, 678, 3, 2,
	2, 2, 677, 675, 3, 2, 2, 2, 678, 679, 7, 95, 2, 2, 679, 166, 3, 2, 2, 2,
	680, 683, 5, 177, 89, 2, 681, 683, 7, 97, 2, 2, 682, 680, 3, 2, 2, 2, 682,
	681, 3, 2, 2, 2, 683, 168, 3, 2, 2, 2, 684, 687, 5, 167, 84, 2, 685, 687,
	5, 179, 90, 2, 686, 684, 3, 2, 2, 2, 686, 685, 3, 2, 2, 2, 687, 693, 3,
	2, 2, 2, 688, 692, 5, 167, 84, 2, 689, 692, 5, 179, 90, 2, 690, 692, 9,
	9, 2, 2, 691, 688, 3, 2, 2, 2, 691, 689, 3, 2, 2, 2, 691, 690, 3, 2, 2,
	2, 692, 695, 3, 2, 2, 2, 693, 691, 3, 2, 2, 2, 693, 694, 3, 2, 2, 2, 694,
	170, 3, 2, 2, 2, 695, 693, 3, 2, 2, 2, 696, 700, 5, 167, 84, 2, 697, 700,
	7, 47, 2, 2, 698, 700, 5, 179, 90, 2, 699, 696, 3, 2, 2, 2, 699, 697, 3,
	2, 2, 2, 699, 698, 3, 2, 2, 2, 700, 172, 3, 2, 2, 2, 701, 710, 5, 177,
	89, 2, 702, 705, 5, 171, 86, 2, 703, 705, 7, 48, 2, 2, 704, 702, 3, 2,
	2, 2, 704, 703, 3, 2, 2, 2, 705, 708, 3, 2, 2, 2, 706, 704, 3, 2, 2, 2,
	706, 707, 3, 2, 2, 2, 707, 709, 3, 2, 2, 2, 708, 706, 3, 2, 2, 2, 709,
	711, 5, 171, 86, 2, 710, 706, 3, 2, 2, 2, 710, 711, 3, 2, 2, 2, 711, 174,
	3, 2, 2, 2, 712, 715, 5, 167, 84, 2, 713, 715, 5, 179, 90, 2, 714, 712,
	3, 2, 2, 2, 714, 713, 3, 2, 2, 2, 715, 724, 3, 2, 2, 2, 716, 719, 5, 171,
	86, 2, 717, 719, 7, 48, 2, 2, 718, 716, 3, 2, 2, 2, 718, 717, 3, 2, 2,
	2, 719, 722, 3, 2, 2, 2, 720, 718, 3, 2, 2, 2, 720, 721, 3, 2, 2, 2, 721,
	723, 3, 2, 2, 2, 722, 720, 3, 2, 2, 2, 723, 725, 5, 171, 86, 2, 724, 720,
	3, 2, 2, 2, 724, 725, 3, 2, 2, 2, 725, 176, 3, 2, 2, 2, 726, 727, 9, 10,
	2, 2, 727, 178, 3, 2, 2, 2, 728, 729, 4, 50, 59, 2, 729, 180, 3, 2, 2,
	2, 730, 732, 9, 11, 2, 2, 731, 730, 3, 2, 2, 2, 732, 733, 3, 2, 2, 2, 733,
	731, 3, 2, 2, 2, 733, 734, 3, 2, 2, 2, 734, 735, 3, 2, 2, 2, 735, 736,
	8, 91, 2, 2, 736, 182, 3, 2, 2, 2, 46, 2, 467, 469, 475, 497, 505, 509,
	515, 520, 526, 533, 535, 540, 546, 555, 562, 566, 588, 593, 598, 600, 608,
	610, 622, 626, 630, 644, 648, 652, 666, 675, 682, 686, 691, 693, 699, 704,
	706, 710, 714, 718, 720, 724, 733, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'BASE'", "'PREFIX'", "'SELECT'", "'DISTINCT'", "'REDUCED'", "'*'",
	"'CONSTRUCT'", "'DESCRIBE'", "'ASK'", "'FROM'", "'NAMED'", "'WHERE'", "'ORDER'",
	"'BY'", "'ASC'", "'DESC'", "'LIMIT'", "'OFFSET'", "'{'", "'.'", "'}'",
	"'OPTIONAL'", "'GRAPH'", "'UNION'", "'FILTER'", "'('", "','", "')'", "';'",
	"'A'", "'['", "']'", "'||'", "'&&'", "'='", "'!='", "'<'", "'>'", "'<='",
	"'>='", "'+'", "'-'", "'/'", "'!'", "'STR'", "'LANG'", "'LANGMATCHES'",
	"'DATATYPE'", "'BOUND'", "'SAMETERM'", "'ISIRI'", "'ISURI'", "'ISBLANK'",
	"'ISLITERAL'", "'REGEX'", "'^^'", "'true'", "'false'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "IRI_REF", "PNAME_NS", "PNAME_LN", "BLANK_NODE_LABEL",
	"VAR1", "VAR2", "LANGTAG", "INTEGER", "DECIMAL", "DOUBLE", "INTEGER_POSITIVE",
	"DECIMAL_POSITIVE", "DOUBLE_POSITIVE", "INTEGER_NEGATIVE", "DECIMAL_NEGATIVE",
	"DOUBLE_NEGATIVE", "EXPONENT", "STRING_LITERAL1", "STRING_LITERAL2", "STRING_LITERAL_LONG1",
	"STRING_LITERAL_LONG2", "ECHAR", "NIL", "ANON", "PN_CHARS_U", "VARNAME",
	"PN_PREFIX", "PN_LOCAL", "WS",
}

var lexerRuleNames = []string{
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
	"DOUBLE_NEGATIVE", "EXPONENT", "STRING_LITERAL1", "STRING_LITERAL2", "STRING_LITERAL_LONG1",
	"STRING_LITERAL_LONG2", "ECHAR", "NIL", "ANON", "PN_CHARS_U", "VARNAME",
	"PN_CHARS", "PN_PREFIX", "PN_LOCAL", "PN_CHARS_BASE", "DIGIT", "WS",
}

type SparqlLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewSparqlLexer(input antlr.CharStream) *SparqlLexer {

	l := new(SparqlLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
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
