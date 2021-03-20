package tracking

import (
	"time"

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
type TrackingResponse struct {
	Meta struct {
		Code    int    `json:"code"`
		Type    string `json:"type"`
		Message string `json:"message"`
	} `json:"meta"`
	Data struct {
		Id              string     `json:"id"`
		TrackingNumber  string     `json:"tracking_number"`
		CarrierCode     string     `json:"carrier_code"`
		Status          string     `json:"status"`
		CreatedAt       time.Time  `json:"created_at"`
		UpdatedAt       time.Time  `json:"updated_at"`
		OrderCreateTime time.Time  `json:"order_create_time"`
		CustomerEmail   []string   `json:"customer_email"`
		Title           string     `json:"title"`
		OrderID         string     `json:"order_id"`
		Comment         string     `json:"comment"`
		CustomerName    string     `json:"customer_name"`
		Archived        bool       `json:"archived"`
		OriginalCountry string     `json:"original_country"`
		ItemTimeLength  int        `json:"itemTimeLength"`
		StayTimeLength  int        `json:"stayTimeLength"`
		OriginInfo      OriginInfo `json:"origin_info"`
		ServiceCode     string     `json:"service_code"`
		LastEvent       string     `json:"lastEvent"`
		LastUpdateTime  time.Time  `json:"lastUpdateTime"`
	} `json:"data"`
}

type OriginInfo struct {
	ItemReceived       time.Time   `json:"ItemReceived"`
	ItemDispatched     time.Time   `json:"ItemDispatched"`
	DepartfromAirport  time.Time   `json:"DepartfromAirport"`
	ArrivalfromAbroad  time.Time   `json:"ArrivalfromAbroad"`
	CustomsClearance   time.Time   `json:"CustomsClearance"`
	DestinationArrived time.Time   `json:"DestinationArrived"`
	Weblink            string      `json:"weblink"`
	Phone              string      `json:"phone"`
	CarrierCode        string      `json:"carrier_code"`
	TrackInfo          []TrackInfo `json:"trackinfo"`
}

type TrackInfo struct {
	Date              time.Time `json:"Date"`
	StatusDescription string    `json:"StatusDescription"`
	Details           string    `json:"Details"`
	CheckpointStatus  string    `json:"checkpoint_status"`
}

//--------------FUNCTIONS-------------------

func CreateTracking(tracking TrackingCreation) TrackingResponse {
	url := config.BasePath + "/trackings/realtime"
	trackingResponse := TrackingResponse{}

	err := DoRequest("POST", url, tracking, &trackingResponse)
	if err != nil {
		panic(err)
	}

	return trackingResponse
}
