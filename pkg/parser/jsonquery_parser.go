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

package parser // JsonQuery

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

type JsonQueryParser struct {
	*antlr.BaseParser
}

var JsonQueryParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func jsonqueryParserInit() {
	staticData := &JsonQueryParserStaticData
	staticData.LiteralNames = []string{
		"", "'('", "')'", "'pr'", "'.'", "'-'", "'['", "']'", "", "", "", "",
		"'null'", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "'\\n'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "NOT", "AND_OPERATOR", "OR_OPERATOR",
		"BOOLEAN", "NULL", "IN", "EQ", "NE", "GT", "LT", "GE", "LE", "CO", "SW",
		"EW", "ATTRNAME", "VERSION", "STRING", "DOUBLE", "INT", "EXP", "NEWLINE",
		"COMMA", "SP",
	}
	staticData.RuleNames = []string{
		"query", "attrPath", "subAttr", "value", "listStrings", "subListOfStrings",
		"listDoubles", "subListOfDoubles", "listInts", "subListOfInts", "listBooleans",
		"subListOfBooleans", "functionArg",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 31, 148, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 1, 0, 1, 0, 3, 0, 29, 8, 0, 1, 0, 3, 0,
		32, 8, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 5, 0, 53, 8, 0, 10,
		0, 12, 0, 56, 9, 0, 1, 0, 1, 0, 1, 0, 3, 0, 61, 8, 0, 1, 0, 1, 0, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 5, 0, 73, 8, 0, 10, 0, 12, 0,
		76, 9, 0, 1, 1, 1, 1, 3, 1, 80, 8, 1, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 3, 3, 91, 8, 3, 1, 3, 1, 3, 3, 3, 95, 8, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 3, 3, 101, 8, 3, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 3, 5, 111, 8, 5, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 3, 7, 121, 8, 7, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		3, 9, 131, 8, 9, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11,
		3, 11, 141, 8, 11, 1, 12, 1, 12, 1, 12, 3, 12, 146, 8, 12, 1, 12, 0, 1,
		0, 13, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 0, 1, 1, 0, 13, 22,
		160, 0, 60, 1, 0, 0, 0, 2, 77, 1, 0, 0, 0, 4, 81, 1, 0, 0, 0, 6, 100, 1,
		0, 0, 0, 8, 102, 1, 0, 0, 0, 10, 110, 1, 0, 0, 0, 12, 112, 1, 0, 0, 0,
		14, 120, 1, 0, 0, 0, 16, 122, 1, 0, 0, 0, 18, 130, 1, 0, 0, 0, 20, 132,
		1, 0, 0, 0, 22, 140, 1, 0, 0, 0, 24, 145, 1, 0, 0, 0, 26, 28, 6, 0, -1,
		0, 27, 29, 5, 8, 0, 0, 28, 27, 1, 0, 0, 0, 28, 29, 1, 0, 0, 0, 29, 31,
		1, 0, 0, 0, 30, 32, 5, 31, 0, 0, 31, 30, 1, 0, 0, 0, 31, 32, 1, 0, 0, 0,
		32, 33, 1, 0, 0, 0, 33, 34, 5, 1, 0, 0, 34, 35, 3, 0, 0, 0, 35, 36, 5,
		2, 0, 0, 36, 61, 1, 0, 0, 0, 37, 38, 3, 2, 1, 0, 38, 39, 5, 31, 0, 0, 39,
		40, 5, 3, 0, 0, 40, 61, 1, 0, 0, 0, 41, 42, 3, 2, 1, 0, 42, 43, 5, 31,
		0, 0, 43, 44, 7, 0, 0, 0, 44, 45, 5, 31, 0, 0, 45, 46, 3, 6, 3, 0, 46,
		61, 1, 0, 0, 0, 47, 48, 3, 2, 1, 0, 48, 54, 5, 1, 0, 0, 49, 50, 3, 24,
		12, 0, 50, 51, 5, 30, 0, 0, 51, 53, 1, 0, 0, 0, 52, 49, 1, 0, 0, 0, 53,
		56, 1, 0, 0, 0, 54, 52, 1, 0, 0, 0, 54, 55, 1, 0, 0, 0, 55, 57, 1, 0, 0,
		0, 56, 54, 1, 0, 0, 0, 57, 58, 3, 24, 12, 0, 58, 59, 5, 2, 0, 0, 59, 61,
		1, 0, 0, 0, 60, 26, 1, 0, 0, 0, 60, 37, 1, 0, 0, 0, 60, 41, 1, 0, 0, 0,
		60, 47, 1, 0, 0, 0, 61, 74, 1, 0, 0, 0, 62, 63, 10, 5, 0, 0, 63, 64, 5,
		31, 0, 0, 64, 65, 5, 9, 0, 0, 65, 66, 5, 31, 0, 0, 66, 73, 3, 0, 0, 6,
		67, 68, 10, 4, 0, 0, 68, 69, 5, 31, 0, 0, 69, 70, 5, 10, 0, 0, 70, 71,
		5, 31, 0, 0, 71, 73, 3, 0, 0, 5, 72, 62, 1, 0, 0, 0, 72, 67, 1, 0, 0, 0,
		73, 76, 1, 0, 0, 0, 74, 72, 1, 0, 0, 0, 74, 75, 1, 0, 0, 0, 75, 1, 1, 0,
		0, 0, 76, 74, 1, 0, 0, 0, 77, 79, 5, 23, 0, 0, 78, 80, 3, 4, 2, 0, 79,
		78, 1, 0, 0, 0, 79, 80, 1, 0, 0, 0, 80, 3, 1, 0, 0, 0, 81, 82, 5, 4, 0,
		0, 82, 83, 3, 2, 1, 0, 83, 5, 1, 0, 0, 0, 84, 101, 5, 11, 0, 0, 85, 101,
		5, 12, 0, 0, 86, 101, 5, 24, 0, 0, 87, 101, 5, 25, 0, 0, 88, 101, 5, 26,
		0, 0, 89, 91, 5, 5, 0, 0, 90, 89, 1, 0, 0, 0, 90, 91, 1, 0, 0, 0, 91, 92,
		1, 0, 0, 0, 92, 94, 5, 27, 0, 0, 93, 95, 5, 28, 0, 0, 94, 93, 1, 0, 0,
		0, 94, 95, 1, 0, 0, 0, 95, 101, 1, 0, 0, 0, 96, 101, 3, 16, 8, 0, 97, 101,
		3, 12, 6, 0, 98, 101, 3, 8, 4, 0, 99, 101, 3, 20, 10, 0, 100, 84, 1, 0,
		0, 0, 100, 85, 1, 0, 0, 0, 100, 86, 1, 0, 0, 0, 100, 87, 1, 0, 0, 0, 100,
		88, 1, 0, 0, 0, 100, 90, 1, 0, 0, 0, 100, 96, 1, 0, 0, 0, 100, 97, 1, 0,
		0, 0, 100, 98, 1, 0, 0, 0, 100, 99, 1, 0, 0, 0, 101, 7, 1, 0, 0, 0, 102,
		103, 5, 6, 0, 0, 103, 104, 3, 10, 5, 0, 104, 9, 1, 0, 0, 0, 105, 106, 5,
		25, 0, 0, 106, 107, 5, 30, 0, 0, 107, 111, 3, 10, 5, 0, 108, 109, 5, 25,
		0, 0, 109, 111, 5, 7, 0, 0, 110, 105, 1, 0, 0, 0, 110, 108, 1, 0, 0, 0,
		111, 11, 1, 0, 0, 0, 112, 113, 5, 6, 0, 0, 113, 114, 3, 14, 7, 0, 114,
		13, 1, 0, 0, 0, 115, 116, 5, 26, 0, 0, 116, 117, 5, 30, 0, 0, 117, 121,
		3, 14, 7, 0, 118, 119, 5, 26, 0, 0, 119, 121, 5, 7, 0, 0, 120, 115, 1,
		0, 0, 0, 120, 118, 1, 0, 0, 0, 121, 15, 1, 0, 0, 0, 122, 123, 5, 6, 0,
		0, 123, 124, 3, 18, 9, 0, 124, 17, 1, 0, 0, 0, 125, 126, 5, 27, 0, 0, 126,
		127, 5, 30, 0, 0, 127, 131, 3, 18, 9, 0, 128, 129, 5, 27, 0, 0, 129, 131,
		5, 7, 0, 0, 130, 125, 1, 0, 0, 0, 130, 128, 1, 0, 0, 0, 131, 19, 1, 0,
		0, 0, 132, 133, 5, 6, 0, 0, 133, 134, 3, 22, 11, 0, 134, 21, 1, 0, 0, 0,
		135, 136, 5, 11, 0, 0, 136, 137, 5, 30, 0, 0, 137, 141, 3, 22, 11, 0, 138,
		139, 5, 11, 0, 0, 139, 141, 5, 7, 0, 0, 140, 135, 1, 0, 0, 0, 140, 138,
		1, 0, 0, 0, 141, 23, 1, 0, 0, 0, 142, 146, 3, 0, 0, 0, 143, 146, 3, 2,
		1, 0, 144, 146, 3, 6, 3, 0, 145, 142, 1, 0, 0, 0, 145, 143, 1, 0, 0, 0,
		145, 144, 1, 0, 0, 0, 146, 25, 1, 0, 0, 0, 15, 28, 31, 54, 60, 72, 74,
		79, 90, 94, 100, 110, 120, 130, 140, 145,
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

// JsonQueryParserInit initializes any static state used to implement JsonQueryParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewJsonQueryParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func JsonQueryParserInit() {
	staticData := &JsonQueryParserStaticData
	staticData.once.Do(jsonqueryParserInit)
}

// NewJsonQueryParser produces a new parser instance for the optional input antlr.TokenStream.
func NewJsonQueryParser(input antlr.TokenStream) *JsonQueryParser {
	JsonQueryParserInit()
	this := new(JsonQueryParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &JsonQueryParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "JsonQuery.g4"

	return this
}

// JsonQueryParser tokens.
const (
	JsonQueryParserEOF          = antlr.TokenEOF
	JsonQueryParserT__0         = 1
	JsonQueryParserT__1         = 2
	JsonQueryParserT__2         = 3
	JsonQueryParserT__3         = 4
	JsonQueryParserT__4         = 5
	JsonQueryParserT__5         = 6
	JsonQueryParserT__6         = 7
	JsonQueryParserNOT          = 8
	JsonQueryParserAND_OPERATOR = 9
	JsonQueryParserOR_OPERATOR  = 10
	JsonQueryParserBOOLEAN      = 11
	JsonQueryParserNULL         = 12
	JsonQueryParserIN           = 13
	JsonQueryParserEQ           = 14
	JsonQueryParserNE           = 15
	JsonQueryParserGT           = 16
	JsonQueryParserLT           = 17
	JsonQueryParserGE           = 18
	JsonQueryParserLE           = 19
	JsonQueryParserCO           = 20
	JsonQueryParserSW           = 21
	JsonQueryParserEW           = 22
	JsonQueryParserATTRNAME     = 23
	JsonQueryParserVERSION      = 24
	JsonQueryParserSTRING       = 25
	JsonQueryParserDOUBLE       = 26
	JsonQueryParserINT          = 27
	JsonQueryParserEXP          = 28
	JsonQueryParserNEWLINE      = 29
	JsonQueryParserCOMMA        = 30
	JsonQueryParserSP           = 31
)

// JsonQueryParser rules.
const (
	JsonQueryParserRULE_query             = 0
	JsonQueryParserRULE_attrPath          = 1
	JsonQueryParserRULE_subAttr           = 2
	JsonQueryParserRULE_value             = 3
	JsonQueryParserRULE_listStrings       = 4
	JsonQueryParserRULE_subListOfStrings  = 5
	JsonQueryParserRULE_listDoubles       = 6
	JsonQueryParserRULE_subListOfDoubles  = 7
	JsonQueryParserRULE_listInts          = 8
	JsonQueryParserRULE_subListOfInts     = 9
	JsonQueryParserRULE_listBooleans      = 10
	JsonQueryParserRULE_subListOfBooleans = 11
	JsonQueryParserRULE_functionArg       = 12
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQueryContext() *QueryContext {
	var p = new(QueryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_query
	return p
}

func InitEmptyQueryContext(p *QueryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_query
}

func (*QueryContext) IsQueryContext() {}

func NewQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QueryContext {
	var p = new(QueryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonQueryParserRULE_query

	return p
}

func (s *QueryContext) GetParser() antlr.Parser { return s.parser }

func (s *QueryContext) CopyAll(ctx *QueryContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *QueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type CompareExpContext struct {
	QueryContext
	op antlr.Token
}

func NewCompareExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CompareExpContext {
	var p = new(CompareExpContext)

	InitEmptyQueryContext(&p.QueryContext)
	p.parser = parser
	p.CopyAll(ctx.(*QueryContext))

	return p
}

func (s *CompareExpContext) GetOp() antlr.Token { return s.op }

func (s *CompareExpContext) SetOp(v antlr.Token) { s.op = v }

func (s *CompareExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompareExpContext) AttrPath() IAttrPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttrPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttrPathContext)
}

func (s *CompareExpContext) AllSP() []antlr.TerminalNode {
	return s.GetTokens(JsonQueryParserSP)
}

func (s *CompareExpContext) SP(i int) antlr.TerminalNode {
	return s.GetToken(JsonQueryParserSP, i)
}

func (s *CompareExpContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *CompareExpContext) EQ() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserEQ, 0)
}

func (s *CompareExpContext) NE() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserNE, 0)
}

func (s *CompareExpContext) GT() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserGT, 0)
}

func (s *CompareExpContext) LT() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserLT, 0)
}

func (s *CompareExpContext) GE() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserGE, 0)
}

func (s *CompareExpContext) LE() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserLE, 0)
}

func (s *CompareExpContext) CO() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserCO, 0)
}

func (s *CompareExpContext) SW() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserSW, 0)
}

func (s *CompareExpContext) EW() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserEW, 0)
}

func (s *CompareExpContext) IN() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserIN, 0)
}

func (s *CompareExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitCompareExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type AndLogicalExpContext struct {
	QueryContext
}

func NewAndLogicalExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndLogicalExpContext {
	var p = new(AndLogicalExpContext)

	InitEmptyQueryContext(&p.QueryContext)
	p.parser = parser
	p.CopyAll(ctx.(*QueryContext))

	return p
}

func (s *AndLogicalExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndLogicalExpContext) AllQuery() []IQueryContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IQueryContext); ok {
			len++
		}
	}

	tst := make([]IQueryContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IQueryContext); ok {
			tst[i] = t.(IQueryContext)
			i++
		}
	}

	return tst
}

func (s *AndLogicalExpContext) Query(i int) IQueryContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQueryContext); ok {
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

	return t.(IQueryContext)
}

func (s *AndLogicalExpContext) AllSP() []antlr.TerminalNode {
	return s.GetTokens(JsonQueryParserSP)
}

func (s *AndLogicalExpContext) SP(i int) antlr.TerminalNode {
	return s.GetToken(JsonQueryParserSP, i)
}

func (s *AndLogicalExpContext) AND_OPERATOR() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserAND_OPERATOR, 0)
}

func (s *AndLogicalExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitAndLogicalExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParenExpContext struct {
	QueryContext
}

func NewParenExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenExpContext {
	var p = new(ParenExpContext)

	InitEmptyQueryContext(&p.QueryContext)
	p.parser = parser
	p.CopyAll(ctx.(*QueryContext))

	return p
}

func (s *ParenExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenExpContext) Query() IQueryContext {
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

func (s *ParenExpContext) NOT() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserNOT, 0)
}

func (s *ParenExpContext) SP() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserSP, 0)
}

func (s *ParenExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitParenExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type PresentExpContext struct {
	QueryContext
}

func NewPresentExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PresentExpContext {
	var p = new(PresentExpContext)

	InitEmptyQueryContext(&p.QueryContext)
	p.parser = parser
	p.CopyAll(ctx.(*QueryContext))

	return p
}

func (s *PresentExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PresentExpContext) AttrPath() IAttrPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttrPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttrPathContext)
}

func (s *PresentExpContext) SP() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserSP, 0)
}

func (s *PresentExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitPresentExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type CallExpContext struct {
	QueryContext
}

func NewCallExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CallExpContext {
	var p = new(CallExpContext)

	InitEmptyQueryContext(&p.QueryContext)
	p.parser = parser
	p.CopyAll(ctx.(*QueryContext))

	return p
}

func (s *CallExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallExpContext) AttrPath() IAttrPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttrPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttrPathContext)
}

func (s *CallExpContext) AllFunctionArg() []IFunctionArgContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFunctionArgContext); ok {
			len++
		}
	}

	tst := make([]IFunctionArgContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFunctionArgContext); ok {
			tst[i] = t.(IFunctionArgContext)
			i++
		}
	}

	return tst
}

func (s *CallExpContext) FunctionArg(i int) IFunctionArgContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionArgContext); ok {
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

	return t.(IFunctionArgContext)
}

func (s *CallExpContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(JsonQueryParserCOMMA)
}

func (s *CallExpContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(JsonQueryParserCOMMA, i)
}

func (s *CallExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitCallExp(s)

	default:
		return t.VisitChildren(s)
	}
}

type OrLogicalExpContext struct {
	QueryContext
}

func NewOrLogicalExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrLogicalExpContext {
	var p = new(OrLogicalExpContext)

	InitEmptyQueryContext(&p.QueryContext)
	p.parser = parser
	p.CopyAll(ctx.(*QueryContext))

	return p
}

func (s *OrLogicalExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrLogicalExpContext) AllQuery() []IQueryContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IQueryContext); ok {
			len++
		}
	}

	tst := make([]IQueryContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IQueryContext); ok {
			tst[i] = t.(IQueryContext)
			i++
		}
	}

	return tst
}

func (s *OrLogicalExpContext) Query(i int) IQueryContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQueryContext); ok {
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

	return t.(IQueryContext)
}

func (s *OrLogicalExpContext) AllSP() []antlr.TerminalNode {
	return s.GetTokens(JsonQueryParserSP)
}

func (s *OrLogicalExpContext) SP(i int) antlr.TerminalNode {
	return s.GetToken(JsonQueryParserSP, i)
}

func (s *OrLogicalExpContext) OR_OPERATOR() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserOR_OPERATOR, 0)
}

func (s *OrLogicalExpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitOrLogicalExp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonQueryParser) Query() (localctx IQueryContext) {
	return p.query(0)
}

func (p *JsonQueryParser) query(_p int) (localctx IQueryContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewQueryContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IQueryContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 0
	p.EnterRecursionRule(localctx, 0, JsonQueryParserRULE_query, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(60)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		localctx = NewParenExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		p.SetState(28)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == JsonQueryParserNOT {
			{
				p.SetState(27)
				p.Match(JsonQueryParserNOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(31)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == JsonQueryParserSP {
			{
				p.SetState(30)
				p.Match(JsonQueryParserSP)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(33)
			p.Match(JsonQueryParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(34)
			p.query(0)
		}
		{
			p.SetState(35)
			p.Match(JsonQueryParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewPresentExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(37)
			p.AttrPath()
		}
		{
			p.SetState(38)
			p.Match(JsonQueryParserSP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(39)
			p.Match(JsonQueryParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewCompareExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(41)
			p.AttrPath()
		}
		{
			p.SetState(42)
			p.Match(JsonQueryParserSP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(43)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*CompareExpContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8380416) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*CompareExpContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(44)
			p.Match(JsonQueryParserSP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(45)
			p.Value()
		}

	case 4:
		localctx = NewCallExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(47)
			p.AttrPath()
		}
		{
			p.SetState(48)
			p.Match(JsonQueryParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(54)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(49)
					p.FunctionArg()
				}
				{
					p.SetState(50)
					p.Match(JsonQueryParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(56)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		{
			p.SetState(57)
			p.FunctionArg()
		}
		{
			p.SetState(58)
			p.Match(JsonQueryParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(72)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
			case 1:
				localctx = NewAndLogicalExpContext(p, NewQueryContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JsonQueryParserRULE_query)
				p.SetState(62)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(63)
					p.Match(JsonQueryParserSP)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(64)
					p.Match(JsonQueryParserAND_OPERATOR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(65)
					p.Match(JsonQueryParserSP)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(66)
					p.query(6)
				}

			case 2:
				localctx = NewOrLogicalExpContext(p, NewQueryContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JsonQueryParserRULE_query)
				p.SetState(67)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(68)
					p.Match(JsonQueryParserSP)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(69)
					p.Match(JsonQueryParserOR_OPERATOR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(70)
					p.Match(JsonQueryParserSP)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(71)
					p.query(5)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(76)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
		if p.HasError() {
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
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAttrPathContext is an interface to support dynamic dispatch.
type IAttrPathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ATTRNAME() antlr.TerminalNode
	SubAttr() ISubAttrContext

	// IsAttrPathContext differentiates from other interfaces.
	IsAttrPathContext()
}

type AttrPathContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttrPathContext() *AttrPathContext {
	var p = new(AttrPathContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_attrPath
	return p
}

func InitEmptyAttrPathContext(p *AttrPathContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_attrPath
}

func (*AttrPathContext) IsAttrPathContext() {}

func NewAttrPathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttrPathContext {
	var p = new(AttrPathContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonQueryParserRULE_attrPath

	return p
}

func (s *AttrPathContext) GetParser() antlr.Parser { return s.parser }

func (s *AttrPathContext) ATTRNAME() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserATTRNAME, 0)
}

func (s *AttrPathContext) SubAttr() ISubAttrContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubAttrContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISubAttrContext)
}

func (s *AttrPathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttrPathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttrPathContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitAttrPath(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonQueryParser) AttrPath() (localctx IAttrPathContext) {
	localctx = NewAttrPathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, JsonQueryParserRULE_attrPath)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(77)
		p.Match(JsonQueryParserATTRNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == JsonQueryParserT__3 {
		{
			p.SetState(78)
			p.SubAttr()
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

// ISubAttrContext is an interface to support dynamic dispatch.
type ISubAttrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AttrPath() IAttrPathContext

	// IsSubAttrContext differentiates from other interfaces.
	IsSubAttrContext()
}

type SubAttrContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubAttrContext() *SubAttrContext {
	var p = new(SubAttrContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_subAttr
	return p
}

func InitEmptySubAttrContext(p *SubAttrContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_subAttr
}

func (*SubAttrContext) IsSubAttrContext() {}

func NewSubAttrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubAttrContext {
	var p = new(SubAttrContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonQueryParserRULE_subAttr

	return p
}

func (s *SubAttrContext) GetParser() antlr.Parser { return s.parser }

func (s *SubAttrContext) AttrPath() IAttrPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttrPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttrPathContext)
}

func (s *SubAttrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubAttrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubAttrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitSubAttr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonQueryParser) SubAttr() (localctx ISubAttrContext) {
	localctx = NewSubAttrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, JsonQueryParserRULE_subAttr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(81)
		p.Match(JsonQueryParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(82)
		p.AttrPath()
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

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_value
	return p
}

func InitEmptyValueContext(p *ValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_value
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonQueryParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) CopyAll(ctx *ValueContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ListOfDoublesContext struct {
	ValueContext
}

func NewListOfDoublesContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListOfDoublesContext {
	var p = new(ListOfDoublesContext)

	InitEmptyValueContext(&p.ValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueContext))

	return p
}

func (s *ListOfDoublesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListOfDoublesContext) ListDoubles() IListDoublesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListDoublesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListDoublesContext)
}

func (s *ListOfDoublesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitListOfDoubles(s)

	default:
		return t.VisitChildren(s)
	}
}

type ListOfStringsContext struct {
	ValueContext
}

func NewListOfStringsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListOfStringsContext {
	var p = new(ListOfStringsContext)

	InitEmptyValueContext(&p.ValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueContext))

	return p
}

func (s *ListOfStringsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListOfStringsContext) ListStrings() IListStringsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListStringsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListStringsContext)
}

func (s *ListOfStringsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitListOfStrings(s)

	default:
		return t.VisitChildren(s)
	}
}

type BooleanContext struct {
	ValueContext
}

func NewBooleanContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BooleanContext {
	var p = new(BooleanContext)

	InitEmptyValueContext(&p.ValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueContext))

	return p
}

func (s *BooleanContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserBOOLEAN, 0)
}

func (s *BooleanContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitBoolean(s)

	default:
		return t.VisitChildren(s)
	}
}

type NullContext struct {
	ValueContext
}

func NewNullContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NullContext {
	var p = new(NullContext)

	InitEmptyValueContext(&p.ValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueContext))

	return p
}

func (s *NullContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NullContext) NULL() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserNULL, 0)
}

func (s *NullContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitNull(s)

	default:
		return t.VisitChildren(s)
	}
}

type StringContext struct {
	ValueContext
}

func NewStringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringContext {
	var p = new(StringContext)

	InitEmptyValueContext(&p.ValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueContext))

	return p
}

func (s *StringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringContext) STRING() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserSTRING, 0)
}

func (s *StringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitString(s)

	default:
		return t.VisitChildren(s)
	}
}

type DoubleContext struct {
	ValueContext
}

func NewDoubleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DoubleContext {
	var p = new(DoubleContext)

	InitEmptyValueContext(&p.ValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueContext))

	return p
}

func (s *DoubleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DoubleContext) DOUBLE() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserDOUBLE, 0)
}

func (s *DoubleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitDouble(s)

	default:
		return t.VisitChildren(s)
	}
}

type ListOfBooleansContext struct {
	ValueContext
}

func NewListOfBooleansContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListOfBooleansContext {
	var p = new(ListOfBooleansContext)

	InitEmptyValueContext(&p.ValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueContext))

	return p
}

func (s *ListOfBooleansContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListOfBooleansContext) ListBooleans() IListBooleansContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListBooleansContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListBooleansContext)
}

func (s *ListOfBooleansContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitListOfBooleans(s)

	default:
		return t.VisitChildren(s)
	}
}

type VersionContext struct {
	ValueContext
}

func NewVersionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VersionContext {
	var p = new(VersionContext)

	InitEmptyValueContext(&p.ValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueContext))

	return p
}

func (s *VersionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VersionContext) VERSION() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserVERSION, 0)
}

func (s *VersionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitVersion(s)

	default:
		return t.VisitChildren(s)
	}
}

type LongContext struct {
	ValueContext
}

func NewLongContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LongContext {
	var p = new(LongContext)

	InitEmptyValueContext(&p.ValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueContext))

	return p
}

func (s *LongContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LongContext) INT() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserINT, 0)
}

func (s *LongContext) EXP() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserEXP, 0)
}

func (s *LongContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitLong(s)

	default:
		return t.VisitChildren(s)
	}
}

type ListOfIntsContext struct {
	ValueContext
}

func NewListOfIntsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListOfIntsContext {
	var p = new(ListOfIntsContext)

	InitEmptyValueContext(&p.ValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueContext))

	return p
}

func (s *ListOfIntsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListOfIntsContext) ListInts() IListIntsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListIntsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListIntsContext)
}

func (s *ListOfIntsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitListOfInts(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonQueryParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, JsonQueryParserRULE_value)
	var _la int

	p.SetState(100)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		localctx = NewBooleanContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(84)
			p.Match(JsonQueryParserBOOLEAN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewNullContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(85)
			p.Match(JsonQueryParserNULL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewVersionContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(86)
			p.Match(JsonQueryParserVERSION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewStringContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(87)
			p.Match(JsonQueryParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewDoubleContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(88)
			p.Match(JsonQueryParserDOUBLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		localctx = NewLongContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		p.SetState(90)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == JsonQueryParserT__4 {
			{
				p.SetState(89)
				p.Match(JsonQueryParserT__4)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(92)
			p.Match(JsonQueryParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(94)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(93)
				p.Match(JsonQueryParserEXP)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case 7:
		localctx = NewListOfIntsContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(96)
			p.ListInts()
		}

	case 8:
		localctx = NewListOfDoublesContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(97)
			p.ListDoubles()
		}

	case 9:
		localctx = NewListOfStringsContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(98)
			p.ListStrings()
		}

	case 10:
		localctx = NewListOfBooleansContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(99)
			p.ListBooleans()
		}

	case antlr.ATNInvalidAltNumber:
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

// IListStringsContext is an interface to support dynamic dispatch.
type IListStringsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SubListOfStrings() ISubListOfStringsContext

	// IsListStringsContext differentiates from other interfaces.
	IsListStringsContext()
}

type ListStringsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListStringsContext() *ListStringsContext {
	var p = new(ListStringsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_listStrings
	return p
}

func InitEmptyListStringsContext(p *ListStringsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_listStrings
}

func (*ListStringsContext) IsListStringsContext() {}

func NewListStringsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListStringsContext {
	var p = new(ListStringsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonQueryParserRULE_listStrings

	return p
}

func (s *ListStringsContext) GetParser() antlr.Parser { return s.parser }

func (s *ListStringsContext) SubListOfStrings() ISubListOfStringsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubListOfStringsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISubListOfStringsContext)
}

func (s *ListStringsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListStringsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListStringsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitListStrings(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonQueryParser) ListStrings() (localctx IListStringsContext) {
	localctx = NewListStringsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, JsonQueryParserRULE_listStrings)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(102)
		p.Match(JsonQueryParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(103)
		p.SubListOfStrings()
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

// ISubListOfStringsContext is an interface to support dynamic dispatch.
type ISubListOfStringsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode
	COMMA() antlr.TerminalNode
	SubListOfStrings() ISubListOfStringsContext

	// IsSubListOfStringsContext differentiates from other interfaces.
	IsSubListOfStringsContext()
}

type SubListOfStringsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubListOfStringsContext() *SubListOfStringsContext {
	var p = new(SubListOfStringsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_subListOfStrings
	return p
}

func InitEmptySubListOfStringsContext(p *SubListOfStringsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_subListOfStrings
}

func (*SubListOfStringsContext) IsSubListOfStringsContext() {}

func NewSubListOfStringsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubListOfStringsContext {
	var p = new(SubListOfStringsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonQueryParserRULE_subListOfStrings

	return p
}

func (s *SubListOfStringsContext) GetParser() antlr.Parser { return s.parser }

func (s *SubListOfStringsContext) STRING() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserSTRING, 0)
}

func (s *SubListOfStringsContext) COMMA() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserCOMMA, 0)
}

func (s *SubListOfStringsContext) SubListOfStrings() ISubListOfStringsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubListOfStringsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISubListOfStringsContext)
}

func (s *SubListOfStringsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubListOfStringsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubListOfStringsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitSubListOfStrings(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonQueryParser) SubListOfStrings() (localctx ISubListOfStringsContext) {
	localctx = NewSubListOfStringsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, JsonQueryParserRULE_subListOfStrings)
	p.SetState(110)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(105)
			p.Match(JsonQueryParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(106)
			p.Match(JsonQueryParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(107)
			p.SubListOfStrings()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(108)
			p.Match(JsonQueryParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(109)
			p.Match(JsonQueryParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
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

// IListDoublesContext is an interface to support dynamic dispatch.
type IListDoublesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SubListOfDoubles() ISubListOfDoublesContext

	// IsListDoublesContext differentiates from other interfaces.
	IsListDoublesContext()
}

type ListDoublesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListDoublesContext() *ListDoublesContext {
	var p = new(ListDoublesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_listDoubles
	return p
}

func InitEmptyListDoublesContext(p *ListDoublesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_listDoubles
}

func (*ListDoublesContext) IsListDoublesContext() {}

func NewListDoublesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListDoublesContext {
	var p = new(ListDoublesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonQueryParserRULE_listDoubles

	return p
}

func (s *ListDoublesContext) GetParser() antlr.Parser { return s.parser }

func (s *ListDoublesContext) SubListOfDoubles() ISubListOfDoublesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubListOfDoublesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISubListOfDoublesContext)
}

func (s *ListDoublesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListDoublesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListDoublesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitListDoubles(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonQueryParser) ListDoubles() (localctx IListDoublesContext) {
	localctx = NewListDoublesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, JsonQueryParserRULE_listDoubles)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(112)
		p.Match(JsonQueryParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(113)
		p.SubListOfDoubles()
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

// ISubListOfDoublesContext is an interface to support dynamic dispatch.
type ISubListOfDoublesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DOUBLE() antlr.TerminalNode
	COMMA() antlr.TerminalNode
	SubListOfDoubles() ISubListOfDoublesContext

	// IsSubListOfDoublesContext differentiates from other interfaces.
	IsSubListOfDoublesContext()
}

type SubListOfDoublesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubListOfDoublesContext() *SubListOfDoublesContext {
	var p = new(SubListOfDoublesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_subListOfDoubles
	return p
}

func InitEmptySubListOfDoublesContext(p *SubListOfDoublesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_subListOfDoubles
}

func (*SubListOfDoublesContext) IsSubListOfDoublesContext() {}

func NewSubListOfDoublesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubListOfDoublesContext {
	var p = new(SubListOfDoublesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonQueryParserRULE_subListOfDoubles

	return p
}

func (s *SubListOfDoublesContext) GetParser() antlr.Parser { return s.parser }

func (s *SubListOfDoublesContext) DOUBLE() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserDOUBLE, 0)
}

func (s *SubListOfDoublesContext) COMMA() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserCOMMA, 0)
}

func (s *SubListOfDoublesContext) SubListOfDoubles() ISubListOfDoublesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubListOfDoublesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISubListOfDoublesContext)
}

func (s *SubListOfDoublesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubListOfDoublesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubListOfDoublesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitSubListOfDoubles(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonQueryParser) SubListOfDoubles() (localctx ISubListOfDoublesContext) {
	localctx = NewSubListOfDoublesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, JsonQueryParserRULE_subListOfDoubles)
	p.SetState(120)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(115)
			p.Match(JsonQueryParserDOUBLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(116)
			p.Match(JsonQueryParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(117)
			p.SubListOfDoubles()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(118)
			p.Match(JsonQueryParserDOUBLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(119)
			p.Match(JsonQueryParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
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

// IListIntsContext is an interface to support dynamic dispatch.
type IListIntsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SubListOfInts() ISubListOfIntsContext

	// IsListIntsContext differentiates from other interfaces.
	IsListIntsContext()
}

type ListIntsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListIntsContext() *ListIntsContext {
	var p = new(ListIntsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_listInts
	return p
}

func InitEmptyListIntsContext(p *ListIntsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_listInts
}

func (*ListIntsContext) IsListIntsContext() {}

func NewListIntsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListIntsContext {
	var p = new(ListIntsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonQueryParserRULE_listInts

	return p
}

func (s *ListIntsContext) GetParser() antlr.Parser { return s.parser }

func (s *ListIntsContext) SubListOfInts() ISubListOfIntsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubListOfIntsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISubListOfIntsContext)
}

func (s *ListIntsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListIntsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListIntsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitListInts(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonQueryParser) ListInts() (localctx IListIntsContext) {
	localctx = NewListIntsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, JsonQueryParserRULE_listInts)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(122)
		p.Match(JsonQueryParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(123)
		p.SubListOfInts()
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

// ISubListOfIntsContext is an interface to support dynamic dispatch.
type ISubListOfIntsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT() antlr.TerminalNode
	COMMA() antlr.TerminalNode
	SubListOfInts() ISubListOfIntsContext

	// IsSubListOfIntsContext differentiates from other interfaces.
	IsSubListOfIntsContext()
}

type SubListOfIntsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubListOfIntsContext() *SubListOfIntsContext {
	var p = new(SubListOfIntsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_subListOfInts
	return p
}

func InitEmptySubListOfIntsContext(p *SubListOfIntsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_subListOfInts
}

func (*SubListOfIntsContext) IsSubListOfIntsContext() {}

func NewSubListOfIntsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubListOfIntsContext {
	var p = new(SubListOfIntsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonQueryParserRULE_subListOfInts

	return p
}

func (s *SubListOfIntsContext) GetParser() antlr.Parser { return s.parser }

func (s *SubListOfIntsContext) INT() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserINT, 0)
}

func (s *SubListOfIntsContext) COMMA() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserCOMMA, 0)
}

func (s *SubListOfIntsContext) SubListOfInts() ISubListOfIntsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubListOfIntsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISubListOfIntsContext)
}

func (s *SubListOfIntsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubListOfIntsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubListOfIntsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitSubListOfInts(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonQueryParser) SubListOfInts() (localctx ISubListOfIntsContext) {
	localctx = NewSubListOfIntsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, JsonQueryParserRULE_subListOfInts)
	p.SetState(130)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(125)
			p.Match(JsonQueryParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(126)
			p.Match(JsonQueryParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(127)
			p.SubListOfInts()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(128)
			p.Match(JsonQueryParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(129)
			p.Match(JsonQueryParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
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

// IListBooleansContext is an interface to support dynamic dispatch.
type IListBooleansContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SubListOfBooleans() ISubListOfBooleansContext

	// IsListBooleansContext differentiates from other interfaces.
	IsListBooleansContext()
}

type ListBooleansContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListBooleansContext() *ListBooleansContext {
	var p = new(ListBooleansContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_listBooleans
	return p
}

func InitEmptyListBooleansContext(p *ListBooleansContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_listBooleans
}

func (*ListBooleansContext) IsListBooleansContext() {}

func NewListBooleansContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListBooleansContext {
	var p = new(ListBooleansContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonQueryParserRULE_listBooleans

	return p
}

func (s *ListBooleansContext) GetParser() antlr.Parser { return s.parser }

func (s *ListBooleansContext) SubListOfBooleans() ISubListOfBooleansContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubListOfBooleansContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISubListOfBooleansContext)
}

func (s *ListBooleansContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListBooleansContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListBooleansContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitListBooleans(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonQueryParser) ListBooleans() (localctx IListBooleansContext) {
	localctx = NewListBooleansContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, JsonQueryParserRULE_listBooleans)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(132)
		p.Match(JsonQueryParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(133)
		p.SubListOfBooleans()
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

// ISubListOfBooleansContext is an interface to support dynamic dispatch.
type ISubListOfBooleansContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BOOLEAN() antlr.TerminalNode
	COMMA() antlr.TerminalNode
	SubListOfBooleans() ISubListOfBooleansContext

	// IsSubListOfBooleansContext differentiates from other interfaces.
	IsSubListOfBooleansContext()
}

type SubListOfBooleansContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubListOfBooleansContext() *SubListOfBooleansContext {
	var p = new(SubListOfBooleansContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_subListOfBooleans
	return p
}

func InitEmptySubListOfBooleansContext(p *SubListOfBooleansContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_subListOfBooleans
}

func (*SubListOfBooleansContext) IsSubListOfBooleansContext() {}

func NewSubListOfBooleansContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubListOfBooleansContext {
	var p = new(SubListOfBooleansContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonQueryParserRULE_subListOfBooleans

	return p
}

func (s *SubListOfBooleansContext) GetParser() antlr.Parser { return s.parser }

func (s *SubListOfBooleansContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserBOOLEAN, 0)
}

func (s *SubListOfBooleansContext) COMMA() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserCOMMA, 0)
}

func (s *SubListOfBooleansContext) SubListOfBooleans() ISubListOfBooleansContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubListOfBooleansContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISubListOfBooleansContext)
}

func (s *SubListOfBooleansContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubListOfBooleansContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubListOfBooleansContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitSubListOfBooleans(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonQueryParser) SubListOfBooleans() (localctx ISubListOfBooleansContext) {
	localctx = NewSubListOfBooleansContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, JsonQueryParserRULE_subListOfBooleans)
	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(135)
			p.Match(JsonQueryParserBOOLEAN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(136)
			p.Match(JsonQueryParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(137)
			p.SubListOfBooleans()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(138)
			p.Match(JsonQueryParserBOOLEAN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(139)
			p.Match(JsonQueryParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
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

// IFunctionArgContext is an interface to support dynamic dispatch.
type IFunctionArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Query() IQueryContext
	AttrPath() IAttrPathContext
	Value() IValueContext

	// IsFunctionArgContext differentiates from other interfaces.
	IsFunctionArgContext()
}

type FunctionArgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionArgContext() *FunctionArgContext {
	var p = new(FunctionArgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_functionArg
	return p
}

func InitEmptyFunctionArgContext(p *FunctionArgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_functionArg
}

func (*FunctionArgContext) IsFunctionArgContext() {}

func NewFunctionArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionArgContext {
	var p = new(FunctionArgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonQueryParserRULE_functionArg

	return p
}

func (s *FunctionArgContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionArgContext) Query() IQueryContext {
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

func (s *FunctionArgContext) AttrPath() IAttrPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttrPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttrPathContext)
}

func (s *FunctionArgContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *FunctionArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionArgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitFunctionArg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonQueryParser) FunctionArg() (localctx IFunctionArgContext) {
	localctx = NewFunctionArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, JsonQueryParserRULE_functionArg)
	p.SetState(145)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(142)
			p.query(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(143)
			p.AttrPath()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(144)
			p.Value()
		}

	case antlr.ATNInvalidAltNumber:
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

func (p *JsonQueryParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 0:
		var t *QueryContext = nil
		if localctx != nil {
			t = localctx.(*QueryContext)
		}
		return p.Query_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *JsonQueryParser) Query_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
