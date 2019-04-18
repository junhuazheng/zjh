package main

import (
	"fmt"
)

type People interface {
	Speak(string) string
}

type Stuent struct {
}

func (stu Stuent) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

func main() {
	var peop People = &Stuent{}
	think := "bitch"
	fmt.Println(peop.Speak(think))
}

/*
指针类型的结构体对象可以同时调用结构体值类型和指针类型对应的方法。
而值类型的结构体对象只能调用值类型对应的接口方法。
*/
