/*
select实现吃饭睡觉打豆豆：
eat()函数会启动1个协程，该协程先睡几秒，时间不定，然后喊你吃饭
main()函数中的sleep是个定时器，每3秒喊你吃饭1次，select则处理3种情况：

1、从eatCh中读数据，代表有人喊我吃饭，我要吃饭了
2、从sleep.C中读到数据，代表闹钟时间到了，我要睡觉
3、default是，没人喊我吃饭，也不到睡觉时间，我就打豆豆
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func eat() chan string {

	out := make(chan string)
	go func() {
		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
		out <- "Mom call you eating"
		close(out)
	}()
	return out
}

func main() {
	eatCh := eat()
	sleep := time.NewTimer(time.Second * 3)
	select {
	case s := <-eatCh:
		fmt.Println(s)
	case <-sleep.C:
		fmt.Println("Time to sleep")
	default:
		fmt.Println("Beat dou")
	}
}

/*
由于前两个case都要等待一会，所以不能执行，所以执行default，运行结果一直都是如下：
输出结果为：
Beat dou

把default和下面的打印注释掉，多运行几次，有时候会吃饭，有时候会睡觉
*/
