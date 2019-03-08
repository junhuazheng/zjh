package main

import "fmt"

func main() {

	var r int = 11
	switch {
	case r > 22:
		fmt.Println(r)
	case r < 30:
		fmt.Println(r)
	}
}
