package main

import (
	"driver-service/internal/storage"
	"driver-service/pb"
	"driver-service/service"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

func main() {
	db, err := storage.OpenSql(os.Getenv("driver_name"), os.Getenv("postgres_url"))
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}
	defer db.Close()

	
	lis, err := net.Listen("tcp", os.Getenv("server_url"))
	if err != nil {
		log.Fatal("Unable to listen :", err)
	}
	defer lis.Close()
	

	serv := service.NewDriverServer(db)
	s := grpc.NewServer()
	pb.RegisterDriverServiceServer(s, serv)

	log.Println("Server is listening on port ", os.Getenv("server_url"))
	if err = s.Serve(lis); err != nil {
		log.Fatal("Unable to serve :", err)
	}

}
