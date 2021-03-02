package domains

// GeoLocation type
type GeoLocation struct {
	IpAddress string `json:"ip_address"`
	CountryCode       string `json:"country_code" `
	Country       string `json:"country" `
	City       string `json:"city" `
	Latitude   float64  `json:"latitude"`
	Longitude   float64  `json:"latitude"`
	MysteryValue int64   `json:"mystery_value"`
}

