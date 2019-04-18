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
	var peo People = student{}
	think := "hi"
	fmt.Println(peo.Speak(think))
}

/*
与上一个例子不同的是
func (stu Student) Speak(think string) (talk string)

上个例子是
func (stu *Student) Speak(think string) (talk string)
导致编译失败，值类型Student{}未实现接口People的方法，不能定义为People类型。
*/
