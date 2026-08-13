package db

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/JuniorCorzo/symphony-lofi/internal/config"
)

// RemoteSong represents the JSON schema for songs fetched from remote servers.
type RemoteSong struct {
	URL      string `json:"url"`
	Category string `json:"category"`
}

// loadSavedETag reads the locally cached HTTP ETag from disk.
func loadSavedETag() string {
	data, _ := os.ReadFile(config.GetFilePath(".etag"))
	return string(data)
}

// saveETag persists a new ETag value to disk for future conditional requests.
func saveETag(etag string) {
	if etag == "" {
		return
	}
	_ = os.WriteFile(config.GetFilePath(".etag"), []byte(etag), 0644)
}

// fetchRemote performs a conditional HTTP GET request using the saved ETag.
// Returns nil response and no error when the server responds 304 Not Modified.
func fetchRemote(jsonURL, savedETag string) (*http.Response, error) {
	req, err := http.NewRequest("GET", jsonURL, nil)
	if err != nil {
		return nil, err
	}

	if savedETag != "" {
		req.Header.Set("If-None-Match", savedETag)
	}

	client := &http.Client{Timeout: 10 * time.Second}
	return client.Do(req)
}

// decodeSongs decodes the HTTP response body as a JSON array of RemoteSong.
func decodeSongs(body io.Reader) ([]RemoteSong, error) {
	var songs []RemoteSong
	if err := json.NewDecoder(body).Decode(&songs); err != nil {
		return nil, err
	}
	return songs, nil
}

// SyncSongsFromRemote fetches remote JSON playlists using HTTP ETag conditional requests (304 Not Modified).
// If updated, it inserts new songs into SQLite and caches the new ETag.
func SyncSongsFromRemote(repo *SongRepository, jsonURL string) {
	if jsonURL == "" {
		return
	}

	savedETag := loadSavedETag()

	resp, err := fetchRemote(jsonURL, savedETag)
	if err != nil {
		slog.Warn("Could not connect to remote song list server, using SQLite cache", "error", err)
		return
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusNotModified:
		slog.Info("Remote song list unchanged (304 Not Modified)")
	case http.StatusOK:
		songs, err := decodeSongs(resp.Body)
		if err != nil {
			slog.Error("Failed decoding remote songs JSON response", "error", err)
			return
		}

		_ = repo.SaveSongsBulk(songs)
		saveETag(resp.Header.Get("ETag"))
		slog.Info("Successfully synced remote songs into SQLite database", "count", len(songs))
	default:
		slog.Warn("Unexpected HTTP status from remote song server", "status", resp.StatusCode)
	}
}
