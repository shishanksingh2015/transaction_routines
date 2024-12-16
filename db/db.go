package db

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib" // Use pgx as database/sql driver
	"log"
)

func ConnectDb(config Config) *sql.DB {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName)
	db, err := sql.Open(config.DBDriver, connStr)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	return db
}
