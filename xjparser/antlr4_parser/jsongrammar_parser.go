// Code generated from JsonGrammar.g4 by ANTLR 4.13.2. DO NOT EDIT.

package antlr4_parser // JsonGrammar
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type JsonGrammarParser struct {
	*antlr.BaseParser
}

var JsonGrammarParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func jsongrammarParserInit() {
	staticData := &JsonGrammarParserStaticData
	staticData.LiteralNames = []string{
		"", "':'", "'{'", "','", "'}'", "'['", "']'", "", "", "", "", "'null'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "INT", "FLOAT", "STRING", "BOOL", "NULL",
		"WS",
	}
	staticData.RuleNames = []string{
		"prog", "expr", "json_expr", "json_key_value", "json_object", "json_array",
		"json_value",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 12, 63, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 3,
		2, 22, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4, 32,
		8, 4, 10, 4, 12, 4, 35, 9, 4, 3, 4, 37, 8, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1,
		5, 1, 5, 5, 5, 45, 8, 5, 10, 5, 12, 5, 48, 9, 5, 3, 5, 50, 8, 5, 1, 5,
		1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 61, 8, 6, 1, 6, 0,
		0, 7, 0, 2, 4, 6, 8, 10, 12, 0, 0, 66, 0, 14, 1, 0, 0, 0, 2, 16, 1, 0,
		0, 0, 4, 21, 1, 0, 0, 0, 6, 23, 1, 0, 0, 0, 8, 27, 1, 0, 0, 0, 10, 40,
		1, 0, 0, 0, 12, 60, 1, 0, 0, 0, 14, 15, 3, 2, 1, 0, 15, 1, 1, 0, 0, 0,
		16, 17, 3, 4, 2, 0, 17, 18, 5, 0, 0, 1, 18, 3, 1, 0, 0, 0, 19, 22, 3, 8,
		4, 0, 20, 22, 3, 10, 5, 0, 21, 19, 1, 0, 0, 0, 21, 20, 1, 0, 0, 0, 22,
		5, 1, 0, 0, 0, 23, 24, 5, 9, 0, 0, 24, 25, 5, 1, 0, 0, 25, 26, 3, 12, 6,
		0, 26, 7, 1, 0, 0, 0, 27, 36, 5, 2, 0, 0, 28, 33, 3, 6, 3, 0, 29, 30, 5,
		3, 0, 0, 30, 32, 3, 6, 3, 0, 31, 29, 1, 0, 0, 0, 32, 35, 1, 0, 0, 0, 33,
		31, 1, 0, 0, 0, 33, 34, 1, 0, 0, 0, 34, 37, 1, 0, 0, 0, 35, 33, 1, 0, 0,
		0, 36, 28, 1, 0, 0, 0, 36, 37, 1, 0, 0, 0, 37, 38, 1, 0, 0, 0, 38, 39,
		5, 4, 0, 0, 39, 9, 1, 0, 0, 0, 40, 49, 5, 5, 0, 0, 41, 46, 3, 12, 6, 0,
		42, 43, 5, 3, 0, 0, 43, 45, 3, 12, 6, 0, 44, 42, 1, 0, 0, 0, 45, 48, 1,
		0, 0, 0, 46, 44, 1, 0, 0, 0, 46, 47, 1, 0, 0, 0, 47, 50, 1, 0, 0, 0, 48,
		46, 1, 0, 0, 0, 49, 41, 1, 0, 0, 0, 49, 50, 1, 0, 0, 0, 50, 51, 1, 0, 0,
		0, 51, 52, 5, 6, 0, 0, 52, 11, 1, 0, 0, 0, 53, 61, 5, 7, 0, 0, 54, 61,
		5, 8, 0, 0, 55, 61, 5, 9, 0, 0, 56, 61, 5, 10, 0, 0, 57, 61, 5, 11, 0,
		0, 58, 61, 3, 8, 4, 0, 59, 61, 3, 10, 5, 0, 60, 53, 1, 0, 0, 0, 60, 54,
		1, 0, 0, 0, 60, 55, 1, 0, 0, 0, 60, 56, 1, 0, 0, 0, 60, 57, 1, 0, 0, 0,
		60, 58, 1, 0, 0, 0, 60, 59, 1, 0, 0, 0, 61, 13, 1, 0, 0, 0, 6, 21, 33,
		36, 46, 49, 60,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// JsonGrammarParserInit initializes any static state used to implement JsonGrammarParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewJsonGrammarParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func JsonGrammarParserInit() {
	staticData := &JsonGrammarParserStaticData
	staticData.once.Do(jsongrammarParserInit)
}

// NewJsonGrammarParser produces a new parser instance for the optional input antlr.TokenStream.
func NewJsonGrammarParser(input antlr.TokenStream) *JsonGrammarParser {
	JsonGrammarParserInit()
	this := new(JsonGrammarParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &JsonGrammarParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "JsonGrammar.g4"

	return this
}

// JsonGrammarParser tokens.
const (
	JsonGrammarParserEOF    = antlr.TokenEOF
	JsonGrammarParserT__0   = 1
	JsonGrammarParserT__1   = 2
	JsonGrammarParserT__2   = 3
	JsonGrammarParserT__3   = 4
	JsonGrammarParserT__4   = 5
	JsonGrammarParserT__5   = 6
	JsonGrammarParserINT    = 7
	JsonGrammarParserFLOAT  = 8
	JsonGrammarParserSTRING = 9
	JsonGrammarParserBOOL   = 10
	JsonGrammarParserNULL   = 11
	JsonGrammarParserWS     = 12
)

// JsonGrammarParser rules.
const (
	JsonGrammarParserRULE_prog           = 0
	JsonGrammarParserRULE_expr           = 1
	JsonGrammarParserRULE_json_expr      = 2
	JsonGrammarParserRULE_json_key_value = 3
	JsonGrammarParserRULE_json_object    = 4
	JsonGrammarParserRULE_json_array     = 5
	JsonGrammarParserRULE_json_value     = 6
)

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext

	// IsProgContext differentiates from other interfaces.
	IsProgContext()
}

type ProgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgContext() *ProgContext {
	var p = new(ProgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonGrammarParserRULE_prog
	return p
}

func InitEmptyProgContext(p *ProgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonGrammarParserRULE_prog
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonGrammarParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonGrammarListener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonGrammarListener); ok {
		listenerT.ExitProg(s)
	}
}

func (p *JsonGrammarParser) Prog() (localctx IProgContext) {
	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, JsonGrammarParserRULE_prog)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(14)
		p.Expr()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Json_expr() IJson_exprContext
	EOF() antlr.TerminalNode

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonGrammarParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonGrammarParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonGrammarParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) Json_expr() IJson_exprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJson_exprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJson_exprContext)
}

func (s *ExprContext) EOF() antlr.TerminalNode {
	return s.GetToken(JsonGrammarParserEOF, 0)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonGrammarListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonGrammarListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (p *JsonGrammarParser) Expr() (localctx IExprContext) {
	localctx = NewExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, JsonGrammarParserRULE_expr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(16)
		p.Json_expr()
	}
	{
		p.SetState(17)
		p.Match(JsonGrammarParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IJson_exprContext is an interface to support dynamic dispatch.
type IJson_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsJson_exprContext differentiates from other interfaces.
	IsJson_exprContext()
}

type Json_exprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJson_exprContext() *Json_exprContext {
	var p = new(Json_exprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonGrammarParserRULE_json_expr
	return p
}

func InitEmptyJson_exprContext(p *Json_exprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonGrammarParserRULE_json_expr
}

func (*Json_exprContext) IsJson_exprContext() {}

func NewJson_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Json_exprContext {
	var p = new(Json_exprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonGrammarParserRULE_json_expr

	return p
}

func (s *Json_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Json_exprContext) CopyAll(ctx *Json_exprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Json_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Json_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type JsonArrayExprContext struct {
	Json_exprContext
}

func NewJsonArrayExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *JsonArrayExprContext {
	var p = new(JsonArrayExprContext)

	InitEmptyJson_exprContext(&p.Json_exprContext)
	p.parser = parser
	p.CopyAll(ctx.(*Json_exprContext))

	return p
}

func (s *JsonArrayExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JsonArrayExprContext) Json_array() IJson_arrayContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJson_arrayContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJson_arrayContext)
}

func (s *JsonArrayExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonGrammarListener); ok {
		listenerT.EnterJsonArrayExpr(s)
	}
}

func (s *JsonArrayExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonGrammarListener); ok {
		listenerT.ExitJsonArrayExpr(s)
	}
}

type JsonObjectExprContext struct {
	Json_exprContext
}

func NewJsonObjectExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *JsonObjectExprContext {
	var p = new(JsonObjectExprContext)

	InitEmptyJson_exprContext(&p.Json_exprContext)
	p.parser = parser
	p.CopyAll(ctx.(*Json_exprContext))

	return p
}

func (s *JsonObjectExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JsonObjectExprContext) Json_object() IJson_objectContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJson_objectContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJson_objectContext)
}

func (s *JsonObjectExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonGrammarListener); ok {
		listenerT.EnterJsonObjectExpr(s)
	}
}

func (s *JsonObjectExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonGrammarListener); ok {
		listenerT.ExitJsonObjectExpr(s)
	}
}

func (p *JsonGrammarParser) Json_expr() (localctx IJson_exprContext) {
	localctx = NewJson_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, JsonGrammarParserRULE_json_expr)
	p.SetState(21)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case JsonGrammarParserT__1:
		localctx = NewJsonObjectExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(19)
			p.Json_object()
		}

	case JsonGrammarParserT__4:
		localctx = NewJsonArrayExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(20)
			p.Json_array()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IJson_key_valueContext is an interface to support dynamic dispatch.
type IJson_key_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetKey returns the key token.
	GetKey() antlr.Token

	// SetKey sets the key token.
	SetKey(antlr.Token)

	// GetValue returns the value rule contexts.
	GetValue() IJson_valueContext

	// SetValue sets the value rule contexts.
	SetValue(IJson_valueContext)

	// Getter signatures
	STRING() antlr.TerminalNode
	Json_value() IJson_valueContext

	// IsJson_key_valueContext differentiates from other interfaces.
	IsJson_key_valueContext()
}

type Json_key_valueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	key    antlr.Token
	value  IJson_valueContext
}

func NewEmptyJson_key_valueContext() *Json_key_valueContext {
	var p = new(Json_key_valueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonGrammarParserRULE_json_key_value
	return p
}

func InitEmptyJson_key_valueContext(p *Json_key_valueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonGrammarParserRULE_json_key_value
}

func (*Json_key_valueContext) IsJson_key_valueContext() {}

func NewJson_key_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Json_key_valueContext {
	var p = new(Json_key_valueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonGrammarParserRULE_json_key_value

	return p
}

func (s *Json_key_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Json_key_valueContext) GetKey() antlr.Token { return s.key }

func (s *Json_key_valueContext) SetKey(v antlr.Token) { s.key = v }

func (s *Json_key_valueContext) GetValue() IJson_valueContext { return s.value }

func (s *Json_key_valueContext) SetValue(v IJson_valueContext) { s.value = v }

func (s *Json_key_valueContext) STRING() antlr.TerminalNode {
	return s.GetToken(JsonGrammarParserSTRING, 0)
}

func (s *Json_key_valueContext) Json_value() IJson_valueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJson_valueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJson_valueContext)
}

func (s *Json_key_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Json_key_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Json_key_valueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonGrammarListener); ok {
		listenerT.EnterJson_key_value(s)
	}
}

func (s *Json_key_valueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonGrammarListener); ok {
		listenerT.ExitJson_key_value(s)
	}
}

func (p *JsonGrammarParser) Json_key_value() (localctx IJson_key_valueContext) {
	localctx = NewJson_key_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, JsonGrammarParserRULE_json_key_value)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(23)

		var _m = p.Match(JsonGrammarParserSTRING)

		localctx.(*Json_key_valueContext).key = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(24)
		p.Match(JsonGrammarParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(25)

		var _x = p.Json_value()

		localctx.(*Json_key_valueContext).value = _x
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IJson_objectContext is an interface to support dynamic dispatch.
type IJson_objectContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllJson_key_value() []IJson_key_valueContext
	Json_key_value(i int) IJson_key_valueContext

	// IsJson_objectContext differentiates from other interfaces.
	IsJson_objectContext()
}

type Json_objectContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJson_objectContext() *Json_objectContext {
	var p = new(Json_objectContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonGrammarParserRULE_json_object
	return p
}

func InitEmptyJson_objectContext(p *Json_objectContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonGrammarParserRULE_json_object
}

func (*Json_objectContext) IsJson_objectContext() {}

func NewJson_objectContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Json_objectContext {
	var p = new(Json_objectContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonGrammarParserRULE_json_object

	return p
}

func (s *Json_objectContext) GetParser() antlr.Parser { return s.parser }

func (s *Json_objectContext) AllJson_key_value() []IJson_key_valueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IJson_key_valueContext); ok {
			len++
		}
	}

	tst := make([]IJson_key_valueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IJson_key_valueContext); ok {
			tst[i] = t.(IJson_key_valueContext)
			i++
		}
	}

	return tst
}

func (s *Json_objectContext) Json_key_value(i int) IJson_key_valueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJson_key_valueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJson_key_valueContext)
}

func (s *Json_objectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Json_objectContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Json_objectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonGrammarListener); ok {
		listenerT.EnterJson_object(s)
	}
}

func (s *Json_objectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonGrammarListener); ok {
		listenerT.ExitJson_object(s)
	}
}

func (p *JsonGrammarParser) Json_object() (localctx IJson_objectContext) {
	localctx = NewJson_objectContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, JsonGrammarParserRULE_json_object)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(27)
		p.Match(JsonGrammarParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(36)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == JsonGrammarParserSTRING {
		{
			p.SetState(28)
			p.Json_key_value()
		}
		p.SetState(33)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == JsonGrammarParserT__2 {
			{
				p.SetState(29)
				p.Match(JsonGrammarParserT__2)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(30)
				p.Json_key_value()
			}

			p.SetState(35)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(38)
		p.Match(JsonGrammarParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IJson_arrayContext is an interface to support dynamic dispatch.
type IJson_arrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllJson_value() []IJson_valueContext
	Json_value(i int) IJson_valueContext

	// IsJson_arrayContext differentiates from other interfaces.
	IsJson_arrayContext()
}

type Json_arrayContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJson_arrayContext() *Json_arrayContext {
	var p = new(Json_arrayContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonGrammarParserRULE_json_array
	return p
}

func InitEmptyJson_arrayContext(p *Json_arrayContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonGrammarParserRULE_json_array
}

func (*Json_arrayContext) IsJson_arrayContext() {}

func NewJson_arrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Json_arrayContext {
	var p = new(Json_arrayContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonGrammarParserRULE_json_array

	return p
}

func (s *Json_arrayContext) GetParser() antlr.Parser { return s.parser }

func (s *Json_arrayContext) AllJson_value() []IJson_valueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IJson_valueContext); ok {
			len++
		}
	}

	tst := make([]IJson_valueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IJson_valueContext); ok {
			tst[i] = t.(IJson_valueContext)
			i++
		}
	}

	return tst
}

func (s *Json_arrayContext) Json_value(i int) IJson_valueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJson_valueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJson_valueContext)
}

func (s *Json_arrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Json_arrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Json_arrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonGrammarListener); ok {
		listenerT.EnterJson_array(s)
	}
}

func (s *Json_arrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonGrammarListener); ok {
		listenerT.ExitJson_array(s)
	}
}

func (p *JsonGrammarParser) Json_array() (localctx IJson_arrayContext) {
	localctx = NewJson_arrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, JsonGrammarParserRULE_json_array)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(40)
		p.Match(JsonGrammarParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(49)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4004) != 0 {
		{
			p.SetState(41)
			p.Json_value()
		}
		p.SetState(46)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == JsonGrammarParserT__2 {
			{
				p.SetState(42)
				p.Match(JsonGrammarParserT__2)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(43)
				p.Json_value()
			}

			p.SetState(48)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(51)
		p.Match(JsonGrammarParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IJson_valueContext is an interface to support dynamic dispatch.
type IJson_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsJson_valueContext differentiates from other interfaces.
	IsJson_valueContext()
}

type Json_valueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJson_valueContext() *Json_valueContext {
	var p = new(Json_valueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonGrammarParserRULE_json_value
	return p
}

func InitEmptyJson_valueContext(p *Json_valueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = JsonGrammarParserRULE_json_value
}

func (*Json_valueContext) IsJson_valueContext() {}

func NewJson_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Json_valueContext {
	var p = new(Json_valueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = JsonGrammarParserRULE_json_value

	return p
}

func (s *Json_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Json_valueContext) CopyAll(ctx *Json_valueContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Json_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Json_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type NullValueContext struct {
	Json_valueContext
}

func NewNullValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NullValueContext {
	var p = new(NullValueContext)

	InitEmptyJson_valueContext(&p.Json_valueContext)
	p.parser = parser
	p.CopyAll(ctx.(*Json_valueContext))

	return p
}

func (s *NullValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NullValueContext) NULL() antlr.TerminalNode {
	return s.GetToken(JsonGrammarParserNULL, 0)
}

func (s *NullValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonGrammarListener); ok {
		listenerT.EnterNullValue(s)
	}
}

func (s *NullValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonGrammarListener); ok {
		listenerT.ExitNullValue(s)
	}
}

type ObjectValueContext struct {
	Json_valueContext
}

func NewObjectValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ObjectValueContext {
	var p = new(ObjectValueContext)

	InitEmptyJson_valueContext(&p.Json_valueContext)
	p.parser = parser
	p.CopyAll(ctx.(*Json_valueContext))

	return p
}

func (s *ObjectValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectValueContext) Json_object() IJson_objectContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJson_objectContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJson_objectContext)
}

func (s *ObjectValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonGrammarListener); ok {
		listenerT.EnterObjectValue(s)
	}
}

func (s *ObjectValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonGrammarListener); ok {
		listenerT.ExitObjectValue(s)
	}
}

type BoolValueContext struct {
	Json_valueContext
}

func NewBoolValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoolValueContext {
	var p = new(BoolValueContext)

	InitEmptyJson_valueContext(&p.Json_valueContext)
	p.parser = parser
	p.CopyAll(ctx.(*Json_valueContext))

	return p
}

func (s *BoolValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolValueContext) BOOL() antlr.TerminalNode {
	return s.GetToken(JsonGrammarParserBOOL, 0)
}

func (s *BoolValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonGrammarListener); ok {
		listenerT.EnterBoolValue(s)
	}
}

func (s *BoolValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonGrammarListener); ok {
		listenerT.ExitBoolValue(s)
	}
}

type FloatValueContext struct {
	Json_valueContext
}

func NewFloatValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FloatValueContext {
	var p = new(FloatValueContext)

	InitEmptyJson_valueContext(&p.Json_valueContext)
	p.parser = parser
	p.CopyAll(ctx.(*Json_valueContext))

	return p
}

func (s *FloatValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatValueContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(JsonGrammarParserFLOAT, 0)
}

func (s *FloatValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonGrammarListener); ok {
		listenerT.EnterFloatValue(s)
	}
}

func (s *FloatValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonGrammarListener); ok {
		listenerT.ExitFloatValue(s)
	}
}

type StringValueContext struct {
	Json_valueContext
}

func NewStringValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringValueContext {
	var p = new(StringValueContext)

	InitEmptyJson_valueContext(&p.Json_valueContext)
	p.parser = parser
	p.CopyAll(ctx.(*Json_valueContext))

	return p
}

func (s *StringValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringValueContext) STRING() antlr.TerminalNode {
	return s.GetToken(JsonGrammarParserSTRING, 0)
}

func (s *StringValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonGrammarListener); ok {
		listenerT.EnterStringValue(s)
	}
}

func (s *StringValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonGrammarListener); ok {
		listenerT.ExitStringValue(s)
	}
}

type IntValueContext struct {
	Json_valueContext
}

func NewIntValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntValueContext {
	var p = new(IntValueContext)

	InitEmptyJson_valueContext(&p.Json_valueContext)
	p.parser = parser
	p.CopyAll(ctx.(*Json_valueContext))

	return p
}

func (s *IntValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntValueContext) INT() antlr.TerminalNode {
	return s.GetToken(JsonGrammarParserINT, 0)
}

func (s *IntValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonGrammarListener); ok {
		listenerT.EnterIntValue(s)
	}
}

func (s *IntValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonGrammarListener); ok {
		listenerT.ExitIntValue(s)
	}
}

type ArrayValueContext struct {
	Json_valueContext
}

func NewArrayValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayValueContext {
	var p = new(ArrayValueContext)

	InitEmptyJson_valueContext(&p.Json_valueContext)
	p.parser = parser
	p.CopyAll(ctx.(*Json_valueContext))

	return p
}

func (s *ArrayValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayValueContext) Json_array() IJson_arrayContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJson_arrayContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJson_arrayContext)
}

func (s *ArrayValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonGrammarListener); ok {
		listenerT.EnterArrayValue(s)
	}
}

func (s *ArrayValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonGrammarListener); ok {
		listenerT.ExitArrayValue(s)
	}
}

func (p *JsonGrammarParser) Json_value() (localctx IJson_valueContext) {
	localctx = NewJson_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, JsonGrammarParserRULE_json_value)
	p.SetState(60)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case JsonGrammarParserINT:
		localctx = NewIntValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(53)
			p.Match(JsonGrammarParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case JsonGrammarParserFLOAT:
		localctx = NewFloatValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(54)
			p.Match(JsonGrammarParserFLOAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case JsonGrammarParserSTRING:
		localctx = NewStringValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(55)
			p.Match(JsonGrammarParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case JsonGrammarParserBOOL:
		localctx = NewBoolValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(56)
			p.Match(JsonGrammarParserBOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case JsonGrammarParserNULL:
		localctx = NewNullValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(57)
			p.Match(JsonGrammarParserNULL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case JsonGrammarParserT__1:
		localctx = NewObjectValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(58)
			p.Json_object()
		}

	case JsonGrammarParserT__4:
		localctx = NewArrayValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(59)
			p.Json_array()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
