/*
尽管Go编译器产生的是本地可执行代码，这些代码仍旧运行在Go的runtime当中
（这部分的代码可以在runtime包中找到）

这个runtime类似Java和.NET语言所用到的虚拟器，它负责管理包括内存分配、垃圾回收
栈处理、goroutine、channel、切片、map和反射等等。

runtime调度器，关于runtime包几个方法：

1、Gosched ： 让当前线程让出cpu以让其他线程运行，它不会挂起当前线程，因此当前线程未来会继续执行

2、NumCPU ： 返回当前系统的cpu核数量。

3、GOMAXPROCS ： 设置最大的可同时使用的CPU核数。

4、Goexit ： 退出当前goroutine（但是defer语句会照常执行）

5、NumGoroutine ： 返回正在执行和排队的任务总数

6、GOOs ： 目标操作系统
*/

//NumCPU

package main

import (
	"fmt"
	"runtime"
)

func main() {

	fmt.Println("cpus:", runtime.NumCPU())
	fmt.Println("goroot:", runtime.GOROOT())
	fmt.Println("archice:", runtime.GOOS)
}
