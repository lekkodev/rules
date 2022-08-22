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
