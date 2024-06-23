package service

import (
	"client-service/internal/method"
	"client-service/pb"
	"context"
	"database/sql"
	"log"

	"github.com/google/uuid"
)

type Server struct {
	pb.UnimplementedClientServiceServer
	DB *sql.DB
}

func NewServer(db *sql.DB) *Server {
	return &Server{DB: db}
}

func (s *Server) CreateClient(ctx context.Context, req *pb.ClRequest) (*pb.ClResponse, error) {
	req.Id = uuid.New().String()
	req.Time=method.GetDateWithTime()

	if err := method.StoreNewClient(s.DB, req); err != nil {
		log.Println("Unable to store New client:",err)
		return nil, err
	}

	return &pb.ClResponse{Id: req.Id}, nil
}

func (s *Server) UpdateClient(ctx context.Context, req *pb.ClRequest)(*pb.Response, error){

	if err :=method.UpdateClientInDatabase(s.DB, req); err != nil {
		log.Println("Unable to update client:",err)
		return nil, err
	}
	return &pb.Response{Message: "Updated succesfully"}, nil
}

func (s *Server) DeleteClient(ctx context.Context, req *pb.ClResponse)(*pb.Response, error){

	if err :=method.DeleteClientFromDatabase(s.DB, req); err != nil {
		log.Println("failed to delete client from database:",err)
		return nil, err
	}
	return &pb.Response{Message: "Deleted Succesfully"}, nil
}

func (s *Server) GetClientById(ctx context.Context, req *pb.ClResponse)(*pb.ClRequest, error){

	return method.GetClientFromDatabse(s.DB, req), nil
}

func (s *Server) GetClientLocation(ctx context.Context, req *pb.ClResponse)(*pb.ClAddress,error){
	
	return method.GetClientAddressFromDatabase(s.DB, req), nil
}

