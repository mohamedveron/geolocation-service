package service

import (
	"encoding/csv"
	"github.com/mohamedveron/geo-location_sdk/domains"
	"log"
	"os"
	"regexp"
	"strconv"
)

func (s *Service) ReadCsvData() [][]string {

	f, err := os.Open(s.filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+"data_dump.csv", err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+"data_dump.csv", err)
	}

	return records
}

func (s *Service) Validator(row []string, gl *domains.GeoLocation) bool {

	valid := true

	ipFlag, err := regexp.MatchString("\\b(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\b",
		row[0])

	codeFlag, err := regexp.MatchString("^[a-zA-Z]+$", row[1])
	countryFlag, err := regexp.MatchString("^[a-zA-Z]+$", row[2])
	cityFlag, err := regexp.MatchString("^[a-zA-Z]+$", row[3])

	if !ipFlag || !codeFlag || !countryFlag || !cityFlag {
		//fmt.Println("note valid string value")
		valid = false
	}

	gl.Latitude, err = strconv.ParseFloat(row[4], 64)
	gl.Longitude, err = strconv.ParseFloat(row[5], 64)
	if err != nil {
		//fmt.Println("wrong value for latitude or longitude" + err.Error())
		valid = false
	}

	if gl.MysteryValue, err = strconv.ParseInt(row[6], 10, 64); err != nil {
		//fmt.Println("wrong value for mystery value" + err.Error())
		valid = false
	}

	return valid
}

func (s *Service) CheckIPDuplicate(ip string) bool {

	if _, ok := s.ipMap[ip]; ok {
		return true
	}

	return false
}
