//使用http POST 方法可以直接使用http.Post或者http.PostForm

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	resp, err := http.Post("http://www.maimaiche.com/loginRegister/login.do",
		"application/x-www-form-urlencoded"), strings.NewReader("mobile=xxxxxxxx&isRemberPwd=1")

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
