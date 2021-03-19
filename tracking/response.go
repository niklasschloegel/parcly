package tracking

import "time"

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
