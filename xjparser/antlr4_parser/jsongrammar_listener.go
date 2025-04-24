// Code generated from JsonGrammar.g4 by ANTLR 4.13.2. DO NOT EDIT.

package antlr4_parser // JsonGrammar
import "github.com/antlr4-go/antlr/v4"

// JsonGrammarListener is a complete listener for a parse tree produced by JsonGrammarParser.
type JsonGrammarListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterJsonObjectExpr is called when entering the JsonObjectExpr production.
	EnterJsonObjectExpr(c *JsonObjectExprContext)

	// EnterJsonArrayExpr is called when entering the JsonArrayExpr production.
	EnterJsonArrayExpr(c *JsonArrayExprContext)

	// EnterJson_key_value is called when entering the json_key_value production.
	EnterJson_key_value(c *Json_key_valueContext)

	// EnterJson_object is called when entering the json_object production.
	EnterJson_object(c *Json_objectContext)

	// EnterJson_array is called when entering the json_array production.
	EnterJson_array(c *Json_arrayContext)

	// EnterIntValue is called when entering the IntValue production.
	EnterIntValue(c *IntValueContext)

	// EnterFloatValue is called when entering the FloatValue production.
	EnterFloatValue(c *FloatValueContext)

	// EnterStringValue is called when entering the StringValue production.
	EnterStringValue(c *StringValueContext)

	// EnterBoolValue is called when entering the BoolValue production.
	EnterBoolValue(c *BoolValueContext)

	// EnterNullValue is called when entering the NullValue production.
	EnterNullValue(c *NullValueContext)

	// EnterObjectValue is called when entering the ObjectValue production.
	EnterObjectValue(c *ObjectValueContext)

	// EnterArrayValue is called when entering the ArrayValue production.
	EnterArrayValue(c *ArrayValueContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitJsonObjectExpr is called when exiting the JsonObjectExpr production.
	ExitJsonObjectExpr(c *JsonObjectExprContext)

	// ExitJsonArrayExpr is called when exiting the JsonArrayExpr production.
	ExitJsonArrayExpr(c *JsonArrayExprContext)

	// ExitJson_key_value is called when exiting the json_key_value production.
	ExitJson_key_value(c *Json_key_valueContext)

	// ExitJson_object is called when exiting the json_object production.
	ExitJson_object(c *Json_objectContext)

	// ExitJson_array is called when exiting the json_array production.
	ExitJson_array(c *Json_arrayContext)

	// ExitIntValue is called when exiting the IntValue production.
	ExitIntValue(c *IntValueContext)

	// ExitFloatValue is called when exiting the FloatValue production.
	ExitFloatValue(c *FloatValueContext)

	// ExitStringValue is called when exiting the StringValue production.
	ExitStringValue(c *StringValueContext)

	// ExitBoolValue is called when exiting the BoolValue production.
	ExitBoolValue(c *BoolValueContext)

	// ExitNullValue is called when exiting the NullValue production.
	ExitNullValue(c *NullValueContext)

	// ExitObjectValue is called when exiting the ObjectValue production.
	ExitObjectValue(c *ObjectValueContext)

	// ExitArrayValue is called when exiting the ArrayValue production.
	ExitArrayValue(c *ArrayValueContext)
}
