package mpv

/*
#include <mpv/client.h>
#include <stdlib.h>

static inline int execute_mpv_command(mpv_handle *ctx, const char **args) {
	return mpv_command(ctx, args);
}
*/
import "C"

import (
	"fmt"
	"unsafe"
)

func (mpv *MpvInstance) Command(args ...string) error {
	handle := mpv.handle
	if handle == nil {
		return fmt.Errorf("mpv: contexto no inicializado")
	}

	cArgs := make([]*C.char, len(args)+1)
	for i, arg := range args {
		cArgs[i] = C.CString(arg)
		defer C.free(unsafe.Pointer(cArgs[i]))
	}

	cArgs[len(args)] = nil

	status := C.execute_mpv_command(handle, (**C.char)(unsafe.Pointer(&cArgs[0])))
	if status < 0 {
		return fmt.Errorf("mpv error: %d", status)
	}

	return nil
}

func (mpv *MpvInstance) GetPropertyString(property string) (string, error) {
	handle := mpv.handle

	cProperty := C.CString(property)
	defer C.free(unsafe.Pointer(cProperty))

	result := C.mpv_get_property_string(handle, cProperty)

	if result == nil {
		return "", fmt.Errorf("mpv: no se pudo obtener la propiedad '%s'", property)
	}
	defer C.mpv_free(unsafe.Pointer(result))

	return C.GoString(result), nil
}

func (mpv *MpvInstance) ObserverProperty(propertyID PropertyID) {
	handle := mpv.handle

	property := PropertyCatalog[propertyID]
	cPropertyName := C.CString(string(property.Name))

	defer C.free(unsafe.Pointer(cPropertyName))

	C.mpv_observe_property(
		handle,
		C.uint64_t(propertyID),
		cPropertyName,
		property.Format)
}

func (mpv *MpvInstance) LoadSong(url string) error {
	return mpv.Command("loadfile", url)
}

func (mpv *MpvInstance) PreloadSong(url string) error {
	return mpv.Command("loadfile", url, "append")
}

func (mpv *MpvInstance) TogglePause() error {
	return mpv.Command("cycle", "pause")
}
