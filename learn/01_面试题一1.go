package main

import "fmt"

func main() {
	defer_call()
}

func defer_call() {

	defer func() {
		fmt.Println("打印前")
	}()

	defer func() {
		fmt.Println("打印中")
	}()

	defer func() {
		fmt.Println("打印后")
	}()

	panic("触发异常")
}

/*在合格案例中，触发异常这几个字打印的顺序其实是不确定的。
defer，panic，recover一般都会配套使用来捕捉异常。*/

/*多个defer中，第一个defer是最后执行的。
有点像往瓶子里放石头，最先放的那个往往最后一个拿起来*/

/*输出结果为：
打印后
打印中
打印前
触发异常
*/
