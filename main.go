package main

import (
	"os"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		printUsage("not enough args passed")
	}

	_, err := build(args[1:])

	if err != nil {
		println(err.Error())
	}
}
