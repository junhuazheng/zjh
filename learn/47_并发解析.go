/*
Golang : 不要通过共享内存来通信，而应该通过通信来共享内存。
1、通过goroutine与sync.Mutex进行并发同步
*/
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var count int = 0

func counter(lock *sync.Mutex) {
	lock.Lock()
	count++
	fmt.Println(count)
	lock.Unlock()
}

func main() {
	lock := &sync.Mutex{}
	for i := 0; i < 10; i++ {
		//传递指针为了防止函数内的锁和调用的锁不一致
		go counter(lock)
	}
	for {
		lock.Lock()
		c := count
		lock.Unlock()
		//把时间片给别的goroutine 未来某个时刻运行该groutine
		runtime.Gosched()
		if c >= 10 {
			fmt.Println("goroutine end")
			break
		}
	}
}
