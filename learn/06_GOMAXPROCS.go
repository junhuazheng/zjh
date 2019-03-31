package main

import (
	"fmt"
	"runtime"
)

func main() {

	//n := runtime.GOMAXPROCS(4) 指定以4核运算
	n := runtime.GOMAXPROCS(1) //指定以单核运算
	fmt.Println("n = ", n)
}
