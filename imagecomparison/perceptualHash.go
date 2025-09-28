package imagecomparison

import (
	"image"
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
