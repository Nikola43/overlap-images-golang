package main

import (
	"image"
	"image/color"
	"image/draw"
)

type (
	//ImageLayer is a struct
	ImageLayer struct {
		Image image.Image
		XPos  int
		YPos  int
	}

	//BgProperty is background property struct
	BgProperty struct {
		Width   int
		Length  int
		BgColor color.Color
	}
)

// GenerateBanner is a function that combine images and texts into one image
func GenerateBanner(imgs []ImageLayer, width, length int) (*image.RGBA, error) {
	bgImg := image.NewRGBA(image.Rect(0, 0, width, width))

	//looping image layer, higher array index = upper layer
	for _, img := range imgs {
		//set image offset
		offset := image.Pt(img.XPos, img.YPos)

		//combine the image
		draw.Draw(bgImg, img.Image.Bounds().Add(offset), img.Image, image.Point{}, draw.Over)
	}

	return bgImg, nil
}
