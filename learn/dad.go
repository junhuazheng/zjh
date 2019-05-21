//声明及定义的类型

package main

import "fmt"

type Chairman struct {
	Name string
	age  int
}

func main() {
	var Leader Chairman
	var DeputyCadress_1 *Chairman
	DeputyCadress_2 := new(Chairman) //用内置函数new进行了初始化&{}
	fmt.Println(Leader)
	fmt.Println(DeputyCadress_1)
	fmt.Println(DeputyCadress_2)
	fmt.Println(DeputyCadress_1 == nil)
	DeputyCadress_1 = new(Chairman) //这种先声明再初始化的方式和直接初始化效果相同
	fmt.Println(DeputyCadress_1)
	fmt.Println(DeputyCadress_1 == nil)
}
