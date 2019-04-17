package main

import "fmt"

type People struct {
}

type Teacher struct {
	People
}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showb")
}

func (p *People) ShowB() {
	fmt.Println("showB")
}

func main() {
	t := Teacher{}
	t.ShowA() //等价于 t.People.ShowA()
	// fmt.Println("----------")
	// t.People.ShowA()

}

/*
Go中没有继承，上面这种写法叫组合

上面的t.ShowA()等价于t.People.ShowA()
*/
