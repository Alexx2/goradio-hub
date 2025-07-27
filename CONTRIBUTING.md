# Contributing to GoRadio Hub 🎵

Thank you for your interest in contributing to GoRadio Hub! We welcome contributions from developers of all skill levels.

## 🚀 Quick Start for Contributors

### Development Setup
```bash
# Fork and clone the repository
git clone https://github.com/yourusername/goradio.git
cd goradio

# Check dependencies
./scripts/check-dependencies.sh

# Install Go dependencies
go mod tidy

# Build and test
go build -o goradio
./goradio
```

## 🎯 Ways to Contribute

### 🎧 Radio Stations
- **Add new working stations** - Test URLs thoroughly before submitting
- **Fix broken streams** - Update URLs for stations that stopped working
- **Improve metadata** - Better descriptions and genre classifications
- **Regional stations** - Add local/regional radio stations

### 🎨 UI/UX Improvements
- **ASCII art** - Create new logos or improve existing ones
- **Color schemes** - Enhance terminal color adaptation
- **Layout improvements** - Better responsive design
- **Accessibility** - Support for different terminal sizes and capabilities

### 🔧 Features & Bug Fixes
- **Audio enhancements** - Better metadata parsing, volume controls
- **New features** - Favorites, playlists, recording capabilities
- **Performance** - Optimize memory usage and startup time
- **Cross-platform** - Improve Windows/macOS compatibility

### 📚 Documentation
- **README improvements** - Better installation guides, troubleshooting
- **Code documentation** - Add comments and examples
- **User guides** - Create tutorials and usage examples

## 📋 Development Guidelines

### Code Style
- Use `go fmt` for consistent formatting
- Follow [Effective Go](https://golang.org/doc/effective_go.html) conventions
- Add comments for exported functions and types
- Keep functions small and focused (< 50 lines when possible)
- Use meaningful variable and function names

### Testing
- Test all new radio stations manually before submitting
- Ensure cross-platform compatibility (Linux, macOS, Windows)
- Test in different terminal emulators when possible
- Verify ASCII art renders correctly

### Git Workflow
1. **Fork** the repository to your GitHub account
2. **Clone** your fork locally
3. **Create** a feature branch: `git checkout -b feature/awesome-feature`
4. **Make** your changes with clear, atomic commits
5. **Test** your changes thoroughly
6. **Push** to your fork: `git push origin feature/awesome-feature`
7. **Submit** a Pull Request with a clear description

### Commit Messages
Use conventional commit format:
```
feat: add new vaporwave radio station
fix: resolve mpv playback issue on macOS
docs: update installation instructions
style: improve ASCII art alignment
```

## 🎵 Adding Radio Stations

### Station Requirements
- **Working stream URL** - Must be accessible and stable
- **Proper format** - MP3, AAC, or OGG preferred
- **Legal compliance** - Only add legitimate, authorized streams
- **Quality** - Minimum 128kbps bitrate recommended

### Station Format
```go
{
    Name:        "Station Name",           // Keep under 20 characters
    URL:         "https://stream.url",     // Direct stream URL (not playlist)
    Genre:       "Genre",                  // Single word preferred
    Description: "Brief description",      // Under 50 characters
},
```

### Testing New Stations
```bash
# Test stream connectivity
curl -I --connect-timeout 5 "https://your.stream.url"

# Test with mpv
mpv --no-video --really-quiet "https://your.stream.url"
```

## 🐸 ASCII Art Guidelines

### Logo Requirements
- Use only standard ASCII characters for maximum compatibility
- Keep width under 50 characters for responsive layout
- Test in multiple terminal emulators
- Maintain alignment and symmetry
- Follow existing style patterns

### Testing ASCII Art
```bash
# Test logos in different scenarios
go run test_logos.go ui.go stations.go

# Test in minimal terminal
COLUMNS=80 LINES=24 ./goradio
```

## 🛠️ Project Structure

```
goradio/
├── main.go              # Application entry point & TUI logic
├── ui.go                # Interface rendering & ASCII art
├── player.go            # Audio playback with mpv
├── stations.go          # Radio station definitions
├── go.mod              # Go module dependencies
├── README.md           # Project documentation
├── CONTRIBUTING.md     # This file
├── LICENSE             # MIT license
├── .gitignore          # Git ignore rules
└── scripts/            # Utility scripts
    ├── check-dependencies.sh
    └── setup-github.sh
```

## 🔍 Code Review Process

### Pull Request Guidelines
- **Clear title** - Describe what the PR does
- **Detailed description** - Explain the changes and why
- **Screenshots** - Include terminal screenshots for UI changes
- **Testing notes** - Describe how you tested the changes
- **Breaking changes** - Clearly mark any breaking changes

### Review Criteria
- Code follows project conventions
- Changes are well-tested
- Documentation is updated if needed
- No unnecessary dependencies added
- Backward compatibility maintained

## 🐛 Bug Reports

Please include:
- **Go version**: `go version`
- **Operating system**: `uname -a` (Linux/macOS) or Windows version
- **Terminal emulator**: gnome-terminal, alacritty, etc.
- **GoRadio Hub version**: `./goradio --version`
- **Steps to reproduce**: Detailed reproduction steps
- **Expected behavior**: What should happen
- **Actual behavior**: What actually happens
- **Logs**: Any error messages or logs

## 💡 Feature Requests

When suggesting new features:
- **Use case**: Describe the problem you're trying to solve
- **Proposed solution**: Your idea for how to solve it
- **Alternatives**: Other solutions you've considered
- **Implementation**: Any thoughts on how to implement it

## 🎖️ Recognition

Contributors will be:
- Listed in the project's contributors section
- Credited in release notes for significant contributions
- Given credit in commit messages and PR descriptions

## 📜 Code of Conduct

### Our Standards
- **Be respectful** - Treat everyone with respect and kindness
- **Be inclusive** - Welcome people of all backgrounds and skill levels
- **Be collaborative** - Work together towards common goals
- **Be constructive** - Provide helpful feedback and suggestions
- **Have fun** - Remember we're building something awesome together! 🎵

### Unacceptable Behavior
- Harassment, discrimination, or hate speech
- Trolling, insulting, or personal attacks
- Sharing private information without permission
- Any behavior that would be inappropriate in a professional setting

## 📞 Getting Help

- **GitHub Issues**: For bug reports and feature requests
- **GitHub Discussions**: For questions and general discussion
- **Discord**: [Join our community](https://discord.gg/goradio) (if applicable)

## 🙏 Thank You

Thank you for contributing to GoRadio Hub! Every contribution helps make the project better for everyone.

---

**Made with 🎵 and 🐸 - FEELS GOOD MAN!**
