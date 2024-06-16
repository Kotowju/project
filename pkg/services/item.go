package services

import (
    "project/pkg/models"
    "project/pkg/db"
)

// CreateItem inserts a new item into the database.
func CreateItem(item models.Item) error {
    conn, err := db.Connect()
    if err != nil {
        return err
    }
    defer conn.Close()

    query := "INSERT INTO items (sector, name, description) VALUES (?, ?, ?)"
    _, err = conn.Exec(query, item.Sector, item.Name, item.Description)
    if err != nil {
        return err
    }
    return nil
}
