//get请求可以直接http.Get方法
package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"reflect"
)

func main() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		// handle error
		log.Println(err)
		return
	}

	defer resp.Body.Close()

	headers := resp.Header
	for k, v := range headers {
		fmt.Printf("k = %v, v = %v\n", k, v)
	}

	fmt.Printf("resp status %s, statusCode %d\n", resp.Status, resp.StatusCode)

	fmt.Printf("resp Proto %s\n", resp.Proto)

	fmt.Printf("resp Cntent length %d\n", resp.ContentLength)

	fmt.Printf("resp transfer encoding %v\n", resp.TransferEncoding)

	fmt.Printf("resp Uncompressed %t\n", resp.Uncompressed)
	
	fmt.Println(reflect.TypeOf(resp.Body)) //*http.gzipReader
	
	buf := bytes.NewBuffer(make([byte, 0, 512]))
	
	length, _ := buf.ReadFrom(resp.Body)
	
	fmt.Println(len(buf.Bytes()))
	fmt.Println(length)
	fmt.Println(string(buf.Bytes()))
}
