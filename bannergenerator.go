package main

import (
	bannergenerator "github.com/arrafiv/img-txt-combiner"
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
func GenerateBanner(imgs []ImageLayer) (*image.RGBA, error) {
	bg := bannergenerator.BgProperty{
		Width:   512,
		Length:  512,
		BgColor: color.Transparent,
	}

	//create image's background
	bgImg := image.NewRGBA(image.Rect(0, 0, bg.Width, bg.Length))

	//set the background color
	draw.Draw(bgImg, bgImg.Bounds(), image.Transparent, image.ZP, draw.Src)

	//looping image layer, higher array index = upper layer
	for _, img := range imgs {
		//set image offset
		offset := image.Pt(img.XPos, img.YPos)

		//combine the image
		draw.Draw(bgImg, img.Image.Bounds().Add(offset), img.Image, image.ZP, draw.Over)
	}

	return bgImg, nil
}
