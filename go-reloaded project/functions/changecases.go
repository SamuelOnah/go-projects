package functions

import (
	"strings"
)

func Changecases(word string) string {
	words := strings.Fields(word)
	for i := 1; i < len(words); i++ {
		switch words[i] {
		case "(up)":
			words[i-1] = strings.ToUpper(words[i-1])
			words = append(words[:i], words[i+1:]...)

		case "(low)":
			words[i-1] = strings.ToLower(words[i-1])
			words = append(words[:i], words[i+1:]...)
		case "(cap)":
			words[i-1] = strings.ToLower(words[i-1])
			words[i-1] = strings.Title(words[i-1])
			words = append(words[:i], words[i+1:]...)
		}
	}
	return strings.Join(words, " ")
}
