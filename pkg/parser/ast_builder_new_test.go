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

	rules "buf.build/gen/go/lekkodev/cli/protocolbuffers/go/lekko/rules/v1beta3"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

type astTestCase struct {
	rule       string
	jsonResult string
	error      bool
}

func TestASTParserV3(t *testing.T) {
	tests := []astTestCase{
		{
			`region IN ['USA', 'EUROPE']`,
			`{"atom": {"context_key": "region", "comparison_operator": "COMPARISON_OPERATOR_CONTAINED_WITHIN", "comparison_value": ["USA", "EUROPE"]}}`,
			false,
		},
		{
			`region IN ["USA", "EUROPE"]`, // double quotes vs. single quotes above
			`{"atom": {"context_key": "region", "comparison_operator": "COMPARISON_OPERATOR_CONTAINED_WITHIN", "comparison_value": ["USA", "EUROPE"]}}`,
			false,
		},
		{
			`region.name IN ['USA', 'EUROPE']`, // dot expr context keys are not supported
			``,
			true,
		},
		{
			`key pr and number < 63`,
			`{"logical_expression": {"rules": [{"atom": {"context_key": "key", "comparison_operator": "COMPARISON_OPERATOR_PRESENT"}}, {"atom": {"context_key": "number", "comparison_operator": "COMPARISON_OPERATOR_LESS_THAN", "comparison_value": 63}}], "logical_operator": "LOGICAL_OPERATOR_AND"}}`,
			false,
		},
		{
			`prefix sw "pre" or suffix ew "ix" or hello == "george"`,
			`{"logical_expression": { "rules": [{"atom": {"context_key": "prefix", "comparison_operator": "COMPARISON_OPERATOR_STARTS_WITH", "comparison_value": "pre"}}, {"atom": {"context_key": "suffix", "comparison_operator": "COMPARISON_OPERATOR_ENDS_WITH", "comparison_value": "ix"}}, {"atom": {"context_key": "hello", "comparison_operator": "COMPARISON_OPERATOR_EQUALS", "comparison_value": "george"}}], "logical_operator": "LOGICAL_OPERATOR_OR"}}`,
			false,
		},
		{
			`prefix sw "pre" and suffix ew "ix" and hello == "george" and org_id in [1]`,
			`{"logical_expression": { "rules": [{"atom": {"context_key": "prefix", "comparison_operator": "COMPARISON_OPERATOR_STARTS_WITH", "comparison_value": "pre"}}, {"atom": {"context_key": "suffix", "comparison_operator": "COMPARISON_OPERATOR_ENDS_WITH", "comparison_value": "ix"}}, {"atom": {"context_key": "hello", "comparison_operator": "COMPARISON_OPERATOR_EQUALS", "comparison_value": "george"}}, {"atom": {"context_key": "org_id", "comparison_operator": "COMPARISON_OPERATOR_CONTAINED_WITHIN", "comparison_value": [1]}}], "logical_operator": "LOGICAL_OPERATOR_AND"}}`,
			false,
		},
		{
			`prefix sw "pre" or suffix ew "ix" and hello == "george"`,
			`{"logical_expression": {"rules": [{"atom": {"context_key": "prefix", "comparison_operator": "COMPARISON_OPERATOR_STARTS_WITH", "comparison_value": "pre"}}, {"logical_expression": { "rules": [{"atom": {"context_key": "suffix", "comparison_operator": "COMPARISON_OPERATOR_ENDS_WITH", "comparison_value": "ix"}}, {"atom": {"context_key": "hello", "comparison_operator": "COMPARISON_OPERATOR_EQUALS", "comparison_value": "george"}}], "logical_operator": "LOGICAL_OPERATOR_AND"}}], "logical_operator": "LOGICAL_OPERATOR_OR"}}`,
			false,
		},
		{
			`not (prefix sw "pre" or suffix ew "ix") and hello == "george"`,
			`{"logical_expression": {"rules": [{"not": {"logical_expression": { "rules": [{"atom": {"context_key": "prefix", "comparison_operator": "COMPARISON_OPERATOR_STARTS_WITH", "comparison_value": "pre"}}, {"atom": {"context_key": "suffix", "comparison_operator": "COMPARISON_OPERATOR_ENDS_WITH", "comparison_value": "ix"}}], "logical_operator": "LOGICAL_OPERATOR_OR"}}}, {"atom": {"context_key": "hello", "comparison_operator": "COMPARISON_OPERATOR_EQUALS", "comparison_value": "george"}}], "logical_operator": "LOGICAL_OPERATOR_AND"}}`,
			false,
		},
		{
			`prefix sw "pre" or (suffix ew "ix" and hello == "george")`,
			`{"logical_expression": {"rules": [{"atom": {"context_key": "prefix", "comparison_operator": "COMPARISON_OPERATOR_STARTS_WITH", "comparison_value": "pre"}}, {"logical_expression": { "rules": [{"atom": {"context_key": "suffix", "comparison_operator": "COMPARISON_OPERATOR_ENDS_WITH", "comparison_value": "ix"}}, {"atom": {"context_key": "hello", "comparison_operator": "COMPARISON_OPERATOR_EQUALS", "comparison_value": "george"}}], "logical_operator": "LOGICAL_OPERATOR_AND"}}], "logical_operator": "LOGICAL_OPERATOR_OR"}}`,
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
		{
			// This function is not in the supported options
			`not_impl_func(id)`,
			"",
			true,
		},
		{
			`bucket(id, 0)`,
			`{"call_expression": {"bucket": {"context_key": "id", "threshold": 0}}}`,
			false,
		},
		{
			`bucket(id, 75.125)`,
			`{"call_expression": {"bucket": {"context_key": "id", "threshold": 75125}}}`,
			false,
		},
		{
			// Not enough args
			`bucket(id)`,
			"",
			true,
		},
		{
			// Too many args
			`bucket(id, 10000, "extra")`,
			"",
			true,
		},
		{
			`bucket("wrong", "types")`,
			"",
			true,
		},
		{
			// Out of range, negative
			`bucket(id, -1)`,
			"",
			true,
		},
		{
			// Out of range, > 100
			`bucket(id, 100.001)`,
			"",
			true,
		},
	}
	for _, tt := range tests {
		result, err := BuildASTV3(tt.rule)
		if tt.error {
			require.Error(t, err, fmt.Sprintf("expected error didn't get one rule: %s", tt.rule))
		} else {
			require.NoError(t, err, "error during parsing! %s", tt.rule)
			expectedRule := &rules.Rule{}
			err := protojson.Unmarshal([]byte(tt.jsonResult), expectedRule)
			require.NoError(t, err, "we shouldn't get this, test malformed: %s", tt.jsonResult)
			require.True(t, proto.Equal(result, expectedRule), "results differ, got: %v, expected: %v", result, expectedRule)
		}
	}
}
