package service

import (
	"fmt"
	"github.com/mohamedveron/geo-location_sdk/domains"
	"time"
)

func (s *Service) DataIngestor() (int, int, error) {

	locations := []*domains.GeoLocation{}
	countValid, countInValid := 0,0

	// pipeline of data ingestion
	importedData := s.ReadCsvData();

	startTime := time.Now()
	for idx, _ := range importedData {

		gl := &domains.GeoLocation{}

		if !s.CheckIPDuplicate(importedData[idx][0]) && s.Validator(importedData[idx], gl) {

			gl.IpAddress = importedData[idx][0]
			gl.CountryCode =  importedData[idx][1]
			gl.Country =      importedData[idx][2]
			gl.City =     importedData[idx][3]

			locations = append(locations, gl)

			s.ipMap[importedData[idx][0]] = true

			if len(locations) < 25000{
				err := s.persistence.CreateGeoLocation(gl)

				if err != nil {
					return countValid, countInValid, err
				}
			}

			countValid++
		}else{
			countInValid++
		}
	}

	//s.insertValidLocationsToDB(locations)

	fmt.Println(countInValid, " : ", countValid , time.Since(startTime) )

	return countValid, countInValid, nil
}

func (s *Service) insertValidLocationsToDB(locations []*domains.GeoLocation) error {

	for idx, _ := range locations {

		err := s.persistence.CreateGeoLocation(locations[idx])

		if err != nil {
			return err
		}

		if idx >1000 {
			break
		}
	}

	return nil
}
