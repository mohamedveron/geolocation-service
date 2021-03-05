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

	location, err := s.svc.GetLocations(ipAddress)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(location)
}
