package db

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type RemoteSong struct {
	URL      string `json:"url"`
	Category string `json:"category"`
}

// SyncSongsFromRemote descarta descargas innecesarias usando ETag HTTP (304 Not Modified)
func SyncSongsFromRemote(repo *SongRepository, jsonURL string) {
	if jsonURL == "" {
		return
	}

	configDir, err := os.UserConfigDir()
	if err != nil {
		configDir = os.TempDir()
	}

	etagPath := filepath.Join(configDir, "symphony-lofi", ".etag")
	savedETag, _ := os.ReadFile(etagPath)

	req, err := http.NewRequest("GET", jsonURL, nil)
	if err != nil {
		slog.Error("Failed to create HTTP request for song sync", "error", err)
		return
	}

	if len(savedETag) > 0 {
		req.Header.Set("If-None-Match", string(savedETag))
	}

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		slog.Warn("Could not connect to remote song list server, using SQLite cache", "error", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotModified {
		slog.Info("Remote song list unchanged (304 Not Modified)")
		return
	}

	if resp.StatusCode == http.StatusOK {
		var remoteSongs []RemoteSong
		if err := json.NewDecoder(resp.Body).Decode(&remoteSongs); err != nil {
			slog.Error("Failed decoding remote songs JSON response", "error", err)
			return
		}

		for _, song := range remoteSongs {
			if song.URL != "" {
				_ = repo.SaveSong(song.URL, song.Category)
			}
		}

		newETag := resp.Header.Get("ETag")
		if newETag != "" {
			_ = os.WriteFile(etagPath, []byte(newETag), 0644)
		}

		slog.Info("Successfully synced remote songs into SQLite database", "count", len(remoteSongs))
	}
}
