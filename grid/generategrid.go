package main

import (
	"fmt"
	"github.com/h2non/bimg"
	"image/jpeg"
	"image/png"
	"os"
	"strconv"
	"strings"
)
import gim "github.com/ozankasikci/go-image-merge"

func main() {

	generatedType := "rare"

	/*
	//for i := 4445; i < 6933; i++ {
	for i := 6934; i < 8266; i++ {
		counterString := strconv.Itoa(i)
		ResizeImage("./../Pinguins/"+generatedType+"/original/baby_"+counterString+".png", "./../Pinguins/"+generatedType+"/resized/", 72)
	}
	os.Exit(0)
	*/

	/*
		for i := 0; i < 2; i++ {
			counterString := strconv.Itoa(i)
			AddLabel("./generatedResized/baby_" + counterString + ".png")
		}
		os.Exit(1)
	*/

	// accepts *Grid instances, grid unit count x, grid unit count y
	// returns an *image.RGBA object
	grids := []*gim.Grid{}

	for i := 6934; i < 8266; i++ {
		counterString := strconv.Itoa(i)
		currentFile := "./../Pinguins/"+generatedType+"/resized/baby_" + counterString + ".png"

		g := &gim.Grid{
			//ImageFilePath: "generatedOR/baby_" + counterString + ".png",
			ImageFilePath: currentFile,
		}
		grids = append(grids, g)
		fmt.Println(counterString)

		if i%1024 == 0 {

			rgba, err := gim.New(grids, 32, 32).Merge()
			if err != nil {
				panic(err)
			}

			// save the output to jpg or png
			file, err2 := os.Create("./../Pinguins/"+generatedType+"/merged/" + counterString + ".png")
			if err2 != nil {
				panic(err2)
			}
			err = jpeg.Encode(file, rgba, &jpeg.Options{Quality: 100})
			err = png.Encode(file, rgba)
			grids = []*gim.Grid{}
		}

	}
}

func AddLabel(path string) {
	/*
		buffer, err := bimg.Read(path)
		if err != nil {
			fmt.Println("os.Stderr, err")
			fmt.Fprintln(os.Stderr, err)
		}


			watermark := bimg.Watermark{
				Text:       "1",
				Opacity:    1,
				Width:      48,
				DPI:        100,
				Background: bimg.Color{R: 255, G: 255, B: 255},
			}

		newImage, err := bimg.NewImage(buffer).Resize(48, 48)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}


			newImage2, err2 := bimg.NewImage(newImage).Watermark(watermark)
			if err != nil {
				fmt.Println("os.Stderr, err2")
				fmt.Fprintln(os.Stderr, err2)
			}


		size, err := bimg.NewImage(newImage).Size()
		if size.Width == size && size.Height == 48 {
			fmt.Println("The image size is valid " + path)
		}

		s := strings.Split(path, "/")
		bimg.Write("./generatedResizedLabel/"+s[len(s)-1], newImage)
	*/
}

func ResizeImage(path string, out string, width int) {
	buffer, err := bimg.Read(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	newImage, err := bimg.NewImage(buffer).Resize(width, width)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	/*
		watermark := bimg.Watermark{
			Text:       "1",
			Opacity:    0.5,
			Width:      25,
			DPI:        100,
			Background: bimg.Color{R: 255, G: 255, B: 255},
		}





		newImage, err := bimg.NewImage(buffer).Watermark(watermark)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	*/

	size, err := bimg.NewImage(newImage).Size()
	if size.Width == width && size.Height == width {
		fmt.Println("The image size is valid " + path)
	}

	s := strings.Split(path, "/")
	bimg.Write(out+s[len(s)-1], newImage)
}
