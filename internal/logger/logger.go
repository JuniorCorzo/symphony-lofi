package logger

import (
	"log/slog"
	"os"

	"github.com/JuniorCorzo/symphony-lofi/internal/config"
)

var Log *slog.Logger

// InitLogger sets up a professional structured logger writing to ~/.config/symphony-lofi/app.log
func InitLogger() (*os.File, error) {
	logPath := config.GetFilePath("app.log")
	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}

	handler := slog.NewTextHandler(file, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})

	Log = slog.New(handler)
	slog.SetDefault(Log)

	Log.Info("Symphony-Lofi logger initialized", "log_file", logPath)
	return file, nil
}
