package main

import (
	"fmt"
	"os"
)

func main() {
	var FontStyle string

	if len(os.Args) == 2 {
		FontStyle = "standard.txt"
	} else if len(os.Args) == 3 {
		FontStyle = os.Args[2]
	} else {
		fmt.Println("Usage: go run . [Your text] [Fontstyle]")
		return
	}
	bannerMap, err := LoadBanner(FontStyle)
	if err != nil {
		fmt.Println("Oops. File not found")
		return
	}
	lines := splitLines(os.Args[1])
	for _, line := range lines {
		if line == "" {
			fmt.Println()
			continue
		}
		_, err := validate(line)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Print(generateAscii(line, bannerMap))
	}
}
