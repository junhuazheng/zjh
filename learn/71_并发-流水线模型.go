/*
Golang的并发核心思路：
Golang并发核心思路是关注数据流动。
数据流动的过程交给channel，数据处理的每个环节都交给goroutine
把这些流程画起来，有始有终形成一条线，那就能构成流水线模型。

流水线
在Golang中，流水线由多个阶段组成，每个阶段之间通过channel连接，每个节点可以由多个同时运行的gotouine组成。

举个例子，设计一个程序：
计算一个整数切片中元素的平方值并把它打印出来。
非并发的方式是使用for遍历整个切片，然后计算平方，打印结果。

使用流水线模型视线这个简单的功能，从流水线的角度，可以分为3个阶段：
1、遍历切片，这是生产者
2、计算平方值
3、打印结果，这是消费者。

下面的代码：
producer()负责生产数据，他会把数据写入通道，并把它写数据的通道返回。
square()负责从某个通道读数字，然后计算平方，将结果写入通道，并把他的输出通道返回。
main()负责启动producer和square，并且还是消费者，读取square的结果并打印出来。
*/

package main

import "fmt"

func producer(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
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

func main() {
	in := producer(1, 2, 3, 4)
	ch := square(in)

	//消费者
	for ret := range ch {
		fmt.Printf("%3d", ret)
	}
	// fmt.Println()
}

/*
输出结果为：
1 4 9 16
*/

/*流水线的特点
1、每个阶段吧数据通过channel传递给下一个阶段。
2、每个阶段要创建1个goroutine和1个channel，这个goroutine向里面写数据，函数要返回这个通道
3、有1个函数来组织流水线，我们的例子中是main函数。*/
