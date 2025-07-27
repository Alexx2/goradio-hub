// +build ignore

package main

import (
	"fmt"
	"log"
)

// Simple test program to verify stations are loaded correctly
func main() {
	stations := GetStations()
	
	fmt.Printf("GoRadio Hub - Station Test\n")
	fmt.Printf("======================\n\n")
	fmt.Printf("Total stations loaded: %d\n\n", len(stations))
	
	// Group by genre
	genreCount := make(map[string]int)
	for _, station := range stations {
		genreCount[station.Genre]++
	}
	
	fmt.Printf("Stations by genre:\n")
	for genre, count := range genreCount {
		fmt.Printf("  %-20s: %d stations\n", genre, count)
	}
	
	fmt.Printf("\nFirst 5 stations:\n")
	for i, station := range stations[:5] {
		fmt.Printf("%d. %s (%s)\n", i+1, station.Name, station.Genre)
		fmt.Printf("   URL: %s\n", station.URL)
		fmt.Printf("   Description: %s\n\n", station.Description)
	}
	
	// Test genre filtering
	lofiStations := GetStationsByGenre("Lofi Hip Hop")
	fmt.Printf("Lofi Hip Hop stations: %d\n", len(lofiStations))
	
	vaporwaveStations := GetStationsByGenre("Vaporwave")
	fmt.Printf("Vaporwave stations: %d\n", len(vaporwaveStations))
	
	ambientStations := GetStationsByGenre("Ambient")
	fmt.Printf("Ambient stations: %d\n", len(ambientStations))
	
	log.Println("Station test completed successfully!")
}
