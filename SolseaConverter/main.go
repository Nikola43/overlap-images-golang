package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type SolseaMergedMetadata []SolseaMergedMetadatum

func UnmarshalSolseaMergedMetadata(data []byte) (SolseaMergedMetadata, error) {
	var r SolseaMergedMetadata
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SolseaMergedMetadata) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type SolseaMergedMetadatum struct {
	Image      string      `json:"image"`
	Attributes []Attribute `json:"attributes"`
	Name       string      `json:"name"`
	CustomID   string      `json:"custom_id"`
}


func main() {

	// Open our jsonFile
	filePath := "/Users/lie/go/src/github.com/nikola43/overlay_images/SolseaConverter/CACHE_METADATA.json"



	data := &UploadMetadataSolana{}

	buffer, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(buffer, &data)
	if err != nil {
		panic(err)
	}

	for key, value := range data.Items {
		fmt.Println("Key:", key, "Value:", value)

		url := value.Link
		method := "GET"

		client := &http.Client {
		}
		req, err := http.NewRequest(method, url, nil)

		if err != nil {
			fmt.Println(err)
			return
		}
		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		//fmt.Println(string(body))

		meta := new(MetaplexMetadata)

		err = json.Unmarshal(body, &meta)
		if err != nil {
			panic(err)
		}
		fmt.Println(meta)

		time.Sleep(5 * time.Second)
	}



	os.Exit(0)
}