//不使用oop的思路计算面积
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

func SquareArea(a Square) float64 {
	return a.length * a.width
}

func TrapezoidArea(a Trapezoid) float64 {
	return (a.UpperLength + a.BelowLength) * a.Height / 2
}

func CirleArea(a Cirle) float64 {
	return Pi * a.radius * a.radius
}

func SquarePerimeter(a Square) float64 {
	return (a.width + a.length) * 2
}

func TrapezoidPerimeter(a Trapezoid) float64 {
	return a.UpperLength + a.BelowLength + a.Waist*2
}

func CirlePerimeter(a Cirle) float64 {
	return 2 * Pi * a.radius
}

func main() {
	a := Square{10, 20}
	fmt.Printf("正方形的面积是[%v]\n", SquareArea(a))
	fmt.Printf("正方形的周长是[%v]\n", SquarePerimeter(a))

	b := Trapezoid{10, 20, 5, 15}
	fmt.Printf("梯形的面积是[%v]\n", TrapezoidArea(b))
	fmt.Printf("梯形的周长是[%v]\n", TrapezoidPerimeter(b))

	c := Cirle{5}
	fmt.Printf("圆形的面积是[%v]\n", CirleArea(c))
	fmt.Printf("圆形的周长是[%v]\n", CirlePerimeter(c))
}
