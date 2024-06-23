package handlers

import (
	"encoding/json"
	auth "gateway-service/auth/jwt"
	"gateway-service/genpb/pbclient"
	"gateway-service/internal/clientmodel"
	"gateway-service/internal/email"
	"gateway-service/internal/file"
	"log"
	"net/http"
)

type ClientService struct {
	Client pbclient.ClientServiceClient
}

func NewClientService(cl pbclient.ClientServiceClient) *ClientService {
	return &ClientService{Client: cl}
}

func (cs *ClientService) RegisterNewClient(w http.ResponseWriter, r *http.Request) {
	var clJsonReq clientmodel.ClientRegisterRequest
	err := json.NewDecoder(r.Body).Decode(&clJsonReq)
	if err != nil {
		log.Println("Unable to Decode request body while registering:", err)
		http.Error(w, "Invalid Request !", http.StatusBadRequest)
		return
	}

	if err := clientmodel.ClientRegisterValidation(clJsonReq); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var clPbReq = pbclient.ClRequest{
		Name:       clJsonReq.Name,
		Email:      clJsonReq.Email,
		Phone:      clJsonReq.Phone,
		City:       clJsonReq.City,
		Street:     clJsonReq.Street,
		HomeNumber: clJsonReq.Home_number,
	}
	resp, err := cs.Client.CreateClient(r.Context(), &clPbReq)
	if err != nil {
		log.Println("Unable to Get response from client-service on CreateClient method :", err)
		http.Error(w, "Something went wrong !", http.StatusInternalServerError)
		return
	}

	var clientLogin = clientmodel.ClientLogin{
		Id:    resp.Id,
		Email: clJsonReq.Email,
	}
	err = file.WriteNewClientToFile("./clients.json", clientLogin)
	if err != nil {
		log.Println("Failed to write newclient to json file:", err)
		http.Error(w, "Something went wrong !", http.StatusInternalServerError)
		return
	}
	if err := email.SendGomail(clJsonReq.Name, clJsonReq.Email, resp.Id); err != nil {
		log.Println("Failed to send an email to client:", err)
		http.Error(w, "Something went wrong !", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	if err = json.NewEncoder(w).Encode(resp); err != nil {
		log.Println("Unable to encode response from client service :", err)
		http.Error(w, "Something went wrong !", http.StatusInternalServerError)
		return
	}
}

func (cs *ClientService) LoginClientToapplication(w http.ResponseWriter, r *http.Request) {
	var cLog clientmodel.ClientLogin
	err := json.NewDecoder(r.Body).Decode(&cLog)
	if err != nil {
		log.Println("Unable to Decode request body while login:", err)
		http.Error(w, "Invalid Request !", http.StatusBadRequest)
		return
	}
	if err = clientmodel.ValidateEmail(cLog.Email); err != nil {
		http.Error(w, "Invalid email !", http.StatusBadRequest)
		return
	}

	if err :=file.CheckClientFromFile("./clients.json", cLog); err != nil {
		log.Println("Unable to find clients while login:", err)
		http.Error(w, err.Error() , http.StatusBadRequest)
		return
	}

	jwtHandler :=auth.JWTHandler{
		Company: "Delivery-service",
		Email: cLog.Email,
		Role: "client",
		Id: cLog.Id,
	}

	token , err :=jwtHandler.GenerateToken()
	if err != nil {
		log.Println("failed to generate jwt token:", err)
		http.Error(w, "something went wrong !", http.StatusBadRequest)
		return
	}

	var res = clientmodel.LoginResponse{
		Status: "Successfull",
		Token: token,
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(200)
	if err =json.NewEncoder(w).Encode(res); err != nil {
		log.Println("failed to encode response token:", err)
		http.Error(w, "something went wrong !", http.StatusBadRequest)
		return
	}
	
}
