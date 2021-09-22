package main

import (
	"encoding/json"
	"image"
	"image/png"
	_ "image/png"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const (
	BASE_PATH        = "BabypunksOriginal"
	BACKGROUNDS_PATH = BASE_PATH + "/Backgrounds"
	CLOTHES_PATH     = BASE_PATH + "/Clothes"
	DUMMIES_PATH     = BASE_PATH + "/Pacifiers"
	EYES_PATH        = BASE_PATH + "/Eyes"
	HAIRS_PATH       = BASE_PATH + "/Hairs"
	HATS_PATH        = BASE_PATH + "/Hats"
	HANDS_PATH       = BASE_PATH + "/Hands"
	STROKES_PATH     = BASE_PATH + "/Strokes"
	SKINS_PATH       = BASE_PATH + "/Skins"
	BACKS_PATH       = BASE_PATH + "/backs"
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

func main() {
	var backgrounds []string
	var pacifiers []string
	var hairs []string
	var clothes []string
	var hats []string
	var eyes []string
	var skins []string
	var strokes []string




	var backgroundImagePath string
	var strokeImagePath string
	var skinImagePath string
	var clotheImagePath string
	var hatImagePath string
	var hairImagePath string
	var pacifierImagePath string
	var eyesImagePath string




	/*

		// UNCOMMON
		if rarityPercent >= 50 && rarityPercent < 78 {
			// cojo ropa uncommon
			strokes = GetFiles(STROKES_PATH, "Common")
		}

		// RARE
		if rarityPercent >= 78 && rarityPercent < 93 {
			// cojo ropa rare
			strokes = GetFiles(STROKES_PATH, "Common")
		}

		// MYTHICAL
		if rarityPercent >= 94 && rarityPercent < 98 {
			// cojo ropa mythical
			strokes = GetFiles(STROKES_PATH, "Common")
		}

		// LEGEND
		if rarityPercent >= 98 && rarityPercent < 100 {
			// cojo ropa legend
			strokes = GetFiles(STROKES_PATH, "Common")
		}
	*/


	COMMON_START_RANGE := 0
	COMMON_END_RANGE := 10

	elements := make([]*MetaplexMetadata, 0)

	for i := COMMON_START_RANGE; i < COMMON_END_RANGE; i++ {

		metadata := new(MetaplexMetadata)
		metadata.Symbol = "BabyPunk"
		metadata.Description = "BabyPunk"
		metadata.SellerFeeBasisPoints = 500
		metadata.ExternalURL = "https://nft.babypunks.com"
		metadata.Edition = "2021"
		metadata.Properties.Category = "image"
		metadata.Properties.Creators = append(metadata.Properties.Creators, Creator{
			Address: "B4s4rHrjWDTxVDZv7s9SZcstuYvRT2dtLnicAkfkHds1",
			Share:   100,
		})


		rarityPercent := GenerateRandomNumber(1, 50)

		// COMMON
		if rarityPercent >= 1 && rarityPercent < 50 {
			backgrounds = GetFiles(BACKGROUNDS_PATH, "Common")
			strokes = GetFiles(STROKES_PATH, "Common")
			skins = GetFiles(SKINS_PATH, "Common")

			backgroundImagePath = backgrounds[GenerateRandomNumber(0, len(backgrounds)-1)]
			strokeImagePath = strokes[GenerateRandomNumber(0, len(strokes)-1)]
			skinImagePath = skins[GenerateRandomNumber(0, len(skins)-1)]

			AddAttributes(metadata, "Background", backgroundImagePath)
			AddAttributes(metadata, "Stroke", strokeImagePath)
			AddAttributes(metadata, "Skin", skinImagePath)


			clotheRarityPercent := GenerateRandomNumber(1, 100)
			pacifierRarityPercent := GenerateRandomNumber(1, 100)
			hatRarityPercent := GenerateRandomNumber(1, 100)
			hairRarityPercent := GenerateRandomNumber(1, 100)
			eyesRarityPercent := GenerateRandomNumber(1, 100)

			hairOrHatPercent := GenerateRandomNumber(1, 100)



			hatImagePath = hats[GenerateRandomNumber(0, len(hats)-1)]
			hairImagePath = hairs[GenerateRandomNumber(0, len(hairs)-1)]
			pacifierImagePath = pacifiers[GenerateRandomNumber(0, len(pacifiers)-1)]

			/*
				CLOTHES
			*/
			if clotheRarityPercent >= 1 && clotheRarityPercent < 50 {
				clothes = GetFiles(CLOTHES_PATH, "Common")
			} else if clotheRarityPercent >= 50 && clotheRarityPercent < 80 {
				clothes = GetFiles(CLOTHES_PATH, "Uncommon")
			} else {
				clothes = GetFiles(CLOTHES_PATH, "Rare")
			}
			clotheImagePath = clothes[GenerateRandomNumber(0, len(clothes)-1)]
			AddAttributes(metadata, "Clothes", clotheImagePath)
			/// -----------------------------------------------------------------------------

			/*
				PACIFIERS
			*/
			if pacifierRarityPercent > 1 && pacifierRarityPercent < 70 {
				pacifiers = GetFiles(DUMMIES_PATH, "Common")
			} else {
				pacifiers = GetFiles(DUMMIES_PATH, "Uncommon")
			}
			pacifierImagePath = pacifiers[GenerateRandomNumber(0, len(pacifiers)-1)]
			AddAttributes(metadata, "Pacifier", pacifierImagePath)
			/// -----------------------------------------------------------------------------

			/*
				HAIR_OR_HAT
			*/
			if hairOrHatPercent >= 1 && hatRarityPercent < 30 {
				/*
					HATS
				*/
				if hatRarityPercent >= 1 && hatRarityPercent < 50 {
					hats = GetFiles(HATS_PATH, "Common")
				} else if hatRarityPercent >= 50 && hatRarityPercent <= 80 {
					hats = GetFiles(HATS_PATH, "Uncommon")
				} else {
					hats = GetFiles(HATS_PATH, "Rare")
				}

				hatImagePath = hats[GenerateRandomNumber(0, len(hats)-1)]
				AddAttributes(metadata, "Hat", hatImagePath)

			} else if hairOrHatPercent >= 30 && hairOrHatPercent < 60 {
				/*
					HAIR
				*/
				if hairRarityPercent >= 1 && hairRarityPercent < 50 {
					hairs = GetFiles(HAIRS_PATH, "Common")
				} else if hairRarityPercent >= 50 && hairRarityPercent < 80 {
					hairs = GetFiles(HAIRS_PATH, "Uncommon")
				} else {
					hairs = GetFiles(HAIRS_PATH, "Rare")
				}
				hairImagePath = hairs[GenerateRandomNumber(0, len(hairs)-1)]
				AddAttributes(metadata, "Hair", hairImagePath)
			} else {

			}

			/*
				GLASSESS
			*/
			if eyesRarityPercent >= 1 && eyesRarityPercent < 50 {

				if eyesRarityPercent >= 1 && eyesRarityPercent < 70 {
					eyes = GetFiles(EYES_PATH, "Common")
				} else {
					eyes = GetFiles(EYES_PATH, "Uncommon")
				}
				eyesImagePath = eyes[GenerateRandomNumber(0, len(eyes)-1)]
				AddAttributes(metadata, "Eyes", eyesImagePath)
			} else {

			}
		}

		res, err := GenerateBanner(layers, 512, 512)
		if err != nil {
			log.Printf("Error generating banner: %+v\n", err)
		}

		counterString := strconv.Itoa(i)
		f, _ := os.Create("./Babies/common/original/" + counterString + ".png")
		png.Encode(f, res)

		metadata.Name = "BabyPunk #" + counterString
		metadata.Image = "./Babies/common/original/" + counterString + ".png"

		metadata.Properties.Files = append(metadata.Properties.Files, File{
			URI:  "./Babies/common/original/" + counterString + ".png",
			Type: "image",
		})

		elements = append(elements, metadata)


		log.Println(counterString)
	}

	j2, _ := json.Marshal(elements)

	// to append to a file
	// create the file if it doesn't exists with O_CREATE, Set the file up for read write, add the append flag and set the permission
	f2, errr2 := os.OpenFile("./Babies/common/original/whole.json", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0660)
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
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
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
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(max-min+1) + min
	//fmt.Println(n)

	return n
}
