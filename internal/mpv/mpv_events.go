package mpv

/*
#include <mpv/client.h>
*/
import "C"

type EventID int

const (
	EventNone             EventID = C.MPV_EVENT_NONE
	EventShutdown         EventID = C.MPV_EVENT_SHUTDOWN
	EventLogMessage       EventID = C.MPV_EVENT_LOG_MESSAGE
	EventGetPropertyReply EventID = C.MPV_EVENT_GET_PROPERTY_REPLY
	EventSetPropertyReply EventID = C.MPV_EVENT_SET_PROPERTY_REPLY
	EventCommandReply     EventID = C.MPV_EVENT_COMMAND_REPLY
	EventStartFile        EventID = C.MPV_EVENT_START_FILE
	EventEndFile          EventID = C.MPV_EVENT_END_FILE
	EventFileLoaded       EventID = C.MPV_EVENT_FILE_LOADED
	EventClientMessage    EventID = C.MPV_EVENT_CLIENT_MESSAGE
	EventVideoReconfig    EventID = C.MPV_EVENT_VIDEO_RECONFIG
	EventAudioReconfig    EventID = C.MPV_EVENT_AUDIO_RECONFIG
	EventSeek             EventID = C.MPV_EVENT_SEEK
	EventPlaybackRestart  EventID = C.MPV_EVENT_PLAYBACK_RESTART
	EventPropertyChange   EventID = C.MPV_EVENT_PROPERTY_CHANGE
	EventQueueOverflow   EventID = C.MPV_EVENT_QUEUE_OVERFLOW
	EventHook             EventID = C.MPV_EVENT_HOOK
)

func (e EventID) String() string {
	switch e {
	case EventNone:
		return "MPV_EVENT_NONE"
	case EventShutdown:
		return "MPV_EVENT_SHUTDOWN"
	case EventLogMessage:
		return "MPV_EVENT_LOG_MESSAGE"
	case EventGetPropertyReply:
		return "MPV_EVENT_GET_PROPERTY_REPLY"
	case EventSetPropertyReply:
		return "MPV_EVENT_SET_PROPERTY_REPLY"
	case EventCommandReply:
		return "MPV_EVENT_COMMAND_REPLY"
	case EventStartFile:
		return "MPV_EVENT_START_FILE"
	case EventEndFile:
		return "MPV_EVENT_END_FILE"
	case EventFileLoaded:
		return "MPV_EVENT_FILE_LOADED"
	case EventClientMessage:
		return "MPV_EVENT_CLIENT_MESSAGE"
	case EventVideoReconfig:
		return "MPV_EVENT_VIDEO_RECONFIG"
	case EventAudioReconfig:
		return "MPV_EVENT_AUDIO_RECONFIG"
	case EventSeek:
		return "MPV_EVENT_SEEK"
	case EventPlaybackRestart:
		return "MPV_EVENT_PLAYBACK_RESTART"
	case EventPropertyChange:
		return "MPV_EVENT_PROPERTY_CHANGE"
	case EventQueueOverflow:
		return "MPV_EVENT_QUEUE_OVERFLOW"
	case EventHook:
		return "MPV_EVENT_HOOK"
	default:
		return "UNKNOWN_EVENT"
	}
}
