package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
	"xjparser/xjparser/antlr4_parser"

	"github.com/stretchr/testify/assert"
)

func TestWronglyFormattedJsonStrings(t *testing.T) {
	nWronglyFormattedFiles := 33
	for i := range nWronglyFormattedFiles {
		var builder strings.Builder
		builder.WriteString("test_files/fail")
		builder.WriteString(strconv.Itoa(i + 1))
		builder.WriteString(".json")
		filename := builder.String()
		fmt.Println("Testing wrongly formatted file: ", filename)
		data, err := os.ReadFile(filename)
		assert.Nil(t, err)
		_, err = antlr4_parser.Parse(string(data))
		assert.NotNil(t, err)
	}
}

func TestCorrectlyFormattedJsonStrings(t *testing.T) {
	nCorrectlyFormattedFiles := 3
	for i := range nCorrectlyFormattedFiles {
		var builder strings.Builder
		builder.WriteString("test_files/pass")
		builder.WriteString(strconv.Itoa(i + 1))
		builder.WriteString(".json")
		filename := builder.String()
		fmt.Println("Testing correctly formatted file: ", filename)
		data, err := os.ReadFile(filename)
		assert.Nil(t, err)
		_, err = antlr4_parser.Parse(string(data))
		assert.Nil(t, err)
	}
}
