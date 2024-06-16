package main

import (
	"log"
	"net/http"
	"project/pkg/db"
	"project/pkg/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/keys", handlers.CreateKey).Methods("POST")
	r.HandleFunc("/items", handlers.CreateItem).Methods("POST")
	r.HandleFunc("/policies", handlers.CreatePolicy).Methods("POST")
	// Add other routes as necessary

	// Initialize ClickHouse connection
	dbConn, err := db.Connect()
	if err != nil {
		log.Fatalf("Error connecting to ClickHouse: %v", err)
	}
	defer dbConn.Close()

	// Start the HTTP server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Error starting HTTP server: %v", err)
	}
}
