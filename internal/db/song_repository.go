package db

import (
	"database/sql"
	"log/slog"
	"strings"
)

// SongRepository handles database persistence operations for songs.
type SongRepository struct {
	db *sql.DB
}

// NewSongRepository constructs a new SongRepository instance.
func NewSongRepository(db *sql.DB) *SongRepository {
	return &SongRepository{db: db}
}

// GetSongs retrieves all stored songs in randomized order.
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

// SaveSong inserts a song URL and category into the database, ignoring duplicates.
func (songRepository *SongRepository) SaveSong(url, category string) error {
	query := `INSERT OR IGNORE INTO songs (url, category) VALUES (?, ?)`

	_, err := songRepository.db.Exec(query, url, category)
	if err != nil {
		slog.Error("Failed to insert song into database", "url", url, "error", err)
		return err
	}

	return nil
}

// buildBulkInsertQuery builds a parameterized multi-row INSERT query and its arguments from a list of RemoteSong.
// Returns an empty string and nil args if all songs have empty URLs.
func buildBulkInsertQuery(songs []RemoteSong) (string, []any) {
	query := strings.Builder{}
	query.WriteString("INSERT OR IGNORE INTO songs (url, category) VALUES ")

	args := make([]any, 0, len(songs)*2)
	first := true
	for _, song := range songs {
		if song.URL == "" {
			continue
		}
		if !first {
			query.WriteString(",")
		}
		query.WriteString("(?,?)")
		args = append(args, song.URL, song.Category)
		first = false
	}

	return query.String(), args
}

// SaveSongsBulk performs an atomic bulk insert by constructing a single multi-row INSERT statement,
// minimizing SQLite roundtrips to exactly one call regardless of slice size.
func (songRepository *SongRepository) SaveSongsBulk(songs []RemoteSong) error {
	if len(songs) == 0 {
		return nil
	}

	query, args := buildBulkInsertQuery(songs)
	if len(args) == 0 {
		return nil
	}

	if _, err := songRepository.db.Exec(query, args...); err != nil {
		slog.Error("Failed executing bulk insert", "error", err)
		return err
	}

	slog.Info("Successfully executed bulk insert", "count", len(args)/2)
	return nil
}
