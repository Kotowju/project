package services

import (
    "project/pkg/models"
    "project/pkg/db"
)

// CreatePolicy inserts a new policy into the database.
func CreatePolicy(policy models.Policy) error {
    conn, err := db.Connect()
    if err != nil {
        return err
    }
    defer conn.Close()

    query := "INSERT INTO policies (name, definition) VALUES (?, ?)"
    _, err = conn.Exec(query, policy.Name, policy.Definition)
    if err != nil {
        return err
    }
    return nil
}
