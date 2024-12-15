package db

import (
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib" // Use pgx as database/sql driver
	"log"
)

func ConnectDb(config Config) *sql.DB {
	db, err := sql.Open("pgx", config.DBSource)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	return db
}
