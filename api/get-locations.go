package api

import (
	"encoding/json"
	"net/http"
)

func (s *Server) GetLocations(w http.ResponseWriter, r *http.Request) {

	var filters GeoLocationRequestData
	if err := json.NewDecoder(r.Body).Decode(&filters); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// add filters
	ipAddress := ""
	if filters.IpAddress != nil{
		ipAddress = *filters.IpAddress
	}

	locations, err := s.svc.GetLocations(ipAddress)

	responseList := GeoLocationResponseData{}
	res := []GeoLocation{}

	for idx, _ := range locations {

		loc := GeoLocation{
			City:        &locations[idx].City,
			Country:     &locations[idx].Country,
			CountryCode: &locations[idx].CountryCode,
			IpAddress:   &locations[idx].City,
			Latitude:    nil,
			Longitude:   nil,
		}

		res = append(res, loc)
	}

	responseList.Locations = res

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseList)
}
