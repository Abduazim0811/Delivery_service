package main

import (
	"gateway-service/api/routers"
	"log"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {

	mux :=routers.ClientRoutes()

	log.Println("Server is listening on port:",os.Getenv("server_url"))
	http.ListenAndServe(os.Getenv("server_url"), mux)
}
