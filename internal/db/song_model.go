package db

// Song represents a track stored in the SQLite database.
type Song struct {
	ID       int64  `json:"ID"`
	URL      string `json:"URL"`
	Category string `json:"Category"`
}
