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
	"testing"

	rulesv1beta2 "buf.build/gen/go/lekkodev/cli/protocolbuffers/go/lekko/rules/v1beta2"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

type astTestCase struct {
	rule       string
	jsonResult string
	error      bool
}

func TestASTParser(t *testing.T) {
	tests := []astTestCase{
		{
			`region IN ["USA", "EUROPE"]`,
			`{"atom": {"context_key": "region", "comparison_operator": "COMPARISON_OPERATOR_CONTAINED_WITHIN", "comparison_value": ["USA", "EUROPE"]}}`,
			false,
		},
		{
			`key pr and number < 63`,
			`{"logical_expression": {"first_rule": {"atom": {"context_key": "key", "comparison_operator": "COMPARISON_OPERATOR_PRESENT"}}, "second_rule": {"atom": {"context_key": "number", "comparison_operator": "COMPARISON_OPERATOR_LESS_THAN", "comparison_value": 63}}, "logical_operator": "LOGICAL_OPERATOR_AND"}}`,
			false,
		},
		{
			`prefix sw "pre" or suffix ew "ix" or hello == "george"`,
			`{"logical_expression": {"first_rule": {"logical_expression": { "first_rule": {"atom": {"context_key": "prefix", "comparison_operator": "COMPARISON_OPERATOR_STARTS_WITH", "comparison_value": "pre"}}, "second_rule": {"atom": {"context_key": "suffix", "comparison_operator": "COMPARISON_OPERATOR_ENDS_WITH", "comparison_value": "ix"}}, "logical_operator": "LOGICAL_OPERATOR_OR"}}, "second_rule": {"atom": {"context_key": "hello", "comparison_operator": "COMPARISON_OPERATOR_EQUALS", "comparison_value": "george"}}, "logical_operator": "LOGICAL_OPERATOR_OR"}}`,
			false,
		},
		{
			`prefix sw "pre" or suffix ew "ix" and hello == "george"`,
			`{"logical_expression": {"first_rule": {"atom": {"context_key": "prefix", "comparison_operator": "COMPARISON_OPERATOR_STARTS_WITH", "comparison_value": "pre"}}, "second_rule": {"logical_expression": { "first_rule": {"atom": {"context_key": "suffix", "comparison_operator": "COMPARISON_OPERATOR_ENDS_WITH", "comparison_value": "ix"}}, "second_rule": {"atom": {"context_key": "hello", "comparison_operator": "COMPARISON_OPERATOR_EQUALS", "comparison_value": "george"}}, "logical_operator": "LOGICAL_OPERATOR_AND"}}, "logical_operator": "LOGICAL_OPERATOR_OR"}}`,
			false,
		},
		{
			`not (prefix sw "pre" or suffix ew "ix") and hello == "george"`,
			`{"logical_expression": {"first_rule": {"not": {"logical_expression": { "first_rule": {"atom": {"context_key": "prefix", "comparison_operator": "COMPARISON_OPERATOR_STARTS_WITH", "comparison_value": "pre"}}, "second_rule": {"atom": {"context_key": "suffix", "comparison_operator": "COMPARISON_OPERATOR_ENDS_WITH", "comparison_value": "ix"}}, "logical_operator": "LOGICAL_OPERATOR_OR"}}}, "second_rule": {"atom": {"context_key": "hello", "comparison_operator": "COMPARISON_OPERATOR_EQUALS", "comparison_value": "george"}}, "logical_operator": "LOGICAL_OPERATOR_AND"}}`,
			false,
		},
		{
			`prefix sw "pre" or (suffix ew "ix" and hello == "george")`,
			`{"logical_expression": {"first_rule": {"atom": {"context_key": "prefix", "comparison_operator": "COMPARISON_OPERATOR_STARTS_WITH", "comparison_value": "pre"}}, "second_rule": {"logical_expression": { "first_rule": {"atom": {"context_key": "suffix", "comparison_operator": "COMPARISON_OPERATOR_ENDS_WITH", "comparison_value": "ix"}}, "second_rule": {"atom": {"context_key": "hello", "comparison_operator": "COMPARISON_OPERATOR_EQUALS", "comparison_value": "george"}}, "logical_operator": "LOGICAL_OPERATOR_AND"}}, "logical_operator": "LOGICAL_OPERATOR_OR"}}`,
			false,
		},
		{
			`prefix sw []`,
			"",
			true,
		},
		{
			// We should consider treating in as a contains if it applies to a string value.
			// For now, lets error.
			`prefix in "string"`,
			"",
			true,
		},
		{
			`prefix > ""`,
			"",
			true,
		},
	}
	for _, tt := range tests {
		result, err := BuildAST(tt.rule)
		if tt.error {
			require.Error(t, err, fmt.Sprintf("expected error didn't get one rule: %s", tt.rule))
		} else {
			require.NoError(t, err, "error during parsing! %s", tt.rule)
			expectedRule := &rulesv1beta2.Rule{}
			err := protojson.Unmarshal([]byte(tt.jsonResult), expectedRule)
			require.NoError(t, err, "we shouldn't get this, test malformed: %s", tt.jsonResult)
			require.True(t, proto.Equal(result, expectedRule), "results differ, got: %v, expected: %v", result, expectedRule)
		}
	}
}
