package clients

import (
	"gateway-service/genpb/pbclient"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DialClient(clientService_url string) (pbclient.ClientServiceClient, error) {

	conn, err := grpc.NewClient(clientService_url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("Failed NewGrpc client to connect client-service")
		return nil, err
	}
	client := pbclient.NewClientServiceClient(conn)
	return client, nil
}
