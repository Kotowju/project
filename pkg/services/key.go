package services

import (
	"project/pkg/db"
	"project/pkg/models"
)

// tworzenie klucza w bazie
func CreateKey(key models.Key) error {
	dbConn, err := db.Connect()
	if err != nil {
		return err
	}
	defer dbConn.Close()
	return nil
}

// usuwanie kluacz z DB przez ID
func DeleteKey(id int) error {
	dbConn, err := db.Connect()
	if err != nil {
		return err
	}
	defer dbConn.Close()
	return nil
}

// Update klucza w bazie
func UpdateKey(key models.Key) error {
	dbConn, err := db.Connect()
	if err != nil {
		return err
	}
	defer dbConn.Close()
	return nil
}
