/*优化：使用等待组替代通道简化同步
通道的内部实现代码在Go语言开发包的src/runtime/chan.go中，经过分析后大概了解到通道也是用常见的互斥量等进行同步
因此通道虽然是一个语言级特性，但也不是被神话的特性
通道的运行和使用都要比传统互斥量、等待组（sync.WaitGroup）有一定的消耗。
使用等待组调整来实现同步：*/
package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

//套接字接收
func soRe(conn net.Conn, wg *sync.WaitGroup) {

	buf := make([]byte, 1024)

	for {

		_, err := conn.Read(buf)

		if err != nil {
			break
		}
	}
	//函数已经结束，发送通知（接收完成后，使用wg.Done()方法将等待组计数器减一）
	wg.Done()
}

func main() {

	//连接地址
	conn, err := net.Dial("tcp", "www.111.com:80")
	if err != nil {
		fmt.Println("net.Dial err = ", err)
		return
	}

	//退出通道（声明退出同步用的等待组）
	var wg sync.WaitGroup

	//添加一个任务（为等待组的计数器加1，表示需要完成一个任务）
	wg.Add(1)

	go soRe(conn, &wg) //将等待组的指针传入函数

	time.Sleep(time.Second)

	conn.Close()

	//等待goroutine退出完毕
	wg.Wait()
	fmt.Println("recv done")
}
