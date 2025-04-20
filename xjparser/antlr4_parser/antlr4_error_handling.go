package antlr4_parser

import (
	"log"

	"github.com/antlr4-go/antlr/v4"
)

type JsonParsingError struct {
	msg string
}

func (parsingError JsonParsingError) Error() string {
	return parsingError.msg
}

type JsonDuplicatedKeyError struct {
	key string
}

func (parsingError JsonDuplicatedKeyError) Error() string {
	return "Duplicated JSON key: " + parsingError.key
}

type JsonErrorListener struct {
}

func (listener *JsonErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	panic(JsonParsingError{msg: msg})
}

func (listener *JsonErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	log.Println("ANTLR4 reported ambiguity when parsing JSON string")
}

func (listener *JsonErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	log.Println("ANTLR4 only able to parse JSON string had to fallback from SLL to LL parsing mode")
}

func (listener *JsonErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs *antlr.ATNConfigSet) {
	log.Println("ANTLR4 only able to parse JSON string using LL parsing mode")
}
