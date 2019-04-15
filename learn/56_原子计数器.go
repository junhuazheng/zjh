/*使用 sync/atomic 包来实现由多个 goroutine 访问的原子计数器。

使用一个无符号整数表示计数器（正数）

为了模拟并发更新，将启动50个 goroutine ， 每个增量计数器大约是 1 毫秒。

为了原子地递增计数器，这里使用 AddUint64() 函数，在ops计数器的内存地址上使用 & 语法。
在增量之间每等待一秒，允许一些操作累计。

为了安全地使用计数器，同时它仍然被其他 goroutine 更新， 通过 LoadUint64
提取一个当前值得副本 opsFinal。 如上所述，需要将获取值得内存地址&ops给这个函数。

运行程序显示执行了大约40000次操作。*/

package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {

	var ops uint64 = 0

	for i := 0; i < 50; i++ {
		go func() {
			for {
				atomic.AddUint64(&ops, 1)

				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	opsFinal := atomic.LoadInt64(&ops)

	fmt.Println("ops:", opsFinal)
}
