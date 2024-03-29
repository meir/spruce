package main

import (
	"fmt"
	"os"

	"github.com/meir/spruce/pkg/states"
	"github.com/meir/spruce/pkg/structure"
)

func main() {
	data, err := os.ReadFile("./examples/index.spr")
	if err != nil {
		panic(err)
	}

	tokens := structure.NewTokens(string(data))
	lexer := structure.NewLexer(tokens, []structure.Node{
		states.NewElementAttributeNode(),
		states.NewElementContentNode(),
		states.NewElementNode(),
		states.NewStringNode(),
	})
	asts := lexer.Parse()
	output := lexer.Format(asts)
	fmt.Println(output)
}
