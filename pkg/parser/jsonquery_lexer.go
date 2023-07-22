// Copyright 2022 Lekko Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated from JsonQuery.g4 by ANTLR 4.13.0. DO NOT EDIT.

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

type JsonQueryLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var JsonQueryLexerLexerStaticData struct {
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

func jsonquerylexerLexerInit() {
	staticData := &JsonQueryLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'('", "')'", "'pr'", "'.'", "'['", "']'", "", "", "", "", "'null'",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"'\\n'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "NOT", "AND_OPERATOR", "OR_OPERATOR", "BOOLEAN",
		"NULL", "IN", "EQ", "NE", "GT", "LT", "GE", "LE", "CO", "SW", "EW",
		"ATTRNAME", "VERSION", "STRING", "DOUBLE", "LONG", "INT", "EXP", "NEWLINE",
		"COMMA", "SP",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "NOT", "AND_OPERATOR",
		"OR_OPERATOR", "BOOLEAN", "NULL", "IN", "EQ", "NE", "GT", "LT", "GE",
		"LE", "CO", "SW", "EW", "ATTRNAME", "ATTR_NAME_CHAR", "DIGIT", "ALPHA",
		"VERSION", "STRING", "ESC", "UNICODE", "HEX", "DOUBLE", "LONG", "INT",
		"EXP", "NEWLINE", "COMMA", "SP",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 31, 308, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4,
		1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 96, 8, 6, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 106, 8, 7, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 114, 8, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1,
		9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 125, 8, 9, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 136, 8, 11, 1, 12, 1, 12, 1,
		12, 1, 12, 1, 12, 1, 12, 3, 12, 144, 8, 12, 1, 13, 1, 13, 1, 13, 1, 13,
		1, 13, 1, 13, 3, 13, 152, 8, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 3,
		14, 159, 8, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 166, 8, 15, 1,
		16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 3, 16, 174, 8, 16, 1, 17, 1, 17,
		1, 17, 1, 17, 1, 17, 1, 17, 3, 17, 182, 8, 17, 1, 18, 1, 18, 1, 18, 1,
		18, 3, 18, 188, 8, 18, 1, 19, 1, 19, 1, 19, 1, 19, 3, 19, 194, 8, 19, 1,
		20, 1, 20, 1, 20, 1, 20, 3, 20, 200, 8, 20, 1, 21, 1, 21, 5, 21, 204, 8,
		21, 10, 21, 12, 21, 207, 9, 21, 1, 22, 1, 22, 1, 22, 3, 22, 212, 8, 22,
		1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1,
		26, 1, 26, 1, 26, 5, 26, 227, 8, 26, 10, 26, 12, 26, 230, 9, 26, 1, 26,
		1, 26, 1, 26, 1, 26, 5, 26, 236, 8, 26, 10, 26, 12, 26, 239, 9, 26, 1,
		26, 3, 26, 242, 8, 26, 1, 27, 1, 27, 1, 27, 3, 27, 247, 8, 27, 1, 28, 1,
		28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 30, 3, 30, 258, 8, 30,
		1, 30, 1, 30, 1, 30, 4, 30, 263, 8, 30, 11, 30, 12, 30, 264, 1, 30, 3,
		30, 268, 8, 30, 1, 31, 3, 31, 271, 8, 31, 1, 31, 1, 31, 3, 31, 275, 8,
		31, 1, 32, 1, 32, 1, 32, 5, 32, 280, 8, 32, 10, 32, 12, 32, 283, 9, 32,
		3, 32, 285, 8, 32, 1, 33, 1, 33, 3, 33, 289, 8, 33, 1, 33, 1, 33, 1, 34,
		1, 34, 1, 35, 1, 35, 5, 35, 297, 8, 35, 10, 35, 12, 35, 300, 9, 35, 1,
		36, 1, 36, 5, 36, 304, 8, 36, 10, 36, 12, 36, 307, 9, 36, 0, 0, 37, 1,
		1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11,
		23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20,
		41, 21, 43, 22, 45, 0, 47, 0, 49, 0, 51, 23, 53, 24, 55, 0, 57, 0, 59,
		0, 61, 25, 63, 26, 65, 27, 67, 28, 69, 29, 71, 30, 73, 31, 1, 0, 10, 3,
		0, 45, 45, 58, 58, 95, 95, 2, 0, 65, 90, 97, 122, 2, 0, 34, 34, 92, 92,
		2, 0, 39, 39, 92, 92, 8, 0, 34, 34, 47, 47, 92, 92, 98, 98, 102, 102, 110,
		110, 114, 114, 116, 116, 3, 0, 48, 57, 65, 70, 97, 102, 1, 0, 48, 57, 1,
		0, 49, 57, 2, 0, 69, 69, 101, 101, 2, 0, 43, 43, 45, 45, 343, 0, 1, 1,
		0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1,
		0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17,
		1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0,
		25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0,
		0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0,
		0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0,
		0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1,
		0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 1, 75,
		1, 0, 0, 0, 3, 77, 1, 0, 0, 0, 5, 79, 1, 0, 0, 0, 7, 82, 1, 0, 0, 0, 9,
		84, 1, 0, 0, 0, 11, 86, 1, 0, 0, 0, 13, 95, 1, 0, 0, 0, 15, 105, 1, 0,
		0, 0, 17, 113, 1, 0, 0, 0, 19, 124, 1, 0, 0, 0, 21, 126, 1, 0, 0, 0, 23,
		135, 1, 0, 0, 0, 25, 143, 1, 0, 0, 0, 27, 151, 1, 0, 0, 0, 29, 158, 1,
		0, 0, 0, 31, 165, 1, 0, 0, 0, 33, 173, 1, 0, 0, 0, 35, 181, 1, 0, 0, 0,
		37, 187, 1, 0, 0, 0, 39, 193, 1, 0, 0, 0, 41, 199, 1, 0, 0, 0, 43, 201,
		1, 0, 0, 0, 45, 211, 1, 0, 0, 0, 47, 213, 1, 0, 0, 0, 49, 215, 1, 0, 0,
		0, 51, 217, 1, 0, 0, 0, 53, 241, 1, 0, 0, 0, 55, 243, 1, 0, 0, 0, 57, 248,
		1, 0, 0, 0, 59, 254, 1, 0, 0, 0, 61, 257, 1, 0, 0, 0, 63, 270, 1, 0, 0,
		0, 65, 284, 1, 0, 0, 0, 67, 286, 1, 0, 0, 0, 69, 292, 1, 0, 0, 0, 71, 294,
		1, 0, 0, 0, 73, 301, 1, 0, 0, 0, 75, 76, 5, 40, 0, 0, 76, 2, 1, 0, 0, 0,
		77, 78, 5, 41, 0, 0, 78, 4, 1, 0, 0, 0, 79, 80, 5, 112, 0, 0, 80, 81, 5,
		114, 0, 0, 81, 6, 1, 0, 0, 0, 82, 83, 5, 46, 0, 0, 83, 8, 1, 0, 0, 0, 84,
		85, 5, 91, 0, 0, 85, 10, 1, 0, 0, 0, 86, 87, 5, 93, 0, 0, 87, 12, 1, 0,
		0, 0, 88, 89, 5, 110, 0, 0, 89, 90, 5, 111, 0, 0, 90, 96, 5, 116, 0, 0,
		91, 92, 5, 78, 0, 0, 92, 93, 5, 79, 0, 0, 93, 96, 5, 84, 0, 0, 94, 96,
		5, 33, 0, 0, 95, 88, 1, 0, 0, 0, 95, 91, 1, 0, 0, 0, 95, 94, 1, 0, 0, 0,
		96, 14, 1, 0, 0, 0, 97, 98, 5, 97, 0, 0, 98, 99, 5, 110, 0, 0, 99, 106,
		5, 100, 0, 0, 100, 101, 5, 65, 0, 0, 101, 102, 5, 78, 0, 0, 102, 106, 5,
		68, 0, 0, 103, 104, 5, 38, 0, 0, 104, 106, 5, 38, 0, 0, 105, 97, 1, 0,
		0, 0, 105, 100, 1, 0, 0, 0, 105, 103, 1, 0, 0, 0, 106, 16, 1, 0, 0, 0,
		107, 108, 5, 111, 0, 0, 108, 114, 5, 114, 0, 0, 109, 110, 5, 79, 0, 0,
		110, 114, 5, 82, 0, 0, 111, 112, 5, 124, 0, 0, 112, 114, 5, 124, 0, 0,
		113, 107, 1, 0, 0, 0, 113, 109, 1, 0, 0, 0, 113, 111, 1, 0, 0, 0, 114,
		18, 1, 0, 0, 0, 115, 116, 5, 116, 0, 0, 116, 117, 5, 114, 0, 0, 117, 118,
		5, 117, 0, 0, 118, 125, 5, 101, 0, 0, 119, 120, 5, 102, 0, 0, 120, 121,
		5, 97, 0, 0, 121, 122, 5, 108, 0, 0, 122, 123, 5, 115, 0, 0, 123, 125,
		5, 101, 0, 0, 124, 115, 1, 0, 0, 0, 124, 119, 1, 0, 0, 0, 125, 20, 1, 0,
		0, 0, 126, 127, 5, 110, 0, 0, 127, 128, 5, 117, 0, 0, 128, 129, 5, 108,
		0, 0, 129, 130, 5, 108, 0, 0, 130, 22, 1, 0, 0, 0, 131, 132, 5, 73, 0,
		0, 132, 136, 5, 78, 0, 0, 133, 134, 5, 105, 0, 0, 134, 136, 5, 110, 0,
		0, 135, 131, 1, 0, 0, 0, 135, 133, 1, 0, 0, 0, 136, 24, 1, 0, 0, 0, 137,
		138, 5, 101, 0, 0, 138, 144, 5, 113, 0, 0, 139, 140, 5, 69, 0, 0, 140,
		144, 5, 81, 0, 0, 141, 142, 5, 61, 0, 0, 142, 144, 5, 61, 0, 0, 143, 137,
		1, 0, 0, 0, 143, 139, 1, 0, 0, 0, 143, 141, 1, 0, 0, 0, 144, 26, 1, 0,
		0, 0, 145, 146, 5, 110, 0, 0, 146, 152, 5, 101, 0, 0, 147, 148, 5, 78,
		0, 0, 148, 152, 5, 69, 0, 0, 149, 150, 5, 33, 0, 0, 150, 152, 5, 61, 0,
		0, 151, 145, 1, 0, 0, 0, 151, 147, 1, 0, 0, 0, 151, 149, 1, 0, 0, 0, 152,
		28, 1, 0, 0, 0, 153, 154, 5, 103, 0, 0, 154, 159, 5, 116, 0, 0, 155, 156,
		5, 71, 0, 0, 156, 159, 5, 84, 0, 0, 157, 159, 5, 62, 0, 0, 158, 153, 1,
		0, 0, 0, 158, 155, 1, 0, 0, 0, 158, 157, 1, 0, 0, 0, 159, 30, 1, 0, 0,
		0, 160, 161, 5, 108, 0, 0, 161, 166, 5, 116, 0, 0, 162, 163, 5, 76, 0,
		0, 163, 166, 5, 84, 0, 0, 164, 166, 5, 60, 0, 0, 165, 160, 1, 0, 0, 0,
		165, 162, 1, 0, 0, 0, 165, 164, 1, 0, 0, 0, 166, 32, 1, 0, 0, 0, 167, 168,
		5, 103, 0, 0, 168, 174, 5, 101, 0, 0, 169, 170, 5, 71, 0, 0, 170, 174,
		5, 69, 0, 0, 171, 172, 5, 62, 0, 0, 172, 174, 5, 61, 0, 0, 173, 167, 1,
		0, 0, 0, 173, 169, 1, 0, 0, 0, 173, 171, 1, 0, 0, 0, 174, 34, 1, 0, 0,
		0, 175, 176, 5, 108, 0, 0, 176, 182, 5, 101, 0, 0, 177, 178, 5, 76, 0,
		0, 178, 182, 5, 69, 0, 0, 179, 180, 5, 60, 0, 0, 180, 182, 5, 61, 0, 0,
		181, 175, 1, 0, 0, 0, 181, 177, 1, 0, 0, 0, 181, 179, 1, 0, 0, 0, 182,
		36, 1, 0, 0, 0, 183, 184, 5, 99, 0, 0, 184, 188, 5, 111, 0, 0, 185, 186,
		5, 67, 0, 0, 186, 188, 5, 79, 0, 0, 187, 183, 1, 0, 0, 0, 187, 185, 1,
		0, 0, 0, 188, 38, 1, 0, 0, 0, 189, 190, 5, 115, 0, 0, 190, 194, 5, 119,
		0, 0, 191, 192, 5, 83, 0, 0, 192, 194, 5, 87, 0, 0, 193, 189, 1, 0, 0,
		0, 193, 191, 1, 0, 0, 0, 194, 40, 1, 0, 0, 0, 195, 196, 5, 101, 0, 0, 196,
		200, 5, 119, 0, 0, 197, 198, 5, 69, 0, 0, 198, 200, 5, 87, 0, 0, 199, 195,
		1, 0, 0, 0, 199, 197, 1, 0, 0, 0, 200, 42, 1, 0, 0, 0, 201, 205, 3, 49,
		24, 0, 202, 204, 3, 45, 22, 0, 203, 202, 1, 0, 0, 0, 204, 207, 1, 0, 0,
		0, 205, 203, 1, 0, 0, 0, 205, 206, 1, 0, 0, 0, 206, 44, 1, 0, 0, 0, 207,
		205, 1, 0, 0, 0, 208, 212, 7, 0, 0, 0, 209, 212, 3, 47, 23, 0, 210, 212,
		3, 49, 24, 0, 211, 208, 1, 0, 0, 0, 211, 209, 1, 0, 0, 0, 211, 210, 1,
		0, 0, 0, 212, 46, 1, 0, 0, 0, 213, 214, 2, 48, 57, 0, 214, 48, 1, 0, 0,
		0, 215, 216, 7, 1, 0, 0, 216, 50, 1, 0, 0, 0, 217, 218, 3, 65, 32, 0, 218,
		219, 5, 46, 0, 0, 219, 220, 3, 65, 32, 0, 220, 221, 5, 46, 0, 0, 221, 222,
		3, 65, 32, 0, 222, 52, 1, 0, 0, 0, 223, 228, 5, 34, 0, 0, 224, 227, 3,
		55, 27, 0, 225, 227, 8, 2, 0, 0, 226, 224, 1, 0, 0, 0, 226, 225, 1, 0,
		0, 0, 227, 230, 1, 0, 0, 0, 228, 226, 1, 0, 0, 0, 228, 229, 1, 0, 0, 0,
		229, 231, 1, 0, 0, 0, 230, 228, 1, 0, 0, 0, 231, 242, 5, 34, 0, 0, 232,
		237, 5, 39, 0, 0, 233, 236, 3, 55, 27, 0, 234, 236, 8, 3, 0, 0, 235, 233,
		1, 0, 0, 0, 235, 234, 1, 0, 0, 0, 236, 239, 1, 0, 0, 0, 237, 235, 1, 0,
		0, 0, 237, 238, 1, 0, 0, 0, 238, 240, 1, 0, 0, 0, 239, 237, 1, 0, 0, 0,
		240, 242, 5, 39, 0, 0, 241, 223, 1, 0, 0, 0, 241, 232, 1, 0, 0, 0, 242,
		54, 1, 0, 0, 0, 243, 246, 5, 92, 0, 0, 244, 247, 7, 4, 0, 0, 245, 247,
		3, 57, 28, 0, 246, 244, 1, 0, 0, 0, 246, 245, 1, 0, 0, 0, 247, 56, 1, 0,
		0, 0, 248, 249, 5, 117, 0, 0, 249, 250, 3, 59, 29, 0, 250, 251, 3, 59,
		29, 0, 251, 252, 3, 59, 29, 0, 252, 253, 3, 59, 29, 0, 253, 58, 1, 0, 0,
		0, 254, 255, 7, 5, 0, 0, 255, 60, 1, 0, 0, 0, 256, 258, 5, 45, 0, 0, 257,
		256, 1, 0, 0, 0, 257, 258, 1, 0, 0, 0, 258, 259, 1, 0, 0, 0, 259, 260,
		3, 65, 32, 0, 260, 262, 5, 46, 0, 0, 261, 263, 7, 6, 0, 0, 262, 261, 1,
		0, 0, 0, 263, 264, 1, 0, 0, 0, 264, 262, 1, 0, 0, 0, 264, 265, 1, 0, 0,
		0, 265, 267, 1, 0, 0, 0, 266, 268, 3, 67, 33, 0, 267, 266, 1, 0, 0, 0,
		267, 268, 1, 0, 0, 0, 268, 62, 1, 0, 0, 0, 269, 271, 5, 45, 0, 0, 270,
		269, 1, 0, 0, 0, 270, 271, 1, 0, 0, 0, 271, 272, 1, 0, 0, 0, 272, 274,
		3, 65, 32, 0, 273, 275, 3, 67, 33, 0, 274, 273, 1, 0, 0, 0, 274, 275, 1,
		0, 0, 0, 275, 64, 1, 0, 0, 0, 276, 285, 5, 48, 0, 0, 277, 281, 7, 7, 0,
		0, 278, 280, 7, 6, 0, 0, 279, 278, 1, 0, 0, 0, 280, 283, 1, 0, 0, 0, 281,
		279, 1, 0, 0, 0, 281, 282, 1, 0, 0, 0, 282, 285, 1, 0, 0, 0, 283, 281,
		1, 0, 0, 0, 284, 276, 1, 0, 0, 0, 284, 277, 1, 0, 0, 0, 285, 66, 1, 0,
		0, 0, 286, 288, 7, 8, 0, 0, 287, 289, 7, 9, 0, 0, 288, 287, 1, 0, 0, 0,
		288, 289, 1, 0, 0, 0, 289, 290, 1, 0, 0, 0, 290, 291, 3, 65, 32, 0, 291,
		68, 1, 0, 0, 0, 292, 293, 5, 10, 0, 0, 293, 70, 1, 0, 0, 0, 294, 298, 5,
		44, 0, 0, 295, 297, 5, 32, 0, 0, 296, 295, 1, 0, 0, 0, 297, 300, 1, 0,
		0, 0, 298, 296, 1, 0, 0, 0, 298, 299, 1, 0, 0, 0, 299, 72, 1, 0, 0, 0,
		300, 298, 1, 0, 0, 0, 301, 305, 5, 32, 0, 0, 302, 304, 3, 69, 34, 0, 303,
		302, 1, 0, 0, 0, 304, 307, 1, 0, 0, 0, 305, 303, 1, 0, 0, 0, 305, 306,
		1, 0, 0, 0, 306, 74, 1, 0, 0, 0, 307, 305, 1, 0, 0, 0, 33, 0, 95, 105,
		113, 124, 135, 143, 151, 158, 165, 173, 181, 187, 193, 199, 205, 211, 226,
		228, 235, 237, 241, 246, 257, 264, 267, 270, 274, 281, 284, 288, 298, 305,
		0,
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

// JsonQueryLexerInit initializes any static state used to implement JsonQueryLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewJsonQueryLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func JsonQueryLexerInit() {
	staticData := &JsonQueryLexerLexerStaticData
	staticData.once.Do(jsonquerylexerLexerInit)
}

// NewJsonQueryLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewJsonQueryLexer(input antlr.CharStream) *JsonQueryLexer {
	JsonQueryLexerInit()
	l := new(JsonQueryLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &JsonQueryLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "JsonQuery.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// JsonQueryLexer tokens.
const (
	JsonQueryLexerT__0         = 1
	JsonQueryLexerT__1         = 2
	JsonQueryLexerT__2         = 3
	JsonQueryLexerT__3         = 4
	JsonQueryLexerT__4         = 5
	JsonQueryLexerT__5         = 6
	JsonQueryLexerNOT          = 7
	JsonQueryLexerAND_OPERATOR = 8
	JsonQueryLexerOR_OPERATOR  = 9
	JsonQueryLexerBOOLEAN      = 10
	JsonQueryLexerNULL         = 11
	JsonQueryLexerIN           = 12
	JsonQueryLexerEQ           = 13
	JsonQueryLexerNE           = 14
	JsonQueryLexerGT           = 15
	JsonQueryLexerLT           = 16
	JsonQueryLexerGE           = 17
	JsonQueryLexerLE           = 18
	JsonQueryLexerCO           = 19
	JsonQueryLexerSW           = 20
	JsonQueryLexerEW           = 21
	JsonQueryLexerATTRNAME     = 22
	JsonQueryLexerVERSION      = 23
	JsonQueryLexerSTRING       = 24
	JsonQueryLexerDOUBLE       = 25
	JsonQueryLexerLONG         = 26
	JsonQueryLexerINT          = 27
	JsonQueryLexerEXP          = 28
	JsonQueryLexerNEWLINE      = 29
	JsonQueryLexerCOMMA        = 30
	JsonQueryLexerSP           = 31
)
