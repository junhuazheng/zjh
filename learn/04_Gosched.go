package main

import (
	"fmt"
	"runtime"
)

func main() {

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("子协程")
		}
	}()

	for i := 0; i < 2; i++ {
		fmt.Println("主协程")
		//让出时间片，先让别的协程执行，它执行完再回来执行此协程
		runtime.Gosched()
	}
}
