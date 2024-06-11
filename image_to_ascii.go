package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
)

func printUsage(errorMessage string) {
	fmt.Printf(
		"Error: %s\n"+
			"USAGE:\n\t"+
			"image-to-ascii [OPTIONS] [PATH]\n\t"+
			"image-to-ascii [OPTIONS] -p [PATH]\n\t"+
			"image-to-ascii [OPTIONS] --path [PATH]\n"+
			"",
		errorMessage)
}

type Config struct {
	path         string
	resizeFactor float64
}

func build(args []string) (*Config, error) {
	values, err := getValues(args)

	if err != nil {
		return nil, err
	}

	fmt.Println("build ran successfully!\n", values)
	return nil, nil
}

func getValues(args []string) ([]string, error) {
	var options []string
	var path string

	for _, v := range args {
		if isValidOption(v) {
			options = append(options, v)
		} else if isValidPath(v) {
			path = v
		}
	}

	if len(path) == 0 {
		return []string{}, fmt.Errorf("no path specified")
	}

	return append([]string{path}, options...), nil
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
