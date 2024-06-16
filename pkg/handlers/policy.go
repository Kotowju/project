package handlers

import (
    "encoding/json"
    "net/http"
    "project/pkg/models"
    "project/pkg/services"
)

// CreatePolicy handles the creation of a new policy.
func CreatePolicy(w http.ResponseWriter, r *http.Request) {
    var policy models.Policy
    if err := json.NewDecoder(r.Body).Decode(&policy); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := services.CreatePolicy(policy); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
}
