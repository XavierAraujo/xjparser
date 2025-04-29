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
		"BOOL", "NULL", "WS", "HEX_VALUE", "HEX_CHAR", "INT_NUMBER", "FLOAT_NUMBER",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 12, 149, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5,
		1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 52, 8, 7, 1, 7, 4, 7, 55,
		8, 7, 11, 7, 12, 7, 56, 1, 7, 1, 7, 1, 7, 3, 7, 62, 8, 7, 1, 7, 4, 7, 65,
		8, 7, 11, 7, 12, 7, 66, 3, 7, 69, 8, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 5,
		8, 76, 8, 8, 10, 8, 12, 8, 79, 9, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 92, 8, 9, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 11, 4, 11, 100, 8, 11, 11, 11, 12, 11, 101, 1, 11, 1, 11, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14,
		5, 14, 117, 8, 14, 10, 14, 12, 14, 120, 9, 14, 1, 14, 4, 14, 123, 8, 14,
		11, 14, 12, 14, 124, 1, 14, 5, 14, 128, 8, 14, 10, 14, 12, 14, 131, 9,
		14, 1, 15, 5, 15, 134, 8, 15, 10, 15, 12, 15, 137, 9, 15, 1, 15, 4, 15,
		140, 8, 15, 11, 15, 12, 15, 141, 1, 15, 1, 15, 4, 15, 146, 8, 15, 11, 15,
		12, 15, 147, 0, 0, 16, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15,
		8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 0, 27, 0, 29, 0, 31, 0, 1, 0, 8,
		2, 0, 69, 69, 101, 101, 2, 0, 43, 43, 45, 45, 1, 0, 48, 57, 4, 0, 9, 10,
		13, 13, 34, 34, 92, 92, 7, 0, 34, 34, 92, 92, 98, 98, 102, 102, 110, 110,
		114, 114, 116, 116, 3, 0, 9, 10, 13, 13, 32, 32, 3, 0, 45, 45, 48, 57,
		65, 70, 1, 0, 49, 57, 161, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1,
		0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13,
		1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0,
		21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 1, 33, 1, 0, 0, 0, 3, 35, 1, 0, 0, 0,
		5, 37, 1, 0, 0, 0, 7, 39, 1, 0, 0, 0, 9, 41, 1, 0, 0, 0, 11, 43, 1, 0,
		0, 0, 13, 45, 1, 0, 0, 0, 15, 68, 1, 0, 0, 0, 17, 70, 1, 0, 0, 0, 19, 91,
		1, 0, 0, 0, 21, 93, 1, 0, 0, 0, 23, 99, 1, 0, 0, 0, 25, 105, 1, 0, 0, 0,
		27, 113, 1, 0, 0, 0, 29, 118, 1, 0, 0, 0, 31, 135, 1, 0, 0, 0, 33, 34,
		5, 58, 0, 0, 34, 2, 1, 0, 0, 0, 35, 36, 5, 123, 0, 0, 36, 4, 1, 0, 0, 0,
		37, 38, 5, 44, 0, 0, 38, 6, 1, 0, 0, 0, 39, 40, 5, 125, 0, 0, 40, 8, 1,
		0, 0, 0, 41, 42, 5, 91, 0, 0, 42, 10, 1, 0, 0, 0, 43, 44, 5, 93, 0, 0,
		44, 12, 1, 0, 0, 0, 45, 46, 3, 29, 14, 0, 46, 14, 1, 0, 0, 0, 47, 69, 3,
		31, 15, 0, 48, 49, 3, 31, 15, 0, 49, 51, 7, 0, 0, 0, 50, 52, 7, 1, 0, 0,
		51, 50, 1, 0, 0, 0, 51, 52, 1, 0, 0, 0, 52, 54, 1, 0, 0, 0, 53, 55, 7,
		2, 0, 0, 54, 53, 1, 0, 0, 0, 55, 56, 1, 0, 0, 0, 56, 54, 1, 0, 0, 0, 56,
		57, 1, 0, 0, 0, 57, 69, 1, 0, 0, 0, 58, 59, 3, 29, 14, 0, 59, 61, 7, 0,
		0, 0, 60, 62, 7, 1, 0, 0, 61, 60, 1, 0, 0, 0, 61, 62, 1, 0, 0, 0, 62, 64,
		1, 0, 0, 0, 63, 65, 7, 2, 0, 0, 64, 63, 1, 0, 0, 0, 65, 66, 1, 0, 0, 0,
		66, 64, 1, 0, 0, 0, 66, 67, 1, 0, 0, 0, 67, 69, 1, 0, 0, 0, 68, 47, 1,
		0, 0, 0, 68, 48, 1, 0, 0, 0, 68, 58, 1, 0, 0, 0, 69, 16, 1, 0, 0, 0, 70,
		77, 5, 34, 0, 0, 71, 76, 8, 3, 0, 0, 72, 73, 5, 92, 0, 0, 73, 76, 7, 4,
		0, 0, 74, 76, 3, 25, 12, 0, 75, 71, 1, 0, 0, 0, 75, 72, 1, 0, 0, 0, 75,
		74, 1, 0, 0, 0, 76, 79, 1, 0, 0, 0, 77, 75, 1, 0, 0, 0, 77, 78, 1, 0, 0,
		0, 78, 80, 1, 0, 0, 0, 79, 77, 1, 0, 0, 0, 80, 81, 5, 34, 0, 0, 81, 18,
		1, 0, 0, 0, 82, 83, 5, 116, 0, 0, 83, 84, 5, 114, 0, 0, 84, 85, 5, 117,
		0, 0, 85, 92, 5, 101, 0, 0, 86, 87, 5, 102, 0, 0, 87, 88, 5, 97, 0, 0,
		88, 89, 5, 108, 0, 0, 89, 90, 5, 115, 0, 0, 90, 92, 5, 101, 0, 0, 91, 82,
		1, 0, 0, 0, 91, 86, 1, 0, 0, 0, 92, 20, 1, 0, 0, 0, 93, 94, 5, 110, 0,
		0, 94, 95, 5, 117, 0, 0, 95, 96, 5, 108, 0, 0, 96, 97, 5, 108, 0, 0, 97,
		22, 1, 0, 0, 0, 98, 100, 7, 5, 0, 0, 99, 98, 1, 0, 0, 0, 100, 101, 1, 0,
		0, 0, 101, 99, 1, 0, 0, 0, 101, 102, 1, 0, 0, 0, 102, 103, 1, 0, 0, 0,
		103, 104, 6, 11, 0, 0, 104, 24, 1, 0, 0, 0, 105, 106, 5, 92, 0, 0, 106,
		107, 5, 117, 0, 0, 107, 108, 1, 0, 0, 0, 108, 109, 3, 27, 13, 0, 109, 110,
		3, 27, 13, 0, 110, 111, 3, 27, 13, 0, 111, 112, 3, 27, 13, 0, 112, 26,
		1, 0, 0, 0, 113, 114, 7, 6, 0, 0, 114, 28, 1, 0, 0, 0, 115, 117, 5, 45,
		0, 0, 116, 115, 1, 0, 0, 0, 117, 120, 1, 0, 0, 0, 118, 116, 1, 0, 0, 0,
		118, 119, 1, 0, 0, 0, 119, 122, 1, 0, 0, 0, 120, 118, 1, 0, 0, 0, 121,
		123, 7, 7, 0, 0, 122, 121, 1, 0, 0, 0, 123, 124, 1, 0, 0, 0, 124, 122,
		1, 0, 0, 0, 124, 125, 1, 0, 0, 0, 125, 129, 1, 0, 0, 0, 126, 128, 7, 2,
		0, 0, 127, 126, 1, 0, 0, 0, 128, 131, 1, 0, 0, 0, 129, 127, 1, 0, 0, 0,
		129, 130, 1, 0, 0, 0, 130, 30, 1, 0, 0, 0, 131, 129, 1, 0, 0, 0, 132, 134,
		5, 45, 0, 0, 133, 132, 1, 0, 0, 0, 134, 137, 1, 0, 0, 0, 135, 133, 1, 0,
		0, 0, 135, 136, 1, 0, 0, 0, 136, 139, 1, 0, 0, 0, 137, 135, 1, 0, 0, 0,
		138, 140, 7, 2, 0, 0, 139, 138, 1, 0, 0, 0, 140, 141, 1, 0, 0, 0, 141,
		139, 1, 0, 0, 0, 141, 142, 1, 0, 0, 0, 142, 143, 1, 0, 0, 0, 143, 145,
		5, 46, 0, 0, 144, 146, 7, 2, 0, 0, 145, 144, 1, 0, 0, 0, 146, 147, 1, 0,
		0, 0, 147, 145, 1, 0, 0, 0, 147, 148, 1, 0, 0, 0, 148, 32, 1, 0, 0, 0,
		16, 0, 51, 56, 61, 66, 68, 75, 77, 91, 101, 118, 124, 129, 135, 141, 147,
		1, 6, 0, 0,
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
