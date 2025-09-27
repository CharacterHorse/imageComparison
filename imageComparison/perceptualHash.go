package imageComparison

import (
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

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

func SaveImage(img image.Image, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	ext := strings.ToLower(filePath[strings.LastIndex(filePath, "."):])
	switch ext {
	case ".jpg", ".jpeg":
		err = jpeg.Encode(file, img, &jpeg.Options{Quality: 100})
	case ".png":
		err = png.Encode(file, img)
	}
	return err
}
