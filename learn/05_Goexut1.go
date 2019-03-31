package main

import (
	"fmt"
	"runtime"
)

func test() {
	defer fmt.Println("ccc")

	runtime.Goexit() //终止所在的协程

	// return 终止函数
	fmt.Println("ddd")
}

func main() {

	go func() {

		fmt.Println("aaa")

		test()

		fmt.Println("bbb")
	}()

	//特地写一个死循环，目的不让主协程结束
	for {
	}

}
