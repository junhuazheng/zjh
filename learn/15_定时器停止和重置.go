package main

import (
	"fmt"
	"time"
)

func main() {

	timer := time.NewTimer(3 * time.Second)

	timer.Reset(time.Second)
	<-timer.C
	fmt.Println("时间到")

}

func main01() {

	timer := time.NewTimer(3 * time.Second)

	go func() {
		<-timer.C
		fmt.Println("定时器时间到，子协程可以打印了")
	}()

	timer.Stop()

	for {

	}
}
