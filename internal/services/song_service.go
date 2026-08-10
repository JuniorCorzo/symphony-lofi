package services

import (
	"log/slog"

	"github.com/JuniorCorzo/symphony-lofi/internal/mpv"
)

type SongService struct {
	playlist *Playlist
	player   *mpv.MpvInstance
}

func NewSongService(playlist *Playlist, player *mpv.MpvInstance) *SongService {
	return &SongService{
		playlist: playlist,
		player:   player,
	}
}

func (service *SongService) PlaySong() {
	currentSong, ok := service.playlist.Current()
	if !ok || currentSong.URL == "" {
		slog.Warn("No current song found in playlist to play")
		return
	}

	slog.Info("Loading and playing song", "id", currentSong.ID, "url", currentSong.URL, "category", currentSong.Category)
	if err := service.player.LoadSong(currentSong.URL); err != nil {
		slog.Error("Failed to load song in mpv player", "url", currentSong.URL, "error", err)
	}
}

func (service *SongService) Next() {
	nextSong, ok := service.playlist.Next()
	if !ok {
		slog.Warn("Playlist is empty, cannot advance to next song")
		return
	}

	slog.Info("Advanced to next song in playlist", "id", nextSong.ID, "url", nextSong.URL)
	service.PlaySong()
}

func (service *SongService) TogglePause() {
	slog.Info("Toggling play/pause state in mpv player")
	if err := service.player.TogglePause(); err != nil {
		slog.Error("Failed to toggle pause in mpv player", "error", err)
	}
}
