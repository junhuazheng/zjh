//通道ch是可以进行遍历的，遍历的结果就是接收到的数据
//数据类型就是通道的数据类型
//通过for遍历获得的变量只有一个

//使用for从通道中接收数据：

package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int)

	go func() {

		for i := 3; i >= 0; i-- {

			ch <- i

			time.Sleep(time.Second)
		}
	}()

	for data := range ch {

		fmt.Println(data)

		if data == 0 {
			break
		}
	}
}
