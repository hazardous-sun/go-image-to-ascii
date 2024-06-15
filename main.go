package main

import (
	"os"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		printUsage("not enough args passed")
	}

	config, err := build(args[1:])

	if err != nil {
		panic(err)
	}

	err = imageToAscii(config)

	if err != nil {
		panic(err)
	}
}
