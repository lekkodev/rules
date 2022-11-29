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
	"strings"
)

type StringOperation struct {
	NullOperation
}

func (o *StringOperation) EQ(left Operand, right Operand) (bool, error) {
	return apply[string, string](left, right, func(l string, r string) bool { return l == r })
}

func (o *StringOperation) NE(left Operand, right Operand) (bool, error) {
	return apply[string, string](left, right, func(l string, r string) bool { return l != r })
}

func (o *StringOperation) GT(left Operand, right Operand) (bool, error) {
	return apply[string, string](left, right, func(l string, r string) bool { return l > r })
}

func (o *StringOperation) LT(left Operand, right Operand) (bool, error) {
	return apply[string, string](left, right, func(l string, r string) bool { return l < r })
}

func (o *StringOperation) GE(left Operand, right Operand) (bool, error) {
	return apply[string, string](left, right, func(l string, r string) bool { return l >= r })
}

func (o *StringOperation) LE(left Operand, right Operand) (bool, error) {
	return apply[string, string](left, right, func(l string, r string) bool { return l <= r })
}

func (o *StringOperation) CO(left Operand, right Operand) (bool, error) {
	return apply[string, string](left, right, func(l string, r string) bool { return strings.Contains(l, r) })
}

func (o *StringOperation) SW(left Operand, right Operand) (bool, error) {
	return apply[string, string](left, right, func(l string, r string) bool { return strings.HasPrefix(l, r) })
}

func (o *StringOperation) EW(left Operand, right Operand) (bool, error) {
	return apply[string, string](left, right, func(l string, r string) bool { return strings.HasSuffix(l, r) })
}

func (o *StringOperation) IN(left Operand, right Operand) (bool, error) {
	return apply[string, []string](left, right, func(l string, r []string) bool {
		for i := range r {
			if l == r[i] {
				return true
			}
		}
		return false
	})
}
