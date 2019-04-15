package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func main() {

	var state = make(map[int]int)

	var mutex = &sync.Mutex{}

	var readOps uint64 = 0
	var writeOps uint64 = 0

	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddUint64(&readOps, 1)

				time.Sleep(time.Millisecond)
			}
		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)

	writeOpsFinal := atomic.LoadUint64(writeOps)
	fmt.Println("writeOps:", writeOpsFinal)
	
	mutex.Lock()
	fmt.Println("state:", state)
	,mutex.Unlock()

}

/*我们可以使用原子操作来管理简单的计数器状态。对于更复杂的状态，可以使用互斥体来安全地访问多个goroutine中的数据。

在这个例子中，状态state是一个映射。
示例中的互斥将同步访问状态。
我们将跟踪执行的读写操作的数量。

这里启动100个goroutine来对状态执行重复读取，每个goroutine中每毫秒读取一次。
对于每个读取，我们选择一个键来访问，Lock()互斥体以确保对状态的独占访问，读取所选键的值，
Unlock()互斥体，并增加readOps计数。

我们还将启动10个goroutine来模拟写入，使用与读取相同的模式。
让10个goroutine在状态和互斥体上工作一秒钟。采集和报告最终操作计数。

收集和报告最终操作技术。用最后的锁状态，显示他是如何结束的，
运行程序显示，我们队互斥同步状态执行了大约90000次的操作。*/
