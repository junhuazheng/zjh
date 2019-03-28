//var 通道变量 chan 通道类型

//通道实例 := make(chan 数据类型)

ch1 := make(chan int)         //创建一个整型类型的通道
ch2 := make(chan interface{}) //创建一个空接口类型的通道，可以存放任意格式

type Equip struct { /*一些字段*/
}
ch2 := make(chan *Equip) //创建Equip指针类型的通道，可以存放*Equip

//使用通道发送数据
//通道创建后，就可以使用通道进行发送和接收操作。

//1、通道发送数据的格式
//通道的发送使用特殊的操作符 <- ，将数据通过通道发送的格式为：
// 通道变量 <- 值
//通道变量：通过make创建好的通道实例
//值：可以是变量、常量、表达式或者函数返回值等。值的类型必须与ch通道的元素类型一致

//2、通过通道发送数据的例子
//使用make创建一个通道后，就可以使用 <- 向通道发送数据：

//创建一个空接口通道
ch := make(chan interface{})
//将0放入通道
ch <- 0
//将hellp字符串放入通道中
ch <- "hello"

//3、发送将持续阻塞知道数据被接收
//把数据往通道中发送时如果接收方一直都没有接收，那么发送操作将持续阻塞，
//GO程序运行时能智能地发现一些永远无法发送成功的语句并作出提示

package main

import (
	"fmt"
)

func main() {
	
	ch := make(chan int)
	
	ch <- 0
}

//运行代码，报错：
fatal error:all goroutines are asleep - deadlock!
//运行时发现所有的goroutine（包括main）都处于等待goroutine。
//也就是说所有的goroutine中的channel并没有形成发送和接收对应的代码

package main

import "fmt"

func main() {
	
	ch := make(chan int)
	
	go func() {
		
		fmt.Println("start goroutine")
		
		ch <- 0
		
		fmt.Println("exit goroutine")
	}()
	
	fmt.Println("wait goroutine")
	
	<- ch
	
	fmt.Println("all done")
}