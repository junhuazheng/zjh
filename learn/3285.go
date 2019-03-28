//使用go关键字，将running()函数并发执行，每个一秒打印一次计数器
//而main的goroutine则等待用户输入，两个行为可以同时进行

package main

import (
	"fmt"
	"time"
)

func running() {

	var times int
	//构建一个无线循环
	for {
		times++

		//输出times变量的值
		fmt.Println("tick", times)

		//延迟1秒
		time.Sleep(time.Second)
	}
}

func main() {

	//并发执行程序（使用go关键字让running()函数并发运行）
	go running()

	//接受命令行输入，不做任何事情
	var input string
	//接受用户输入，直到按Enter键时将输入的内容写入input变量中并返回，整个程序终止
	fmt.Scanln(&input)
}

//代码执行后，命令行会不断地输出tick,同时可以使用fnt.Scanln()接受用户输入。两个环节可以同时进行

//这个例子中，Go程序在启动时，运行时（runtime）会默认为main()函数创建一个goroutine。
//在main()函数的goroutine中执行到go running 语句时，归属于running()函数的goroutine被创建
//running()函数开始在自己的goroutine中执行。此时main()继续执行，两个goroutine通过Go程序的调度机制同时运作
