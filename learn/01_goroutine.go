package main

import (
	"fmt"
	"time"
)

func newTest() {
	for {
		fmt.Println("this is a newTest")
		time.Sleep(time.Second)
	}

}

func main() {

	go newTest() //新建一个协程，新建一个任务

	for {
		fmt.Println("this is a main goroutine")
		time.Sleep(time.Second)
	}

}
