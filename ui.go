package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// LogoType represents different logo options
type LogoType int

const (
	LogoOriginal LogoType = iota
	LogoPepe
	LogoNone
)

func (l LogoType) String() string {
	switch l {
	case LogoOriginal:
		return "GoRadio Hub"
	case LogoPepe:
		return "Pepe"
	case LogoNone:
		return "None"
	default:
		return "GoRadio Hub"
	}
}

// ASCII art for GoRadio logo - Compact and clean
const GoRadioLogo = `
  ╔══════════════════════════════════════════════════╗
  ║           G O R A D I O   H U B                  ║
  ║         ~ Terminal Music Player v1.0 ~           ║
  ╚══════════════════════════════════════════════════╝
`

// ASCII art for Pepe the Frog - Perfectly aligned
const PepeLogo = `
      ╔═══════════════════════════════════╗
      ║          ,-._____.-.              ║
      ║         /  o     o  \             ║
      ║        |       >      |           ║
      ║        |   \_______/   |           ║
      ║         \             /           ║
      ║          '-.______.-'             ║
      ║                                   ║
      ║   🐸 F E E L S  G O O D  M A N   ║
      ║        ~ Bringing Da Vibes ~      ║
      ╚═══════════════════════════════════╝
`

// Style definitions
var (
	// Base colors that adapt to terminal
	primaryColor   = lipgloss.AdaptiveColor{Light: "#874BFE", Dark: "#7D56F4"}
	secondaryColor = lipgloss.AdaptiveColor{Light: "#43BF6D", Dark: "#73F59F"}
	accentColor    = lipgloss.AdaptiveColor{Light: "#F25D94", Dark: "#F25D94"}
	textColor      = lipgloss.AdaptiveColor{Light: "#0F0F0F", Dark: "#FAFAFA"}
	mutedColor     = lipgloss.AdaptiveColor{Light: "#6B7280", Dark: "#9CA3AF"}
	borderColor    = lipgloss.AdaptiveColor{Light: "#D1D5DB", Dark: "#374151"}

	// Logo styles
	logoStyle = lipgloss.NewStyle().
			Foreground(primaryColor).
			Bold(true).
			Align(lipgloss.Center).
			Margin(1, 0, 2, 0)

	// Pepe style - bold green
	pepeStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#00FF00")).
			Bold(true).
			Align(lipgloss.Center).
			Margin(1, 0, 2, 0)

	// Title styles
	titleStyle = lipgloss.NewStyle().
			Foreground(primaryColor).
			Bold(true).
			Align(lipgloss.Center).
			Margin(0, 0, 1, 0)

	subtitleStyle = lipgloss.NewStyle().
			Foreground(mutedColor).
			Italic(true).
			Align(lipgloss.Center).
			Margin(0, 0, 2, 0)

	// Station list styles
	selectedItemStyle = lipgloss.NewStyle().
				Foreground(textColor).
				Background(primaryColor).
				Bold(true).
				Padding(0, 1)

	normalItemStyle = lipgloss.NewStyle().
			Foreground(textColor).
			Padding(0, 1)

	// Info styles
	nowPlayingStyle = lipgloss.NewStyle().
			Foreground(secondaryColor).
			Bold(true).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(borderColor).
			Padding(1, 2).
			Margin(1, 0)

	stationInfoStyle = lipgloss.NewStyle().
				Foreground(accentColor).
				Bold(true).
				Margin(0, 0, 1, 0)

	// Status styles
	statusStyle = lipgloss.NewStyle().
			Foreground(mutedColor).
			Italic(true).
			Margin(1, 0)

	// Help text
	helpStyle = lipgloss.NewStyle().
			Foreground(mutedColor).
			Margin(2, 0, 0, 0)

	// Container styles
	containerStyle = lipgloss.NewStyle().
			Padding(1, 2).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(borderColor)

	// Panel styles for layout - Much wider for better spacing
	leftPanelStyle = lipgloss.NewStyle().
			Width(60).
			Height(30).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(borderColor).
			Padding(1, 2)

	rightPanelStyle = lipgloss.NewStyle().
			Width(70).
			Height(30).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(borderColor).
			Padding(1, 2)
)

// RenderLogo renders the selected ASCII logo
func RenderLogo(logoType LogoType) string {
	switch logoType {
	case LogoOriginal:
		return logoStyle.Render(GoRadioLogo)
	case LogoPepe:
		return pepeStyle.Render(PepeLogo)
	case LogoNone:
		return ""
	default:
		return logoStyle.Render(GoRadioLogo)
	}
}

// RenderLogoSelector renders the logo selection info
func RenderLogoSelector(currentLogo LogoType) string {
	selectorText := fmt.Sprintf("Logo: %s (Press 'l' to cycle)", currentLogo.String())
	return lipgloss.NewStyle().
		Foreground(mutedColor).
		Italic(true).
		Align(lipgloss.Center).
		Margin(0, 0, 1, 0).
		Render(selectorText)
}

// RenderTitle renders a styled title
func RenderTitle(text string) string {
	return titleStyle.Render(text)
}

// RenderSubtitle renders a styled subtitle
func RenderSubtitle(text string) string {
	return subtitleStyle.Render(text)
}

// RenderStationList renders the station selection list
func RenderStationList(stations []RadioStation, selected int, startIdx int, visibleCount int) string {
	var items []string
	
	// Calculate visible range
	endIdx := startIdx + visibleCount
	if endIdx > len(stations) {
		endIdx = len(stations)
	}
	
	for i := startIdx; i < endIdx; i++ {
		station := stations[i]
		prefix := "  "
		
		// Add numbering
		line := fmt.Sprintf("%s%d. %s", prefix, i+1, station.Name)
		
		// Add genre info
		line += fmt.Sprintf(" (%s)", station.Genre)
		
		if i == selected {
			line = selectedItemStyle.Render(line)
		} else {
			line = normalItemStyle.Render(line)
		}
		items = append(items, line)
	}
	
	return strings.Join(items, "\n")
}

// RenderNowPlaying renders the currently playing station info
func RenderNowPlaying(station *RadioStation, song string, status string) string {
	if station == nil {
		return statusStyle.Render("No station selected")
	}
	
	content := []string{
		stationInfoStyle.Render("♪ " + station.Name),
		fmt.Sprintf("Genre: %s", station.Genre),
		fmt.Sprintf("Status: %s", status),
	}
	
	if song != "" {
		content = append(content, "")
		content = append(content, fmt.Sprintf("♫ Now Playing: %s", song))
	}
	
	return nowPlayingStyle.Render(strings.Join(content, "\n"))
}

// RenderStationInfo renders detailed station information
func RenderStationInfo(station *RadioStation) string {
	if station == nil {
		return lipgloss.NewStyle().Foreground(mutedColor).Render("Select a station to see details")
	}
	
	content := []string{
		titleStyle.Render(station.Name),
		fmt.Sprintf("Genre: %s", station.Genre),
		"",
		"Description:",
		station.Description,
		"",
		fmt.Sprintf("Stream URL: %s", station.URL),
	}
	
	return strings.Join(content, "\n")
}

// RenderHelp renders help text
func RenderHelp() string {
	helpText := []string{
		"Controls:",
		"↑/↓ or j/k  Navigate stations",
		"Enter/Space  Play/Stop station", 
		"l           Cycle logo (GoRadio Hub/Pepe/None)",
		"q           Quit",
		"?           Toggle this help",
	}
	
	return helpStyle.Render(strings.Join(helpText, "\n"))
}

// RenderStatus renders a status message
func RenderStatus(message string) string {
	return statusStyle.Render(message)
}

// CreateLayout creates a two-panel layout
func CreateLayout(leftContent, rightContent string) string {
	left := leftPanelStyle.Render(leftContent)
	right := rightPanelStyle.Render(rightContent)
	
	return lipgloss.JoinHorizontal(lipgloss.Top, left, right)
}

// Music note animations for visual flair
var musicNotes = []string{"♪", "♫", "♬", "♭", "♮", "♯"}

// GetMusicNote returns a random music note for animations
func GetMusicNote(index int) string {
	return musicNotes[index%len(musicNotes)]
}

// WaveAnimation creates a simple wave animation
func WaveAnimation(step int) string {
	waves := []string{
		"▁▂▃▄▅▆▇█",
		"▂▃▄▅▆▇█▁",
		"▃▄▅▆▇█▁▂", 
		"▄▅▆▇█▁▂▃",
		"▅▆▇█▁▂▃▄",
		"▆▇█▁▂▃▄▅",
		"▇█▁▂▃▄▅▆",
		"█▁▂▃▄▅▆▇",
	}
	
	return waves[step%len(waves)]
}
