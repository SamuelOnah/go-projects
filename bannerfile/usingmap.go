package main

import (
	"fmt"
	"os"
	"strings"
)


func main (){
	bannerfile := os.Args[2]

	data, err := os.ReadFile(bannerfile)
	if err != nil {
		fmt.Println("error reading banner file", err)
		return 
	}

	banner := strings.Split(string(data), "\n")

	var bannerfont = make(map[rune][]string)

	for i := ' '; i <= '~'; i++{
		start :=  int(i -32)*9 + 1
		end := start + 8

		bannerfont[i] = banner[start:end]
	}
}