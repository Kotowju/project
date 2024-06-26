package db

import (
	"database/sql"
	"log"

	_ "github.com/ClickHouse/clickhouse-go"
)

// Połączenie z bazą danych
func Connect() (*sql.DB, error) {
	db, err := sql.Open("clickhouse", "tcp://clickhouse:9000?username=default&password=&database=default")
	if err != nil {
		log.Println("Error opening ClickHouse connection:", err)
		return nil, err
	}
	return db, nil
}
