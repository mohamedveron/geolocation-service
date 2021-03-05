package service

import "github.com/mohamedveron/geolocation-service/domains"

func (s *Service) GetLocations(ip string) ([]domains.GeoLocation, error){

	locations, err := s.persistence.GetLocationsByIP(ip)

	if err != nil {
		return nil, err
	}

	return locations, err
}
