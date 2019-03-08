package main

import "fmt"
import "time"

func main() {
	i := 0

	for { //for后面不写任何东西，这个循环条件永远为真，死循环
		i++
		time.Sleep(time.Second) //延时1秒

		if i == 5 {
			break //跳出循环，如果嵌套多个循环，跳出最近的那个内循环
		}
		fmt.Println("i = ", i)
	}
}
