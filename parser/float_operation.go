package parser

import (
	"reflect"
)

type FloatOperation struct {
	NullOperation
}

func (o *FloatOperation) EQ(left Operand, right Operand) (bool, error) {
	return applyWithTransform[float64, float64](left, right, toFloat, toFloat, func(l float64, r float64) bool { return l == r })
}

func (o *FloatOperation) NE(left Operand, right Operand) (bool, error) {
	return applyWithTransform[float64, float64](left, right, toFloat, toFloat, func(l float64, r float64) bool { return l != r })
}

func (o *FloatOperation) GT(left Operand, right Operand) (bool, error) {
	return applyWithTransform[float64, float64](left, right, toFloat, toFloat, func(l float64, r float64) bool { return l > r })
}

func (o *FloatOperation) LT(left Operand, right Operand) (bool, error) {
	return applyWithTransform[float64, float64](left, right, toFloat, toFloat, func(l float64, r float64) bool { return l < r })
}

func (o *FloatOperation) GE(left Operand, right Operand) (bool, error) {
	return applyWithTransform[float64, float64](left, right, toFloat, toFloat, func(l float64, r float64) bool { return l >= r })
}

func (o *FloatOperation) LE(left Operand, right Operand) (bool, error) {
	return applyWithTransform[float64, float64](left, right, toFloat, toFloat, func(l float64, r float64) bool { return l <= r })
}

func (o *FloatOperation) IN(left Operand, right Operand) (bool, error) {
	return applyWithTransform[float64, []float64](left, right, toFloat, func(right Operand) ([]float64, error) {
		rv := reflect.ValueOf(right)
		if rv.Kind() == reflect.Slice {
			ret := make([]float64, rv.Len())
			for i := 0; i < rv.Len(); i = i + 1 {
				val, err := toFloat(Operand(rv.Index(i).Interface()))
				if err != nil {
					return nil, err
				}
				ret[i] = val
			}
			return ret, nil
		} else {
			return nil, newErrInvalidOperand(right, []float64{})
		}
	}, func(l float64, r []float64) bool {
		for i := range r {
			if l == r[i] {
				return true
			}
		}
		return false
	})
}
