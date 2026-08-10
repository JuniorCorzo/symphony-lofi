package services

import "github.com/JuniorCorzo/symphony-lofi/internal/db"

type Playlist struct {
	songs        []db.Song
	currentIndex int
}

func NewPlaylist(songs []db.Song) *Playlist {
	return &Playlist{
		songs:        songs,
		currentIndex: 0,
	}
}

func (playlist *Playlist) Current() (db.Song, bool) {
	if len(playlist.songs) == 0 {
		return db.Song{}, false
	}
	return playlist.songs[playlist.currentIndex], true
}

func (playlist *Playlist) PeekNext() (db.Song, bool) {
	if len(playlist.songs) == 0 {
		return db.Song{}, false
	}
	nextIdx := (playlist.currentIndex + 1) % len(playlist.songs)
	return playlist.songs[nextIdx], true
}

func (playlist *Playlist) Next() (db.Song, bool) {
	if len(playlist.songs) == 0 {
		return db.Song{}, false
	}

	playlist.currentIndex = (playlist.currentIndex + 1) % len(playlist.songs)

	return playlist.songs[playlist.currentIndex], true
}
