package tracking

import (
	"fmt"

	"github.com/niklasschloegel/parcly/config"
)

//----------------STRUCTS-------------------
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
	Meta struct {
		Code    int    `json:"code"`
		Type    string `json:"type"`
		Message string `json:"message"`
	} `json:"meta"`
	Data struct {
		Items []TrackingData `json:"items"`
	} `json:"data"`
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

type DeletionResponse struct {
	Meta struct {
		Code    int    `json:"code"`
		Type    string `json:"type"`
		Message string `json:"message"`
	} `json:"meta"`
	Data []interface{} `json:"data"`
}

func (t TrackingData) ShortInfo() string {
	var title string
	if t.Comment != "" {
		title = t.Comment
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

//--------------FUNCTIONS-------------------

func CreateTracking(tracking TrackingCreation) TrackingData {
	url := config.BasePath + "/trackings/realtime"
	trackingResponse := TrackingItems{}

	err := DoRequest("POST", url, tracking, &trackingResponse)
	if err != nil {
		panic(err)
	}

	return trackingResponse.Data.Items[0]
}

func GetTrackings() []TrackingData {
	url := config.BasePath + "/trackings/get"
	trackingResponse := TrackingItems{}

	err := DoRequest("GET", url, nil, &trackingResponse)
	if err != nil {
		panic(err)
	}

	return trackingResponse.Data.Items
}

func DeleteTracking(carrierCode, trackingNumber string) {
	url := fmt.Sprintf("%s/trackings/%s/%s", config.BasePath, carrierCode, trackingNumber)
	deletionResponse := DeletionResponse{}

	err := DoRequest("DELETE", url, nil, &deletionResponse)
	if err != nil {
		panic(err)
	}
}
