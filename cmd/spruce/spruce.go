package main

import (
	"github.com/meir/spruce/internal/spruce"
	_ "github.com/meir/spruce/internal/spruce/ast"
)

func main() {
	file := "./examples/index.spr"

	ctx, err := spruce.Parse(file)
	if err != nil {
		panic(err)
	}

	result := spruce.String(ctx)
	println(result)
}
