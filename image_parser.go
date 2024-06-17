package main

import (
	"fmt"
	"github.com/nfnt/resize"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"image/png"
	"net/http"
	"os"
)

// ImageData
/*
Holds the image metadata used to resize and transform it into grayscale.
*/
type ImageData struct {
	img       image.Image
	height    int
	width     int
	grayscale image.Gray
	format    string
}

// The function responsible for handling the logic of converting the image into ASCII
func imageToAscii(config Config) error {
	// 1- open image
	file, err := os.Open(config.path)

	if err != nil {
		return fmt.Errorf("error: could not open file: \n\t" + err.Error())
	}

	defer file.Close() // close the file after imageToAscii returns

	// 2 - collect metadata

	metadata, err := collectMetadata(file)

	if err != nil {
		return fmt.Errorf("error: could not collect metadata from file: \n\t" + err.Error())
	}

	// 3- resize image

	resizeImage(&metadata, config)

	// 4- turn the image into grayscale
	removeColor(&metadata)

	// 5- Iterate over each pixel and print each value to the cli as an ASCII char
	printPixelsValues(metadata, config)

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
	buffer := make([]byte, 512)
	n, err := file.Read(buffer)

	if err != nil {
		return "", err
	}

	if n < len(buffer) {
		buffer = buffer[:n]
	}

	contentType := http.DetectContentType(buffer[:n])

	switch contentType {
	case "image/png":
		return "png", nil
	case "image/gif":
		return "gif", nil
	case "image/jpeg":
		return "jpeg", nil
	default:
		return "", fmt.Errorf("error: unknown image format: %s \n\t"+
			"buffer: %v", contentType, buffer)
	}
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
func removeColor(metadata *ImageData) {
	gray := image.NewGray(metadata.img.Bounds())
	draw.Draw(gray, gray.Bounds(), metadata.img, image.Point{}, draw.Src)
	metadata.grayscale = *gray
}

func printPixelsValues(metadata ImageData, config Config) {
	for y := 0; y < metadata.height; y++ {
		for x := 0; x < metadata.width; x++ {
			fmt.Print(getChar(metadata.grayscale.At(x, y), config.reverse))
		}
		fmt.Println()
	}
}

func getChar(grayscale color.Color, reverse bool) string {
	chars := []string{
		" ",
		"□",
		"▧",
		"▥",
		"▩",
		"▦",
		"▣",
		"■",
	}

	r, g, b, intensity := grayscale.RGBA()
	fmt.Println(r, g, b, intensity)

	if reverse {
		switch {
		case intensity > 225:
			return chars[0]
		case intensity > 193:
			return chars[1]
		case intensity > 161:
			return chars[2]
		case intensity > 129:
			return chars[3]
		case intensity > 97:
			return chars[4]
		case intensity > 65:
			return chars[5]
		case intensity > 33:
			return chars[6]
		default:
			return chars[7]
		}
	}

	switch {
	case intensity < 32:
		return chars[0]
	case intensity < 64:
		return chars[1]
	case intensity < 96:
		return chars[2]
	case intensity < 128:
		return chars[3]
	case intensity < 160:
		return chars[4]
	case intensity < 192:
		return chars[5]
	case intensity < 224:
		return chars[6]
	default:
		return chars[7]
	}
}
