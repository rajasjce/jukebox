package controllers

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

// Helper function to set up the test database and tables
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("failed to open test database: %v", err)
	}

	_, err = db.Exec(`
		CREATE TABLE albums (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT,
			release_date TEXT,
			genre TEXT,
			price REAL,
			description TEXT
		);
		CREATE TABLE album_musicians (
			album_id INTEGER,
			musician_id INTEGER
		);
		CREATE TABLE musicians (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			musician_type TEXT NOT NULL
		);
	`)
	if err != nil {
		t.Fatalf("failed to create test tables: %v", err)
	}

	return db
}
