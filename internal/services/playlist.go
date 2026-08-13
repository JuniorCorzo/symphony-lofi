// Package services implements application domain services and playlist management.
package services

import "github.com/JuniorCorzo/symphony-lofi/internal/db"

// Playlist manages the track queue and playback position.
type Playlist struct {
	songs        []db.Song
	currentIndex int
}

// NewPlaylist constructs a new Playlist instance with a slice of songs.
func NewPlaylist(songs []db.Song) *Playlist {
	return &Playlist{
		songs:        songs,
		currentIndex: 0,
	}
}

// Current returns the track at the current position.
func (playlist *Playlist) Current() (db.Song, bool) {
	if len(playlist.songs) == 0 {
		return db.Song{}, false
	}
	return playlist.songs[playlist.currentIndex], true
}

// PeekNext returns the next track in the queue without advancing the position.
func (playlist *Playlist) PeekNext() (db.Song, bool) {
	if len(playlist.songs) == 0 {
		return db.Song{}, false
	}
	nextIdx := (playlist.currentIndex + 1) % len(playlist.songs)
	return playlist.songs[nextIdx], true
}

// Next advances to the next track in circular modulo order and returns it.
func (playlist *Playlist) Next() (db.Song, bool) {
	if len(playlist.songs) == 0 {
		return db.Song{}, false
	}

	playlist.currentIndex = (playlist.currentIndex + 1) % len(playlist.songs)

	return playlist.songs[playlist.currentIndex], true
}
