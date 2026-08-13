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

### 🐧 Linux

#### 1. Install Prerequisites (`libmpv` & C Compiler)

- **Fedora / RHEL**:
  ```bash
  sudo dnf install mpv-devel gcc
  ```
- **Ubuntu / Debian / Pop!_OS**:
  ```bash
  sudo apt update && sudo apt install libmpv-dev gcc
  ```
- **Arch Linux / Manjaro**:
  ```bash
  sudo pacman -S mpv gcc
  ```

#### 2. Build & Install to PATH

```bash
# Clone the repository
git clone https://github.com/JuniorCorzo/symphony-lofi.git
cd symphony-lofi

# Compile and install to your local user bin (~/.local/bin)
go build -ldflags="-s -w" -o ~/.local/bin/symphony-lofi ./cmd/symphony-lofi

# Run anywhere
symphony-lofi
```

---

### 🪟 Windows

Because `symphony-lofi` interfaces with `libmpv` through CGO, you need a C compiler (`gcc`) and `mpv` header/dll libraries provided via **MSYS2**.

#### 1. Install Prerequisites via MSYS2

1. Download and install [MSYS2](https://www.msys2.org/).
2. Open the **MSYS2 MINGW64** shell and run:
   ```bash
   pacman -S mingw-w64-x86_64-gcc mingw-w64-x86_64-pkg-config mingw-w64-x86_64-mpv
   ```
3. Add `C:\msys64\mingw64\bin` to your Windows `PATH` environment variable so Go can locate `gcc` and `mpv-2.dll`.

#### 2. Build & Run (PowerShell or CMD)

```powershell
# Clone the repository
git clone https://github.com/JuniorCorzo/symphony-lofi.git
cd symphony-lofi

# Compile executable
go build -ldflags="-s -w" -o symphony-lofi.exe ./cmd/symphony-lofi

# Run player
.\symphony-lofi.exe
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
