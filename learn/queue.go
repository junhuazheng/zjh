/*
队列是一种非常常见的数据结构，和我们日常生活中排队是基本一致的，都遵循FIFO(First In First Out)的原则

队列可以使用链表或者数组实现，使用链表的优点是扩容简单，缺点是无法通过索引定位元素
使用数组则相反，扩容不容易但是可以通过索引定位元素，这次我们使用双向链表实现

*/

package queue

type Queue interface {

	//获取当前链表长度
	Length() int

	//获取当前链表容量
	Capacity() int

	//获取当前链表头结点
	Front() *Node

	//获取当前链表尾结点。
	Rear() *Node

	//入列
	Enqueue(value interface{}) bool

	//出列
	Dequeue() interface{}
}

//front和rear结点不保存具体值，只是用来指示真正头尾结点的位置。
