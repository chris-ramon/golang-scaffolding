package db

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/chris-ramon/golang-scaffolding/config"
)

type DB struct {
	db *sql.DB
}

func (d *DB) Ping() error {
	if err := d.db.PingContext(context.Background()); err != nil {
		return err
	}

	return nil
}

func (d *DB) Close() error {
	return d.db.Close()
}

func New(dbConfig *config.DBConfig) (*DB, error) {
	connStr := fmt.Sprintf("user=%s password=%s host=%s dbname=%s sslmode=%s",
		dbConfig.User, dbConfig.PWD, dbConfig.Host, dbConfig.Name, dbConfig.SSLMode)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return &DB{
		db: db,
	}, nil
}
