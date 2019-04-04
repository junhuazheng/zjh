package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

type Client struct {
	C    chan string
	Name string
	Addr string
}

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

func HandleConn(conn net.Conn) {
	defer conn.Close()

	cliAddr := conn.RemoteAddr().String()

	cli := Client(make(chan string), cliAddr, cliAddr)

	onlineMap[cliAddr] = cli

	go WriteMsgToClient(cli, conn)

	messaage <- MakeMsg(cli, "login")

	cli.C <- MakeMsg(cli, "I am here")

	isQuit := make(chan bool)
	
	hasData := make(chan bool)

	go func() {
		buf := make([]byte, 2048)

		for {
			n, err := conn.Read(buf)
			if n == 0 {
				isQuit <- true
				fmt.Println("conn.Read err = ", err)
				return
			}
			
			msg := string(buf[:n-1])
			
			if len(msg) == 3 && msg == "who" {
				conn.Write([]byte("user list:\n"))
				for _, tmp := range onlineMap {
					msg = tmp.Addr + ":" + tmp.Name + "\"
					conn.Write([]byte(msg))
				}
			}else if len(msg) >= 8 && msg[:6] == "rename" {
				name := strings.Split(msg, "|")[1]
				cli.Name = name
				onlineMap[cliAddr] = cli
				conn.Write([]byte("rename ok\n"))
			}else {
				messaage <- MakeMsg(cli, msg)
			}
		}

	}()
	
	for {
		select {
			case <- isQuit:
			delete(onlineMap, cliAddr)
			messaage <- MakeMsg(cli, "login out")
			
			return
			
			case <- hasData:
			
			case <- time.After(30*time.Second):
			delete(onlineMap, cliAddr)
			messaage <- MakeMsg(cli, "time out leave out")
			return
		}
	}
}

func Manager() {
	onlineMap = make(map[string]Client)
	
	for {
		msg := <-messaage
		
		for _, cli := range onlineMap {
			cli.C <- msg
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println("net.Listen err = ", err)
		return
	}

	defer listener.Close()

	go Manager()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err = ", err)
			continue
		}

		go HandleConn(conn)
	}
}
