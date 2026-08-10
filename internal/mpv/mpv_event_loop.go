package mpv

/*
#include <mpv/client.h>
*/
import "C"

type EventData struct {
	ID        EventID
	PropID    PropertyID
	PropName  string
	StringVal string
	DoubleVal float64
	FlagVal   bool
}

func (mpv *MpvInstance) ListenEvents(handler func(EventData)) {
	eventsToObserve := []PropertyID{
		PropertyIDTimePos,
		PropertyIDDuration,
		PropertyIDMediaTitle,
		PropertyIDMetadataArtist,
		PropertyIDVolume,
		PropertyIDPause,
	}

	for _, propID := range eventsToObserve {
		mpv.ObserverProperty(propID)
	}

	for {
		cEvent := C.mpv_wait_event(mpv.handle, -1)
		if cEvent == nil || cEvent.event_id == C.MPV_EVENT_NONE {
			continue
		}
		if cEvent.event_id == C.MPV_EVENT_SHUTDOWN {
			break
		}

		eventData := EventData{
			ID: EventID(cEvent.event_id),
		}

		if cEvent.event_id == C.MPV_EVENT_PROPERTY_CHANGE {
			cProp := (*C.mpv_event_property)(cEvent.data)
			if cProp == nil {
				continue
			}

			eventData.PropID = PropertyID(cEvent.reply_userdata)
			eventData.PropName = C.GoString(cProp.name)

			switch cProp.format {
			case C.MPV_FORMAT_STRING:
				if cProp.data != nil {
					eventData.StringVal = C.GoString(*(**C.char)(cProp.data))
				}
			case C.MPV_FORMAT_DOUBLE:
				if cProp.data != nil {
					eventData.DoubleVal = float64(*(*C.double)(cProp.data))
				}
			case C.MPV_FORMAT_FLAG:
				if cProp.data != nil {
					eventData.FlagVal = int(*(*C.int)(cProp.data)) != 0
				}

			}
		}

		if handler != nil {
			handler(eventData)
		}
	}
}
