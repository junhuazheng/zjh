//面向过程实现方法
//面向过程中没有“方法”概念，只能通过结构体和函数，由使用者使用函数参数和调用关系来形成接近方法的概念：
package main

import "fmt"

type Bag struct {
	items []int
}

type Aag struct {
	tems []int
}

func Insert(b *Bag, itemid int) { //面向过程实现方法
	b.items = append(b.items, itemid) //用append()将itemid添加到Bag的itens成员中
}

func (a *Aag) Onsert(temid int) { //面对对象的方法
	a.tems = append(a.tems, temid)
}

func main() {

	bag := new(Bag)

	Insert(bag, 111)

	a := new(Aag)

	a.Onsert(222)

	fmt.Println(bag)
	fmt.Println(a.tems)

}
