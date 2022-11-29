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

// Code generated from JsonQuery.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
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

var jsonquerylexerLexerStaticData struct {
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

func jsonquerylexerLexerInit() {
	staticData := &jsonquerylexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "'('", "')'", "'pr'", "'.'", "'-'", "'['", "']'", "", "", "", "'null'",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "'\\n'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "NOT", "LOGICAL_OPERATOR", "BOOLEAN",
		"NULL", "IN", "EQ", "NE", "GT", "LT", "GE", "LE", "CO", "SW", "EW",
		"ATTRNAME", "VERSION", "STRING", "DOUBLE", "INT", "EXP", "NEWLINE",
		"COMMA", "SP",
	}
	staticData.ruleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "NOT", "LOGICAL_OPERATOR",
		"BOOLEAN", "NULL", "IN", "EQ", "NE", "GT", "LT", "GE", "LE", "CO", "SW",
		"EW", "ATTRNAME", "ATTR_NAME_CHAR", "DIGIT", "ALPHA", "VERSION", "STRING",
		"ESC", "UNICODE", "HEX", "DOUBLE", "INT", "EXP", "NEWLINE", "COMMA",
		"SP",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 30, 279, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 1, 0,
		1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5,
		1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 95, 8, 7, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 8, 3, 8, 102, 8, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1,
		9, 1, 9, 1, 9, 1, 9, 3, 9, 113, 8, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 124, 8, 11, 1, 12, 1, 12, 1, 12, 1,
		12, 1, 12, 1, 12, 3, 12, 132, 8, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13,
		1, 13, 3, 13, 140, 8, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 147,
		8, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 154, 8, 15, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 16, 1, 16, 3, 16, 162, 8, 16, 1, 17, 1, 17, 1, 17,
		1, 17, 1, 17, 1, 17, 3, 17, 170, 8, 17, 1, 18, 1, 18, 1, 18, 1, 18, 3,
		18, 176, 8, 18, 1, 19, 1, 19, 1, 19, 1, 19, 3, 19, 182, 8, 19, 1, 20, 1,
		20, 1, 20, 1, 20, 3, 20, 188, 8, 20, 1, 21, 1, 21, 5, 21, 192, 8, 21, 10,
		21, 12, 21, 195, 9, 21, 1, 22, 1, 22, 1, 22, 3, 22, 200, 8, 22, 1, 23,
		1, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 26, 1,
		26, 1, 26, 5, 26, 215, 8, 26, 10, 26, 12, 26, 218, 9, 26, 1, 26, 1, 26,
		1, 27, 1, 27, 1, 27, 3, 27, 225, 8, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1,
		28, 1, 28, 1, 29, 1, 29, 1, 30, 3, 30, 236, 8, 30, 1, 30, 1, 30, 1, 30,
		4, 30, 241, 8, 30, 11, 30, 12, 30, 242, 1, 30, 3, 30, 246, 8, 30, 1, 31,
		1, 31, 1, 31, 5, 31, 251, 8, 31, 10, 31, 12, 31, 254, 9, 31, 3, 31, 256,
		8, 31, 1, 32, 1, 32, 3, 32, 260, 8, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1,
		34, 1, 34, 5, 34, 268, 8, 34, 10, 34, 12, 34, 271, 9, 34, 1, 35, 1, 35,
		5, 35, 275, 8, 35, 10, 35, 12, 35, 278, 9, 35, 0, 0, 36, 1, 1, 3, 2, 5,
		3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25,
		13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43,
		22, 45, 0, 47, 0, 49, 0, 51, 23, 53, 24, 55, 0, 57, 0, 59, 0, 61, 25, 63,
		26, 65, 27, 67, 28, 69, 29, 71, 30, 1, 0, 9, 3, 0, 45, 45, 58, 58, 95,
		95, 2, 0, 65, 90, 97, 122, 2, 0, 34, 34, 92, 92, 8, 0, 34, 34, 47, 47,
		92, 92, 98, 98, 102, 102, 110, 110, 114, 114, 116, 116, 3, 0, 48, 57, 65,
		70, 97, 102, 1, 0, 48, 57, 1, 0, 49, 57, 2, 0, 69, 69, 101, 101, 2, 0,
		43, 43, 45, 45, 305, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0,
		0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0,
		0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1,
		0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29,
		1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0,
		37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0,
		0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0,
		0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0,
		0, 0, 1, 73, 1, 0, 0, 0, 3, 75, 1, 0, 0, 0, 5, 77, 1, 0, 0, 0, 7, 80, 1,
		0, 0, 0, 9, 82, 1, 0, 0, 0, 11, 84, 1, 0, 0, 0, 13, 86, 1, 0, 0, 0, 15,
		94, 1, 0, 0, 0, 17, 101, 1, 0, 0, 0, 19, 112, 1, 0, 0, 0, 21, 114, 1, 0,
		0, 0, 23, 123, 1, 0, 0, 0, 25, 131, 1, 0, 0, 0, 27, 139, 1, 0, 0, 0, 29,
		146, 1, 0, 0, 0, 31, 153, 1, 0, 0, 0, 33, 161, 1, 0, 0, 0, 35, 169, 1,
		0, 0, 0, 37, 175, 1, 0, 0, 0, 39, 181, 1, 0, 0, 0, 41, 187, 1, 0, 0, 0,
		43, 189, 1, 0, 0, 0, 45, 199, 1, 0, 0, 0, 47, 201, 1, 0, 0, 0, 49, 203,
		1, 0, 0, 0, 51, 205, 1, 0, 0, 0, 53, 211, 1, 0, 0, 0, 55, 221, 1, 0, 0,
		0, 57, 226, 1, 0, 0, 0, 59, 232, 1, 0, 0, 0, 61, 235, 1, 0, 0, 0, 63, 255,
		1, 0, 0, 0, 65, 257, 1, 0, 0, 0, 67, 263, 1, 0, 0, 0, 69, 265, 1, 0, 0,
		0, 71, 272, 1, 0, 0, 0, 73, 74, 5, 40, 0, 0, 74, 2, 1, 0, 0, 0, 75, 76,
		5, 41, 0, 0, 76, 4, 1, 0, 0, 0, 77, 78, 5, 112, 0, 0, 78, 79, 5, 114, 0,
		0, 79, 6, 1, 0, 0, 0, 80, 81, 5, 46, 0, 0, 81, 8, 1, 0, 0, 0, 82, 83, 5,
		45, 0, 0, 83, 10, 1, 0, 0, 0, 84, 85, 5, 91, 0, 0, 85, 12, 1, 0, 0, 0,
		86, 87, 5, 93, 0, 0, 87, 14, 1, 0, 0, 0, 88, 89, 5, 110, 0, 0, 89, 90,
		5, 111, 0, 0, 90, 95, 5, 116, 0, 0, 91, 92, 5, 78, 0, 0, 92, 93, 5, 79,
		0, 0, 93, 95, 5, 84, 0, 0, 94, 88, 1, 0, 0, 0, 94, 91, 1, 0, 0, 0, 95,
		16, 1, 0, 0, 0, 96, 97, 5, 97, 0, 0, 97, 98, 5, 110, 0, 0, 98, 102, 5,
		100, 0, 0, 99, 100, 5, 111, 0, 0, 100, 102, 5, 114, 0, 0, 101, 96, 1, 0,
		0, 0, 101, 99, 1, 0, 0, 0, 102, 18, 1, 0, 0, 0, 103, 104, 5, 116, 0, 0,
		104, 105, 5, 114, 0, 0, 105, 106, 5, 117, 0, 0, 106, 113, 5, 101, 0, 0,
		107, 108, 5, 102, 0, 0, 108, 109, 5, 97, 0, 0, 109, 110, 5, 108, 0, 0,
		110, 111, 5, 115, 0, 0, 111, 113, 5, 101, 0, 0, 112, 103, 1, 0, 0, 0, 112,
		107, 1, 0, 0, 0, 113, 20, 1, 0, 0, 0, 114, 115, 5, 110, 0, 0, 115, 116,
		5, 117, 0, 0, 116, 117, 5, 108, 0, 0, 117, 118, 5, 108, 0, 0, 118, 22,
		1, 0, 0, 0, 119, 120, 5, 73, 0, 0, 120, 124, 5, 78, 0, 0, 121, 122, 5,
		105, 0, 0, 122, 124, 5, 110, 0, 0, 123, 119, 1, 0, 0, 0, 123, 121, 1, 0,
		0, 0, 124, 24, 1, 0, 0, 0, 125, 126, 5, 101, 0, 0, 126, 132, 5, 113, 0,
		0, 127, 128, 5, 69, 0, 0, 128, 132, 5, 81, 0, 0, 129, 130, 5, 61, 0, 0,
		130, 132, 5, 61, 0, 0, 131, 125, 1, 0, 0, 0, 131, 127, 1, 0, 0, 0, 131,
		129, 1, 0, 0, 0, 132, 26, 1, 0, 0, 0, 133, 134, 5, 110, 0, 0, 134, 140,
		5, 101, 0, 0, 135, 136, 5, 78, 0, 0, 136, 140, 5, 69, 0, 0, 137, 138, 5,
		33, 0, 0, 138, 140, 5, 61, 0, 0, 139, 133, 1, 0, 0, 0, 139, 135, 1, 0,
		0, 0, 139, 137, 1, 0, 0, 0, 140, 28, 1, 0, 0, 0, 141, 142, 5, 103, 0, 0,
		142, 147, 5, 116, 0, 0, 143, 144, 5, 71, 0, 0, 144, 147, 5, 84, 0, 0, 145,
		147, 5, 62, 0, 0, 146, 141, 1, 0, 0, 0, 146, 143, 1, 0, 0, 0, 146, 145,
		1, 0, 0, 0, 147, 30, 1, 0, 0, 0, 148, 149, 5, 108, 0, 0, 149, 154, 5, 116,
		0, 0, 150, 151, 5, 76, 0, 0, 151, 154, 5, 84, 0, 0, 152, 154, 5, 60, 0,
		0, 153, 148, 1, 0, 0, 0, 153, 150, 1, 0, 0, 0, 153, 152, 1, 0, 0, 0, 154,
		32, 1, 0, 0, 0, 155, 156, 5, 103, 0, 0, 156, 162, 5, 101, 0, 0, 157, 158,
		5, 71, 0, 0, 158, 162, 5, 69, 0, 0, 159, 160, 5, 62, 0, 0, 160, 162, 5,
		61, 0, 0, 161, 155, 1, 0, 0, 0, 161, 157, 1, 0, 0, 0, 161, 159, 1, 0, 0,
		0, 162, 34, 1, 0, 0, 0, 163, 164, 5, 108, 0, 0, 164, 170, 5, 101, 0, 0,
		165, 166, 5, 76, 0, 0, 166, 170, 5, 69, 0, 0, 167, 168, 5, 60, 0, 0, 168,
		170, 5, 61, 0, 0, 169, 163, 1, 0, 0, 0, 169, 165, 1, 0, 0, 0, 169, 167,
		1, 0, 0, 0, 170, 36, 1, 0, 0, 0, 171, 172, 5, 99, 0, 0, 172, 176, 5, 111,
		0, 0, 173, 174, 5, 67, 0, 0, 174, 176, 5, 79, 0, 0, 175, 171, 1, 0, 0,
		0, 175, 173, 1, 0, 0, 0, 176, 38, 1, 0, 0, 0, 177, 178, 5, 115, 0, 0, 178,
		182, 5, 119, 0, 0, 179, 180, 5, 83, 0, 0, 180, 182, 5, 87, 0, 0, 181, 177,
		1, 0, 0, 0, 181, 179, 1, 0, 0, 0, 182, 40, 1, 0, 0, 0, 183, 184, 5, 101,
		0, 0, 184, 188, 5, 119, 0, 0, 185, 186, 5, 69, 0, 0, 186, 188, 5, 87, 0,
		0, 187, 183, 1, 0, 0, 0, 187, 185, 1, 0, 0, 0, 188, 42, 1, 0, 0, 0, 189,
		193, 3, 49, 24, 0, 190, 192, 3, 45, 22, 0, 191, 190, 1, 0, 0, 0, 192, 195,
		1, 0, 0, 0, 193, 191, 1, 0, 0, 0, 193, 194, 1, 0, 0, 0, 194, 44, 1, 0,
		0, 0, 195, 193, 1, 0, 0, 0, 196, 200, 7, 0, 0, 0, 197, 200, 3, 47, 23,
		0, 198, 200, 3, 49, 24, 0, 199, 196, 1, 0, 0, 0, 199, 197, 1, 0, 0, 0,
		199, 198, 1, 0, 0, 0, 200, 46, 1, 0, 0, 0, 201, 202, 2, 48, 57, 0, 202,
		48, 1, 0, 0, 0, 203, 204, 7, 1, 0, 0, 204, 50, 1, 0, 0, 0, 205, 206, 3,
		63, 31, 0, 206, 207, 5, 46, 0, 0, 207, 208, 3, 63, 31, 0, 208, 209, 5,
		46, 0, 0, 209, 210, 3, 63, 31, 0, 210, 52, 1, 0, 0, 0, 211, 216, 5, 34,
		0, 0, 212, 215, 3, 55, 27, 0, 213, 215, 8, 2, 0, 0, 214, 212, 1, 0, 0,
		0, 214, 213, 1, 0, 0, 0, 215, 218, 1, 0, 0, 0, 216, 214, 1, 0, 0, 0, 216,
		217, 1, 0, 0, 0, 217, 219, 1, 0, 0, 0, 218, 216, 1, 0, 0, 0, 219, 220,
		5, 34, 0, 0, 220, 54, 1, 0, 0, 0, 221, 224, 5, 92, 0, 0, 222, 225, 7, 3,
		0, 0, 223, 225, 3, 57, 28, 0, 224, 222, 1, 0, 0, 0, 224, 223, 1, 0, 0,
		0, 225, 56, 1, 0, 0, 0, 226, 227, 5, 117, 0, 0, 227, 228, 3, 59, 29, 0,
		228, 229, 3, 59, 29, 0, 229, 230, 3, 59, 29, 0, 230, 231, 3, 59, 29, 0,
		231, 58, 1, 0, 0, 0, 232, 233, 7, 4, 0, 0, 233, 60, 1, 0, 0, 0, 234, 236,
		5, 45, 0, 0, 235, 234, 1, 0, 0, 0, 235, 236, 1, 0, 0, 0, 236, 237, 1, 0,
		0, 0, 237, 238, 3, 63, 31, 0, 238, 240, 5, 46, 0, 0, 239, 241, 7, 5, 0,
		0, 240, 239, 1, 0, 0, 0, 241, 242, 1, 0, 0, 0, 242, 240, 1, 0, 0, 0, 242,
		243, 1, 0, 0, 0, 243, 245, 1, 0, 0, 0, 244, 246, 3, 65, 32, 0, 245, 244,
		1, 0, 0, 0, 245, 246, 1, 0, 0, 0, 246, 62, 1, 0, 0, 0, 247, 256, 5, 48,
		0, 0, 248, 252, 7, 6, 0, 0, 249, 251, 7, 5, 0, 0, 250, 249, 1, 0, 0, 0,
		251, 254, 1, 0, 0, 0, 252, 250, 1, 0, 0, 0, 252, 253, 1, 0, 0, 0, 253,
		256, 1, 0, 0, 0, 254, 252, 1, 0, 0, 0, 255, 247, 1, 0, 0, 0, 255, 248,
		1, 0, 0, 0, 256, 64, 1, 0, 0, 0, 257, 259, 7, 7, 0, 0, 258, 260, 7, 8,
		0, 0, 259, 258, 1, 0, 0, 0, 259, 260, 1, 0, 0, 0, 260, 261, 1, 0, 0, 0,
		261, 262, 3, 63, 31, 0, 262, 66, 1, 0, 0, 0, 263, 264, 5, 10, 0, 0, 264,
		68, 1, 0, 0, 0, 265, 269, 5, 44, 0, 0, 266, 268, 5, 32, 0, 0, 267, 266,
		1, 0, 0, 0, 268, 271, 1, 0, 0, 0, 269, 267, 1, 0, 0, 0, 269, 270, 1, 0,
		0, 0, 270, 70, 1, 0, 0, 0, 271, 269, 1, 0, 0, 0, 272, 276, 5, 32, 0, 0,
		273, 275, 3, 67, 33, 0, 274, 273, 1, 0, 0, 0, 275, 278, 1, 0, 0, 0, 276,
		274, 1, 0, 0, 0, 276, 277, 1, 0, 0, 0, 277, 72, 1, 0, 0, 0, 278, 276, 1,
		0, 0, 0, 27, 0, 94, 101, 112, 123, 131, 139, 146, 153, 161, 169, 175, 181,
		187, 193, 199, 214, 216, 224, 235, 242, 245, 252, 255, 259, 269, 276, 0,
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
	staticData := &jsonquerylexerLexerStaticData
	staticData.once.Do(jsonquerylexerLexerInit)
}

// NewJsonQueryLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewJsonQueryLexer(input antlr.CharStream) *JsonQueryLexer {
	JsonQueryLexerInit()
	l := new(JsonQueryLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &jsonquerylexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "JsonQuery.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// JsonQueryLexer tokens.
const (
	JsonQueryLexerT__0             = 1
	JsonQueryLexerT__1             = 2
	JsonQueryLexerT__2             = 3
	JsonQueryLexerT__3             = 4
	JsonQueryLexerT__4             = 5
	JsonQueryLexerT__5             = 6
	JsonQueryLexerT__6             = 7
	JsonQueryLexerNOT              = 8
	JsonQueryLexerLOGICAL_OPERATOR = 9
	JsonQueryLexerBOOLEAN          = 10
	JsonQueryLexerNULL             = 11
	JsonQueryLexerIN               = 12
	JsonQueryLexerEQ               = 13
	JsonQueryLexerNE               = 14
	JsonQueryLexerGT               = 15
	JsonQueryLexerLT               = 16
	JsonQueryLexerGE               = 17
	JsonQueryLexerLE               = 18
	JsonQueryLexerCO               = 19
	JsonQueryLexerSW               = 20
	JsonQueryLexerEW               = 21
	JsonQueryLexerATTRNAME         = 22
	JsonQueryLexerVERSION          = 23
	JsonQueryLexerSTRING           = 24
	JsonQueryLexerDOUBLE           = 25
	JsonQueryLexerINT              = 26
	JsonQueryLexerEXP              = 27
	JsonQueryLexerNEWLINE          = 28
	JsonQueryLexerCOMMA            = 29
	JsonQueryLexerSP               = 30
)