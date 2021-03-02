package service

import "github.com/mohamedveron/geo-location_sdk/persistence"

type Service struct {
	persistence      *persistence.Persistence
	ipMap            map[string]bool
	filePath		 string
}

func NewService(persistence *persistence.Persistence, filePath string) *Service {

	dict := make(map[string]bool)

	return &Service{
		persistence:      persistence,
		ipMap: dict,
		filePath: filePath,
	}
}
