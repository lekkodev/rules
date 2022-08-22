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
