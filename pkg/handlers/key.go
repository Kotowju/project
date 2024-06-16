package handlers

import (
    "encoding/json"
    "net/http"
    "project/pkg/models"
    "project/pkg/services"
)

// CreateKey handles the creation of a new key.
func CreateKey(w http.ResponseWriter, r *http.Request) {
    var key models.Key
    if err := json.NewDecoder(r.Body).Decode(&key); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := services.CreateKey(key); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
}
