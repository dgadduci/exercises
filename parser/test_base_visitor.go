// Code generated from Test.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Test

import "github.com/antlr4-go/antlr/v4"

type BaseTestVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseTestVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTestVisitor) VisitStat(ctx *StatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTestVisitor) VisitAssignstat(ctx *AssignstatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTestVisitor) VisitValueExpr(ctx *ValueExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTestVisitor) VisitAddMinus(ctx *AddMinusContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTestVisitor) VisitMulDiv(ctx *MulDivContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTestVisitor) VisitValue(ctx *ValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTestVisitor) VisitNumber(ctx *NumberContext) interface{} {
	return v.VisitChildren(ctx)
}
