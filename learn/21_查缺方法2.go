/*指针类型的接受器由一个结构体的指针组成，更接近于面向对象中的this或者self
由于指针的特性，调用方法时，修改接收器指针的任意成员变量，在方法结束后，修改都是有效的。*/
package main

import "fmt"

type Prop struct {
	value int
}

type Point struct {
	X int
	Y int
}

func (po Point) Add(ot Point) Point {
	return Point{po.X + ot.X, po.Y + ot.Y}
}

func (p *Prop) Set(v int) {
	p.value = v
}

func main() {
	p := new(Prop)

	p.Set(100)

	fmt.Println(p.value)

	po1 := Point{1, 1}
	po2 := Point{2, 2}

	result := po1.Add(po2)

	fmt.Println(result)

}

/*当方法作用于费指针接收器时，Go语言会在代码运行时讲接收器的值复制一份，在分支真接收器的方法中
可以获取接收器的成员值，但修改后无效。*/
