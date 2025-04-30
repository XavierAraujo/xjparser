package main

import (
	"fmt"
	"xjparser/xjparser"
)

func main() {
	input := "{\"bla\": {\"blo\": \"bli\"}}"
	fmt.Println("Parsing string: ", input)
	xjparser.Parse(input)
}
