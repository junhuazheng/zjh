/*
读写锁 RWMutex
读写锁是针对读写操作的互斥锁。
读写锁与互斥锁最大的不同就是可以分别对  读、写 进行锁定。
一般用在大量读操作、少量写操作的情况：
func (rw *RWMutex) Lock()
func (rw *RWMutex) Unlock()

func (rw *RWMutex) RLock()
func (rw *RWMUtex) RUnlock()

读锁定(Rlock)， 对读操作进行锁定
读解锁(RUnlock)， 对读锁定进行解锁

写锁定(Lock)， 对写操作进行锁定
写解锁(Unlock)， 对写锁定进行解锁

不要混用锁定和解锁
因为对未读锁定的读写锁进行读解锁或对未写锁定的读写锁进行写解锁将会引起运行时错误。

读写锁：
1、同时只能有一个goroutine能够获得写锁定。
2、同时可以有人以多个goroutine获得读锁定。
3、同时只能存在写锁定或读锁定（读和写互斥）

也就是说，当有一个goroutine获得写锁定，其他无论是读锁定还是写锁定都将阻塞直到写解锁；

当有一个goroutine获得读锁定，其他读锁定仍然可以继续；

当有一个或者人以多个读锁定，写锁定将等待所有读锁定解锁后才能进行写锁定；

所以说这里的读锁定（RLock）目的其实是告诉写锁定：很多人正在读取数据，需要等待他们读（读解锁）完再来写（写锁定）。
*/

package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var count int
var rw sync.RWMutex

func read(n int, ch chan struct{}) {
	rw.RLock()
	fmt.Printf("goroutine %d 进入读操作...\n", n)
	v := count
	fmt.Printf("goroutine %d 读取结束，值为： %d\n", n, v)
	rw.RUnlock()
	ch <- struct{}{}
}

func write(n int, ch chan struct{}) {
	rw.Lock()
	fmt.Printf("goroutine %d 进入写操作...\n", n)
	v := rand.Intn(1000)
	count = v
	fmt.Printf("goroutine %d 写入结束，新值为： %d\n", n, v)
	rw.Unlock()
	ch <- struct{}{}
}

func main() {
	ch := make(chan struct{}, 10)
	for i := 0; i < 5; i++ {
		go read(i, ch)
	}
	for i := 0; i < 5; i++ {
		go write(i, ch)
	}
	for i := 0; i < 10; i++ {
		<-ch
	}
}
