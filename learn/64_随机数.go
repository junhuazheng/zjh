//结合时间模块初始化种子，简单的随机数生成
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	//以时间作为初始化种子
	rand.Seed(time.Now().UnixNano())
}

func main() {

	for i := 0; i < 10; i++ {
		a := rand.Int() //随机抽取整数，但是，int()里面不能放入参数
		fmt.Println(a)
	}

	for i := 0; i < 10; i++ {
		a := rand.Intn(100) //100范围内随机抽取
		fmt.Println(a)
	}

	for i := 0; i < 10; i++ {
		a := rand.Float32() //随机抽取浮点数
		fmt.Println(a)
	}
}
