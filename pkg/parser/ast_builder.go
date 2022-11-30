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

	rulesv1beta2 "github.com/lekkodev/cli/pkg/gen/proto/go/lekko/rules/v1beta2"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"google.golang.org/protobuf/types/known/structpb"
)

type ASTBuilder struct {
	antlr.ParseTreeVisitor
}

func NewASTBuilder() *ASTBuilder {
	return &ASTBuilder{}
}

func BuildAST(rule string) (*rulesv1beta2.Rule, error) {
	tree, err := lexAndParse(rule)
	if err != nil {
		return nil, err
	}
	switch v := NewASTBuilder().Visit(tree).(type) {
	case error:
		return nil, v
	case *rulesv1beta2.Rule:
		return v, nil
	default:
		return nil, fmt.Errorf("unknown type during AST building: %v", v)
	}

}

func (a *ASTBuilder) Visit(tree antlr.ParseTree) interface{} {
	fmt.Println("visiting")
	switch val := tree.(type) {
	case *LogicalExpContext:
		return val.Accept(a)
	case *CompareExpContext:
		return val.Accept(a)
	case *ParenExpContext:
		return val.Accept(a)
	case *PresentExpContext:
		return val.Accept(a)
	default:
		return fmt.Errorf("invalid tree type: %v", val)
	}
}

func (a *ASTBuilder) VisitParenExp(ctx *ParenExpContext) interface{} {
	v := ctx.Query().Accept(a)
	if err, ok := v.(error); ok {
		return err
	}
	r, ok := v.(*rulesv1beta2.Rule)
	if !ok {
		return fmt.Errorf("invalid type during AST building: %v", v)
	}
	if ctx.NOT() != nil {
		return &rulesv1beta2.Rule{Rule: &rulesv1beta2.Rule_Not{Not: r}}
	}
	return v
}

func (a *ASTBuilder) VisitLogicalExp(ctx *LogicalExpContext) interface{} {
	fmt.Println("visiting logical")
	left := ctx.Query(0).Accept(a)
	if err, ok := left.(error); ok {
		return err
	}
	leftR, ok := left.(*rulesv1beta2.Rule)
	if !ok {
		return fmt.Errorf("invalid type during AST building: %v", left)
	}

	right := ctx.Query(1).Accept(a)
	if err, ok := right.(error); ok {
		return err
	}
	rightR, ok := right.(*rulesv1beta2.Rule)
	if !ok {
		return fmt.Errorf("invalid type during AST building: %v", right)
	}
	op := ctx.LOGICAL_OPERATOR().GetText()
	exp := &rulesv1beta2.Rule_LogicalExpression{
		LogicalExpression: &rulesv1beta2.LogicalExpression{
			FirstRule:  leftR,
			SecondRule: rightR,
		},
	}

	switch op {
	case "or":
		exp.LogicalExpression.LogicalOperator = rulesv1beta2.LogicalOperator_LOGICAL_OPERATOR_OR
	case "and":
		exp.LogicalExpression.LogicalOperator = rulesv1beta2.LogicalOperator_LOGICAL_OPERATOR_AND
	}
	return &rulesv1beta2.Rule{Rule: exp}
}

func (a *ASTBuilder) VisitPresentExp(ctx *PresentExpContext) interface{} {
	return &rulesv1beta2.Rule{
		Rule: &rulesv1beta2.Rule_Atom{
			Atom: &rulesv1beta2.Atom{
				ContextKey:         ctx.AttrPath().Accept(a).(string),
				ComparisonOperator: rulesv1beta2.ComparisonOperator_COMPARISON_OPERATOR_PRESENT,
			},
		},
	}
}

func (a *ASTBuilder) VisitCompareExp(ctx *CompareExpContext) (ret interface{}) {
	fmt.Println("visiting compare")
	key := ctx.AttrPath().Accept(a).(string)
	value := ctx.Value().Accept(a)

	fmt.Printf("got: %v %v\n", key, value)

	if err, ok := value.(error); ok {
		return err
	}
	valueR, ok := value.(*structpb.Value)
	if !ok {
		return fmt.Errorf("invalid type during AST building: %v", value)
	}

	atom := &rulesv1beta2.Atom{
		ContextKey:      key,
		ComparisonValue: valueR,
	}
	// TODO: Do a type check on each righthanded operator.
	switch ctx.op.GetTokenType() {
	case JsonQueryParserEQ:
		atom.ComparisonOperator = rulesv1beta2.ComparisonOperator_COMPARISON_OPERATOR_EQUALS
	case JsonQueryParserNE:
		// We need to special case not equal to return equals with a surrounding not.
		atom.ComparisonOperator = rulesv1beta2.ComparisonOperator_COMPARISON_OPERATOR_EQUALS
		return &rulesv1beta2.Rule{
			Rule: &rulesv1beta2.Rule_Not{
				Not: &rulesv1beta2.Rule{
					Rule: &rulesv1beta2.Rule_Atom{
						Atom: atom,
					},
				},
			},
		}
	case JsonQueryParserGT:
		atom.ComparisonOperator = rulesv1beta2.ComparisonOperator_COMPARISON_OPERATOR_GREATER_THAN
		if _, ok := valueR.GetKind().(*structpb.Value_NumberValue); !ok {
			return fmt.Errorf("invalid type for operator %v %T", atom.ComparisonOperator, valueR)
		}
	case JsonQueryParserLT:
		atom.ComparisonOperator = rulesv1beta2.ComparisonOperator_COMPARISON_OPERATOR_LESS_THAN
		if _, ok := valueR.GetKind().(*structpb.Value_NumberValue); !ok {
			return fmt.Errorf("invalid type for operator %v %T", atom.ComparisonOperator, valueR)
		}
	case JsonQueryParserLE:
		atom.ComparisonOperator = rulesv1beta2.ComparisonOperator_COMPARISON_OPERATOR_LESS_THAN_OR_EQUALS
		if _, ok := valueR.GetKind().(*structpb.Value_NumberValue); !ok {
			return fmt.Errorf("invalid type for operator %v %T", atom.ComparisonOperator, valueR)
		}
	case JsonQueryParserGE:
		atom.ComparisonOperator = rulesv1beta2.ComparisonOperator_COMPARISON_OPERATOR_GREATER_THAN
		if _, ok := valueR.GetKind().(*structpb.Value_NumberValue); !ok {
			return fmt.Errorf("invalid type for operator %v %T", atom.ComparisonOperator, valueR)
		}
	case JsonQueryParserCO:
		atom.ComparisonOperator = rulesv1beta2.ComparisonOperator_COMPARISON_OPERATOR_CONTAINS
		if _, ok := valueR.GetKind().(*structpb.Value_StringValue); !ok {
			return fmt.Errorf("invalid type for operator %v %T", atom.ComparisonOperator, valueR)
		}
	case JsonQueryParserSW:
		atom.ComparisonOperator = rulesv1beta2.ComparisonOperator_COMPARISON_OPERATOR_STARTS_WITH
		if _, ok := valueR.GetKind().(*structpb.Value_StringValue); !ok {
			return fmt.Errorf("invalid type for operator %v %T", atom.ComparisonOperator, valueR)
		}
	case JsonQueryParserEW:
		atom.ComparisonOperator = rulesv1beta2.ComparisonOperator_COMPARISON_OPERATOR_ENDS_WITH
		if _, ok := valueR.GetKind().(*structpb.Value_StringValue); !ok {
			return fmt.Errorf("invalid type for operator %v %T", atom.ComparisonOperator, valueR)
		}
	case JsonQueryParserIN:
		atom.ComparisonOperator = rulesv1beta2.ComparisonOperator_COMPARISON_OPERATOR_CONTAINED_WITHIN
		if _, ok := valueR.GetKind().(*structpb.Value_ListValue); !ok {
			return fmt.Errorf("invalid type for operator %v %T", atom.ComparisonOperator, valueR)
		}
	default:
		return fmt.Errorf("invalid token: %v", ctx.op.GetTokenType())
	}
	fmt.Printf("got to the end: %v\n", atom)
	return &rulesv1beta2.Rule{
		Rule: &rulesv1beta2.Rule_Atom{
			Atom: atom,
		},
	}
}

func (a *ASTBuilder) VisitAttrPath(ctx *AttrPathContext) interface{} {
	if ctx.SubAttr() == nil || ctx.SubAttr().IsEmpty() {
		return ctx.ATTRNAME().GetText()
	}
	return ctx.ATTRNAME().GetText() + ctx.SubAttr().Accept(a).(string)
}

func (a *ASTBuilder) VisitSubAttr(ctx *SubAttrContext) interface{} {
	return ctx.AttrPath().Accept(a)
}

func (a *ASTBuilder) VisitBoolean(ctx *BooleanContext) interface{} {
	val, err := strconv.ParseBool(ctx.GetText())
	if err != nil {
		return err
	}
	return structpb.NewBoolValue(val)
}

func (a *ASTBuilder) VisitNull(ctx *NullContext) interface{} {
	// TODO: not sure what to do here or if we even want to support it.
	return nil
}

func (a *ASTBuilder) VisitString(ctx *StringContext) interface{} {
	return structpb.NewStringValue(getString(ctx.GetText()))
}

func (a *ASTBuilder) VisitDouble(ctx *DoubleContext) interface{} {
	val, err := strconv.ParseFloat(ctx.GetText(), 64)
	if err != nil {
		return err
	}
	return structpb.NewNumberValue(val)
}

func (a *ASTBuilder) VisitVersion(ctx *VersionContext) interface{} {
	return structpb.NewStringValue(ctx.VERSION().GetText())
}

func (a *ASTBuilder) VisitLong(ctx *LongContext) interface{} {
	val, err := strconv.ParseInt(ctx.GetText(), 10, 64)
	if err != nil {
		return err
	}
	return structpb.NewNumberValue(float64(val))
}

func (a *ASTBuilder) VisitListOfInts(ctx *ListOfIntsContext) interface{} {
	return ctx.ListInts().Accept(a)
}

func (a *ASTBuilder) VisitListInts(ctx *ListIntsContext) interface{} {
	return ctx.SubListOfInts().Accept(a)
}

func (a *ASTBuilder) VisitSubListOfInts(ctx *SubListOfIntsContext) interface{} {
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

func (a *ASTBuilder) VisitListOfDoubles(ctx *ListOfDoublesContext) interface{} {
	return ctx.ListDoubles().Accept(a)
}

func (a *ASTBuilder) VisitListDoubles(ctx *ListDoublesContext) interface{} {
	return ctx.SubListOfDoubles().Accept(a)
}

func (a *ASTBuilder) VisitSubListOfDoubles(ctx *SubListOfDoublesContext) interface{} {
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

func (a *ASTBuilder) VisitListOfStrings(ctx *ListOfStringsContext) interface{} {
	return ctx.ListStrings().Accept(a)
}

func (a *ASTBuilder) VisitListStrings(ctx *ListStringsContext) interface{} {
	return ctx.SubListOfStrings().Accept(a)
}

func (a *ASTBuilder) VisitSubListOfStrings(ctx *SubListOfStringsContext) interface{} {
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
