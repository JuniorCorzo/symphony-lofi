package db

import (
	"database/sql"
	"log/slog"
)

type SongRepository struct {
	db *sql.DB
}

func NewSongRepository(db *sql.DB) *SongRepository {
	return &SongRepository{db: db}
}

func (songRepository *SongRepository) GetSongs() ([]Song, error) {
	query := `SELECT id, url, category FROM songs ORDER BY RANDOM()`

	rows, err := songRepository.db.Query(query)
	if err != nil {
		slog.Error("Failed executing query to get songs", "error", err)
		return nil, err
	}
	defer rows.Close()

	var songs []Song
	for rows.Next() {
		var s Song
		if err := rows.Scan(&s.ID, &s.URL, &s.Category); err != nil {
			slog.Error("Failed scanning row into Song struct", "error", err)
			continue
		}
		songs = append(songs, s)
	}

	if err := rows.Err(); err != nil {
		slog.Error("Error during rows iteration", "error", err)
		return nil, err
	}

	return songs, nil
}

func (songRepository *SongRepository) SaveSong(url, category string) error {
	query := `INSERT OR IGNORE INTO songs (url, category) VALUES (?, ?)`

	_, err := songRepository.db.Exec(query, url, category)
	if err != nil {
		slog.Error("Failed to insert song into database", "url", url, "error", err)
		return err
	}

	return nil
}
