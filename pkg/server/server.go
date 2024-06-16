// pkg/server/server.go

package server

import (
    "net/http"
    "project/pkg/handlers"
    "github.com/gorilla/mux"
)

// StartServer initializes and starts the HTTP server.
func StartServer() error {
    r := mux.NewRouter()
    r.HandleFunc("/users", handlers.CreateUser).Methods("POST")
    r.HandleFunc("/keys", handlers.CreateKey).Methods("POST")
    r.HandleFunc("/items", handlers.CreateItem).Methods("POST")
    r.HandleFunc("/policies", handlers.CreatePolicy).Methods("POST")
    // Add other routes as necessary

    // Start the HTTP server
    err := http.ListenAndServe(":8080", r)
    if err != nil {
        return err
    }
    return nil
}
