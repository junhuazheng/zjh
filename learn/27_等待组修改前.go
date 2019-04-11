/*下面代码中尝试使用套接字的TCP协议连接一个网址，连接上后，进行数据接收
等待一段时间后主动关闭套接字。等待套接字所在的goroutine自然结束。
连接、关闭、同步goroutine主流程部分*/
package main

import (
	"fmt"
	"net"
	"net/http"
	"time"
)

//套接字接收部分
//连接后，需要不停地接收数据
func soRe(conn net.Conn, exit chan string) {

	//创建一个接收的缓冲
	buf := make([]byte, 1024)

	//不停接收数据
	for {

		//读
		_, err := conn.Read(buf)

		if err != nil {
			break
		}
	}

	//函数已结束，发送通知
	exit <- "recv exit"
}

func main() {

	//连接一个地址
	conn, err := net.Dial("tcp", "www.163.com:80")
	if err != nil {
		fmt.Println("net.Dial err = ", err)
		return
	}

	//创建退出通道
	exit := make(chan string)

	//套接字接收
	go soRe(conn, exit)

	time.Sleep(time.Second)

	//主动关闭套接字
	conn.Close()

	//等待goroutine退出完毕
	fmt.Println(<-exit)

}
