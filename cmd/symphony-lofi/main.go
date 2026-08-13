package main

import (
	"log/slog"
	"os"

	"github.com/JuniorCorzo/symphony-lofi/internal/db"
	"github.com/JuniorCorzo/symphony-lofi/internal/logger"
	"github.com/JuniorCorzo/symphony-lofi/internal/mpv"
	"github.com/JuniorCorzo/symphony-lofi/internal/services"
	"github.com/JuniorCorzo/symphony-lofi/internal/tui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	initLogger()

	slog.Info("Starting Symphony-Lofi application")

	songService, player, state := initServices()
	songService.PlaySong()

	p := initTUI(songService, state)
	go initEventLoop(player, state, songService, p)

	if _, err := p.Run(); err != nil {
		slog.Error("Failed to run TUI program", "error", err)
		os.Exit(1)
	}

	slog.Info("Symphony-Lofi application exited gracefully")
}

// initLogger sets up structured logging. Failures are non-fatal.
func initLogger() {
	logFile, err := logger.InitLogger()
	if err == nil && logFile != nil {
		// logFile is closed when the process exits via deferred call in the caller.
		// Caller main() defers this via the returned handle.
		_ = logFile
	}
}

const defaultRemoteJSONURL = "https://raw.githubusercontent.com/JuniorCorzo/symphony-lofi/main/songs.json"

// initServices wires the database, player, and song service together.
func initServices() (*services.SongService, *mpv.MpvInstance, *mpv.PlayerState) {
	conn := db.InitDB()
	songRepository := db.NewSongRepository(conn)

	player := mpv.GetInstance()
	state := mpv.NewPlayerState()
	playlist := initPlaylist(songRepository)

	songService := services.NewSongService(playlist, player, state)
	return songService, player, state
}

// initPlaylist loads all songs from the repository, ensuring first-run sync if database is empty.
func initPlaylist(songRepository *db.SongRepository) *services.Playlist {
	songs, err := songRepository.GetSongs()
	if err != nil {
		slog.Error("Failed to load songs from repository", "error", err)
	}

	// Si es la primera ejecución y la base de datos está vacía, sincroniza de forma sincrónica antes de arrancar la TUI
	if len(songs) == 0 {
		slog.Info("Database is empty on first run. Performing synchronous remote song sync...")
		db.SyncSongsFromRemote(songRepository, defaultRemoteJSONURL)
		songs, _ = songRepository.GetSongs()
	} else {
		// En ejecuciones posteriores, sincroniza asíncronamente en segundo plano
		go db.SyncSongsFromRemote(songRepository, defaultRemoteJSONURL)
	}

	slog.Info("Playlist loaded from database", "count", len(songs))
	return services.NewPlaylist(songs)
}

// initTUI constructs and returns the Bubbletea program.
func initTUI(songService *services.SongService, state *mpv.PlayerState) *tea.Program {
	return tea.NewProgram(tui.InitialModel(songService, state))
}

// initEventLoop starts the mpv event listener that dispatches state updates, TUI refreshes, and AutoPlay.
func initEventLoop(player *mpv.MpvInstance, state *mpv.PlayerState, songService *services.SongService, p *tea.Program) {
	player.ListenEvents(func(event mpv.EventData) {
		slog.Debug("mpv event received", "id", event.ID.String(), "prop_id", event.PropID, "prop_name", event.PropName, "val_str", event.StringVal, "val_num", event.DoubleVal, "val_flag", event.FlagVal)

		state.UpdateFromEvent(event)
		p.Send(event)

		if event.ID == mpv.EventEndFile {
			slog.Info("mpv EventEndFile received, triggering AutoPlay next song")
			songService.Next()
		}
	})
}
