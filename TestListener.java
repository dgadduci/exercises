// Generated from Test.g4 by ANTLR 4.13.2
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link TestParser}.
 */
public interface TestListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link TestParser#program}.
	 * @param ctx the parse tree
	 */
	void enterProgram(TestParser.ProgramContext ctx);
	/**
	 * Exit a parse tree produced by {@link TestParser#program}.
	 * @param ctx the parse tree
	 */
	void exitProgram(TestParser.ProgramContext ctx);
	/**
	 * Enter a parse tree produced by {@link TestParser#stat}.
	 * @param ctx the parse tree
	 */
	void enterStat(TestParser.StatContext ctx);
	/**
	 * Exit a parse tree produced by {@link TestParser#stat}.
	 * @param ctx the parse tree
	 */
	void exitStat(TestParser.StatContext ctx);
	/**
	 * Enter a parse tree produced by {@link TestParser#assignstat}.
	 * @param ctx the parse tree
	 */
	void enterAssignstat(TestParser.AssignstatContext ctx);
	/**
	 * Exit a parse tree produced by {@link TestParser#assignstat}.
	 * @param ctx the parse tree
	 */
	void exitAssignstat(TestParser.AssignstatContext ctx);
	/**
	 * Enter a parse tree produced by {@link TestParser#operation}.
	 * @param ctx the parse tree
	 */
	void enterOperation(TestParser.OperationContext ctx);
	/**
	 * Exit a parse tree produced by {@link TestParser#operation}.
	 * @param ctx the parse tree
	 */
	void exitOperation(TestParser.OperationContext ctx);
	/**
	 * Enter a parse tree produced by {@link TestParser#value}.
	 * @param ctx the parse tree
	 */
	void enterValue(TestParser.ValueContext ctx);
	/**
	 * Exit a parse tree produced by {@link TestParser#value}.
	 * @param ctx the parse tree
	 */
	void exitValue(TestParser.ValueContext ctx);
	/**
	 * Enter a parse tree produced by {@link TestParser#number}.
	 * @param ctx the parse tree
	 */
	void enterNumber(TestParser.NumberContext ctx);
	/**
	 * Exit a parse tree produced by {@link TestParser#number}.
	 * @param ctx the parse tree
	 */
	void exitNumber(TestParser.NumberContext ctx);
}