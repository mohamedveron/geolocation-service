package persistence

import (
	"github.com/mohamedveron/geolocation-service/domains"
)

func (db *Persistence) GetPhoneNumbers() ([]domains.GeoLocation, error) {
	phonesList := []domains.GeoLocation{}

	rows, err := db.database.Query("SELECT name, phone FROM customer")

	if err != nil {
		return phonesList, err
	}

	var name string
	var phone string

	for rows.Next() {
		p := domains.GeoLocation{}
		rows.Scan(&name, &phone)

		phonesList = append(phonesList, p)
	}

	return phonesList, nil
}
