//GO语言结构体的方法
//将背包及放入背包的物品中使用Go语言的结构体和方法方式编写：为*Bag创建一个方法，代码如下：

type Bag struct {
	items []int
}

func (b *Bag) Insert(itemid int) { //Insert（itemid int）的写法与函数一致。（b *Bag）表示接收器，即Insert作用的对象实例
	b.items = append(b.items, itemid)
}

func amin() {

	b := new(Bag)

	b.Insert(1001) //在Insert()转换为方法后，我们就可以愉快地像其他语言一样，用面向对象的方法来调用b的Insert
}

//每个方法只能有一个接收器
