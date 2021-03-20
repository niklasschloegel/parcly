package tracking

import (
	"fmt"
	"sort"

	"github.com/niklasschloegel/parcly/config"
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

func (c Carrier) Info() string {
	return fmt.Sprintf("[%s] - code: %s", c.Name, c.Code)
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

func GetCarriers() []Carrier {
	url := config.BasePath + "/carriers"
	carrierResponse := CarrierResponse{}

	err := DoRequest("GET", url, nil, &carrierResponse)
	if err != nil {
		panic(err)
	}

	carriers := carrierResponse.Carriers

	sort.Slice(carriers, func(i, j int) bool {
		return carriers[i].Name < carriers[j].Name
	})

	return carriers
}
