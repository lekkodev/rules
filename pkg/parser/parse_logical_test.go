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

	"github.com/stretchr/testify/require"
)

func TestNestedObject(t *testing.T) {
	tests := []testCase{
		{
			`x.a eq 1 and x.b.c eq 2`,
			map[string]interface{}{
				"x": map[string]interface{}{
					"a": 1,
					"b": map[string]interface{}{
						"c": 2.0,
					},
				},
			},
			true,
			false,
		},
		{
			`x.a eq 1 and x.b.c eq 2`,
			map[string]interface{}{
				"x": map[string]interface{}{
					"a": 1,
					"b": map[string]interface{}{
						"c": 2,
					},
				},
			},
			true,
			false,
		},
		{
			`x.a eq 1 and x.b.c eq 2`,
			map[string]interface{}{
				"x": map[string]interface{}{
					"b": map[string]interface{}{
						"c": 2,
					},
				},
			},
			false,
			false,
		},
		{
			`x.a eq 1 or x.b.c eq 2`,
			map[string]interface{}{
				"x": map[string]interface{}{
					"b": map[string]interface{}{
						"c": 2,
					},
				},
			},
			true,
			false,
		},
		{
			`x.a eq 1 or x.b.c eq 2`,
			nil,
			false,
			false,
		},
		{
			`x.a eq 1 or x.b.c eq 2`,
			obj{},
			false,
			false,
		},
		{
			`x.a eq 1 or x.b.c gt true`,
			obj{},
			false,
			true,
		},
		{
			`x.a eq 1 or x.b.c gt true`,
			map[string]interface{}{
				"x": map[string]interface{}{
					"a": 1,
				},
			},
			true,
			false,
		},
		{
			`x.b.c gt true or x.a eq 1`,
			map[string]interface{}{
				"x": map[string]interface{}{
					"a": 1,
				},
			},
			false,
			true,
		},
	}

	for _, tt := range tests {
		result, err := eval(t, tt.rule, tt.input)
		if tt.hasError {
			require.Error(t, err)
			continue
		}
		require.NoError(t, err, fmt.Sprintf("unexpected error rule: %s input: %v", tt.rule, tt.input))
		require.Equal(t, tt.result, result, fmt.Sprintf("invalid value rule: %s input: %v", tt.rule, tt.input))
	}
}

func TestLogicalExpWithAnd(t *testing.T) {
	tests := []testCase{
		{
			`x eq 1 or y gt 1`,
			obj{
				"x": 1,
				"y": 0,
			},
			true,
			false,
		},
		{
			`x eq 1 and y gt 1`,
			obj{
				"x": 1,
				"y": 2,
			},
			true,
			false,
		},
		{
			`x eq 1 and y gt 1`,
			obj{
				"x": 1,
				"y": 1,
			},
			false,
			false,
		},
		{
			`x eq 1 and not (y gt 1)`,
			obj{
				"x": 1,
				"y": 1,
			},
			true,
			false,
		},
		{
			`x eq 1 and not (y gt 1)`,
			obj{
				"x": 1,
				"y": 2,
			},
			false,
			false,
		},
		{
			`(x eq 1 and y gt 1) and z eq 3`,
			obj{
				"x": 1,
				"y": 2,
				"z": 3,
			},
			true,
			false,
		},
		{
			`(x eq 1 and y gt 1) and z eq 3 or a gt 4`,
			obj{
				"x": 1,
				"y": 2,
				"a": 5,
			},
			true,
			false,
		},
		{
			`(x eq 1 and y gt 1) and (z eq 3 or a gt 4)`,
			obj{
				"x": 1,
				"y": 2,
				"a": 5,
			},
			true,
			false,
		},
	}

	for _, tt := range tests {
		require.Equal(t, tt.result, Evaluate(tt.rule, tt.input), tt.rule, fmt.Sprintf("invalid result rule: %s input: %v", tt.rule, tt.input))
		require.Equal(t, tt.result, Evaluate(fmt.Sprintf("(%s)", tt.rule), tt.input), tt.rule)
	}
}
