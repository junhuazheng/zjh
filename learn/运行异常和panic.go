package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	String string
	Input  string
)

func main() {
	f := bufio.NewReader(os.Stdin) //读取输入的内容

	for {
		fmt.Print("请输入您的用户名>")
		Input, _ = f.ReadString('\n') //定义一行输入的内容分隔符
		if len(Input) == 1 {
			continue //如果用户输入的是一个空行就让用户继续输入。
		}

		fmt.Printf("您输入的是:%s", Input)
		fmt.Scan(Input, &String)
		if String == "stop" {
			break
		}
		if String == "z123" {
			fmt.Println("欢迎登陆！")
		} else {
			panic("您输入的用户不存在") //要求用户输入一个字符串，一旦输入的字符串不是“zhengjunhua”就让程序直接崩溃
		}
	}
}
