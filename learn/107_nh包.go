//需要在请求的时候设置头参数、cookie之类的数据，就可以使用htyp.Do方法。
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {

	client := &http.Client{}

	req, err := http.NewRequest("POST", "http://www.baidu.com; charset=UTF-8", strings.NewReader("mobile=xxx&isRemberPwd=1"))
	if err != nil {
		log.Println(err)
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	resp, err := client.Do(req)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicln(err)
		return
	}

	fmt.Println(resp.Header.Get("Content-Type"))

	type Result struct {
		Msg    string
		Status string
		Obj    string
	}

	result := &Result{}
	json.Unmarshal(body, result) //解析json字符串

	if result.Status == "1" {
		fmt.Println(result.Msg)
	} else {
		fmt.Println("login error")
	}
	fmt.Println(result)
}
