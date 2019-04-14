package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "tow"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}

/*Go语言的选择（select）可等待多个通道操作。
 对于这个例子，将选择两个通道。
每个通道将在一段时间后开始接收值， 以模拟阻塞在并发goroutines中执行的RPC操作。
我们将使用select同时等待这两个值，在每个值到达时打印他们。
执行实例程序得打的值是“one”， 然后是“two”。注意，总执行时间只有1~2秒，因为1-2秒Sleeps同时执行*/
