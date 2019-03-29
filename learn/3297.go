/*Go语言RPC（模拟远程过程调用）

服务器开发中使用RPC（Remote Procedure Call，远程过程调用）简化进程间通信的过程。
RPC能有效地封装通信过程，让远程的数据收发通信过程看起来就像本地的函数调用一样。

下面使用通道代替Socket实现RPC的过程。
客户端与服务器运行在同一个进程，服务器和客户端在两个goroutine中运行*/

package main

import (
	"errors"
	"fmt"
	"time"
)

func RPCClient(ch chan string, req string) (string, error) {
	ch <- req

	select {
	case ack := <-ch:
		return ack, nil
	case <-time.After(time.Second):
		return "", errors.New("Time out")
	}
}

func RPCServer(ch chan string) {
	for {

		data := <-ch

		fmt.Println("server received:", data)

		ch <- "roger"
	}
}

func main() {

	ch := make(chan string)

	go RPCServer(ch)

	recv, err := RPCClient(ch, "hi")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("client received", recv)
	}
}
