package main

import "fmt"

func main() {

	var a = "hello"
	switch a {
	case "hello":
		fmt.Println(1)
	case "world":
		fmt.Println(2)
	default:
		fmt.Println(3)
	}
}
