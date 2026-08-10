// Package tui
package tui

import (
	"fmt"

	"github.com/JuniorCorzo/symphony-lofi/internal/mpv"
	"github.com/JuniorCorzo/symphony-lofi/internal/services"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	// Catppuccin Mocha Palette & Styles
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#1e1e2e")). // Base
			Background(lipgloss.Color("#cba6f7")). // Mauve
			Padding(0, 1)

	labelStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#a6adc8")). // Subtext0
			Width(9)

	artistStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#a6e3a1")). // Green
			Bold(true)

	trackStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#f5c2e7")). // Pink
			Bold(true)

	timeStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#94e2d5")) // Teal

	statusStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#f9e2af")). // Yellow
			Bold(true)

	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#6c7086")) // Overlay0

	boxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#cba6f7")). // Mauve Border
			Padding(1, 2).
			Width(52)
)

type model struct {
	state       *mpv.PlayerState
	songService *services.SongService
}

func InitialModel(
	songService *services.SongService,
	state *mpv.PlayerState,
) model {
	return model{
		songService: songService,
		state:       state,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case mpv.EventData:
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "n":
			m.songService.Next()
		case " ":
			m.songService.TogglePause()
		}
	}

	return m, nil
}

func formatTime(seconds float64) string {
	if seconds < 0 {
		seconds = 0
	}
	m := int(seconds) / 60
	s := int(seconds) % 60
	return fmt.Sprintf("%02d:%02d", m, s)
}

func (m model) View() string {
	snap := m.state.Snapshot()
	artist := snap.Artist
	if artist == "" {
		artist = "Unknown Artist"
	}

	title := snap.SongTitle
	if title == "" {
		title = "Loading Track..."
	}

	status := "▶ PLAYING"
	if snap.IsPaused {
		status = "⏸ PAUSED"
	}

	posStr := formatTime(snap.Position)
	durStr := formatTime(snap.Duration)

	header := titleStyle.Render("🎵 SYMPHONY LOFI")
	statusText := statusStyle.Render(status)
	artistText := labelStyle.Render("Artist:") + artistStyle.Render(artist)
	trackText := labelStyle.Render("Track:") + trackStyle.Render(title)
	timeText := labelStyle.Render("Time:") + timeStyle.Render(fmt.Sprintf("%s / %s", posStr, durStr))
	statusRow := labelStyle.Render("Status:") + statusText

	content := fmt.Sprintf("%s\n\n%s\n%s\n%s\n%s\n\n%s",
		header,
		artistText,
		trackText,
		timeText,
		statusRow,
		helpStyle.Render("[Space] Play/Pause • [n] Next • [q] Quit"),
	)

	return "\n" + boxStyle.Render(content) + "\n"
}
