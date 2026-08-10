package db

import (
	"database/sql"
	"log/slog"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

func InitDB() *sql.DB {
	configDir, err := os.UserConfigDir()
	if err != nil {
		configDir = os.TempDir()
	}

	appDir := filepath.Join(configDir, "symphony-lofi")
	_ = os.MkdirAll(appDir, 0o755)

	dbPath := filepath.Join(appDir, "songs.db")

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

	// Seed default lofi streams if empty
	seedQuery := `INSERT OR IGNORE INTO songs (url, category) VALUES
		('https://music.youtube.com/watch?v=nSXGgI5W3LU', 'lofi'),
		('https://music.youtube.com/watch?v=jfKfPfyJRdk', 'lofi'),
		('https://music.youtube.com/watch?v=5qap5aO4i9A', 'lofi');`

	if _, err := db.Exec(seedQuery); err != nil {
		slog.Error("Failed to seed default songs", "error", err)
	} else {
		slog.Info("Default Lofi playlist seeded")
	}
}
