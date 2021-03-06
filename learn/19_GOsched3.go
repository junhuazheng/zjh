/*
Gosched再说明
Gosched用来让出CPU的时间片。就像是接力赛跑，A跑了一会
碰到runtime.Gosched()，就会把时间片给B，让B接着跑。
*/
package main

import (
	"fmt"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(1)
}

func say(s string) {
	for i := 0; i < 2; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}

/*
输出结果为：
hello
world
hello

注意：
1、先输出了hello，然后输出了world
2、hello输出了两个，world输出了一个，因为第二个hello食醋胡玩，主线程就退出了。

把代码中的runtime.Gosched()注释掉，执行结果是：
hello
hello
因为say("hello")这句话占用了时间，等它执行完，线程也结束了，say("world")就没有机会了

这里同时可以看出，go中的goroutine并不是同时在运行。事实上，如果没有在代码中通过
runtime.GOMAXPROCS(n)其中n是整数，指定使用多核的话，goroutines都是在一个线程的，
他们之间通不停地让出时间片轮流运行，打到类似同时运行的效果
*/
