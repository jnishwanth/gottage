# 🏡 Gottage — Virtual Cottage TUI

**Gottage** is a cozy Virtual Cottage focus app for your terminal built in Go using [Bubbletea](https://github.com/charmbracelet/bubbletea), [Lipgloss](https://github.com/charmbracelet/lipgloss), and [CLIAMP](https://github.com/cliamp).

![Cozy Bit-Art](https://img.shields.io/badge/Art-TrueColor%2024bit-orange)
![Bubbletea](https://img.shields.io/badge/TUI-Bubbletea-blue)
![License](https://img.shields.io/badge/License-MIT-green)

---

## ✨ Features

- **🏡 Multi-Scene Focus Bit-Art Engine**:
  - High-density 24-bit TrueColor pixel canvas using Unicode half-blocks (`▀`), supporting **74x68 pixel resolution** in terminal!
  - **4 Science-Backed Research Scenes** (Cycle with `v`):
    1. **Rainy City Lofi Cafe** 🏙️ (*Lofi Girl / LifeAt*): High-rise window looking onto a rain-swept neon skyline with coffee steam.
    2. **Forest Sanctuary Cabin** 🌲 (*Spirit City / Biophilic Nature*): Deep pine forest window with swaying trees and sleeping fox/cat.
    3. **Classic Cozy Cottage** 🏡 (*Original Gottage preserved!*): Warm brick fireplace, desk lamp, rug, and arched window.
    4. **Celestial Observatory** 🌌 (*Astronomical Study*): Massive circular skylight facing constellations and glowing crescent moon.
  - **Dynamic Sky & Time-of-Day**: Dawn 🌅, Day ☀️, Sunset 🌇, Night 🌙, or Auto-sync with your real local machine clock ⏰.
  - **Dynamic Weather System**: Falling Raindrops 🌧️, Swirling Snow ❄️, Thunderstorm Lightning Flashes 🌩️, and Cozy Fog 🌫️.
  - **Interactive Lighting Engine**: Desk Lamp cone glow 💡, Fireplace radial warmth 🔥, and Lantern glow.

- **🎧 CLIAMP Music Player & Ambient Soundscape**:
  - Full integration with **CLIAMP** retro terminal music player!
  - Query real-time playing track, artist, volume, and playback state.
  - Curated 24/7 Lofi Radio presets (Lofi Girl, Coffee Shop Beats, Chiptune & 8-Bit Chill, Rainy Day Piano).
  - Built-in procedural ambient sound engine for Rain 🌧️, Fireplace Crackle 🔥, and Wind 🍃.

- **⏱️ Cozy Productivity Suite**:
  - **Focus Timer**: Pomodoro Countdown (25m Focus, 5m Short Break, 15m Long Break) with Big ASCII Digital Clock, progress bar, and session counter.
  - **Task Notebook**: Interactive To-Do list to write down and check off your daily goals 📝.
  - **Ambiance Control Panel**: TUI dashboard to adjust lighting, weather, time, cat breed, and sound mix.

---

## 🚀 Quick Start

### Prerequisites
- Go 1.20+ installed.
- (Optional) [cliamp](https://github.com/cliamp) for lofi music playback.

### Running Gottage

```bash
# Clone & run directly
go run main.go

# Or build the binary
go build -o gottage main.go
./gottage
```

### Command Line Options

```bash
# Render a single frame preview to stdout
./gottage --demo

# Render one named scene (accepts a unique substring, so the emoji is optional)
./gottage --demo --scene "Kirokaze Cozy Room"

# Launch with specific initial weather
./gottage --weather rain

# Launch with specific time of day
./gottage --time night
```

---

## ⌨️ Keybindings

| Key | Action |
| --- | --- |
| `1` | Switch to **Cottage Bit-Art** view |
| `2` | Switch to **Focus Timer** |
| `3` | Switch to **Task Notebook** |
| `4` | Switch to **Ambiance & Audio** |
| `Tab` | Cycle through views |
| `v` | Cycle Scene Preset (Rainy City, Forest Cabin, Classic, Observatory) |
| `l` | Toggle Desk Lamp (ON / OFF) |
| `f` | Toggle Fireplace (ON / OFF) |
| `s` | Toggle Fairy Lights (ON / OFF) |
| `w` | Cycle Weather (Clear, Rain, Snow, Storm, Fog) |
| `t` | Cycle Time of Day (Dawn, Day, Sunset, Night, Auto) |
| `c` | Cycle Cat Breed (Ginger Tabby, Tuxedo, Calico, Black) |
| `Space` | Play / Pause CLIAMP music player |
| `n` / `p` | Cycle Lofi Radio stream presets |
| `r` | Toggle Rain ambient sound |
| `m` | Toggle Fireplace ambient sound |
| `x` | Start / Pause Pomodoro Timer |
| `z` | Reset Pomodoro Timer |
| `a` | Add new Task (in Tasks View) |
| `Enter` | Toggle Task completion |
| `d` | Delete Task |
| `?` | Open / Close Keyboard Shortcuts Help Modal |
| `q` | Quit Gottage |

---

## 🏗️ Architecture

```
gottage/
├── main.go               # Application entry point & CLI flags
├── pkg/
│   ├── art/              # Dynamic Bit-Art Canvas & Scene Engine
│   │   ├── canvas.go     # 24-bit TrueColor ANSI half-block pixel canvas
│   │   └── scene.go      # Dynamic Cottage room, weather, fire, cat & lighting
│   ├── audio/            # Audio & CLIAMP Integration
│   │   ├── cliamp.go     # CLIAMP daemon controller & preset radio streams
│   │   └── ambient.go    # Native procedural ambient sound generator (Rain/Fire/Wind)
│   ├── features/         # Focus Timer & Task Notebook
│   │   ├── timer.go      # Pomodoro timer logic
│   │   └── tasks.go      # Task list manager
│   └── ui/               # Bubbletea TUI Model, Views & Lipgloss Styles
│       ├── model.go      # Bubbletea Model & Event handling
│       ├── view.go       # View layout renderer & sidebar
│       └── styles.go     # Color palettes & Lipgloss design tokens
```

---

## 📜 License

MIT License © 2026
