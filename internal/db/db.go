package db

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

const DefaultDBPath = "database.sqlite"

func InitDb() (*db, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("Error getting user home directory: %v", err)
	}

	appDir := filepath.Join(homeDir, ".1bookmark")
	if err := os.MkdirAll(appDir, 0755); err != nil {
		return nil, fmt.Errorf("Error creating app directory: %v", err)
	}

	dbPath := filepath.Join(appDir, DefaultDBPath)
	conn, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, fmt.Errorf("Error opening database: %v", err)
	}

	if err := conn.Ping(); err != nil {
		return nil, fmt.Errorf("Error connecting to database: %v", err)
	}

	return &db{conn: conn}, nil
}

type db struct {
	conn *sql.DB
}

func (d *db) Close() error {
	return d.conn.Close()
}
