/*速率限制是控制资源利用和维持服务质量的重要机制。

通过goroutines，channel和ticker都可以优雅地支持速率限制。

首先来看一下基本速率限制。假设想限制对传入请求的处理。我们会在相同名称的通道上放送这些要求。*/

package main

import (
	"fmt"
	"time"
)

func main() {

	request := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		request <- i
	}
	close(request)

	limiter := time.Tick(time.Millisecond * 200)

	for req := range request {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			burstyLimiter <- t
		}
	}()

	burstyRequest := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequest <- i
	}
	close(burstyRequest)

	for req := range burstyRequest {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}

/*这里限制通道将每200毫秒接收一个值。这是速率限制方案中的调节器
通过在服务每个请求之前紫色来自限制器信通的接受，我们限制自己每200毫秒接收一个请求。

我们可能虚妄在速率限制方案中允许短脉冲喘请求，同事保持总体速率限制。可以通过缓冲的限制器通道来实现
这个bustyLimter通道将允许最多3个事件的突发。

每200毫秒，将尝试向bustyLimiter添加一个新值，最大限制为3、现在模拟5个更多的传入请求。
这些传入请求中的前3个位超过bustyLimiter值。

运行程序后，就会看到第一批请求每200毫秒处理一次。
对于第二批请求，程序会立即服务前3个，因为突发速率限制，然后剩余2服务都具有200ms延迟。*/
