package main

import (
	"fmt"
	"github.com/nfnt/resize"
	"image"
	"image/png"
	"os"
)

type ImageData struct {
	img    image.Image
	height int
	width  int
	format string
}

func imageToAscii(config Config) {
	// 1- open image
	file, err := os.Open(config.path)

	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	defer file.Close() // close the file after imageToAscii returns

	// 2 - identify image type

	// 3- resize image
	// 4- turn the image bw
}

/*
Tries to convert the file to Image and returns some metadata about it.
Will fail if the file is an unsupported format.
*/
func convertImage(file os.File) (ImageData, error) {
	var img image.Image
	var height, width int
	var format string

	return ImageData{
		img:    img,
		height: height,
		width:  width,
		format: format,
	}, nil
}

func getFormat(file *os.File) (*image.Image, error) {
	if isJPG(file) {
		return nil, nil
	} else if isPNG(file) {
		return nil, nil
	} else if isGIF(file) {
		return nil, nil
	} else if isWEBP(file) {
		return nil, nil
	} else {
		return nil, fmt.Errorf("format not supported")
	}
}

func isJPG(file *os.File) bool {
	buffer := make([]byte, 2)
	_, err := file.ReadAt(buffer, 0)

	if err != nil {
		return false
	} else if buffer[0] == 0xff && buffer[1] == 0xd8 {
		return true
	} else {
		return false
	}
}

func isPNG(file *os.File) bool {
	buffer := make([]byte, 8)
	_, err := file.ReadAt(buffer, 0)

	if err != nil {
		return false
	} else if buffer[0] == 0x89 && buffer[1] == 'P' && buffer[2] == 'N' && buffer[3] == 'G' &&
		buffer[4] == '\r' && buffer[5] == '\n' && buffer[6] == 0x1a && buffer[7] == '\n' {
		return true
	} else {
		return false
	}
}

func isGIF(file *os.File) bool {
	return false
}

func isWEBP(file *os.File) bool {
	return false
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
