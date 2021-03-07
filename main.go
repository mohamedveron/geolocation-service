package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/mohamedveron/geolocation-service/api"
	"github.com/mohamedveron/geolocation-service/persistence"
	"github.com/mohamedveron/geolocation-service/service"
	"net/http"
)

func main() {

	// db configuration
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

	//numberOfGoroutines := 10
	//serviceLayer.RunDataIngestor(numberOfGoroutines)
	server := api.NewServer(serviceLayer)

	// prepare router
	r := chi.NewRouter()
	r.Route("/", func(r chi.Router) {
		r.Mount("/", api.Handler(server))
	})

	srv := &http.Server{
		Handler: r,
		Addr:    ":9090",
	}

	fmt.Println("server starting on port 9090...")

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Println("server failed to start", "error", err)
	}

}
