package imageComparison

import (
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
