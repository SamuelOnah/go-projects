
package main 

import "fmt"

func rev(s string) string {
	if s == "" {
		return ""
	}
	return rev(s[1:]) + s[:1]
}

func main(){
	x := "miriam"
	fmt.Println(rev(x))
}


package main 

import (
	"fmt"
	"strings"
)
func count(word string) map[string] int {
	m := make(map[string]int)
	for _, i := range strings.Fields(word) {
		m[i]++
	}
	return m
}

func main(){
	fmt.Println(count("sam, ogo, sam, unapapa"))
}


