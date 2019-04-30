// http.PostForm方法
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {

	postParam := url.Values{
		"mobile":      {"xxxxx"},
		"isRemberPwd": {"1"},
	}

	resp, err := http.PostForm("http:/www.maimaiche.com/loginRegister/login.do")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))
}
