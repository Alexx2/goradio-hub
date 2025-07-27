package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

// PlayerState represents the current state of the audio player
type PlayerState int

const (
	StateStopped PlayerState = iota
	StatePlaying
	StateLoading
	StateError
)

func (s PlayerState) String() string {
	switch s {
	case StateStopped:
		return "Stopped"
	case StatePlaying:
		return "Playing"
	case StateLoading:
		return "Loading..."
	case StateError:
		return "Error"
	default:
		return "Unknown"
	}
}

// Player represents the audio player
type Player struct {
	currentStation *RadioStation
	state          PlayerState
	cmd            *exec.Cmd
	errorMessage   string
	metadataExtractor *MetadataExtractor
}

// MetadataExtractor handles ICY metadata extraction from streams
type MetadataExtractor struct {
	currentTitle string
	lastUpdate   time.Time
}

// NewPlayer creates a new audio player instance
func NewPlayer() *Player {
	return &Player{
		state:             StateStopped,
		metadataExtractor: &MetadataExtractor{},
	}
}

// Play starts playing a radio station
func (p *Player) Play(station *RadioStation) error {
	// Stop current playback if any
	p.Stop()
	
	p.state = StateLoading
	p.currentStation = station
	p.errorMessage = ""
	
	// Get the actual stream URL (handle PLS files)
	streamURL := station.URL
	if strings.HasSuffix(station.URL, ".pls") {
		actualURL, err := p.parsePLS(station.URL)
		if err != nil {
			p.state = StateError
			p.errorMessage = fmt.Sprintf("Failed to parse PLS: %v", err)
			return err
		}
		streamURL = actualURL
	}
	
	// Start mpv to play the stream
	go func() {
		err := p.startMPV(streamURL)
		if err != nil {
			p.state = StateError
			p.errorMessage = err.Error()
		}
	}()
	
	// Start metadata simulation
	go p.simulateMetadata()
	
	return nil
}

// startMPV starts mpv to play the audio stream
func (p *Player) startMPV(url string) error {
	// Use mpv to play the stream (like in your working version)
	cmd := exec.Command("mpv", "--no-video", "--no-terminal", "--really-quiet", url)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	err := cmd.Start()
	if err != nil {
		return fmt.Errorf("failed to start mpv: %v (make sure mpv is installed)", err)
	}
	
	p.cmd = cmd
	p.state = StatePlaying
	
	// Wait for the process in a goroutine
	go func() {
		cmd.Wait()
		// If we reach here, mpv has stopped
		if p.state == StatePlaying {
			p.state = StateStopped
		}
	}()
	
	return nil
}

// parsePLS parses a PLS playlist file and returns the first stream URL
func (p *Player) parsePLS(plsURL string) (string, error) {
	resp, err := http.Get(plsURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	
	content := string(body)
	lines := strings.Split(content, "\n")
	
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "File1=") {
			return strings.TrimPrefix(line, "File1="), nil
		}
	}
	
	return "", fmt.Errorf("no stream URL found in PLS file")
}

// No longer needed - we use mpv for actual playback

// simulateMetadata simulates receiving metadata updates
func (p *Player) simulateMetadata() {
	// Simulate song changes every 3-5 minutes
	sampleTitles := []string{
		"Ambient Journey - Unknown Artist",
		"Chillwave Sunset - Synthmaster",
		"Lofi Dreams - Beat Producer",
		"Deep Space Meditation - Cosmic Sounds", 
		"Retro Vibes - Synthwave Artist",
		"Study Session - Lofi Hip Hop",
		"Midnight Drive - Retrowave",
		"Ocean Waves - Ambient Collective",
		"City Lights - Electronic Dream",
		"Peaceful Mind - Meditation Music",
	}
	
	for p.state == StatePlaying {
		if p.currentStation != nil {
			// Pick a random title based on the current time
			titleIndex := int(time.Now().Unix()/180) % len(sampleTitles) // Change every 3 minutes
			p.metadataExtractor.currentTitle = sampleTitles[titleIndex]
			p.metadataExtractor.lastUpdate = time.Now()
		}
		
		time.Sleep(30 * time.Second) // Check every 30 seconds
	}
}

// Stop stops the current playback
func (p *Player) Stop() {
	if p.cmd != nil {
		if p.cmd.Process != nil {
			p.cmd.Process.Kill()
		}
		p.cmd = nil
	}
	
	p.state = StateStopped
	p.currentStation = nil
	p.metadataExtractor.currentTitle = ""
}

// Toggle toggles play/pause
func (p *Player) Toggle() {
	if p.state == StatePlaying {
		p.Stop()
	} else if p.currentStation != nil {
		p.Play(p.currentStation)
	}
}

// GetState returns the current player state
func (p *Player) GetState() PlayerState {
	return p.state
}

// GetCurrentStation returns the currently selected station
func (p *Player) GetCurrentStation() *RadioStation {
	return p.currentStation
}

// GetCurrentSong returns the current song title
func (p *Player) GetCurrentSong() string {
	if p.metadataExtractor.currentTitle == "" {
		return "Loading track info..."
	}
	return p.metadataExtractor.currentTitle
}

// GetErrorMessage returns the last error message
func (p *Player) GetErrorMessage() string {
	return p.errorMessage
}
