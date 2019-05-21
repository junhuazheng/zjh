package main

import "fmt"

type Point struct {
	X, Y float64
}

//想要修改p的值就得传指针类型“*Point”，如果这里传递不是指针的话将无法对其属性进行更改
func (p *Point) ScaleBy(factor float64) {

	p.X *= factor //等价于: X = X * factor， 对对象p进行操作，修改其私有属性
	p.Y *= factor
}

func main() {
	p := Point{100, 200}
	p.ScaleBy(2) //直接调用
	fmt.Println(p)

	p1 := Point{100, 200} //声明结构体后再用指针指向
	p2 := &p1
	p2.ScaleBy(2)
	fmt.Println(p2)
}
