package services

import (
	"log/slog"

	"github.com/JuniorCorzo/symphony-lofi/internal/mpv"
)

// SongService coordinates playlist actions with the mpv playback engine.
type SongService struct {
	playlist *Playlist
	player   *mpv.MpvInstance
	state    *mpv.PlayerState
}

// NewSongService constructs a new SongService instance.
func NewSongService(playlist *Playlist, player *mpv.MpvInstance, state *mpv.PlayerState) *SongService {
	return &SongService{
		playlist: playlist,
		player:   player,
		state:    state,
	}
}

// PlaySong loads and plays the current track, pre-buffering the next track in mpv's queue.
func (service *SongService) PlaySong() {
	currentSong, ok := service.playlist.Current()
	if !ok || currentSong.URL == "" {
		slog.Warn("No current song found in playlist to play")
		return
	}

	if service.state != nil {
		service.state.SetLoading()
	}

	slog.Info("Loading and playing song", "id", currentSong.ID, "url", currentSong.URL, "category", currentSong.Category)
	if err := service.player.LoadSong(currentSong.URL); err != nil {
		slog.Error("Failed to load song in mpv player", "url", currentSong.URL, "error", err)
	}

	// Preload next track in queue asynchronously
	if nextTrack, hasNext := service.playlist.PeekNext(); hasNext {
		slog.Info("Pre-buffering next song in mpv queue", "url", nextTrack.URL)
		_ = service.player.PreloadSong(nextTrack.URL)
	}
}

// Next advances to the next track in the playlist and starts playback.
func (service *SongService) Next() {
	nextSong, ok := service.playlist.Next()
	if !ok {
		slog.Warn("Playlist is empty, cannot advance to next song")
		return
	}

	slog.Info("Advanced to next song in playlist", "id", nextSong.ID, "url", nextSong.URL)
	service.PlaySong()
}

// TogglePause toggles the play/pause state in the mpv instance.
func (service *SongService) TogglePause() {
	slog.Info("Toggling play/pause state in mpv player")
	if err := service.player.TogglePause(); err != nil {
		slog.Error("Failed to toggle pause in mpv player", "error", err)
	}
}
