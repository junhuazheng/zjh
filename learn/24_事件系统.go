//事件注册的过程就是将事件名称和响应函数关联并保存起来
package main

import "fmt"

var evenByName = make(map[string][]func(interface{}))

//注册事件，提供事件名和回调函数
func RegisterEvent(name string, callback func(interface{})) {

	list := evenByName[name]

	list = append(list, callback)
	
	evenByName[name] = list
}

func CallEvent(name string, param interface{}) {
	
	list := evenByName[name]
	
	for _, callback := range list {
		callback(package)
	}
}

type Actor struct {	
}

func (a *Actor) OEvent(param interface{}) {
	fmt.Println("actor event:", param)
}

func GlobalEvent(param interface{}) {
	fmt.Println("global event:", param)
}

func main() {
	a := new(Actor)
	
	RegisterEvent("OnSkill", a.OnEvent)
	
	RegisterEvent("OnSkill", GlobalEvent)
	
	CallEvent("OnSkill", 100)
}
