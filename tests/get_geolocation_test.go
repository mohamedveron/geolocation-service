package tests

import (
	"github.com/mohamedveron/geo-location_sdk/persistence"
	"github.com/mohamedveron/geo-location_sdk/service"
	"testing"
)

func TestGetAllNumbers(t *testing.T){
	persistenceLayer := persistence.NewPersistence("./sample.db")
	serviceLayer := service.NewService(persistenceLayer)
	_, err := serviceLayer.GetPhoneNumbers("", "")

	if err != nil {
		t.Error("error while getting phone numbers list")
	}
}
