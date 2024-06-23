package drivers

import (
	"gateway-service/genpb/pbdriver"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DialDriver(driverService_url string) (pbdriver.DriverServiceClient, error){
	conn, err := grpc.NewClient(driverService_url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("Failed to dial driverclient:",err)
		return nil, err
	}
	client := pbdriver.NewDriverServiceClient(conn)

	return client, nil
}
