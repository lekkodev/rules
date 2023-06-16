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
	"encoding/json"
	"fmt"

	rules "buf.build/gen/go/lekkodev/cli/protocolbuffers/go/lekko/rules/v1beta3"
)

func RuleToString(rule *rules.Rule) (string, error) {
	switch v := rule.GetRule().(type) {
	case *rules.Rule_Atom:
		if v.Atom.ComparisonOperator == rules.ComparisonOperator_COMPARISON_OPERATOR_PRESENT {
			return fmt.Sprintf("%s pr", v.Atom.ContextKey), nil
		}
		op, err := OpToString(v.Atom.ComparisonOperator)
		if err != nil {
			return "", err
		}
		b, err := json.Marshal(rule.GetAtom().ComparisonValue)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%s %s %s", v.Atom.ContextKey, op, string(b)), nil
	case *rules.Rule_BoolConst:
		if v.BoolConst {
			return "true", nil
		}
		return "false", nil
	case *rules.Rule_LogicalExpression:
		op := ""
		switch v.LogicalExpression.GetLogicalOperator() {
		case rules.LogicalOperator_LOGICAL_OPERATOR_UNSPECIFIED:
			return "", fmt.Errorf("unknown type %T of rule %v", rule, rule)
		case rules.LogicalOperator_LOGICAL_OPERATOR_AND:
			op = "and"
		case rules.LogicalOperator_LOGICAL_OPERATOR_OR:
			op = "or"
		}
		res := ""
		for _, rule := range v.LogicalExpression.GetRules() {
			new, err := RuleToString(rule)
			if rule.GetLogicalExpression() != nil {
				// If there is a sub-logical expression, than the tree needed
				// parens.
				new = fmt.Sprintf("(%s)", new)
			}
			if err != nil {
				return "", err
			}
			if res != "" {
				res = fmt.Sprintf("%s %s %s", res, op, new)
			} else {
				res = new
			}
		}
		return res, nil
	case *rules.Rule_Not:
		inner, err := RuleToString(v.Not)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("not(%s)", inner), nil
	case *rules.Rule_CallExpression:
		switch f := rule.GetCallExpression().GetFunction().(type) {
		case *rules.CallExpression_Bucket_:
			return fmt.Sprintf("bucket(%s, %d)", f.Bucket.GetContextKey(), f.Bucket.GetThreshold()), nil
		default:
			return "", fmt.Errorf("unknown function type %T of rule %v", f, rule)
		}
	default:
		return "", fmt.Errorf("unknown type %T of rule %v", rule.GetRule(), rule.GetRule())
	}
}

func OpToString(op rules.ComparisonOperator) (string, error) {
	switch op {
	case rules.ComparisonOperator_COMPARISON_OPERATOR_UNSPECIFIED:
		return "", fmt.Errorf("invalid unspecified operator")
	case rules.ComparisonOperator_COMPARISON_OPERATOR_EQUALS:
		return "==", nil
	case rules.ComparisonOperator_COMPARISON_OPERATOR_NOT_EQUALS:
		return "!=", nil
	case rules.ComparisonOperator_COMPARISON_OPERATOR_LESS_THAN:
		return "<", nil
	case rules.ComparisonOperator_COMPARISON_OPERATOR_LESS_THAN_OR_EQUALS:
		return "<=", nil
	case rules.ComparisonOperator_COMPARISON_OPERATOR_GREATER_THAN:
		return ">", nil
	case rules.ComparisonOperator_COMPARISON_OPERATOR_GREATER_THAN_OR_EQUALS:
		return ">=", nil
	case rules.ComparisonOperator_COMPARISON_OPERATOR_CONTAINED_WITHIN:
		return "in", nil
	case rules.ComparisonOperator_COMPARISON_OPERATOR_STARTS_WITH:
		return "sw", nil
	case rules.ComparisonOperator_COMPARISON_OPERATOR_ENDS_WITH:
		return "ew", nil
	case rules.ComparisonOperator_COMPARISON_OPERATOR_CONTAINS:
		return "co", nil
	default:
		return "", fmt.Errorf("unknown operator %v", op)
	}
}
