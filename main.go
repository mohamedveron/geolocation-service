package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/mohamedveron/geo-location_sdk/persistence"
	"github.com/mohamedveron/geo-location_sdk/service"
	"net/http"
)

func main() {

	dbSettings := persistence.DBSettings{
		Host:     "localhost",
		Port:     5432,
		Username: "postgres",
		Password: "123456",
		DbName:   "entities",
	}

	//initialize the service layers
	persistenceLayer := persistence.NewPersistence(&dbSettings)
	serviceLayer := service.NewService(persistenceLayer, "data_dump.csv")

	// start the importing process

	numberOfGoroutines := 10
	serviceLayer.RunDataIngestor(numberOfGoroutines)
	//server := api.NewServer(serviceLayer)

	// prepare router
	r := chi.NewRouter()

	srv := &http.Server{
		Handler: r,
		Addr:    ":8080",
	}

	fmt.Println("server starting on port 8080...")

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Println("server failed to start", "error", err)
	}

}
