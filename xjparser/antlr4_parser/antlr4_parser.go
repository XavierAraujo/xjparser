package antlr4_parser

import (
	"log"
	"strconv"
	"xjparser/xjparser"

	"github.com/antlr4-go/antlr/v4"
	"github.com/golang-collections/collections/stack"
)

func Parse(inputStr string) (jsonType *xjparser.JsonValue, err error) {
	defer func() {
		if r := recover(); r != nil {
			jsonType = nil
			err = r.(error)
		}
	}()

	errorListener := JsonErrorListener{}

	input := antlr.NewInputStream(inputStr)
	lexer := NewJsonGrammarLexer(input)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(&errorListener)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := NewJsonGrammarParser(stream)
	parser.RemoveErrorListeners()
	parser.AddErrorListener(&errorListener)

	listener := &BasicJsonGrammarListener{}
	antlr.ParseTreeWalkerDefault.Walk(listener, parser.Expr())
	return &listener.topLevelElement, nil
}

type BasicJsonGrammarListener struct {
	*BaseJsonGrammarListener
	topLevelElement    xjparser.JsonValue
	currentEntityStack stack.Stack
}

func (listener *BasicJsonGrammarListener) EnterJsonObjectExpr(ctx *JsonObjectExprContext) {
	// top level object
	topLevelObject := xjparser.NewJsonObject()
	listener.currentEntityStack.Push(topLevelObject)
	listener.topLevelElement = topLevelObject
}

func (listener *BasicJsonGrammarListener) EnterJsonArrayExpr(ctx *JsonArrayExprContext) {
	// top level array
	topLevelArray := xjparser.NewJsonArray()
	listener.currentEntityStack.Push(topLevelArray)
	listener.topLevelElement = topLevelArray
}

func (listener *BasicJsonGrammarListener) EnterJson_object(ctx *Json_objectContext) {
	jsonArray, ok := listener.currentEntityStack.Peek().(*xjparser.JsonArray)
	if !ok {
		// Ignore. If the top of the stack is not an array this object was already
		// initialized when EnterJson_key_value() was called.
		return
	}

	// If we are inside an array and a new object is found we need to initialize it
	// and put on the top of visiting stack.
	newObject := xjparser.NewJsonObject()
	listener.currentEntityStack.Push(newObject)
	jsonArray.Add(*newObject)
}

func (listener *BasicJsonGrammarListener) ExitJson_object(ctx *Json_objectContext) {
	listener.currentEntityStack.Pop()
}

func (listener *BasicJsonGrammarListener) ExitJson_array(ctx *Json_arrayContext) {
	listener.currentEntityStack.Pop()
}

func (listener *BasicJsonGrammarListener) EnterJson_key_value(ctx *Json_key_valueContext) {
	// Add newly found key-value pair to the current entity
	currentObject := listener.currentEntityStack.Peek().(*xjparser.JsonObject)

	key := getStringTokenWithoutQuotes(ctx.key.GetText())
	if currentObject.HasKey(key) {
		panic(JsonDuplicatedKeyError{key})
	}

	switch ctx.value.(type) {
	case *IntValueContext:
		intValue, _ := strconv.ParseInt(ctx.value.GetText(), 10, 64)
		// TODO: Error validation for bit size. Validation for numeric value done by the grammar already
		currentObject.SetKey(key, xjparser.NewJsonInt(intValue))
	case *FloatValueContext:
		floatValue, _ := strconv.ParseFloat(ctx.value.GetText(), 64)
		// TODO: Error validation for bit size. Validation for numeric value done by the grammar already
		currentObject.SetKey(key, xjparser.NewJsonFloat(floatValue))
	case *StringValueContext:
		strValue := getStringTokenWithoutQuotes(ctx.value.GetText())
		currentObject.SetKey(key, xjparser.NewJsonStr(strValue))
	case *BoolValueContext:
		boolValue := true
		if ctx.value.GetText() == "false" {
			boolValue = false
		}
		currentObject.SetKey(key, xjparser.NewJsonBool(boolValue))
	case *NullValueContext:
		currentObject.SetKey(key, xjparser.NewJsonNull())
	case *ObjectValueContext:
		// If the new value is an object we add it to the top of the stack since the next
		// key-value pairs found in the AST may belong to it
		newObject := xjparser.NewJsonObject()
		currentObject.SetKey(key, newObject)
		listener.currentEntityStack.Push(newObject)
	case *ArrayValueContext:
		// If the new value is an array we add it to the top of the stack since the next
		// objects found in the AST may belong to it
		newArray := xjparser.NewJsonArray()
		currentObject.SetKey(key, newArray)
		listener.currentEntityStack.Push(newArray)
	default:
		log.Panic(JsonParsingError{"Unexpected JSON type"})
	}
}

func getStringTokenWithoutQuotes(stringToken string) string {
	// This is guaranteed to be safe due to the way we define the string token in the grammar file
	return stringToken[1 : len(stringToken)-1]
}
