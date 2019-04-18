package main

import (
	"fmt"
)

type People interface {
	Speak(string) string
}

type Stuent struct {
}

func (stu *student) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

func main() {
	var peo People = Stuent{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
}

/*
编译失败，值类型Student{}未实现接口People的方法，不能定义为People类型。
*/
