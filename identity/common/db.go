package common

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func Connect(conn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
		return nil, err
	}
	return db, nil
}
