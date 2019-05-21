//接受者函数
package main

import "fmt"

const Pi = 3.14

type Square struct { //定义一个长方形
	length, width float64
}

type Trapezoid struct { //定义一个梯形
	UpperLength float64
	BelowLength float64
	Height      float64
	Waist       float64
}

type Cirle struct {
	radius float64
}

func (a Square) Area() float64 {
	return a.length * a.width
}

func (a Trapezoid) Area() float64 {
	return (a.UpperLength + a.BelowLength) * a.Height / 2
}

func (a Cirle) Area() float64 {
	return Pi * a.radius * a.radius
}

func (a Square) Perimeter() float64 {
	return (a.width + a.length) * 2
}

func (a Trapezoid) Perimeter() float64 {
	return a.UpperLength + a.BelowLength + a.Waist*2
}

func (a Cirle) Perimeter() float64 {
	return 2 * Pi * a.radius
}

func main() {
	a := Square{10, 20}
	fmt.Printf("正方形的面积是[%v]\n", a.Area())
	fmt.Printf("正方形的周长是[%v]\n", a.Perimeter())

	b := Trapezoid{10, 20, 5, 15}
	fmt.Printf("梯形的面积是[%v]\n", b.Area())
	fmt.Printf("梯形的周长是[%v]\n", b.Perimeter())

	c := Cirle{5}
	fmt.Printf("圆形的面积是[%v]\n", c.Area())
	fmt.Printf("圆形的周长是[%v]\n", c.Perimeter())
}
