package method

import (
	"client-service/pb"
	"database/sql"
	"log"
	"time"
)

func StoreNewClient(db *sql.DB, req *pb.ClRequest) error {
	
	query := `
		insert into clients(id, name, email, phone, created_at)
		values($1, $2, $3, $4, $5)
	`
	_, err := db.Exec(query, req.Id,req.Name, req.Email, req.Phone ,req.Time)
	if err != nil {
		log.Println("failed to insert client to database:", err)
		return err
	}

	query = `
		insert into locations(user_id, city, street, home_number)
		values($1, $2, $3, $4)
	`

	_, err = db.Exec(query, req.Id, req.City, req.Street, req.HomeNumber)
	if err != nil {
		log.Println("failed to insert client locations:", err)
		return err
	}
	return nil
}

func UpdateClientInDatabase(db *sql.DB, req *pb.ClRequest)error{
	query :=`
		update locations
		set city=$1, street=$2, home_number=$3
		where user_id = $4
	`
	_,err :=db.Exec(query, req.City, req.Street, req.HomeNumber, req.Id)
	if err != nil {
		log.Println("Unable to update locations:",err)
		return err
	}

	query =`
		update clients 
		set name =$1, email =$2, phone = $3
		where id = $4
	`

	_,err =db.Exec(query, req.Name, req.Email,req.Phone, req.Id)
	if err != nil {
		log.Println("Unable to update clients:",err)
		return err
	}

	return nil
}

func DeleteClientFromDatabase(db *sql.DB, req *pb.ClResponse)error{
	query :=` delete from locations
		where user_id = $1
	`
	_,err :=db.Exec(query, req.Id)
	if err != nil {
		log.Println("Unable to delete client locations:",err)
		return err
	}

	query =	`
		delete from clients
		where id = $1
	`
	_, err =db.Exec(query, req.Id)
	if err != nil {
		log.Println("Unable to delete client:",err)
		return err
	}
	return nil
}

func GetClientFromDatabse(db *sql.DB, req *pb.ClResponse)(res *pb.ClRequest){
	query :=`
		select c.id, c.name, c.email, c.phone ,c.created_at, l.city, l.street, l.home_number from clients c
		join locations l on c.id = l.user_id
		where c.id =$1
	`
	var id, name, email, phone, time, city, street, home string
	row :=db.QueryRow(query,req.Id)
	if err :=row.Scan(&id, &name, &email,&phone, &time, &city, &street, &home); err != nil {
		log.Fatal("Unable to get client information:",err)
	}
	return &pb.ClRequest{Id: id, Name: name, Email: email, Time: time, City: city, Street: street, HomeNumber: home}
}

func GetClientAddressFromDatabase(db *sql.DB, req *pb.ClResponse)(*pb.ClAddress){
	query :=`
		select c.id, c.phone , l.city, l.street, l.home_number from clients s
		join locations l on c.id =l.user_id
		where c.id = $1
	`
	var id, phone, city, street, home string
	row :=db.QueryRow(query, req.Id)
	if err :=row.Scan(&id, &phone, &city, &street, &home); err != nil {
		log.Fatal("failed to get client locations:",err)
	}
	return &pb.ClAddress{Id: id, City: city, Street: street, HomeNumber: home}
}

	func GetDateWithTime()string{
		timeStr := time.Now().Format("2006-01-02 15:04:05 Monday")
		return timeStr
	}
