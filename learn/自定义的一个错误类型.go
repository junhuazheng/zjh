package main

import (
	"errors"
	"fmt"
)

var (
	CustomError       error  //用于定义错误的变量
	PromptInformation string //用于定义提示信息的变量
)

func main() {

	PromptInformation = "这是自定义的一个错误类型"

	//errors包的New方法可以创建一个error类型的数据，需要传入一个字符串类型用于提示信息。
	CustomError = errors.New(PromptInformation)

	fmt.Printf("error: %v", CustomError)
}
