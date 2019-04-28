/*
1、什么是高并发？
高并发(High Concurrency)是互联网分布式系统架构设计中必须考虑的因素之一
他通常是指，通过设计保证系统能够同时并行处理很多请求。

2、Go为什么能做到高并发
goroutine是Go并行设计的核心。goroutine是到底其实就是协程，但是它比线程更小。
几十个goroutine可能体现在底层也就是五六个线程，go语言内部帮你实现了这些goroutine之间的内存共享。
执行goroutine只需极少的栈内存（大概是4~5KB），当然也会根据相应的数据伸缩。
也正因为如此，可同时运行成千上万个并发任务。goroutine更易用、更高效、更轻便。
协程更轻量，占内存更小，这就是他能做到高并发的前提。

3、go web开发中怎么做到高并发的能力
学习go的HTTP代码。先创建一个简单的web服务。
*/

package main

import (
	"fmt"
	"log"
	"net/http"
)

func response(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!") //这个写入到w的是输出到客户端的
}

func main() {
	http.HandleFunc("/", response)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}
