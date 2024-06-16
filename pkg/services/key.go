package services

import (
	"project/pkg/db"
	"project/pkg/models"
	"project/pkg/vault"
)

// CreateKey stores the key's secret in Vault and the key's metadata in the database.
func CreateKey(key models.Key) error {
	// Store the secret in Vault
	err := vault.StoreSecret(key.Name, key.Secret)
	if err != nil {
		return err
	}

	conn, err := db.Connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	query := "INSERT INTO keys (name, username) VALUES (?, ?)"
	_, err = conn.Exec(query, key.Name, key.Username)
	if err != nil {
		return err
	}
	return nil
}
