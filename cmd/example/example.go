package main

import (
	"fmt"
	"xjparser/xjparser/antlr4_parser"
)

func main() {
	input := "{  \"bla\": {\"blo\":\"bli\"}}"
	fmt.Println("parsing: ", input)
	antlr4_parser.Parse(input)
}
