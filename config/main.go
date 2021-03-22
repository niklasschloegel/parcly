package config

const (
	BasePath = "https://api.tracktry.com/v1"
)

var TrackingStatuses = []string{"pending", "notfound", "transit", "pickup",
	"delivered", "undelivered", "exception", "expired"}
var TracktryApiKey string
