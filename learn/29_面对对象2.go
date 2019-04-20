/*
继承
其他面对对象的语言，是使用extends关键字来表示继承的
而go语言中继承的设计非常简单，不需要额外的关键字
*/

package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

type Stuent struct {
	Person
	id    int
	score int
}

func (student *Stuent) showInfo() {
	fmt.Println("I am a student ...")
}

func (student *Stuent) read() {
	fmt.Println("read book")
}

func (perosn *Person) showInfo() {
	fmt.Printf("My name is %s, age is %d\n", perosn.name, perosn.age)
}

func (perosn *Person) selAge(age int) {
	perosn.age = age
}

func main() {

	student := Stuent{Person{"jake", 16}, 1001, 99}

	student.showInfo()
	student.selAge(22)
	fmt.Println(student)
	student.read()
}

/*
可以看到，在学生结构体中，我们使用了Person作为匿名变量
意思就是，Student继承自Person。我们除了继承了成员变量和成员方法之外，
还可以为学生结构体增加成员方法，重写成员方法。*/
