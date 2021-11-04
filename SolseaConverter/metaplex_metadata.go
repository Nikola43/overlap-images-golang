package main

import "encoding/json"


func UnmarshalUploadMetadataSolana(data []byte) (UploadMetadataSolana, error) {
	var r UploadMetadataSolana
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *UploadMetadataSolana) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type UploadMetadataSolana struct {
	Program Program         `json:"program"`
	Items   map[string]Item `json:"items"`
}

type Item struct {
	Link    string `json:"link"`
	Name    string `json:"name"`
	OnChain bool   `json:"onChain"`
}

type Program struct {
	UUID   string `json:"uuid"`
	Config string `json:"config"`
}



func UnmarshalConfirmCustomerSaleOrderRequest(data []byte) (MetaplexMetadata, error) {
	var r MetaplexMetadata
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MetaplexMetadata) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func (r *MetaplexMetadata) ToJson() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type MetaplexMetadata struct {
	Name                 string      `json:"name"`
	Symbol               string      `json:"symbol"`
	Edition              string      `json:"edition"`
	Description          string      `json:"description"`
	SellerFeeBasisPoints int64       `json:"seller_fee_basis_points"`
	Image                string      `json:"image"`
	ExternalURL          string      `json:"external_url"`
	Collection           Collection  `json:"collection"`
	Attributes           []Attribute `json:"attributes"`
	Properties           Properties  `json:"properties"`
}

type Collection struct {
	Name   string `json:"name"`
	Family string `json:"family"`
}

type Attribute struct {
	TraitType string `json:"trait_type"`
	Value     string `json:"value"`
}

type Properties struct {
	Files    []File    `json:"files"`
	Category string    `json:"category"`
	Creators []Creator `json:"creators"`
}

type Creator struct {
	Address  string `json:"address"`
	Verified bool   `json:"verified"`
	Share    int64  `json:"share"`
}

type File struct {
	URI  string `json:"uri"`
	Type string `json:"type"`
}
