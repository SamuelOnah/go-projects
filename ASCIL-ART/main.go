package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . your text here fontstyle")
		return
	}
	bannerFiles := os.Args[2]
	inputfile := strings.ReplaceAll(os.Args[1], "\\n", "\n")
	data, err := os.ReadFile(bannerFiles)
	if err != nil {
		fmt.Println("Oops. File to read not found!")
		return
	}

	reader := strings.ReplaceAll(string(data), "\r\n", "\n")
	fontLines := strings.Split(reader, "\n")
	inputLines := strings.Split(inputfile, "\n")
	for _, line := range inputLines {
		if line == "" {
			fmt.Println()
			continue
		}

		for lines := 1; lines <= 8; lines++ {
			for _, char := range line {
				if char < 32 || char > 126 {
					continue
				}
				start := int(char-32)*9 + lines

				fmt.Print(fontLines[start])
				//fmt.Printf("Index: %d -> %s\n", start, fontLines[start])
			}
			fmt.Println()
		}

	}

}
