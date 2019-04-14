package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string, 1)

	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "result 2"
	}()

	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout 2")
	}

}

/*在这个示例中，假设正在执行一个外部调用，2秒后在通道c1上返回其结果。
这里是select实现超市。res := <-c1 等待结果和 <-Time。
等待在超市1秒后发送一个值，由于选择继续准备好第一个接收，如果操作超过允许的1秒，则将按超时情况处理
如果允许更长的超时，如3s，那么c2的接收将成功，这里将会打印结果。
运行此程序显示第一个操作超时和第二个操作超时*/
