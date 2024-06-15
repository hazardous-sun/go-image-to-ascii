package main

import (
	"fmt"
	"github.com/nfnt/resize"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
)

// ImageData
/*
Holds the image metadata used to resize and transform it into grayscale.
*/
type ImageData struct {
	img    image.Image
	height int
	width  int
	format string
}

// The function responsible for handling the logic of converting the image into ASCII
func imageToAscii(config Config) error {
	// 1- open image
	file, err := os.Open(config.path)

	if err != nil {
		return fmt.Errorf("error opening file: " + err.Error())
	}

	defer file.Close() // close the file after imageToAscii returns

	// 2 - collect metadata

	metadata, err := collectMetadata(file)

	if err != nil {
		return fmt.Errorf("error getting metadata from file: " + err.Error())
	}

	// 3- resize image

	resizeImage(&metadata, config)

	// 4- turn the image into grayscale
	metadata.img, err = removeColor(metadata.img)

	if err != nil {
		return fmt.Errorf("error removing color from image: " + err.Error())
	}

	printPixelsValues(metadata)

	return nil
}

/*
Tries to convert the file to Image and returns some metadata about it.
*/
func collectMetadata(file *os.File) (ImageData, error) {
	format, err := getFormat(file)

	if err != nil {
		return ImageData{}, err
	}

	img, err := loadImage(file, format)

	if err != nil {
		return ImageData{}, err
	}

	bounds := img.Bounds()
	width := bounds.Max.X
	height := bounds.Max.Y

	return ImageData{
		img:    img,
		height: height,
		width:  width,
		format: format,
	}, nil
}

/*
Tries to collect the format of the image.
Will fail if the format is not supported.
*/
func getFormat(file *os.File) (string, error) {
	if isJPEG(file) {
		return "jpeg", nil
	} else if isPNG(file) {
		return "png", nil
	} else if isGIF(file) {
		return "gif", nil
	} else {
		return "", fmt.Errorf("format not supported")
	}
}

/*
Checks if the file is in the JPEG format.
*/
func isJPEG(file *os.File) bool {
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

/*
Checks if the file is in the PNG format.
*/
func isPNG(file *os.File) bool {
	buffer := make([]byte, 8)
	_, err := file.ReadAt(buffer, 0)

	if err != nil {
		return false
	} else if string(buffer) == "PNG\r\n0x1a\n" {
		return true
	} else {
		return false
	}
}

/*
Checks if the file is in the GIF format.
*/
func isGIF(file *os.File) bool {
	buffer := make([]byte, 6)
	_, err := file.ReadAt(buffer, 0)

	if err != nil {
		return false
	}

	return string(buffer) == "GIF89a"
}

/*
Decodes the image considering its format.
*/
func loadImage(file *os.File, format string) (image.Image, error) {
	switch format {
	case "jpeg":
		image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
		return jpeg.Decode(file)
	case "png":
		image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
		return png.Decode(file)
	case "gif":
		image.RegisterFormat("gif", "gif", gif.Decode, gif.DecodeConfig)
		return gif.Decode(file)
	default:
		return nil, fmt.Errorf("unknown format: " + format)
	}
}

/*
Resizes the image based on the value of resizeFactor maintained in ImageData.
*/
func resizeImage(metadata *ImageData, config Config) {
	bounds := metadata.img.Bounds()
	newWidth := uint(float64(bounds.Dx()) * config.resizeFactor)
	newHeight := uint(float64(bounds.Dy()) * config.resizeFactor)

	metadata.img = resize.Resize(newWidth, newHeight, metadata.img, resize.Lanczos3)
}

/*
Transforms the image into grayscale
*/
func removeColor(img image.Image) (image.Image, error) {
	gray := image.NewGray(img.Bounds())
	for y := range gray.Pix {
		for x := range gray.Pix[y:] {
			r, g, b, _ := img.At(x, y).RGBA()
			// Calculate average for grayscale value
			grayVal := (uint8(r) + uint8(g) + uint8(b)) / 3
			gray.Pix[y*gray.Stride+x] = grayVal
		}
	}
	return gray, nil
}

func printPixelsValues(metadata ImageData) {
	for y := 0; y < metadata.height; y++ {
		for x := 0; x < metadata.width; x++ {
			fmt.Println(metadata.img.At(x, y))
		}
	}
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
