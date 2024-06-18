package services

import (
	"project/pkg/db"
	"project/pkg/models"
)

// tworzenie nowej policy w bazie
func CreatePolicy(policy models.Policy) error {
	dbConn, err := db.Connect()
	if err != nil {
		return err
	}
	defer dbConn.Close()
	return nil
}

// usuwanie polityki z bazy
func DeletePolicy(id int) error {
	dbConn, err := db.Connect()
	if err != nil {
		return err
	}
	defer dbConn.Close()
	return nil
}

// Update policy w bazie
func UpdatePolicy(policy models.Policy) error {
	dbConn, err := db.Connect()
	if err != nil {
		return err
	}
	defer dbConn.Close()
	return nil
}
