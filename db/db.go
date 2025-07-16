package db

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

func DBConnnect() (*sql.DB, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return nil, err
	}
	// Ensure app directory exists: ~/.config/lexo
	appDir := filepath.Join(configDir, "tasks")
	if err := os.MkdirAll(appDir, 0700); err != nil {
		log.Fatalf("savePos: cannot create app directory %q: %v", appDir, err)
		return nil, err
	}
	db, err := sql.Open("sqlite3", appDir+"/tasks.db")
	if err != nil {
		return nil, err
	}

	// Create table
	_, err = db.Exec(`
    CREATE TABLE IF NOT EXISTS tasks (
      id INTEGER PRIMARY KEY AUTOINCREMENT,
      description TEXT NOT NULL,
      complete INTEGER NOT NULL CHECK(complete IN (0,1)),
      created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP
    )`)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return db, nil
}
