//初始化内嵌匿名结构体
//在前面描述车辆和引擎的例子中，有时考虑编写代码的便利性
//会将结构体直接定义在嵌入的结构体中
//也就是说，结构体的定义不会被外部引用到。在初始化这个被嵌入的结构体时，就需要再次声明结构才能赋予数据。

package main

import "fmt"

type Whell struct {
	Size int
}

type Car struct {
	Wheel
	Engine struct { //原来的Engine结构体被直接定义在Car的结构体中。这种嵌入的写法就是将原来的结构体类型转换为struct{...}
		Power int
		Type  string
	}
}

func main() {

	c := Car{

		Wheel: Wheel{
			Size: 178,
		},
		Engine: struct {
			Power int
			Type  string
		}{
			Type:  "1.4T",
			Power: 143,
		},
	}
	fmt.Fprintf("%+v\n", c)
}
