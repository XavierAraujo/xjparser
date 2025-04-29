// Code generated from JsonGrammar.g4 by ANTLR 4.13.2. DO NOT EDIT.

package antlr4_parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type JsonGrammarLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var JsonGrammarLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func jsongrammarlexerLexerInit() {
	staticData := &JsonGrammarLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "':'", "'{'", "','", "'}'", "'['", "']'", "", "", "", "", "'null'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "INT", "FLOAT", "STRING", "BOOL", "NULL",
		"WS",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "INT", "FLOAT", "STRING",
		"BOOL", "NULL", "WS", "HEX_VALUE", "HEX_CHAR",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 12, 126, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 1, 0, 1, 0, 1, 1,
		1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 5, 6, 43, 8,
		6, 10, 6, 12, 6, 46, 9, 6, 1, 6, 4, 6, 49, 8, 6, 11, 6, 12, 6, 50, 1, 6,
		5, 6, 54, 8, 6, 10, 6, 12, 6, 57, 9, 6, 1, 7, 5, 7, 60, 8, 7, 10, 7, 12,
		7, 63, 9, 7, 1, 7, 4, 7, 66, 8, 7, 11, 7, 12, 7, 67, 1, 7, 5, 7, 71, 8,
		7, 10, 7, 12, 7, 74, 9, 7, 1, 7, 1, 7, 4, 7, 78, 8, 7, 11, 7, 12, 7, 79,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 5, 8, 87, 8, 8, 10, 8, 12, 8, 90, 9, 8, 1,
		8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 103,
		8, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 4, 11, 111, 8, 11, 11,
		11, 12, 11, 112, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 13, 1, 13, 0, 0, 14, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11,
		6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 0, 27, 0, 1, 0, 6,
		1, 0, 49, 57, 1, 0, 48, 57, 4, 0, 9, 10, 13, 13, 34, 34, 92, 92, 7, 0,
		34, 34, 92, 92, 98, 98, 102, 102, 110, 110, 114, 114, 116, 116, 3, 0, 9,
		10, 13, 13, 32, 32, 3, 0, 45, 45, 48, 57, 65, 70, 135, 0, 1, 1, 0, 0, 0,
		0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0,
		0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0,
		0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 1, 29, 1, 0,
		0, 0, 3, 31, 1, 0, 0, 0, 5, 33, 1, 0, 0, 0, 7, 35, 1, 0, 0, 0, 9, 37, 1,
		0, 0, 0, 11, 39, 1, 0, 0, 0, 13, 44, 1, 0, 0, 0, 15, 61, 1, 0, 0, 0, 17,
		81, 1, 0, 0, 0, 19, 102, 1, 0, 0, 0, 21, 104, 1, 0, 0, 0, 23, 110, 1, 0,
		0, 0, 25, 116, 1, 0, 0, 0, 27, 124, 1, 0, 0, 0, 29, 30, 5, 58, 0, 0, 30,
		2, 1, 0, 0, 0, 31, 32, 5, 123, 0, 0, 32, 4, 1, 0, 0, 0, 33, 34, 5, 44,
		0, 0, 34, 6, 1, 0, 0, 0, 35, 36, 5, 125, 0, 0, 36, 8, 1, 0, 0, 0, 37, 38,
		5, 91, 0, 0, 38, 10, 1, 0, 0, 0, 39, 40, 5, 93, 0, 0, 40, 12, 1, 0, 0,
		0, 41, 43, 5, 45, 0, 0, 42, 41, 1, 0, 0, 0, 43, 46, 1, 0, 0, 0, 44, 42,
		1, 0, 0, 0, 44, 45, 1, 0, 0, 0, 45, 48, 1, 0, 0, 0, 46, 44, 1, 0, 0, 0,
		47, 49, 7, 0, 0, 0, 48, 47, 1, 0, 0, 0, 49, 50, 1, 0, 0, 0, 50, 48, 1,
		0, 0, 0, 50, 51, 1, 0, 0, 0, 51, 55, 1, 0, 0, 0, 52, 54, 7, 1, 0, 0, 53,
		52, 1, 0, 0, 0, 54, 57, 1, 0, 0, 0, 55, 53, 1, 0, 0, 0, 55, 56, 1, 0, 0,
		0, 56, 14, 1, 0, 0, 0, 57, 55, 1, 0, 0, 0, 58, 60, 5, 45, 0, 0, 59, 58,
		1, 0, 0, 0, 60, 63, 1, 0, 0, 0, 61, 59, 1, 0, 0, 0, 61, 62, 1, 0, 0, 0,
		62, 65, 1, 0, 0, 0, 63, 61, 1, 0, 0, 0, 64, 66, 7, 0, 0, 0, 65, 64, 1,
		0, 0, 0, 66, 67, 1, 0, 0, 0, 67, 65, 1, 0, 0, 0, 67, 68, 1, 0, 0, 0, 68,
		72, 1, 0, 0, 0, 69, 71, 7, 1, 0, 0, 70, 69, 1, 0, 0, 0, 71, 74, 1, 0, 0,
		0, 72, 70, 1, 0, 0, 0, 72, 73, 1, 0, 0, 0, 73, 75, 1, 0, 0, 0, 74, 72,
		1, 0, 0, 0, 75, 77, 5, 46, 0, 0, 76, 78, 7, 1, 0, 0, 77, 76, 1, 0, 0, 0,
		78, 79, 1, 0, 0, 0, 79, 77, 1, 0, 0, 0, 79, 80, 1, 0, 0, 0, 80, 16, 1,
		0, 0, 0, 81, 88, 5, 34, 0, 0, 82, 87, 8, 2, 0, 0, 83, 84, 5, 92, 0, 0,
		84, 87, 7, 3, 0, 0, 85, 87, 3, 25, 12, 0, 86, 82, 1, 0, 0, 0, 86, 83, 1,
		0, 0, 0, 86, 85, 1, 0, 0, 0, 87, 90, 1, 0, 0, 0, 88, 86, 1, 0, 0, 0, 88,
		89, 1, 0, 0, 0, 89, 91, 1, 0, 0, 0, 90, 88, 1, 0, 0, 0, 91, 92, 5, 34,
		0, 0, 92, 18, 1, 0, 0, 0, 93, 94, 5, 116, 0, 0, 94, 95, 5, 114, 0, 0, 95,
		96, 5, 117, 0, 0, 96, 103, 5, 101, 0, 0, 97, 98, 5, 102, 0, 0, 98, 99,
		5, 97, 0, 0, 99, 100, 5, 108, 0, 0, 100, 101, 5, 115, 0, 0, 101, 103, 5,
		101, 0, 0, 102, 93, 1, 0, 0, 0, 102, 97, 1, 0, 0, 0, 103, 20, 1, 0, 0,
		0, 104, 105, 5, 110, 0, 0, 105, 106, 5, 117, 0, 0, 106, 107, 5, 108, 0,
		0, 107, 108, 5, 108, 0, 0, 108, 22, 1, 0, 0, 0, 109, 111, 7, 4, 0, 0, 110,
		109, 1, 0, 0, 0, 111, 112, 1, 0, 0, 0, 112, 110, 1, 0, 0, 0, 112, 113,
		1, 0, 0, 0, 113, 114, 1, 0, 0, 0, 114, 115, 6, 11, 0, 0, 115, 24, 1, 0,
		0, 0, 116, 117, 5, 92, 0, 0, 117, 118, 5, 117, 0, 0, 118, 119, 1, 0, 0,
		0, 119, 120, 3, 27, 13, 0, 120, 121, 3, 27, 13, 0, 121, 122, 3, 27, 13,
		0, 122, 123, 3, 27, 13, 0, 123, 26, 1, 0, 0, 0, 124, 125, 7, 5, 0, 0, 125,
		28, 1, 0, 0, 0, 12, 0, 44, 50, 55, 61, 67, 72, 79, 86, 88, 102, 112, 1,
		6, 0, 0,
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

// JsonGrammarLexerInit initializes any static state used to implement JsonGrammarLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewJsonGrammarLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func JsonGrammarLexerInit() {
	staticData := &JsonGrammarLexerLexerStaticData
	staticData.once.Do(jsongrammarlexerLexerInit)
}

// NewJsonGrammarLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewJsonGrammarLexer(input antlr.CharStream) *JsonGrammarLexer {
	JsonGrammarLexerInit()
	l := new(JsonGrammarLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &JsonGrammarLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "JsonGrammar.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// JsonGrammarLexer tokens.
const (
	JsonGrammarLexerT__0   = 1
	JsonGrammarLexerT__1   = 2
	JsonGrammarLexerT__2   = 3
	JsonGrammarLexerT__3   = 4
	JsonGrammarLexerT__4   = 5
	JsonGrammarLexerT__5   = 6
	JsonGrammarLexerINT    = 7
	JsonGrammarLexerFLOAT  = 8
	JsonGrammarLexerSTRING = 9
	JsonGrammarLexerBOOL   = 10
	JsonGrammarLexerNULL   = 11
	JsonGrammarLexerWS     = 12
)
