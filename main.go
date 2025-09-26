package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

func DecodeImage(file *os.File, ext string) (image.Image, error) {
	var decodedImage image.Image
	var err error

	switch ext {
	case ".jpg", ".jpeg":
		decodedImage, err = jpeg.Decode(file)
	case ".png":
		decodedImage, err = png.Decode(file)
	}
	if err != nil {
		return nil, err
	}
	return decodedImage, nil
}

func OpenImage(filePath string) (image.Image, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	ext := strings.ToLower(filePath[strings.LastIndex(filePath, "."):])
	return DecodeImage(file, ext)
}

func ToGray(img image.Image) *image.Gray {
	bounds := img.Bounds()
	newImg := image.NewGray(bounds)

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			originalColor := img.At(x, y)
			newImg.Set(x, y, originalColor)
		}
	}

	return newImg
}

func main() {
	fmt.Println("Image Decoder Example")
	img, err := OpenImage("/home/raghav/Pictures/steam.png")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Image opened successfully:", img.Bounds())
}
