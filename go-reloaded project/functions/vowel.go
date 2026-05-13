package functions

import (
	"strings"
)

func ChangeVowels(s string) string {
	word := strings.Fields(s)

	for i := 0; i < len(word)-1; i++ {

		switch word[i] {
		case "A":
			if i+1 < len(word) {
				vowel := word[i+1]

				if len(vowel) > 0 &&
					(vowel[0] == 'a' ||
						vowel[0] == 'e' ||
						vowel[0] == 'i' ||
						vowel[0] == 'o' ||
						vowel[0] == 'u' ||
						vowel[0] == 'h') {
					word[i] = "An"

				}
			}
		case "a":
			if i+1 < len(word) {
				vowel := word[i+1]

				if len(vowel) > 0 &&
					(vowel[0] == 'a' ||
						vowel[0] == 'e' ||
						vowel[0] == 'i' ||
						vowel[0] == 'o' ||
						vowel[0] == 'u' ||
						vowel[0] == 'h') {
					word[i] = "an"

				}
			}

		}
	}
	return strings.Join(word, " ")
}
