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
		"query", "subquery", "attrPath", "subAttr", "value", "listStrings",
		"subListOfStrings", "listNumbers", "subListOfNumbers", "listBooleans",
		"subListOfBooleans", "functionArg",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 31, 132, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 3, 1, 30, 8, 1, 1, 1, 3,
		1, 33, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 54, 8, 1, 10,
		1, 12, 1, 57, 9, 1, 1, 1, 1, 1, 1, 1, 3, 1, 62, 8, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 74, 8, 1, 10, 1, 12, 1,
		77, 9, 1, 1, 2, 1, 2, 3, 2, 81, 8, 2, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 95, 8, 4, 1, 5, 1, 5, 1, 5,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 105, 8, 6, 1, 7, 1, 7, 1, 7, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 115, 8, 8, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 3, 10, 125, 8, 10, 1, 11, 1, 11, 1, 11, 3, 11, 130,
		8, 11, 1, 11, 0, 1, 2, 12, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 0,
		2, 1, 0, 12, 21, 1, 0, 25, 26, 141, 0, 24, 1, 0, 0, 0, 2, 61, 1, 0, 0,
		0, 4, 78, 1, 0, 0, 0, 6, 82, 1, 0, 0, 0, 8, 94, 1, 0, 0, 0, 10, 96, 1,
		0, 0, 0, 12, 104, 1, 0, 0, 0, 14, 106, 1, 0, 0, 0, 16, 114, 1, 0, 0, 0,
		18, 116, 1, 0, 0, 0, 20, 124, 1, 0, 0, 0, 22, 129, 1, 0, 0, 0, 24, 25,
		3, 2, 1, 0, 25, 26, 5, 0, 0, 1, 26, 1, 1, 0, 0, 0, 27, 29, 6, 1, -1, 0,
		28, 30, 5, 7, 0, 0, 29, 28, 1, 0, 0, 0, 29, 30, 1, 0, 0, 0, 30, 32, 1,
		0, 0, 0, 31, 33, 5, 31, 0, 0, 32, 31, 1, 0, 0, 0, 32, 33, 1, 0, 0, 0, 33,
		34, 1, 0, 0, 0, 34, 35, 5, 1, 0, 0, 35, 36, 3, 2, 1, 0, 36, 37, 5, 2, 0,
		0, 37, 62, 1, 0, 0, 0, 38, 39, 3, 4, 2, 0, 39, 40, 5, 31, 0, 0, 40, 41,
		5, 3, 0, 0, 41, 62, 1, 0, 0, 0, 42, 43, 3, 4, 2, 0, 43, 44, 5, 31, 0, 0,
		44, 45, 7, 0, 0, 0, 45, 46, 5, 31, 0, 0, 46, 47, 3, 8, 4, 0, 47, 62, 1,
		0, 0, 0, 48, 49, 3, 4, 2, 0, 49, 55, 5, 1, 0, 0, 50, 51, 3, 22, 11, 0,
		51, 52, 5, 30, 0, 0, 52, 54, 1, 0, 0, 0, 53, 50, 1, 0, 0, 0, 54, 57, 1,
		0, 0, 0, 55, 53, 1, 0, 0, 0, 55, 56, 1, 0, 0, 0, 56, 58, 1, 0, 0, 0, 57,
		55, 1, 0, 0, 0, 58, 59, 3, 22, 11, 0, 59, 60, 5, 2, 0, 0, 60, 62, 1, 0,
		0, 0, 61, 27, 1, 0, 0, 0, 61, 38, 1, 0, 0, 0, 61, 42, 1, 0, 0, 0, 61, 48,
		1, 0, 0, 0, 62, 75, 1, 0, 0, 0, 63, 64, 10, 5, 0, 0, 64, 65, 5, 31, 0,
		0, 65, 66, 5, 8, 0, 0, 66, 67, 5, 31, 0, 0, 67, 74, 3, 2, 1, 6, 68, 69,
		10, 4, 0, 0, 69, 70, 5, 31, 0, 0, 70, 71, 5, 9, 0, 0, 71, 72, 5, 31, 0,
		0, 72, 74, 3, 2, 1, 5, 73, 63, 1, 0, 0, 0, 73, 68, 1, 0, 0, 0, 74, 77,
		1, 0, 0, 0, 75, 73, 1, 0, 0, 0, 75, 76, 1, 0, 0, 0, 76, 3, 1, 0, 0, 0,
		77, 75, 1, 0, 0, 0, 78, 80, 5, 22, 0, 0, 79, 81, 3, 6, 3, 0, 80, 79, 1,
		0, 0, 0, 80, 81, 1, 0, 0, 0, 81, 5, 1, 0, 0, 0, 82, 83, 5, 4, 0, 0, 83,
		84, 3, 4, 2, 0, 84, 7, 1, 0, 0, 0, 85, 95, 5, 10, 0, 0, 86, 95, 5, 11,
		0, 0, 87, 95, 5, 23, 0, 0, 88, 95, 5, 24, 0, 0, 89, 95, 5, 25, 0, 0, 90,
		95, 5, 26, 0, 0, 91, 95, 3, 14, 7, 0, 92, 95, 3, 10, 5, 0, 93, 95, 3, 18,
		9, 0, 94, 85, 1, 0, 0, 0, 94, 86, 1, 0, 0, 0, 94, 87, 1, 0, 0, 0, 94, 88,
		1, 0, 0, 0, 94, 89, 1, 0, 0, 0, 94, 90, 1, 0, 0, 0, 94, 91, 1, 0, 0, 0,
		94, 92, 1, 0, 0, 0, 94, 93, 1, 0, 0, 0, 95, 9, 1, 0, 0, 0, 96, 97, 5, 5,
		0, 0, 97, 98, 3, 12, 6, 0, 98, 11, 1, 0, 0, 0, 99, 100, 5, 24, 0, 0, 100,
		101, 5, 30, 0, 0, 101, 105, 3, 12, 6, 0, 102, 103, 5, 24, 0, 0, 103, 105,
		5, 6, 0, 0, 104, 99, 1, 0, 0, 0, 104, 102, 1, 0, 0, 0, 105, 13, 1, 0, 0,
		0, 106, 107, 5, 5, 0, 0, 107, 108, 3, 16, 8, 0, 108, 15, 1, 0, 0, 0, 109,
		110, 7, 1, 0, 0, 110, 111, 5, 30, 0, 0, 111, 115, 3, 16, 8, 0, 112, 113,
		7, 1, 0, 0, 113, 115, 5, 6, 0, 0, 114, 109, 1, 0, 0, 0, 114, 112, 1, 0,
		0, 0, 115, 17, 1, 0, 0, 0, 116, 117, 5, 5, 0, 0, 117, 118, 3, 20, 10, 0,
		118, 19, 1, 0, 0, 0, 119, 120, 5, 10, 0, 0, 120, 121, 5, 30, 0, 0, 121,
		125, 3, 20, 10, 0, 122, 123, 5, 10, 0, 0, 123, 125, 5, 6, 0, 0, 124, 119,
		1, 0, 0, 0, 124, 122, 1, 0, 0, 0, 125, 21, 1, 0, 0, 0, 126, 130, 3, 2,
		1, 0, 127, 130, 3, 4, 2, 0, 128, 130, 3, 8, 4, 0, 129, 126, 1, 0, 0, 0,
		129, 127, 1, 0, 0, 0, 129, 128, 1, 0, 0, 0, 130, 23, 1, 0, 0, 0, 12, 29,
		32, 55, 61, 73, 75, 80, 94, 104, 114, 124, 129,
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
	JsonQueryParserNOT          = 7
	JsonQueryParserAND_OPERATOR = 8
	JsonQueryParserOR_OPERATOR  = 9
	JsonQueryParserBOOLEAN      = 10
	JsonQueryParserNULL         = 11
	JsonQueryParserIN           = 12
	JsonQueryParserEQ           = 13
	JsonQueryParserNE           = 14
	JsonQueryParserGT           = 15
	JsonQueryParserLT           = 16
	JsonQueryParserGE           = 17
	JsonQueryParserLE           = 18
	JsonQueryParserCO           = 19
	JsonQueryParserSW           = 20
	JsonQueryParserEW           = 21
	JsonQueryParserATTRNAME     = 22
	JsonQueryParserVERSION      = 23
	JsonQueryParserSTRING       = 24
	JsonQueryParserDOUBLE       = 25
	JsonQueryParserLONG         = 26
	JsonQueryParserINT          = 27
	JsonQueryParserEXP          = 28
	JsonQueryParserNEWLINE      = 29
	JsonQueryParserCOMMA        = 30
	JsonQueryParserSP           = 31
)

// JsonQueryParser rules.
const (
	JsonQueryParserRULE_query             = 0
	JsonQueryParserRULE_subquery          = 1
	JsonQueryParserRULE_attrPath          = 2
	JsonQueryParserRULE_subAttr           = 3
	JsonQueryParserRULE_value             = 4
	JsonQueryParserRULE_listStrings       = 5
	JsonQueryParserRULE_subListOfStrings  = 6
	JsonQueryParserRULE_listNumbers       = 7
	JsonQueryParserRULE_subListOfNumbers  = 8
	JsonQueryParserRULE_listBooleans      = 9
	JsonQueryParserRULE_subListOfBooleans = 10
	JsonQueryParserRULE_functionArg       = 11
)

// IQueryContext is an interface to support dynamic dispatch.
type IQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Subquery() ISubqueryContext
	EOF() antlr.TerminalNode

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

func (s *QueryContext) Subquery() ISubqueryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubqueryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISubqueryContext)
}

func (s *QueryContext) EOF() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserEOF, 0)
}

func (s *QueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QueryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitQuery(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonQueryParser) Query() (localctx IQueryContext) {
	localctx = NewQueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, JsonQueryParserRULE_query)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(24)
		p.subquery(0)
	}
	{
		p.SetState(25)
		p.Match(JsonQueryParserEOF)
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

// ISubqueryContext is an interface to support dynamic dispatch.
type ISubqueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSubqueryContext differentiates from other interfaces.
	IsSubqueryContext()
}

type SubqueryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubqueryContext() *SubqueryContext {
	var p = new(SubqueryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_subquery
	return p
}

func InitEmptySubqueryContext(p *SubqueryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_subquery
}

func (*SubqueryContext) IsSubqueryContext() {}

func NewSubqueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubqueryContext {
	var p = new(SubqueryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonQueryParserRULE_subquery

	return p
}

func (s *SubqueryContext) GetParser() antlr.Parser { return s.parser }

func (s *SubqueryContext) CopyAll(ctx *SubqueryContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *SubqueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubqueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type CompareExpContext struct {
	SubqueryContext
	op antlr.Token
}

func NewCompareExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CompareExpContext {
	var p = new(CompareExpContext)

	InitEmptySubqueryContext(&p.SubqueryContext)
	p.parser = parser
	p.CopyAll(ctx.(*SubqueryContext))

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
	SubqueryContext
}

func NewAndLogicalExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndLogicalExpContext {
	var p = new(AndLogicalExpContext)

	InitEmptySubqueryContext(&p.SubqueryContext)
	p.parser = parser
	p.CopyAll(ctx.(*SubqueryContext))

	return p
}

func (s *AndLogicalExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndLogicalExpContext) AllSubquery() []ISubqueryContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISubqueryContext); ok {
			len++
		}
	}

	tst := make([]ISubqueryContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISubqueryContext); ok {
			tst[i] = t.(ISubqueryContext)
			i++
		}
	}

	return tst
}

func (s *AndLogicalExpContext) Subquery(i int) ISubqueryContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubqueryContext); ok {
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

	return t.(ISubqueryContext)
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
	SubqueryContext
}

func NewParenExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenExpContext {
	var p = new(ParenExpContext)

	InitEmptySubqueryContext(&p.SubqueryContext)
	p.parser = parser
	p.CopyAll(ctx.(*SubqueryContext))

	return p
}

func (s *ParenExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenExpContext) Subquery() ISubqueryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubqueryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISubqueryContext)
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
	SubqueryContext
}

func NewPresentExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PresentExpContext {
	var p = new(PresentExpContext)

	InitEmptySubqueryContext(&p.SubqueryContext)
	p.parser = parser
	p.CopyAll(ctx.(*SubqueryContext))

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
	SubqueryContext
}

func NewCallExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CallExpContext {
	var p = new(CallExpContext)

	InitEmptySubqueryContext(&p.SubqueryContext)
	p.parser = parser
	p.CopyAll(ctx.(*SubqueryContext))

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
	SubqueryContext
}

func NewOrLogicalExpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrLogicalExpContext {
	var p = new(OrLogicalExpContext)

	InitEmptySubqueryContext(&p.SubqueryContext)
	p.parser = parser
	p.CopyAll(ctx.(*SubqueryContext))

	return p
}

func (s *OrLogicalExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrLogicalExpContext) AllSubquery() []ISubqueryContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISubqueryContext); ok {
			len++
		}
	}

	tst := make([]ISubqueryContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISubqueryContext); ok {
			tst[i] = t.(ISubqueryContext)
			i++
		}
	}

	return tst
}

func (s *OrLogicalExpContext) Subquery(i int) ISubqueryContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubqueryContext); ok {
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

	return t.(ISubqueryContext)
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

func (p *JsonQueryParser) Subquery() (localctx ISubqueryContext) {
	return p.subquery(0)
}

func (p *JsonQueryParser) subquery(_p int) (localctx ISubqueryContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewSubqueryContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ISubqueryContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, JsonQueryParserRULE_subquery, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		localctx = NewParenExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		p.SetState(29)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == JsonQueryParserNOT {
			{
				p.SetState(28)
				p.Match(JsonQueryParserNOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(32)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == JsonQueryParserSP {
			{
				p.SetState(31)
				p.Match(JsonQueryParserSP)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(34)
			p.Match(JsonQueryParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(35)
			p.subquery(0)
		}
		{
			p.SetState(36)
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
			p.SetState(38)
			p.AttrPath()
		}
		{
			p.SetState(39)
			p.Match(JsonQueryParserSP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(40)
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
			p.SetState(42)
			p.AttrPath()
		}
		{
			p.SetState(43)
			p.Match(JsonQueryParserSP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(44)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*CompareExpContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4190208) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*CompareExpContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(45)
			p.Match(JsonQueryParserSP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(46)
			p.Value()
		}

	case 4:
		localctx = NewCallExpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(48)
			p.AttrPath()
		}
		{
			p.SetState(49)
			p.Match(JsonQueryParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(55)
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
					p.SetState(50)
					p.FunctionArg()
				}
				{
					p.SetState(51)
					p.Match(JsonQueryParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}
			p.SetState(57)
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
			p.SetState(58)
			p.FunctionArg()
		}
		{
			p.SetState(59)
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
	p.SetState(75)
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
			p.SetState(73)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
			case 1:
				localctx = NewAndLogicalExpContext(p, NewSubqueryContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JsonQueryParserRULE_subquery)
				p.SetState(63)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(64)
					p.Match(JsonQueryParserSP)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(65)
					p.Match(JsonQueryParserAND_OPERATOR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(66)
					p.Match(JsonQueryParserSP)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(67)
					p.subquery(6)
				}

			case 2:
				localctx = NewOrLogicalExpContext(p, NewSubqueryContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, JsonQueryParserRULE_subquery)
				p.SetState(68)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(69)
					p.Match(JsonQueryParserSP)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(70)
					p.Match(JsonQueryParserOR_OPERATOR)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(71)
					p.Match(JsonQueryParserSP)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(72)
					p.subquery(5)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(77)
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
	p.EnterRule(localctx, 4, JsonQueryParserRULE_attrPath)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(78)
		p.Match(JsonQueryParserATTRNAME)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(80)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == JsonQueryParserT__3 {
		{
			p.SetState(79)
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
	p.EnterRule(localctx, 6, JsonQueryParserRULE_subAttr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(82)
		p.Match(JsonQueryParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(83)
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

type ListOfNumbersContext struct {
	ValueContext
}

func NewListOfNumbersContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListOfNumbersContext {
	var p = new(ListOfNumbersContext)

	InitEmptyValueContext(&p.ValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValueContext))

	return p
}

func (s *ListOfNumbersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListOfNumbersContext) ListNumbers() IListNumbersContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListNumbersContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListNumbersContext)
}

func (s *ListOfNumbersContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitListOfNumbers(s)

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

func (s *LongContext) LONG() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserLONG, 0)
}

func (s *LongContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitLong(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonQueryParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, JsonQueryParserRULE_value)
	p.SetState(94)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		localctx = NewBooleanContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(85)
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
			p.SetState(86)
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
			p.SetState(87)
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
			p.SetState(88)
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
			p.SetState(89)
			p.Match(JsonQueryParserDOUBLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		localctx = NewLongContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(90)
			p.Match(JsonQueryParserLONG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		localctx = NewListOfNumbersContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(91)
			p.ListNumbers()
		}

	case 8:
		localctx = NewListOfStringsContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(92)
			p.ListStrings()
		}

	case 9:
		localctx = NewListOfBooleansContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(93)
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
	p.EnterRule(localctx, 10, JsonQueryParserRULE_listStrings)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(96)
		p.Match(JsonQueryParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(97)
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
	p.EnterRule(localctx, 12, JsonQueryParserRULE_subListOfStrings)
	p.SetState(104)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(99)
			p.Match(JsonQueryParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(100)
			p.Match(JsonQueryParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(101)
			p.SubListOfStrings()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(102)
			p.Match(JsonQueryParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(103)
			p.Match(JsonQueryParserT__5)
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

// IListNumbersContext is an interface to support dynamic dispatch.
type IListNumbersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SubListOfNumbers() ISubListOfNumbersContext

	// IsListNumbersContext differentiates from other interfaces.
	IsListNumbersContext()
}

type ListNumbersContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListNumbersContext() *ListNumbersContext {
	var p = new(ListNumbersContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_listNumbers
	return p
}

func InitEmptyListNumbersContext(p *ListNumbersContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_listNumbers
}

func (*ListNumbersContext) IsListNumbersContext() {}

func NewListNumbersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListNumbersContext {
	var p = new(ListNumbersContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonQueryParserRULE_listNumbers

	return p
}

func (s *ListNumbersContext) GetParser() antlr.Parser { return s.parser }

func (s *ListNumbersContext) SubListOfNumbers() ISubListOfNumbersContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubListOfNumbersContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISubListOfNumbersContext)
}

func (s *ListNumbersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListNumbersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListNumbersContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitListNumbers(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonQueryParser) ListNumbers() (localctx IListNumbersContext) {
	localctx = NewListNumbersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, JsonQueryParserRULE_listNumbers)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(106)
		p.Match(JsonQueryParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(107)
		p.SubListOfNumbers()
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

// ISubListOfNumbersContext is an interface to support dynamic dispatch.
type ISubListOfNumbersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetNum returns the num token.
	GetNum() antlr.Token

	// SetNum sets the num token.
	SetNum(antlr.Token)

	// Getter signatures
	COMMA() antlr.TerminalNode
	SubListOfNumbers() ISubListOfNumbersContext
	LONG() antlr.TerminalNode
	DOUBLE() antlr.TerminalNode

	// IsSubListOfNumbersContext differentiates from other interfaces.
	IsSubListOfNumbersContext()
}

type SubListOfNumbersContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	num    antlr.Token
}

func NewEmptySubListOfNumbersContext() *SubListOfNumbersContext {
	var p = new(SubListOfNumbersContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_subListOfNumbers
	return p
}

func InitEmptySubListOfNumbersContext(p *SubListOfNumbersContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonQueryParserRULE_subListOfNumbers
}

func (*SubListOfNumbersContext) IsSubListOfNumbersContext() {}

func NewSubListOfNumbersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubListOfNumbersContext {
	var p = new(SubListOfNumbersContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonQueryParserRULE_subListOfNumbers

	return p
}

func (s *SubListOfNumbersContext) GetParser() antlr.Parser { return s.parser }

func (s *SubListOfNumbersContext) GetNum() antlr.Token { return s.num }

func (s *SubListOfNumbersContext) SetNum(v antlr.Token) { s.num = v }

func (s *SubListOfNumbersContext) COMMA() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserCOMMA, 0)
}

func (s *SubListOfNumbersContext) SubListOfNumbers() ISubListOfNumbersContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubListOfNumbersContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISubListOfNumbersContext)
}

func (s *SubListOfNumbersContext) LONG() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserLONG, 0)
}

func (s *SubListOfNumbersContext) DOUBLE() antlr.TerminalNode {
	return s.GetToken(JsonQueryParserDOUBLE, 0)
}

func (s *SubListOfNumbersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubListOfNumbersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubListOfNumbersContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case JsonQueryVisitor:
		return t.VisitSubListOfNumbers(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *JsonQueryParser) SubListOfNumbers() (localctx ISubListOfNumbersContext) {
	localctx = NewSubListOfNumbersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, JsonQueryParserRULE_subListOfNumbers)
	var _la int

	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(109)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*SubListOfNumbersContext).num = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == JsonQueryParserDOUBLE || _la == JsonQueryParserLONG) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*SubListOfNumbersContext).num = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(110)
			p.Match(JsonQueryParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(111)
			p.SubListOfNumbers()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(112)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*SubListOfNumbersContext).num = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == JsonQueryParserDOUBLE || _la == JsonQueryParserLONG) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*SubListOfNumbersContext).num = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(113)
			p.Match(JsonQueryParserT__5)
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
	p.EnterRule(localctx, 18, JsonQueryParserRULE_listBooleans)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(116)
		p.Match(JsonQueryParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(117)
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
	p.EnterRule(localctx, 20, JsonQueryParserRULE_subListOfBooleans)
	p.SetState(124)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(119)
			p.Match(JsonQueryParserBOOLEAN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(120)
			p.Match(JsonQueryParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(121)
			p.SubListOfBooleans()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(122)
			p.Match(JsonQueryParserBOOLEAN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(123)
			p.Match(JsonQueryParserT__5)
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
	Subquery() ISubqueryContext
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

func (s *FunctionArgContext) Subquery() ISubqueryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISubqueryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISubqueryContext)
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
	p.EnterRule(localctx, 22, JsonQueryParserRULE_functionArg)
	p.SetState(129)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(126)
			p.subquery(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(127)
			p.AttrPath()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(128)
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
	case 1:
		var t *SubqueryContext = nil
		if localctx != nil {
			t = localctx.(*SubqueryContext)
		}
		return p.Subquery_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *JsonQueryParser) Subquery_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
