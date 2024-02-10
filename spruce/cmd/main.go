package main

import (
	"fmt"
	"os"

	"github.com/meir/spruce/pkg/structure"
)

func main() {
	data, err := os.ReadFile("./examples/index.spr")
	if err != nil {
		panic(err)
	}

	tokens := structure.NewTokens(string(data))
	for {
		t := tokens.Next()
		if t.Str == "\n" {
			panic(fmt.Errorf("unexpected newline at %v:%v", t.Line, t.Start))
		}
	}
}
