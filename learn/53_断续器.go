/*计时器是当想在未来做一些事情， tickers是用于定期做一些事情、
这里是一个例行程序，周期性执行直到停止。

代码机使用与计时器的机制类似：发送值到通道。
这里我们将使用通道上的一个范围内来迭代值，这些之每500ms到达。
代码可以像计时器一样停止，当代码停止后，他不会在其他通道上接收任何更多的值。我们将在1600ms吼停止。
当运行这个程序时，ticker应该在我们停止之前打3次。*/

package main

import (
	"fmt"
	"time"
)

func main() {

	ticker := time.NewTicker(time.Millisecond * 500)

	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	time.Sleep(time.Millisecond * 1600)

	ticker.Stop()

	fmt.Println("Ticker stopped")
}

/* time.NewTicker 再利用专用通道 ticker.C来实现断续器*/
