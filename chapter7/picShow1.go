package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
)

func main() {

	rectangle := "rectangle.png"

	rectImage := image.NewRGBA(image.Rect(0, 0, 200, 200))
	green := color.RGBA{0, 100, 0, 255}

	draw.Draw(rectImage, rectImage.Bounds(), &image.Uniform{green}, image.ZP, draw.Src)

	file, err := os.Create(rectangle)
	if err != nil {
		log.Fatalf("failed create file: %s", err)
	}
	png.Encode(file, rectImage)
}
