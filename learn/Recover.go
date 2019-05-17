package main

import (
	"fmt"
)

//定义一个让程序运行时崩溃的函数
func badCall() {
	panic("bad end")
}

func test() {

	//recover函数可以处理错误，做法是打印这个错误的输出并且不让程序崩溃。
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("Panicing %s\n", e)
		}
	}()

	/*调用这个运行时崩溃的函数，下面的代码不会执行，而是直接结束函数
	结束函数之后就会触发defer关键字，因此会被recover函数捕捉*/
	badCall()
	fmt.Printf("After bad call\r\n")
}

func main() {
	fmt.Printf("Calling test\r\n")

	//调用函数，如果程序没有崩溃，就可以执行下一行代码
	test()

	fmt.Printf("Test completed\r\n")
}
