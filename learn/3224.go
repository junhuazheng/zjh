//Go语言提供的http包里也大量使用了类型方法。GO语言使用http包进行HTTP的请求
//使用http包的NewRequest()方法可以创建一个HTTP请求，填充请求中的http头（req.Header）
//再调用http.Client的Do包方法，将传入的HTTP请求发送出去。

//创建一个HTTP请求，并且设定HTTP头
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	client := &http.Client{}

	req, err := http.NewRequest("POST", "http://www.163.com/",
		strings.NewReader("key = value"))

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}

	req.Header.Add("User-Agent", "myClient")

	resp, err := client.Do(req)
	
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}
	
	data, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))
	
	dafer resp.Body.Close()

}
