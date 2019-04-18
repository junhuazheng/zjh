package main

import (
	"fmt"
)

type People interface {
	Show()
}

type Stuent struct{}

func (stu *Stuent) Show() {

}

func live() People {
	var stu *Stuent
	return stu
}

func main() {
	if live() == nil {
		fmt.Println("AAAAA")
	} else {
		fmt.Println("BBBBB")
	}
}
