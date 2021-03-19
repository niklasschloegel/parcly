package tracking

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
