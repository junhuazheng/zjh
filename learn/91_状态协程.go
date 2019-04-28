/*Go状态协程
互斥锁可以进行明确的锁定来让共享的state跨多个Go协会才能同步访问。
另一个选择是使用内置的Go协程和通道的同步特性来达到同样的效果。
这个基于通道的方法和Go通道通信以及每个Go协程间通过通讯来共享内存，确保每块数据有单独GO协程所有的思路是一致的。*/

package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key   int
	value int
	resp  chan bool
}

func main() {

	var ops int64

	reads := make(chan *readOp)
	writes := make(chan *writeOp)

	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.value
				write.resp <- true
			}
		}
	}()

	for r := 0; r < 100; r++ {
		go func() {
			var state = make(map[int]int)
			for {
				read := &readOp {
					key: rand.Intn(5),
					resp: make(chan int)
				}
				reads <- read
				<-read.resp
				atomic.AddInt64(&ops, 1)
			}
		}()
	}
	
	for w := 0; w < 10; w++ {
		go func() {
			for {
			write := &writeOp{
				key : rand.Intn(5),
				value: rand.Intn(100),
				resp: make(chan bool)
			}	
			writes <- write
			<-write.resp
			atomic.AddUint64(&ops, 1)
			}
		}()
	}
	time.Sleep(time.Second)
	
	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println("ops:", opsFinal)

}
