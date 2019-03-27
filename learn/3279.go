//使用通道代替Socket视线RPC的过程
//客户端与服务器运行在同一个进程，服务器和客户端在两个goroutine中运行

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

	go PRCSever(ch)

	recv, err := RPCClient(ch, "hi")

	if err != nil {
		fmt.Println()
	} else {
		fmt.Println("client received", recv)
	}
}
