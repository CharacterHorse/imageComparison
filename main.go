package main

import (
	"fmt"
)

func main() {
	fmt.Println("Image Decoder Example")
	img, err := OpenImage("/home/raghav/Pictures/spider-man.jpg")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Image opened successfully:", img.Bounds())
}
