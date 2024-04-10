package main

import "github.com/meir/spruce/internal/build"

func main() {
	// get current directory
	dir := "."
	output_dir := "./output"

	build.Build(dir, output_dir)
}
