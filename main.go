// package main

// import(
// 	"fmt"
// 	"strconv"
// )
// func main(){
// 	v1, _:= strconv.ParseInt("1E", 16, 64)
// 	v2, _ := strconv.ParseInt("1010", 2, 64)

// 	fmt.Println(v1, v2 )
// }

// package main

// import "fmt"

// func pun(n string) bool {
// 	return n == "," || n == "." || n == "!"
// }

// func main() {
// 	fmt.Println(pun(","))
// 	fmt.Println(pun("."))
// 	fmt.Println(pun("x"))
// }

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	b := strings.ToLower("hELLO")
// 	fmt.Println(strings.Title(b))
// }

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	w, n := []string{"this", "is", "so", "exciting"}, 2
// 	for i := len(w) - n; i < len(w); i++ {
// 		w[i] = strings.ToUpper(w[i])
// 	}
// 	fmt.Println(w)
// }

package main

import "fmt"

func main() {
	word := "samuel"
	fmt.Println(string(word[0]))
}
