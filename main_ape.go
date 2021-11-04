package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/AvraamMavridis/randomcolor"
	"github.com/h2non/bimg"
	gim "github.com/ozankasikci/go-image-merge"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	_ "image/png"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

/*
for ok := true; ok; ok = extrasImagePath == "Bubblegum" {
					extrasImagePath = extras[GenerateRandomNumber(0, len(extras)-1)]
				}
*/

/*
	accessoriesPercent := GenerateRandomNumber(1, 100)
	hatsPercent := GenerateRandomNumber(1, 100)
	hairPercent := GenerateRandomNumber(1, 100)

	fmt.Println(accessoriesPercent)
	fmt.Println(hatsPercent)
	fmt.Println(hairPercent)
*/

/*
	1. Background
	2. Stroke
	3. Skin
	4. Clothes
	5. Hats
	6. Glasses
	7. Pacifiers
	8. Hairs
	9. BackObjects
	10. HandObjects
	11. HeadObjects
*/
/*
	rarityPercent := map[string]int{
		"common":   50,
		"uncommon": 28,
		"rare":     15,
		"mythical": 5,
		"legend":   2,
	}
*/

var layers = make([]ImageLayer, 0)
var src cryptoSource
var rnd *rand.Rand

var (
	colorB = [3]float64{248, 54, 0}
	colorA = [3]float64{254, 140, 0}
)

var (
	max = float64(0)
)

func linearGradient(x, y float64) (uint8, uint8, uint8) {
	d := x / max
	r := colorA[0] + d*(colorB[0]-colorA[0])
	g := colorA[1] + d*(colorB[1]-colorA[1])
	b := colorA[2] + d*(colorB[2]-colorA[2])
	return uint8(r), uint8(g), uint8(b)
}

func GenerateImage(filename string, width, height int) *image.RGBA {
	var r1 = randomcolor.GetRandomColorInRgb()
	var r2 = randomcolor.GetRandomColorInRgb()
	max = float64(width)

	colorA = [3]float64{float64(r1.Red), float64(r1.Green), float64(r1.Blue)}
	colorB = [3]float64{float64(r2.Red), float64(r2.Green), float64(r2.Blue)}

	var w, h int = width, height
	dst := image.NewRGBA(image.Rect(0, 0, w, h)) //*NRGBA (image.Image interface)

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			r, g, b := linearGradient(float64(x), float64(y))
			c := color.RGBA{

				r,
				g,
				b,
				255,
			}
			dst.Set(x, y, c)
		}
	}

	//img, _ := os.Create(filename)
	//defer img.Close()
	//png.Encode(img, dst) //Encode writes the Image m to w in PNG format.

	return dst
}

const (
	BASE_PATH        = "penguins_resouces"
	BACKGROUNDS_PATH = BASE_PATH + "/Backgrounds"
	CLOTHES_PATH     = BASE_PATH + "/Clothes"
	EYES_PATH        = BASE_PATH + "/Eyes"
	GLASSESS_PATH    = BASE_PATH + "/Glasses"
	HAIRS_PATH       = BASE_PATH + "/Hairs"
	HATS_PATH        = BASE_PATH + "/Hats"
	MOUTHS_PATH      = BASE_PATH + "/Mouths"
	NECKS_PATH       = BASE_PATH + "/Necks"
	SKINS_PATH       = BASE_PATH + "/Skins"
	OBJECTS_PATH     = BASE_PATH + "/Objects"
)

var COLLECTION_NAME = "PENGUINS"

func main() {
	// for generate random number
	rnd = rand.New(src)

	// folders
	generatedFolder := COLLECTION_NAME + "_GENERATED"
	originalFolder := "originalFolder"
	metadataFolder := "metadataFolder"
	resizedFolder := "resizedFolder"
	mergedFolder := "mergedFolder"

	fmt.Println(metadataFolder)

	// create generated folders structure
	CreateGeneratedCollectionFolder(generatedFolder, mergedFolder, originalFolder, metadataFolder, resizedFolder)

	var backgrounds []string
	var hairs []string
	var clothes []string
	var hats []string
	var eyes []string
	var skins []string
	var necks []string
	var objects []string
	var mouths []string
	var glasses []string

	var backgroundImagePath string
	var skinImagePath string
	var clotheImagePath string
	var hatImagePath string
	var hairImagePath string
	var eyesImagePath string
	var necksImagePath string
	var objectsImagePath string
	var mouthsImagePath string
	var glassesImagePath string

	COMMON_START_RANGE := 0
	COMMON_END_RANGE := 8888

	elements := make([]*HarmonyMetadata, 0)

	for i := COMMON_START_RANGE; i <= COMMON_END_RANGE; i++ {
		layers = make([]ImageLayer, 0)

		metadata := HarmonyMetadata{
		}

		clothesPercent := GenerateRandomNumber(1, 100)
		skinPercent := GenerateRandomNumber(1, 100)
		hairsPercent := GenerateRandomNumber(1, 100)
		hatsPercent := GenerateRandomNumber(1, 100)
		eyesRarityPercent := GenerateRandomNumber(1, 100)
		necksPercent := GenerateRandomNumber(1, 100)
		objectsPercent := GenerateRandomNumber(1, 100)
		mouthsPercent := GenerateRandomNumber(1, 100)
		glassesPercent := GenerateRandomNumber(1, 100)
		backgroundsPercent := GenerateRandomNumber(1, 100)

		clothesType := ""
		skinType := ""
		hairType := ""
		hatType := ""
		neckType := ""
		objectType := ""
		mouthType := ""
		eyesType := ""
		glassesType := ""
		backgroundType := ""

		if backgroundsPercent >= 0 && backgroundsPercent < 50 {
			backgroundType = "Common"
		} else if backgroundsPercent >= 50 && backgroundsPercent < 78 {
			backgroundType = "Uncommon"
		} else if backgroundsPercent >= 78 && backgroundsPercent < 93 {
			backgroundType = "Rare"
		} else if backgroundsPercent >= 93 && backgroundsPercent < 98 {
			backgroundType = "Mythical"
		} else {
			backgroundType = "Legend"
		}

		if glassesPercent >= 0 && glassesPercent < 50 {
			glassesType = "Common"
		} else if glassesPercent >= 50 && glassesPercent < 78 {
			glassesType = "Uncommon"
		} else if glassesPercent >= 78 && glassesPercent < 93 {
			glassesType = "Rare"
		} else if glassesPercent >= 93 && glassesPercent < 98 {
			glassesType = "Mythical"
		} else {
			glassesType = "Legend"
		}

		if eyesRarityPercent >= 0 && eyesRarityPercent < 50 {
			eyesType = "Common"
		} else if eyesRarityPercent >= 50 && eyesRarityPercent < 78 {
			eyesType = "Uncommon"
		} else if eyesRarityPercent >= 78 && eyesRarityPercent < 93 {
			eyesType = "Rare"
		} else if eyesRarityPercent >= 93 && eyesRarityPercent < 98 {
			eyesType = "Mythical"
		} else {
			eyesType = "Legend"
		}

		if mouthsPercent >= 0 && mouthsPercent < 50 {
			mouthType = "Common"
		} else if mouthsPercent >= 50 && mouthsPercent < 78 {
			mouthType = "Uncommon"
		} else if mouthsPercent >= 78 && mouthsPercent < 95 {
			mouthType = "Rare"
		} else {
			mouthType = "Mythical"
		}

		if objectsPercent >= 0 && objectsPercent < 33 {
			objectType = "Rare"
		} else if objectsPercent >= 33 && objectsPercent < 66 {
			objectType = "Mythical"
		} else {
			objectType = "Legend"
		}

		if skinPercent >= 0 && skinPercent < 50 {
			skinType = "Common"
		} else if skinPercent >= 50 && skinPercent < 78 {
			skinType = "Uncommon"
		} else if skinPercent >= 78 && skinPercent < 93 {
			skinType = "Rare"
		} else if skinPercent >= 93 && skinPercent < 98 {
			skinType = "Mythical"
		} else {
			skinType = "Legend"
		}

		if clothesPercent >= 0 && clothesPercent < 50 {
			clothesType = "Common"
		} else if clothesPercent >= 50 && clothesPercent < 78 {
			clothesType = "Uncommon"
		} else if clothesPercent >= 78 && clothesPercent < 93 {
			clothesType = "Rare"
		} else if clothesPercent >= 93 && clothesPercent < 98 {
			clothesType = "Mythical"
		} else {
			clothesType = "Legend"
		}

		if hairsPercent >= 0 && hairsPercent < 50 {
			hairType = "Uncommon"
		} else if hairsPercent >= 50 && hairsPercent < 85 {
			hairType = "Rare"
		} else {
			hairType = "Mythical"
		}

		if hatsPercent >= 0 && hatsPercent < 50 {
			hatType = "Common"
		} else if hatsPercent >= 50 && hatsPercent < 78 {
			hatType = "Uncommon"
		} else if hatsPercent >= 78 && hatsPercent < 93 {
			hatType = "Rare"
		} else if hatsPercent >= 93 && hatsPercent < 98 {
			hatType = "Mythical"
		} else {
			hatType = "Legend"
		}

		if necksPercent >= 0 && necksPercent < 50 {
			neckType = "Common"
		} else if necksPercent >= 50 && necksPercent < 78 {
			neckType = "Uncommon"
		} else if necksPercent >= 78 && necksPercent < 93 {
			neckType = "Rare"
		} else if necksPercent >= 93 && necksPercent < 98 {
			neckType = "Mythical"
		} else if necksPercent >= 98 && necksPercent <= 100 {
			neckType = "Legend"
		}

		backgrounds = GetFiles(BACKGROUNDS_PATH, backgroundType)
		skins = GetFiles(SKINS_PATH, skinType)
		eyes = GetFiles(EYES_PATH, eyesType)
		mouths = GetFiles(MOUTHS_PATH, mouthType)
		clothes = GetFiles(CLOTHES_PATH, clothesType)

		backgroundImagePath = backgrounds[GenerateRandomNumber(0, len(backgrounds)-1)]
		skinImagePath = skins[GenerateRandomNumber(0, len(skins)-1)]
		eyesImagePath = eyes[GenerateRandomNumber(0, len(eyes)-1)]
		mouthsImagePath = mouths[GenerateRandomNumber(0, len(mouths)-1)]
		clotheImagePath = clothes[GenerateRandomNumber(0, len(clothes)-1)]

		AddElement(backgroundImagePath)
		AddElement(skinImagePath)
		AddElement(eyesImagePath)
		AddElement(mouthsImagePath)
		AddElement(clotheImagePath)

		metadata.Background = GetObjectName(backgroundImagePath)
		metadata.Skin = GetObjectName(skinImagePath)
		metadata.Eye = GetObjectName(eyesImagePath)
		metadata.Mouth = GetObjectName(mouthsImagePath)
		metadata.Clothes = GetObjectName(clotheImagePath)

		hasHairOrHatPercent := GenerateRandomNumber(0, 100)
		if hasHairOrHatPercent >= 0 && hasHairOrHatPercent < 40 {

			gPercent := GenerateRandomNumber(1, 100)
			if gPercent < 50 {
				glasses = GetFiles(GLASSESS_PATH, glassesType)
				glassesImagePath = glasses[GenerateRandomNumber(0, len(glasses)-1)]
				AddElement(glassesImagePath)
				metadata.Glass = GetObjectName(glassesImagePath)
			}

			hats = GetFiles(HATS_PATH, hatType)
			hatImagePath = hats[GenerateRandomNumber(0, len(hats)-1)]
			metadata.Hat = GetObjectName(hatImagePath)

		} else if hasHairOrHatPercent >= 40 && hasHairOrHatPercent < 80 {

			hairs = GetFiles(HAIRS_PATH, hairType)
			hairImagePath = hairs[GenerateRandomNumber(0, len(hairs)-1)]
			AddElement(hairImagePath)
			metadata.Hair = GetObjectName(hairImagePath)

			gPercent := GenerateRandomNumber(1, 100)
			if gPercent < 50 {
				glasses = GetFiles(GLASSESS_PATH, glassesType)
				glassesImagePath = glasses[GenerateRandomNumber(0, len(glasses)-1)]
				AddElement(glassesImagePath)
				metadata.Glass = GetObjectName(glassesImagePath)
			}

		} else {
			gPercent := GenerateRandomNumber(1, 100)
			if gPercent < 50 {
				glasses = GetFiles(GLASSESS_PATH, glassesType)
				glassesImagePath = glasses[GenerateRandomNumber(0, len(glasses)-1)]
				AddElement(glassesImagePath)
				metadata.Glass = GetObjectName(glassesImagePath)
			}
		}

		nPercent := GenerateRandomNumber(0, 100)
		if nPercent < 60 {
			necks = GetFiles(NECKS_PATH, neckType)
			necksImagePath = necks[GenerateRandomNumber(0, len(necks)-1)]
			AddElement(necksImagePath)
			metadata.Neck = GetObjectName(necksImagePath)
		}

		oPercent := GenerateRandomNumber(0, 100)
		if oPercent < 60 {
			objects = GetFiles(OBJECTS_PATH, objectType)
			fmt.Println(GetObjectName(objectsImagePath))

			objectsImagePath = objects[GenerateRandomNumber(0, len(objects)-1)]

			if GetObjectName(hairImagePath) == "Green mohawk" || GetObjectName(hairImagePath) == "Pink mohawk" {

				for {
					objectsPercent = GenerateRandomNumber(1, 100)
					if objectsPercent >= 0 && objectsPercent < 33 {
						objectType = "Rare"
					} else if objectsPercent >= 33 && objectsPercent < 66 {
						objectType = "Mythical"
					} else {
						objectType = "Legend"
					}
					objects = GetFiles(OBJECTS_PATH, objectType)
					objectsImagePath = objects[GenerateRandomNumber(0, len(objects)-1)]
					if GetObjectName(objectsImagePath) != "Wings and aura" { // the condition stops matching
						break // break out of the loop
					}
				}

			}

			AddElement(objectsImagePath)
			metadata.RandomObject = GetObjectName(objectsImagePath)
		}

		res, err2 := GenerateBanner(layers, 512, 512)
		if err2 != nil {
			log.Printf("Error generating banner: %+v\n", err2)
		}

		counterString := strconv.Itoa(i)
		buff := new(bytes.Buffer)
		err := png.Encode(buff, res)
		if err != nil {
			fmt.Println("failed to create buffer", err)
		}

		options := bimg.Options{
			Width:     512,
			Height:    512,
			Crop:      true,
			Quality:   100,
			Rotate:    0,
			Interlace: true,
		}

		newImage, err3 := bimg.NewImage(buff.Bytes()).Process(options)
		if err3 != nil {
			fmt.Fprintln(os.Stderr, err3)
		}

		if bimg.NewImage(newImage).Type() != "png" {
			panic(errors.New("Erro exporting"))
		}

		r := bimg.Write("PENGUINS_GENERATED/originalFolder/"+counterString+".png", newImage)
		if r != nil {
			fmt.Println(r)
		}

		elements = append(elements, &metadata)

		log.Println(counterString)

		//meta, _ := json.Marshal(metadata)

		meta, _ := json.MarshalIndent(metadata, "", "    ")

		// to append to a file
		// create the file if it doesn't exists with O_CREATE, Set the file up for read write, add the append flag and set the permission
		f3, errr3 := os.OpenFile("./PENGUINS_GENERATED/metadataFolder/"+counterString+".json", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0660)
		defer f3.Close()
		if errr3 != nil {
			log.Fatal(errr3)
		}
		// write to file, f.Write()
		f3.Write(meta)

	}

	//j2, _ := json.Marshal(elements)
	j2, _ := json.MarshalIndent(elements, "", "    ")

	// to append to a file
	// create the file if it doesn't exists with O_CREATE, Set the file up for read write, add the append flag and set the permission
	f2, errr2 := os.OpenFile("./PENGUINS_GENERATED/metadataFolder/whole.json", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0660)
	defer f2.Close()
	if errr2 != nil {
		log.Fatal(errr2)
	}
	// write to file, f.Write()
	f2.Write(j2)

	/*
			// to append to a file
		// create the file if it doesn't exists with O_CREATE, Set the file up for read write, add the append flag and set the permission
		f, errr := os.OpenFile("./Babies/common/original/"+counterString+".json", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0660)
		defer f.Close()
		if errr != nil {
			log.Fatal(errr)
		}
		// write to file, f.Write()
		f.Write(j)
	*/

	ResizeImages(generatedFolder+"/"+originalFolder, generatedFolder+"/"+resizedFolder, COMMON_END_RANGE, 256)
	MergeImages(generatedFolder+"/"+resizedFolder, generatedFolder+"/"+mergedFolder, 1000, 9, 9, 81)
}

func AddAttributes(metadata *MetaplexMetadata, name string, value string) {
	metadata.Attributes = append(metadata.Attributes, Attribute{
		TraitType: name,
		Value:     GetObjectName(value),
	})
}

func AddElement(layer string) {
	img, e := openImage(layer)
	if e != nil {
		panic("e")
		panic(e)
	}

	layers = append(layers, ImageLayer{
		Image: img,
		XPos:  0,
		YPos:  0,
	})
}

func openImage(path string) (image.Image, error) {
	p := filepath.FromSlash(path)

	var file, err = os.OpenFile(p, os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	imageFile, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return imageFile, err
}

func GetFiles(root string, fileType string) []string {
	var files []string

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, ".png") {
			t := GetObjectRarity(path)
			if fileType == "All" {
				files = append(files, path)
			} else {
				if t == fileType {
					files = append(files, path)
				}
			}
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	return files
}

func GenerateRandomNumber(min, max int) int {
	ns := make([]int, 0)

	if max == 0 {
		return 0
	}

	for i := 0; i < 100; i++ {
		//fmt.Println("max")
		//fmt.Println(max)
		n := rnd.Intn(max)
		ns = append(ns, n)
	}
	n2 := rnd.Intn(len(ns))

	return ns[n2]
}

func ResizeImages(collectionName, generatedType string, amount int, width int) {
	for i := 0; i < amount; i++ {
		counterString := strconv.Itoa(i)
		ResizeImage("./"+collectionName+"/"+counterString+".png", "./PENGUINS_GENERATED/resizedFolder/", width)
	}
}

func MergeImages(collectionName, generatedType string, amount int, imageCountDX, imageCountDY int, divider int) {

	grids := make([]*gim.Grid, 0)

	for i := 0; i < amount; i++ {
		counterString := strconv.Itoa(i)
		currentFile := collectionName + "/" + counterString + ".png"

		g := &gim.Grid{
			ImageFilePath: currentFile,
		}
		grids = append(grids, g)

		fmt.Println(counterString)

		if i != 0 && i%divider == 0 {
			rgba, err := gim.New(grids, imageCountDX, imageCountDY).Merge()
			if err != nil {
				panic(err)
			}

			// save the output to jpg or png
			file, err2 := os.Create("./PENGUINS_GENERATED/mergedFolder/" + counterString + ".png")
			if err2 != nil {
				panic(err2)
			}
			err = jpeg.Encode(file, rgba, &jpeg.Options{Quality: 100})
			err = png.Encode(file, rgba)
			grids = []*gim.Grid{}
		}
	}
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

	size, err := bimg.NewImage(newImage).Size()
	if size.Width == width && size.Height == width {
		fmt.Println("The image size is valid " + path)
	}

	s := strings.Split(path, "/")
	bimg.Write(out+s[len(s)-1], newImage)
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

/*
	var colorInRGB = randomcolor.GetRandomColorInRgb()
	col := color.RGBA{uint8(colorInRGB.Red), uint8(colorInRGB.Green), uint8(colorInRGB.Blue), 0xaa}
	c := image.NewRGBA(image.Rect(0, 0, 1024, 1024))
	draw.Draw(c, c.Bounds(), image.NewUniform(col), image.ZP, draw.Src)

	layers = append(layers, ImageLayer{
		Image: c,
		XPos:  0,
		YPos:  0,
	})
*/

/*
	layers = append(layers, ImageLayer{
		Image: GenerateImage("t", 1024, 1024),
		XPos:  0,
		YPos:  0,
	})
*/

//AddElement(GetFiles("BabypunksOriginal/baseback", "All")[0])

func CreateGeneratedCollectionFolder(generatedFolder, mergedFolder, originalFolder, metadataFolder, resizedFolder string) {
	// clean folder
	// remove generated folder
	removeGeneratedFolderErr := os.Remove(generatedFolder)
	if removeGeneratedFolderErr != nil {
		fmt.Println(removeGeneratedFolderErr)
	}

	//Create a folder/directory at a full qualified path
	err := os.Mkdir(generatedFolder, 0755)
	if err != nil {
		log.Fatal(err)
	}

	//Create a folder/directory at a full qualified path
	err = os.Mkdir(generatedFolder+"/"+mergedFolder, 0755)
	if err != nil {
		log.Fatal(err)
	}

	//Create a folder/directory at a full qualified path
	err = os.Mkdir(generatedFolder+"/"+originalFolder, 0755)
	if err != nil {
		log.Fatal(err)
	}

	//Create a folder/directory at a full qualified path
	err = os.Mkdir(generatedFolder+"/"+metadataFolder, 0755)
	if err != nil {
		log.Fatal(err)
	}

	//Create a folder/directory at a full qualified path
	err = os.Mkdir(generatedFolder+"/"+resizedFolder, 0755)
	if err != nil {
		log.Fatal(err)
	}
}
