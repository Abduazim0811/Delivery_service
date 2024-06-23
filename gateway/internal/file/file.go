package file

import (
	"encoding/json"
	"fmt"
	"gateway-service/internal/clientmodel"
	"os"
)

func WriteNewClientToFile(filepath string ,cl clientmodel.ClientLogin)error{
	f, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer f.Close()

	var clients []clientmodel.ClientLogin
	err = json.NewDecoder(f).Decode(&clients)
	if  err != nil {
		clients = []clientmodel.ClientLogin{}
	}

	for _, client := range clients {
		if client.Id == cl.Id {
			return fmt.Errorf("user with ID '%s' already exists", cl.Id)
		}
	}
	clients = append(clients, cl)
	f, err = os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer f.Close()
	err = json.NewEncoder(f).Encode(clients)
	if  err != nil {
		return fmt.Errorf("failed to encode users to JSON: %v", err)
	}
	return nil
}

func CheckClientFromFile(filepath string, req clientmodel.ClientLogin)error{
	f, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer f.Close()

	var clients []clientmodel.ClientLogin
	err = json.NewDecoder(f).Decode(&clients)
	if err != nil {
		return fmt.Errorf("failed to decode file: %v", err)
	}

	for _, client := range clients {
		if client.Id == req.Id {
			return nil
		}
	}
	return fmt.Errorf("no client found with this id and email !")
}