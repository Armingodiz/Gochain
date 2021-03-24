package main

import (
	"log"
	"net/http"
	"os"
	"github.com/gorilla/handlers"
)

func main() {
	port := os.Args[1] // inputted port
	//router := bet.NewRouter(address)

	//allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	//allowedMethods := handlers.AllowedMethods([]string{"GET", "POST"})

	// launch server
	log.Fatal(http.ListenAndServe(":"+port,
		handlers.CORS(allowedOrigins, allowedMethods)(router)))
}
