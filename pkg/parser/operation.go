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
	"errors"
	"fmt"
	"reflect"
)

func apply[L Operand, R Operand](left Operand, right Operand, applyFunc func(L, R) bool) (bool, error) {
	if left == nil {
		// We may get here if the context value isn't specified. To handle this case,
		// we are going to set the left value to a zero value. This will result
		// in behavior as if the left value was a zero value.
		// THIS IS DANGEROUS FOR PRIMITIVES. We will try to handle primitives manually.
		var l L
		// This is a safe cast since we declared L to be an Operand.
		left = Operand(l)
	}
	l, r, err := get[L, R](left, right)
	if err != nil {
		return false, err
	}
	return applyFunc(l, r), nil
}

func applyWithTransform[L Operand, R Operand](left Operand, right Operand, leftTransform func(left Operand) (L, error), rightTransform func(right Operand) (R, error), applyFunc func(L, R) bool) (bool, error) {
	if left == nil {
		// We may get here if the context value isn't specified. To handle this case,
		// we are going to set the left value to a zero value. This will result
		// in behavior as if the left value was a zero value.
		// THIS IS DANGEROUS FOR PRIMITIVES. We will try to handle primitives manually.
		var l L
		// This is a safe cast since we declared L to be an Operand.
		left = Operand(l)
	}
	l, err := leftTransform(left)
	if err != nil {
		return false, err
	}

	r, err := rightTransform(right)
	if err != nil {
		return false, err
	}

	return applyFunc(l, r), nil
}

func get[L Operand, R Operand](left Operand, right Operand) (L, R, error) {
	// we need to instantiate these default values to get valid return values
	var leftVal L
	var rightVal R
	var ok bool
	leftVal, ok = left.(L)
	if !ok {
		return leftVal, rightVal, newErrInvalidOperand(left, leftVal)
	}
	rightVal, ok = right.(R)
	if !ok {
		return leftVal, rightVal, newErrInvalidOperand(right, rightVal)
	}
	return leftVal, rightVal, nil
}

func toFloat(op Operand) (float64, error) {
	switch val := op.(type) {
	case int:
		return float64(val), nil
	case float64:
		return val, nil
	}
	var exp float64
	return 0, newErrInvalidOperand(op, exp)
}

func toInt(op Operand) (int, error) {
	switch val := op.(type) {
	case int:
		return val, nil
	case float64:
		return int(val), nil
	case int32:
		return int(val), nil
	case int64:
		return int(val), nil
	}
	var exp int
	return 0, newErrInvalidOperand(op, exp)
}

type Operand interface{}

var (
	ErrInvalidOperation   = errors.New("Invalid operation on the type")
	ErrEvalOperandMissing = errors.New("Operand not present")
)

type ErrInvalidOperand struct {
	Val     interface{}
	typeObj interface{}
}

func newErrInvalidOperand(val Operand, typeObj interface{}) *ErrInvalidOperand {
	return &ErrInvalidOperand{
		Val:     val,
		typeObj: typeObj,
	}
}

func (e *ErrInvalidOperand) Error() string {
	return fmt.Sprintf("Operand %v is not the correc type. Expected: %s, Actual: %s",
		e.Val,
		reflect.TypeOf(e.typeObj).String(),
		reflect.TypeOf(e.Val).String(),
	)
}

type Operation interface {
	EQ(left Operand, right Operand) (bool, error)
	NE(left Operand, right Operand) (bool, error)
	GT(left Operand, right Operand) (bool, error)
	LT(left Operand, right Operand) (bool, error)
	GE(left Operand, right Operand) (bool, error)
	LE(left Operand, right Operand) (bool, error)
	CO(left Operand, right Operand) (bool, error)
	SW(left Operand, right Operand) (bool, error)
	EW(left Operand, right Operand) (bool, error)
	IN(left Operand, right Operand) (bool, error)
}
