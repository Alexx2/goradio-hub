package main

import (
	"fmt"
	"log"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

// Model represents the state of our TUI application
type Model struct {
	stations      []RadioStation
	selected      int
	startIdx      int
	visibleCount  int
	player        *Player
	showHelp      bool
	lastUpdate    time.Time
	animationStep int
	quitting      bool
	currentLogo   LogoType
}

// tickMsg is sent every second for animations and updates
type tickMsg time.Time

// Init initializes the model
func (m Model) Init() tea.Cmd {
	return tea.Batch(
		tick(),
		tea.EnterAltScreen,
	)
}

// tick sends a tick message every second
func tick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

// Update handles messages and updates the model
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			m.player.Stop()
			return m, tea.Quit
			
		case "up", "k":
			if m.selected > 0 {
				m.selected--
				if m.selected < m.startIdx {
					m.startIdx = m.selected
				}
			}
			
		case "down", "j":
			if m.selected < len(m.stations)-1 {
				m.selected++
				if m.selected >= m.startIdx+m.visibleCount {
					m.startIdx = m.selected - m.visibleCount + 1
				}
			}
			
		case "enter", " ":
			station := &m.stations[m.selected]
			if m.player.GetCurrentStation() == station && m.player.GetState() == StatePlaying {
				m.player.Stop()
			} else {
				err := m.player.Play(station)
				if err != nil {
					// Error is handled in the player
				}
			}
			
		case "l":
			// Cycle through logo types
			m.currentLogo = LogoType((int(m.currentLogo) + 1) % 3)
			
		case "?":
			m.showHelp = !m.showHelp
		}
		
	case tickMsg:
		m.animationStep++
		m.lastUpdate = time.Time(msg)
		return m, tick()
		
	case tea.WindowSizeMsg:
		// Adjust visible count based on window height
		m.visibleCount = max(5, min(20, msg.Height-10))
		
	case tea.QuitMsg:
		m.player.Stop()
		return m, tea.Quit
	}
	
	return m, nil
}

// View renders the TUI
func (m Model) View() string {
	if m.quitting {
		return "Thanks for using GoRadio Hub! 🎵\n"
	}
	
	// Left panel: Station list
	var leftContent string
	
	// Add logo
	logo := RenderLogo(m.currentLogo)
	if logo != "" {
		leftContent += logo
		leftContent += "\n"
	}
	
	// Add logo selector info
	leftContent += RenderLogoSelector(m.currentLogo)
	leftContent += "\n"
	
	// Add subtitle with animation
	subtitle := fmt.Sprintf("🎵 Radio Stations %s", WaveAnimation(m.animationStep))
	leftContent += RenderSubtitle(subtitle)
	leftContent += "\n"
	
	// Add station list
	leftContent += RenderStationList(m.stations, m.selected, m.startIdx, m.visibleCount)
	
	// Right panel content
	var rightContent string
	
	if m.showHelp {
		rightContent += RenderTitle("Help")
		rightContent += "\n"
		rightContent += RenderHelp()
	} else {
		// Current station info
		currentStation := m.player.GetCurrentStation()
		
		// Now playing section
		status := m.player.GetState().String()
		if m.player.GetState() == StateError {
			status += ": " + m.player.GetErrorMessage()
		}
		
		song := ""
		if m.player.GetState() == StatePlaying {
			song = m.player.GetCurrentSong()
		}
		
		rightContent += RenderNowPlaying(currentStation, song, status)
		rightContent += "\n"
		
		// Station details
		if currentStation != nil {
			rightContent += RenderStationInfo(currentStation)
		} else {
			// Show info about selected station
			selectedStation := &m.stations[m.selected]
			rightContent += RenderStationInfo(selectedStation)
		}
	}
	
	// Create layout
	layout := CreateLayout(leftContent, rightContent)
	
	// Add help hint at bottom
	if !m.showHelp {
		layout += "\n" + RenderStatus("Press ? for help, l to cycle logo, Enter/Space to play/stop, q to quit")
	}
	
	return layout
}

// NewModel creates a new model instance
func NewModel() Model {
	stations := GetStations()
	
	return Model{
		stations:     stations,
		selected:     0,
		startIdx:     0,
		visibleCount: 15,
		player:       NewPlayer(),
		showHelp:     false,
		lastUpdate:   time.Now(),
		currentLogo:  LogoOriginal, // Start with original GoRadio Hub logo
	}
}

func main() {
	// Create the model
	m := NewModel()
	
	// Create the program
	p := tea.NewProgram(m, tea.WithAltScreen())
	
	// Run the program
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

// Helper function for max
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Helper function for min
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Version information
const (
	Version = "1.0.0"
	AppName = "GoRadio Hub"
)

// init function to set up the application
func init() {
	// Set up logging to a file for debugging
	f, err := os.OpenFile("goradio.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	log.SetOutput(f)
	
	log.Printf("Starting %s v%s", AppName, Version)
}
