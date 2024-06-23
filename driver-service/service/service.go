package service

import (
	"context"
	"database/sql"
	"driver-service/internal/method"
	"driver-service/pb"

	"github.com/google/uuid"
)

type DriverServer struct {
	pb.UnimplementedDriverServiceServer
	DB *sql.DB
}

func NewDriverServer(db *sql.DB) *DriverServer {
	return &DriverServer{DB: db}
}

func (d *DriverServer) CreateDriver(ctx context.Context, req *pb.DRequest) (*pb.DResponse, error) {
	req.Id = uuid.New().String()
	req.JoinDate = method.GetDateWithTime()
	req.Active="active"

	method.StoreNewDriverToDatabase(d.DB, req)

	return &pb.DResponse{Id: req.Id}, nil
}

func (d *DriverServer) GetActiveDriverByLocation(ctx context.Context, req *pb.LocationRequest)(*pb.LocationResponse, error){

	resp, err :=method.GetActiveDriver(d.DB, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

