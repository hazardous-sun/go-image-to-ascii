package main

import (
	"fmt"
	"github.com/nfnt/resize"
	"image"
	"image/png"
	"os"
)

func imageToAscii(config Config) {
	// 1- open image
	file, err := os.Open(imagePath)

	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	defer file.Close()

	// 2 - identify image type

	// 3- resize image
	// 4- turn the image bw
}

func resizeImage(img image.Image, factor float64) image.Image {
	bounds := img.Bounds()
	newWidth := uint(float64(bounds.Dx()) * factor)
	newHeight := uint(float64(bounds.Dy()) * factor)

	return resize.Resize(newWidth, newHeight, img, resize.Lanczos3)
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
