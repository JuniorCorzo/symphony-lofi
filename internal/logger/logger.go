package logger

import (
	"log/slog"
	"os"
	"path/filepath"
)

var Log *slog.Logger

// InitLogger sets up a professional structured logger writing to ~/.config/symphony-lofi/app.log
func InitLogger() (*os.File, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		configDir = os.TempDir()
	}

	appDir := filepath.Join(configDir, "symphony-lofi")
	if err := os.MkdirAll(appDir, 0755); err != nil {
		return nil, err
	}

	logPath := filepath.Join(appDir, "app.log")
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
