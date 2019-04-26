package main

import "fmt"
/*
goroutine在退出方面，不像线程和进程，不能通过某种手段强制关闭他们，只能等待goroutine主动退出。
*/

/*
1、使用 for-range 退出
for-range是使用频率很高的结构，常用它来遍历数据
range能够感知channel的变比，当cahnnel被发送数据的协程关闭时，range就会结束，接着退出for循环

并发中的使用场景：
当协程只从1个channel读取数据，然后进行处理，处理后协程退出。
*/

go func(in <-chan int) {
	
	for n := range in {
		fmt.Printf("Process %d\n", n)
	}
}(inCh)
//当in通道被关闭时，协程可自动退出

/*
2、使用 ,ok 退出
for-select也是使用频率很高的结构，select提供了多路复用的能力
所以for-select可以让函数具有持续多路处理多个channel的能力
但是select没有感知channel的关闭，这就引出了两个问题：
1）、继续在关闭的通道上读 ，会读到通道传输数据类型的零值，如果是指针类型，读到nil，继续处理还会产生nil
2）、继续在关闭的通道上写，将会panic

问题2可以这样解决：
通道只由发送方关闭，接收方不可关闭，即某个写通道只由使用该select的协程关闭，select中就不存在继续在关闭的通道上写数据的问题。

问题1可以使用 ,ok 来检测通道的关闭，使用情况有两种。*/

//第一种：如果某个通道关闭后，需要退出协程，直接return即可。
go func() {
	
	for {
		select {
			case n, ok := <- in:
			if !ok{
				return
			}
			fmt.Printf("Process:%d\n", n)
			processedCnt++
			case <-t.C:
			fmt.Printf("Wroking,processedCnt = %d\n", processedCnt)
		}
	}
}()

//第二种，如果某个通道关闭了，不再处理该通道，而是继续处理其他case，退出是等待所有的可读通道关闭。
//我们需要使用select的一个特征：select不会再nil的通道上进行等待。这种情况，把只读通道设置为nil即可解决。
go func() {
	
	for {
		select {
			case n,ok := <- in1:
			if !ok {
				in1 = nil
			}
			
			case m, ok := <-in2:
			if !ok{
				in2 = nil
			}
			case <-t.C:
			fmt.Printf("W,p = %d\n", p)
		}
		
		if in1 == nil&&in2 == nil {
			return
		}
	}
}()

/*
3、使用退出通道退出
使用 ,ok 来退出使用for-select协会才能

解决是    *当读入数据的通道关闭时没数据读时程序的正常结束*

想想下面这两种场景 ,ok 还适用吗？
1）、接收的协程要退出了，如果它直接退出，不告知发送协程，发送协程将阻塞
2）、启动了一个工作协程处理数据，如何通知它退出？

使用一个专门的通道，发送退出信息，可以解决这类问题。
*/

func worker(stopCh <-chan struct{}) {
	go func() {
		defer fmt.Println("worker exit")
		
		for {
			select{
				case <-stopCh:
				fmt.Println("Recv stop signal")
				return
				case <-t.C
				fmt.Println("working")
			}
		}
	}()
	return
}