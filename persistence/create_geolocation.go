package persistence

import (
	"fmt"
	"github.com/mohamedveron/geo-location_sdk/domains"
)

func (db Persistence) CreateGeoLocation(row *domains.GeoLocation) error {

	sqlStatement := `INSERT INTO geolocation (ip_address, country, country_code, city, latitude, longitude, mystery_value)
			VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err := db.database.Exec(sqlStatement, row.IpAddress, row.Country, row.CountryCode, row.City, row.Latitude, row.Longitude, row.MysteryValue)

	if err != nil {
		fmt.Print(err)
		panic(err)
	}

	return err
}
