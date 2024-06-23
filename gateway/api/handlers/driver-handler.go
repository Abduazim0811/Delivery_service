package handlers

import (
	"gateway-service/genpb/pbdriver"
	"net/http"
)

type DriverService struct {
	Driver  pbdriver.DriverServiceClient
}

func NewDriverService(dr pbdriver.DriverServiceClient)*DriverService{
	return &DriverService{Driver: dr }
}

func (d *DriverService) AddNewDriver(w http.ResponseWriter, r *http.Request){
	
}