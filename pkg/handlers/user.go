package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"project/pkg/models"
	"project/pkg/services"
)

// CreateUser handles the creation of a new user.
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := services.CreateUser(user); err != nil {
		// Log the error for debugging purposes
		fmt.Println("Error creating user:", err)

		// Depending on the error type, you might want to return different status codes
		// For database-related errors, consider returning 500 Internal Server Error
		// For validation errors or specific conditions, return 400 Bad Request or other appropriate codes

		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
