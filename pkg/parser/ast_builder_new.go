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

package parser

import (
	"fmt"
	"strconv"

	rulesv1beta3 "buf.build/gen/go/lekkodev/cli/protocolbuffers/go/lekko/rules/v1beta3"

	"github.com/antlr4-go/antlr/v4"
	"google.golang.org/protobuf/types/known/structpb"
)

// ASTBuilderV3 correctly parses into the rulesv1beta3 model which supports n-ary ands and ors.
type ASTBuilderV3 struct {
	antlr.ParseTreeVisitor
}

func NewASTBuilderV3() *ASTBuilderV3 {
	return &ASTBuilderV3{}
}

func BuildASTV3(rule string) (*rulesv1beta3.Rule, error) {
	tree, err := lexAndParse(rule)
	if err != nil {
		return nil, err
	}
	switch v := NewASTBuilderV3().Visit(tree).(type) {
	case error:
		return nil, v
	case *rulesv1beta3.Rule:
		return v, nil
	default:
		return nil, fmt.Errorf("unknown type during AST building: %v %T", v, v)
	}

}

func (a *ASTBuilderV3) Visit(tree antlr.ParseTree) interface{} {
	switch val := tree.(type) {
	case *AndLogicalExpContext:
		return val.Accept(a)
	case *OrLogicalExpContext:
		return val.Accept(a)
	case *CompareExpContext:
		return val.Accept(a)
	case *ParenExpContext:
		return val.Accept(a)
	case *PresentExpContext:
		return val.Accept(a)
	case *CallExpContext:
		return val.Accept(a)
	default:
		return fmt.Errorf("invalid tree type: %v", val)
	}
}

func (a *ASTBuilderV3) VisitParenExp(ctx *ParenExpContext) interface{} {
	v := ctx.Query().Accept(a)
	if err, ok := v.(error); ok {
		return err
	}
	r, ok := v.(*rulesv1beta3.Rule)
	if !ok {
		return fmt.Errorf("invalid type during AST building: %v", v)
	}
	if ctx.NOT() != nil {
		return &rulesv1beta3.Rule{Rule: &rulesv1beta3.Rule_Not{Not: r}}
	}
	return v
}

func (a *ASTBuilderV3) VisitAndLogicalExp(ctx *AndLogicalExpContext) interface{} {
	left := ctx.Query(0).Accept(a)
	if err, ok := left.(error); ok {
		return err
	}
	leftR, ok := left.(*rulesv1beta3.Rule)
	if !ok {
		return fmt.Errorf("invalid type during AST building: %v", left)
	}

	right := ctx.Query(1).Accept(a)
	if err, ok := right.(error); ok {
		return err
	}
	rightR, ok := right.(*rulesv1beta3.Rule)
	if !ok {
		return fmt.Errorf("invalid type during AST building: %v", right)
	}

	// special case the left being an already constructed andlogicalexp for n-ary trees.
	if leftExpr := leftR.GetLogicalExpression(); leftExpr != nil {
		if leftExpr.LogicalOperator == rulesv1beta3.LogicalOperator_LOGICAL_OPERATOR_AND {
			leftExpr.Rules = append(leftExpr.Rules, rightR)
			return left
		}
	}

	return &rulesv1beta3.Rule{
		Rule: &rulesv1beta3.Rule_LogicalExpression{
			LogicalExpression: &rulesv1beta3.LogicalExpression{
				Rules:           []*rulesv1beta3.Rule{leftR, rightR},
				LogicalOperator: rulesv1beta3.LogicalOperator_LOGICAL_OPERATOR_AND,
			},
		},
	}
}

func (a *ASTBuilderV3) VisitOrLogicalExp(ctx *OrLogicalExpContext) interface{} {
	left := ctx.Query(0).Accept(a)
	if err, ok := left.(error); ok {
		return err
	}
	leftR, ok := left.(*rulesv1beta3.Rule)
	if !ok {
		return fmt.Errorf("invalid type during AST building: %v", left)
	}

	right := ctx.Query(1).Accept(a)
	if err, ok := right.(error); ok {
		return err
	}
	rightR, ok := right.(*rulesv1beta3.Rule)
	if !ok {
		return fmt.Errorf("invalid type during AST building: %v", right)
	}

	// special case the left being an already constructed orlogicalexp for n-ary trees.
	if leftExpr := leftR.GetLogicalExpression(); leftExpr != nil {
		if leftExpr.LogicalOperator == rulesv1beta3.LogicalOperator_LOGICAL_OPERATOR_OR {
			leftExpr.Rules = append(leftExpr.Rules, rightR)
			return left
		}
	}

	return &rulesv1beta3.Rule{
		Rule: &rulesv1beta3.Rule_LogicalExpression{
			LogicalExpression: &rulesv1beta3.LogicalExpression{
				Rules:           []*rulesv1beta3.Rule{leftR, rightR},
				LogicalOperator: rulesv1beta3.LogicalOperator_LOGICAL_OPERATOR_OR,
			},
		},
	}
}

func (a *ASTBuilderV3) VisitPresentExp(ctx *PresentExpContext) interface{} {
	return &rulesv1beta3.Rule{
		Rule: &rulesv1beta3.Rule_Atom{
			Atom: &rulesv1beta3.Atom{
				ContextKey:         ctx.AttrPath().Accept(a).(string),
				ComparisonOperator: rulesv1beta3.ComparisonOperator_COMPARISON_OPERATOR_PRESENT,
			},
		},
	}
}

func (a *ASTBuilderV3) VisitCompareExp(ctx *CompareExpContext) (ret interface{}) {
	key := ctx.AttrPath().Accept(a).(string)
	value := ctx.Value().Accept(a)

	if err, ok := value.(error); ok {
		return err
	}
	valueR, ok := value.(*structpb.Value)
	if !ok {
		return fmt.Errorf("invalid type during AST building: %v", value)
	}

	atom := &rulesv1beta3.Atom{
		ContextKey:      key,
		ComparisonValue: valueR,
	}
	// TODO: Do a type check on each righthanded operator.
	switch ctx.op.GetTokenType() {
	case JsonQueryParserEQ:
		atom.ComparisonOperator = rulesv1beta3.ComparisonOperator_COMPARISON_OPERATOR_EQUALS
	case JsonQueryParserNE:
		atom.ComparisonOperator = rulesv1beta3.ComparisonOperator_COMPARISON_OPERATOR_NOT_EQUALS
	case JsonQueryParserGT:
		atom.ComparisonOperator = rulesv1beta3.ComparisonOperator_COMPARISON_OPERATOR_GREATER_THAN
		if _, ok := valueR.GetKind().(*structpb.Value_NumberValue); !ok {
			return fmt.Errorf("invalid type for operator %v %T", atom.ComparisonOperator, valueR)
		}
	case JsonQueryParserLT:
		atom.ComparisonOperator = rulesv1beta3.ComparisonOperator_COMPARISON_OPERATOR_LESS_THAN
		if _, ok := valueR.GetKind().(*structpb.Value_NumberValue); !ok {
			return fmt.Errorf("invalid type for operator %v %T", atom.ComparisonOperator, valueR)
		}
	case JsonQueryParserLE:
		atom.ComparisonOperator = rulesv1beta3.ComparisonOperator_COMPARISON_OPERATOR_LESS_THAN_OR_EQUALS
		if _, ok := valueR.GetKind().(*structpb.Value_NumberValue); !ok {
			return fmt.Errorf("invalid type for operator %v %T", atom.ComparisonOperator, valueR)
		}
	case JsonQueryParserGE:
		atom.ComparisonOperator = rulesv1beta3.ComparisonOperator_COMPARISON_OPERATOR_GREATER_THAN_OR_EQUALS
		if _, ok := valueR.GetKind().(*structpb.Value_NumberValue); !ok {
			return fmt.Errorf("invalid type for operator %v %T", atom.ComparisonOperator, valueR)
		}
	case JsonQueryParserCO:
		atom.ComparisonOperator = rulesv1beta3.ComparisonOperator_COMPARISON_OPERATOR_CONTAINS
		if _, ok := valueR.GetKind().(*structpb.Value_StringValue); !ok {
			return fmt.Errorf("invalid type for operator %v %T", atom.ComparisonOperator, valueR)
		}
	case JsonQueryParserSW:
		atom.ComparisonOperator = rulesv1beta3.ComparisonOperator_COMPARISON_OPERATOR_STARTS_WITH
		if _, ok := valueR.GetKind().(*structpb.Value_StringValue); !ok {
			return fmt.Errorf("invalid type for operator %v %T", atom.ComparisonOperator, valueR)
		}
	case JsonQueryParserEW:
		atom.ComparisonOperator = rulesv1beta3.ComparisonOperator_COMPARISON_OPERATOR_ENDS_WITH
		if _, ok := valueR.GetKind().(*structpb.Value_StringValue); !ok {
			return fmt.Errorf("invalid type for operator %v %T", atom.ComparisonOperator, valueR)
		}
	case JsonQueryParserIN:
		atom.ComparisonOperator = rulesv1beta3.ComparisonOperator_COMPARISON_OPERATOR_CONTAINED_WITHIN
		if _, ok := valueR.GetKind().(*structpb.Value_ListValue); !ok {
			return fmt.Errorf("invalid type for operator %v %T", atom.ComparisonOperator, valueR)
		}
	default:
		return fmt.Errorf("invalid token: %v", ctx.op.GetTokenType())
	}
	return &rulesv1beta3.Rule{
		Rule: &rulesv1beta3.Rule_Atom{
			Atom: atom,
		},
	}
}

func (a *ASTBuilderV3) VisitCallExp(ctx *CallExpContext) interface{} {
	funcName := ctx.AttrPath().Accept(a).(string)

	funcVisitor, ok := functionVisitors[funcName]
	if !ok {
		return fmt.Errorf("invalid function name for call expression during AST building: %v", funcName)
	}

	return funcVisitor(a, ctx)
}

func (a *ASTBuilderV3) VisitAttrPath(ctx *AttrPathContext) interface{} {
	if ctx.SubAttr() == nil || ctx.SubAttr().IsEmpty() {
		return ctx.ATTRNAME().GetText()
	}
	return ctx.ATTRNAME().GetText() + ctx.SubAttr().Accept(a).(string)
}

func (a *ASTBuilderV3) VisitSubAttr(ctx *SubAttrContext) interface{} {
	return ctx.AttrPath().Accept(a)
}

func (a *ASTBuilderV3) VisitBoolean(ctx *BooleanContext) interface{} {
	val, err := strconv.ParseBool(ctx.GetText())
	if err != nil {
		return err
	}
	return structpb.NewBoolValue(val)
}

func (a *ASTBuilderV3) VisitNull(ctx *NullContext) interface{} {
	// TODO: not sure what to do here or if we even want to support it.
	return nil
}

func (a *ASTBuilderV3) VisitString(ctx *StringContext) interface{} {
	return structpb.NewStringValue(getString(ctx.GetText()))
}

func (a *ASTBuilderV3) VisitDouble(ctx *DoubleContext) interface{} {
	val, err := strconv.ParseFloat(ctx.GetText(), 64)
	if err != nil {
		return err
	}
	return structpb.NewNumberValue(val)
}

func (a *ASTBuilderV3) VisitVersion(ctx *VersionContext) interface{} {
	return structpb.NewStringValue(ctx.VERSION().GetText())
}

func (a *ASTBuilderV3) VisitLong(ctx *LongContext) interface{} {
	val, err := strconv.ParseInt(ctx.GetText(), 10, 64)
	if err != nil {
		return err
	}
	return structpb.NewNumberValue(float64(val))
}

func (a *ASTBuilderV3) VisitListOfInts(ctx *ListOfIntsContext) interface{} {
	return ctx.ListInts().Accept(a)
}

func (a *ASTBuilderV3) VisitListInts(ctx *ListIntsContext) interface{} {
	return ctx.SubListOfInts().Accept(a)
}

func (a *ASTBuilderV3) VisitSubListOfInts(ctx *SubListOfIntsContext) interface{} {
	val, err := strconv.ParseInt(ctx.INT().GetText(), 10, 64)
	if err != nil {
		return err
	}
	if ctx.SubListOfInts() == nil || ctx.SubListOfInts().IsEmpty() {
		res, err := structpb.NewList([]interface{}{float64(val)})
		if err != nil {
			return err
		}
		return structpb.NewListValue(res)
	}

	rest := ctx.SubListOfInts().Accept(a)
	if err, ok := rest.(error); ok {
		return err
	}
	restL, ok := rest.(*structpb.Value)
	if !ok {
		return fmt.Errorf("unknown type when parsing list of strings: %v", rest)
	}
	restL.GetListValue().Values = append([]*structpb.Value{structpb.NewNumberValue(float64(val))}, restL.GetListValue().Values...)
	return restL
}

func (a *ASTBuilderV3) VisitListOfDoubles(ctx *ListOfDoublesContext) interface{} {
	return ctx.ListDoubles().Accept(a)
}

func (a *ASTBuilderV3) VisitListDoubles(ctx *ListDoublesContext) interface{} {
	return ctx.SubListOfDoubles().Accept(a)
}

func (a *ASTBuilderV3) VisitSubListOfDoubles(ctx *SubListOfDoublesContext) interface{} {
	val, err := strconv.ParseFloat(ctx.DOUBLE().GetText(), 64)
	if err != nil {
		return err
	}
	if ctx.SubListOfDoubles() == nil || ctx.SubListOfDoubles().IsEmpty() {
		res, err := structpb.NewList([]interface{}{val})
		if err != nil {
			return err
		}
		return structpb.NewListValue(res)
	}

	rest := ctx.SubListOfDoubles().Accept(a)
	if err, ok := rest.(error); ok {
		return err
	}
	restL, ok := rest.(*structpb.Value)
	if !ok {
		return fmt.Errorf("unknown type when parsing list of doubles: %v", rest)
	}
	restL.GetListValue().Values = append([]*structpb.Value{structpb.NewNumberValue(val)}, restL.GetListValue().Values...)
	return restL
}

func (a *ASTBuilderV3) VisitListOfStrings(ctx *ListOfStringsContext) interface{} {
	return ctx.ListStrings().Accept(a)
}

func (a *ASTBuilderV3) VisitListStrings(ctx *ListStringsContext) interface{} {
	return ctx.SubListOfStrings().Accept(a)
}

func (a *ASTBuilderV3) VisitSubListOfStrings(ctx *SubListOfStringsContext) interface{} {
	val := getString(ctx.STRING().GetText())
	if ctx.SubListOfStrings() == nil || ctx.SubListOfStrings().IsEmpty() {
		res, err := structpb.NewList([]interface{}{val})
		if err != nil {
			return err
		}
		return structpb.NewListValue(res)
	}

	rest := ctx.SubListOfStrings().Accept(a)
	if err, ok := rest.(error); ok {
		return err
	}
	restL, ok := rest.(*structpb.Value)
	if !ok {
		return fmt.Errorf("unknown type when parsing list of strings: %v", rest)
	}
	restL.GetListValue().Values = append([]*structpb.Value{structpb.NewStringValue(val)}, restL.GetListValue().Values...)
	return restL
}

func (a *ASTBuilderV3) VisitListOfBooleans(ctx *ListOfBooleansContext) interface{} {
	return ctx.ListBooleans().Accept(a)
}

func (a *ASTBuilderV3) VisitListBooleans(ctx *ListBooleansContext) interface{} {
	return ctx.SubListOfBooleans().Accept(a)
}

func (a *ASTBuilderV3) VisitSubListOfBooleans(ctx *SubListOfBooleansContext) interface{} {
	val, err := strconv.ParseBool(ctx.BOOLEAN().GetText())
	if err != nil {
		return err
	}
	if ctx.SubListOfBooleans() == nil || ctx.SubListOfBooleans().IsEmpty() {
		res, err := structpb.NewList([]interface{}{val})
		if err != nil {
			return err
		}
		return structpb.NewListValue(res)
	}

	rest := ctx.SubListOfBooleans().Accept(a)
	if err, ok := rest.(error); ok {
		return err
	}
	restL, ok := rest.(*structpb.Value)
	if !ok {
		return fmt.Errorf("unknown type when parsing list of booleans: %v", rest)
	}
	restL.GetListValue().Values = append([]*structpb.Value{structpb.NewBoolValue(val)}, restL.GetListValue().Values...)
	return restL
}

func (a *ASTBuilderV3) VisitFunctionArg(ctx *FunctionArgContext) interface{} {
	if ctx.Query() != nil {
		return ctx.Query().Accept(a)
	}
	if ctx.AttrPath() != nil {
		return ctx.AttrPath().Accept(a)
	}
	if ctx.Value() != nil {
		return ctx.Value().Accept(a)
	}

	return fmt.Errorf("unknown type of function arg when parsing: %v", ctx.GetText())
}
