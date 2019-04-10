/*sync.Map有以下特性：
1、无需初始化，直接声明即可
2、sync.Map不能使用map的方式进行取值和设置等操作，而是使用sync.Map的发布方法进行调用。
Stroe表示存储，Load表示获取，Delete表示删除。
3、使用Range配合一个回调函数进行遍历操作，通过回调函数返回内部遍历出来的值。
Range参数中的回调函数的返回值功能是：需要继续迭代遍历时，返回true；终止迭代遍历时，返回false
并发安全sync.Map演示代码：*/
package main

import (
	"fmt"
	"sync"
)

func main() {

	var scene sync.Map

	scene.Store("gre", 11)
	scene.Store("lon", 22)
	scene.Store("qwe", 33)

	fmt.Println(scene.Load("qwe"))

	scene.Delete("qwe")

	scene.Range(func(k, v interface{}) bool {
		fmt.Println("iterate:", k, v)
		return true
	})
}

//sync.Map没有提供获取map数量的方法，替代方法是获取时遍历自行计算数量。
//sync.Map为了保证并发安全有一些性能损失，因此在非并发情况下，使用map相比使用sync.Map会有更好的性能。
