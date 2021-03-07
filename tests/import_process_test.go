package tests

import (
	"fmt"
	"github.com/mohamedveron/geolocation-service/persistence"
	"github.com/mohamedveron/geolocation-service/service"
	"testing"
)

func TestImportGeolocationData(t *testing.T){

	dbSettings := persistence.DBSettings{
		Host:     "geolocation.czqumefsqwp6.eu-central-1.rds.amazonaws.com",
		Port:     5432,
		Username: "postgres",
		Password: "123456789",
		DbName:   "postgres",
	}

	//initialize the service layers
	persistenceLayer := persistence.NewPersistence(&dbSettings)
	serviceLayer := service.NewService(persistenceLayer, "../data_dump.csv")

	// start the importing process

	numberOfGoroutines := 10
	countInValid, countValid, duration, err := serviceLayer.RunDataIngestor(numberOfGoroutines)

	if err != nil {
		t.Errorf("Import geolocation data process failed")
	}

	fmt.Println("accepted entries :" , countInValid, " , discarded entries: ", countValid , " , time elapsed : " , duration )

}
