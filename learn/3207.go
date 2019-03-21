package main

//Go语言中的方法（Menthod）是一种作用于特定类型变量（这种特定类型变量叫做接收器Receiver）的函数
//在Go语言中，接收器的类型可以使任何类型，不仅仅是结构体，任何类型都可以拥有方法
//本例子中将会使用背包作为“对象”，将武平放入背包的过程作为“方法”
//通过面向过程的方式和Go语言中结构体的方式来理解“方法”的概念
//面向过程中没有“方法”的概念，只能通过结构体和函数，由使用者使用函数参数和调用关系来形成接近“方法“的概念

//声明Bag结构，这个结构体包含一个整型切片类型的items的成员
type Bag struct {
	items []int
}

//将一个物品放入背包的过程
//定义了Insert()函数，这个函数拥有两个参数，一个是背包指针（*Bag），一个是物品Id（itemid）
func Insert(b *Bag, itemid int) {
	b.items = append(b.items, itemid) //用append()将itermid添加到Bag的items成员中，模拟往背包添加物品的过程
}

func main() {

	bag := new(Bag) //创建背包实例bag

	Insert(bag, 1001) //调用Inster()函数，第一个参数放入背包，第二个参数放入物品ID
}

//Insert()函数将*Bag参数放在第一位，强调Insert会操作*Bag结构体。
