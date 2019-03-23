//声明结构体内嵌

//计算机图形学中的颜色有两种类型，一中是包含红、绿、蓝三原色的基础颜色；
//另一种实在基础颜色之外增加透明度的颜色。透明度在颜色中叫Alpha，范围为0~1之间
//0表示完全透明，1表示不透明。
//使用传统的结构体字段的方法定义基础颜色和带有透明度颜色的过程代码如下：

package main

import "fmt"

type BasicColor struct {
	R, G, B float32
}

type Color struct {
	Basic BasicColor

	Alpha float32
}

func main() {

	var c Color //实例化一个完整的颜色结构

	//访问基础颜色并赋值
	c.Basic.R = 1
	c.Basic.G = 1
	c.Basic.B = 0

	c.Alpha = 1

	fmt.Printf("%+v", c)

}
