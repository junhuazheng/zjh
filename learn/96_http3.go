package main

import (
	"fmt"
	"net"
)

func main() {

	//主动连接服务器
	conn, err := net.Dial("tcp", ":8000")
	if err != nil {
		fmt.Println("dial err = ", err)
		return
	}
	defer conn.Close()
	
	requestBuf := "GET / go HTTp/ 1.1\r\..."
	
	//先发请求包，服务器才会响应包
	conn.Write([]byte(requestBuf))
	
	//接受服务器回复的响应包
	buf := make([]byte, 1024*4)
	n, err := conn.Read(buf)
	if n == 0 {
		fmt.Println("Read err = ", err)
		return
	}
	
	//打印响应报文
	fmt.Printf("#%v#", string(buf[:n]))
}
