/*producer()保持不变，负责生产数据。

squre()也不变，负责计算平方值。

修改main()，启动3个square，这3个squre从producer生成的通道读数据
这是FAN-OUT。

增加merge()，入参是3个square各自写数据的通道，给这3个通道分别启动1个协程
把数据写入到自己创建的通道，并返回该通道，这是FAN-IN*/

package main

import (
	"fmt"
	"sync"
)

func producer(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, i := range nums {
			out <- i
		}
	}()
	return out
}

func square(inCh <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range inCh {
			out <- n * n
		}
	}()
	return out
}

func merge(cs ...<-chan int) <-chan int {
	out := make(chan int)

	var wg sync.WaitGroup

	collect := func(in <-chan int) {
		defer wg.Done()
		for n := range in {
			out <- n
		}
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go collect(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	/*错误方式：直接等待是bug，死锁，以为merge写了out，main没有读
	wg.Wait()
	close(out)*/

	return out
}

func main() {
	in := producer(1, 2, 3, 4)

	//FAN-OUT
	c1 := square(in)
	c2 := square(in)
	c3 := square(in)

	for ret := range merge(c1, c2, c3) {
		fmt.Printf("%3d", ret)
	}
	fmt.Println()
}

/*
FAN模式能够更好的利用CPU，提供更好的并发，提高Golang程序的并发性能

在协程比较费时时，FAN模式可以减少程序运行时间，同样的时间可以处理更多的数据。

但是，FAN模式不一定能提高性能
FAN模式可以提高CPU利用率
FAN模式不一定能提升效率，降低程序运行时间

适当使用带缓冲通道可以提高程序性能
我们当前的程序的瓶颈在FAN-IN，squre函数很快就完成
merge函数它把3个数据写入到1个通道时出现了瓶颈
再修改下代码
merge()中的out修改为：
out := make(chan int, 100)
使用带缓冲通道后，程序的性能有了较大的提升。
*/
