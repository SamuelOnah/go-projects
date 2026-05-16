package main

import (
	"fmt"
	"os"
	"strings"
)

func LoadBanner(filename string) (map[rune][]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error: file doesn't exist")
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("file is empty")
	}

	characters := strings.ReplaceAll(string(data), "\r\n", "\n")

	lines := strings.Split(characters, "\n")

	banner := make(map[rune][]string)

	currentChar := rune(32)

	for i := 1; i < len(lines); i += 9 {
		if i+8 > len(lines) {
			break
		}

		charArt := lines[i : i+8]

		banner[currentChar] = charArt

		currentChar++
	}

	if len(banner) != 95 {
		return nil, fmt.Errorf("expected 95 characters, got %d", len(banner))
	}

	return banner, nil
}
