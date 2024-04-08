package main

import (
	"fmt"

	"github.com/meir/spruce/internal/spruce"
)

func main() {
	file, err := spruce.Parse("./examples/index.spr")
	if err != nil {
		panic(err)
	}

	fmt.Println(file.String())
}
