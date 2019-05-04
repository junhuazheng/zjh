/*
package main

import (
	"fmt"
)

type Rectangle struct {
	width, height float64
}

func area(r Rectangle) float64 {
	return r.width * r.height
}

func main() {
	r1 := Rectangle{12, 3}
	r2 := Rectangle{9, 2}

	fmt.Println("Area of r1 is :", area(r1))
	fmt.Println("Area of r2 is :", area(r2))
}
*/

package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func main() {

	r1 := Rectangle{12, 3}
	r2 := Rectangle{9, 2}

	c1 := Circle{12}
	c2 := Circle{22}

	fmt.Println("Area of r1 is:", r1.area())
	fmt.Println("Area of r2 is:", r2.area())
	fmt.Println("Area of c1 is:", c1.area())
	fmt.Println("Area of c2 is:", c2.area())
}

/*
1、虽然method的名字一模一样，但是如果接收者不一样，那么method就不一样
2、method里面可以访问接收者的字段
3、调用method通过 . 访问，就像struct里面访问字段一样

上面的例子中，method area()分别属于REctangle和Circle
于是他们的Receiver就变成了Rectangle和Circle
或者说，这个area()方法是由REctangle/Circl发出的

例子中的method，此处的方法的Receiver是以值传递，而非引用传递
是的，Receiver还可以是指针，两者的差别在于，指针作为Receiver会对实例对象的内容发生操作。
而普通类型作为Receivcer仅仅是以副本作为操作对象，并不对原实例发生操作。
*/
