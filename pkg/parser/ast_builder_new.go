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

	rules "buf.build/gen/go/lekkodev/cli/protocolbuffers/go/lekko/rules/v1beta3"

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

func BuildASTV3(rule string) (*rules.Rule, error) {
	tree, err := lexAndParse(rule)
	if err != nil {
		return nil, err
	}
	switch v := NewASTBuilderV3().Visit(tree).(type) {
	case error:
		return nil, v
	case *rules.Rule:
		return v, nil
	default:
		return nil, fmt.Errorf("unknown type during AST building: %v %T", v, v)
	}

}

func (a *ASTBuilderV3) Visit(tree antlr.ParseTree) interface{} {
	val, ok := tree.(*QueryContext)
	if !ok {
		return fmt.Errorf("invalid tree type: %v", val)
	}
	return val.Accept(a)
}

func (a *ASTBuilderV3) VisitQuery(ctx *QueryContext) interface{} {
	switch val := ctx.Subquery().(type) {
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
		return fmt.Errorf("invalid subquery type: %v", val)
	}
}

func (a *ASTBuilderV3) VisitParenExp(ctx *ParenExpContext) interface{} {
	v := ctx.Subquery().Accept(a)
	if err, ok := v.(error); ok {
		return err
	}
	r, ok := v.(*rules.Rule)
	if !ok {
		return fmt.Errorf("invalid type during AST building: %v", v)
	}
	if ctx.NOT() != nil {
		return &rules.Rule{Rule: &rules.Rule_Not{Not: r}}
	}
	return v
}

func (a *ASTBuilderV3) VisitAndLogicalExp(ctx *AndLogicalExpContext) interface{} {
	left := ctx.Subquery(0).Accept(a)
	if err, ok := left.(error); ok {
		return err
	}
	leftR, ok := left.(*rules.Rule)
	if !ok {
		return fmt.Errorf("invalid type during AST building: %v", left)
	}

	right := ctx.Subquery(1).Accept(a)
	if err, ok := right.(error); ok {
		return err
	}
	rightR, ok := right.(*rules.Rule)
	if !ok {
		return fmt.Errorf("invalid type during AST building: %v", right)
	}

	// special case the left being an already constructed andlogicalexp for n-ary trees.
	if leftExpr := leftR.GetLogicalExpression(); leftExpr != nil {
		if leftExpr.LogicalOperator == rules.LogicalOperator_LOGICAL_OPERATOR_AND {
			leftExpr.Rules = append(leftExpr.Rules, rightR)
			return left
		}
	}

	return &rules.Rule{
		Rule: &rules.Rule_LogicalExpression{
			LogicalExpression: &rules.LogicalExpression{
				Rules:           []*rules.Rule{leftR, rightR},
				LogicalOperator: rules.LogicalOperator_LOGICAL_OPERATOR_AND,
			},
		},
	}
}

func (a *ASTBuilderV3) VisitOrLogicalExp(ctx *OrLogicalExpContext) interface{} {
	left := ctx.Subquery(0).Accept(a)
	if err, ok := left.(error); ok {
		return err
	}
	leftR, ok := left.(*rules.Rule)
	if !ok {
		return fmt.Errorf("invalid type during AST building: %v", left)
	}

	right := ctx.Subquery(1).Accept(a)
	if err, ok := right.(error); ok {
		return err
	}
	rightR, ok := right.(*rules.Rule)
	if !ok {
		return fmt.Errorf("invalid type during AST building: %v", right)
	}

	// special case the left being an already constructed orlogicalexp for n-ary trees.
	if leftExpr := leftR.GetLogicalExpression(); leftExpr != nil {
		if leftExpr.LogicalOperator == rules.LogicalOperator_LOGICAL_OPERATOR_OR {
			leftExpr.Rules = append(leftExpr.Rules, rightR)
			return left
		}
	}

	return &rules.Rule{
		Rule: &rules.Rule_LogicalExpression{
			LogicalExpression: &rules.LogicalExpression{
				Rules:           []*rules.Rule{leftR, rightR},
				LogicalOperator: rules.LogicalOperator_LOGICAL_OPERATOR_OR,
			},
		},
	}
}

func (a *ASTBuilderV3) VisitPresentExp(ctx *PresentExpContext) interface{} {
	return &rules.Rule{
		Rule: &rules.Rule_Atom{
			Atom: &rules.Atom{
				ContextKey:         ctx.AttrPath().Accept(a).(string),
				ComparisonOperator: rules.ComparisonOperator_COMPARISON_OPERATOR_PRESENT,
			},
		},
	}
}

func (a *ASTBuilderV3) VisitCompareExp(ctx *CompareExpContext) (ret interface{}) {
	key := ctx.AttrPath().Accept(a)
	if err, ok := key.(error); ok {
		return err
	}
	ctxKey, ok := key.(string)
	if !ok {
		return fmt.Errorf("unknown type when visiting cmp expr: %v", key)
	}
	value := ctx.Value().Accept(a)

	if err, ok := value.(error); ok {
		return err
	}
	valueR, ok := value.(*structpb.Value)
	if !ok {
		return fmt.Errorf("invalid type during AST building: %v", value)
	}

	atom := &rules.Atom{
		ContextKey:      ctxKey,
		ComparisonValue: valueR,
	}
	// TODO: Do a type check on each righthanded operator.
	switch ctx.op.GetTokenType() {
	case JsonQueryParserEQ:
		atom.ComparisonOperator = rules.ComparisonOperator_COMPARISON_OPERATOR_EQUALS
	case JsonQueryParserNE:
		atom.ComparisonOperator = rules.ComparisonOperator_COMPARISON_OPERATOR_NOT_EQUALS
	case JsonQueryParserGT:
		atom.ComparisonOperator = rules.ComparisonOperator_COMPARISON_OPERATOR_GREATER_THAN
		if _, ok := valueR.GetKind().(*structpb.Value_NumberValue); !ok {
			return fmt.Errorf("invalid type for operator %v %T", atom.ComparisonOperator, valueR)
		}
	case JsonQueryParserLT:
		atom.ComparisonOperator = rules.ComparisonOperator_COMPARISON_OPERATOR_LESS_THAN
		if _, ok := valueR.GetKind().(*structpb.Value_NumberValue); !ok {
			return fmt.Errorf("invalid type for operator %v %T", atom.ComparisonOperator, valueR)
		}
	case JsonQueryParserLE:
		atom.ComparisonOperator = rules.ComparisonOperator_COMPARISON_OPERATOR_LESS_THAN_OR_EQUALS
		if _, ok := valueR.GetKind().(*structpb.Value_NumberValue); !ok {
			return fmt.Errorf("invalid type for operator %v %T", atom.ComparisonOperator, valueR)
		}
	case JsonQueryParserGE:
		atom.ComparisonOperator = rules.ComparisonOperator_COMPARISON_OPERATOR_GREATER_THAN_OR_EQUALS
		if _, ok := valueR.GetKind().(*structpb.Value_NumberValue); !ok {
			return fmt.Errorf("invalid type for operator %v %T", atom.ComparisonOperator, valueR)
		}
	case JsonQueryParserCO:
		atom.ComparisonOperator = rules.ComparisonOperator_COMPARISON_OPERATOR_CONTAINS
		if _, ok := valueR.GetKind().(*structpb.Value_StringValue); !ok {
			return fmt.Errorf("invalid type for operator %v %T", atom.ComparisonOperator, valueR)
		}
	case JsonQueryParserSW:
		atom.ComparisonOperator = rules.ComparisonOperator_COMPARISON_OPERATOR_STARTS_WITH
		if _, ok := valueR.GetKind().(*structpb.Value_StringValue); !ok {
			return fmt.Errorf("invalid type for operator %v %T", atom.ComparisonOperator, valueR)
		}
	case JsonQueryParserEW:
		atom.ComparisonOperator = rules.ComparisonOperator_COMPARISON_OPERATOR_ENDS_WITH
		if _, ok := valueR.GetKind().(*structpb.Value_StringValue); !ok {
			return fmt.Errorf("invalid type for operator %v %T", atom.ComparisonOperator, valueR)
		}
	case JsonQueryParserIN:
		atom.ComparisonOperator = rules.ComparisonOperator_COMPARISON_OPERATOR_CONTAINED_WITHIN
		if _, ok := valueR.GetKind().(*structpb.Value_ListValue); !ok {
			return fmt.Errorf("invalid type for operator %v %T", atom.ComparisonOperator, valueR)
		}
	default:
		return fmt.Errorf("invalid token: %v", ctx.op.GetTokenType())
	}
	return &rules.Rule{
		Rule: &rules.Rule_Atom{
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
	return fmt.Errorf("invalid attribute '%s': please remove period", ctx.ATTRNAME().GetText())
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
	val, err := strconv.ParseFloat(ctx.GetText(), 64)
	if err != nil {
		return err
	}
	return structpb.NewNumberValue(val)
}

func (a *ASTBuilderV3) VisitListOfNumbers(ctx *ListOfNumbersContext) interface{} {
	return ctx.ListNumbers().Accept(a)
}

func (a *ASTBuilderV3) VisitListNumbers(ctx *ListNumbersContext) interface{} {
	return ctx.SubListOfNumbers().Accept(a)
}

func (a *ASTBuilderV3) VisitSubListOfNumbers(ctx *SubListOfNumbersContext) interface{} {
	var val float64
	var err error
	switch ctx.num.GetTokenType() {
	case JsonQueryParserLONG:
		val, err = strconv.ParseFloat(ctx.LONG().GetText(), 64)
		if err != nil {
			return err
		}
	case JsonQueryParserDOUBLE:
		val, err = strconv.ParseFloat(ctx.DOUBLE().GetText(), 64)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("invalid token: %v", ctx.num.GetTokenType())
	}
	if ctx.SubListOfNumbers() == nil || ctx.SubListOfNumbers().IsEmpty() {
		res, err := structpb.NewList([]interface{}{val})
		if err != nil {
			return err
		}
		return structpb.NewListValue(res)
	}

	rest := ctx.SubListOfNumbers().Accept(a)
	if err, ok := rest.(error); ok {
		return err
	}
	restL, ok := rest.(*structpb.Value)
	if !ok {
		return fmt.Errorf("unknown type when parsing list of numbers: %v", rest)
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
	if ctx.Subquery() != nil {
		return ctx.Subquery().Accept(a)
	}
	if ctx.AttrPath() != nil {
		return ctx.AttrPath().Accept(a)
	}
	if ctx.Value() != nil {
		return ctx.Value().Accept(a)
	}

	return fmt.Errorf("unknown type of function arg when parsing: %v", ctx.GetText())
}
