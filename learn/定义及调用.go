package main

import (
	"fmt"
	"strings"
)

type Human struct {
	Name   string
	string //内置的匿名字段，我们用他来定义家庭住址
	phone  int
}

type Student struct {
	Human     //匿名字段，其类型就是上面的自定义类型
	Classroom string
	Score     float64 //考试成绩
}

type Teacher struct {
	Human
	Position string  //职位信息
	salary   float64 //薪资情况
}

//结构体Human的自我介绍方法
func (h Human) SayHi() {
	fmt.Printf("Hello, my name is[%s] . My phone number is [%d] . My home address is [%s]\n",
		h.Name, h.phone, h.string)
}

func (h Human) Sing(Name string) {
	fmt.Printf("[%s] : 在那山的那边海的那边有一群小精灵......\n", Name)
}

func (t Teacher) SayHi() {
	fmt.Printf("Anywaym, I'm your [%s]teacher, you can call me [%s], my salary is...[%f]\n",
		t.Position, t.Human.Name, t.salary)
}

type Superman interface {
	SayHi()
	Sing(Name string) //该接口也包含“Sing(Name string)”方法
}

func main() {

	zjh := Student{Human{"郑俊华", "广州", 69696991}, "三年级一班", 99} //将"zjh"实例
	mht := Student{Human{"麻花藤", "深圳", 23232331}, "一年级一班", 59}
	cat := Teacher{Human{"猫", "广州", 34343441}, "Golang", 20000}

	var zhengjunhua Superman //声明一个变量，其类型为我们定义的接口。

	zhengjunhua = zjh
	zhengjunhua.SayHi()
	zhengjunhua.Sing("高音唱")

	fmt.Println("\n", strings.Repeat("*", 30), "我是分割线", strings.Repeat("*", 30), "\n")

	zheng := make([]Superman, 3)
	zheng[0], zheng[1], zheng[2] = mht, cat, zjh

	for _, value := range zheng { //同时调用3个方法
		value.SayHi()
	}
}

/*
通过上面的代码我们可以知道，interface可以被任意的对象实现。同理，一个对象可以实现任意多个interface。
你会发现interface就是一组抽象方法的集合，他必须由其他非interface类型实现，而不能自我实现，go通过
interface实现了duck-typing：即“当看到一只鸟走起来像鸭子、游泳像鸭子、叫起来也像鸭子，那么这只鸟
就可以被称为鸭子”。
*/
