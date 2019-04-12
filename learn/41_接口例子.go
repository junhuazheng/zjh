package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	redius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (c circle) area() float64 {
	return math.Pi * c.redius * c.redius
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.redius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {

	r := rect{width: 3, height: 4}
	c := circle{redius: 5}

	measure(r)
	measure(c)
}

/*Go支持在struct类型上定义方法。area方法有一个*rect类型的接收器。
可以为指针或值接收器定义方法，这里是一个值接收器的例子。

接口是方法签名的命名集合。这里是集合形状（geometry）的基本接口。
对于这个例子，将在rect和circle类型上实现这个接口。

要在Go中实现一个接口，需要实现接口中的所有方法。这里在rect上实现geometry的一个实例，一级circle的实现

如果变量具有接口类型，那么可以调用命名接口中的方法。
circle和rect结构类型都实现了几何（geometry）接口，因此可以使用这些就结构体的实例上调用maesure（）函数*/
