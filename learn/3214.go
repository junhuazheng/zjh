//当方法作用于非指针接收器时，Go语言会在代码运行时将接收器的值复制一份。
//非指针接收器的方法中可以获取接收器成员的值，但修改后无效

//点（Point）使用结构体描述时，为点添加Add()方法，这个方法不能修改Point的成员X、Y变量，而是在计算后返回新的Point对象
//Point属于小内存对象，在函数返回值的复制过程中可以极大地提高代码运行效率：

package main

import "fmt"

//定义点结构
type Point struct {
	X int
	Y int
}

//非指针接收器的加方法(为Point结构定义一个Add（）方法。传入和返回都是点的结构，可以方便地视线U盾讴歌点连续相加的效果。
func (p Point) Add(other Point) Point {

	//成员值与参数相加后返回新的结构
	return Point{p.X + other.X, p.Y + other.Y}
}

func main() {

	//初始化点
	p1 := Point{1, 1}
	p2 := Point{3, 3}

	//p1p2相加后返回结果
	result := p1.Add(p2)

	//输出结果
	fmt.Println(result)
}

//由于例子中使用了非指针接收器，Add()方法变得类似于只读的方法，Add()方法内部不会对成员进行任何修改
