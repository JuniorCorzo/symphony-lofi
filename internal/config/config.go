// Package config provides application directory and path resolution utilities.
package config

import (
	"os"
	"path/filepath"
	"sync"
)

var (
	appDir string
	once   sync.Once
)

// GetAppDir returns the absolute path to the application configuration directory (~/.config/symphony-lofi).
// It thread-safely creates the directory if it does not already exist.
func GetAppDir() string {
	once.Do(func() {
		configDir, err := os.UserConfigDir()
		if err != nil {
			configDir = os.TempDir()
		}
		appDir = filepath.Join(configDir, "symphony-lofi")
		_ = os.MkdirAll(appDir, 0755)
	})
	return appDir
}

// GetFilePath returns the absolute path for a specific file inside the application configuration directory.
func GetFilePath(filename string) string {
	return filepath.Join(GetAppDir(), filename)
}
