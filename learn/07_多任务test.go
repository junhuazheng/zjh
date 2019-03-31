package main

import (
	"fmt"
	"time"
)

func Printer(str string) {
	for _, data := range str {
		fmt.Printf("%c", data)
		time.Sleep(time.Second)
	}
	fmt.Printf("\n")
}

func main() {

	go func() {
		Printer("aaa")

	}()

	go func() {
		Printer("bbb")
	}()

	for {

	}
}
