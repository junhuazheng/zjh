//使用通道接收数据

//通道接收同样使用 <- ，通道接收有如下特性：

/*1、通道的收发操作在不同的两个goroutine间进行
由于通道的数据在没有接收方处理时，数据发送方也会持续阻塞，因此通道的接收必定在另外一个goroutine中进行*/

/*2、接收持续阻塞直到发送方发送数据
如果接收方接收时，通道中没有发送方发送数据，接收方也会发生阻塞，直到发送方发送数据为止*/

/*3、每次接收一个元素
通道一次只能接收一个数据元素*/

//通道接收数据的4种写法

//1)阻塞模式接收数据时，讲接收变量作为<-操作符的左值
// data := <-ch
//执行该语句时将会阻塞，直到接收到数据并赋值给data变量

/*2)非阻塞魔石接收数据
使用非阻塞方式从通道接收数据时，语句不会发生阻塞：*/
// data, ok := <- ch
/*data：表示接收到的数据。未接收到数据时，data为通道类型的零值
OK： 表示是否接收到数据
非阻塞的通道接收方法可能造成高的CPU占用
因此使用非常少。如果需要实现接收超时检测，可以配合select和计时器channel进行*/

/*3)接收任意数据，忽略接收的数据
阻塞接收数据后，忽略从通道返回的数据：*/
// <- ch
/*执行该语句时将会发生阻塞，直到接收到数据，但接受到的数据会被忽略。
这个方式实际上只是通过通道在goroutine间阻塞收发实现并发同步。*/

/*使用通道做并发同步的写法*/
package main

import "fmt"

func main() {

	//构建一个同步用的通道
	ch := make(chan int)

	go func() {

		fmt.Println("start goroutine0")

		/*匿名goroutine即将结束时，通过通道通知main的goroutine，
		这一句会一直阻塞直到main的goroutine接收为止*/
		ch <- 0

		fmt.Println("exit goroutine")
	}()

	fmt.Println("wait goroutine")

	//开启goroutine后，马上通过通道等待匿名goroutine结束
	<-ch

	fmt.Println("all done")
}
