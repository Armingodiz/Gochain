package main

import (
	"github.com/ArminGodiz/Gochain/loan"
	"github.com/gorilla/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Args[1] // inputted port
	router := loan.NewRouter(port)

	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST"})
	// launch server
	log.Fatal(http.ListenAndServe(":"+port,
		handlers.CORS(allowedOrigins, allowedMethods)(router)))
}
