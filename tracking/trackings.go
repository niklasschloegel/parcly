/*
Copyright © 2021 Niklas Schlögel <niklasschloegel@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package tracking

import (
	"errors"
	"fmt"
	"os"
	"sort"

	"github.com/niklasschloegel/parcly/config"
)

//----------------STRUCTS-------------------
type AnyResponse struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

//------------Tracking Creation-------------
type TrackingCreation struct {
	TrackingNumber           string `json:"tracking_number"`
	CarrierCode              string `json:"carrier_code"`
	DestinationCode          string `json:"destination_code,omitempty"`
	TrackingShipDate         string `json:"tracking_ship_date,omitempty"`
	TrackingPostalCode       string `json:"tracking_postal_code,omitempty"`
	TrackingAccountNumber    string `json:"tracking_account_number,omitempty"`
	SpecialNumberDestination string `json:"specialNumberDestination,omitempty"`
	Order                    string `json:"order,omitempty"`
	OrderCreateTime          string `json:"order_create_time,omitempty"`
	Lang                     string `json:"lang,omitempty"`
	AutoCorrect              string `json:"auto_correct,omitempty"`
	Comment                  string `json:"comment,omitempty"`
}

//------------Tracking Response------------
type TrackingItems struct {
	Meta Meta `json:"meta"`
	Data struct {
		Items []TrackingData `json:"items"`
	} `json:"data"`
}

type Meta struct {
	Code    int    `json:"code"`
	Type    string `json:"type"`
	Message string `json:"message"`
}

type TrackingData struct {
	Id              string      `json:"id"`
	TrackingNumber  string      `json:"tracking_number"`
	CarrierCode     string      `json:"carrier_code"`
	Status          string      `json:"status"`
	CreatedAt       string      `json:"created_at"`
	UpdatedAt       string      `json:"updated_at"`
	OrderCreateTime string      `json:"order_create_time"`
	Title           string      `json:"title"`
	OrderID         string      `json:"order_id"`
	Comment         string      `json:"comment"`
	CustomerName    string      `json:"customer_name"`
	Archived        bool        `json:"archived"`
	OriginalCountry string      `json:"original_country"`
	ItemTimeLength  int         `json:"itemTimeLength"`
	StayTimeLength  int         `json:"stayTimeLength"`
	OriginInfo      CarrierInfo `json:"origin_info"`
	DestinationInfo CarrierInfo `json:"destination_info"`
	ServiceCode     string      `json:"service_code"`
	LastEvent       string      `json:"lastEvent"`
	LastUpdateTime  string      `json:"lastUpdateTime"`
}
type CarrierInfo struct {
	ItemReceived       string      `json:"ItemReceived"`
	ItemDispatched     string      `json:"ItemDispatched"`
	DepartfromAirport  string      `json:"DepartfromAirport"`
	ArrivalfromAbroad  string      `json:"ArrivalfromAbroad"`
	CustomsClearance   string      `json:"CustomsClearance"`
	DestinationArrived string      `json:"DestinationArrived"`
	Weblink            string      `json:"weblink"`
	Phone              string      `json:"phone"`
	CarrierCode        string      `json:"carrier_code"`
	TrackInfo          []TrackInfo `json:"trackinfo"`
}

type TrackInfo struct {
	Date              string `json:"Date"`
	StatusDescription string `json:"StatusDescription"`
	Details           string `json:"Details"`
	CheckpointStatus  string `json:"checkpoint_status"`
}

func (t TrackingData) ShortInfo() string {
	var title string
	if t.Title != "" {
		title = t.Title
	} else {
		title = "A parcel"
	}
	if t.OrderID != "" {
		title += fmt.Sprintf("(#%s)", t.OrderID)
	}
	return fmt.Sprintf("[%s]\t@ %s - %s from %s (%s)", t.Status, t.UpdatedAt, title, t.CarrierCode, t.TrackingNumber)
}

func (t TrackingData) Info() string {
	info := t.ShortInfo() + "\n\tHistory:\n"
	for _, trackInfo := range append(t.OriginInfo.TrackInfo, t.DestinationInfo.TrackInfo...) {
		info += trackInfo.Info() + "\n"
	}
	return info
}

func (t TrackInfo) Info() string {
	return fmt.Sprintf("\t[%s] @ %s\n\t- %s", t.CheckpointStatus, t.Date, t.StatusDescription)
}

//------------Tracking Deletion-------------
type DeletionResponse struct {
	Meta Meta          `json:"meta"`
	Data []interface{} `json:"data"`
}

//------------Tracking Update-------------
type UpdateTracking struct {
	Title            string `json:"title,omitempty"`
	LogisticsChannel string `json:"logistics_channel,omitempty"`
	CustomerName     string `json:"customer_name,omitempty"`
	CustomerEmail    string `json:"customer_email,omitempty"`
	CustomerPhone    string `json:"customer_phone,omitempty"`
	OrderId          string `json:"order_id,omitempty"`
	DestinationCode  string `json:"destination_code,omitempty"`
	Archived         string `json:"archived,omitempty"`
	Status           int    `json:"status,omitempty"`
}

type UpdateResponse struct {
	Meta Meta         `json:"meta"`
	Data TrackingData `json:"data"`
}

//--------------FUNCTIONS-------------------

func CreateTracking(tracking TrackingCreation) TrackingData {
	url := config.BasePath + "/trackings/realtime"
	trackingResponse := TrackingItems{}

	err := DoRequest("POST", url, tracking, &trackingResponse)
	if err != nil {
		errOut := fmt.Errorf("error: %v", err)
		fmt.Println(errOut.Error())
		os.Exit(-1)
	}

	if tracking.Comment != "" {
		// straight Update, because title cannot be set on first request
		updateTracking := UpdateTracking{Title: tracking.Comment}
		EditTrackings(tracking.CarrierCode, tracking.TrackingNumber, updateTracking)
	}

	return trackingResponse.Data.Items[0]
}

func GetTrackings() []TrackingData {
	url := config.BasePath + "/trackings/get"
	trackingResponse := TrackingItems{}

	err := DoRequest("GET", url, nil, &trackingResponse)
	if err != nil {
		errOut := fmt.Errorf("error: %v", err)
		fmt.Println(errOut.Error())
		os.Exit(-1)
	}

	return trackingResponse.Data.Items
}

func EditTrackings(carrierCode, trackingNumber string, update UpdateTracking) {
	url := fmt.Sprintf("%s/trackings/%s/%s", config.BasePath, carrierCode, trackingNumber)
	updateResponse := UpdateResponse{}

	err := DoRequest("PUT", url, update, &updateResponse)
	if err != nil {
		errOut := fmt.Errorf("error: %v", err)
		fmt.Println(errOut.Error())
		os.Exit(-1)
	}
}

func DeleteTracking(carrierCode, trackingNumber string) {
	url := fmt.Sprintf("%s/trackings/%s/%s", config.BasePath, carrierCode, trackingNumber)
	deletionResponse := DeletionResponse{}

	err := DoRequest("DELETE", url, nil, &deletionResponse)
	if err != nil {
		errOut := fmt.Errorf("error: %v", err)
		fmt.Println(errOut.Error())
		os.Exit(-1)
	}
}

func FilterStatus(pattern string, trackings []TrackingData) ([]TrackingData, error) {
	filteredTrackings := []TrackingData{}

	var possibleStatuses sort.StringSlice = config.TrackingStatuses
	possibleStatuses.Sort()
	index := possibleStatuses.Search(pattern)
	if possibleStatuses[index] != pattern {
		errorMessage := fmt.Sprintf("No status '%s' found", pattern)
		return filteredTrackings, errors.New(errorMessage)
	}

	for _, t := range trackings {
		if t.Status == pattern {
			filteredTrackings = append(filteredTrackings, t)
		}
	}
	return filteredTrackings, nil
}

func FilterCarrier(pattern string, trackings []TrackingData) ([]TrackingData, error) {
	filteredTrackings := []TrackingData{}

	possibleCarriers := GetCarriers()
	var carrierCodes sort.StringSlice = []string{}
	for _, c := range possibleCarriers {
		carrierCodes = append(carrierCodes, c.Code)
	}
	carrierCodes.Sort()
	index := carrierCodes.Search(pattern)
	if carrierCodes[index] != pattern {
		errorMessage := fmt.Sprintf("No carrier found with code '%s'", pattern)
		return filteredTrackings, errors.New(errorMessage)
	}

	for _, t := range trackings {
		if t.CarrierCode == pattern {
			filteredTrackings = append(filteredTrackings, t)
		}
	}

	return filteredTrackings, nil

}
