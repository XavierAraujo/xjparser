// Code generated from JsonGrammar.g4 by ANTLR 4.13.2. DO NOT EDIT.

package antlr4_parser // JsonGrammar
import "github.com/antlr4-go/antlr/v4"

// BaseJsonGrammarListener is a complete listener for a parse tree produced by JsonGrammarParser.
type BaseJsonGrammarListener struct{}

var _ JsonGrammarListener = &BaseJsonGrammarListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseJsonGrammarListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseJsonGrammarListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseJsonGrammarListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseJsonGrammarListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BaseJsonGrammarListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BaseJsonGrammarListener) ExitProg(ctx *ProgContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseJsonGrammarListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseJsonGrammarListener) ExitExpr(ctx *ExprContext) {}

// EnterJsonObjectExpr is called when production JsonObjectExpr is entered.
func (s *BaseJsonGrammarListener) EnterJsonObjectExpr(ctx *JsonObjectExprContext) {}

// ExitJsonObjectExpr is called when production JsonObjectExpr is exited.
func (s *BaseJsonGrammarListener) ExitJsonObjectExpr(ctx *JsonObjectExprContext) {}

// EnterJsonArrayExpr is called when production JsonArrayExpr is entered.
func (s *BaseJsonGrammarListener) EnterJsonArrayExpr(ctx *JsonArrayExprContext) {}

// ExitJsonArrayExpr is called when production JsonArrayExpr is exited.
func (s *BaseJsonGrammarListener) ExitJsonArrayExpr(ctx *JsonArrayExprContext) {}

// EnterJson_key_value is called when production json_key_value is entered.
func (s *BaseJsonGrammarListener) EnterJson_key_value(ctx *Json_key_valueContext) {}

// ExitJson_key_value is called when production json_key_value is exited.
func (s *BaseJsonGrammarListener) ExitJson_key_value(ctx *Json_key_valueContext) {}

// EnterJson_object is called when production json_object is entered.
func (s *BaseJsonGrammarListener) EnterJson_object(ctx *Json_objectContext) {}

// ExitJson_object is called when production json_object is exited.
func (s *BaseJsonGrammarListener) ExitJson_object(ctx *Json_objectContext) {}

// EnterJson_array is called when production json_array is entered.
func (s *BaseJsonGrammarListener) EnterJson_array(ctx *Json_arrayContext) {}

// ExitJson_array is called when production json_array is exited.
func (s *BaseJsonGrammarListener) ExitJson_array(ctx *Json_arrayContext) {}

// EnterIntValue is called when production IntValue is entered.
func (s *BaseJsonGrammarListener) EnterIntValue(ctx *IntValueContext) {}

// ExitIntValue is called when production IntValue is exited.
func (s *BaseJsonGrammarListener) ExitIntValue(ctx *IntValueContext) {}

// EnterFloatValue is called when production FloatValue is entered.
func (s *BaseJsonGrammarListener) EnterFloatValue(ctx *FloatValueContext) {}

// ExitFloatValue is called when production FloatValue is exited.
func (s *BaseJsonGrammarListener) ExitFloatValue(ctx *FloatValueContext) {}

// EnterStringValue is called when production StringValue is entered.
func (s *BaseJsonGrammarListener) EnterStringValue(ctx *StringValueContext) {}

// ExitStringValue is called when production StringValue is exited.
func (s *BaseJsonGrammarListener) ExitStringValue(ctx *StringValueContext) {}

// EnterBoolValue is called when production BoolValue is entered.
func (s *BaseJsonGrammarListener) EnterBoolValue(ctx *BoolValueContext) {}

// ExitBoolValue is called when production BoolValue is exited.
func (s *BaseJsonGrammarListener) ExitBoolValue(ctx *BoolValueContext) {}

// EnterNullValue is called when production NullValue is entered.
func (s *BaseJsonGrammarListener) EnterNullValue(ctx *NullValueContext) {}

// ExitNullValue is called when production NullValue is exited.
func (s *BaseJsonGrammarListener) ExitNullValue(ctx *NullValueContext) {}

// EnterObjectValue is called when production ObjectValue is entered.
func (s *BaseJsonGrammarListener) EnterObjectValue(ctx *ObjectValueContext) {}

// ExitObjectValue is called when production ObjectValue is exited.
func (s *BaseJsonGrammarListener) ExitObjectValue(ctx *ObjectValueContext) {}

// EnterArrayValue is called when production ArrayValue is entered.
func (s *BaseJsonGrammarListener) EnterArrayValue(ctx *ArrayValueContext) {}

// ExitArrayValue is called when production ArrayValue is exited.
func (s *BaseJsonGrammarListener) ExitArrayValue(ctx *ArrayValueContext) {}
