# Symphony Lofi 🎵

A lightweight, terminal-based Lofi music player written in **Go**, powered by **libmpv**, **Bubbletea**, **Lipgloss**, and **SQLite**. Styled with the **Catppuccin Mocha** color palette.

![Theme](https://img.shields.io/badge/Theme-Catppuccin--Mocha-cba6f7?style=for-the-badge)
![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![License](https://img.shields.io/badge/License-MIT-blue?style=for-the-badge)

---

## 🎨 Preview

```text
╭──────────────────────────────────────────────────╮
│  🎵 SYMPHONY LOFI                                │
│                                                  │
│  Artist:   Purrple Cat                           │
│  Track:    Equinox                               │
│  Time:     01:45 / 03:20                         │
│  Status:   ▶ PLAYING                             │
│                                                  │
│  [Space] Play/Pause • [n] Next • [q] Quit        │
╰──────────────────────────────────────────────────╯
```

---

## ✨ Features

- **Catppuccin Mocha TUI**: Styled terminal UI using `Bubbletea` and `Lipgloss`.
- **CGO libmpv Engine**: Hardware-accelerated audio playback with `libmpv` bindings.
- **AutoPlay & Pre-buffering**: Asynchronous track pre-buffering (`loadfile append`) for instant playback transitions.
- **Embedded SQLite & Remote ETag Sync**: Stores tracks in local SQLite (`~/.config/symphony-lofi/songs.db`) with HTTP ETag (304 Not Modified) synchronization support.
- **Automatic Browser Cookies**: Auto-detects installed browsers (`Firefox`, `Brave`, `Chrome`, `Chromium`) to bypass YouTube rate limits (HTTP 429).
- **Structured Slog Logging**: App lifecycle logs written cleanly to `~/.config/symphony-lofi/app.log`.

---

## 🎹 Keyboard Controls

| Key | Action |
| --- | --- |
| `Space` | Toggle Play / Pause |
| `n` | Skip to Next Song |
| `q` / `Ctrl+C` | Quit Application |

---

## 🚀 Installation & Usage

### Prerequisites

Ensure you have `mpv` C development libraries installed:

- **Fedora**: `sudo dnf install mpv-devel gcc`
- **Ubuntu / Debian**: `sudo apt install libmpv-dev gcc`
- **Arch Linux**: `sudo pacman -S mpv gcc`

### Run Locally

```bash
# Clone repository
git clone https://github.com/JuniorCorzo/symphony-lofi.git
cd symphony-lofi

# Run application
go run ./cmd/symphony-lofi/
```

---

## 💬 Credits & Acknowledgments

Special thanks to **[Purrple Cat](https://purrplecat.com/)** for creating amazing royalty-free Lofi music that powers focus and chill vibes.

- **Music by Purrple Cat**: [purrplecat.com](https://purrplecat.com/)
- **Spotify**: [Purrple Cat on Spotify](https://open.spotify.com/artist/73aKnLT4O8G2pBEIBl6A3E)
- **YouTube**: [Purrple Cat Music Channel](https://www.youtube.com/c/PurrpleCatMusic)

---

## 📄 License

This project is open-source under the [MIT License](LICENSE).
