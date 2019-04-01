package main

import "fmt"

func main() {

	ch := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch) //不需要再写数据时，关闭channel
	}()

	for num := range ch {
		fmt.Println("num = ", num)
	}

	// for {
	// 	如果ok为true，说明管道没有关闭，为false说明管道已经关闭
	// 	if num, ok := <-ch; ok == true {
	// 		fmt.Println("num = ", num)
	// 	} else { 管道关闭
	// 		break
	// 	}
	// }
}
