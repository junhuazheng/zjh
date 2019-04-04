package main

import (
	"fmt"
	"net/http"
)

func HandConn(w http.ResponseWriter, req *http.Request) {

	w.Write([]byte("hello go"))
	fmt.Println("req.Method = ", req.Method)
	fmt.Println("req.Body = ", req.Body)
	fmt.Println("req.URL = ", req.URL)
	fmt.Println("req.Header = ", req.Header)
}

func main() {

	http.HandleFunc("/", HandConn)

	http.ListenAndServe(":8000", nil)
}
