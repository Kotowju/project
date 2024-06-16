package services

import (
    "project/pkg/models"
    "project/pkg/db"
)

// CreateUser inserts a new user into the database.
func CreateUser(user models.User) error {
    conn, err := db.Connect()
    if err != nil {
        return err
    }
    defer conn.Close()

    query := "INSERT INTO users (username, email) VALUES (?, ?)"
    _, err = conn.Exec(query, user.Username, user.Email)
    if err != nil {
        return err
    }
    return nil
}
