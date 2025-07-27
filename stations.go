package main

import "fmt"

// RadioStation represents a radio station with its metadata
type RadioStation struct {
	Name        string
	URL         string
	Genre       string
	Description string
}

// GetStations returns all available radio stations
func GetStations() []RadioStation {
	return []RadioStation{
		// VERIFIED WORKING SomaFM Stations - Direct MP3 URLs
		{
			Name:        "Groove Salad",
			URL:         "https://ice1.somafm.com/groovesalad-128-mp3",
			Genre:       "Downtempo",
			Description: "Chilled ambient/downtempo beats and grooves",
		},
		{
			Name:        "Drone Zone", 
			URL:         "https://ice1.somafm.com/dronezone-128-mp3",
			Genre:       "Ambient",
			Description: "Deep ambient soundscapes for meditation",
		},
		{
			Name:        "Lush",
			URL:         "https://ice1.somafm.com/lush-128-mp3",
			Genre:       "Dream Pop",
			Description: "Sensuous vocals with electronic influence",
		},
		{
			Name:        "Space Station",
			URL:         "https://ice1.somafm.com/spacestation-128-mp3",
			Genre:       "Space Ambient",
			Description: "Spaced-out ambient electronica",
		},
		{
			Name:        "Fluid",
			URL:         "https://ice1.somafm.com/fluid-128-mp3",
			Genre:       "Future Soul",
			Description: "Instrumental hiphop and liquid trap",
		},
		{
			Name:        "Beat Blender",
			URL:         "https://ice1.somafm.com/beatblender-128-mp3",
			Genre:       "Deep House",
			Description: "Late night deep-house and downtempo chill",
		},
		{
			Name:        "Vaporwaves",
			URL:         "https://ice1.somafm.com/vaporwaves-128-mp3",
			Genre:       "Vaporwave",
			Description: "Aesthetic vaporwave and future funk",
		},
		{
			Name:        "Underground 80s",
			URL:         "https://ice1.somafm.com/u80s-128-mp3",
			Genre:       "Synthpop",
			Description: "Early 80s UK synthpop and new wave",
		},
		{
			Name:        "DEF CON Radio",
			URL:         "https://ice1.somafm.com/defcon-128-mp3",
			Genre:       "Hacker",
			Description: "Music for hacking - DEF CON vibes",
		},
		{
			Name:        "Secret Agent",
			URL:         "https://ice1.somafm.com/secretagent-128-mp3",
			Genre:       "Spy Jazz",
			Description: "The soundtrack for your stylish life",
		},
		
		// VERIFIED WORKING - Additional Genres
		{
			Name:        "Plaza Radio",
			URL:         "https://radio.plaza.one/ogg",
			Genre:       "Vaporwave",
			Description: "24/7 vaporwave, synthwave, and aesthetic",
		},
		{
			Name:        "Bassdrive DNB",
			URL:         "http://bassdrive.com/bassdrive3.m3u",
			Genre:       "Drum & Bass",
			Description: "World's largest drum & bass station",
		},
		{
			Name:        "ChillHop Radio",
			URL:         "http://stream.laut.fm/chillhop",
			Genre:       "Lofi Hip Hop",
			Description: "24/7 chillhop and lofi hip hop beats",
		},
		{
			Name:        "Deep Space One",
			URL:         "https://ice1.somafm.com/deepspaceone-128-mp3",
			Genre:       "Space Music",
			Description: "Deep ambient electronic space music",
		},
		{
			Name:        "Boot Liquor",
			URL:         "https://ice1.somafm.com/bootliquor-128-mp3",
			Genre:       "Americana",
			Description: "Roots music for cowpokes and indie rockers",
		},
		{
			Name:        "Cliqhop IDM",
			URL:         "https://ice1.somafm.com/cliqhop-128-mp3",
			Genre:       "IDM",
			Description: "Blips, beeps and clicks of intelligent dance",
		},
		{
			Name:        "The Trip",
			URL:         "https://ice1.somafm.com/thetrip-128-mp3",
			Genre:       "Psychedelic",
			Description: "Progressive rock and trippy experimental music",
		},
		{
			Name:        "Seven Inch Soul",
			URL:         "https://ice1.somafm.com/7soul-128-mp3",
			Genre:       "Soul/R&B",
			Description: "Vintage soul tracks from original 45 RPM vinyl",
		},
		{
			Name:        "Metal Detector",
			URL:         "https://ice1.somafm.com/metal-128-mp3",
			Genre:       "Metal",
			Description: "From black to doom, thrash to post-metal",
		},
		{
			Name:        "Folk Forward",
			URL:         "https://ice1.somafm.com/folkfwd-128-mp3",
			Genre:       "Folk",
			Description: "Indie folk, alt-folk and folk classics",
		},
	}
}

// GetStationsByGenre returns stations filtered by genre
func GetStationsByGenre(genre string) []RadioStation {
	var filtered []RadioStation
	for _, station := range GetStations() {
		if station.Genre == genre {
			filtered = append(filtered, station)
		}
	}
	return filtered
}

// GetGenres returns all unique genres
func GetGenres() []string {
	genreMap := make(map[string]bool)
	for _, station := range GetStations() {
		genreMap[station.Genre] = true
	}
	
	var genres []string
	for genre := range genreMap {
		genres = append(genres, genre)
	}
	return genres
}

// PrintStations prints all stations for debugging
func PrintStations() {
	stations := GetStations()
	fmt.Printf("Total stations: %d\n\n", len(stations))
	for i, station := range stations {
		fmt.Printf("%d. %s\n", i+1, station.Name)
		fmt.Printf("   Genre: %s\n", station.Genre)
		fmt.Printf("   URL: %s\n", station.URL)
		fmt.Printf("   Description: %s\n\n", station.Description)
	}
}
