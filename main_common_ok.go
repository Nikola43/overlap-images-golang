package main

import (
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	_ "image/png"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/h2non/bimg"
	gim "github.com/ozankasikci/go-image-merge"
)

const (
	BASE_PATH              = "WOMAN"
	ARM_ACCENT_PATH        = BASE_PATH + "/ARM_ACCENT"
	DRESS_UPPER_PATH       = BASE_PATH + "/DRESS_UPPER"
	HEAD_PATH              = BASE_PATH + "/HEAD"
	LEGS_PATH              = BASE_PATH + "/LEGS"
	LOWER_BODY_ACCENT_PATH = BASE_PATH + "/LOWER_BODY_ACCENT"
	NECK_PATH              = BASE_PATH + "/NECK"
	SHOES_PATH             = BASE_PATH + "/SHOES"
	WEAPON_PATH            = BASE_PATH + "/WEAPON"
	WINGS_PATH             = BASE_PATH + "/WINGS"
)

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

func main() {

	rnd = rand.New(src)

	armAccents := GetFiles(ARM_ACCENT_PATH, "All")
	dressess := GetFiles(DRESS_UPPER_PATH, "All")
	heads := GetFiles(HEAD_PATH, "All")
	legs := GetFiles(LEGS_PATH, "All")
	lowerBodys := GetFiles(LOWER_BODY_ACCENT_PATH, "All")
	necks := GetFiles(NECK_PATH, "All")
	skins := GetFiles(SHOES_PATH, "All")
	weapons := GetFiles(WEAPON_PATH, "All")
	wings := GetFiles(WINGS_PATH, "All")

	COMMON_START_RANGE := 0
	COMMON_END_RANGE := 1000
	/*
		UNCOMMON_START_RANGE := 0
		UNCOMMON_END_RANGE := 4444

		RARE_START_RANGE := 0
		RARE_END_RANGE := 4444

		MYTHICAL_START_RANGE := 0
		MYTHICAL_END_RANGE := 4444

		LEGEND_START_RANGE := 0
		LEGEND_END_RANGE := 4444
	*/

	elements := make([]*MetaplexMetadata, 0)

	for i := COMMON_START_RANGE; i < COMMON_END_RANGE; i++ {
		layers = make([]ImageLayer, 0)

		baseWomanImagePath := "./WOMAN/WOMAN_BASE.png"
		armAccentsImagePath := armAccents[GenerateRandomNumber(0, len(armAccents)-1)]
		dressessImagePath := dressess[GenerateRandomNumber(0, len(dressess)-1)]
		headsImagePath := heads[GenerateRandomNumber(0, len(heads)-1)]
		legsImagePath := legs[GenerateRandomNumber(0, len(legs)-1)]
		lowerBodysImagePath := lowerBodys[GenerateRandomNumber(0, len(lowerBodys)-1)]
		necksImagePath := necks[GenerateRandomNumber(0, len(necks)-1)]
		skinsImagePath := skins[GenerateRandomNumber(0, len(skins)-1)]
		weaponsImagePath := weapons[GenerateRandomNumber(0, len(weapons)-1)]
		wingsImagePath := wings[GenerateRandomNumber(0, len(wings)-1)]

		AddElement(wingsImagePath)
		AddElement(weaponsImagePath)
		AddElement(baseWomanImagePath)
		AddElement(dressessImagePath)
		AddElement(skinsImagePath)
		AddElement(armAccentsImagePath)
		AddElement(headsImagePath)
		AddElement(legsImagePath)
		AddElement(lowerBodysImagePath)
		AddElement(necksImagePath)

		res, err := GenerateBanner(layers, 2480, 3508)
		if err != nil {
			log.Printf("Error generating banner: %+v\n", err)
		}

		counterString := strconv.Itoa(i)
		f, _ := os.Create("./womans/common/original/" + counterString + ".png")
		png.Encode(f, res)

		metadata := new(MetaplexMetadata)
		metadata.Name = "Woman #" + counterString
		metadata.Symbol = "Woman"
		metadata.Description = "Woman"
		metadata.SellerFeeBasisPoints = 500
		metadata.Image = "./Babies/common/original/" + counterString + ".png"
		metadata.ExternalURL = "https://nft.woman.com"
		metadata.Edition = "2021"

		AddAttributes(metadata, "ARM_ACCENT", armAccentsImagePath)
		AddAttributes(metadata, "ARM_ACCENT", armAccentsImagePath)
		AddAttributes(metadata, "DRESS_UPPER", dressessImagePath)
		AddAttributes(metadata, "HEAD", headsImagePath)
		AddAttributes(metadata, "LEGS", legsImagePath)
		AddAttributes(metadata, "LOWER", lowerBodysImagePath)
		AddAttributes(metadata, "NECK", necksImagePath)
		AddAttributes(metadata, "SHOES", skinsImagePath)
		AddAttributes(metadata, "WEAPON", weaponsImagePath)
		AddAttributes(metadata, "WINGS", wingsImagePath)

		metadata.Properties.Category = "image"
		metadata.Properties.Creators = append(metadata.Properties.Creators, Creator{
			Address: "B4s4rHrjWDTxVDZv7s9SZcstuYvRT2dtLnicAkfkHds1",
			Share:   100,
		})
		metadata.Properties.Files = append(metadata.Properties.Files, File{
			URI:  "./womans/common/original/" + counterString + ".png",
			Type: "image",
		})

		elements = append(elements, metadata)

		//fmt.Println()

		log.Println(counterString)
	}

	j2, _ := json.Marshal(elements)

	// to append to a file
	// create the file if it doesn't exists with O_CREATE, Set the file up for read write, add the append flag and set the permission
	f2, errr2 := os.OpenFile("./womans/common/original/whole.json", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0660)
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

func GetObjectName(object string) string {
	s := strings.Split(object, "/")
	objectName := s[len(s)-1]
	objectName = strings.Replace(objectName, ".png", "", 4)
	return objectName
}

func GetObjectRarity(object string) string {

	if strings.Contains(object, "Uncommon") {
		return "Uncommon"
	}

	if strings.Contains(object, "Common") {
		return "Common"
	}

	if strings.Contains(object, "Rare") {
		return "Rare"
	}

	if strings.Contains(object, "Mythical") {
		return "Mythical"
	}

	if strings.Contains(object, "Legend") {
		return "Legend"
	}

	return ""
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
	/*
		for _, file := range files {
			fmt.Println(file)
		}
	*/
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
	//for i := 4445; i < 6933; i++ {
	for i := 1; i <= amount; i++ {
		counterString := strconv.Itoa(i)
		ResizeImage("./"+collectionName+"/"+generatedType+"/original/"+counterString+".png", "./"+collectionName+"/"+generatedType+"/resized/", width)
	}
}

func MergeImages(collectionName, generatedType string, amount int, imageCountDX, imageCountDY int, divider int) {

	grids := make([]*gim.Grid, 0)

	for i := 1; i < amount; i++ {
		counterString := strconv.Itoa(i)
		currentFile := collectionName + "/" + generatedType + "/resized/" + counterString + ".png"

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
			file, err2 := os.Create(collectionName + "/" + generatedType + "/merged/" + counterString + ".png")
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
