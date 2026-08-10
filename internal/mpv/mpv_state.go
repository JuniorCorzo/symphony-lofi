package mpv

import (
	"strings"
	"sync"
)

type PlayerState struct {
	mu        sync.RWMutex
	SongTitle string  `json:"song_title"`
	Artist    string  `json:"artist"`
	Duration  float64 `json:"duration"`
	Position  float64 `json:"position"`
	Volume    float64 `json:"volume"`
	IsPaused  bool    `json:"is_paused"`
}

func NewPlayerState() *PlayerState {
	return &PlayerState{
		Volume: 100,
	}
}

func (s *PlayerState) SetLoading() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.SongTitle = "Cargando canción..."
	s.Artist = "Cargando..."
}

func (s *PlayerState) UpdateFromEvent(event EventData) {
	s.mu.Lock()
	defer s.mu.Unlock()

	switch event.PropID {
	case PropertyIDTimePos:
		s.Position = event.DoubleVal
	case PropertyIDDuration:
		s.Duration = event.DoubleVal
	case PropertyIDMediaTitle:
		title := event.StringVal
		if strings.HasPrefix(title, "http://") || strings.HasPrefix(title, "https://") || strings.Contains(title, "watch?v=") {
			s.SongTitle = "Cargando canción..."
			s.Artist = "Cargando..."
		} else if parts := strings.SplitN(title, " - ", 2); len(parts) == 2 {
			s.Artist = parts[0]
			s.SongTitle = parts[1]
		} else {
			s.SongTitle = title
		}
	case PropertyIDMetadataArtist:
		if event.StringVal != "" && !strings.HasPrefix(event.StringVal, "http") {
			s.Artist = event.StringVal
		}
	case PropertyIDVolume:
		s.Volume = event.DoubleVal
	case PropertyIDPause:
		s.IsPaused = event.FlagVal
	}
}

func (s *PlayerState) Snapshot() PlayerState {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return PlayerState{
		SongTitle: s.SongTitle,
		Artist:    s.Artist,
		Duration:  s.Duration,
		Position:  s.Position,
		Volume:    s.Volume,
		IsPaused:  s.IsPaused,
	}
}
