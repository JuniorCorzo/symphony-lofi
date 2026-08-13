// Package db handles SQLite database connections, migrations, and song repositories.
package db

import (
	"database/sql"
	"log/slog"

	"github.com/JuniorCorzo/symphony-lofi/internal/config"
	_ "modernc.org/sqlite"
)

// InitDB initializes and opens the SQLite database at ~/.config/symphony-lofi/songs.db.
// It automatically executes schema migrations if the songs table does not exist.
func InitDB() *sql.DB {
	dbPath := config.GetFilePath("songs.db")

	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		slog.Error("Failed to open SQLite database", "path", dbPath, "error", err)
		return nil
	}

	slog.Info("Connected to SQLite database", "path", dbPath)
	createTablesIfNotExists(db)

	return db
}

func createTablesIfNotExists(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS songs (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            url TEXT NOT NULL UNIQUE,
            category TEXT DEFAULT 'lofi',
            created_at DATETIME DEFAULT CURRENT_TIMESTAMP
        );`

	if _, err := db.Exec(query); err != nil {
		slog.Error("Failed to execute DDL migrations for songs table", "error", err)
		return
	}
	slog.Debug("Ensured songs table migration complete")
}
