package main

import (
	"fmt"
	"github.com/nfnt/resize"
	"os"
	"strconv"
	"strings"
)

// Used to print error messages to the CLI
func printUsage(errorMessage string) {
	fmt.Printf(
		"Error: %s\n"+
			"USAGE:\n\t"+
			"image-to-ascii [OPTIONS] [PATH] [RESIZE_FACTOR]\n"+
			"OPTIONS:\n\t"+
			"-r | --reverse\tReverses the image ASCII values\n\t"+
			"--l2\tUses Lanczos2 interpolation method\n\t"+
			"--l3\tUses Lanczos3 interpolation method\n\t"+
			"--bc\tUses Bicubic interpolation method\n\t"+
			"--bl\tUses Bilinear interpolation method\n\t"+
			"--nn\tUses Nearest Neighbor interpolation method\n\t"+
			"--mn\tUses Mitchell Netravali interpolation method\n",
		errorMessage)
}

// Config
/*
The configuration used during runtime. Contains the path to the image, the resize factor and the options.
*/
type Config struct {
	path          string
	resizeFactor  float64
	reverse       bool
	interpolation resize.InterpolationFunction
	options       []string
}

/*
Returns a Config struct.
May fail depending on the inputs provided.
*/
func build(args []string) (Config, error) {
	options, resizeFactor, err := getValues(args)

	if err != nil {
		config := Config{options: options}
		config.analyzeOptions()
		return config, err
	}

	config := Config{
		path:          options[0],
		resizeFactor:  resizeFactor,
		reverse:       false,
		interpolation: resize.Bicubic,
		options:       options[1:],
	}
	config.analyzeOptions()
	return config, nil
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
		return options, 0, fmt.Errorf("no path specified")
	}

	if resizeFactor == 0 {
		return options, 0, fmt.Errorf("resize factor cannot be zero")
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

/*
Checks which adjustments needs to be done for each option passed.
*/
func (c *Config) analyzeOptions() {
	for _, option := range c.options {
		switch option {
		case "-r":
			c.reverse = true
		case "--reverse":
			c.reverse = true
		case "--l3":
			c.interpolation = resize.Lanczos3
		case "--l2":
			c.interpolation = resize.Lanczos2
		case "--bc":
			c.interpolation = resize.Bicubic
		case "--bl":
			c.interpolation = resize.Bilinear
		case "--nn":
			c.interpolation = resize.NearestNeighbor
		case "--mn":
			c.interpolation = resize.MitchellNetravali
		case "-h":
			printUsage("Help page called")
			os.Exit(0)
		case "--help":
			printUsage("Help page called")
			os.Exit(0)
		default:
			fmt.Printf("Invalid option: '%s'\n", option)
		}
	}
}
