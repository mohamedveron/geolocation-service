package service

import (
	"fmt"
	"github.com/mohamedveron/geolocation-service/domains"
	"golang.org/x/sync/errgroup"
	"time"
)

func (s *Service) RunDataIngestor(numberOfGoroutines int) (int, int, time.Duration, error) {

	countValid, countInValid := 0,0

	// pipeline of data ingestion
	importedData := s.ReadCsvData();

	// save locations to channel for concurrent access
	ch := make(chan domains.GeoLocation)

	errg := errgroup.Group{}

	for i := 0; i < numberOfGoroutines; i++ {
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

			s.ipMap[importedData[idx][0]] = true

			countValid++

			// add valid location to the channel
			ch <- *gl

		}else{
			countInValid++
		}
	}

	close(ch)

	err := errg.Wait()
	if err != nil {
		e := "Error while persist locations "
		return countValid, countInValid, time.Since(startTime), fmt.Errorf("%s: %s", e, err)
	}

	fmt.Println(countInValid, " : ", countValid , time.Since(startTime) )

	return countValid, countInValid, time.Since(startTime), nil
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
