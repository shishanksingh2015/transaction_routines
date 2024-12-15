package db

import (
	"context"
	"database/sql"
)

//go:generate mockgen -source=base_db.go  -destination=../mocks/db/base_db.go -package=db
type BaseDB interface {
	DB(ctx context.Context) *sql.DB
	Insert(ctx context.Context, query string, data ...any) error
}

type baseDb struct {
	db *sql.DB
}

func NewBaseDB(baseDB *sql.DB) BaseDB {
	return baseDb{db: baseDB}
}

func (b baseDb) DB(ctx context.Context) *sql.DB {
	return b.db
}

func (b baseDb) Insert(ctx context.Context, query string, data ...any) error {
	_, err := b.db.ExecContext(ctx, query, data...)
	if err != nil {
		return err
	}

	return nil
}
