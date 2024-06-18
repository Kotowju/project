package services

import (
	"project/pkg/db"
	"project/pkg/models"
)

// tworzy nowy item w bazie danych
func CreateItem(item models.Item) error {
	dbConn, err := db.Connect()
	if err != nil {
		return err
	}
	defer dbConn.Close()
	return nil
}

// usuwa item z DB na podstawie ID
func DeleteItem(id int) error {
	dbConn, err := db.Connect()
	if err != nil {
		return err
	}
	defer dbConn.Close()
	return nil
}

// Update itemu w bazie
func UpdateItem(item models.Item) error {
	dbConn, err := db.Connect()
	if err != nil {
		return err
	}
	defer dbConn.Close()
	return nil
}
