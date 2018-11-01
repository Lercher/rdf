// Code generated from /home/lercher/go/src/github.com/lercher/rdf/sparql/Sparql.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Sparql

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 89, 625,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 9, 44,
	4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9, 49, 4,
	50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54, 4, 55,
	9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4, 58, 9, 58, 4, 59, 9, 59, 4, 60, 9,
	60, 4, 61, 9, 61, 4, 62, 9, 62, 4, 63, 9, 63, 4, 64, 9, 64, 4, 65, 9, 65,
	4, 66, 9, 66, 4, 67, 9, 67, 4, 68, 9, 68, 4, 69, 9, 69, 4, 70, 9, 70, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 5, 2, 146, 10, 2, 3, 2, 3, 2, 3, 3, 5, 3, 151,
	10, 3, 3, 3, 7, 3, 154, 10, 3, 12, 3, 14, 3, 157, 11, 3, 3, 4, 3, 4, 3,
	4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 5, 6, 168, 10, 6, 3, 6, 6, 6, 171,
	10, 6, 13, 6, 14, 6, 172, 3, 6, 5, 6, 176, 10, 6, 3, 6, 7, 6, 179, 10,
	6, 12, 6, 14, 6, 182, 11, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 7, 7,
	190, 10, 7, 12, 7, 14, 7, 193, 11, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 6,
	8, 200, 10, 8, 13, 8, 14, 8, 201, 3, 8, 5, 8, 205, 10, 8, 3, 8, 7, 8, 208,
	10, 8, 12, 8, 14, 8, 211, 11, 8, 3, 8, 5, 8, 214, 10, 8, 3, 8, 3, 8, 3,
	9, 3, 9, 7, 9, 220, 10, 9, 12, 9, 14, 9, 223, 11, 9, 3, 9, 3, 9, 3, 10,
	3, 10, 3, 10, 5, 10, 230, 10, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3,
	13, 3, 13, 3, 14, 5, 14, 240, 10, 14, 3, 14, 3, 14, 3, 15, 5, 15, 245,
	10, 15, 3, 15, 5, 15, 248, 10, 15, 3, 16, 3, 16, 5, 16, 252, 10, 16, 3,
	16, 3, 16, 5, 16, 256, 10, 16, 5, 16, 258, 10, 16, 3, 17, 3, 17, 3, 17,
	6, 17, 263, 10, 17, 13, 17, 14, 17, 264, 3, 18, 3, 18, 3, 18, 3, 18, 5,
	18, 271, 10, 18, 5, 18, 273, 10, 18, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20,
	3, 20, 3, 21, 3, 21, 5, 21, 283, 10, 21, 3, 21, 3, 21, 5, 21, 287, 10,
	21, 3, 21, 5, 21, 290, 10, 21, 3, 21, 5, 21, 293, 10, 21, 7, 21, 295, 10,
	21, 12, 21, 14, 21, 298, 11, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 5,
	22, 305, 10, 22, 5, 22, 307, 10, 22, 3, 23, 3, 23, 3, 23, 5, 23, 312, 10,
	23, 3, 24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26,
	7, 26, 324, 10, 26, 12, 26, 14, 26, 327, 11, 26, 3, 27, 3, 27, 3, 27, 3,
	28, 3, 28, 3, 28, 5, 28, 335, 10, 28, 3, 29, 3, 29, 3, 29, 3, 30, 3, 30,
	3, 30, 3, 30, 3, 30, 7, 30, 345, 10, 30, 12, 30, 14, 30, 348, 11, 30, 3,
	30, 3, 30, 5, 30, 352, 10, 30, 3, 31, 3, 31, 5, 31, 356, 10, 31, 3, 31,
	3, 31, 3, 32, 3, 32, 3, 32, 5, 32, 363, 10, 32, 5, 32, 365, 10, 32, 3,
	33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 5, 33, 373, 10, 33, 3, 34, 3, 34,
	3, 34, 3, 34, 3, 34, 3, 34, 5, 34, 381, 10, 34, 7, 34, 383, 10, 34, 12,
	34, 14, 34, 386, 11, 34, 3, 35, 5, 35, 389, 10, 35, 3, 36, 3, 36, 3, 36,
	7, 36, 394, 10, 36, 12, 36, 14, 36, 397, 11, 36, 3, 37, 3, 37, 3, 38, 3,
	38, 5, 38, 403, 10, 38, 3, 39, 3, 39, 5, 39, 407, 10, 39, 3, 40, 3, 40,
	3, 40, 3, 40, 3, 41, 3, 41, 6, 41, 415, 10, 41, 13, 41, 14, 41, 416, 3,
	41, 3, 41, 3, 42, 3, 42, 5, 42, 423, 10, 42, 3, 43, 3, 43, 5, 43, 427,
	10, 43, 3, 44, 3, 44, 5, 44, 431, 10, 44, 3, 45, 3, 45, 3, 46, 3, 46, 3,
	46, 3, 46, 3, 46, 3, 46, 5, 46, 441, 10, 46, 3, 47, 3, 47, 3, 48, 3, 48,
	3, 48, 7, 48, 448, 10, 48, 12, 48, 14, 48, 451, 11, 48, 3, 49, 3, 49, 3,
	49, 7, 49, 456, 10, 49, 12, 49, 14, 49, 459, 11, 49, 3, 50, 3, 50, 3, 51,
	3, 51, 3, 51, 3, 51, 3, 51, 3, 51, 3, 51, 3, 51, 3, 51, 3, 51, 3, 51, 3,
	51, 3, 51, 5, 51, 476, 10, 51, 3, 52, 3, 52, 3, 53, 3, 53, 3, 53, 3, 53,
	3, 53, 3, 53, 3, 53, 7, 53, 487, 10, 53, 12, 53, 14, 53, 490, 11, 53, 3,
	54, 3, 54, 3, 54, 3, 54, 3, 54, 7, 54, 497, 10, 54, 12, 54, 14, 54, 500,
	11, 54, 3, 55, 3, 55, 3, 55, 3, 55, 3, 55, 3, 55, 3, 55, 5, 55, 509, 10,
	55, 3, 56, 3, 56, 3, 56, 3, 56, 3, 56, 3, 56, 3, 56, 5, 56, 518, 10, 56,
	3, 57, 3, 57, 3, 57, 3, 57, 3, 58, 3, 58, 3, 58, 3, 58, 3, 58, 3, 58, 3,
	58, 3, 58, 3, 58, 3, 58, 3, 58, 3, 58, 3, 58, 3, 58, 3, 58, 3, 58, 3, 58,
	3, 58, 3, 58, 3, 58, 3, 58, 3, 58, 3, 58, 3, 58, 3, 58, 3, 58, 3, 58, 3,
	58, 3, 58, 3, 58, 3, 58, 3, 58, 3, 58, 3, 58, 3, 58, 3, 58, 3, 58, 3, 58,
	3, 58, 3, 58, 3, 58, 3, 58, 3, 58, 3, 58, 3, 58, 3, 58, 3, 58, 3, 58, 3,
	58, 3, 58, 3, 58, 3, 58, 3, 58, 3, 58, 3, 58, 5, 58, 579, 10, 58, 3, 59,
	3, 59, 3, 59, 3, 59, 3, 59, 3, 59, 3, 59, 5, 59, 588, 10, 59, 3, 59, 3,
	59, 3, 60, 3, 60, 5, 60, 594, 10, 60, 3, 61, 3, 61, 3, 61, 3, 61, 5, 61,
	600, 10, 61, 3, 62, 3, 62, 3, 62, 5, 62, 605, 10, 62, 3, 63, 3, 63, 3,
	64, 3, 64, 3, 65, 3, 65, 3, 66, 3, 66, 3, 67, 3, 67, 3, 68, 3, 68, 5, 68,
	619, 10, 68, 3, 69, 3, 69, 3, 70, 3, 70, 3, 70, 2, 2, 71, 2, 4, 6, 8, 10,
	12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46,
	48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82,
	84, 86, 88, 90, 92, 94, 96, 98, 100, 102, 104, 106, 108, 110, 112, 114,
	116, 118, 120, 122, 124, 126, 128, 130, 132, 134, 136, 138, 2, 12, 3, 2,
	6, 7, 3, 2, 17, 18, 3, 2, 65, 66, 3, 2, 68, 70, 3, 2, 71, 73, 3, 2, 74,
	76, 3, 2, 59, 60, 3, 2, 78, 79, 3, 2, 62, 63, 4, 2, 64, 64, 84, 84, 2,
	653, 2, 140, 3, 2, 2, 2, 4, 150, 3, 2, 2, 2, 6, 158, 3, 2, 2, 2, 8, 161,
	3, 2, 2, 2, 10, 165, 3, 2, 2, 2, 12, 186, 3, 2, 2, 2, 14, 197, 3, 2, 2,
	2, 16, 217, 3, 2, 2, 2, 18, 226, 3, 2, 2, 2, 20, 231, 3, 2, 2, 2, 22, 233,
	3, 2, 2, 2, 24, 236, 3, 2, 2, 2, 26, 239, 3, 2, 2, 2, 28, 244, 3, 2, 2,
	2, 30, 257, 3, 2, 2, 2, 32, 259, 3, 2, 2, 2, 34, 272, 3, 2, 2, 2, 36, 274,
	3, 2, 2, 2, 38, 277, 3, 2, 2, 2, 40, 280, 3, 2, 2, 2, 42, 301, 3, 2, 2,
	2, 44, 311, 3, 2, 2, 2, 46, 313, 3, 2, 2, 2, 48, 316, 3, 2, 2, 2, 50, 320,
	3, 2, 2, 2, 52, 328, 3, 2, 2, 2, 54, 334, 3, 2, 2, 2, 56, 336, 3, 2, 2,
	2, 58, 351, 3, 2, 2, 2, 60, 353, 3, 2, 2, 2, 62, 359, 3, 2, 2, 2, 64, 372,
	3, 2, 2, 2, 66, 374, 3, 2, 2, 2, 68, 388, 3, 2, 2, 2, 70, 390, 3, 2, 2,
	2, 72, 398, 3, 2, 2, 2, 74, 402, 3, 2, 2, 2, 76, 406, 3, 2, 2, 2, 78, 408,
	3, 2, 2, 2, 80, 412, 3, 2, 2, 2, 82, 422, 3, 2, 2, 2, 84, 426, 3, 2, 2,
	2, 86, 430, 3, 2, 2, 2, 88, 432, 3, 2, 2, 2, 90, 440, 3, 2, 2, 2, 92, 442,
	3, 2, 2, 2, 94, 444, 3, 2, 2, 2, 96, 452, 3, 2, 2, 2, 98, 460, 3, 2, 2,
	2, 100, 462, 3, 2, 2, 2, 102, 477, 3, 2, 2, 2, 104, 479, 3, 2, 2, 2, 106,
	491, 3, 2, 2, 2, 108, 508, 3, 2, 2, 2, 110, 517, 3, 2, 2, 2, 112, 519,
	3, 2, 2, 2, 114, 578, 3, 2, 2, 2, 116, 580, 3, 2, 2, 2, 118, 591, 3, 2,
	2, 2, 120, 595, 3, 2, 2, 2, 122, 604, 3, 2, 2, 2, 124, 606, 3, 2, 2, 2,
	126, 608, 3, 2, 2, 2, 128, 610, 3, 2, 2, 2, 130, 612, 3, 2, 2, 2, 132,
	614, 3, 2, 2, 2, 134, 618, 3, 2, 2, 2, 136, 620, 3, 2, 2, 2, 138, 622,
	3, 2, 2, 2, 140, 145, 5, 4, 3, 2, 141, 146, 5, 10, 6, 2, 142, 146, 5, 12,
	7, 2, 143, 146, 5, 14, 8, 2, 144, 146, 5, 16, 9, 2, 145, 141, 3, 2, 2,
	2, 145, 142, 3, 2, 2, 2, 145, 143, 3, 2, 2, 2, 145, 144, 3, 2, 2, 2, 146,
	147, 3, 2, 2, 2, 147, 148, 7, 2, 2, 3, 148, 3, 3, 2, 2, 2, 149, 151, 5,
	6, 4, 2, 150, 149, 3, 2, 2, 2, 150, 151, 3, 2, 2, 2, 151, 155, 3, 2, 2,
	2, 152, 154, 5, 8, 5, 2, 153, 152, 3, 2, 2, 2, 154, 157, 3, 2, 2, 2, 155,
	153, 3, 2, 2, 2, 155, 156, 3, 2, 2, 2, 156, 5, 3, 2, 2, 2, 157, 155, 3,
	2, 2, 2, 158, 159, 7, 3, 2, 2, 159, 160, 7, 61, 2, 2, 160, 7, 3, 2, 2,
	2, 161, 162, 7, 4, 2, 2, 162, 163, 7, 62, 2, 2, 163, 164, 7, 61, 2, 2,
	164, 9, 3, 2, 2, 2, 165, 167, 7, 5, 2, 2, 166, 168, 9, 2, 2, 2, 167, 166,
	3, 2, 2, 2, 167, 168, 3, 2, 2, 2, 168, 175, 3, 2, 2, 2, 169, 171, 5, 88,
	45, 2, 170, 169, 3, 2, 2, 2, 171, 172, 3, 2, 2, 2, 172, 170, 3, 2, 2, 2,
	172, 173, 3, 2, 2, 2, 173, 176, 3, 2, 2, 2, 174, 176, 7, 8, 2, 2, 175,
	170, 3, 2, 2, 2, 175, 174, 3, 2, 2, 2, 176, 180, 3, 2, 2, 2, 177, 179,
	5, 18, 10, 2, 178, 177, 3, 2, 2, 2, 179, 182, 3, 2, 2, 2, 180, 178, 3,
	2, 2, 2, 180, 181, 3, 2, 2, 2, 181, 183, 3, 2, 2, 2, 182, 180, 3, 2, 2,
	2, 183, 184, 5, 26, 14, 2, 184, 185, 5, 28, 15, 2, 185, 11, 3, 2, 2, 2,
	186, 187, 7, 9, 2, 2, 187, 191, 5, 60, 31, 2, 188, 190, 5, 18, 10, 2, 189,
	188, 3, 2, 2, 2, 190, 193, 3, 2, 2, 2, 191, 189, 3, 2, 2, 2, 191, 192,
	3, 2, 2, 2, 192, 194, 3, 2, 2, 2, 193, 191, 3, 2, 2, 2, 194, 195, 5, 26,
	14, 2, 195, 196, 5, 28, 15, 2, 196, 13, 3, 2, 2, 2, 197, 204, 7, 10, 2,
	2, 198, 200, 5, 86, 44, 2, 199, 198, 3, 2, 2, 2, 200, 201, 3, 2, 2, 2,
	201, 199, 3, 2, 2, 2, 201, 202, 3, 2, 2, 2, 202, 205, 3, 2, 2, 2, 203,
	205, 7, 8, 2, 2, 204, 199, 3, 2, 2, 2, 204, 203, 3, 2, 2, 2, 205, 209,
	3, 2, 2, 2, 206, 208, 5, 18, 10, 2, 207, 206, 3, 2, 2, 2, 208, 211, 3,
	2, 2, 2, 209, 207, 3, 2, 2, 2, 209, 210, 3, 2, 2, 2, 210, 213, 3, 2, 2,
	2, 211, 209, 3, 2, 2, 2, 212, 214, 5, 26, 14, 2, 213, 212, 3, 2, 2, 2,
	213, 214, 3, 2, 2, 2, 214, 215, 3, 2, 2, 2, 215, 216, 5, 28, 15, 2, 216,
	15, 3, 2, 2, 2, 217, 221, 7, 11, 2, 2, 218, 220, 5, 18, 10, 2, 219, 218,
	3, 2, 2, 2, 220, 223, 3, 2, 2, 2, 221, 219, 3, 2, 2, 2, 221, 222, 3, 2,
	2, 2, 222, 224, 3, 2, 2, 2, 223, 221, 3, 2, 2, 2, 224, 225, 5, 26, 14,
	2, 225, 17, 3, 2, 2, 2, 226, 229, 7, 12, 2, 2, 227, 230, 5, 20, 11, 2,
	228, 230, 5, 22, 12, 2, 229, 227, 3, 2, 2, 2, 229, 228, 3, 2, 2, 2, 230,
	19, 3, 2, 2, 2, 231, 232, 5, 24, 13, 2, 232, 21, 3, 2, 2, 2, 233, 234,
	7, 13, 2, 2, 234, 235, 5, 24, 13, 2, 235, 23, 3, 2, 2, 2, 236, 237, 5,
	134, 68, 2, 237, 25, 3, 2, 2, 2, 238, 240, 7, 14, 2, 2, 239, 238, 3, 2,
	2, 2, 239, 240, 3, 2, 2, 2, 240, 241, 3, 2, 2, 2, 241, 242, 5, 40, 21,
	2, 242, 27, 3, 2, 2, 2, 243, 245, 5, 32, 17, 2, 244, 243, 3, 2, 2, 2, 244,
	245, 3, 2, 2, 2, 245, 247, 3, 2, 2, 2, 246, 248, 5, 30, 16, 2, 247, 246,
	3, 2, 2, 2, 247, 248, 3, 2, 2, 2, 248, 29, 3, 2, 2, 2, 249, 251, 5, 36,
	19, 2, 250, 252, 5, 38, 20, 2, 251, 250, 3, 2, 2, 2, 251, 252, 3, 2, 2,
	2, 252, 258, 3, 2, 2, 2, 253, 255, 5, 38, 20, 2, 254, 256, 5, 36, 19, 2,
	255, 254, 3, 2, 2, 2, 255, 256, 3, 2, 2, 2, 256, 258, 3, 2, 2, 2, 257,
	249, 3, 2, 2, 2, 257, 253, 3, 2, 2, 2, 258, 31, 3, 2, 2, 2, 259, 260, 7,
	15, 2, 2, 260, 262, 7, 16, 2, 2, 261, 263, 5, 34, 18, 2, 262, 261, 3, 2,
	2, 2, 263, 264, 3, 2, 2, 2, 264, 262, 3, 2, 2, 2, 264, 265, 3, 2, 2, 2,
	265, 33, 3, 2, 2, 2, 266, 267, 9, 3, 2, 2, 267, 273, 5, 112, 57, 2, 268,
	271, 5, 54, 28, 2, 269, 271, 5, 88, 45, 2, 270, 268, 3, 2, 2, 2, 270, 269,
	3, 2, 2, 2, 271, 273, 3, 2, 2, 2, 272, 266, 3, 2, 2, 2, 272, 270, 3, 2,
	2, 2, 273, 35, 3, 2, 2, 2, 274, 275, 7, 19, 2, 2, 275, 276, 7, 68, 2, 2,
	276, 37, 3, 2, 2, 2, 277, 278, 7, 20, 2, 2, 278, 279, 7, 68, 2, 2, 279,
	39, 3, 2, 2, 2, 280, 282, 7, 21, 2, 2, 281, 283, 5, 42, 22, 2, 282, 281,
	3, 2, 2, 2, 282, 283, 3, 2, 2, 2, 283, 296, 3, 2, 2, 2, 284, 287, 5, 44,
	23, 2, 285, 287, 5, 52, 27, 2, 286, 284, 3, 2, 2, 2, 286, 285, 3, 2, 2,
	2, 287, 289, 3, 2, 2, 2, 288, 290, 7, 22, 2, 2, 289, 288, 3, 2, 2, 2, 289,
	290, 3, 2, 2, 2, 290, 292, 3, 2, 2, 2, 291, 293, 5, 42, 22, 2, 292, 291,
	3, 2, 2, 2, 292, 293, 3, 2, 2, 2, 293, 295, 3, 2, 2, 2, 294, 286, 3, 2,
	2, 2, 295, 298, 3, 2, 2, 2, 296, 294, 3, 2, 2, 2, 296, 297, 3, 2, 2, 2,
	297, 299, 3, 2, 2, 2, 298, 296, 3, 2, 2, 2, 299, 300, 7, 23, 2, 2, 300,
	41, 3, 2, 2, 2, 301, 306, 5, 64, 33, 2, 302, 304, 7, 22, 2, 2, 303, 305,
	5, 42, 22, 2, 304, 303, 3, 2, 2, 2, 304, 305, 3, 2, 2, 2, 305, 307, 3,
	2, 2, 2, 306, 302, 3, 2, 2, 2, 306, 307, 3, 2, 2, 2, 307, 43, 3, 2, 2,
	2, 308, 312, 5, 46, 24, 2, 309, 312, 5, 50, 26, 2, 310, 312, 5, 48, 25,
	2, 311, 308, 3, 2, 2, 2, 311, 309, 3, 2, 2, 2, 311, 310, 3, 2, 2, 2, 312,
	45, 3, 2, 2, 2, 313, 314, 7, 24, 2, 2, 314, 315, 5, 40, 21, 2, 315, 47,
	3, 2, 2, 2, 316, 317, 7, 25, 2, 2, 317, 318, 5, 86, 44, 2, 318, 319, 5,
	40, 21, 2, 319, 49, 3, 2, 2, 2, 320, 325, 5, 40, 21, 2, 321, 322, 7, 26,
	2, 2, 322, 324, 5, 40, 21, 2, 323, 321, 3, 2, 2, 2, 324, 327, 3, 2, 2,
	2, 325, 323, 3, 2, 2, 2, 325, 326, 3, 2, 2, 2, 326, 51, 3, 2, 2, 2, 327,
	325, 3, 2, 2, 2, 328, 329, 7, 27, 2, 2, 329, 330, 5, 54, 28, 2, 330, 53,
	3, 2, 2, 2, 331, 335, 5, 112, 57, 2, 332, 335, 5, 114, 58, 2, 333, 335,
	5, 56, 29, 2, 334, 331, 3, 2, 2, 2, 334, 332, 3, 2, 2, 2, 334, 333, 3,
	2, 2, 2, 335, 55, 3, 2, 2, 2, 336, 337, 5, 134, 68, 2, 337, 338, 5, 58,
	30, 2, 338, 57, 3, 2, 2, 2, 339, 352, 7, 83, 2, 2, 340, 341, 7, 28, 2,
	2, 341, 346, 5, 92, 47, 2, 342, 343, 7, 29, 2, 2, 343, 345, 5, 92, 47,
	2, 344, 342, 3, 2, 2, 2, 345, 348, 3, 2, 2, 2, 346, 344, 3, 2, 2, 2, 346,
	347, 3, 2, 2, 2, 347, 349, 3, 2, 2, 2, 348, 346, 3, 2, 2, 2, 349, 350,
	7, 30, 2, 2, 350, 352, 3, 2, 2, 2, 351, 339, 3, 2, 2, 2, 351, 340, 3, 2,
	2, 2, 352, 59, 3, 2, 2, 2, 353, 355, 7, 21, 2, 2, 354, 356, 5, 62, 32,
	2, 355, 354, 3, 2, 2, 2, 355, 356, 3, 2, 2, 2, 356, 357, 3, 2, 2, 2, 357,
	358, 7, 23, 2, 2, 358, 61, 3, 2, 2, 2, 359, 364, 5, 64, 33, 2, 360, 362,
	7, 22, 2, 2, 361, 363, 5, 62, 32, 2, 362, 361, 3, 2, 2, 2, 362, 363, 3,
	2, 2, 2, 363, 365, 3, 2, 2, 2, 364, 360, 3, 2, 2, 2, 364, 365, 3, 2, 2,
	2, 365, 63, 3, 2, 2, 2, 366, 367, 5, 84, 43, 2, 367, 368, 5, 66, 34, 2,
	368, 373, 3, 2, 2, 2, 369, 370, 5, 76, 39, 2, 370, 371, 5, 68, 35, 2, 371,
	373, 3, 2, 2, 2, 372, 366, 3, 2, 2, 2, 372, 369, 3, 2, 2, 2, 373, 65, 3,
	2, 2, 2, 374, 375, 5, 74, 38, 2, 375, 384, 5, 70, 36, 2, 376, 380, 7, 31,
	2, 2, 377, 378, 5, 74, 38, 2, 378, 379, 5, 70, 36, 2, 379, 381, 3, 2, 2,
	2, 380, 377, 3, 2, 2, 2, 380, 381, 3, 2, 2, 2, 381, 383, 3, 2, 2, 2, 382,
	376, 3, 2, 2, 2, 383, 386, 3, 2, 2, 2, 384, 382, 3, 2, 2, 2, 384, 385,
	3, 2, 2, 2, 385, 67, 3, 2, 2, 2, 386, 384, 3, 2, 2, 2, 387, 389, 5, 66,
	34, 2, 388, 387, 3, 2, 2, 2, 388, 389, 3, 2, 2, 2, 389, 69, 3, 2, 2, 2,
	390, 395, 5, 72, 37, 2, 391, 392, 7, 29, 2, 2, 392, 394, 5, 72, 37, 2,
	393, 391, 3, 2, 2, 2, 394, 397, 3, 2, 2, 2, 395, 393, 3, 2, 2, 2, 395,
	396, 3, 2, 2, 2, 396, 71, 3, 2, 2, 2, 397, 395, 3, 2, 2, 2, 398, 399, 5,
	82, 42, 2, 399, 73, 3, 2, 2, 2, 400, 403, 5, 86, 44, 2, 401, 403, 7, 32,
	2, 2, 402, 400, 3, 2, 2, 2, 402, 401, 3, 2, 2, 2, 403, 75, 3, 2, 2, 2,
	404, 407, 5, 80, 41, 2, 405, 407, 5, 78, 40, 2, 406, 404, 3, 2, 2, 2, 406,
	405, 3, 2, 2, 2, 407, 77, 3, 2, 2, 2, 408, 409, 7, 33, 2, 2, 409, 410,
	5, 66, 34, 2, 410, 411, 7, 34, 2, 2, 411, 79, 3, 2, 2, 2, 412, 414, 7,
	28, 2, 2, 413, 415, 5, 82, 42, 2, 414, 413, 3, 2, 2, 2, 415, 416, 3, 2,
	2, 2, 416, 414, 3, 2, 2, 2, 416, 417, 3, 2, 2, 2, 417, 418, 3, 2, 2, 2,
	418, 419, 7, 30, 2, 2, 419, 81, 3, 2, 2, 2, 420, 423, 5, 84, 43, 2, 421,
	423, 5, 76, 39, 2, 422, 420, 3, 2, 2, 2, 422, 421, 3, 2, 2, 2, 423, 83,
	3, 2, 2, 2, 424, 427, 5, 88, 45, 2, 425, 427, 5, 90, 46, 2, 426, 424, 3,
	2, 2, 2, 426, 425, 3, 2, 2, 2, 427, 85, 3, 2, 2, 2, 428, 431, 5, 88, 45,
	2, 429, 431, 5, 134, 68, 2, 430, 428, 3, 2, 2, 2, 430, 429, 3, 2, 2, 2,
	431, 87, 3, 2, 2, 2, 432, 433, 9, 4, 2, 2, 433, 89, 3, 2, 2, 2, 434, 441,
	5, 134, 68, 2, 435, 441, 5, 120, 61, 2, 436, 441, 5, 122, 62, 2, 437, 441,
	5, 130, 66, 2, 438, 441, 5, 138, 70, 2, 439, 441, 7, 83, 2, 2, 440, 434,
	3, 2, 2, 2, 440, 435, 3, 2, 2, 2, 440, 436, 3, 2, 2, 2, 440, 437, 3, 2,
	2, 2, 440, 438, 3, 2, 2, 2, 440, 439, 3, 2, 2, 2, 441, 91, 3, 2, 2, 2,
	442, 443, 5, 94, 48, 2, 443, 93, 3, 2, 2, 2, 444, 449, 5, 96, 49, 2, 445,
	446, 7, 35, 2, 2, 446, 448, 5, 96, 49, 2, 447, 445, 3, 2, 2, 2, 448, 451,
	3, 2, 2, 2, 449, 447, 3, 2, 2, 2, 449, 450, 3, 2, 2, 2, 450, 95, 3, 2,
	2, 2, 451, 449, 3, 2, 2, 2, 452, 457, 5, 98, 50, 2, 453, 454, 7, 36, 2,
	2, 454, 456, 5, 98, 50, 2, 455, 453, 3, 2, 2, 2, 456, 459, 3, 2, 2, 2,
	457, 455, 3, 2, 2, 2, 457, 458, 3, 2, 2, 2, 458, 97, 3, 2, 2, 2, 459, 457,
	3, 2, 2, 2, 460, 461, 5, 100, 51, 2, 461, 99, 3, 2, 2, 2, 462, 475, 5,
	102, 52, 2, 463, 464, 7, 37, 2, 2, 464, 476, 5, 102, 52, 2, 465, 466, 7,
	38, 2, 2, 466, 476, 5, 102, 52, 2, 467, 468, 7, 39, 2, 2, 468, 476, 5,
	102, 52, 2, 469, 470, 7, 40, 2, 2, 470, 476, 5, 102, 52, 2, 471, 472, 7,
	41, 2, 2, 472, 476, 5, 102, 52, 2, 473, 474, 7, 42, 2, 2, 474, 476, 5,
	102, 52, 2, 475, 463, 3, 2, 2, 2, 475, 465, 3, 2, 2, 2, 475, 467, 3, 2,
	2, 2, 475, 469, 3, 2, 2, 2, 475, 471, 3, 2, 2, 2, 475, 473, 3, 2, 2, 2,
	475, 476, 3, 2, 2, 2, 476, 101, 3, 2, 2, 2, 477, 478, 5, 104, 53, 2, 478,
	103, 3, 2, 2, 2, 479, 488, 5, 106, 54, 2, 480, 481, 7, 43, 2, 2, 481, 487,
	5, 106, 54, 2, 482, 483, 7, 44, 2, 2, 483, 487, 5, 106, 54, 2, 484, 487,
	5, 126, 64, 2, 485, 487, 5, 128, 65, 2, 486, 480, 3, 2, 2, 2, 486, 482,
	3, 2, 2, 2, 486, 484, 3, 2, 2, 2, 486, 485, 3, 2, 2, 2, 487, 490, 3, 2,
	2, 2, 488, 486, 3, 2, 2, 2, 488, 489, 3, 2, 2, 2, 489, 105, 3, 2, 2, 2,
	490, 488, 3, 2, 2, 2, 491, 498, 5, 108, 55, 2, 492, 493, 7, 8, 2, 2, 493,
	497, 5, 108, 55, 2, 494, 495, 7, 45, 2, 2, 495, 497, 5, 108, 55, 2, 496,
	492, 3, 2, 2, 2, 496, 494, 3, 2, 2, 2, 497, 500, 3, 2, 2, 2, 498, 496,
	3, 2, 2, 2, 498, 499, 3, 2, 2, 2, 499, 107, 3, 2, 2, 2, 500, 498, 3, 2,
	2, 2, 501, 502, 7, 46, 2, 2, 502, 509, 5, 110, 56, 2, 503, 504, 7, 43,
	2, 2, 504, 509, 5, 110, 56, 2, 505, 506, 7, 44, 2, 2, 506, 509, 5, 110,
	56, 2, 507, 509, 5, 110, 56, 2, 508, 501, 3, 2, 2, 2, 508, 503, 3, 2, 2,
	2, 508, 505, 3, 2, 2, 2, 508, 507, 3, 2, 2, 2, 509, 109, 3, 2, 2, 2, 510,
	518, 5, 112, 57, 2, 511, 518, 5, 114, 58, 2, 512, 518, 5, 118, 60, 2, 513,
	518, 5, 120, 61, 2, 514, 518, 5, 122, 62, 2, 515, 518, 5, 130, 66, 2, 516,
	518, 5, 88, 45, 2, 517, 510, 3, 2, 2, 2, 517, 511, 3, 2, 2, 2, 517, 512,
	3, 2, 2, 2, 517, 513, 3, 2, 2, 2, 517, 514, 3, 2, 2, 2, 517, 515, 3, 2,
	2, 2, 517, 516, 3, 2, 2, 2, 518, 111, 3, 2, 2, 2, 519, 520, 7, 28, 2, 2,
	520, 521, 5, 92, 47, 2, 521, 522, 7, 30, 2, 2, 522, 113, 3, 2, 2, 2, 523,
	524, 7, 47, 2, 2, 524, 525, 7, 28, 2, 2, 525, 526, 5, 92, 47, 2, 526, 527,
	7, 30, 2, 2, 527, 579, 3, 2, 2, 2, 528, 529, 7, 48, 2, 2, 529, 530, 7,
	28, 2, 2, 530, 531, 5, 92, 47, 2, 531, 532, 7, 30, 2, 2, 532, 579, 3, 2,
	2, 2, 533, 534, 7, 49, 2, 2, 534, 535, 7, 28, 2, 2, 535, 536, 5, 92, 47,
	2, 536, 537, 7, 29, 2, 2, 537, 538, 5, 92, 47, 2, 538, 539, 7, 30, 2, 2,
	539, 579, 3, 2, 2, 2, 540, 541, 7, 50, 2, 2, 541, 542, 7, 28, 2, 2, 542,
	543, 5, 92, 47, 2, 543, 544, 7, 30, 2, 2, 544, 579, 3, 2, 2, 2, 545, 546,
	7, 51, 2, 2, 546, 547, 7, 28, 2, 2, 547, 548, 5, 88, 45, 2, 548, 549, 7,
	30, 2, 2, 549, 579, 3, 2, 2, 2, 550, 551, 7, 52, 2, 2, 551, 552, 7, 28,
	2, 2, 552, 553, 5, 92, 47, 2, 553, 554, 7, 29, 2, 2, 554, 555, 5, 92, 47,
	2, 555, 556, 7, 30, 2, 2, 556, 579, 3, 2, 2, 2, 557, 558, 7, 53, 2, 2,
	558, 559, 7, 28, 2, 2, 559, 560, 5, 92, 47, 2, 560, 561, 7, 30, 2, 2, 561,
	579, 3, 2, 2, 2, 562, 563, 7, 54, 2, 2, 563, 564, 7, 28, 2, 2, 564, 565,
	5, 92, 47, 2, 565, 566, 7, 30, 2, 2, 566, 579, 3, 2, 2, 2, 567, 568, 7,
	55, 2, 2, 568, 569, 7, 28, 2, 2, 569, 570, 5, 92, 47, 2, 570, 571, 7, 30,
	2, 2, 571, 579, 3, 2, 2, 2, 572, 573, 7, 56, 2, 2, 573, 574, 7, 28, 2,
	2, 574, 575, 5, 92, 47, 2, 575, 576, 7, 30, 2, 2, 576, 579, 3, 2, 2, 2,
	577, 579, 5, 116, 59, 2, 578, 523, 3, 2, 2, 2, 578, 528, 3, 2, 2, 2, 578,
	533, 3, 2, 2, 2, 578, 540, 3, 2, 2, 2, 578, 545, 3, 2, 2, 2, 578, 550,
	3, 2, 2, 2, 578, 557, 3, 2, 2, 2, 578, 562, 3, 2, 2, 2, 578, 567, 3, 2,
	2, 2, 578, 572, 3, 2, 2, 2, 578, 577, 3, 2, 2, 2, 579, 115, 3, 2, 2, 2,
	580, 581, 7, 57, 2, 2, 581, 582, 7, 28, 2, 2, 582, 583, 5, 92, 47, 2, 583,
	584, 7, 29, 2, 2, 584, 587, 5, 92, 47, 2, 585, 586, 7, 29, 2, 2, 586, 588,
	5, 92, 47, 2, 587, 585, 3, 2, 2, 2, 587, 588, 3, 2, 2, 2, 588, 589, 3,
	2, 2, 2, 589, 590, 7, 30, 2, 2, 590, 117, 3, 2, 2, 2, 591, 593, 5, 134,
	68, 2, 592, 594, 5, 58, 30, 2, 593, 592, 3, 2, 2, 2, 593, 594, 3, 2, 2,
	2, 594, 119, 3, 2, 2, 2, 595, 599, 5, 132, 67, 2, 596, 600, 7, 67, 2, 2,
	597, 598, 7, 58, 2, 2, 598, 600, 5, 134, 68, 2, 599, 596, 3, 2, 2, 2, 599,
	597, 3, 2, 2, 2, 599, 600, 3, 2, 2, 2, 600, 121, 3, 2, 2, 2, 601, 605,
	5, 124, 63, 2, 602, 605, 5, 126, 64, 2, 603, 605, 5, 128, 65, 2, 604, 601,
	3, 2, 2, 2, 604, 602, 3, 2, 2, 2, 604, 603, 3, 2, 2, 2, 605, 123, 3, 2,
	2, 2, 606, 607, 9, 5, 2, 2, 607, 125, 3, 2, 2, 2, 608, 609, 9, 6, 2, 2,
	609, 127, 3, 2, 2, 2, 610, 611, 9, 7, 2, 2, 611, 129, 3, 2, 2, 2, 612,
	613, 9, 8, 2, 2, 613, 131, 3, 2, 2, 2, 614, 615, 9, 9, 2, 2, 615, 133,
	3, 2, 2, 2, 616, 619, 7, 61, 2, 2, 617, 619, 5, 136, 69, 2, 618, 616, 3,
	2, 2, 2, 618, 617, 3, 2, 2, 2, 619, 135, 3, 2, 2, 2, 620, 621, 9, 10, 2,
	2, 621, 137, 3, 2, 2, 2, 622, 623, 9, 11, 2, 2, 623, 139, 3, 2, 2, 2, 67,
	145, 150, 155, 167, 172, 175, 180, 191, 201, 204, 209, 213, 221, 229, 239,
	244, 247, 251, 255, 257, 264, 270, 272, 282, 286, 289, 292, 296, 304, 306,
	311, 325, 334, 346, 351, 355, 362, 364, 372, 380, 384, 388, 395, 402, 406,
	416, 422, 426, 430, 440, 449, 457, 475, 486, 488, 496, 498, 508, 517, 578,
	587, 593, 599, 604, 618,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'BASE'", "'PREFIX'", "'SELECT'", "'DISTINCT'", "'REDUCED'", "'*'",
	"'CONSTRUCT'", "'DESCRIBE'", "'ASK'", "'FROM'", "'NAMED'", "'WHERE'", "'ORDER'",
	"'BY'", "'ASC'", "'DESC'", "'LIMIT'", "'OFFSET'", "'{'", "'.'", "'}'",
	"'OPTIONAL'", "'GRAPH'", "'UNION'", "'FILTER'", "'('", "','", "')'", "';'",
	"'A'", "'['", "']'", "'||'", "'&&'", "'='", "'!='", "'<'", "'>'", "'<='",
	"'>='", "'+'", "'-'", "'/'", "'!'", "'STR'", "'LANG'", "'LANGMATCHES'",
	"'DATATYPE'", "'BOUND'", "'SAMETERM'", "'ISIRI'", "'ISURI'", "'ISBLANK'",
	"'ISLITERAL'", "'REGEX'", "'^^'", "'TRUE'", "'FALSE'",
}
var symbolicNames = []string{
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

var ruleNames = []string{
	"query", "prologue", "baseDecl", "prefixDecl", "selectQuery", "constructQuery",
	"describeQuery", "askQuery", "datasetClause", "defaultGraphClause", "namedGraphClause",
	"sourceSelector", "whereClause", "solutionModifier", "limitOffsetClauses",
	"orderClause", "orderCondition", "limitClause", "offsetClause", "groupGraphPattern",
	"triplesBlock", "graphPatternNotTriples", "optionalGraphPattern", "graphGraphPattern",
	"groupOrUnionGraphPattern", "filter", "constraint", "functionCall", "argList",
	"constructTemplate", "constructTriples", "triplesSameSubject", "propertyListNotEmpty",
	"propertyList", "objectList", "object", "verb", "triplesNode", "blankNodePropertyList",
	"collection", "graphNode", "varOrTerm", "varOrIRIref", "varx", "graphTerm",
	"expression", "conditionalOrExpression", "conditionalAndExpression", "valueLogical",
	"relationalExpression", "numericExpression", "additiveExpression", "multiplicativeExpression",
	"unaryExpression", "primaryExpression", "brackettedExpression", "builtInCall",
	"regexExpression", "iriRefOrFunction", "rdfLiteral", "numericLiteral",
	"numericLiteralUnsigned", "numericLiteralPositive", "numericLiteralNegative",
	"booleanLiteral", "rdfstring", "iriRef", "prefixedName", "blankNode",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type SparqlParser struct {
	*antlr.BaseParser
}

func NewSparqlParser(input antlr.TokenStream) *SparqlParser {
	this := new(SparqlParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
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
	p.RuleIndex = SparqlParserRULE_query
	return p
}

func (*QueryContext) IsQueryContext() {}

func NewQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QueryContext {
	var p = new(QueryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_query

	return p
}

func (s *QueryContext) GetParser() antlr.Parser { return s.parser }

func (s *QueryContext) Prologue() IPrologueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrologueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrologueContext)
}

func (s *QueryContext) EOF() antlr.TerminalNode {
	return s.GetToken(SparqlParserEOF, 0)
}

func (s *QueryContext) SelectQuery() ISelectQueryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectQueryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectQueryContext)
}

func (s *QueryContext) ConstructQuery() IConstructQueryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstructQueryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstructQueryContext)
}

func (s *QueryContext) DescribeQuery() IDescribeQueryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDescribeQueryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDescribeQueryContext)
}

func (s *QueryContext) AskQuery() IAskQueryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAskQueryContext)(nil)).Elem(), 0)

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
		p.SetState(138)
		p.Prologue()
	}
	p.SetState(143)
	p.GetErrorHandler().Sync(p)

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
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(145)
		p.Match(SparqlParserEOF)
	}

	return localctx
}

// IPrologueContext is an interface to support dynamic dispatch.
type IPrologueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrologueContext differentiates from other interfaces.
	IsPrologueContext()
}

type PrologueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrologueContext() *PrologueContext {
	var p = new(PrologueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SparqlParserRULE_prologue
	return p
}

func (*PrologueContext) IsPrologueContext() {}

func NewPrologueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrologueContext {
	var p = new(PrologueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_prologue

	return p
}

func (s *PrologueContext) GetParser() antlr.Parser { return s.parser }

func (s *PrologueContext) BaseDecl() IBaseDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBaseDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBaseDeclContext)
}

func (s *PrologueContext) AllPrefixDecl() []IPrefixDeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPrefixDeclContext)(nil)).Elem())
	var tst = make([]IPrefixDeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPrefixDeclContext)
		}
	}

	return tst
}

func (s *PrologueContext) PrefixDecl(i int) IPrefixDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrefixDeclContext)(nil)).Elem(), i)

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
	p.SetState(148)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SparqlParserT__0 {
		{
			p.SetState(147)
			p.BaseDecl()
		}

	}
	p.SetState(153)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SparqlParserT__1 {
		{
			p.SetState(150)
			p.PrefixDecl()
		}

		p.SetState(155)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
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

	// IsBaseDeclContext differentiates from other interfaces.
	IsBaseDeclContext()
}

type BaseDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	iri    antlr.Token
}

func NewEmptyBaseDeclContext() *BaseDeclContext {
	var p = new(BaseDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SparqlParserRULE_baseDecl
	return p
}

func (*BaseDeclContext) IsBaseDeclContext() {}

func NewBaseDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BaseDeclContext {
	var p = new(BaseDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
		p.SetState(156)
		p.Match(SparqlParserT__0)
	}
	{
		p.SetState(157)

		var _m = p.Match(SparqlParserIRI_REF)

		localctx.(*BaseDeclContext).iri = _m
	}

	return localctx
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

	// IsPrefixDeclContext differentiates from other interfaces.
	IsPrefixDeclContext()
}

type PrefixDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	prefix antlr.Token
	iri    antlr.Token
}

func NewEmptyPrefixDeclContext() *PrefixDeclContext {
	var p = new(PrefixDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SparqlParserRULE_prefixDecl
	return p
}

func (*PrefixDeclContext) IsPrefixDeclContext() {}

func NewPrefixDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrefixDeclContext {
	var p = new(PrefixDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
		p.SetState(159)
		p.Match(SparqlParserT__1)
	}
	{
		p.SetState(160)

		var _m = p.Match(SparqlParserPNAME_NS)

		localctx.(*PrefixDeclContext).prefix = _m
	}
	{
		p.SetState(161)

		var _m = p.Match(SparqlParserIRI_REF)

		localctx.(*PrefixDeclContext).iri = _m
	}

	return localctx
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

	// IsSelectQueryContext differentiates from other interfaces.
	IsSelectQueryContext()
}

type SelectQueryContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	mod     antlr.Token
	_varx   IVarxContext
	vars    []IVarxContext
	varstar antlr.Token
}

func NewEmptySelectQueryContext() *SelectQueryContext {
	var p = new(SelectQueryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SparqlParserRULE_selectQuery
	return p
}

func (*SelectQueryContext) IsSelectQueryContext() {}

func NewSelectQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectQueryContext {
	var p = new(SelectQueryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhereClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWhereClauseContext)
}

func (s *SelectQueryContext) SolutionModifier() ISolutionModifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISolutionModifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISolutionModifierContext)
}

func (s *SelectQueryContext) AllDatasetClause() []IDatasetClauseContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDatasetClauseContext)(nil)).Elem())
	var tst = make([]IDatasetClauseContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDatasetClauseContext)
		}
	}

	return tst
}

func (s *SelectQueryContext) DatasetClause(i int) IDatasetClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDatasetClauseContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDatasetClauseContext)
}

func (s *SelectQueryContext) AllVarx() []IVarxContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVarxContext)(nil)).Elem())
	var tst = make([]IVarxContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVarxContext)
		}
	}

	return tst
}

func (s *SelectQueryContext) Varx(i int) IVarxContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarxContext)(nil)).Elem(), i)

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
		p.SetState(163)
		p.Match(SparqlParserT__2)
	}
	p.SetState(165)
	p.GetErrorHandler().Sync(p)
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

	switch p.GetTokenStream().LA(1) {
	case SparqlParserVAR1, SparqlParserVAR2:
		p.SetState(168)
		p.GetErrorHandler().Sync(p)
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
			_la = p.GetTokenStream().LA(1)
		}

	case SparqlParserT__5:
		{
			p.SetState(172)

			var _m = p.Match(SparqlParserT__5)

			localctx.(*SelectQueryContext).varstar = _m
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(178)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SparqlParserT__9 {
		{
			p.SetState(175)
			p.DatasetClause()
		}

		p.SetState(180)
		p.GetErrorHandler().Sync(p)
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

	return localctx
}

// IConstructQueryContext is an interface to support dynamic dispatch.
type IConstructQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstructQueryContext differentiates from other interfaces.
	IsConstructQueryContext()
}

type ConstructQueryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstructQueryContext() *ConstructQueryContext {
	var p = new(ConstructQueryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SparqlParserRULE_constructQuery
	return p
}

func (*ConstructQueryContext) IsConstructQueryContext() {}

func NewConstructQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstructQueryContext {
	var p = new(ConstructQueryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_constructQuery

	return p
}

func (s *ConstructQueryContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstructQueryContext) ConstructTemplate() IConstructTemplateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstructTemplateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstructTemplateContext)
}

func (s *ConstructQueryContext) WhereClause() IWhereClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhereClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWhereClauseContext)
}

func (s *ConstructQueryContext) SolutionModifier() ISolutionModifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISolutionModifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISolutionModifierContext)
}

func (s *ConstructQueryContext) AllDatasetClause() []IDatasetClauseContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDatasetClauseContext)(nil)).Elem())
	var tst = make([]IDatasetClauseContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDatasetClauseContext)
		}
	}

	return tst
}

func (s *ConstructQueryContext) DatasetClause(i int) IDatasetClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDatasetClauseContext)(nil)).Elem(), i)

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
		p.SetState(184)
		p.Match(SparqlParserT__6)
	}
	{
		p.SetState(185)
		p.ConstructTemplate()
	}
	p.SetState(189)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SparqlParserT__9 {
		{
			p.SetState(186)
			p.DatasetClause()
		}

		p.SetState(191)
		p.GetErrorHandler().Sync(p)
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

	return localctx
}

// IDescribeQueryContext is an interface to support dynamic dispatch.
type IDescribeQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDescribeQueryContext differentiates from other interfaces.
	IsDescribeQueryContext()
}

type DescribeQueryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDescribeQueryContext() *DescribeQueryContext {
	var p = new(DescribeQueryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SparqlParserRULE_describeQuery
	return p
}

func (*DescribeQueryContext) IsDescribeQueryContext() {}

func NewDescribeQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DescribeQueryContext {
	var p = new(DescribeQueryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_describeQuery

	return p
}

func (s *DescribeQueryContext) GetParser() antlr.Parser { return s.parser }

func (s *DescribeQueryContext) SolutionModifier() ISolutionModifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISolutionModifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISolutionModifierContext)
}

func (s *DescribeQueryContext) AllDatasetClause() []IDatasetClauseContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDatasetClauseContext)(nil)).Elem())
	var tst = make([]IDatasetClauseContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDatasetClauseContext)
		}
	}

	return tst
}

func (s *DescribeQueryContext) DatasetClause(i int) IDatasetClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDatasetClauseContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDatasetClauseContext)
}

func (s *DescribeQueryContext) WhereClause() IWhereClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhereClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWhereClauseContext)
}

func (s *DescribeQueryContext) AllVarOrIRIref() []IVarOrIRIrefContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVarOrIRIrefContext)(nil)).Elem())
	var tst = make([]IVarOrIRIrefContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVarOrIRIrefContext)
		}
	}

	return tst
}

func (s *DescribeQueryContext) VarOrIRIref(i int) IVarOrIRIrefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarOrIRIrefContext)(nil)).Elem(), i)

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
		p.SetState(195)
		p.Match(SparqlParserT__7)
	}
	p.SetState(202)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SparqlParserIRI_REF, SparqlParserPNAME_NS, SparqlParserPNAME_LN, SparqlParserVAR1, SparqlParserVAR2:
		p.SetState(197)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = (((_la-59)&-(0x1f+1)) == 0 && ((1<<uint((_la-59)))&((1<<(SparqlParserIRI_REF-59))|(1<<(SparqlParserPNAME_NS-59))|(1<<(SparqlParserPNAME_LN-59))|(1<<(SparqlParserVAR1-59))|(1<<(SparqlParserVAR2-59)))) != 0) {
			{
				p.SetState(196)
				p.VarOrIRIref()
			}

			p.SetState(199)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case SparqlParserT__5:
		{
			p.SetState(201)
			p.Match(SparqlParserT__5)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(207)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SparqlParserT__9 {
		{
			p.SetState(204)
			p.DatasetClause()
		}

		p.SetState(209)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(211)
	p.GetErrorHandler().Sync(p)
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

	return localctx
}

// IAskQueryContext is an interface to support dynamic dispatch.
type IAskQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAskQueryContext differentiates from other interfaces.
	IsAskQueryContext()
}

type AskQueryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAskQueryContext() *AskQueryContext {
	var p = new(AskQueryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SparqlParserRULE_askQuery
	return p
}

func (*AskQueryContext) IsAskQueryContext() {}

func NewAskQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AskQueryContext {
	var p = new(AskQueryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_askQuery

	return p
}

func (s *AskQueryContext) GetParser() antlr.Parser { return s.parser }

func (s *AskQueryContext) WhereClause() IWhereClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhereClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWhereClauseContext)
}

func (s *AskQueryContext) AllDatasetClause() []IDatasetClauseContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDatasetClauseContext)(nil)).Elem())
	var tst = make([]IDatasetClauseContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDatasetClauseContext)
		}
	}

	return tst
}

func (s *AskQueryContext) DatasetClause(i int) IDatasetClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDatasetClauseContext)(nil)).Elem(), i)

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
		p.SetState(215)
		p.Match(SparqlParserT__8)
	}
	p.SetState(219)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SparqlParserT__9 {
		{
			p.SetState(216)
			p.DatasetClause()
		}

		p.SetState(221)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(222)
		p.WhereClause()
	}

	return localctx
}

// IDatasetClauseContext is an interface to support dynamic dispatch.
type IDatasetClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDatasetClauseContext differentiates from other interfaces.
	IsDatasetClauseContext()
}

type DatasetClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDatasetClauseContext() *DatasetClauseContext {
	var p = new(DatasetClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SparqlParserRULE_datasetClause
	return p
}

func (*DatasetClauseContext) IsDatasetClauseContext() {}

func NewDatasetClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DatasetClauseContext {
	var p = new(DatasetClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_datasetClause

	return p
}

func (s *DatasetClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *DatasetClauseContext) DefaultGraphClause() IDefaultGraphClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefaultGraphClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDefaultGraphClauseContext)
}

func (s *DatasetClauseContext) NamedGraphClause() INamedGraphClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INamedGraphClauseContext)(nil)).Elem(), 0)

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
		p.SetState(224)
		p.Match(SparqlParserT__9)
	}
	p.SetState(227)
	p.GetErrorHandler().Sync(p)

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
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IDefaultGraphClauseContext is an interface to support dynamic dispatch.
type IDefaultGraphClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefaultGraphClauseContext differentiates from other interfaces.
	IsDefaultGraphClauseContext()
}

type DefaultGraphClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefaultGraphClauseContext() *DefaultGraphClauseContext {
	var p = new(DefaultGraphClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SparqlParserRULE_defaultGraphClause
	return p
}

func (*DefaultGraphClauseContext) IsDefaultGraphClauseContext() {}

func NewDefaultGraphClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefaultGraphClauseContext {
	var p = new(DefaultGraphClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_defaultGraphClause

	return p
}

func (s *DefaultGraphClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *DefaultGraphClauseContext) SourceSelector() ISourceSelectorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISourceSelectorContext)(nil)).Elem(), 0)

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
		p.SetState(229)
		p.SourceSelector()
	}

	return localctx
}

// INamedGraphClauseContext is an interface to support dynamic dispatch.
type INamedGraphClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNamedGraphClauseContext differentiates from other interfaces.
	IsNamedGraphClauseContext()
}

type NamedGraphClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNamedGraphClauseContext() *NamedGraphClauseContext {
	var p = new(NamedGraphClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SparqlParserRULE_namedGraphClause
	return p
}

func (*NamedGraphClauseContext) IsNamedGraphClauseContext() {}

func NewNamedGraphClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NamedGraphClauseContext {
	var p = new(NamedGraphClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_namedGraphClause

	return p
}

func (s *NamedGraphClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *NamedGraphClauseContext) SourceSelector() ISourceSelectorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISourceSelectorContext)(nil)).Elem(), 0)

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
		p.SetState(231)
		p.Match(SparqlParserT__10)
	}
	{
		p.SetState(232)
		p.SourceSelector()
	}

	return localctx
}

// ISourceSelectorContext is an interface to support dynamic dispatch.
type ISourceSelectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSourceSelectorContext differentiates from other interfaces.
	IsSourceSelectorContext()
}

type SourceSelectorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySourceSelectorContext() *SourceSelectorContext {
	var p = new(SourceSelectorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SparqlParserRULE_sourceSelector
	return p
}

func (*SourceSelectorContext) IsSourceSelectorContext() {}

func NewSourceSelectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SourceSelectorContext {
	var p = new(SourceSelectorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_sourceSelector

	return p
}

func (s *SourceSelectorContext) GetParser() antlr.Parser { return s.parser }

func (s *SourceSelectorContext) IriRef() IIriRefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIriRefContext)(nil)).Elem(), 0)

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
		p.SetState(234)
		p.IriRef()
	}

	return localctx
}

// IWhereClauseContext is an interface to support dynamic dispatch.
type IWhereClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWhereClauseContext differentiates from other interfaces.
	IsWhereClauseContext()
}

type WhereClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhereClauseContext() *WhereClauseContext {
	var p = new(WhereClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SparqlParserRULE_whereClause
	return p
}

func (*WhereClauseContext) IsWhereClauseContext() {}

func NewWhereClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhereClauseContext {
	var p = new(WhereClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_whereClause

	return p
}

func (s *WhereClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *WhereClauseContext) GroupGraphPattern() IGroupGraphPatternContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGroupGraphPatternContext)(nil)).Elem(), 0)

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
	p.SetState(237)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SparqlParserT__11 {
		{
			p.SetState(236)
			p.Match(SparqlParserT__11)
		}

	}
	{
		p.SetState(239)
		p.GroupGraphPattern()
	}

	return localctx
}

// ISolutionModifierContext is an interface to support dynamic dispatch.
type ISolutionModifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSolutionModifierContext differentiates from other interfaces.
	IsSolutionModifierContext()
}

type SolutionModifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySolutionModifierContext() *SolutionModifierContext {
	var p = new(SolutionModifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SparqlParserRULE_solutionModifier
	return p
}

func (*SolutionModifierContext) IsSolutionModifierContext() {}

func NewSolutionModifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SolutionModifierContext {
	var p = new(SolutionModifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_solutionModifier

	return p
}

func (s *SolutionModifierContext) GetParser() antlr.Parser { return s.parser }

func (s *SolutionModifierContext) OrderClause() IOrderClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOrderClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOrderClauseContext)
}

func (s *SolutionModifierContext) LimitOffsetClauses() ILimitOffsetClausesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILimitOffsetClausesContext)(nil)).Elem(), 0)

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
	p.SetState(242)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SparqlParserT__12 {
		{
			p.SetState(241)
			p.OrderClause()
		}

	}
	p.SetState(245)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SparqlParserT__16 || _la == SparqlParserT__17 {
		{
			p.SetState(244)
			p.LimitOffsetClauses()
		}

	}

	return localctx
}

// ILimitOffsetClausesContext is an interface to support dynamic dispatch.
type ILimitOffsetClausesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLimitOffsetClausesContext differentiates from other interfaces.
	IsLimitOffsetClausesContext()
}

type LimitOffsetClausesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLimitOffsetClausesContext() *LimitOffsetClausesContext {
	var p = new(LimitOffsetClausesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SparqlParserRULE_limitOffsetClauses
	return p
}

func (*LimitOffsetClausesContext) IsLimitOffsetClausesContext() {}

func NewLimitOffsetClausesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LimitOffsetClausesContext {
	var p = new(LimitOffsetClausesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_limitOffsetClauses

	return p
}

func (s *LimitOffsetClausesContext) GetParser() antlr.Parser { return s.parser }

func (s *LimitOffsetClausesContext) LimitClause() ILimitClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILimitClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILimitClauseContext)
}

func (s *LimitOffsetClausesContext) OffsetClause() IOffsetClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOffsetClauseContext)(nil)).Elem(), 0)

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
	p.SetState(255)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SparqlParserT__16:
		{
			p.SetState(247)
			p.LimitClause()
		}
		p.SetState(249)
		p.GetErrorHandler().Sync(p)
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
		_la = p.GetTokenStream().LA(1)

		if _la == SparqlParserT__16 {
			{
				p.SetState(252)
				p.LimitClause()
			}

		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IOrderClauseContext is an interface to support dynamic dispatch.
type IOrderClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOrderClauseContext differentiates from other interfaces.
	IsOrderClauseContext()
}

type OrderClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrderClauseContext() *OrderClauseContext {
	var p = new(OrderClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SparqlParserRULE_orderClause
	return p
}

func (*OrderClauseContext) IsOrderClauseContext() {}

func NewOrderClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrderClauseContext {
	var p = new(OrderClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_orderClause

	return p
}

func (s *OrderClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *OrderClauseContext) AllOrderCondition() []IOrderConditionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOrderConditionContext)(nil)).Elem())
	var tst = make([]IOrderConditionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOrderConditionContext)
		}
	}

	return tst
}

func (s *OrderClauseContext) OrderCondition(i int) IOrderConditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOrderConditionContext)(nil)).Elem(), i)

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
		p.SetState(257)
		p.Match(SparqlParserT__12)
	}
	{
		p.SetState(258)
		p.Match(SparqlParserT__13)
	}
	p.SetState(260)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SparqlParserT__14)|(1<<SparqlParserT__15)|(1<<SparqlParserT__25))) != 0) || (((_la-45)&-(0x1f+1)) == 0 && ((1<<uint((_la-45)))&((1<<(SparqlParserT__44-45))|(1<<(SparqlParserT__45-45))|(1<<(SparqlParserT__46-45))|(1<<(SparqlParserT__47-45))|(1<<(SparqlParserT__48-45))|(1<<(SparqlParserT__49-45))|(1<<(SparqlParserT__50-45))|(1<<(SparqlParserT__51-45))|(1<<(SparqlParserT__52-45))|(1<<(SparqlParserT__53-45))|(1<<(SparqlParserT__54-45))|(1<<(SparqlParserIRI_REF-45))|(1<<(SparqlParserPNAME_NS-45))|(1<<(SparqlParserPNAME_LN-45))|(1<<(SparqlParserVAR1-45))|(1<<(SparqlParserVAR2-45)))) != 0) {
		{
			p.SetState(259)
			p.OrderCondition()
		}

		p.SetState(262)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IOrderConditionContext is an interface to support dynamic dispatch.
type IOrderConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOrderConditionContext differentiates from other interfaces.
	IsOrderConditionContext()
}

type OrderConditionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrderConditionContext() *OrderConditionContext {
	var p = new(OrderConditionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SparqlParserRULE_orderCondition
	return p
}

func (*OrderConditionContext) IsOrderConditionContext() {}

func NewOrderConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrderConditionContext {
	var p = new(OrderConditionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_orderCondition

	return p
}

func (s *OrderConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *OrderConditionContext) BrackettedExpression() IBrackettedExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBrackettedExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBrackettedExpressionContext)
}

func (s *OrderConditionContext) Constraint() IConstraintContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstraintContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstraintContext)
}

func (s *OrderConditionContext) Varx() IVarxContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarxContext)(nil)).Elem(), 0)

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

	p.SetState(270)
	p.GetErrorHandler().Sync(p)

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
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ILimitClauseContext is an interface to support dynamic dispatch.
type ILimitClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLimitClauseContext differentiates from other interfaces.
	IsLimitClauseContext()
}

type LimitClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLimitClauseContext() *LimitClauseContext {
	var p = new(LimitClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SparqlParserRULE_limitClause
	return p
}

func (*LimitClauseContext) IsLimitClauseContext() {}

func NewLimitClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LimitClauseContext {
	var p = new(LimitClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
		p.SetState(272)
		p.Match(SparqlParserT__16)
	}
	{
		p.SetState(273)
		p.Match(SparqlParserINTEGER)
	}

	return localctx
}

// IOffsetClauseContext is an interface to support dynamic dispatch.
type IOffsetClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOffsetClauseContext differentiates from other interfaces.
	IsOffsetClauseContext()
}

type OffsetClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOffsetClauseContext() *OffsetClauseContext {
	var p = new(OffsetClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SparqlParserRULE_offsetClause
	return p
}

func (*OffsetClauseContext) IsOffsetClauseContext() {}

func NewOffsetClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OffsetClauseContext {
	var p = new(OffsetClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
		p.SetState(275)
		p.Match(SparqlParserT__17)
	}
	{
		p.SetState(276)
		p.Match(SparqlParserINTEGER)
	}

	return localctx
}

// IGroupGraphPatternContext is an interface to support dynamic dispatch.
type IGroupGraphPatternContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGroupGraphPatternContext differentiates from other interfaces.
	IsGroupGraphPatternContext()
}

type GroupGraphPatternContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGroupGraphPatternContext() *GroupGraphPatternContext {
	var p = new(GroupGraphPatternContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SparqlParserRULE_groupGraphPattern
	return p
}

func (*GroupGraphPatternContext) IsGroupGraphPatternContext() {}

func NewGroupGraphPatternContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GroupGraphPatternContext {
	var p = new(GroupGraphPatternContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_groupGraphPattern

	return p
}

func (s *GroupGraphPatternContext) GetParser() antlr.Parser { return s.parser }

func (s *GroupGraphPatternContext) AllTriplesBlock() []ITriplesBlockContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITriplesBlockContext)(nil)).Elem())
	var tst = make([]ITriplesBlockContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITriplesBlockContext)
		}
	}

	return tst
}

func (s *GroupGraphPatternContext) TriplesBlock(i int) ITriplesBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITriplesBlockContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITriplesBlockContext)
}

func (s *GroupGraphPatternContext) AllGraphPatternNotTriples() []IGraphPatternNotTriplesContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IGraphPatternNotTriplesContext)(nil)).Elem())
	var tst = make([]IGraphPatternNotTriplesContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IGraphPatternNotTriplesContext)
		}
	}

	return tst
}

func (s *GroupGraphPatternContext) GraphPatternNotTriples(i int) IGraphPatternNotTriplesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGraphPatternNotTriplesContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IGraphPatternNotTriplesContext)
}

func (s *GroupGraphPatternContext) AllFilter() []IFilterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFilterContext)(nil)).Elem())
	var tst = make([]IFilterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFilterContext)
		}
	}

	return tst
}

func (s *GroupGraphPatternContext) Filter(i int) IFilterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFilterContext)(nil)).Elem(), i)

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
		p.SetState(278)
		p.Match(SparqlParserT__18)
	}
	p.SetState(280)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SparqlParserT__25 || _la == SparqlParserT__30 || (((_la-57)&-(0x1f+1)) == 0 && ((1<<uint((_la-57)))&((1<<(SparqlParserT__56-57))|(1<<(SparqlParserT__57-57))|(1<<(SparqlParserIRI_REF-57))|(1<<(SparqlParserPNAME_NS-57))|(1<<(SparqlParserPNAME_LN-57))|(1<<(SparqlParserBLANK_NODE_LABEL-57))|(1<<(SparqlParserVAR1-57))|(1<<(SparqlParserVAR2-57))|(1<<(SparqlParserINTEGER-57))|(1<<(SparqlParserDECIMAL-57))|(1<<(SparqlParserDOUBLE-57))|(1<<(SparqlParserINTEGER_POSITIVE-57))|(1<<(SparqlParserDECIMAL_POSITIVE-57))|(1<<(SparqlParserDOUBLE_POSITIVE-57))|(1<<(SparqlParserINTEGER_NEGATIVE-57))|(1<<(SparqlParserDECIMAL_NEGATIVE-57))|(1<<(SparqlParserDOUBLE_NEGATIVE-57))|(1<<(SparqlParserSTRING_LITERAL1-57))|(1<<(SparqlParserSTRING_LITERAL2-57))|(1<<(SparqlParserNIL-57))|(1<<(SparqlParserANON-57)))) != 0) {
		{
			p.SetState(279)
			p.TriplesBlock()
		}

	}
	p.SetState(294)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SparqlParserT__18)|(1<<SparqlParserT__21)|(1<<SparqlParserT__22)|(1<<SparqlParserT__24))) != 0 {
		p.SetState(284)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case SparqlParserT__18, SparqlParserT__21, SparqlParserT__22:
			{
				p.SetState(282)
				p.GraphPatternNotTriples()
			}

		case SparqlParserT__24:
			{
				p.SetState(283)
				p.Filter()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}
		p.SetState(287)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SparqlParserT__19 {
			{
				p.SetState(286)
				p.Match(SparqlParserT__19)
			}

		}
		p.SetState(290)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SparqlParserT__25 || _la == SparqlParserT__30 || (((_la-57)&-(0x1f+1)) == 0 && ((1<<uint((_la-57)))&((1<<(SparqlParserT__56-57))|(1<<(SparqlParserT__57-57))|(1<<(SparqlParserIRI_REF-57))|(1<<(SparqlParserPNAME_NS-57))|(1<<(SparqlParserPNAME_LN-57))|(1<<(SparqlParserBLANK_NODE_LABEL-57))|(1<<(SparqlParserVAR1-57))|(1<<(SparqlParserVAR2-57))|(1<<(SparqlParserINTEGER-57))|(1<<(SparqlParserDECIMAL-57))|(1<<(SparqlParserDOUBLE-57))|(1<<(SparqlParserINTEGER_POSITIVE-57))|(1<<(SparqlParserDECIMAL_POSITIVE-57))|(1<<(SparqlParserDOUBLE_POSITIVE-57))|(1<<(SparqlParserINTEGER_NEGATIVE-57))|(1<<(SparqlParserDECIMAL_NEGATIVE-57))|(1<<(SparqlParserDOUBLE_NEGATIVE-57))|(1<<(SparqlParserSTRING_LITERAL1-57))|(1<<(SparqlParserSTRING_LITERAL2-57))|(1<<(SparqlParserNIL-57))|(1<<(SparqlParserANON-57)))) != 0) {
			{
				p.SetState(289)
				p.TriplesBlock()
			}

		}

		p.SetState(296)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(297)
		p.Match(SparqlParserT__20)
	}

	return localctx
}

// ITriplesBlockContext is an interface to support dynamic dispatch.
type ITriplesBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTriplesBlockContext differentiates from other interfaces.
	IsTriplesBlockContext()
}

type TriplesBlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTriplesBlockContext() *TriplesBlockContext {
	var p = new(TriplesBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SparqlParserRULE_triplesBlock
	return p
}

func (*TriplesBlockContext) IsTriplesBlockContext() {}

func NewTriplesBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TriplesBlockContext {
	var p = new(TriplesBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_triplesBlock

	return p
}

func (s *TriplesBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *TriplesBlockContext) TriplesSameSubject() ITriplesSameSubjectContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITriplesSameSubjectContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITriplesSameSubjectContext)
}

func (s *TriplesBlockContext) TriplesBlock() ITriplesBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITriplesBlockContext)(nil)).Elem(), 0)

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
		p.SetState(299)
		p.TriplesSameSubject()
	}
	p.SetState(304)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SparqlParserT__19 {
		{
			p.SetState(300)
			p.Match(SparqlParserT__19)
		}
		p.SetState(302)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SparqlParserT__25 || _la == SparqlParserT__30 || (((_la-57)&-(0x1f+1)) == 0 && ((1<<uint((_la-57)))&((1<<(SparqlParserT__56-57))|(1<<(SparqlParserT__57-57))|(1<<(SparqlParserIRI_REF-57))|(1<<(SparqlParserPNAME_NS-57))|(1<<(SparqlParserPNAME_LN-57))|(1<<(SparqlParserBLANK_NODE_LABEL-57))|(1<<(SparqlParserVAR1-57))|(1<<(SparqlParserVAR2-57))|(1<<(SparqlParserINTEGER-57))|(1<<(SparqlParserDECIMAL-57))|(1<<(SparqlParserDOUBLE-57))|(1<<(SparqlParserINTEGER_POSITIVE-57))|(1<<(SparqlParserDECIMAL_POSITIVE-57))|(1<<(SparqlParserDOUBLE_POSITIVE-57))|(1<<(SparqlParserINTEGER_NEGATIVE-57))|(1<<(SparqlParserDECIMAL_NEGATIVE-57))|(1<<(SparqlParserDOUBLE_NEGATIVE-57))|(1<<(SparqlParserSTRING_LITERAL1-57))|(1<<(SparqlParserSTRING_LITERAL2-57))|(1<<(SparqlParserNIL-57))|(1<<(SparqlParserANON-57)))) != 0) {
			{
				p.SetState(301)
				p.TriplesBlock()
			}

		}

	}

	return localctx
}

// IGraphPatternNotTriplesContext is an interface to support dynamic dispatch.
type IGraphPatternNotTriplesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGraphPatternNotTriplesContext differentiates from other interfaces.
	IsGraphPatternNotTriplesContext()
}

type GraphPatternNotTriplesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGraphPatternNotTriplesContext() *GraphPatternNotTriplesContext {
	var p = new(GraphPatternNotTriplesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SparqlParserRULE_graphPatternNotTriples
	return p
}

func (*GraphPatternNotTriplesContext) IsGraphPatternNotTriplesContext() {}

func NewGraphPatternNotTriplesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GraphPatternNotTriplesContext {
	var p = new(GraphPatternNotTriplesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_graphPatternNotTriples

	return p
}

func (s *GraphPatternNotTriplesContext) GetParser() antlr.Parser { return s.parser }

func (s *GraphPatternNotTriplesContext) OptionalGraphPattern() IOptionalGraphPatternContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionalGraphPatternContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionalGraphPatternContext)
}

func (s *GraphPatternNotTriplesContext) GroupOrUnionGraphPattern() IGroupOrUnionGraphPatternContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGroupOrUnionGraphPatternContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGroupOrUnionGraphPatternContext)
}

func (s *GraphPatternNotTriplesContext) GraphGraphPattern() IGraphGraphPatternContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGraphGraphPatternContext)(nil)).Elem(), 0)

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

	p.SetState(309)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SparqlParserT__21:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(306)
			p.OptionalGraphPattern()
		}

	case SparqlParserT__18:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(307)
			p.GroupOrUnionGraphPattern()
		}

	case SparqlParserT__22:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(308)
			p.GraphGraphPattern()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IOptionalGraphPatternContext is an interface to support dynamic dispatch.
type IOptionalGraphPatternContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOptionalGraphPatternContext differentiates from other interfaces.
	IsOptionalGraphPatternContext()
}

type OptionalGraphPatternContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionalGraphPatternContext() *OptionalGraphPatternContext {
	var p = new(OptionalGraphPatternContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SparqlParserRULE_optionalGraphPattern
	return p
}

func (*OptionalGraphPatternContext) IsOptionalGraphPatternContext() {}

func NewOptionalGraphPatternContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionalGraphPatternContext {
	var p = new(OptionalGraphPatternContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_optionalGraphPattern

	return p
}

func (s *OptionalGraphPatternContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionalGraphPatternContext) GroupGraphPattern() IGroupGraphPatternContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGroupGraphPatternContext)(nil)).Elem(), 0)

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
		p.SetState(311)
		p.Match(SparqlParserT__21)
	}
	{
		p.SetState(312)
		p.GroupGraphPattern()
	}

	return localctx
}

// IGraphGraphPatternContext is an interface to support dynamic dispatch.
type IGraphGraphPatternContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGraphGraphPatternContext differentiates from other interfaces.
	IsGraphGraphPatternContext()
}

type GraphGraphPatternContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGraphGraphPatternContext() *GraphGraphPatternContext {
	var p = new(GraphGraphPatternContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SparqlParserRULE_graphGraphPattern
	return p
}

func (*GraphGraphPatternContext) IsGraphGraphPatternContext() {}

func NewGraphGraphPatternContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GraphGraphPatternContext {
	var p = new(GraphGraphPatternContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_graphGraphPattern

	return p
}

func (s *GraphGraphPatternContext) GetParser() antlr.Parser { return s.parser }

func (s *GraphGraphPatternContext) VarOrIRIref() IVarOrIRIrefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarOrIRIrefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarOrIRIrefContext)
}

func (s *GraphGraphPatternContext) GroupGraphPattern() IGroupGraphPatternContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGroupGraphPatternContext)(nil)).Elem(), 0)

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
		p.SetState(314)
		p.Match(SparqlParserT__22)
	}
	{
		p.SetState(315)
		p.VarOrIRIref()
	}
	{
		p.SetState(316)
		p.GroupGraphPattern()
	}

	return localctx
}

// IGroupOrUnionGraphPatternContext is an interface to support dynamic dispatch.
type IGroupOrUnionGraphPatternContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGroupOrUnionGraphPatternContext differentiates from other interfaces.
	IsGroupOrUnionGraphPatternContext()
}

type GroupOrUnionGraphPatternContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGroupOrUnionGraphPatternContext() *GroupOrUnionGraphPatternContext {
	var p = new(GroupOrUnionGraphPatternContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SparqlParserRULE_groupOrUnionGraphPattern
	return p
}

func (*GroupOrUnionGraphPatternContext) IsGroupOrUnionGraphPatternContext() {}

func NewGroupOrUnionGraphPatternContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GroupOrUnionGraphPatternContext {
	var p = new(GroupOrUnionGraphPatternContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_groupOrUnionGraphPattern

	return p
}

func (s *GroupOrUnionGraphPatternContext) GetParser() antlr.Parser { return s.parser }

func (s *GroupOrUnionGraphPatternContext) AllGroupGraphPattern() []IGroupGraphPatternContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IGroupGraphPatternContext)(nil)).Elem())
	var tst = make([]IGroupGraphPatternContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IGroupGraphPatternContext)
		}
	}

	return tst
}

func (s *GroupOrUnionGraphPatternContext) GroupGraphPattern(i int) IGroupGraphPatternContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGroupGraphPatternContext)(nil)).Elem(), i)

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
		p.SetState(318)
		p.GroupGraphPattern()
	}
	p.SetState(323)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SparqlParserT__23 {
		{
			p.SetState(319)
			p.Match(SparqlParserT__23)
		}
		{
			p.SetState(320)
			p.GroupGraphPattern()
		}

		p.SetState(325)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFilterContext is an interface to support dynamic dispatch.
type IFilterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFilterContext differentiates from other interfaces.
	IsFilterContext()
}

type FilterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFilterContext() *FilterContext {
	var p = new(FilterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SparqlParserRULE_filter
	return p
}

func (*FilterContext) IsFilterContext() {}

func NewFilterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FilterContext {
	var p = new(FilterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_filter

	return p
}

func (s *FilterContext) GetParser() antlr.Parser { return s.parser }

func (s *FilterContext) Constraint() IConstraintContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstraintContext)(nil)).Elem(), 0)

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
		p.SetState(326)
		p.Match(SparqlParserT__24)
	}
	{
		p.SetState(327)
		p.Constraint()
	}

	return localctx
}

// IConstraintContext is an interface to support dynamic dispatch.
type IConstraintContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

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
	p.RuleIndex = SparqlParserRULE_constraint
	return p
}

func (*ConstraintContext) IsConstraintContext() {}

func NewConstraintContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstraintContext {
	var p = new(ConstraintContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_constraint

	return p
}

func (s *ConstraintContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstraintContext) BrackettedExpression() IBrackettedExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBrackettedExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBrackettedExpressionContext)
}

func (s *ConstraintContext) BuiltInCall() IBuiltInCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBuiltInCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBuiltInCallContext)
}

func (s *ConstraintContext) FunctionCall() IFunctionCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionCallContext)(nil)).Elem(), 0)

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

	p.SetState(332)
	p.GetErrorHandler().Sync(p)

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
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFunctionCallContext is an interface to support dynamic dispatch.
type IFunctionCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionCallContext differentiates from other interfaces.
	IsFunctionCallContext()
}

type FunctionCallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionCallContext() *FunctionCallContext {
	var p = new(FunctionCallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SparqlParserRULE_functionCall
	return p
}

func (*FunctionCallContext) IsFunctionCallContext() {}

func NewFunctionCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionCallContext {
	var p = new(FunctionCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_functionCall

	return p
}

func (s *FunctionCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionCallContext) IriRef() IIriRefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIriRefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIriRefContext)
}

func (s *FunctionCallContext) ArgList() IArgListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgListContext)(nil)).Elem(), 0)

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
		p.SetState(334)
		p.IriRef()
	}
	{
		p.SetState(335)
		p.ArgList()
	}

	return localctx
}

// IArgListContext is an interface to support dynamic dispatch.
type IArgListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgListContext differentiates from other interfaces.
	IsArgListContext()
}

type ArgListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgListContext() *ArgListContext {
	var p = new(ArgListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SparqlParserRULE_argList
	return p
}

func (*ArgListContext) IsArgListContext() {}

func NewArgListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgListContext {
	var p = new(ArgListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_argList

	return p
}

func (s *ArgListContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgListContext) NIL() antlr.TerminalNode {
	return s.GetToken(SparqlParserNIL, 0)
}

func (s *ArgListContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ArgListContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

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
	p.SetState(349)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SparqlParserNIL:
		{
			p.SetState(337)
			p.Match(SparqlParserNIL)
		}

	case SparqlParserT__25:
		{
			p.SetState(338)
			p.Match(SparqlParserT__25)
		}
		{
			p.SetState(339)
			p.Expression()
		}
		p.SetState(344)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == SparqlParserT__26 {
			{
				p.SetState(340)
				p.Match(SparqlParserT__26)
			}
			{
				p.SetState(341)
				p.Expression()
			}

			p.SetState(346)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(347)
			p.Match(SparqlParserT__27)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IConstructTemplateContext is an interface to support dynamic dispatch.
type IConstructTemplateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstructTemplateContext differentiates from other interfaces.
	IsConstructTemplateContext()
}

type ConstructTemplateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstructTemplateContext() *ConstructTemplateContext {
	var p = new(ConstructTemplateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SparqlParserRULE_constructTemplate
	return p
}

func (*ConstructTemplateContext) IsConstructTemplateContext() {}

func NewConstructTemplateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstructTemplateContext {
	var p = new(ConstructTemplateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_constructTemplate

	return p
}

func (s *ConstructTemplateContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstructTemplateContext) ConstructTriples() IConstructTriplesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstructTriplesContext)(nil)).Elem(), 0)

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
		p.SetState(351)
		p.Match(SparqlParserT__18)
	}
	p.SetState(353)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SparqlParserT__25 || _la == SparqlParserT__30 || (((_la-57)&-(0x1f+1)) == 0 && ((1<<uint((_la-57)))&((1<<(SparqlParserT__56-57))|(1<<(SparqlParserT__57-57))|(1<<(SparqlParserIRI_REF-57))|(1<<(SparqlParserPNAME_NS-57))|(1<<(SparqlParserPNAME_LN-57))|(1<<(SparqlParserBLANK_NODE_LABEL-57))|(1<<(SparqlParserVAR1-57))|(1<<(SparqlParserVAR2-57))|(1<<(SparqlParserINTEGER-57))|(1<<(SparqlParserDECIMAL-57))|(1<<(SparqlParserDOUBLE-57))|(1<<(SparqlParserINTEGER_POSITIVE-57))|(1<<(SparqlParserDECIMAL_POSITIVE-57))|(1<<(SparqlParserDOUBLE_POSITIVE-57))|(1<<(SparqlParserINTEGER_NEGATIVE-57))|(1<<(SparqlParserDECIMAL_NEGATIVE-57))|(1<<(SparqlParserDOUBLE_NEGATIVE-57))|(1<<(SparqlParserSTRING_LITERAL1-57))|(1<<(SparqlParserSTRING_LITERAL2-57))|(1<<(SparqlParserNIL-57))|(1<<(SparqlParserANON-57)))) != 0) {
		{
			p.SetState(352)
			p.ConstructTriples()
		}

	}
	{
		p.SetState(355)
		p.Match(SparqlParserT__20)
	}

	return localctx
}

// IConstructTriplesContext is an interface to support dynamic dispatch.
type IConstructTriplesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstructTriplesContext differentiates from other interfaces.
	IsConstructTriplesContext()
}

type ConstructTriplesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstructTriplesContext() *ConstructTriplesContext {
	var p = new(ConstructTriplesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SparqlParserRULE_constructTriples
	return p
}

func (*ConstructTriplesContext) IsConstructTriplesContext() {}

func NewConstructTriplesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstructTriplesContext {
	var p = new(ConstructTriplesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_constructTriples

	return p
}

func (s *ConstructTriplesContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstructTriplesContext) TriplesSameSubject() ITriplesSameSubjectContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITriplesSameSubjectContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITriplesSameSubjectContext)
}

func (s *ConstructTriplesContext) ConstructTriples() IConstructTriplesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstructTriplesContext)(nil)).Elem(), 0)

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
		p.SetState(357)
		p.TriplesSameSubject()
	}
	p.SetState(362)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SparqlParserT__19 {
		{
			p.SetState(358)
			p.Match(SparqlParserT__19)
		}
		p.SetState(360)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SparqlParserT__25 || _la == SparqlParserT__30 || (((_la-57)&-(0x1f+1)) == 0 && ((1<<uint((_la-57)))&((1<<(SparqlParserT__56-57))|(1<<(SparqlParserT__57-57))|(1<<(SparqlParserIRI_REF-57))|(1<<(SparqlParserPNAME_NS-57))|(1<<(SparqlParserPNAME_LN-57))|(1<<(SparqlParserBLANK_NODE_LABEL-57))|(1<<(SparqlParserVAR1-57))|(1<<(SparqlParserVAR2-57))|(1<<(SparqlParserINTEGER-57))|(1<<(SparqlParserDECIMAL-57))|(1<<(SparqlParserDOUBLE-57))|(1<<(SparqlParserINTEGER_POSITIVE-57))|(1<<(SparqlParserDECIMAL_POSITIVE-57))|(1<<(SparqlParserDOUBLE_POSITIVE-57))|(1<<(SparqlParserINTEGER_NEGATIVE-57))|(1<<(SparqlParserDECIMAL_NEGATIVE-57))|(1<<(SparqlParserDOUBLE_NEGATIVE-57))|(1<<(SparqlParserSTRING_LITERAL1-57))|(1<<(SparqlParserSTRING_LITERAL2-57))|(1<<(SparqlParserNIL-57))|(1<<(SparqlParserANON-57)))) != 0) {
			{
				p.SetState(359)
				p.ConstructTriples()
			}

		}

	}

	return localctx
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

	// IsTriplesSameSubjectContext differentiates from other interfaces.
	IsTriplesSameSubjectContext()
}

type TriplesSameSubjectContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	subject    IVarOrTermContext
	properties IPropertyListNotEmptyContext
}

func NewEmptyTriplesSameSubjectContext() *TriplesSameSubjectContext {
	var p = new(TriplesSameSubjectContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SparqlParserRULE_triplesSameSubject
	return p
}

func (*TriplesSameSubjectContext) IsTriplesSameSubjectContext() {}

func NewTriplesSameSubjectContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TriplesSameSubjectContext {
	var p = new(TriplesSameSubjectContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarOrTermContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarOrTermContext)
}

func (s *TriplesSameSubjectContext) PropertyListNotEmpty() IPropertyListNotEmptyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertyListNotEmptyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPropertyListNotEmptyContext)
}

func (s *TriplesSameSubjectContext) TriplesNode() ITriplesNodeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITriplesNodeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITriplesNodeContext)
}

func (s *TriplesSameSubjectContext) PropertyList() IPropertyListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertyListContext)(nil)).Elem(), 0)

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

	p.SetState(370)
	p.GetErrorHandler().Sync(p)

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
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
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

	// IsPropertyListNotEmptyContext differentiates from other interfaces.
	IsPropertyListNotEmptyContext()
}

type PropertyListNotEmptyContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	_verb       IVerbContext
	verbs       []IVerbContext
	_objectList IObjectListContext
	ol          []IObjectListContext
}

func NewEmptyPropertyListNotEmptyContext() *PropertyListNotEmptyContext {
	var p = new(PropertyListNotEmptyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SparqlParserRULE_propertyListNotEmpty
	return p
}

func (*PropertyListNotEmptyContext) IsPropertyListNotEmptyContext() {}

func NewPropertyListNotEmptyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropertyListNotEmptyContext {
	var p = new(PropertyListNotEmptyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVerbContext)(nil)).Elem())
	var tst = make([]IVerbContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVerbContext)
		}
	}

	return tst
}

func (s *PropertyListNotEmptyContext) Verb(i int) IVerbContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVerbContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVerbContext)
}

func (s *PropertyListNotEmptyContext) AllObjectList() []IObjectListContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IObjectListContext)(nil)).Elem())
	var tst = make([]IObjectListContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IObjectListContext)
		}
	}

	return tst
}

func (s *PropertyListNotEmptyContext) ObjectList(i int) IObjectListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjectListContext)(nil)).Elem(), i)

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
	_la = p.GetTokenStream().LA(1)

	for _la == SparqlParserT__28 {
		{
			p.SetState(374)
			p.Match(SparqlParserT__28)
		}
		p.SetState(378)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SparqlParserT__29 || (((_la-59)&-(0x1f+1)) == 0 && ((1<<uint((_la-59)))&((1<<(SparqlParserIRI_REF-59))|(1<<(SparqlParserPNAME_NS-59))|(1<<(SparqlParserPNAME_LN-59))|(1<<(SparqlParserVAR1-59))|(1<<(SparqlParserVAR2-59)))) != 0) {
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
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IPropertyListContext is an interface to support dynamic dispatch.
type IPropertyListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPropertyListContext differentiates from other interfaces.
	IsPropertyListContext()
}

type PropertyListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPropertyListContext() *PropertyListContext {
	var p = new(PropertyListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SparqlParserRULE_propertyList
	return p
}

func (*PropertyListContext) IsPropertyListContext() {}

func NewPropertyListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropertyListContext {
	var p = new(PropertyListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_propertyList

	return p
}

func (s *PropertyListContext) GetParser() antlr.Parser { return s.parser }

func (s *PropertyListContext) PropertyListNotEmpty() IPropertyListNotEmptyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertyListNotEmptyContext)(nil)).Elem(), 0)

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
	p.SetState(386)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SparqlParserT__29 || (((_la-59)&-(0x1f+1)) == 0 && ((1<<uint((_la-59)))&((1<<(SparqlParserIRI_REF-59))|(1<<(SparqlParserPNAME_NS-59))|(1<<(SparqlParserPNAME_LN-59))|(1<<(SparqlParserVAR1-59))|(1<<(SparqlParserVAR2-59)))) != 0) {
		{
			p.SetState(385)
			p.PropertyListNotEmpty()
		}

	}

	return localctx
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

	// IsObjectListContext differentiates from other interfaces.
	IsObjectListContext()
}

type ObjectListContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	_object IObjectContext
	ob      []IObjectContext
}

func NewEmptyObjectListContext() *ObjectListContext {
	var p = new(ObjectListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SparqlParserRULE_objectList
	return p
}

func (*ObjectListContext) IsObjectListContext() {}

func NewObjectListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectListContext {
	var p = new(ObjectListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IObjectContext)(nil)).Elem())
	var tst = make([]IObjectContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IObjectContext)
		}
	}

	return tst
}

func (s *ObjectListContext) Object(i int) IObjectContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjectContext)(nil)).Elem(), i)

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
		p.SetState(388)

		var _x = p.Object()

		localctx.(*ObjectListContext)._object = _x
	}
	localctx.(*ObjectListContext).ob = append(localctx.(*ObjectListContext).ob, localctx.(*ObjectListContext)._object)
	p.SetState(393)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SparqlParserT__26 {
		{
			p.SetState(389)
			p.Match(SparqlParserT__26)
		}
		{
			p.SetState(390)

			var _x = p.Object()

			localctx.(*ObjectListContext)._object = _x
		}
		localctx.(*ObjectListContext).ob = append(localctx.(*ObjectListContext).ob, localctx.(*ObjectListContext)._object)

		p.SetState(395)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
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

	// IsObjectContext differentiates from other interfaces.
	IsObjectContext()
}

type ObjectContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	gn     IGraphNodeContext
}

func NewEmptyObjectContext() *ObjectContext {
	var p = new(ObjectContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SparqlParserRULE_object
	return p
}

func (*ObjectContext) IsObjectContext() {}

func NewObjectContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectContext {
	var p = new(ObjectContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_object

	return p
}

func (s *ObjectContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectContext) GetGn() IGraphNodeContext { return s.gn }

func (s *ObjectContext) SetGn(v IGraphNodeContext) { s.gn = v }

func (s *ObjectContext) GraphNode() IGraphNodeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGraphNodeContext)(nil)).Elem(), 0)

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
		p.SetState(396)

		var _x = p.GraphNode()

		localctx.(*ObjectContext).gn = _x
	}

	return localctx
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

	// IsVerbContext differentiates from other interfaces.
	IsVerbContext()
}

type VerbContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	vi     IVarOrIRIrefContext
	a      antlr.Token
}

func NewEmptyVerbContext() *VerbContext {
	var p = new(VerbContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SparqlParserRULE_verb
	return p
}

func (*VerbContext) IsVerbContext() {}

func NewVerbContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VerbContext {
	var p = new(VerbContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarOrIRIrefContext)(nil)).Elem(), 0)

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

	p.SetState(400)
	p.GetErrorHandler().Sync(p)

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
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITriplesNodeContext is an interface to support dynamic dispatch.
type ITriplesNodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTriplesNodeContext differentiates from other interfaces.
	IsTriplesNodeContext()
}

type TriplesNodeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTriplesNodeContext() *TriplesNodeContext {
	var p = new(TriplesNodeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SparqlParserRULE_triplesNode
	return p
}

func (*TriplesNodeContext) IsTriplesNodeContext() {}

func NewTriplesNodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TriplesNodeContext {
	var p = new(TriplesNodeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_triplesNode

	return p
}

func (s *TriplesNodeContext) GetParser() antlr.Parser { return s.parser }

func (s *TriplesNodeContext) Collection() ICollectionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICollectionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICollectionContext)
}

func (s *TriplesNodeContext) BlankNodePropertyList() IBlankNodePropertyListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlankNodePropertyListContext)(nil)).Elem(), 0)

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

	p.SetState(404)
	p.GetErrorHandler().Sync(p)

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
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IBlankNodePropertyListContext is an interface to support dynamic dispatch.
type IBlankNodePropertyListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlankNodePropertyListContext differentiates from other interfaces.
	IsBlankNodePropertyListContext()
}

type BlankNodePropertyListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlankNodePropertyListContext() *BlankNodePropertyListContext {
	var p = new(BlankNodePropertyListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SparqlParserRULE_blankNodePropertyList
	return p
}

func (*BlankNodePropertyListContext) IsBlankNodePropertyListContext() {}

func NewBlankNodePropertyListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlankNodePropertyListContext {
	var p = new(BlankNodePropertyListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_blankNodePropertyList

	return p
}

func (s *BlankNodePropertyListContext) GetParser() antlr.Parser { return s.parser }

func (s *BlankNodePropertyListContext) PropertyListNotEmpty() IPropertyListNotEmptyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertyListNotEmptyContext)(nil)).Elem(), 0)

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
		p.SetState(406)
		p.Match(SparqlParserT__30)
	}
	{
		p.SetState(407)
		p.PropertyListNotEmpty()
	}
	{
		p.SetState(408)
		p.Match(SparqlParserT__31)
	}

	return localctx
}

// ICollectionContext is an interface to support dynamic dispatch.
type ICollectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCollectionContext differentiates from other interfaces.
	IsCollectionContext()
}

type CollectionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCollectionContext() *CollectionContext {
	var p = new(CollectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SparqlParserRULE_collection
	return p
}

func (*CollectionContext) IsCollectionContext() {}

func NewCollectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CollectionContext {
	var p = new(CollectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_collection

	return p
}

func (s *CollectionContext) GetParser() antlr.Parser { return s.parser }

func (s *CollectionContext) AllGraphNode() []IGraphNodeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IGraphNodeContext)(nil)).Elem())
	var tst = make([]IGraphNodeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IGraphNodeContext)
		}
	}

	return tst
}

func (s *CollectionContext) GraphNode(i int) IGraphNodeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGraphNodeContext)(nil)).Elem(), i)

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
		p.SetState(410)
		p.Match(SparqlParserT__25)
	}
	p.SetState(412)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == SparqlParserT__25 || _la == SparqlParserT__30 || (((_la-57)&-(0x1f+1)) == 0 && ((1<<uint((_la-57)))&((1<<(SparqlParserT__56-57))|(1<<(SparqlParserT__57-57))|(1<<(SparqlParserIRI_REF-57))|(1<<(SparqlParserPNAME_NS-57))|(1<<(SparqlParserPNAME_LN-57))|(1<<(SparqlParserBLANK_NODE_LABEL-57))|(1<<(SparqlParserVAR1-57))|(1<<(SparqlParserVAR2-57))|(1<<(SparqlParserINTEGER-57))|(1<<(SparqlParserDECIMAL-57))|(1<<(SparqlParserDOUBLE-57))|(1<<(SparqlParserINTEGER_POSITIVE-57))|(1<<(SparqlParserDECIMAL_POSITIVE-57))|(1<<(SparqlParserDOUBLE_POSITIVE-57))|(1<<(SparqlParserINTEGER_NEGATIVE-57))|(1<<(SparqlParserDECIMAL_NEGATIVE-57))|(1<<(SparqlParserDOUBLE_NEGATIVE-57))|(1<<(SparqlParserSTRING_LITERAL1-57))|(1<<(SparqlParserSTRING_LITERAL2-57))|(1<<(SparqlParserNIL-57))|(1<<(SparqlParserANON-57)))) != 0) {
		{
			p.SetState(411)
			p.GraphNode()
		}

		p.SetState(414)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(416)
		p.Match(SparqlParserT__27)
	}

	return localctx
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

	// IsGraphNodeContext differentiates from other interfaces.
	IsGraphNodeContext()
}

type GraphNodeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	vt     IVarOrTermContext
	tn     ITriplesNodeContext
}

func NewEmptyGraphNodeContext() *GraphNodeContext {
	var p = new(GraphNodeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SparqlParserRULE_graphNode
	return p
}

func (*GraphNodeContext) IsGraphNodeContext() {}

func NewGraphNodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GraphNodeContext {
	var p = new(GraphNodeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarOrTermContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarOrTermContext)
}

func (s *GraphNodeContext) TriplesNode() ITriplesNodeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITriplesNodeContext)(nil)).Elem(), 0)

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

	p.SetState(420)
	p.GetErrorHandler().Sync(p)

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
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
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

	// IsVarOrTermContext differentiates from other interfaces.
	IsVarOrTermContext()
}

type VarOrTermContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	variable IVarxContext
	gt       IGraphTermContext
}

func NewEmptyVarOrTermContext() *VarOrTermContext {
	var p = new(VarOrTermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SparqlParserRULE_varOrTerm
	return p
}

func (*VarOrTermContext) IsVarOrTermContext() {}

func NewVarOrTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarOrTermContext {
	var p = new(VarOrTermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarxContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarxContext)
}

func (s *VarOrTermContext) GraphTerm() IGraphTermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGraphTermContext)(nil)).Elem(), 0)

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

	p.SetState(424)
	p.GetErrorHandler().Sync(p)

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
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
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

	// IsVarOrIRIrefContext differentiates from other interfaces.
	IsVarOrIRIrefContext()
}

type VarOrIRIrefContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	variable IVarxContext
	iri      IIriRefContext
}

func NewEmptyVarOrIRIrefContext() *VarOrIRIrefContext {
	var p = new(VarOrIRIrefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SparqlParserRULE_varOrIRIref
	return p
}

func (*VarOrIRIrefContext) IsVarOrIRIrefContext() {}

func NewVarOrIRIrefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarOrIRIrefContext {
	var p = new(VarOrIRIrefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarxContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarxContext)
}

func (s *VarOrIRIrefContext) IriRef() IIriRefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIriRefContext)(nil)).Elem(), 0)

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

	p.SetState(428)
	p.GetErrorHandler().Sync(p)

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
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IVarxContext is an interface to support dynamic dispatch.
type IVarxContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVarxContext differentiates from other interfaces.
	IsVarxContext()
}

type VarxContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarxContext() *VarxContext {
	var p = new(VarxContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SparqlParserRULE_varx
	return p
}

func (*VarxContext) IsVarxContext() {}

func NewVarxContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarxContext {
	var p = new(VarxContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
		p.SetState(430)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SparqlParserVAR1 || _la == SparqlParserVAR2) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
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

	// IsGraphTermContext differentiates from other interfaces.
	IsGraphTermContext()
}

type GraphTermContext struct {
	*antlr.BaseParserRuleContext
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
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SparqlParserRULE_graphTerm
	return p
}

func (*GraphTermContext) IsGraphTermContext() {}

func NewGraphTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GraphTermContext {
	var p = new(GraphTermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIriRefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIriRefContext)
}

func (s *GraphTermContext) RdfLiteral() IRdfLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRdfLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRdfLiteralContext)
}

func (s *GraphTermContext) NumericLiteral() INumericLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumericLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumericLiteralContext)
}

func (s *GraphTermContext) BooleanLiteral() IBooleanLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBooleanLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBooleanLiteralContext)
}

func (s *GraphTermContext) BlankNode() IBlankNodeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlankNodeContext)(nil)).Elem(), 0)

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

	p.SetState(438)
	p.GetErrorHandler().Sync(p)

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
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SparqlParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) ConditionalOrExpression() IConditionalOrExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionalOrExpressionContext)(nil)).Elem(), 0)

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
		p.SetState(440)
		p.ConditionalOrExpression()
	}

	return localctx
}

// IConditionalOrExpressionContext is an interface to support dynamic dispatch.
type IConditionalOrExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConditionalOrExpressionContext differentiates from other interfaces.
	IsConditionalOrExpressionContext()
}

type ConditionalOrExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionalOrExpressionContext() *ConditionalOrExpressionContext {
	var p = new(ConditionalOrExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SparqlParserRULE_conditionalOrExpression
	return p
}

func (*ConditionalOrExpressionContext) IsConditionalOrExpressionContext() {}

func NewConditionalOrExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionalOrExpressionContext {
	var p = new(ConditionalOrExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_conditionalOrExpression

	return p
}

func (s *ConditionalOrExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionalOrExpressionContext) AllConditionalAndExpression() []IConditionalAndExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IConditionalAndExpressionContext)(nil)).Elem())
	var tst = make([]IConditionalAndExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IConditionalAndExpressionContext)
		}
	}

	return tst
}

func (s *ConditionalOrExpressionContext) ConditionalAndExpression(i int) IConditionalAndExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionalAndExpressionContext)(nil)).Elem(), i)

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
		p.SetState(442)
		p.ConditionalAndExpression()
	}
	p.SetState(447)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SparqlParserT__32 {
		{
			p.SetState(443)
			p.Match(SparqlParserT__32)
		}
		{
			p.SetState(444)
			p.ConditionalAndExpression()
		}

		p.SetState(449)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IConditionalAndExpressionContext is an interface to support dynamic dispatch.
type IConditionalAndExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConditionalAndExpressionContext differentiates from other interfaces.
	IsConditionalAndExpressionContext()
}

type ConditionalAndExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionalAndExpressionContext() *ConditionalAndExpressionContext {
	var p = new(ConditionalAndExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SparqlParserRULE_conditionalAndExpression
	return p
}

func (*ConditionalAndExpressionContext) IsConditionalAndExpressionContext() {}

func NewConditionalAndExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionalAndExpressionContext {
	var p = new(ConditionalAndExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_conditionalAndExpression

	return p
}

func (s *ConditionalAndExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionalAndExpressionContext) AllValueLogical() []IValueLogicalContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValueLogicalContext)(nil)).Elem())
	var tst = make([]IValueLogicalContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValueLogicalContext)
		}
	}

	return tst
}

func (s *ConditionalAndExpressionContext) ValueLogical(i int) IValueLogicalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueLogicalContext)(nil)).Elem(), i)

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
		p.SetState(450)
		p.ValueLogical()
	}
	p.SetState(455)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SparqlParserT__33 {
		{
			p.SetState(451)
			p.Match(SparqlParserT__33)
		}
		{
			p.SetState(452)
			p.ValueLogical()
		}

		p.SetState(457)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IValueLogicalContext is an interface to support dynamic dispatch.
type IValueLogicalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueLogicalContext differentiates from other interfaces.
	IsValueLogicalContext()
}

type ValueLogicalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueLogicalContext() *ValueLogicalContext {
	var p = new(ValueLogicalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SparqlParserRULE_valueLogical
	return p
}

func (*ValueLogicalContext) IsValueLogicalContext() {}

func NewValueLogicalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueLogicalContext {
	var p = new(ValueLogicalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_valueLogical

	return p
}

func (s *ValueLogicalContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueLogicalContext) RelationalExpression() IRelationalExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRelationalExpressionContext)(nil)).Elem(), 0)

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
		p.SetState(458)
		p.RelationalExpression()
	}

	return localctx
}

// IRelationalExpressionContext is an interface to support dynamic dispatch.
type IRelationalExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRelationalExpressionContext differentiates from other interfaces.
	IsRelationalExpressionContext()
}

type RelationalExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelationalExpressionContext() *RelationalExpressionContext {
	var p = new(RelationalExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SparqlParserRULE_relationalExpression
	return p
}

func (*RelationalExpressionContext) IsRelationalExpressionContext() {}

func NewRelationalExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationalExpressionContext {
	var p = new(RelationalExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_relationalExpression

	return p
}

func (s *RelationalExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationalExpressionContext) AllNumericExpression() []INumericExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INumericExpressionContext)(nil)).Elem())
	var tst = make([]INumericExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INumericExpressionContext)
		}
	}

	return tst
}

func (s *RelationalExpressionContext) NumericExpression(i int) INumericExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumericExpressionContext)(nil)).Elem(), i)

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
		p.SetState(460)
		p.NumericExpression()
	}
	p.SetState(473)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SparqlParserT__34:
		{
			p.SetState(461)
			p.Match(SparqlParserT__34)
		}
		{
			p.SetState(462)
			p.NumericExpression()
		}

	case SparqlParserT__35:
		{
			p.SetState(463)
			p.Match(SparqlParserT__35)
		}
		{
			p.SetState(464)
			p.NumericExpression()
		}

	case SparqlParserT__36:
		{
			p.SetState(465)
			p.Match(SparqlParserT__36)
		}
		{
			p.SetState(466)
			p.NumericExpression()
		}

	case SparqlParserT__37:
		{
			p.SetState(467)
			p.Match(SparqlParserT__37)
		}
		{
			p.SetState(468)
			p.NumericExpression()
		}

	case SparqlParserT__38:
		{
			p.SetState(469)
			p.Match(SparqlParserT__38)
		}
		{
			p.SetState(470)
			p.NumericExpression()
		}

	case SparqlParserT__39:
		{
			p.SetState(471)
			p.Match(SparqlParserT__39)
		}
		{
			p.SetState(472)
			p.NumericExpression()
		}

	case SparqlParserT__26, SparqlParserT__27, SparqlParserT__32, SparqlParserT__33:

	default:
	}

	return localctx
}

// INumericExpressionContext is an interface to support dynamic dispatch.
type INumericExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumericExpressionContext differentiates from other interfaces.
	IsNumericExpressionContext()
}

type NumericExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumericExpressionContext() *NumericExpressionContext {
	var p = new(NumericExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SparqlParserRULE_numericExpression
	return p
}

func (*NumericExpressionContext) IsNumericExpressionContext() {}

func NewNumericExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumericExpressionContext {
	var p = new(NumericExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_numericExpression

	return p
}

func (s *NumericExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *NumericExpressionContext) AdditiveExpression() IAdditiveExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAdditiveExpressionContext)(nil)).Elem(), 0)

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
		p.SetState(475)
		p.AdditiveExpression()
	}

	return localctx
}

// IAdditiveExpressionContext is an interface to support dynamic dispatch.
type IAdditiveExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAdditiveExpressionContext differentiates from other interfaces.
	IsAdditiveExpressionContext()
}

type AdditiveExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAdditiveExpressionContext() *AdditiveExpressionContext {
	var p = new(AdditiveExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SparqlParserRULE_additiveExpression
	return p
}

func (*AdditiveExpressionContext) IsAdditiveExpressionContext() {}

func NewAdditiveExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AdditiveExpressionContext {
	var p = new(AdditiveExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_additiveExpression

	return p
}

func (s *AdditiveExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *AdditiveExpressionContext) AllMultiplicativeExpression() []IMultiplicativeExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMultiplicativeExpressionContext)(nil)).Elem())
	var tst = make([]IMultiplicativeExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMultiplicativeExpressionContext)
		}
	}

	return tst
}

func (s *AdditiveExpressionContext) MultiplicativeExpression(i int) IMultiplicativeExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultiplicativeExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMultiplicativeExpressionContext)
}

func (s *AdditiveExpressionContext) AllNumericLiteralPositive() []INumericLiteralPositiveContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INumericLiteralPositiveContext)(nil)).Elem())
	var tst = make([]INumericLiteralPositiveContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INumericLiteralPositiveContext)
		}
	}

	return tst
}

func (s *AdditiveExpressionContext) NumericLiteralPositive(i int) INumericLiteralPositiveContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumericLiteralPositiveContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INumericLiteralPositiveContext)
}

func (s *AdditiveExpressionContext) AllNumericLiteralNegative() []INumericLiteralNegativeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INumericLiteralNegativeContext)(nil)).Elem())
	var tst = make([]INumericLiteralNegativeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INumericLiteralNegativeContext)
		}
	}

	return tst
}

func (s *AdditiveExpressionContext) NumericLiteralNegative(i int) INumericLiteralNegativeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumericLiteralNegativeContext)(nil)).Elem(), i)

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
		p.SetState(477)
		p.MultiplicativeExpression()
	}
	p.SetState(486)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la-41)&-(0x1f+1)) == 0 && ((1<<uint((_la-41)))&((1<<(SparqlParserT__40-41))|(1<<(SparqlParserT__41-41))|(1<<(SparqlParserINTEGER_POSITIVE-41))|(1<<(SparqlParserDECIMAL_POSITIVE-41))|(1<<(SparqlParserDOUBLE_POSITIVE-41))|(1<<(SparqlParserINTEGER_NEGATIVE-41)))) != 0) || _la == SparqlParserDECIMAL_NEGATIVE || _la == SparqlParserDOUBLE_NEGATIVE {
		p.SetState(484)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case SparqlParserT__40:
			{
				p.SetState(478)
				p.Match(SparqlParserT__40)
			}
			{
				p.SetState(479)
				p.MultiplicativeExpression()
			}

		case SparqlParserT__41:
			{
				p.SetState(480)
				p.Match(SparqlParserT__41)
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
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(488)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IMultiplicativeExpressionContext is an interface to support dynamic dispatch.
type IMultiplicativeExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMultiplicativeExpressionContext differentiates from other interfaces.
	IsMultiplicativeExpressionContext()
}

type MultiplicativeExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiplicativeExpressionContext() *MultiplicativeExpressionContext {
	var p = new(MultiplicativeExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SparqlParserRULE_multiplicativeExpression
	return p
}

func (*MultiplicativeExpressionContext) IsMultiplicativeExpressionContext() {}

func NewMultiplicativeExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiplicativeExpressionContext {
	var p = new(MultiplicativeExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_multiplicativeExpression

	return p
}

func (s *MultiplicativeExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiplicativeExpressionContext) AllUnaryExpression() []IUnaryExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IUnaryExpressionContext)(nil)).Elem())
	var tst = make([]IUnaryExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IUnaryExpressionContext)
		}
	}

	return tst
}

func (s *MultiplicativeExpressionContext) UnaryExpression(i int) IUnaryExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnaryExpressionContext)(nil)).Elem(), i)

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
		p.SetState(489)
		p.UnaryExpression()
	}
	p.SetState(496)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SparqlParserT__5 || _la == SparqlParserT__42 {
		p.SetState(494)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case SparqlParserT__5:
			{
				p.SetState(490)
				p.Match(SparqlParserT__5)
			}
			{
				p.SetState(491)
				p.UnaryExpression()
			}

		case SparqlParserT__42:
			{
				p.SetState(492)
				p.Match(SparqlParserT__42)
			}
			{
				p.SetState(493)
				p.UnaryExpression()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(498)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IUnaryExpressionContext is an interface to support dynamic dispatch.
type IUnaryExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnaryExpressionContext differentiates from other interfaces.
	IsUnaryExpressionContext()
}

type UnaryExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryExpressionContext() *UnaryExpressionContext {
	var p = new(UnaryExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SparqlParserRULE_unaryExpression
	return p
}

func (*UnaryExpressionContext) IsUnaryExpressionContext() {}

func NewUnaryExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryExpressionContext {
	var p = new(UnaryExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_unaryExpression

	return p
}

func (s *UnaryExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryExpressionContext) PrimaryExpression() IPrimaryExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimaryExpressionContext)(nil)).Elem(), 0)

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

	p.SetState(506)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SparqlParserT__43:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(499)
			p.Match(SparqlParserT__43)
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
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPrimaryExpressionContext is an interface to support dynamic dispatch.
type IPrimaryExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimaryExpressionContext differentiates from other interfaces.
	IsPrimaryExpressionContext()
}

type PrimaryExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryExpressionContext() *PrimaryExpressionContext {
	var p = new(PrimaryExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SparqlParserRULE_primaryExpression
	return p
}

func (*PrimaryExpressionContext) IsPrimaryExpressionContext() {}

func NewPrimaryExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryExpressionContext {
	var p = new(PrimaryExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_primaryExpression

	return p
}

func (s *PrimaryExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryExpressionContext) BrackettedExpression() IBrackettedExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBrackettedExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBrackettedExpressionContext)
}

func (s *PrimaryExpressionContext) BuiltInCall() IBuiltInCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBuiltInCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBuiltInCallContext)
}

func (s *PrimaryExpressionContext) IriRefOrFunction() IIriRefOrFunctionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIriRefOrFunctionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIriRefOrFunctionContext)
}

func (s *PrimaryExpressionContext) RdfLiteral() IRdfLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRdfLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRdfLiteralContext)
}

func (s *PrimaryExpressionContext) NumericLiteral() INumericLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumericLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumericLiteralContext)
}

func (s *PrimaryExpressionContext) BooleanLiteral() IBooleanLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBooleanLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBooleanLiteralContext)
}

func (s *PrimaryExpressionContext) Varx() IVarxContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarxContext)(nil)).Elem(), 0)

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

	p.SetState(515)
	p.GetErrorHandler().Sync(p)

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
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IBrackettedExpressionContext is an interface to support dynamic dispatch.
type IBrackettedExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBrackettedExpressionContext differentiates from other interfaces.
	IsBrackettedExpressionContext()
}

type BrackettedExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBrackettedExpressionContext() *BrackettedExpressionContext {
	var p = new(BrackettedExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SparqlParserRULE_brackettedExpression
	return p
}

func (*BrackettedExpressionContext) IsBrackettedExpressionContext() {}

func NewBrackettedExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BrackettedExpressionContext {
	var p = new(BrackettedExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_brackettedExpression

	return p
}

func (s *BrackettedExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *BrackettedExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

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
		p.SetState(517)
		p.Match(SparqlParserT__25)
	}
	{
		p.SetState(518)
		p.Expression()
	}
	{
		p.SetState(519)
		p.Match(SparqlParserT__27)
	}

	return localctx
}

// IBuiltInCallContext is an interface to support dynamic dispatch.
type IBuiltInCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBuiltInCallContext differentiates from other interfaces.
	IsBuiltInCallContext()
}

type BuiltInCallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBuiltInCallContext() *BuiltInCallContext {
	var p = new(BuiltInCallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SparqlParserRULE_builtInCall
	return p
}

func (*BuiltInCallContext) IsBuiltInCallContext() {}

func NewBuiltInCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BuiltInCallContext {
	var p = new(BuiltInCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_builtInCall

	return p
}

func (s *BuiltInCallContext) GetParser() antlr.Parser { return s.parser }

func (s *BuiltInCallContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *BuiltInCallContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BuiltInCallContext) Varx() IVarxContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarxContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarxContext)
}

func (s *BuiltInCallContext) RegexExpression() IRegexExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRegexExpressionContext)(nil)).Elem(), 0)

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

	p.SetState(576)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SparqlParserT__44:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(521)
			p.Match(SparqlParserT__44)
		}
		{
			p.SetState(522)
			p.Match(SparqlParserT__25)
		}
		{
			p.SetState(523)
			p.Expression()
		}
		{
			p.SetState(524)
			p.Match(SparqlParserT__27)
		}

	case SparqlParserT__45:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(526)
			p.Match(SparqlParserT__45)
		}
		{
			p.SetState(527)
			p.Match(SparqlParserT__25)
		}
		{
			p.SetState(528)
			p.Expression()
		}
		{
			p.SetState(529)
			p.Match(SparqlParserT__27)
		}

	case SparqlParserT__46:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(531)
			p.Match(SparqlParserT__46)
		}
		{
			p.SetState(532)
			p.Match(SparqlParserT__25)
		}
		{
			p.SetState(533)
			p.Expression()
		}
		{
			p.SetState(534)
			p.Match(SparqlParserT__26)
		}
		{
			p.SetState(535)
			p.Expression()
		}
		{
			p.SetState(536)
			p.Match(SparqlParserT__27)
		}

	case SparqlParserT__47:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(538)
			p.Match(SparqlParserT__47)
		}
		{
			p.SetState(539)
			p.Match(SparqlParserT__25)
		}
		{
			p.SetState(540)
			p.Expression()
		}
		{
			p.SetState(541)
			p.Match(SparqlParserT__27)
		}

	case SparqlParserT__48:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(543)
			p.Match(SparqlParserT__48)
		}
		{
			p.SetState(544)
			p.Match(SparqlParserT__25)
		}
		{
			p.SetState(545)
			p.Varx()
		}
		{
			p.SetState(546)
			p.Match(SparqlParserT__27)
		}

	case SparqlParserT__49:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(548)
			p.Match(SparqlParserT__49)
		}
		{
			p.SetState(549)
			p.Match(SparqlParserT__25)
		}
		{
			p.SetState(550)
			p.Expression()
		}
		{
			p.SetState(551)
			p.Match(SparqlParserT__26)
		}
		{
			p.SetState(552)
			p.Expression()
		}
		{
			p.SetState(553)
			p.Match(SparqlParserT__27)
		}

	case SparqlParserT__50:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(555)
			p.Match(SparqlParserT__50)
		}
		{
			p.SetState(556)
			p.Match(SparqlParserT__25)
		}
		{
			p.SetState(557)
			p.Expression()
		}
		{
			p.SetState(558)
			p.Match(SparqlParserT__27)
		}

	case SparqlParserT__51:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(560)
			p.Match(SparqlParserT__51)
		}
		{
			p.SetState(561)
			p.Match(SparqlParserT__25)
		}
		{
			p.SetState(562)
			p.Expression()
		}
		{
			p.SetState(563)
			p.Match(SparqlParserT__27)
		}

	case SparqlParserT__52:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(565)
			p.Match(SparqlParserT__52)
		}
		{
			p.SetState(566)
			p.Match(SparqlParserT__25)
		}
		{
			p.SetState(567)
			p.Expression()
		}
		{
			p.SetState(568)
			p.Match(SparqlParserT__27)
		}

	case SparqlParserT__53:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(570)
			p.Match(SparqlParserT__53)
		}
		{
			p.SetState(571)
			p.Match(SparqlParserT__25)
		}
		{
			p.SetState(572)
			p.Expression()
		}
		{
			p.SetState(573)
			p.Match(SparqlParserT__27)
		}

	case SparqlParserT__54:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(575)
			p.RegexExpression()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IRegexExpressionContext is an interface to support dynamic dispatch.
type IRegexExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRegexExpressionContext differentiates from other interfaces.
	IsRegexExpressionContext()
}

type RegexExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRegexExpressionContext() *RegexExpressionContext {
	var p = new(RegexExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SparqlParserRULE_regexExpression
	return p
}

func (*RegexExpressionContext) IsRegexExpressionContext() {}

func NewRegexExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RegexExpressionContext {
	var p = new(RegexExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_regexExpression

	return p
}

func (s *RegexExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *RegexExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *RegexExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

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
		p.SetState(578)
		p.Match(SparqlParserT__54)
	}
	{
		p.SetState(579)
		p.Match(SparqlParserT__25)
	}
	{
		p.SetState(580)
		p.Expression()
	}
	{
		p.SetState(581)
		p.Match(SparqlParserT__26)
	}
	{
		p.SetState(582)
		p.Expression()
	}
	p.SetState(585)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SparqlParserT__26 {
		{
			p.SetState(583)
			p.Match(SparqlParserT__26)
		}
		{
			p.SetState(584)
			p.Expression()
		}

	}
	{
		p.SetState(587)
		p.Match(SparqlParserT__27)
	}

	return localctx
}

// IIriRefOrFunctionContext is an interface to support dynamic dispatch.
type IIriRefOrFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIriRefOrFunctionContext differentiates from other interfaces.
	IsIriRefOrFunctionContext()
}

type IriRefOrFunctionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIriRefOrFunctionContext() *IriRefOrFunctionContext {
	var p = new(IriRefOrFunctionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SparqlParserRULE_iriRefOrFunction
	return p
}

func (*IriRefOrFunctionContext) IsIriRefOrFunctionContext() {}

func NewIriRefOrFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IriRefOrFunctionContext {
	var p = new(IriRefOrFunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_iriRefOrFunction

	return p
}

func (s *IriRefOrFunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *IriRefOrFunctionContext) IriRef() IIriRefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIriRefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIriRefContext)
}

func (s *IriRefOrFunctionContext) ArgList() IArgListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgListContext)(nil)).Elem(), 0)

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
		p.SetState(589)
		p.IriRef()
	}
	p.SetState(591)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SparqlParserT__25 || _la == SparqlParserNIL {
		{
			p.SetState(590)
			p.ArgList()
		}

	}

	return localctx
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

	// IsRdfLiteralContext differentiates from other interfaces.
	IsRdfLiteralContext()
}

type RdfLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	str    IRdfstringContext
	lang   antlr.Token
	iri    IIriRefContext
}

func NewEmptyRdfLiteralContext() *RdfLiteralContext {
	var p = new(RdfLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SparqlParserRULE_rdfLiteral
	return p
}

func (*RdfLiteralContext) IsRdfLiteralContext() {}

func NewRdfLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RdfLiteralContext {
	var p = new(RdfLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRdfstringContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRdfstringContext)
}

func (s *RdfLiteralContext) LANGTAG() antlr.TerminalNode {
	return s.GetToken(SparqlParserLANGTAG, 0)
}

func (s *RdfLiteralContext) IriRef() IIriRefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIriRefContext)(nil)).Elem(), 0)

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
		p.SetState(593)

		var _x = p.Rdfstring()

		localctx.(*RdfLiteralContext).str = _x
	}
	p.SetState(597)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SparqlParserLANGTAG:
		{
			p.SetState(594)

			var _m = p.Match(SparqlParserLANGTAG)

			localctx.(*RdfLiteralContext).lang = _m
		}

	case SparqlParserT__55:
		{
			p.SetState(595)
			p.Match(SparqlParserT__55)
		}
		{
			p.SetState(596)

			var _x = p.IriRef()

			localctx.(*RdfLiteralContext).iri = _x
		}

	case SparqlParserT__5, SparqlParserT__18, SparqlParserT__19, SparqlParserT__20, SparqlParserT__21, SparqlParserT__22, SparqlParserT__24, SparqlParserT__25, SparqlParserT__26, SparqlParserT__27, SparqlParserT__28, SparqlParserT__29, SparqlParserT__30, SparqlParserT__31, SparqlParserT__32, SparqlParserT__33, SparqlParserT__34, SparqlParserT__35, SparqlParserT__36, SparqlParserT__37, SparqlParserT__38, SparqlParserT__39, SparqlParserT__40, SparqlParserT__41, SparqlParserT__42, SparqlParserT__56, SparqlParserT__57, SparqlParserIRI_REF, SparqlParserPNAME_NS, SparqlParserPNAME_LN, SparqlParserBLANK_NODE_LABEL, SparqlParserVAR1, SparqlParserVAR2, SparqlParserINTEGER, SparqlParserDECIMAL, SparqlParserDOUBLE, SparqlParserINTEGER_POSITIVE, SparqlParserDECIMAL_POSITIVE, SparqlParserDOUBLE_POSITIVE, SparqlParserINTEGER_NEGATIVE, SparqlParserDECIMAL_NEGATIVE, SparqlParserDOUBLE_NEGATIVE, SparqlParserSTRING_LITERAL1, SparqlParserSTRING_LITERAL2, SparqlParserNIL, SparqlParserANON:

	default:
	}

	return localctx
}

// INumericLiteralContext is an interface to support dynamic dispatch.
type INumericLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumericLiteralContext differentiates from other interfaces.
	IsNumericLiteralContext()
}

type NumericLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumericLiteralContext() *NumericLiteralContext {
	var p = new(NumericLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SparqlParserRULE_numericLiteral
	return p
}

func (*NumericLiteralContext) IsNumericLiteralContext() {}

func NewNumericLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumericLiteralContext {
	var p = new(NumericLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_numericLiteral

	return p
}

func (s *NumericLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *NumericLiteralContext) NumericLiteralUnsigned() INumericLiteralUnsignedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumericLiteralUnsignedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumericLiteralUnsignedContext)
}

func (s *NumericLiteralContext) NumericLiteralPositive() INumericLiteralPositiveContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumericLiteralPositiveContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumericLiteralPositiveContext)
}

func (s *NumericLiteralContext) NumericLiteralNegative() INumericLiteralNegativeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumericLiteralNegativeContext)(nil)).Elem(), 0)

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

	p.SetState(602)
	p.GetErrorHandler().Sync(p)

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
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// INumericLiteralUnsignedContext is an interface to support dynamic dispatch.
type INumericLiteralUnsignedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumericLiteralUnsignedContext differentiates from other interfaces.
	IsNumericLiteralUnsignedContext()
}

type NumericLiteralUnsignedContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumericLiteralUnsignedContext() *NumericLiteralUnsignedContext {
	var p = new(NumericLiteralUnsignedContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SparqlParserRULE_numericLiteralUnsigned
	return p
}

func (*NumericLiteralUnsignedContext) IsNumericLiteralUnsignedContext() {}

func NewNumericLiteralUnsignedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumericLiteralUnsignedContext {
	var p = new(NumericLiteralUnsignedContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
		p.SetState(604)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-66)&-(0x1f+1)) == 0 && ((1<<uint((_la-66)))&((1<<(SparqlParserINTEGER-66))|(1<<(SparqlParserDECIMAL-66))|(1<<(SparqlParserDOUBLE-66)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// INumericLiteralPositiveContext is an interface to support dynamic dispatch.
type INumericLiteralPositiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumericLiteralPositiveContext differentiates from other interfaces.
	IsNumericLiteralPositiveContext()
}

type NumericLiteralPositiveContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumericLiteralPositiveContext() *NumericLiteralPositiveContext {
	var p = new(NumericLiteralPositiveContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SparqlParserRULE_numericLiteralPositive
	return p
}

func (*NumericLiteralPositiveContext) IsNumericLiteralPositiveContext() {}

func NewNumericLiteralPositiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumericLiteralPositiveContext {
	var p = new(NumericLiteralPositiveContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
		p.SetState(606)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-69)&-(0x1f+1)) == 0 && ((1<<uint((_la-69)))&((1<<(SparqlParserINTEGER_POSITIVE-69))|(1<<(SparqlParserDECIMAL_POSITIVE-69))|(1<<(SparqlParserDOUBLE_POSITIVE-69)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// INumericLiteralNegativeContext is an interface to support dynamic dispatch.
type INumericLiteralNegativeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumericLiteralNegativeContext differentiates from other interfaces.
	IsNumericLiteralNegativeContext()
}

type NumericLiteralNegativeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumericLiteralNegativeContext() *NumericLiteralNegativeContext {
	var p = new(NumericLiteralNegativeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SparqlParserRULE_numericLiteralNegative
	return p
}

func (*NumericLiteralNegativeContext) IsNumericLiteralNegativeContext() {}

func NewNumericLiteralNegativeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumericLiteralNegativeContext {
	var p = new(NumericLiteralNegativeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
		p.SetState(608)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-72)&-(0x1f+1)) == 0 && ((1<<uint((_la-72)))&((1<<(SparqlParserINTEGER_NEGATIVE-72))|(1<<(SparqlParserDECIMAL_NEGATIVE-72))|(1<<(SparqlParserDOUBLE_NEGATIVE-72)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
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
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBooleanLiteralContext() *BooleanLiteralContext {
	var p = new(BooleanLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SparqlParserRULE_booleanLiteral
	return p
}

func (*BooleanLiteralContext) IsBooleanLiteralContext() {}

func NewBooleanLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooleanLiteralContext {
	var p = new(BooleanLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
		p.SetState(610)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SparqlParserT__56 || _la == SparqlParserT__57) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IRdfstringContext is an interface to support dynamic dispatch.
type IRdfstringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRdfstringContext differentiates from other interfaces.
	IsRdfstringContext()
}

type RdfstringContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRdfstringContext() *RdfstringContext {
	var p = new(RdfstringContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SparqlParserRULE_rdfstring
	return p
}

func (*RdfstringContext) IsRdfstringContext() {}

func NewRdfstringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RdfstringContext {
	var p = new(RdfstringContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
		p.SetState(612)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SparqlParserSTRING_LITERAL1 || _la == SparqlParserSTRING_LITERAL2) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
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

	// IsIriRefContext differentiates from other interfaces.
	IsIriRefContext()
}

type IriRefContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	literal  antlr.Token
	prefixed IPrefixedNameContext
}

func NewEmptyIriRefContext() *IriRefContext {
	var p = new(IriRefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SparqlParserRULE_iriRef
	return p
}

func (*IriRefContext) IsIriRefContext() {}

func NewIriRefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IriRefContext {
	var p = new(IriRefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrefixedNameContext)(nil)).Elem(), 0)

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

	p.SetState(616)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SparqlParserIRI_REF:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(614)

			var _m = p.Match(SparqlParserIRI_REF)

			localctx.(*IriRefContext).literal = _m
		}

	case SparqlParserPNAME_NS, SparqlParserPNAME_LN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(615)

			var _x = p.PrefixedName()

			localctx.(*IriRefContext).prefixed = _x
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPrefixedNameContext is an interface to support dynamic dispatch.
type IPrefixedNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrefixedNameContext differentiates from other interfaces.
	IsPrefixedNameContext()
}

type PrefixedNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrefixedNameContext() *PrefixedNameContext {
	var p = new(PrefixedNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SparqlParserRULE_prefixedName
	return p
}

func (*PrefixedNameContext) IsPrefixedNameContext() {}

func NewPrefixedNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrefixedNameContext {
	var p = new(PrefixedNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
		p.SetState(618)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SparqlParserPNAME_NS || _la == SparqlParserPNAME_LN) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IBlankNodeContext is an interface to support dynamic dispatch.
type IBlankNodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlankNodeContext differentiates from other interfaces.
	IsBlankNodeContext()
}

type BlankNodeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlankNodeContext() *BlankNodeContext {
	var p = new(BlankNodeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SparqlParserRULE_blankNode
	return p
}

func (*BlankNodeContext) IsBlankNodeContext() {}

func NewBlankNodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlankNodeContext {
	var p = new(BlankNodeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

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
		p.SetState(620)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SparqlParserBLANK_NODE_LABEL || _la == SparqlParserANON) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
