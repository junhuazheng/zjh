package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println("http.Get err = ", err)
		return
	}

	defer resp.Body.Close()

	fmt.Println("Status = ", resp.Status)
	fmt.Println("StatusCode = ", resp.StatusCode)
	fmt.Println("Header = ", resp.Header)

	buf := make([]byte, 4*1024)
	var tmp string

	for {
		n, err := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("read err = ", err)
			break
		}
		tmp += string(buf[:n])
	}

	fmt.Println("tmp = ", tmp)
}
