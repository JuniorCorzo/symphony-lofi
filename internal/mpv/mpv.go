// Package mpv provides bindings and control for the mpv media player.
package mpv

/*
#cgo LDFLAGS: -lmpv
#include <mpv/client.h>
#include <stdlib.h>

static inline int execute_mpv_command(mpv_handle *ctx, const char **args) {
	return mpv_command(ctx, args);
}
*/
import "C"

import (
	"os"
	"path/filepath"
	"sync"
	"unsafe"
)

type MpvInstance struct {
	handle *C.mpv_handle
}

var (
	instance *MpvInstance
	once     sync.Once
)

func setMpvOption(handle *C.mpv_handle, key, val string) {
	cKey := C.CString(key)
	cVal := C.CString(val)
	defer C.free(unsafe.Pointer(cKey))
	defer C.free(unsafe.Pointer(cVal))
	C.mpv_set_option_string(handle, cKey, cVal)
}

func detectBrowser() string {
	if b := os.Getenv("LOFI_BROWSER"); b != "" {
		return b
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "firefox"
	}

	// Detecta automáticamente el navegador instalado según los directorios del usuario
	if _, err := os.Stat(filepath.Join(homeDir, ".mozilla", "firefox")); err == nil {
		return "firefox"
	}
	if _, err := os.Stat(filepath.Join(homeDir, ".config", "BraveSoftware", "Brave-Browser")); err == nil {
		return "brave"
	}
	if _, err := os.Stat(filepath.Join(homeDir, ".config", "google-chrome")); err == nil {
		return "chrome"
	}
	if _, err := os.Stat(filepath.Join(homeDir, ".config", "chromium")); err == nil {
		return "chromium"
	}

	return "firefox"
}

func GetInstance() *MpvInstance {
	once.Do(func() {
		handle := C.mpv_create()
		if handle != nil {
			setMpvOption(handle, "vo", "null")
			setMpvOption(handle, "vid", "no")

			browser := detectBrowser()
			setMpvOption(handle, "ytdl-raw-options", "cookies-from-browser="+browser)

			C.mpv_initialize(handle)
		}
		instance = &MpvInstance{
			handle: handle,
		}
	})

	return instance
}
