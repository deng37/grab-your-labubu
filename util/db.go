package util

import (
    "database/sql"
    _ "modernc.org/sqlite"
    "log"
    "os"
)

var DB *sql.DB

func InitDB() {
    dbPath := os.Getenv("DATABASE_URL")	// Getting env from Fly.io
    if dbPath == "" {
        dbPath = "./labubu.db" // Local only
    }

    var err error
    DB, err = sql.Open("sqlite", dbPath)
    if err != nil {
        log.Fatalf("Failed open DB: %v", err)
    }

    DB.Exec("PRAGMA journal_mode=WAL;")
    DB.Exec("PRAGMA synchronous=NORMAL;")
    DB.Exec("PRAGMA busy_timeout=5000;")

    createTables()
}

func createTables() {
    query := `
    CREATE TABLE IF NOT EXISTS leaderboard (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_name TEXT,
		user_email TEXT,
		ip_address TEXT UNIQUE,
        duration_ms REAL,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME
    );`

    if _, err := DB.Exec(query); err != nil {
        log.Fatalf("Failed creating DB: %v", err)
    }
}