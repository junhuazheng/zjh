//耦合：
//耦合是指两个或两个以上的体系或运动形式间通过相互作用彼此影响以致联合起来的现象

//解耦：
//解耦就是用数学方法将两种运动分离开来处理问题。

//事件系统：
//事件系统可以将时间派发者与事件处理者解耦
//一个事件系统拥有如下特性：
//1、能够实现事件的一方，可以根据事件Id或者名字注册对应的事件。
//2、事件发起者，会根据注册信息通知这些注册者
//3、一个事件可有多个实现方响应

//事件注册：
//事件系统需要为外部提供一个注册入口。这个注册入口传入注册的事件名称和对应事件名称的响应函数。
//事件注册的过程就是将事件名称和响应函数关联并保存起来，，详细实现参考下面代码的RegisterEcent()函数
package main

var eventByName = make(map[string][]func(interface{}))

func REgisterEvent(name string, callback func(interface{})) {

	list := eventByName[name]

	list = append(list, callback)

	eventByName[name] = list
}

func CallEvent(name string, param interface{}) {

	list := eventByName[name]

	for _, callback := range list {

		callback(param)
	}
}
