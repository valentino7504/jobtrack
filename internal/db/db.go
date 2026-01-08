// Package db provides functions for connecting to the sqlite databas
// and initializing the database of jobs.
package db

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	_ "modernc.org/sqlite"
)

// GetConnection returns a pointer to the sqlite database handle.
func GetConnection() (*sql.DB, error) {
	var dir, path string
	if runtime.GOOS == "windows" {
		appData := os.Getenv("APPDATA")
		dir = filepath.Join(appData, "jobtrack")
	} else {
		home, _ := os.UserHomeDir()
		dir = filepath.Join(home, ".local", "share", "jobtrack")
	}
	path = filepath.Join(dir, "jobs.db")
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return nil, fmt.Errorf("failed to create database directory: %w", err)
	}
	sqliteDB, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, fmt.Errorf("couldn't connect to database: %w", err)
	}
	return sqliteDB, nil
}

// InitDB initializes the database, creating the table if it does not exist.
// It takes a pointer to the database handle
func InitDB(db *sql.DB) error {
	const schemaQuery = `CREATE TABLE IF NOT EXISTS jobs (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			company TEXT NOT NULL,
			position TEXT NOT NULL,
			status TEXT NOT NULL,
			location TEXT,
			applied_at TEXT NOT NULL DEFAULT (DATE('now')),
			salary_range TEXT,
			job_posting_url TEXT,
			created_at TEXT NOT NULL DEFAULT (CURRENT_TIMESTAMP),
			updated_at TEXT NOT NULL DEFAULT (CURRENT_TIMESTAMP)
		);`
	if db == nil {
		return errors.New("The pointer passed to InitDB is nil")
	}
	if err := db.Ping(); err != nil {
		return fmt.Errorf("No connection to the db: %w", err)
	}
	if _, err := db.Exec(schemaQuery); err != nil {
		return fmt.Errorf("Error creating database: %w", err)
	}
	return nil
}
