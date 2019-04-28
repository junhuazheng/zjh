//生成器

package main

import "fmt"

func playerGen(name string) func() (string, int) {

	hp := 150

	//返回创建的闭包
	return func() (string, int) {

		//将变量引用到闭包中
		return name, hp
	}
}

func main() {

	//创建一个玩家生成器
	gen := playerGen("high noon")

	//返回name，hp
	name, hp := gen()

	fmt.Println(name, hp)
}

/*输出结果为：
high noon 150
*/
