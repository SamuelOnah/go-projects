package functions

import (
	"regexp"
)

func FixQuotes(word string) string {
	pattern := regexp.MustCompile(`'\s+(.*?)\s+'`)
	word = pattern.ReplaceAllString(word, "'$1'")
	return word
}
