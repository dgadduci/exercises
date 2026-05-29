// Code generated from Test.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Test

import "github.com/antlr4-go/antlr/v4"

// TestListener is a complete listener for a parse tree produced by TestParser.
type TestListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterStat is called when entering the stat production.
	EnterStat(c *StatContext)

	// EnterAssignstat is called when entering the assignstat production.
	EnterAssignstat(c *AssignstatContext)

	// EnterValueExpr is called when entering the ValueExpr production.
	EnterValueExpr(c *ValueExprContext)

	// EnterAddMinus is called when entering the AddMinus production.
	EnterAddMinus(c *AddMinusContext)

	// EnterMulDiv is called when entering the MulDiv production.
	EnterMulDiv(c *MulDivContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitStat is called when exiting the stat production.
	ExitStat(c *StatContext)

	// ExitAssignstat is called when exiting the assignstat production.
	ExitAssignstat(c *AssignstatContext)

	// ExitValueExpr is called when exiting the ValueExpr production.
	ExitValueExpr(c *ValueExprContext)

	// ExitAddMinus is called when exiting the AddMinus production.
	ExitAddMinus(c *AddMinusContext)

	// ExitMulDiv is called when exiting the MulDiv production.
	ExitMulDiv(c *MulDivContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)
}
