package handlers

import (
    "encoding/json"
    "net/http"
    "project/pkg/models"
    "project/pkg/services"
)

// CreateItem handles the creation of a new item.
func CreateItem(w http.ResponseWriter, r *http.Request) {
    var item models.Item
    if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := services.CreateItem(item); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
}
