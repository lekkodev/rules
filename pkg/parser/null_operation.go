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

type NullOperation struct {
}

func (o *NullOperation) EQ(left Operand, right Operand) (bool, error) {
	return left == nil, nil
}

func (o *NullOperation) NE(left Operand, right Operand) (bool, error) {
	return left != nil, nil
}

func (o *NullOperation) GT(left Operand, right Operand) (bool, error) {
	return false, ErrInvalidOperation
}

func (o *NullOperation) LT(left Operand, right Operand) (bool, error) {
	return false, ErrInvalidOperation
}

func (o *NullOperation) GE(left Operand, right Operand) (bool, error) {
	return false, ErrInvalidOperation
}

func (o *NullOperation) LE(left Operand, right Operand) (bool, error) {
	return false, ErrInvalidOperation
}

func (o *NullOperation) CO(left Operand, right Operand) (bool, error) {
	return false, ErrInvalidOperation
}

func (o *NullOperation) SW(left Operand, right Operand) (bool, error) {
	return false, ErrInvalidOperation
}

func (o *NullOperation) EW(left Operand, right Operand) (bool, error) {
	return false, ErrInvalidOperation
}

func (o *NullOperation) IN(left Operand, right Operand) (bool, error) {
	return false, ErrInvalidOperation
}
