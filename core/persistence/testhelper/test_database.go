package testhelper

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file" // need to migration for migrate library
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/testcontainers/testcontainers-go"
	"log"
	"time"
)

// The implementation uses github.com/testcontainers/testcontainers-go it is easy clean api for integration testing
// RunPostgresContainer Note: We can use dockertest libray to spin containers for testing
// RunPostgresContainer will spin the container and pull the postgres image and will run migration when database is up.
func RunPostgresContainer(ctx context.Context, migrationsPath string) (*sql.DB, testcontainers.Container, error) {
	req := testcontainers.ContainerRequest{
		Image:        "postgres:15",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_PASSWORD": "super-secret",
			"POSTGRES_USER":     "user",
			"POSTGRES_DB":       "transaction-routines",
		},
	}

	postgresContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to start container: %w", err)
	}

	port, err := postgresContainer.MappedPort(ctx, "5432")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get mapped port: %w", err)
	}
	var db *sql.DB
	for i := 0; i < 10; i++ {
		connectionString := fmt.Sprintf("postgres://user:super-secret@0.0.0.0:%s/transaction-routines?sslmode=disable", port.Port())

		// Connect to the database
		db, err = sql.Open("pgx", connectionString)
		if err != nil {
			log.Printf("Error opening connection: %v\n", err)
			continue
		}
		err = db.Ping() // Ping to check if the DB is ready
		if err == nil {
			log.Println("PostgreSQL is ready!")
			break
		}

		log.Printf("Error pinging database: %v\n", err)

		time.Sleep(2 * time.Second) // wait before retrying
	}

	// Run migrations
	if err := runMigrations(db, migrationsPath); err != nil {
		return nil, nil, fmt.Errorf("failed to run migrations: %w", err)
	}

	return db, postgresContainer, nil
}

// runMigrations will run migrations for give path
func runMigrations(db *sql.DB, migrationsPath string) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("failed to create database driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://"+migrationsPath, "transaction-routines", driver)
	if err != nil {
		return fmt.Errorf("failed to create migration instance: %w", err)
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("migration failed: %w", err)
	}

	return nil
}
