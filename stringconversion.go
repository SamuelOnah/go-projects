package main

import (
	"fmt"
	"strings"
)



func main() {
	word := "samuel"
	chief := ("hrh chief dr. godwin ngbede onah")
	

	fmt.Println(strings.ToUpper(chief))
	fmt.Println(strings.Title(chief))
	fmt.Println(strings.ToLower(chief))
	fmt.Println(strings.Fields(chief))
	fmt.Println(strings.TrimRight(chief, " "))
	fmt.Println(word[:len(word)-2] + strings.ToUpper(word[len(word)-2:]))
}

