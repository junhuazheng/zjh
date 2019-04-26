package main

import (
	"fmt"
	"time"
)

func producer(n int) <-chan int {
	out := make(chan int)
	go func() {
		defer func() {
			close(out)
			out = nil
			fmt.Println("producer exit")
		}()

		for i := 0; i < n; i++ {
			fmt.Printf("send %d\n", i)
			out <- i
			time.Sleep(time.Millisecond)
		}
	}()
	return out
}

//consummer read data from in channel , print it, an print
//all process count in each second
func consumer(in <-chan int) <-chan struct{} {
	finish := make(chan struct{})

	t := time.NewTicker(time.Millisecond * 500)
	processedcnt := 0

	go func() {
		defer func() {
			fmt.Println("worker exit")
			finish <- struct{}{}
			close(finish)
		}()

		//in for-select using ok to exit goroutine
		for {
			select {
			case x, ok := <-in:
				if !ok {
					return
				}
				fmt.Printf("Process:%d\n", x)
				processedcnt++
			case <-t.C:
				fmt.Printf("Wroking,processedCnt = %d\n", processedcnt)
			}
		}
	}()

	return finish
}

func main() {

	out := producer(3)
	finish := consumer(out)

	<-finish
	fmt.Println("main exit")
}
