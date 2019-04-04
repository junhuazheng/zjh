package main

import (
	"fmt"
	"net/http"
)

//服务端编写的业务逻辑处理程序
func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func myHandler1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "ni li hai")
}

func myHandler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "I am superman")
}

func main() {

	http.HandleFunc("/go", myHandler)

	http.HandleFunc("/mike", myHandler1)

	http.HandleFunc("/go/mike/yoyo", myHandler2)

	//在指定的地址进行监听，开启一个HTTP
	http.ListenAndServe("127.0.0.1:8000", nil)
}
