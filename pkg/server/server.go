// pkg/server/server.go

package server

import (
	"net/http"
	"project/pkg/handlers"

	"github.com/gorilla/mux"
)

// Start serwera
func StartServer() error {
	r := mux.NewRouter()
	r.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/keys", handlers.CreateKey).Methods("POST")
	r.HandleFunc("/items", handlers.CreateItem).Methods("POST")
	r.HandleFunc("/policies", handlers.CreatePolicy).Methods("POST")

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		return err
	}
	return nil
}
