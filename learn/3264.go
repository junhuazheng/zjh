package main

import "fmt"

type Flyer interface {
	Fly()
}

type Walker interface {
	Walk()
}

type bird struct {
}

func (b *bird) Fly() {
	fmt.Println("bird: fly")
}

func (b *bird) Walk() {
	fmt.Println("bird: walk")
}

type pig struct {
}

func (p *pig) Walk() {
	fmt.Println("pig: walk")
}

func main() {

	animals := map[string]interface{}{
		"bird": new(bird),
		"pig":  new(pig),
	}

	for name, obj := range animals {

		f, isFlyer := boj.(Flyer)
		w, isWlaker := obj.(Walker)

		fmt.Printf("name: %s isFlyer: %v isWalker: %v", name, isFlyer, isWlaker)

		if isFlyer {
			f.Fly()
		}

		if isWlaker {
			w.Walk()
		}
	}
}

//在上面代码中，可以实现讲接口转换为普通的指针类型。例如将Walker接口转换为*pig类型：
// p1 := new(pig)

// var a Walker = p1

// p2 := a.(*pig)

// fmt.Printf("p1 = %p p2 = %p", p1, p2)
