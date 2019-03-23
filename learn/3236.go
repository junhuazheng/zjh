//初始化内嵌结构体
//内嵌结构体初始化时，将结构体内嵌的类型作为字段名像普通结构体一样进行初始化

//车辆结构图的组装和初始化：
package main

import "fmt"

type Whell struct {
	Size int
}

type Engine struct {
	Power int
	Type  string
}

type Car struct {
	Whell
	Engine
}

func main() {

	c := Car{

		Whell: Whell{ //将Car的Whell字段使用Whell结构体进行初始化
			Size: 18,
		},
		Engine: Engine{ //将Car的Engine字段使用Engine结构体进行初始化
			Type:  "1.4T",
			Power: 143,
		},
	}

	fmt.Printf("%+v\n", c)
}
