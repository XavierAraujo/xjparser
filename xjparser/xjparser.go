package xjparser

import (
	"strconv"
	"xjparser/xjparser/antlr4_parser"
	"xjparser/xjparser/util"

	"github.com/antlr4-go/antlr/v4"
)

const _MAX_ALLOWED_NESTED_LEVEL = 19

type BasicJsonGrammarListener struct {
	*antlr4_parser.BaseJsonGrammarListener
	topLevelElement JsonValue
	entityStack     *util.Stack[JsonValue]
}

func Parse(inputStr string) (jsonValue JsonValue, err error) {
	defer func() {
		if r := recover(); r != nil {
			jsonValue = nil
			err = r.(error)
		}
	}()

	errorListener := JsonErrorListener{}

	input := antlr.NewInputStream(inputStr)
	lexer := antlr4_parser.NewJsonGrammarLexer(input)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(&errorListener)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := antlr4_parser.NewJsonGrammarParser(stream)
	parser.RemoveErrorListeners()
	parser.AddErrorListener(&errorListener)

	listener := &BasicJsonGrammarListener{entityStack: util.NewStack[JsonValue]()}
	antlr.ParseTreeWalkerDefault.Walk(listener, parser.Expr())
	return listener.topLevelElement, nil
}

func (listener *BasicJsonGrammarListener) EnterJsonObjectExpr(ctx *antlr4_parser.JsonObjectExprContext) {
	// top level object
	topLevelObject := NewJsonObject()
	listener.pushEntityToStack(topLevelObject)
	listener.topLevelElement = topLevelObject
}

func (listener *BasicJsonGrammarListener) EnterJsonArrayExpr(ctx *antlr4_parser.JsonArrayExprContext) {
	// top level array
	topLevelArray := NewJsonArray()
	listener.pushEntityToStack(topLevelArray)
	listener.topLevelElement = topLevelArray
}

func (listener *BasicJsonGrammarListener) EnterJson_key_value(ctx *antlr4_parser.Json_key_valueContext) {
	// Add newly found key-value pair to the current entity
	currentStackTop, err := listener.entityStack.Peek()
	if err != nil {
		// Should never happen due to the grammar constraints
		panic(JsonParsingError{"Visiting stack is empty"})
	}

	// key-value pairs can only be found inside objects as defined in the grammar
	// so this cast is already pre-validated by the grammar itself
	currentObject := currentStackTop.(*JsonObject)
	key := getStringTokenWithoutQuotes(ctx.GetKey().GetText())
	if currentObject.HasKey(key) {
		panic(JsonDuplicatedKeyError{key})
	}

	switch ctx.GetValue().(type) {
	case *antlr4_parser.IntValueContext:
		intValue, _ := strconv.ParseInt(ctx.GetValue().GetText(), 10, 64)
		// TODO: Error validation for bit size. Validation for numeric value done by the grammar already
		currentObject.SetKey(key, NewJsonInt(intValue))
	case *antlr4_parser.FloatValueContext:
		floatValue, _ := strconv.ParseFloat(ctx.GetValue().GetText(), 64)
		// TODO: Error validation for bit size. Validation for numeric value done by the grammar already
		currentObject.SetKey(key, NewJsonFloat(floatValue))
	case *antlr4_parser.StringValueContext:
		strValue := getStringTokenWithoutQuotes(ctx.GetValue().GetText())
		currentObject.SetKey(key, NewJsonStr(strValue))
	case *antlr4_parser.BoolValueContext:
		boolValue := boolStr2Bool(ctx.GetValue().GetText())
		currentObject.SetKey(key, NewJsonBool(boolValue))
	case *antlr4_parser.NullValueContext:
		currentObject.SetKey(key, NewJsonNull())
	case *antlr4_parser.ObjectValueContext:
		// If the new value is an object we add it to the top of the stack since the next
		// key-value pairs found in the AST may belong to it
		newObject := NewJsonObject()
		currentObject.SetKey(key, newObject)
		listener.pushEntityToStack(newObject)
	case *antlr4_parser.ArrayValueContext:
		// If the new value is an array we add it to the top of the stack since the next
		// objects found in the AST may belong to it
		newArray := NewJsonArray()
		currentObject.SetKey(key, newArray)
		listener.pushEntityToStack(newArray)
	default:
		panic(JsonParsingError{"Unexpected JSON type"})
	}
}

func (listener *BasicJsonGrammarListener) EnterIntValue(ctx *antlr4_parser.IntValueContext) {
	_, isParentKeyValueParserRule := ctx.GetParent().(*antlr4_parser.Json_key_valueContext)
	if isParentKeyValueParserRule {
		return // Ignore. Value already processed on entering the key-value parser rule
	}

	currentStackTop, err := listener.entityStack.Peek()
	if err != nil {
		panic(JsonParsingError{"Visiting stack is empty"}) // Should never happen due to the grammar constraints
	}
	jsonArray := currentStackTop.(*JsonArray)
	intValue, _ := strconv.ParseInt(ctx.GetText(), 10, 64)
	jsonArray.Add(NewJsonInt(intValue))
}

func (listener *BasicJsonGrammarListener) EnterFloatValue(ctx *antlr4_parser.FloatValueContext) {
	_, isParentKeyValueParserRule := ctx.GetParent().(*antlr4_parser.Json_key_valueContext)
	if isParentKeyValueParserRule {
		return // Ignore. Value already processed on entering the key-value parser rule
	}

	currentStackTop, err := listener.entityStack.Peek()
	if err != nil {
		panic(JsonParsingError{"Visiting stack is empty"}) // Should never happen due to the grammar constraints
	}
	jsonArray := currentStackTop.(*JsonArray)
	floatValue, _ := strconv.ParseFloat(ctx.GetText(), 64)
	jsonArray.Add(NewJsonFloat(floatValue))
}

func (listener *BasicJsonGrammarListener) EnterStringValue(ctx *antlr4_parser.StringValueContext) {
	_, isParentKeyValueParserRule := ctx.GetParent().(*antlr4_parser.Json_key_valueContext)
	if isParentKeyValueParserRule {
		return // Ignore. Value already processed on entering the key-value parser rule
	}

	currentStackTop, err := listener.entityStack.Peek()
	if err != nil {
		panic(JsonParsingError{"Visiting stack is empty"}) // Should never happen due to the grammar constraints
	}
	jsonArray := currentStackTop.(*JsonArray)
	strValue := getStringTokenWithoutQuotes(ctx.GetText())
	jsonArray.Add(NewJsonStr(strValue))
}

func (listener *BasicJsonGrammarListener) EnterBoolValue(ctx *antlr4_parser.BoolValueContext) {
	_, isParentKeyValueParserRule := ctx.GetParent().(*antlr4_parser.Json_key_valueContext)
	if isParentKeyValueParserRule {
		return // Ignore. Value already processed on entering the key-value parser rule
	}

	currentStackTop, err := listener.entityStack.Peek()
	if err != nil {
		panic(JsonParsingError{"Visiting stack is empty"}) // Should never happen due to the grammar constraints
	}
	jsonArray := currentStackTop.(*JsonArray)
	boolValue := boolStr2Bool(ctx.GetText())
	jsonArray.Add(NewJsonBool(boolValue))
}

func (listener *BasicJsonGrammarListener) EnterNullValue(ctx *antlr4_parser.NullValueContext) {
	_, isParentKeyValueParserRule := ctx.GetParent().(*antlr4_parser.Json_key_valueContext)
	if isParentKeyValueParserRule {
		return // Ignore. Value already processed on entering the key-value parser rule
	}

	currentStackTop, err := listener.entityStack.Peek()
	if err != nil {
		panic(JsonParsingError{"Visiting stack is empty"}) // Should never happen due to the grammar constraints
	}
	jsonArray := currentStackTop.(*JsonArray)
	jsonArray.Add(NewJsonNull())
}

func (listener *BasicJsonGrammarListener) EnterObjectValue(ctx *antlr4_parser.ObjectValueContext) {
	_, isParentKeyValueParserRule := ctx.GetParent().(*antlr4_parser.Json_key_valueContext)
	if isParentKeyValueParserRule {
		return // Ignore. Value already processed on entering the key-value parser rule
	}

	currentStackTop, err := listener.entityStack.Peek()
	if err != nil {
		panic(JsonParsingError{"Visiting stack is empty"}) // Should never happen due to the grammar constraints
	}
	jsonArray := currentStackTop.(*JsonArray)
	newObject := NewJsonObject()
	jsonArray.Add(newObject)
	listener.pushEntityToStack(newObject)
}

func (listener *BasicJsonGrammarListener) ExitJson_object(ctx *antlr4_parser.Json_objectContext) {
	listener.popEntityFromStack()
}

func (listener *BasicJsonGrammarListener) EnterArrayValue(ctx *antlr4_parser.ArrayValueContext) {
	_, isParentKeyValueParserRule := ctx.GetParent().(*antlr4_parser.Json_key_valueContext)
	if isParentKeyValueParserRule {
		return // Ignore. Value already processed on entering the key-value parser rule
	}

	currentStackTop, err := listener.entityStack.Peek()
	if err != nil {
		panic(JsonParsingError{"Visiting stack is empty"}) // Should never happen due to the grammar constraints
	}
	jsonArray := currentStackTop.(*JsonArray)
	newArray := NewJsonArray()
	jsonArray.Add(newArray)
	listener.pushEntityToStack(newArray)
}

func (listener *BasicJsonGrammarListener) ExitJson_array(ctx *antlr4_parser.Json_arrayContext) {
	listener.popEntityFromStack()
}

func boolStr2Bool(boolStr string) bool {
	boolValue := true
	if boolStr == "false" {
		boolValue = false
	}
	return boolValue
}

func (listener *BasicJsonGrammarListener) pushEntityToStack(jsonValue JsonValue) {
	listener.entityStack.Push(jsonValue)
	if listener.entityStack.Len() > _MAX_ALLOWED_NESTED_LEVEL {
		panic(JsonParsingError{"Reached JSON max nesting level"})
	}
}

func (listener *BasicJsonGrammarListener) popEntityFromStack() {
	listener.entityStack.Pop()
}

func getStringTokenWithoutQuotes(stringToken string) string {
	// This is guaranteed to be safe due to the way we define the string token in the grammar file
	return stringToken[1 : len(stringToken)-1]
}
