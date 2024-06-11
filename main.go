package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		printUsage("not enough args passed")
	}

	config, err := build(args[1:])

	if err != nil {
		println(err.Error())
	}

	fmt.Println("config:", config)
}
