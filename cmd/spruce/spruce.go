package main

import "github.com/meir/spruce/internal/spruce"

func main() {
	// get current directory
	dir := "."
	output_dir := "./build"

	spruce.Build(dir, output_dir)
}
