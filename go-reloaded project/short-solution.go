package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func hexToDecimal(hexStr string) (int64, error) {
	return strconv.ParseInt(hexStr, 16, 64)
}

func binToDecimal(binStr string) (int64, error) {
	return strconv.ParseInt(binStr, 2, 64)
}

func cap(s string) string {
	x := strings.ToLower(s)
	return strings.Title(x)
}

func upperN(words []string, n int) string {
	for i := len(words) -n; i < len(words); i++ {
		words[i] = strings.ToUpper(words[i])
	}
	return strings.Join(words, " ")
}

func punc(n string) bool {
	return n == "," || n == "." || n == "!"
}

func fixArticles(words []string) string {
	word1 := strings.Join(words, " ",)
	word2 := strings.Replace(word1, " ,", ",",-1) 
	word3 := strings.Replace(word2, " !", "!",-1)
	return word3
}

func a2An(w string ) string {
	x := string(w[0])
	y := "aieouhAIEOUH"

	if strings.ContainsAny(x,y) {
		return "a " + w
	}
	return "an " + w 
}

func fixSentenceArticles(text string) string {
	text = strings.ReplaceAll(text, "A a", "An a")
	text = strings.ReplaceAll(text, "A h", "An h")
	return text 

}


func fixSingleQuotes(t string) string {
	x := regexp.MustCompile(`\s*'\s*`)
	return x.ReplaceAllString(t, "'")
}