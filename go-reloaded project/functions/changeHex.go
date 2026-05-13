package functions

import (
	"strconv"
	"strings"
)

func Changehex(word string) string {
	words := strings.Fields(word)

	for m := 1; m < len(words); m++ {
		switch words[m] {
		case "(hex)":
			value, err := strconv.ParseInt(words[m-1], 16, 64)
			if err != nil {
				return "wrong input"
			}
			words[m-1] = strconv.FormatInt(value, 10)
			words = append(words[:m], words[m+1:]...)

		case "(bin)":
			value, err := strconv.ParseInt(words[m-1], 2, 64)
			if err != nil {
				return "wrong input"
			}
			words[m-1] = strconv.FormatInt(value, 10)
			words = append(words[:m], words[m+1:]...)

		}

	}
	return strings.Join(words, " ")
}
