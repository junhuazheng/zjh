package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

//接收文件内容
func RecvFile(fileName string, conn net.Conn) {
	//新建文件
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println("os.Create err = ", err)
		return
	}

	buf := make([]byte, 1024*4)

	//接收多少，写多少，一点不差
	for {
		n, err := conn.Read(buf) //接收对方发过来的文件内容
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件接收完毕")
			} else {
				fmt.Println("conn.Read err = ", err)
			}
			return
		}

		if n == 0 {
			fmt.Println("n == 0 文件接收完毕")
			return
		}

		f.Write(buf[:n]) //往文件里面写入内容
	}

}

func main() {

	//监听
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Listen err = ", err)
		return
	}

	defer listener.Close()

	//阻塞等待用户连接
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("listener.Accept err = ", err)
		return
	}
	defer conn.Close()

	//读取用户发送的文件名
	buf := make([]byte, 1024)
	var n int
	n, err = conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err = ", err)
		return
	}

	fileName := string(buf[:n])

	//回复“ok”
	conn.Write([]byte("ok"))

	//接收文件内容
	RecvFile(fileName, conn)

}
