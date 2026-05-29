// Code generated from Test.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Test

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by TestParser.
type TestVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by TestParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by TestParser#stat.
	VisitStat(ctx *StatContext) interface{}

	// Visit a parse tree produced by TestParser#assignstat.
	VisitAssignstat(ctx *AssignstatContext) interface{}

	// Visit a parse tree produced by TestParser#ValueExpr.
	VisitValueExpr(ctx *ValueExprContext) interface{}

	// Visit a parse tree produced by TestParser#AddMinus.
	VisitAddMinus(ctx *AddMinusContext) interface{}

	// Visit a parse tree produced by TestParser#MulDiv.
	VisitMulDiv(ctx *MulDivContext) interface{}

	// Visit a parse tree produced by TestParser#value.
	VisitValue(ctx *ValueContext) interface{}

	// Visit a parse tree produced by TestParser#number.
	VisitNumber(ctx *NumberContext) interface{}
}
