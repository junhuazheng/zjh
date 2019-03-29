/*goroutine（Go语言并发）如何使用才更加高效？

Go语言原生支持并发是被众人津津乐道的特性。
goroutine早期是infermo操作系统的一个实验性特性，而现在这个特性与操作系统一起，将开发变得越来越简单。
很多刚开始使用Go语言开发的人都很喜欢使用并发特性，而没有考虑并发是否真正能解决他们的问题

1、了解goroutine的生命期时再创建goroutine

在Go语言中，开发者习惯将并发内容与goroutine一一对应地创建goroutine。
开发者很少会考虑goroutine在什么时候能退出和控制goroutine生命期
这就会造成goroutine失控的情况

失控的goroutine：*/
package main

import (
	"fmt"
	"runtime"
)

/*consumer()函数模拟平时业务中放到goroutine中执行的耗时操作。
该函数从其他goroutine中获取和接受数据或者指令，处理后返回结果。*/
func consumer(ch chan int) {

	for {
		data := <-ch

		fmt.Println(data)
	}
}

func main() {

	ch := make(chan int)

	for {
		var dummy string

		fmt.Scan(&dummy)

		go consumer(ch)

		fmt.Println("goroutines:", runtime.NumGoroutine())
	}
}
