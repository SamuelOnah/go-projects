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
	b := strings.ToLower(s)
	return strings.Title(b)
}


func upperN(words []string, n int) string {
	for i := len(words) -n ; i < len(words); i ++ {
		words[i] = strings.ToUpper(words[i])
	}
	return strings.Join(words, " ")
}


func punc(s string) bool {
	return s == "," || s == "." ||  s == "%"
}


func fixArticles(words []string) string {
	word1 := strings.Join(words, " ")
	word2 := strings.ReplaceAll(word1, " ,", ",")
	word3 := strings.ReplaceAll(word2, " !", "!")
	return word3
}

func article(word string) string {
	x := string(word[0])
	y := "aieouhAIEOUH"

	if  strings.ContainsAny(x, y){
		return "An " + word
	}
	return "A " + word 
}

func fixSentenceArticles(text string) string {
	text = strings.ReplaceAll(text, "A a", "An a")
	text = strings.ReplaceAll(text, "A h", "An h")
	return text 
}

func fixSingleQuotes(text string) string {
	x := regexp.MustCompile(`\s*'\s*`)
	return x.ReplaceAllString(text, "'")
}
	


func main(){
	
	fmt.Println(hexToDecimal("1E"))
	fmt.Println(hexToDecimal("FF"))
	fmt.Println(binToDecimal("10"))
	fmt.Println(binToDecimal("1010"))
	fmt.Println(binToDecimal("11111111"))
	fmt.Println(cap("hELLO"))
	fmt.Println(cap("WORLD"))
	fmt.Println(upperN([]string{"How","are","you"}, 2))
	fmt.Println(punc(","))
	fmt.Println(punc("."))
	fmt.Println(punc("x"))
	fmt.Println(fixArticles([]string{"hello", ",", "world", "!"}))
	fmt.Println(article("apple"))
	fmt.Println(article("book"))
	fmt.Println(article("honest"))
	fmt.Println(fixSentenceArticles("There it was. A amazing rock. A honest man."))
	fmt.Println(fixSingleQuotes("' awesome '" ))


}