/*我们经常想在将来的某个时间点执行Go代码，或者在某个时间间隔重复执行。
Go的内置计时器和自动接收器功能使这两项任务变得容易。

定时器代表未来的一个时间。可告诉定时器你想要等待多长时间，
它提供了一个通道，在当将通知时执行对应程序。
在这个实例中，计时器将等待2秒钟。*/

package main

import (
	"fmt"
	"time"
)

func main() {

	timer1 := time.NewTimer(time.Second * 2)

	<-timer1.C

	fmt.Println("Timer 1 expired")

	timer2 := time.NewTimer(time.Second)

	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()
	stop2 := timer2.Stop()

	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
}

/* <-timer1.C 阻塞定时器的通道C，知道它发送一个指定定时器超时的值
如果只是想等待，可以使用time.Sleep。定时器可能起作用的一个原因是在定时器到期之前取消定时器。

第一个定时器将启动程序后约2秒国企，单第二个定时器应该在它到期之前停止，第29行*/

/* time.NewTimer 再利用专用通道 timer.C 来实现计时器*/
