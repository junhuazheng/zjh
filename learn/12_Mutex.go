package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan struct{}, 2)

	var l sync.Mutex
	go func() {
		l.Lock()
		defer l.Unlock()
		fmt.Println("goroutine: 锁定大概2s")
		time.Sleep(time.Second * 2)
		fmt.Println("goroutine: 解锁了，可以开抢了")
		ch <- struct{}{}
	}()

	go func() {
		fmt.Println("goroutine2: 等待解锁")
		l.Lock()
		defer l.Unlock()
		fmt.Println("goroutine2: 我锁定了")
		ch <- struct{}{}
	}()

	//等待goroutine执行结束
	for i := 0; i < 2; i++ {
		<-ch
	}
}

/*
输出结果为：
goroutine2； 等待解锁
goroutine： 锁定大概2s
goroutine：解锁了，可以开抢了
goroutine2： 我锁定了
*/

/*
sync包含一个Locker interface {
	Lock()
	Unlock()
}
该接口只有两个方法，Lock()和Unlock()。整个sync包都是围绕该接口实现。

互斥锁Mutex
互斥锁Mutex是Locker的一中具体实现，有两个方法：
func (m *Mutex) Lock()
func (m *Mutex) Unlock()
一个互斥锁只能同时被一个goroutine锁定。其他goroutine将阻塞
直到互斥锁被解锁，之后再重新争抢对互斥锁的锁定。

对一个未锁定的互斥锁解锁将会产生运行时错误。
*/
