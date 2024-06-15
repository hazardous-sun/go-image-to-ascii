package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Used to print error messages to the CLI
func printUsage(errorMessage string) {
	fmt.Printf(
		"Error: %s\n" +
			"USAGE:\n\t" +
			"image-to-ascii [OPTIONS] [PATH] [RESIZE_FACTOR]\n" +
			errorMessage)
}

// Config
/*
The configuration used during runtime. Contains the path to the image, the resize factor and the options.
*/
type Config struct {
	path         string
	resizeFactor float64
	reverse      bool
	options      []string
}

/*
Returns a Config struct.
May fail depending on the inputs provided.
*/
func build(args []string) (Config, error) {
	values, resizeFactor, err := getValues(args)

	if err != nil {
		return Config{}, err
	}

	return Config{
		path:         values[0],
		resizeFactor: resizeFactor,
		reverse:      false,
		options:      values[1:],
	}, nil
}

/*
Collects the values passed to the main function and checks if they are valid.
Will fail when no path or resize factor are passed.
*/
func getValues(args []string) ([]string, float64, error) {
	var options []string
	var path string
	var resizeFactor float64

	for _, v := range args {
		v = strings.TrimSpace(v)
		if isValidOption(v) {
			options = append(options, v)
		} else if isValidPath(v) {
			path = v
		} else {
			factor, err := getResizeFactor(v)
			if err == nil {
				resizeFactor = factor
			}
		}
	}

	if len(path) == 0 {
		return []string{}, 0, fmt.Errorf("no path specified")
	}

	if resizeFactor == 0 {
		return []string{}, 0, fmt.Errorf("resize factor cannot be zero")
	}

	return append([]string{path}, options...), resizeFactor, nil
}

/*
Checks if the value is a valid option. An option should begin with either "-" or "--"
*/
func isValidOption(option string) bool {
	if len(option) == 0 {
		return false
	}

	return option[0] == '-' || (option[0:2] == "--" && len(option) > 2)
}

/*
Checks if the value is a valid path.
*/
func isValidPath(path string) bool {
	_, err := os.Stat(path)

	if err != nil {
		return false
	}

	return true
}

/*
Checks if the value is a valid float64 and returns it if the parse succeeds.
*/
func getResizeFactor(value string) (float64, error) {
	factor, err := strconv.ParseFloat(value, 64)

	if err != nil {
		return 0, err
	}

	return factor, nil
}
