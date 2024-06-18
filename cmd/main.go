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
	r.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")
	r.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")

	r.HandleFunc("/keys", handlers.CreateKey).Methods("POST")
	r.HandleFunc("/keys/{id}", handlers.DeleteKey).Methods("DELETE")
	r.HandleFunc("/keys/{id}", handlers.UpdateKey).Methods("PUT")

	r.HandleFunc("/items", handlers.CreateItem).Methods("POST")
	r.HandleFunc("/items/{id}", handlers.DeleteItem).Methods("DELETE")
	r.HandleFunc("/items/{id}", handlers.UpdateItem).Methods("PUT")

	r.HandleFunc("/policies", handlers.CreatePolicy).Methods("POST")
	r.HandleFunc("/policies/{id}", handlers.DeletePolicy).Methods("DELETE")
	r.HandleFunc("/policies/{id}", handlers.UpdatePolicy).Methods("PUT")

	// polaczenie ClickHouse
	dbConn, err := db.Connect()
	if err != nil {
		log.Fatalf("Error connecting to ClickHouse: %v", err)
	}
	defer dbConn.Close()

	//  HTTP serwer
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Error starting HTTP server: %v", err)
	}
}
