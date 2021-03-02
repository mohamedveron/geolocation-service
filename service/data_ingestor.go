package service

import (
	"fmt"
	"github.com/mohamedveron/geo-location_sdk/domains"
)

func (s *Service) DataIngestor() (int, int, error) {

	locations := []*domains.GeoLocation{}
	countValid, countInValid := 0,0

	// pipeline of data ingestion
	importedData := s.ReadCsvData();

	for idx, _ := range importedData {

		gl := &domains.GeoLocation{}

		if !s.CheckIPDuplicate(importedData[idx][0]) && s.Validator(importedData[idx], gl) {

			gl.IpAddress = importedData[idx][0]
			gl.CountryCode =  importedData[idx][1]
			gl.Country =      importedData[idx][2]
			gl.City =     importedData[idx][3]

			locations = append(locations, gl)

			s.ipMap[importedData[idx][0]] = true

			countValid++
		}else{
			countInValid++
		}
	}

	fmt.Println(countInValid, " : ", countValid , " :")

	return countValid, countInValid, nil
}
