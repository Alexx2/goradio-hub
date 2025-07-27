// +build ignore

package main

import (
	"fmt"
	"time"
)

// Simple test program to display all logos
func main() {
	fmt.Println("🎵 GoRadio Hub Logo Gallery 🎵")
	fmt.Println("==========================\n")
	
	// Original GoRadio Hub Logo
fmt.Println("1. Original GoRadio Hub Logo:")
	fmt.Println(RenderLogo(LogoOriginal))
	
	time.Sleep(1 * time.Second)
	
	// Pepe Logo
	fmt.Println("\n2. Pepe the Frog Logo (FEELS GOOD MAN):")
	fmt.Println(RenderLogo(LogoPepe))
	
	time.Sleep(1 * time.Second)
	
	// No Logo
	fmt.Println("\n3. No Logo Mode:")
	logo := RenderLogo(LogoNone)
	if logo == "" {
		fmt.Println("(Clean interface - no logo displayed)")
	} else {
		fmt.Println(logo)
	}
	
	fmt.Println("\n🎮 In the app, press 'l' to cycle between these logos!")
	fmt.Println("🐸 Pepe is rendered in bold green for maximum comfy vibes!")
}
