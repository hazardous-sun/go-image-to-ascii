package main

import (
	"fmt"
	"image"
	"image/png"
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
	options      []string
}

/*
Returns a Config struct.
May fail depending on the inputs provided.
*/
func build(args []string) (*Config, error) {
	values, resizeFactor, err := getValues(args)

	if err != nil {
		return nil, err
	}

	config := Config{
		path:         values[0],
		resizeFactor: resizeFactor,
		options:      values[1:],
	}

	return &config, nil
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

func isValidOption(option string) bool {
	if len(option) == 0 {
		return false
	}

	return option[0] == '-' || (option[0:2] == "--" && len(option) > 2)
}

func isValidPath(path string) bool {
	_, err := os.Stat(path)

	if err != nil {
		return false
	}

	return true
}

func getResizeFactor(value string) (float64, error) {
	factor, err := strconv.ParseFloat(value, 64)

	if err != nil {
		return 0, err
	}

	return factor, nil
}

func readImage(imagePath string) {
	// Open the image file
	file, err := os.Open(imagePath)

	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	defer file.Close()

	// Register PNG decoder
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)

	// Decode the image
	img, err := png.Decode(file)
	if err != nil {
		fmt.Println("Error decoding image:", err)
		return
	}

	// Get image bounds
	bounds := img.Bounds()

	// Loop through pixels and access RGB values
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()

			// Convert to uint8 (optional)
			red := uint8(r >> 8)
			green := uint8(g >> 8)
			blue := uint8(b >> 8)

			fmt.Printf("Pixel (%d, %d): R: %d, G: %d, B: %d\n", x, y, red, green, blue)
		}
	}
}
