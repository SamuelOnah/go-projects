package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	v1, _ := strconv.ParseInt("1E", 16, 64)
// 	v2, _ := strconv.ParseInt("1010", 2, 64)
// 	v3, _ := strconv.ParseInt("46", 8, 64)

// 	fmt.Println(v1)
// 	fmt.Println(v2)
// 	fmt.Println(v3)
// }

// import (
// 	"fmt"
// )

// func punc(n string) bool {
// 	return n == "," || n == "." || n == "!"
// }

// func main() {
// 	fmt.Println(punc(","))
// 	fmt.Println(punc("."))
// 	fmt.Println(punc("x"))
// }

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	b := strings.ToLower("hELLO")
// 	fmt.Println(strings.Title(b))
// 	fmt.Println(strings.ToUpper(b))
// }

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

// import (
// 	"fmt"
// 	"strings"
// )

// func joinpunc(s string) string {
// 	s = strings.ReplaceAll(s, " ,", ",")
// 	s = strings.ReplaceAll(s, " !", "!")
// 	s = strings.ReplaceAll(s, " .", ".")
// 	return s
// }

// func main() {
// 	fmt.Println(joinpunc("Hello , world !"))
// }

// import (
// 	"fmt"
// 	"strings"
// )

// func a2An(w string) string {
// 	f := string(w[0])
// 	v := "aieouhAIEOUH"

// 	if strings.ContainsAny(f, v) {
// 		return "an " + w
// 	}
// 	return "a " + w
// }

// func main() {
// 	fmt.Println(a2An("horse"))
// 	fmt.Println(a2An("apple"))
// 	fmt.Println(a2An("book"))
// }

// import (
// 	"fmt"
// 	"strings"
// )

// func fixArt(t string) string {
// 	t = strings.ReplaceAll(t, "A a", "An a")
// 	t = strings.ReplaceAll(t, "A h", "An h")

// 	return t
// }

// func main() {
// 	fmt.Println(fixArt("there it was. A amazing rock. A honest man. A book"))
// }

import (
	"fmt"
	"regexp"
)

func fixSingleQuotes(t string) string {
	re := regexp.MustCompile(`\s*'\s*`)

	return re.ReplaceAllString(t, "'")
}

func main() {
	fmt.Println(fixSingleQuotes(" ' awesome ' "))
}
