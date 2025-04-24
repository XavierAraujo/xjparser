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
		"BOOL", "NULL", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 12, 99, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3,
		1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 4, 6, 39, 8, 6, 11, 6, 12, 6, 40, 1, 6, 5,
		6, 44, 8, 6, 10, 6, 12, 6, 47, 9, 6, 1, 7, 4, 7, 50, 8, 7, 11, 7, 12, 7,
		51, 1, 7, 5, 7, 55, 8, 7, 10, 7, 12, 7, 58, 9, 7, 1, 7, 1, 7, 4, 7, 62,
		8, 7, 11, 7, 12, 7, 63, 1, 8, 1, 8, 1, 8, 1, 8, 5, 8, 70, 8, 8, 10, 8,
		12, 8, 73, 9, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 3, 9, 86, 8, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 4,
		11, 94, 8, 11, 11, 11, 12, 11, 95, 1, 11, 1, 11, 0, 0, 12, 1, 1, 3, 2,
		5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 1,
		0, 4, 1, 0, 49, 57, 1, 0, 48, 57, 3, 0, 9, 10, 13, 13, 34, 34, 3, 0, 9,
		10, 13, 13, 32, 32, 107, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0,
		0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1,
		0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21,
		1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 1, 25, 1, 0, 0, 0, 3, 27, 1, 0, 0, 0, 5,
		29, 1, 0, 0, 0, 7, 31, 1, 0, 0, 0, 9, 33, 1, 0, 0, 0, 11, 35, 1, 0, 0,
		0, 13, 38, 1, 0, 0, 0, 15, 49, 1, 0, 0, 0, 17, 65, 1, 0, 0, 0, 19, 85,
		1, 0, 0, 0, 21, 87, 1, 0, 0, 0, 23, 93, 1, 0, 0, 0, 25, 26, 5, 58, 0, 0,
		26, 2, 1, 0, 0, 0, 27, 28, 5, 123, 0, 0, 28, 4, 1, 0, 0, 0, 29, 30, 5,
		44, 0, 0, 30, 6, 1, 0, 0, 0, 31, 32, 5, 125, 0, 0, 32, 8, 1, 0, 0, 0, 33,
		34, 5, 91, 0, 0, 34, 10, 1, 0, 0, 0, 35, 36, 5, 93, 0, 0, 36, 12, 1, 0,
		0, 0, 37, 39, 7, 0, 0, 0, 38, 37, 1, 0, 0, 0, 39, 40, 1, 0, 0, 0, 40, 38,
		1, 0, 0, 0, 40, 41, 1, 0, 0, 0, 41, 45, 1, 0, 0, 0, 42, 44, 7, 1, 0, 0,
		43, 42, 1, 0, 0, 0, 44, 47, 1, 0, 0, 0, 45, 43, 1, 0, 0, 0, 45, 46, 1,
		0, 0, 0, 46, 14, 1, 0, 0, 0, 47, 45, 1, 0, 0, 0, 48, 50, 7, 0, 0, 0, 49,
		48, 1, 0, 0, 0, 50, 51, 1, 0, 0, 0, 51, 49, 1, 0, 0, 0, 51, 52, 1, 0, 0,
		0, 52, 56, 1, 0, 0, 0, 53, 55, 7, 1, 0, 0, 54, 53, 1, 0, 0, 0, 55, 58,
		1, 0, 0, 0, 56, 54, 1, 0, 0, 0, 56, 57, 1, 0, 0, 0, 57, 59, 1, 0, 0, 0,
		58, 56, 1, 0, 0, 0, 59, 61, 5, 46, 0, 0, 60, 62, 7, 1, 0, 0, 61, 60, 1,
		0, 0, 0, 62, 63, 1, 0, 0, 0, 63, 61, 1, 0, 0, 0, 63, 64, 1, 0, 0, 0, 64,
		16, 1, 0, 0, 0, 65, 71, 5, 34, 0, 0, 66, 70, 8, 2, 0, 0, 67, 68, 5, 92,
		0, 0, 68, 70, 9, 0, 0, 0, 69, 66, 1, 0, 0, 0, 69, 67, 1, 0, 0, 0, 70, 73,
		1, 0, 0, 0, 71, 69, 1, 0, 0, 0, 71, 72, 1, 0, 0, 0, 72, 74, 1, 0, 0, 0,
		73, 71, 1, 0, 0, 0, 74, 75, 5, 34, 0, 0, 75, 18, 1, 0, 0, 0, 76, 77, 5,
		116, 0, 0, 77, 78, 5, 114, 0, 0, 78, 79, 5, 117, 0, 0, 79, 86, 5, 101,
		0, 0, 80, 81, 5, 102, 0, 0, 81, 82, 5, 97, 0, 0, 82, 83, 5, 108, 0, 0,
		83, 84, 5, 115, 0, 0, 84, 86, 5, 101, 0, 0, 85, 76, 1, 0, 0, 0, 85, 80,
		1, 0, 0, 0, 86, 20, 1, 0, 0, 0, 87, 88, 5, 110, 0, 0, 88, 89, 5, 117, 0,
		0, 89, 90, 5, 108, 0, 0, 90, 91, 5, 108, 0, 0, 91, 22, 1, 0, 0, 0, 92,
		94, 7, 3, 0, 0, 93, 92, 1, 0, 0, 0, 94, 95, 1, 0, 0, 0, 95, 93, 1, 0, 0,
		0, 95, 96, 1, 0, 0, 0, 96, 97, 1, 0, 0, 0, 97, 98, 6, 11, 0, 0, 98, 24,
		1, 0, 0, 0, 10, 0, 40, 45, 51, 56, 63, 69, 71, 85, 95, 1, 6, 0, 0,
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
