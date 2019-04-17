package main

import (
	"fmt"
	"runtime"
)

func main() {

	runtime.GOMAXPROCS(1)

	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)
	int_chan <- 1
	string_chan <- "hello"

	select {
	case value := <-int_chan:
		fmt.Println(value)
	case value := <-string_chan:
		panic(value)
	}
}

/*
有可能会发生异常，如果没有select这段代码，就会出线程阻塞
当有select这个语句后，系统会随机抽取一个case进行判断
只要有其中一条语句正常retrun，此程序将立即执行
*/
