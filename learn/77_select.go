package main

import "fmt"

/*举个例子：
1、妈妈喊你吃饭，你去吃饭
2、时间到了，要睡觉
3、没事做，打豆豆
在Golang里，select就是干这个事情的：到吃饭了去吃饭，该睡觉了就睡觉，没事做就打豆豆。

select的功能
在多个通道上进行读或写操作，让函数可以处理多个事情，但1次只处理1个
1、每次执行select，都会只执行其中1个case或者执行default语句。
2、当没有case或者default可以执行时，则随机选择1个case执行
3、case后面跟的必须是读或者写通道的操作，否则编译错误。

select由select和case组成，default不是必须的：*/
func main() {
	readCh := make(chan int, 1)
	writeCh := make(chan int, 1)

	y := 1
	select {
	case x := <-readCh:
		fmt.Printf("Read %d\n", x)
	case writeCh <- y:
		fmt.Printf("write %d\n", y)
	default:
		fmt.Println("Do what you want")
	}
}

/*输出结果为： write 1*/
/*我们创建了 readCh 和 writeCh 2个通道
1、readCh中没有数据，所以case x := <-readCh读不到数据，所以这个case不能执行。
2、writeCh是带缓冲区的通道，它里面是空的，可以写入1个数据，所以case writeCh <- y可以执行
3、有case可以执行，所以default不会执行*/
