package main

import (
	"fmt"
	"time"
)

//全局变量，创建一个channel
var ch = make(chan int)

func Printer(str string) {

	for _, data := range str {
		fmt.Printf("%c", data)
		time.Sleep(time.Second)
	}
	fmt.Printf("\n")
}

func person1() {
	Printer("hello")
	ch <- 666 //给通道写入数据，发送
}

func person2() {
	<-ch //从管道取数据，接收，如果通道没哟数据他就会阻塞
	Printer("world")
}

func main() {
	go person1()

	go person2()

	for {
	}
}
