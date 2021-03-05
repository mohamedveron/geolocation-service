package api

import "github.com/mohamedveron/geolocation-service/service"

type Server struct {
	svc *service.Service
}

func NewServer(svc *service.Service) *Server {
	return &Server{
		svc: svc,
	}
}
