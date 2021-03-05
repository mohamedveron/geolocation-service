package api

import (
	"net/http"
)

func (s *Server) GetLocations(w http.ResponseWriter, r *http.Request) {

	s.svc.GetLocations()
}
