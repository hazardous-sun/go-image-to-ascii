package go_image_to_ascii

import (
	"fmt"
	"image"
	"image/png"
	"os"
)

type Config struct {
	path         string
	resizeFactor float64
}

func generateConfig() (*Config, error) {
	args := os.Args

	if len(args) < 2 {
		return nil, fmt.Errorf("not enough args passed")
	}

	path := getImagePath(args)
}

func printUsage(errorMessage string) {
	fmt.Printf(
		"Error: %s\n USAGE:\n\timage-to-ascii [OPTIONS] -p [PATH]\n\timage-to-ascii [OPTIONS] --path [PATH]\n",
		errorMessage)
}

func getImagePath(args []string) (string, error) {
	var values []string
	var path string

	for _, v := range args {
		if isValidPath(v) {
			path = v
		}
	}

	return args[1], nil
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
