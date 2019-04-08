package main

import (
	"fmt"
	"time"
)

var n Name

func blue(w int) {
	fmt.Println("w = ", w)
	return
}

type Name struct {
	firstName string
	lastName  string
	nickName  string
}

func GetName() (firstName, lastName, nickName string) {
	return
}

func main() {
	n.firstName = "Cuihua"
	var a int
	a = 5
	blue(a)
	GetName()
	fmt.Println(n.firstName)

	for {
		blue(a)
		time.Sleep(time.Second * 2)
	}

}
