/*为了避免goroutine无限制地被创建，
这个例子中，需要为consumer()函数添加合理的退出条件：*/

package main

import (
	"fmt"
	"runtime"
)

func consumer(ch chan int) {

	for {

		data := <-ch

		//为无限循环设置退出条件，这里设置0为退出
		if data == 0 {
			break
		}

		fmt.Println(data)
	}

	fmt.Println("goroutine exit")
}

func main() {

	//传递数据用的通道
	ch := make(chan int)

	for {

		//空变量，什么也不做
		var dummy string

		//获取输入，模拟进程持续运行
		fmt.Scan(&dummy)

		//当命令输入quit时，进入退出处理的流程。
		if dummy == "quit" {

			/*runtime.NumGoroutine返回一个进程的所有goroutine数
			main()的goroutine也被算在里面。因此需要勾出main()的goroutine数
			剩下的goroutine为实际创建的goroutine数，对这些goroutine进行遍历*/
			for i := 0; i < runtime.NumGoroutine()-1; i++ {
				/*并发开启的goroutine都在竞争获取通道中的数据，
				因此只要知道多少个goroutine需要退出，就给通道里发多少个0*/
				ch <- 0
			}

			continue
		}

		go consumer(ch)

		fmt.Println("goroutines:", runtime.NumGoroutine())

	}
}
