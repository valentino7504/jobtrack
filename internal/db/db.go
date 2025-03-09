// Package db provides functions for connecting to the sqlite databas
// and initializing the database of jobs.
package db

import (
	"database/sql"
	"errors"
	"fmt"

	_ "modernc.org/sqlite"
)

// GetConnection returns a pointer to the sqlite database handle.
func GetConnection() (*sql.DB, error) {
	path := "./test.db"
	sqliteDB, err := sql.Open("sqlite", path)
	if err != nil {
		fmt.Println("Couldn't connect to db:", err)
		return nil, err
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
		fmt.Println("Nil DB pointer")
		return errors.New("The pointer passed to InitDB is nil")
	}
	if err := db.Ping(); err != nil {
		fmt.Println("No connection to the db:", err)
		return err
	}
	if _, err := db.Exec(schemaQuery); err != nil {
		fmt.Println("Error creating database:", err)
		return err
	}
	return nil
}
