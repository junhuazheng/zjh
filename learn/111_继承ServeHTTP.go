package main

import (
	"fmt"
	"net/http"
)

type BaseHander struct {
}

func (hander *BaseHander) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	fmt.Println("url path =>", req.URL.Path)
	fmt.Println("url param a =>", req.URL.Query().Get("a"))
	resp.Write([]byte("hello world"))
}

func main() {
	http.ListenAndServe(":8000", &BaseHander{})
}
