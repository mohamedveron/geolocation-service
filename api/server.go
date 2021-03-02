package api

import "github.com/mohamedveron/geo-location_sdk/service"

type Server struct {
	svc *service.Service
}

func NewServer(svc *service.Service) *Server {
	return &Server{
		svc: svc,
	}
}
