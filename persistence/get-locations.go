package persistence

import (
	"fmt"
	"github.com/mohamedveron/geolocation-service/domains"
)

func (db *Persistence) GetLocationsByIP() ([]domains.GeoLocation, error) {
	locationsList := []domains.GeoLocation{}

	rows, err := db.database.Query("SELECT ip_address, country, country_code, city  FROM geolocation")

	if err != nil {
		return locationsList, err
	}

	var ip string
	var country string
	var country_code string
	var city string

	for rows.Next() {
		p := domains.GeoLocation{}
		rows.Scan(&ip, &country, &country_code, &city)
		fmt.Print("here" + ip + country)
		p.IpAddress = ip
		p.Country = country
		p.CountryCode = country_code
		p.City = city
		locationsList = append(locationsList, p)
	}

	return locationsList, nil
}
