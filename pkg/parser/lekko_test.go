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

// Specific tests for lekko use cases.
func TestLekko(t *testing.T) {
	tests := []testCase{
		{
			`region IN ["USA", "EUROPE"]`,
			obj{
				"region": "USA",
			},
			true,
			false,
		},
		{
			`region IN ['USA', 'EUROPE']`,
			obj{
				"region": "USA",
			},
			false,
			true,
		},
		{
			`region IN ["USA", "EUROPE"]`,
			obj{
				"region": "blah",
			},
			false,
			false,
		},
	}
	for _, tt := range tests {
		result, err := eval(t, tt.rule, tt.input)
		if tt.hasError {
			require.Error(t, err, fmt.Sprintf("expected error didn't get one rule: %s input: %v", tt.rule, tt.input))
			continue
		} else {
			require.NoError(t, err, fmt.Sprintf("unexpected error for rule: %s input: %v", tt.rule, tt.input))
			require.Equal(t, result, tt.result, fmt.Sprintf("rule: %s, input :%v", tt.rule, tt.input))
		}
	}
}
