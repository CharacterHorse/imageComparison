package main

import (
	"fmt"
	"github.com/CharacterHorse/imagecomparison/imagecomparison"
)

func main() {
	fmt.Println("Image Decoder Example")
	img, err := imagecomparison.OpenImage("/home/raghav/nyx.jpg")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Image opened successfully:", img.Bounds())
}
