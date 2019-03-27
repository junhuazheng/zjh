//time.After

//延迟回调：

package main

import (
	"fmt"
	"time"
)

func main() {

	exit := make(chan int)

	fmt.Println("strart")

	time.AfterFunc(time.Second, func() {

		fmt.Println("one second after")

		exit <- 0
	})

	<-exit
}

//time.AfterFunc()函数实在time.After基础上增加了到时的回调，方便使用

//而time.After()函数又是在time.NewTimer()函数上进行的封装
