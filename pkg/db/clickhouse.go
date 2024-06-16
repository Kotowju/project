package db

import (
	"database/sql"
	"log"

	_ "github.com/ClickHouse/clickhouse-go"
)

// Connect creates a connection to the ClickHouse database.
func Connect() (*sql.DB, error) {
	db, err := sql.Open("clickhouse", "tcp://clickhouse:9000?username=default&password=&database=default")
	if err != nil {
		log.Println("Error opening ClickHouse connection:", err)
		return nil, err
	}
	return db, nil
}
