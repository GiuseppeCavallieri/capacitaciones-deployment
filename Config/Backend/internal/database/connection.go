package database

import (
	"database/sql"
	"fmt"
	"time"

	"veterinaria/backend/internal/config"

	_ "github.com/jackc/pgx/v5/stdlib"
)

// ConnectDB opens a sql.DB connection using pgx driver.
func ConnectDB(cfg *config.Config) (*sql.DB, error) {
	dsn := cfg.PostgresDSN
	if dsn == "" {
		return nil, fmt.Errorf("empty database DSN (POSTGRES_DSN)")
	}

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(5 * time.Minute)
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)

	// ping to ensure the connection is valid
	deadline := time.Now().Add(5 * time.Second)
	for time.Now().Before(deadline) {
		if err := db.Ping(); err == nil {
			return db, nil
		}
		time.Sleep(200 * time.Millisecond)
	}

	return nil, fmt.Errorf("unable to connect to postgres")
}