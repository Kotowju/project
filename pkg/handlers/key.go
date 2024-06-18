package handlers

import (
	"encoding/json"
	"net/http"
	"project/pkg/models"
	"project/pkg/services"
	"strconv"

	"github.com/gorilla/mux"
)

// tworzenie klucza
func CreateKey(w http.ResponseWriter, r *http.Request) {
	var key models.Key
	if err := json.NewDecoder(r.Body).Decode(&key); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := services.CreateKey(key); err != nil {
		http.Error(w, "Failed to create key", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// usuwanie istniejacego
func DeleteKey(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid key ID", http.StatusBadRequest)
		return
	}

	if err := services.DeleteKey(id); err != nil {
		http.Error(w, "Failed to delete key", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Update istniejacego klucza
func UpdateKey(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid key ID", http.StatusBadRequest)
		return
	}

	var key models.Key
	if err := json.NewDecoder(r.Body).Decode(&key); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	key.ID = id

	if err := services.UpdateKey(key); err != nil {
		http.Error(w, "Failed to update key", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
