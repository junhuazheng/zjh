package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	client := &http.Client{} //实例化HTTP客户端，请求需要通过这个客户端实例发送。

	//使用POST方式向网易的服务器创建一个HTTP请求，第三个参数为HTTP的Body部分
	//Body部分的内容来自字符串，但参数只能接受os.Reader类型
	//因此使用strings.NewReader()创建一个字符串的读取器，返回io.Reader接口作为htpp的Body部分
	//供Newrequest()函数读取。创建请求只是构造一个请求对象，不会连接网络

	req, err := http.NewRequest("POST", "http://www.163.com/",
		strings.NewReader("key = value"))

	//发现错误就打印并退出
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}

	//为创建好的HTTP请求的头部添加User-Agent，作用是表明用户的代理特性。
	req.Header.Add("User-Agent", "myClient")

	//使用客户端处理请求，此时client将HTTP请求发送到网易服务器。服务器响应请求后，将信息返回并保存到resp变量中
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}

	data, err := ioutil.ReadAll(resp.Body) //读取响应的Body部分并打印
	fmt.Println(string(data))

	defer resp.Body.Close()
}

//由于我们构造的请求不是网易服务器所支持的类型，所以服务器返回操作不被运行的405号错误
