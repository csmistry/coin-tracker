package main

import (
	"fmt"
	"net/http"

	"github.com/csmistry/coin-tracker/backend/handlers"
	"github.com/csmistry/coin-tracker/backend/pkg/wallet"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	router := mux.NewRouter()

	//  Define routes
	router.HandleFunc("/addresses", handlers.ListAddresses).Methods("GET")
	router.HandleFunc("/addresses/{id}", handlers.GetAddress).Methods("GET")
	router.HandleFunc("/addresses/{id}", handlers.AddAddress).Methods("POST")
	router.HandleFunc("/addresses/{id}", handlers.RemoveAddress).Methods("DELETE")

	// Create new in-memory wallet
	wallet.Init()

	// Enable CORS
	handler := cors.Default().Handler(router)
	fmt.Println("Serving requests on port :8080")
	http.ListenAndServe(":8080", handler)
}
