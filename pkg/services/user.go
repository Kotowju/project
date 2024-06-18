package services

import (
	"project/pkg/db"
	"project/pkg/models"
)

// tworzenie nowego użytkownika w bazie
func CreateUser(user models.User) error {
	dbConn, err := db.Connect()
	if err != nil {
		return err
	}
	defer dbConn.Close()
	return nil
}

// usuwanie użytkownika z bazy
func DeleteUser(id int) error {
	dbConn, err := db.Connect()
	if err != nil {
		return err
	}
	defer dbConn.Close()
	return nil
}

// Update użytkownika w bazie
func UpdateUser(user models.User) error {
	dbConn, err := db.Connect()
	if err != nil {
		return err
	}
	defer dbConn.Close()
	return nil
}
