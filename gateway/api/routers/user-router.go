package routers

import (
	"gateway-service/api/handlers"
	"gateway-service/grpc-client/clients"
	"gateway-service/grpc-client/drivers"
	"log"
	"net/http"
	"os"
)

func ClientRoutes()*http.ServeMux {
	mux :=http.NewServeMux()

	client, err :=clients.DialClient(os.Getenv("client_service_url"))
	if err != nil {
		log.Println("Failed to dial to client-service:",err)
	}

	driver, err :=drivers.DialDriver(os.Getenv("driver_service_url"))
	if err != nil {
		log.Println("failed to dial to driver-service:",err)
	}

	clientservice :=handlers.NewClientService(client)
	driverservice :=handlers.NewDriverService(driver)

	mux.HandleFunc("POST /register", clientservice.RegisterNewClient)
	mux.HandleFunc("POST /login", clientservice.LoginClientToapplication)

	mux.HandleFunc("POST /driver-add", driverservice.AddNewDriver)
	
	return mux
}
