package main

import (
	"fmt"
	"os"
	// "strings"
)

// func main() {
// 	s, sep := "", ""
// 	for _, arg := range os.Args[1:] {
// 		s += sep + arg
// 		sep = ""
// 	}
// 	fmt.Println(s)
// }

// func main() {
// 	fmt.Println(strings.Join(os.Args[1:], ""))
// }

func main() {
	fmt.Println(os.Args[0])
}
