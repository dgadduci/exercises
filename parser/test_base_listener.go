// Code generated from Test.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Test

import "github.com/antlr4-go/antlr/v4"

// BaseTestListener is a complete listener for a parse tree produced by TestParser.
type BaseTestListener struct{}

var _ TestListener = &BaseTestListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseTestListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseTestListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseTestListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseTestListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseTestListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseTestListener) ExitProgram(ctx *ProgramContext) {}

// EnterStat is called when production stat is entered.
func (s *BaseTestListener) EnterStat(ctx *StatContext) {}

// ExitStat is called when production stat is exited.
func (s *BaseTestListener) ExitStat(ctx *StatContext) {}

// EnterAssignstat is called when production assignstat is entered.
func (s *BaseTestListener) EnterAssignstat(ctx *AssignstatContext) {}

// ExitAssignstat is called when production assignstat is exited.
func (s *BaseTestListener) ExitAssignstat(ctx *AssignstatContext) {}

// EnterValueExpr is called when production ValueExpr is entered.
func (s *BaseTestListener) EnterValueExpr(ctx *ValueExprContext) {}

// ExitValueExpr is called when production ValueExpr is exited.
func (s *BaseTestListener) ExitValueExpr(ctx *ValueExprContext) {}

// EnterAddMinus is called when production AddMinus is entered.
func (s *BaseTestListener) EnterAddMinus(ctx *AddMinusContext) {}

// ExitAddMinus is called when production AddMinus is exited.
func (s *BaseTestListener) ExitAddMinus(ctx *AddMinusContext) {}

// EnterMulDiv is called when production MulDiv is entered.
func (s *BaseTestListener) EnterMulDiv(ctx *MulDivContext) {}

// ExitMulDiv is called when production MulDiv is exited.
func (s *BaseTestListener) ExitMulDiv(ctx *MulDivContext) {}

// EnterValue is called when production value is entered.
func (s *BaseTestListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseTestListener) ExitValue(ctx *ValueContext) {}

// EnterNumber is called when production number is entered.
func (s *BaseTestListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BaseTestListener) ExitNumber(ctx *NumberContext) {}
