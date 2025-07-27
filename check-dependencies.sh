#!/bin/bash

echo "🎵 GoRadio Hub Dependency Check 🎵"
echo "==============================="
echo

# Check for mpv
if command -v mpv >/dev/null 2>&1; then
    echo "✅ mpv is installed: $(mpv --version | head -n1)"
else
    echo "❌ mpv is NOT installed"
    echo "📦 To install mpv:"
    echo "   Ubuntu/Debian: sudo apt install mpv"
    echo "   Arch Linux:    sudo pacman -S mpv"
    echo "   Fedora:        sudo dnf install mpv"
    echo "   macOS:         brew install mpv"
    echo
    echo "⚠️  GoRadio Hub requires mpv to play audio streams!"
    exit 1
fi

# Check for Go
if command -v go >/dev/null 2>&1; then
    echo "✅ Go is installed: $(go version)"
else
    echo "❌ Go is NOT installed"
    echo "📦 Install Go from: https://golang.org/dl/"
    exit 1
fi

echo
echo "🎉 All dependencies are ready!"
echo "🚀 You can now run: ./goradio"
echo "🐸 Press 'l' to see the awesome Pepe ASCII art!"
