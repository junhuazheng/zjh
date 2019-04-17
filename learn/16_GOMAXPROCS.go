/*
GOMAXPROCS
Golang默认所有任务都运行在一个cpu核里
如果要在goroutine中使用多核，可以使用runtime.GOMAXPROCS函数修改，当参数小于1时使用默认值
*/
package main

import (
	"fmt"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(1)
}

func main() {
	//任务逻辑 ...
}
