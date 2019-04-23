package main

/*
for range 中使用闭包
*/

import "fmt"

// func main() {
// 	s := []string{"a", "b", "c"}
// 	for _, v := range s {
// 		go func() {
// 			fmt.Println(v)
// 		}()
// 	}
// 	select {} //阻塞模式
// }

/*
输出结果为：
c
c
c
在没有将变量v的拷贝值传进匿名函数之前，只能获取最后一次循环的值
*/

//修改后
func main() {
	s := []string{"a", "b", "c"}
	for _, v := range s {
		go func(v string) {
			fmt.Println(v)
		}(v) //每次将变量v的拷贝传进函数
	}
	select {}
}

/*
输出结果为
c
b
a
由于使用了go协程，并非顺序输出
*/
