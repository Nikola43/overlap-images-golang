package main

import "encoding/json"

func (r *EthOpenseaMetadata) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func (r *EthOpenseaMetadata) ToJson() string {
	b, _ := json.Marshal(r)
	return string(b)
}

type EthOpenseaMetadata struct {
	Name                 string      `json:"name"`
	Symbol               string      `json:"symbol"`
	Edition              string      `json:"edition"`
	Description          string      `json:"description"`
	SellerFeeBasisPoints int64       `json:"seller_fee_basis_points"`
	Image                string      `json:"image"`
	ExternalURL          string      `json:"external_url"`
	Attributes           []Attribute `json:"attributes"`
	Properties           Properties  `json:"properties"`
}
