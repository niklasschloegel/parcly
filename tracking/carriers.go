package tracking

import (
	"github.com/niklasschloegel/parcly/cmd"
)

type CarrierResponse struct {
	Meta struct {
		Code int `json:"code"`
	} `json:"meta"`
	Carriers []Carrier `json:"data"`
}

type Carrier struct {
	Name     string `json:"name"`
	Code     string `json:"code"`
	Phone    string `json:"phone"`
	Homepage string `json:"homepage"`
	Type     string `json:"type"`
	Picture  string `json:"picture"`
}

type CarrierDetection struct {
	Meta struct {
		Code    int    `json:"code"`
		Type    string `json:"type"`
		Message string `json:"message"`
	} `json:"meta"`
	Data []DetectionData `json:"data"`
}

type DetectionData struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

func GetCarriers() []CarrierResponse {
	url := cmd.BasePath + "/carriers"
	carrierResponses := []CarrierResponse{}

	err := DoRequest("GET", url, nil, &carrierResponses)
	if err != nil {
		panic(err)
	}

	return carrierResponses
}
