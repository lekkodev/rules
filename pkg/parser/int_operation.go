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
	"reflect"
)

type IntOperation struct {
	NullOperation
}

func (o *IntOperation) EQ(left Operand, right Operand) (bool, error) {
	return applyWithTransform[int, int](left, right, toInt, toInt, func(l int, r int) bool { return l == r })
}

func (o *IntOperation) NE(left Operand, right Operand) (bool, error) {
	return applyWithTransform[int, int](left, right, toInt, toInt, func(l int, r int) bool { return l != r })
}

func (o *IntOperation) GT(left Operand, right Operand) (bool, error) {
	return applyWithTransform[int, int](left, right, toInt, toInt, func(l int, r int) bool { return l > r })
}

func (o *IntOperation) LT(left Operand, right Operand) (bool, error) {
	return applyWithTransform[int, int](left, right, toInt, toInt, func(l int, r int) bool { return l < r })
}

func (o *IntOperation) GE(left Operand, right Operand) (bool, error) {
	return applyWithTransform[int, int](left, right, toInt, toInt, func(l int, r int) bool { return l >= r })
}

func (o *IntOperation) LE(left Operand, right Operand) (bool, error) {
	return applyWithTransform[int, int](left, right, toInt, toInt, func(l int, r int) bool { return l <= r })

}

func (o *IntOperation) IN(left Operand, right Operand) (bool, error) {
	return applyWithTransform[int, []int](left, right, toInt, func(right Operand) ([]int, error) {
		rv := reflect.ValueOf(right)
		if rv.Kind() == reflect.Slice {
			ret := make([]int, rv.Len())
			for i := 0; i < rv.Len(); i = i + 1 {
				val, err := toInt(Operand(rv.Index(i).Interface()))
				if err != nil {
					return nil, err
				}
				ret[i] = val
			}
			return ret, nil
		} else {
			return nil, newErrInvalidOperand(right, []int{})
		}
	},
		func(l int, r []int) bool {
			for i := range r {
				if l == r[i] {
					return true
				}
			}
			return false
		})
}
