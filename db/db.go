package db

import (
	"context"
	"database/sql"
	"embed"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "github.com/lib/pq"

	"github.com/chris-ramon/golang-scaffolding/config"
	"github.com/chris-ramon/golang-scaffolding/db/models"
)

//go:embed migrations/*.sql
var fs embed.FS

type db struct {
	sqlDB   *sql.DB
	queries *models.Queries
}

func (d *db) Ping(ctx context.Context) error {
	if err := d.sqlDB.PingContext(ctx); err != nil {
		return err
	}

	return nil
}

func (d *db) Close() error {
	return d.sqlDB.Close()
}

func (d *db) Migrate() error {
	migrations, err := iofs.New(fs, "migrations")
	if err != nil {
		return err
	}
	driver, err := postgres.WithInstance(d.sqlDB, &postgres.Config{})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithInstance("iofs", migrations, "postgres", driver)
	if err != nil {
		return err
	}
	if err := m.Up(); err != migrate.ErrNoChange {
		return err
	}
	return nil
}

func (db *db) Queries() *models.Queries {
	return db.queries
}

func New(dbConfig *config.DBConfig) (*db, error) {
	connStr := fmt.Sprintf("user=%s password=%s host=%s dbname=%s sslmode=%s",
		dbConfig.User, dbConfig.PWD, dbConfig.Host, dbConfig.Name, dbConfig.SSLMode)

	_db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	queries := models.New(_db)

	return &db{
		sqlDB:   _db,
		queries: queries,
	}, nil
}

type DB interface {
	Ping(ctx context.Context) error
	Close() error
	Migrate() error
	Queries() *models.Queries
}
