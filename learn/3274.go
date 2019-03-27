//使用go关键字，将running()函数并发执行，每隔一秒打印一次计数器
//而main的goroutine则等待用户输入，两个行为可以同时进行

package main

import (
	"fmt"
	"time"
)

fun running() {
	
	var times int
	
	for {
		times++
		fmt.Println("tick", times)
		
		times.Sleep(time.Second)
	}
}

func main() {
	
	go running()
	
	var input string
	
	fmt.Scanln(&input)
}