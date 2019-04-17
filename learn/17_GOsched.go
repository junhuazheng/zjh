/*
Gosched
这个函数的作用是：
让当前goroutine让出cpu，当一个goroutine发生阻塞，
Go会自动把与该goroutine处于同一系统线程的其他goroutine转移到另一个系统线程上去，以使这些goroutine不阻塞。
*/

package main

import (
	"fmt"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(1) //使用单核
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
输出结果为：
a: 0
a: 1
b
a: 2
a: 3
*/
