package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"reflect"
)

func main() {

	addr := "www.baidu.com:80"         //定义主机名
	conn, err := net.Dial("tcp", addr) //拨号操作，需要指定协议。
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("访问公司IP地址是:", conn.RemoteAddr().String())

	fmt.Printf("客户端链接的地址及端口是: %v\n", conn.LocalAddr())

	fmt.Println("“conn.LocalAddr()”所对应的数据类型是:", reflect.TypeOf(conn.LocalAddr()))

	fmt.Println("“conn.RemoteAddr().String()”所对应的数据类型是:", reflect.TypeOf(conn.RemoteAddr().String()))

	n, err := conn.Write([]byte("GET / HTTP/1.1\r\n\r\n"))

	if err != nil {
		log.Fatal()
	}

	fmt.Println("向服务端发送的数据大小事:", n)

	buf := make([]byte, 1024)

	n, err = conn.Read(buf)

	if err != nil && err != io.EOF {
		log.Fatal(err)
	}
	fmt.Println(string(buf[:n]))
	conn.Close()
}
