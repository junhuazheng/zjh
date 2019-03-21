//模拟父级构造调用
//黑猫是一种猫，猫是黑猫的一中泛称。同事描述这两种概念时，就是派生，黑猫派生自猫的种类。
//使用结构体描述猫和黑猫的关系时，将猫（Cat）的结构体嵌入到黑猫（BlackCat）中
//表示黑猫拥有猫的特性，然后再使用两个不同的构造函数分别构造出黑猫和猫的两个结构体实例

type Cat struct {
	Color string
	Name  string
}

//定义BlackCat结构，并嵌入Cat结构体。BlackCat拥有Cat的所有成员，实例化后可以自由访问Cat的所有成员。
type BlackCat struct {
	Cat //嵌入Cat，类似于派生
}

//“构造基类”
//NEwCat()函数定义了Cat的构造过程，使用名字作为参数，填充Cat结构体
func NewCat(name string) *Cat {
	return &Cat{
		Name: name,
	}
}

//“构造子类”
//NewBlack()使用color作为参数构造返回BlackCat指针
func NewBlackCat(color string) *BlackCat {
	//实例化BalckCat结构，此时Cat也同时被实例化
	cat := &BlackCat{}
	//填充BlackCat中嵌入的Cat颜色属性。BlackCat没有任何成员，所有的成员都来自Cat
	cat.Color = color
	return cat
}

//这个例子中，Cat结构体类似于面向对象中的“基类”。
//BalckCat嵌入Cat结构体，类似于面向对象中的“派生”
//实例化时，BlackCat中的Cat也会一并被实例化
