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

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by JsonQueryParser.
type JsonQueryVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by JsonQueryParser#query.
	VisitQuery(ctx *QueryContext) interface{}

	// Visit a parse tree produced by JsonQueryParser#compareExp.
	VisitCompareExp(ctx *CompareExpContext) interface{}

	// Visit a parse tree produced by JsonQueryParser#andLogicalExp.
	VisitAndLogicalExp(ctx *AndLogicalExpContext) interface{}

	// Visit a parse tree produced by JsonQueryParser#parenExp.
	VisitParenExp(ctx *ParenExpContext) interface{}

	// Visit a parse tree produced by JsonQueryParser#presentExp.
	VisitPresentExp(ctx *PresentExpContext) interface{}

	// Visit a parse tree produced by JsonQueryParser#callExp.
	VisitCallExp(ctx *CallExpContext) interface{}

	// Visit a parse tree produced by JsonQueryParser#orLogicalExp.
	VisitOrLogicalExp(ctx *OrLogicalExpContext) interface{}

	// Visit a parse tree produced by JsonQueryParser#attrPath.
	VisitAttrPath(ctx *AttrPathContext) interface{}

	// Visit a parse tree produced by JsonQueryParser#subAttr.
	VisitSubAttr(ctx *SubAttrContext) interface{}

	// Visit a parse tree produced by JsonQueryParser#boolean.
	VisitBoolean(ctx *BooleanContext) interface{}

	// Visit a parse tree produced by JsonQueryParser#null.
	VisitNull(ctx *NullContext) interface{}

	// Visit a parse tree produced by JsonQueryParser#version.
	VisitVersion(ctx *VersionContext) interface{}

	// Visit a parse tree produced by JsonQueryParser#string.
	VisitString(ctx *StringContext) interface{}

	// Visit a parse tree produced by JsonQueryParser#double.
	VisitDouble(ctx *DoubleContext) interface{}

	// Visit a parse tree produced by JsonQueryParser#long.
	VisitLong(ctx *LongContext) interface{}

	// Visit a parse tree produced by JsonQueryParser#listOfNumbers.
	VisitListOfNumbers(ctx *ListOfNumbersContext) interface{}

	// Visit a parse tree produced by JsonQueryParser#listOfStrings.
	VisitListOfStrings(ctx *ListOfStringsContext) interface{}

	// Visit a parse tree produced by JsonQueryParser#listOfBooleans.
	VisitListOfBooleans(ctx *ListOfBooleansContext) interface{}

	// Visit a parse tree produced by JsonQueryParser#listStrings.
	VisitListStrings(ctx *ListStringsContext) interface{}

	// Visit a parse tree produced by JsonQueryParser#subListOfStrings.
	VisitSubListOfStrings(ctx *SubListOfStringsContext) interface{}

	// Visit a parse tree produced by JsonQueryParser#listNumbers.
	VisitListNumbers(ctx *ListNumbersContext) interface{}

	// Visit a parse tree produced by JsonQueryParser#subListOfNumbers.
	VisitSubListOfNumbers(ctx *SubListOfNumbersContext) interface{}

	// Visit a parse tree produced by JsonQueryParser#listBooleans.
	VisitListBooleans(ctx *ListBooleansContext) interface{}

	// Visit a parse tree produced by JsonQueryParser#subListOfBooleans.
	VisitSubListOfBooleans(ctx *SubListOfBooleansContext) interface{}

	// Visit a parse tree produced by JsonQueryParser#functionArg.
	VisitFunctionArg(ctx *FunctionArgContext) interface{}
}
