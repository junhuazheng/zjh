//多核测试：
package main

import (
	"fmt"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(4) //使用多核
}

func main() {

	exit := make(chan int)
	go func() {
		defer close(exit)
		go func() {
			fmt.Println("b")
		}()
	}()

	for i := 0; i < 4; i++ {
		fmt.Println("a:", i)

		if i == 1 {
			runtime.Gosched() //切换任务
		}
	}
	<-exit
}

/*
此电脑运行结果为：
b
a: 0
a: 1
a: 2
a: 3
*/
