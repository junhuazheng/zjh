//go关键字后也可以为匿名函数或必报启动goroutine

//1、使用匿名函数或闭包创建goroutine时，除了将函数定义部分写在go的后面，还需要加上匿名函数的调用参数：
go func(参数列表) {
	函数体
}(调用参数列表)

//参数列表：函数体内的参数变量列表
//函数体：匿名函数代码
//调用参数列表：启动goroutine时，需要向匿名函数传递的调用参数

//在main()函数中创建一个匿名函数并为匿名函数启动goroutine。匿名函数没有参数。代码将并行执行定时打印技术的效果
package main

import (
	"fmt"
	"time"
)

func main() {
	
	go func() {
		
		var times int
		
		for {
			times++
			fmt.Pritln("tick, times")
			
			time.Sleep(time.Second)
		}
	}() //括号的功能是调用匿名函数的参数列表。由于前面创建匿名函数没有参数，因此这里的参数列表也是空的
	
	var input string
	fmt.Scanln(&input)
}

//所有goroutine在main()函数结束时会一同结束

//goroutine虽然类似于线程概念，但是从调度性能上没有现成细致，而细致程度取决于Go程序goroutine调度器的实现和运行环境

//终止goroutine的最好方法就是自然返回goroutine对应的函数。虽然可以用
// golang.org/x/net/context包进行goroutine生命期深度控制，但是这种方法仍然处于内部测试阶段
