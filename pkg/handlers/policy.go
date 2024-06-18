package handlers

import (
	"encoding/json"
	"net/http"
	"project/pkg/models"
	"project/pkg/services"
	"strconv"

	"github.com/gorilla/mux"
)

// tworzenie Policy
func CreatePolicy(w http.ResponseWriter, r *http.Request) {
	var policy models.Policy
	if err := json.NewDecoder(r.Body).Decode(&policy); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := services.CreatePolicy(policy); err != nil {
		http.Error(w, "Failed to create policy", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// usuwanie
func DeletePolicy(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid policy ID", http.StatusBadRequest)
		return
	}

	if err := services.DeletePolicy(id); err != nil {
		http.Error(w, "Failed to delete policy", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Update polityki
func UpdatePolicy(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid policy ID", http.StatusBadRequest)
		return
	}

	var policy models.Policy
	if err := json.NewDecoder(r.Body).Decode(&policy); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	policy.ID = id

	if err := services.UpdatePolicy(policy); err != nil {
		http.Error(w, "Failed to update policy", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
