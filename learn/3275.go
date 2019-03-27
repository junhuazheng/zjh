//使用匿名函数创建goroutine

//在main()函数中创建一个匿名函数并未匿名函数启动goroutine
//匿名函数没有参数。代码将并行执行定时打印计数的效果

package main

import (
	"fmt"
	"time"
)

func main() {

	go func() {

		var times int

		for {

			times++
			fmt.Println("tick", times)

			times.Sleep(time.Second)
		}
	}()

	var input string
	fmt.Scanln(&input)
}
