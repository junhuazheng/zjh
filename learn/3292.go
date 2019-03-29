/*循环接收
通道的数据接收可以借用for range 语句进行多个元素的接收操作，格式如下：
for data := range ch {

}
通道ch是可以进行遍历的，遍历的结果就是接收到的数据
数据类型就是通道的数据类型。通过for遍历获得的变量只有一个，即上面例子中的data

使用for从通道中接收数据*/

package main

import (
	"fmt"
	"time"
)

func main() {
	
	ch := make(chan int)
	
	//将匿名函数并发执行
	go func)() {
		
		for i := 3; i >= 0; i-- {
			
			//将3到0之间的数值一次发送到通道中
			ch <- i
			
			//每次发送后暂停一秒
			time.Sleep(time.Second)
		}
	}()
	
	//使用for从通道中接收数据
	for data := range ch {
		
		fmt.Println(data)
		
		//当接收到数值0时，停止接收。如果继续发送，由于接收goroutine已经退出，没有goroutine发送到通道，因此运行时会发生宕机报错
		if data == 0 {
			break
		}
	}
}