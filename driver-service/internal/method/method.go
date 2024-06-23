package method

import (
	"database/sql"
	"driver-service/pb"
	"fmt"
	"log"
	"time"
)

func StoreNewDriverToDatabase(db *sql.DB, req *pb.DRequest) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatalf("Failed to begin transaction: %v", err)
	}
	defer tx.Rollback()

	query := `
		INSERT INTO drivers (id, name, email, phone, working_place, active)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err = tx.Exec(query, req.Id, req.Name, req.Email, req.Phone, req.WorkingPlace, req.Active)
	if err != nil {
		log.Fatalf("Failed to insert new driver to drivers table: %v", err)
	}

	query = `
		INSERT INTO locations (driver_id, city, street, home_number)
		VALUES ($1, $2, $3, $4)
	`
	_, err = tx.Exec(query, req.Id, req.Address.City, req.Address.Street, req.Address.HomeNumber)
	if err != nil {
		log.Fatalf("Failed to insert to locations table: %v", err)
	}

	query = `
		INSERT INTO statuses (driver_id, vehicle, join_date, rating)
		VALUES ($1, $2, $3, $4)
	`
	_, err = tx.Exec(query, req.Id, req.Status.Vehicle, req.JoinDate, req.Status.Rating)
	if err != nil {
		log.Fatalf("Failed to insert to statuses table: %v", err)
	}

	if err = tx.Commit(); err != nil {
		log.Fatalf("Failed to commit transaction: %v", err)
	}
}

func GetActiveDriver(db *sql.DB, req *pb.LocationRequest)(*pb.LocationResponse, error){
	query :=`
		select d.id, d.name, d.phone, s.vehicle from drivers d
		join statuses s on d.id = s.driver_id
		where d.working_place = $1 and d.active = $2
	`	
	var id, name, phone, vehicle string
	row :=db.QueryRow(query, req.Region, "active")
	if err :=row.Scan(&id, &name, &phone, &vehicle); err != nil {
		if err ==sql.ErrNoRows{
			return nil, fmt.Errorf("no available drivers at the moment")
		}
		log.Fatalf("Failed to select driver_id :%v",err)
	}
	return &pb.LocationResponse{Id: id, Name: name, Phone: phone, Vehicle: vehicle}, nil	
}

func GetDateWithTime() string {
	timeStr := time.Now().Format("2006-01-02 15:04:05 Monday")
	return timeStr
}
