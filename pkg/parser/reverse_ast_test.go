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
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRoundTrip(t *testing.T) {
	cases := []string{
		"a == 1 and (b == 2 or c == 3)",
		"a != 2",
		"d == 3",
		"f pr and f != 3",
		"h < 5",
		"j > 5",
		"l <= 5",
		"n >= 5",
		"p co \"foo\"",
		"q sw \"hello\"",
		"r ew \"bar\"",
		"s in [1,2,3]",
		"t pr",
		"u pr and not(u == 3)",
		"u pr and u != 3",
		"v == 1 or (w == 2 and x == 3)",
		"bucket(a, 10)",
		"bucket(a, 7.75)",
		"evaluate_to(\"segments\", \"segment_name\")",
		"x in [1,2.5,3.14,-10]",
	}
	for _, tc := range cases {
		ast, err := BuildASTV3(tc)
		require.NoError(t, err, "failed to build ast %s", tc)
		str, err := RuleToString(ast)
		require.NoError(t, err, "failed to go from ast to string %s %v", tc, ast)
		require.Equal(t, tc, str, "strings don't match, test case: %s, round trip: %s, ast: %v", tc, str, ast)
	}
}

func TestFMT(t *testing.T) {
	cases := [][]string{
		{"a == 1 and b == 2 or c == 3", "(a == 1 and b == 2) or c == 3"},
		{"(g pr and g ne 3)", "g pr and g != 3"},
		{"i lt 5", "i < 5"},
		{"k gt 5", "k > 5"},
		{"m le 5", "m <= 5"},
		{"o ge 5", "o >= 5"},
		{"(NOT(v == 3 and NOT (a == 1)))", "not(v == 3 and not(a == 1))"},
		{"(!(v == 3 && NOT (a == 1 || b != 4)))", "not(v == 3 and not(a == 1 or b != 4))"},
	}
	for _, tc := range cases {
		ast, err := BuildASTV3(tc[0])
		require.NoError(t, err, "failed to build ast %s", tc[0])
		str, err := RuleToString(ast)
		require.NoError(t, err, "failed to go from ast to string %s %v", tc[0], ast)
		require.Equal(t, tc[1], str, "strings don't match, test case: %s, round trip: %s, ast: %v", tc[1], str, ast)
	}
}

func TestNoHTMLEscape(t *testing.T) {
	ast, err := BuildASTV3("foo == \"<bar>\"")
	require.NoError(t, err)
	str, err := RuleToString(ast)
	require.NoError(t, err)
	require.Equal(t, "foo == \"<bar>\"", str)
}
