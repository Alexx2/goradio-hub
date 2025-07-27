# 🎵 GoRadio Hub 

[![Go Version](https://img.shields.io/badge/Go-1.19+-00ADD8?style=flat&logo=go)](https://golang.org/dl/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Release](https://img.shields.io/github/v/release/yourusername/goradio)](https://github.com/yourusername/goradio/releases)
[![Platform](https://img.shields.io/badge/platform-Linux%20%7C%20macOS%20%7C%20Windows-lightgrey)](https://github.com/yourusername/goradio)

**The Ultimate Terminal Radio Experience** - A beautiful TUI radio client with **FIRE ASCII art**, curated radio stations, and that **FEELS GOOD MAN** energy! 🐸

```
  ╔══════════════════════════════════════════════════╗
  ║           G O R A D I O   R A D I O              ║
  ║         ~ Terminal Music Player v1.0 ~           ║
  ╚══════════════════════════════════════════════════╝
```

## ✨ Features

🎨 **Beautiful TUI Interface**  
- **FIRE ASCII Art Logos**: Clean GoRadio Hub design + **Pepe the Frog** in bold green! 🐸
- **THREE Logo Modes**: GoRadio Hub → Pepe → Clean (press `l` to cycle)
- Terminal color adaptation with smooth wave animations
- Modern design with responsive layout

🎧 **Premium Audio Experience**
- **Real Audio Playback** with mpv (works with all stream formats!)
- Real-time song metadata display  
- Stream playback controls with play/pause/stop
- High-quality MP3/AAC/OGG stream support

🎵 **Curated Station Collection**
- **20 Verified Working Stations** across multiple genres
- Complete SomaFM lineup (ambient, downtempo, synthpop, psychedelic)
- Vaporwave & Aesthetic (Plaza Radio, synthwave vibes)
- Lofi Hip Hop & Chillhop for study/work sessions
- Drum & Bass (Bassdrive - world's largest DNB station)
- Multiple genres: Jazz, Soul, Folk, Metal, IDM, and more

⌨️ **Intuitive Controls**
- Vim-style navigation (`j`/`k` or arrow keys)
- Quick play/pause with `Enter`/`Space`
- Logo switching with `l` key
- Built-in help system (`?`)
- Smooth, responsive interface

## 🚀 Quick Start

### Prerequisites
- **Go 1.19+** - [Download](https://golang.org/dl/)
- **mpv** - Media player for audio playback

#### Install mpv:
```bash
# Ubuntu/Debian
sudo apt install mpv

# Arch Linux  
sudo pacman -S mpv

# Fedora
sudo dnf install mpv

# macOS
brew install mpv

# Windows (via Chocolatey)
choco install mpv
```

### Installation

#### Option 1: Download Release (Recommended)
```bash
# Download latest release
curl -LO https://github.com/yourusername/goradio/releases/latest/download/goradio-linux-amd64
chmod +x goradio-linux-amd64
./goradio-linux-amd64
```

#### Option 2: Build from Source
```bash
# Clone repository
git clone https://github.com/yourusername/goradio.git
cd goradio

# Check dependencies
./scripts/check-dependencies.sh

# Build
go mod tidy
go build -o goradio

# Run
./goradio
```

## 🎮 Usage

### Controls
| Key | Action |
|-----|--------|
| `↑`/`↓` or `j`/`k` | Navigate stations |
| `Enter` or `Space` | Play/Stop station |
| `l` | Cycle logo (GoRadio Hub → Pepe → None) |
| `?` | Toggle help screen |
| `q` or `Ctrl+C` | Quit |

### Interface Layout
- **Left Panel**: ASCII logo (switchable) + station browser with genre info
- **Right Panel**: Now playing info, station details, or help screen

### Logo Options
Press `l` to cycle through three epic logo modes:

1. **GoRadio Hub** - Clean styled ASCII art with music notes
2. **Pepe** - REAL Pepe the Frog face in bold green 🐸 ("FEELS GOOD MAN")
3. **None** - Minimalist interface without logo

## 🎵 Station Categories

### SomaFM Stations
- **Groove Salad** - Chilled ambient/downtempo beats
- **Drone Zone** - Deep ambient soundscapes for meditation  
- **Lush** - Sensuous vocals with electronic influence
- **Space Station** - Spaced-out ambient electronica
- **Vaporwaves** - Aesthetic vaporwave and future funk
- **Underground 80s** - Early UK synthpop and new wave
- **DEF CON Radio** - Music for hackers and makers
- **Secret Agent** - Spy jazz for your stylish life
- And more!

### Specialty Genres
- **Plaza Radio** - 24/7 vaporwave and aesthetic vibes 🌆
- **Bassdrive DNB** - World's largest drum & bass station 🥁
- **ChillHop Radio** - Perfect lofi hip hop for study/work 📚
- **Deep Space One** - Ambient electronic space music 🚀
- **The Trip** - Progressive rock and psychedelic experiments
- **Seven Inch Soul** - Vintage soul from original 45 RPM vinyl

## 🛠️ Development

### Building
```bash
# Clone and enter directory
git clone https://github.com/yourusername/goradio.git
cd goradio

# Install dependencies
go mod tidy

# Run tests
go test ./...

# Build
go build -o goradio

# Cross-compile for different platforms
GOOS=linux GOARCH=amd64 go build -o goradio-linux-amd64
GOOS=darwin GOARCH=amd64 go build -o goradio-darwin-amd64
GOOS=windows GOARCH=amd64 go build -o goradio-windows-amd64.exe
```

### Project Structure
```
goradio/
├── main.go           # Main application entry point
├── ui.go             # TUI interface and ASCII art
├── player.go         # Audio playback with mpv
├── stations.go       # Radio station definitions
├── go.mod           # Go module dependencies
├── .gitignore       # Git ignore rules
├── LICENSE          # MIT license
├── README.md        # This file
├── CONTRIBUTING.md  # Contribution guidelines
└── scripts/         # Utility scripts
    ├── check-dependencies.sh
    └── setup-github.sh
```

## 🤝 Contributing

We welcome contributions! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

### Quick Contribution Guide
1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

### Adding Radio Stations
To add new stations, edit `stations.go`:
```go
{
    Name:        "Your Station Name",
    URL:         "http://your.stream.url/stream.mp3",
    Genre:       "Your Genre",
    Description: "Brief description of the station",
},
```

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- **SomaFM** for providing excellent commercial-free radio streams
- **The Go Community** for amazing libraries and tools
- **Charm** (Bubble Tea/Lip Gloss) for beautiful TUI framework
- **MPV** for reliable cross-platform audio playback
- **All the radio stations** for keeping the music flowing 24/7

## 🐸 Fun Facts

- Press `l` to see **Pepe the Frog** ASCII art in your terminal!
- GoRadio adapts to your terminal's color scheme automatically
- All 20 stations are verified working and regularly tested
- Built with ❤️ for terminal music lovers worldwide

---

**Made with 🎵 and 🐸 - FEELS GOOD MAN!**

[![Star History Chart](https://api.star-history.com/svg?repos=yourusername/goradio&type=Date)](https://star-history.com/#yourusername/goradio&Date)
