package service

import "github.com/mohamedveron/geolocation-service/domains"

func (s *Service) GetLocations() ([]domains.GeoLocation, error){

	locations, err := s.persistence.GetLocationsByIP()

	return locations, err
}
