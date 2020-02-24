package app

import (
	"skoltech/pkg/devices"
)

//Server structure
type Server struct {
	devicesService devices.IService
}

//NewServer server constructor
func NewServer(ds devices.IService) *Server {
	return &Server{
		devicesService: ds,
	}
}
