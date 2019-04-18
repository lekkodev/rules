package parser

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type objStack struct {
	items []interface{}
}

func (o *objStack) empty() bool {
	return len(o.items) == 0
}

func (o *objStack) peek() interface{} {
	return o.items[len(o.items)-1]
}

func (o *objStack) pop() interface{} {
	val := o.peek()
	o.items = o.items[:len(o.items)-1]
	return val
}

func (o *objStack) push(item interface{}) {
	o.items = append(o.items, item)
}

func (o *objStack) clear() {
	o.items = nil
}

type JsonQueryVisitorImpl struct {
	antlr.ParseTreeVisitor

	item map[string]interface{}

	stack *objStack

	currentOperation Operation
	leftOp           Operand
	rightOp          Operand
}

func NewJsonQueryVisitorImpl(item map[string]interface{}) *JsonQueryVisitorImpl {
	return &JsonQueryVisitorImpl{
		stack: &objStack{},
		item:  item,
	}
}

func (j *JsonQueryVisitorImpl) Visit(tree antlr.ParseTree) interface{} {
	switch val := tree.(type) {
	case *LogicalExpContext:
		return val.Accept(j).(bool)
	case *CompareExpContext:
		return val.Accept(j).(bool)
	case *ParenExpContext:
		return val.Accept(j).(bool)
	case *PresentExpContext:
		return val.Accept(j).(bool)
	default:
		fmt.Println(reflect.TypeOf(tree))
	}
	return nil
}

func (j *JsonQueryVisitorImpl) VisitParenExp(ctx *ParenExpContext) interface{} {
	result := ctx.Query().Accept(j).(bool)
	if ctx.NOT() != nil {
		return !result
	}
	return result
}

func (j *JsonQueryVisitorImpl) VisitLogicalExp(ctx *LogicalExpContext) interface{} {
	left := ctx.Query(0).Accept(j).(bool)
	op := ctx.LOGICAL_OPERATOR().GetText()
	if op == "or" {
		if left {
			return left
		}
		return ctx.Query(1).Accept(j).(bool)
	}
	// means it is and
	if !left {
		return left
	}
	return ctx.Query(1).Accept(j).(bool)
}

func (j *JsonQueryVisitorImpl) VisitPresentExp(ctx *PresentExpContext) interface{} {
	ctx.AttrPath().Accept(j)
	return j.leftOp != nil
}

func (j *JsonQueryVisitorImpl) VisitCompareExp(ctx *CompareExpContext) interface{} {
	ctx.AttrPath().Accept(j)
	ctx.Value().Accept(j)
	var apply func(Operand, Operand) bool
	currentOp := j.currentOperation
	switch ctx.op.GetTokenType() {
	case JsonQueryParserEQ:
		apply = currentOp.EQ
	case JsonQueryParserNE:
		apply = currentOp.NE
	case JsonQueryParserGT:
		apply = currentOp.GT
	case JsonQueryParserLT:
		apply = currentOp.LT
	case JsonQueryParserLE:
		apply = currentOp.LE
	case JsonQueryParserGE:
		apply = currentOp.GE
	case JsonQueryParserCO:
		apply = currentOp.CO
	case JsonQueryParserSW:
		apply = currentOp.SW
	case JsonQueryParserEW:
		apply = currentOp.EW
	default:
		panic("unknown operation")
	}
	return apply(j.leftOp, j.rightOp)
}

func (j *JsonQueryVisitorImpl) VisitAttrPath(ctx *AttrPathContext) interface{} {
	var item interface{}
	if ctx.SubAttr() == nil || ctx.SubAttr().IsEmpty() {
		if !j.stack.empty() {
			item = j.stack.pop()
		} else {
			item = j.item
		}
		if item == nil {
			return false
		}
		m := item.(map[string]interface{})
		j.leftOp = m[ctx.ATTRNAME().GetText()]
		j.stack.clear()
		return true
	}
	if !j.stack.empty() {
		item = j.stack.peek()
	} else {
		item = j.item
	}
	if item == nil {
		return false
	}
	m := item.(map[string]interface{})
	j.stack.push(m[ctx.ATTRNAME().GetText()])
	return ctx.SubAttr().Accept(j).(bool)
}

func (j *JsonQueryVisitorImpl) VisitSubAttr(ctx *SubAttrContext) interface{} {
	// TODO fix this if it is okay
	return ctx.AttrPath().Accept(j).(bool)
}

func (j *JsonQueryVisitorImpl) VisitBoolean(ctx *BooleanContext) interface{} {
	j.currentOperation = &BoolOperation{}

	val, err := strconv.ParseBool(ctx.GetText())
	if err != nil {
		// TODO set err somewhere
		j.rightOp = nil
		return false
	}
	j.rightOp = val
	return true
}

func (j *JsonQueryVisitorImpl) VisitNull(ctx *NullContext) interface{} {
	j.currentOperation = &NullOperation{}
	j.rightOp = nil
	return true
}

func (j *JsonQueryVisitorImpl) VisitString(ctx *StringContext) interface{} {
	j.currentOperation = &StringOperation{}

	s := ctx.GetText()
	if len(s) > 2 {
		j.rightOp = s[1 : len(s)-1]
	} else {
		j.rightOp = ""
	}
	return true
}

func (j *JsonQueryVisitorImpl) VisitDouble(ctx *DoubleContext) interface{} {
	j.currentOperation = &FloatOperation{}
	val, err := strconv.ParseFloat(ctx.GetText(), 10)
	if err != nil {
		// TODO set err somewhere
		j.rightOp = nil
		return false
	}
	j.rightOp = val
	return true
}

func (j *JsonQueryVisitorImpl) VisitLong(ctx *LongContext) interface{} {
	j.currentOperation = &IntOperation{}
	val, err := strconv.ParseInt(ctx.GetText(), 10, 64)
	if err != nil {
		// TODO set err somewhere
		j.rightOp = nil
		return false
	}
	j.rightOp = int(val)
	return true
}