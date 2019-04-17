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
	t.ShowB()

}

/*因为Teacher没有方法ShowA
所以调用的是People的ShowA
又People没有t.Show()方法
故t.ShowA 等价于 t.People.ShwoA*/
