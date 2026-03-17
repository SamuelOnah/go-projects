package main

import (
	"fmt"
	"strings"
)

func fixQoute(text string) string {
	words := strings.Fields(text)
	text = strings.Join(words, "")
	text = strings.ReplaceAll(text, "' ", "'")
	text = strings.ReplaceAll(text, " '", "'")
	text = strings.ReplaceAll(text, ":'", ": '")

	return text
}

func main() {
	input := "start' test 'example: ' test ' end "
	fixed := fixQoute(input)
	fmt.Println(input)

	fmt.Println(fixed)
}
