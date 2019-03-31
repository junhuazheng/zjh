package main

import (
	"fmt"
	"time"
)

func main() {

	defer fmt.Println("主协程也结束")

	ch := make(chan string)

	go func() {

		defer fmt.Println("子协程调用完毕")

		for i := 0; i < 2; i++ {
			fmt.Println("子协程 i = ", i)
			time.Sleep(time.Second)
		}

		ch <- "我是子协程，该你啦老铁"

	}()

	str := <-ch //没有数据前阻塞
	fmt.Println(str)

}
