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

// Code generated from ./pkg/parser/JsonQuery.g4 by ANTLR 4.12.0. DO NOT EDIT.

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
		"", "'('", "')'", "'pr'", "'.'", "'-'", "'['", "']'", "", "", "", "",
		"'null'", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "'\\n'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "NOT", "AND_OPERATOR", "OR_OPERATOR",
		"BOOLEAN", "NULL", "IN", "EQ", "NE", "GT", "LT", "GE", "LE", "CO", "SW",
		"EW", "ATTRNAME", "VERSION", "STRING", "DOUBLE", "INT", "EXP", "NEWLINE",
		"COMMA", "SP",
	}
	staticData.ruleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "NOT", "AND_OPERATOR",
		"OR_OPERATOR", "BOOLEAN", "NULL", "IN", "EQ", "NE", "GT", "LT", "GE",
		"LE", "CO", "SW", "EW", "ATTRNAME", "ATTR_NAME_CHAR", "DIGIT", "ALPHA",
		"VERSION", "STRING", "ESC", "UNICODE", "HEX", "DOUBLE", "INT", "EXP",
		"NEWLINE", "COMMA", "SP",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 31, 303, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4,
		1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7,
		98, 8, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 108, 8,
		8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 116, 8, 9, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 127, 8, 10, 1, 11,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 138, 8,
		12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 3, 13, 146, 8, 13, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 154, 8, 14, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 15, 3, 15, 161, 8, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16,
		3, 16, 168, 8, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 3, 17, 176,
		8, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 3, 18, 184, 8, 18, 1,
		19, 1, 19, 1, 19, 1, 19, 3, 19, 190, 8, 19, 1, 20, 1, 20, 1, 20, 1, 20,
		3, 20, 196, 8, 20, 1, 21, 1, 21, 1, 21, 1, 21, 3, 21, 202, 8, 21, 1, 22,
		1, 22, 5, 22, 206, 8, 22, 10, 22, 12, 22, 209, 9, 22, 1, 23, 1, 23, 1,
		23, 3, 23, 214, 8, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26,
		1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 5, 27, 229, 8, 27, 10, 27, 12,
		27, 232, 9, 27, 1, 27, 1, 27, 1, 27, 1, 27, 5, 27, 238, 8, 27, 10, 27,
		12, 27, 241, 9, 27, 1, 27, 3, 27, 244, 8, 27, 1, 28, 1, 28, 1, 28, 3, 28,
		249, 8, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1,
		31, 3, 31, 260, 8, 31, 1, 31, 1, 31, 1, 31, 4, 31, 265, 8, 31, 11, 31,
		12, 31, 266, 1, 31, 3, 31, 270, 8, 31, 1, 32, 1, 32, 1, 32, 5, 32, 275,
		8, 32, 10, 32, 12, 32, 278, 9, 32, 3, 32, 280, 8, 32, 1, 33, 1, 33, 3,
		33, 284, 8, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 35, 1, 35, 5, 35, 292, 8,
		35, 10, 35, 12, 35, 295, 9, 35, 1, 36, 1, 36, 5, 36, 299, 8, 36, 10, 36,
		12, 36, 302, 9, 36, 0, 0, 37, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13,
		7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16,
		33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 0, 49, 0, 51,
		0, 53, 24, 55, 25, 57, 0, 59, 0, 61, 0, 63, 26, 65, 27, 67, 28, 69, 29,
		71, 30, 73, 31, 1, 0, 10, 3, 0, 45, 45, 58, 58, 95, 95, 2, 0, 65, 90, 97,
		122, 2, 0, 34, 34, 92, 92, 2, 0, 39, 39, 92, 92, 8, 0, 34, 34, 47, 47,
		92, 92, 98, 98, 102, 102, 110, 110, 114, 114, 116, 116, 3, 0, 48, 57, 65,
		70, 97, 102, 1, 0, 48, 57, 1, 0, 49, 57, 2, 0, 69, 69, 101, 101, 2, 0,
		43, 43, 45, 45, 336, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0,
		0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0,
		0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1,
		0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29,
		1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0,
		37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0,
		0, 45, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 63, 1, 0, 0,
		0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0,
		0, 0, 0, 73, 1, 0, 0, 0, 1, 75, 1, 0, 0, 0, 3, 77, 1, 0, 0, 0, 5, 79, 1,
		0, 0, 0, 7, 82, 1, 0, 0, 0, 9, 84, 1, 0, 0, 0, 11, 86, 1, 0, 0, 0, 13,
		88, 1, 0, 0, 0, 15, 97, 1, 0, 0, 0, 17, 107, 1, 0, 0, 0, 19, 115, 1, 0,
		0, 0, 21, 126, 1, 0, 0, 0, 23, 128, 1, 0, 0, 0, 25, 137, 1, 0, 0, 0, 27,
		145, 1, 0, 0, 0, 29, 153, 1, 0, 0, 0, 31, 160, 1, 0, 0, 0, 33, 167, 1,
		0, 0, 0, 35, 175, 1, 0, 0, 0, 37, 183, 1, 0, 0, 0, 39, 189, 1, 0, 0, 0,
		41, 195, 1, 0, 0, 0, 43, 201, 1, 0, 0, 0, 45, 203, 1, 0, 0, 0, 47, 213,
		1, 0, 0, 0, 49, 215, 1, 0, 0, 0, 51, 217, 1, 0, 0, 0, 53, 219, 1, 0, 0,
		0, 55, 243, 1, 0, 0, 0, 57, 245, 1, 0, 0, 0, 59, 250, 1, 0, 0, 0, 61, 256,
		1, 0, 0, 0, 63, 259, 1, 0, 0, 0, 65, 279, 1, 0, 0, 0, 67, 281, 1, 0, 0,
		0, 69, 287, 1, 0, 0, 0, 71, 289, 1, 0, 0, 0, 73, 296, 1, 0, 0, 0, 75, 76,
		5, 40, 0, 0, 76, 2, 1, 0, 0, 0, 77, 78, 5, 41, 0, 0, 78, 4, 1, 0, 0, 0,
		79, 80, 5, 112, 0, 0, 80, 81, 5, 114, 0, 0, 81, 6, 1, 0, 0, 0, 82, 83,
		5, 46, 0, 0, 83, 8, 1, 0, 0, 0, 84, 85, 5, 45, 0, 0, 85, 10, 1, 0, 0, 0,
		86, 87, 5, 91, 0, 0, 87, 12, 1, 0, 0, 0, 88, 89, 5, 93, 0, 0, 89, 14, 1,
		0, 0, 0, 90, 91, 5, 110, 0, 0, 91, 92, 5, 111, 0, 0, 92, 98, 5, 116, 0,
		0, 93, 94, 5, 78, 0, 0, 94, 95, 5, 79, 0, 0, 95, 98, 5, 84, 0, 0, 96, 98,
		5, 33, 0, 0, 97, 90, 1, 0, 0, 0, 97, 93, 1, 0, 0, 0, 97, 96, 1, 0, 0, 0,
		98, 16, 1, 0, 0, 0, 99, 100, 5, 97, 0, 0, 100, 101, 5, 110, 0, 0, 101,
		108, 5, 100, 0, 0, 102, 103, 5, 65, 0, 0, 103, 104, 5, 78, 0, 0, 104, 108,
		5, 68, 0, 0, 105, 106, 5, 38, 0, 0, 106, 108, 5, 38, 0, 0, 107, 99, 1,
		0, 0, 0, 107, 102, 1, 0, 0, 0, 107, 105, 1, 0, 0, 0, 108, 18, 1, 0, 0,
		0, 109, 110, 5, 111, 0, 0, 110, 116, 5, 114, 0, 0, 111, 112, 5, 79, 0,
		0, 112, 116, 5, 82, 0, 0, 113, 114, 5, 124, 0, 0, 114, 116, 5, 124, 0,
		0, 115, 109, 1, 0, 0, 0, 115, 111, 1, 0, 0, 0, 115, 113, 1, 0, 0, 0, 116,
		20, 1, 0, 0, 0, 117, 118, 5, 116, 0, 0, 118, 119, 5, 114, 0, 0, 119, 120,
		5, 117, 0, 0, 120, 127, 5, 101, 0, 0, 121, 122, 5, 102, 0, 0, 122, 123,
		5, 97, 0, 0, 123, 124, 5, 108, 0, 0, 124, 125, 5, 115, 0, 0, 125, 127,
		5, 101, 0, 0, 126, 117, 1, 0, 0, 0, 126, 121, 1, 0, 0, 0, 127, 22, 1, 0,
		0, 0, 128, 129, 5, 110, 0, 0, 129, 130, 5, 117, 0, 0, 130, 131, 5, 108,
		0, 0, 131, 132, 5, 108, 0, 0, 132, 24, 1, 0, 0, 0, 133, 134, 5, 73, 0,
		0, 134, 138, 5, 78, 0, 0, 135, 136, 5, 105, 0, 0, 136, 138, 5, 110, 0,
		0, 137, 133, 1, 0, 0, 0, 137, 135, 1, 0, 0, 0, 138, 26, 1, 0, 0, 0, 139,
		140, 5, 101, 0, 0, 140, 146, 5, 113, 0, 0, 141, 142, 5, 69, 0, 0, 142,
		146, 5, 81, 0, 0, 143, 144, 5, 61, 0, 0, 144, 146, 5, 61, 0, 0, 145, 139,
		1, 0, 0, 0, 145, 141, 1, 0, 0, 0, 145, 143, 1, 0, 0, 0, 146, 28, 1, 0,
		0, 0, 147, 148, 5, 110, 0, 0, 148, 154, 5, 101, 0, 0, 149, 150, 5, 78,
		0, 0, 150, 154, 5, 69, 0, 0, 151, 152, 5, 33, 0, 0, 152, 154, 5, 61, 0,
		0, 153, 147, 1, 0, 0, 0, 153, 149, 1, 0, 0, 0, 153, 151, 1, 0, 0, 0, 154,
		30, 1, 0, 0, 0, 155, 156, 5, 103, 0, 0, 156, 161, 5, 116, 0, 0, 157, 158,
		5, 71, 0, 0, 158, 161, 5, 84, 0, 0, 159, 161, 5, 62, 0, 0, 160, 155, 1,
		0, 0, 0, 160, 157, 1, 0, 0, 0, 160, 159, 1, 0, 0, 0, 161, 32, 1, 0, 0,
		0, 162, 163, 5, 108, 0, 0, 163, 168, 5, 116, 0, 0, 164, 165, 5, 76, 0,
		0, 165, 168, 5, 84, 0, 0, 166, 168, 5, 60, 0, 0, 167, 162, 1, 0, 0, 0,
		167, 164, 1, 0, 0, 0, 167, 166, 1, 0, 0, 0, 168, 34, 1, 0, 0, 0, 169, 170,
		5, 103, 0, 0, 170, 176, 5, 101, 0, 0, 171, 172, 5, 71, 0, 0, 172, 176,
		5, 69, 0, 0, 173, 174, 5, 62, 0, 0, 174, 176, 5, 61, 0, 0, 175, 169, 1,
		0, 0, 0, 175, 171, 1, 0, 0, 0, 175, 173, 1, 0, 0, 0, 176, 36, 1, 0, 0,
		0, 177, 178, 5, 108, 0, 0, 178, 184, 5, 101, 0, 0, 179, 180, 5, 76, 0,
		0, 180, 184, 5, 69, 0, 0, 181, 182, 5, 60, 0, 0, 182, 184, 5, 61, 0, 0,
		183, 177, 1, 0, 0, 0, 183, 179, 1, 0, 0, 0, 183, 181, 1, 0, 0, 0, 184,
		38, 1, 0, 0, 0, 185, 186, 5, 99, 0, 0, 186, 190, 5, 111, 0, 0, 187, 188,
		5, 67, 0, 0, 188, 190, 5, 79, 0, 0, 189, 185, 1, 0, 0, 0, 189, 187, 1,
		0, 0, 0, 190, 40, 1, 0, 0, 0, 191, 192, 5, 115, 0, 0, 192, 196, 5, 119,
		0, 0, 193, 194, 5, 83, 0, 0, 194, 196, 5, 87, 0, 0, 195, 191, 1, 0, 0,
		0, 195, 193, 1, 0, 0, 0, 196, 42, 1, 0, 0, 0, 197, 198, 5, 101, 0, 0, 198,
		202, 5, 119, 0, 0, 199, 200, 5, 69, 0, 0, 200, 202, 5, 87, 0, 0, 201, 197,
		1, 0, 0, 0, 201, 199, 1, 0, 0, 0, 202, 44, 1, 0, 0, 0, 203, 207, 3, 51,
		25, 0, 204, 206, 3, 47, 23, 0, 205, 204, 1, 0, 0, 0, 206, 209, 1, 0, 0,
		0, 207, 205, 1, 0, 0, 0, 207, 208, 1, 0, 0, 0, 208, 46, 1, 0, 0, 0, 209,
		207, 1, 0, 0, 0, 210, 214, 7, 0, 0, 0, 211, 214, 3, 49, 24, 0, 212, 214,
		3, 51, 25, 0, 213, 210, 1, 0, 0, 0, 213, 211, 1, 0, 0, 0, 213, 212, 1,
		0, 0, 0, 214, 48, 1, 0, 0, 0, 215, 216, 2, 48, 57, 0, 216, 50, 1, 0, 0,
		0, 217, 218, 7, 1, 0, 0, 218, 52, 1, 0, 0, 0, 219, 220, 3, 65, 32, 0, 220,
		221, 5, 46, 0, 0, 221, 222, 3, 65, 32, 0, 222, 223, 5, 46, 0, 0, 223, 224,
		3, 65, 32, 0, 224, 54, 1, 0, 0, 0, 225, 230, 5, 34, 0, 0, 226, 229, 3,
		57, 28, 0, 227, 229, 8, 2, 0, 0, 228, 226, 1, 0, 0, 0, 228, 227, 1, 0,
		0, 0, 229, 232, 1, 0, 0, 0, 230, 228, 1, 0, 0, 0, 230, 231, 1, 0, 0, 0,
		231, 233, 1, 0, 0, 0, 232, 230, 1, 0, 0, 0, 233, 244, 5, 34, 0, 0, 234,
		239, 5, 39, 0, 0, 235, 238, 3, 57, 28, 0, 236, 238, 8, 3, 0, 0, 237, 235,
		1, 0, 0, 0, 237, 236, 1, 0, 0, 0, 238, 241, 1, 0, 0, 0, 239, 237, 1, 0,
		0, 0, 239, 240, 1, 0, 0, 0, 240, 242, 1, 0, 0, 0, 241, 239, 1, 0, 0, 0,
		242, 244, 5, 39, 0, 0, 243, 225, 1, 0, 0, 0, 243, 234, 1, 0, 0, 0, 244,
		56, 1, 0, 0, 0, 245, 248, 5, 92, 0, 0, 246, 249, 7, 4, 0, 0, 247, 249,
		3, 59, 29, 0, 248, 246, 1, 0, 0, 0, 248, 247, 1, 0, 0, 0, 249, 58, 1, 0,
		0, 0, 250, 251, 5, 117, 0, 0, 251, 252, 3, 61, 30, 0, 252, 253, 3, 61,
		30, 0, 253, 254, 3, 61, 30, 0, 254, 255, 3, 61, 30, 0, 255, 60, 1, 0, 0,
		0, 256, 257, 7, 5, 0, 0, 257, 62, 1, 0, 0, 0, 258, 260, 5, 45, 0, 0, 259,
		258, 1, 0, 0, 0, 259, 260, 1, 0, 0, 0, 260, 261, 1, 0, 0, 0, 261, 262,
		3, 65, 32, 0, 262, 264, 5, 46, 0, 0, 263, 265, 7, 6, 0, 0, 264, 263, 1,
		0, 0, 0, 265, 266, 1, 0, 0, 0, 266, 264, 1, 0, 0, 0, 266, 267, 1, 0, 0,
		0, 267, 269, 1, 0, 0, 0, 268, 270, 3, 67, 33, 0, 269, 268, 1, 0, 0, 0,
		269, 270, 1, 0, 0, 0, 270, 64, 1, 0, 0, 0, 271, 280, 5, 48, 0, 0, 272,
		276, 7, 7, 0, 0, 273, 275, 7, 6, 0, 0, 274, 273, 1, 0, 0, 0, 275, 278,
		1, 0, 0, 0, 276, 274, 1, 0, 0, 0, 276, 277, 1, 0, 0, 0, 277, 280, 1, 0,
		0, 0, 278, 276, 1, 0, 0, 0, 279, 271, 1, 0, 0, 0, 279, 272, 1, 0, 0, 0,
		280, 66, 1, 0, 0, 0, 281, 283, 7, 8, 0, 0, 282, 284, 7, 9, 0, 0, 283, 282,
		1, 0, 0, 0, 283, 284, 1, 0, 0, 0, 284, 285, 1, 0, 0, 0, 285, 286, 3, 65,
		32, 0, 286, 68, 1, 0, 0, 0, 287, 288, 5, 10, 0, 0, 288, 70, 1, 0, 0, 0,
		289, 293, 5, 44, 0, 0, 290, 292, 5, 32, 0, 0, 291, 290, 1, 0, 0, 0, 292,
		295, 1, 0, 0, 0, 293, 291, 1, 0, 0, 0, 293, 294, 1, 0, 0, 0, 294, 72, 1,
		0, 0, 0, 295, 293, 1, 0, 0, 0, 296, 300, 5, 32, 0, 0, 297, 299, 3, 69,
		34, 0, 298, 297, 1, 0, 0, 0, 299, 302, 1, 0, 0, 0, 300, 298, 1, 0, 0, 0,
		300, 301, 1, 0, 0, 0, 301, 74, 1, 0, 0, 0, 302, 300, 1, 0, 0, 0, 31, 0,
		97, 107, 115, 126, 137, 145, 153, 160, 167, 175, 183, 189, 195, 201, 207,
		213, 228, 230, 237, 239, 243, 248, 259, 266, 269, 276, 279, 283, 293, 300,
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
	JsonQueryLexerT__0         = 1
	JsonQueryLexerT__1         = 2
	JsonQueryLexerT__2         = 3
	JsonQueryLexerT__3         = 4
	JsonQueryLexerT__4         = 5
	JsonQueryLexerT__5         = 6
	JsonQueryLexerT__6         = 7
	JsonQueryLexerNOT          = 8
	JsonQueryLexerAND_OPERATOR = 9
	JsonQueryLexerOR_OPERATOR  = 10
	JsonQueryLexerBOOLEAN      = 11
	JsonQueryLexerNULL         = 12
	JsonQueryLexerIN           = 13
	JsonQueryLexerEQ           = 14
	JsonQueryLexerNE           = 15
	JsonQueryLexerGT           = 16
	JsonQueryLexerLT           = 17
	JsonQueryLexerGE           = 18
	JsonQueryLexerLE           = 19
	JsonQueryLexerCO           = 20
	JsonQueryLexerSW           = 21
	JsonQueryLexerEW           = 22
	JsonQueryLexerATTRNAME     = 23
	JsonQueryLexerVERSION      = 24
	JsonQueryLexerSTRING       = 25
	JsonQueryLexerDOUBLE       = 26
	JsonQueryLexerINT          = 27
	JsonQueryLexerEXP          = 28
	JsonQueryLexerNEWLINE      = 29
	JsonQueryLexerCOMMA        = 30
	JsonQueryLexerSP           = 31
)
