/*通过container/list包的New方法初始化list
变量名 := list.New()

列表与切片和map不同的是，列表兵没有具体元素类型的限制。
因此，列表的元素可以使任意类型。
这既带来便利，也会引来一些问题。给一个列表放入了非期望类型的值，在取出值后，将interface{}转换为期望类型事将会发生宕机。

双链表支持从队列前方或者后方插入元素，分别对应的方法是PushFront和PushBack。
两个方法都会返回一个*list.Element结构。如果在以后的使用中需要删除插入的元素，则只能通过
*list.Element配合Remove()方法进行珊瑚，这种方法可以让删除更加效率化，也是双链表特性之一。

l := list.New()
l.Pushback("fist")
l.Push.Frout(66)

*list.Element结构记录着列表元素的值及和其他节点之间的关系等信息，删除元素是，需要用到这个结构。*/

package main

import (
	"container/list"
	"fmt"
)

func main() {

	l := list.New()

	l.PushBack("can")

	l.PushFront(11)

	element := l.PushBack("fist")

	l.InsertAfter("high", element)

	l.InsertBefore("nonn", element)

	l.Remove(element)

	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value) //使用遍历返回的*list.Element的Value成员取得放入列表时的原值。
	}

}
