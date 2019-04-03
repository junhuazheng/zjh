package main

import (
	"fmt"
	"net"
)

type Client struct {
	C    chan string
	Name string
	Addr string
}

//保存在线用户 cliAddr ===》Client
var onlineMap map[string]Client

var messaage = make(chan string)

func WriteMsgToClient(cli Client, conn net.Conn) {
	for msg := range cli.C {
		conn.Write([]byte(msg + "\n"))
	}
}

func MakeMsg(cli Client, msg string) (buf string) {
	buf = "[" + cli.Addr + "]" + cli.Name + ":" + msg

	return
}

func HandleConn(conn net.Conn) { //处理用户连接，并且广播用户上线
	defer conn.Close()

	//获取客户端的网络地址
	cliAddr := conn.RemoteAddr().String()

	//创建一个结构体，默认，用户名和网络地址一样
	cli := Client{make(chan string), cliAddr, cliAddr}

	//把结构体添加到map
	onlineMap[cliAddr] = cli

	//新开一个协程，专门给当前客户端发送信息
	go WriteMsgToClient(cli, conn)

	//广播某个人在线
	//"[" + cli.Addr + "]" + cli.Name + ": login"
	messaage <- MakeMsg(cli, "login")

	//提示我是谁
	cli.C <- MakeMsg(cli, "I am here")

	go func() {
		buf := make([]byte, 2048)

		for {
			n, err := conn.Read(buf)
			if n == 0 { //对方断开，或者，出问题
				fmt.Println("conn.Read err = ", err)
				return
			}

			msg := string(buf[:n-1]) //通过windows，nc测试，多一个换行符

			if len(msg) == 3 && msg == "who" {

				//遍历map，给当前用户发送所有成员
				conn.Write([]byte("user list:\n"))
				for _, tmp := range onlineMap {
					msg = tmp.Addr + ":" + tmp.Name + "\n"
					conn.Write([]byte(msg))
				}
			} else {
				//转发此内容
				messaage <- MakeMsg(cli, msg)
			}
		}
	}()

	for {

	}
}

func Manager() {

	//给map分配空间
	onlineMap = make(map[string]Client)

	for {
		msg := <-messaage //没有消息前，阻塞

		//遍历map，给map每个成员都发送此消息
		for _, cli := range onlineMap {
			cli.C <- msg
		}

	}
}

func main() {
	//监听
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println("net.Listen err = ", err)
		return
	}

	defer listener.Close()

	//新开一个协程，转发消息，只要有消息来了，遍历map，给map每个成员都发送此消息
	go Manager()

	//主协程，循环阻塞等待用户连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err = ", err)
			continue
		}

		go HandleConn(conn) //处理用户连接，广播用户上线
	}
}
