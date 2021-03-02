package api

import (
	"net/http"
)

func (s *Server) GetPhoneNumbers(w http.ResponseWriter, r *http.Request) {

	/*var filters Filters
	if err := json.NewDecoder(r.Body).Decode(&filters); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// add filters
	country := ""
	if filters.Country != nil{
		country = *filters.Country
	}

	state := ""
	if filters.State != nil{
		state = *filters.State
	}
	phones, err := s.svc.GetPhoneNumbers(country, state)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(phones)*/
}
