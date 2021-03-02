package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/mohamedveron/geo-location_sdk/api"
	"github.com/mohamedveron/geo-location_sdk/persistence"
	"github.com/mohamedveron/geo-location_sdk/service"
	"net/http"
)

func main() {

	/*f, err := os.Open("data_dump.csv")
	if err != nil {
		log.Fatal("Unable to read input file " + "data_dump.csv", err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for " + "data_dump.csv", err)
	}

	locations := []*domains.GeoLocation{}
	ipMap := make(map[string]bool)

	countValid, countInValid := 0,0

	for idx, _ := range records {

		gl := &domains.GeoLocation{}

		valid := true

		ipFlag , err := regexp.MatchString("\\b(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\b", records[idx][0])
		codeFlag, err := regexp.MatchString("^[a-zA-Z]+$", records[idx][1])
		countryFlag, err := regexp.MatchString("^[a-zA-Z]+$", records[idx][2])
		cityFlag, err := regexp.MatchString("^[a-zA-Z]+$", records[idx][3])

		if !ipFlag || !codeFlag || !countryFlag || !cityFlag{
			//fmt.Println("here str", codeFlag, countryFlag, cityFlag)
			valid = false
		}
		if gl.Latitude, err = strconv.ParseFloat(records[idx][4], 64); err != nil {
			valid = false
		}
		if gl.Longitude, err = strconv.ParseFloat(records[idx][5], 64); err != nil {
			valid = false
		}

		if gl.MysteryValue, err = strconv.ParseInt(records[idx][6], 10,64); err != nil {
			valid = false
		}

		if _, ok := ipMap[records[idx][0]]; !ok && valid {

			gl.IpAddress = records[idx][0]
			gl.CountryCode =  records[idx][1]
			gl.Country =      records[idx][2]
			gl.City =     records[idx][3]

			locations = append(locations, gl)

			ipMap[records[idx][0]] = true

			countValid++
		}else{
			countInValid++
		}

	}
	fmt.Println(countInValid, " : ", countValid , " :")*/

	dbSettings := persistence.DBSettings{
		Host:     "geolocation.czqumefsqwp6.eu-central-1.rds.amazonaws.com",
		Port:     5432,
		Username: "postgres",
		Password: "123456789",
		DbName:   "postgres",
	}

	//initialize the service layers
	persistenceLayer := persistence.NewPersistence(&dbSettings)
	serviceLayer := service.NewService(persistenceLayer, "data_dump.csv")

	serviceLayer.DataIngestor()
	server := api.NewServer(serviceLayer)

	// prepare router
	r := chi.NewRouter()
	r.Route("/", func(r chi.Router) {
		r.Mount("/", api.Handler(server))
	})

	srv := &http.Server{
		Handler: r,
		Addr:    ":8080",
	}

	fmt.Println("server starting on port 8080...")

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Println("server failed to start", "error", err)
	}

}
