// Code generated from JsonQuery.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // JsonQuery

import "github.com/antlr4-go/antlr/v4"

type BaseJsonQueryVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseJsonQueryVisitor) VisitQuery(ctx *QueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonQueryVisitor) VisitCompareExp(ctx *CompareExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonQueryVisitor) VisitAndLogicalExp(ctx *AndLogicalExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonQueryVisitor) VisitParenExp(ctx *ParenExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonQueryVisitor) VisitPresentExp(ctx *PresentExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonQueryVisitor) VisitCallExp(ctx *CallExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonQueryVisitor) VisitOrLogicalExp(ctx *OrLogicalExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonQueryVisitor) VisitAttrPath(ctx *AttrPathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonQueryVisitor) VisitSubAttr(ctx *SubAttrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonQueryVisitor) VisitBoolean(ctx *BooleanContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonQueryVisitor) VisitNull(ctx *NullContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonQueryVisitor) VisitVersion(ctx *VersionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonQueryVisitor) VisitString(ctx *StringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonQueryVisitor) VisitDouble(ctx *DoubleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonQueryVisitor) VisitLong(ctx *LongContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonQueryVisitor) VisitListOfNumbers(ctx *ListOfNumbersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonQueryVisitor) VisitListOfStrings(ctx *ListOfStringsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonQueryVisitor) VisitListOfBooleans(ctx *ListOfBooleansContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonQueryVisitor) VisitListStrings(ctx *ListStringsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonQueryVisitor) VisitSubListOfStrings(ctx *SubListOfStringsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonQueryVisitor) VisitListNumbers(ctx *ListNumbersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonQueryVisitor) VisitSubListOfNumbers(ctx *SubListOfNumbersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonQueryVisitor) VisitListBooleans(ctx *ListBooleansContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonQueryVisitor) VisitSubListOfBooleans(ctx *SubListOfBooleansContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseJsonQueryVisitor) VisitFunctionArg(ctx *FunctionArgContext) interface{} {
	return v.VisitChildren(ctx)
}
