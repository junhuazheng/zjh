package main

import (
	"fmt"
)

/*
1、如何跳出for-select

在select内的break并不能跳出for-select循环。
看下面的例子，consume函数从通道inCh不停读数据
期待在inch关闭后退出for-select循环，但结果是永远没有退出。
*/

func consume(inCh <-chan int) {

	i := 
	for {
		fmt.Printf("for:%d\n", i)
		select {
		case x, open := <-inCh:
			if !open {
				break
			}
			fmt.Printf("read:%d\n", x)
		}
		i++
	}
	fmt.Println("combine-routine exit")

}

/*输出结果为：
for : 0
read : 0
for : 1
read : 1
...never stop
*/

/*既然break不能跳出for-select，那怎么办呢？
1）在满足条件的case内，使用return介绍协程，如果有结尾工作，尝试交给defer
2）在select外for内使用break跳出循环，如上个例子中的combine函数
3）使用goto，goto没有那么可怕，适当使用。*/

/*2、selcet{}永远阻塞

select{}的效果等价于直接从刚创建的通道读数据：
ch := make(chan int)
<-ch

永远阻塞有什么用呢？
防止main函数在子协程干完活之前就退出。*/

/*3、select的应用场景

1）无阻塞的读、写通道。即使通道时带缓存的，也是存在阻塞的情况
使用select可以完美的解决阻塞读写
2）给某个请求/处理/操作，设置超时时间，一旦超时时间内无法完成，则停止处理
3）select本色：多通道处理
*/

