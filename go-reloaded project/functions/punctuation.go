package functions

import (
	"regexp"
)

func FixPunctuation(text string) string {
	before := regexp.MustCompile(`\s+([.,!?:;]+)`)
	text = before.ReplaceAllString(text, `$1`)

	after := regexp.MustCompile(`([.,!?:;]+)(\w)`)
	text = after.ReplaceAllString(text, `$1 $2`)

	return text
}
