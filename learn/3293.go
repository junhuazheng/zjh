/*Go语言并发打印（借助通道实现）
前面的例子创建的都是无缓冲通道。
使用无缓冲通道往里面装入数据时，装入方将被阻塞
直到另外通达在另外一个goroutine中被取出

同样，如果通道中没有放入任何数据，接收方试图从通道中获取数据时，同样也是阻塞
发送和接收的操作是同步完成的

通过一个并发打印的例子，将goroutine和channel放在一起展示他们的用法：*/

package main

import "fmt"

func printer(c chan int) {

	for {

		//从函数参数传入的通道中获取一个整型数值
		data := <-c

		if data == 0 {
			break
		}

		fmt.Println(data)
	}

	//在退出循环时，通过通道通知main()函数已经完成工作。（我搞定了！）
	c <- 0
}

func main() {

	//创建一个整型通道进行跨goroutine的通信
	c := make(chan int)

	//创建一个goroutine，并发执行printer()函数，传入channel
	go printer(c)

	//构建一个数值循环，将1~10的数通过通道传给printer构造出的goroutine
	for i := 1; i <= 10; i++ {
		c <- i
	}

	//给通道传入一个0，表示将前面的数据处理完成后，退出循环。（没数据啦）（通知并发的printer结束循环）
	c <- 0

	/*在数据发送过去后，因为并发和调度的原因，任务会并发执行。
	这里需要等待printer的c<-0返回数据后，才可以退出main()（我搞定啦！）*/
	<-c
}
