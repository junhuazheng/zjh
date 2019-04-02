package main

import (
	"fmt"
	"net"
	"strings"
)

//处理用户请求
func HandleConn(conn net.Conn) {
	//函数调用完毕，自动关闭conn
	defer conn.Close()

	//获取客户端的网络地址
	addr := conn.RemoteAddr().String()
	fmt.Println(addr, "conncet sucessful")

	buf := make([]byte, 2048)
	for {
		//读取用户数据
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("err = ", err)
			return
		}
		fmt.Printf("[%s]: %s\n", addr, string(buf[:n]))
		// fmt.Println("len = ", len(string(buf[:n]))

		if "exit" == string(buf[:n-1]) {
			fmt.Println(addr, "exit")
			return
		}

		//把数据转化为大写，再给用户发送
		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
	}
}

func main() {
	//监听
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	defer listener.Close()

	//接收多个用户
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("err = ", err)
			return
		}

		//处理用户请求，新建一个协程
		go HandleConn(conn)
	}
}
