package service

import (
	"fmt"
	"github.com/mohamedveron/geo-location_sdk/domains"
	"golang.org/x/sync/errgroup"
	"time"
)

func (s *Service) DataIngestor() (int, int, error) {

	locations := []*domains.GeoLocation{}
	countValid, countInValid := 0,0

	// pipeline of data ingestion
	importedData := s.ReadCsvData();

	ch := make(chan domains.GeoLocation)

	errg := errgroup.Group{}
	for i := 0; i < 10; i++ {
		errg.Go(func() error {
			return s.worker(ch)
		})
	}

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

			countValid++

			if len(locations) < 180000 {
				ch <- *gl
			}

		}else{
			countInValid++
		}
	}

	close(ch)

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

func (s *Service) worker(locations chan domains.GeoLocation) error {

	for location := range locations {

		err := s.persistence.CreateGeoLocation(&location)

		if err != nil {
			return err
		}
	}

	return nil
}
