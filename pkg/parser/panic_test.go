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
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNestedError(t *testing.T) {
	err := errors.New("some random error")
	ne := newNestedError(newNestedError(err, "a"), "b")
	require.EqualValues(t, ne.Original(), err)
	m := make(map[string]interface{})
	require.NoError(t, json.Unmarshal([]byte(ne.Error()), &m))
	require.True(t, len(m) > 0)
}

func TestEvaluatorPanic(t *testing.T) {
	ev, err := NewEvaluator(`x eq 1`)
	require.NoError(t, err)
	ev.testHookPanic = func() {
		panic("wait what")
	}
	ret, err := ev.Process(map[string]interface{}{"x": 1})
	require.False(t, ret)
	require.Error(t, err)
}
