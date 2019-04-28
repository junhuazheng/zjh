/*
Go语言的map在并发情况下，只读是线程安全的，同时读写线程不安全。

并发的map读写，也就是说使用了两个并发函数不断地对map进行读和写而发生了竞态问题。
map内部会对这种并发操作进行检查并提前发现。

Go语言提供了一中效率较高的并发安全sync.Map，sync.Map是在sync包下的特殊结构。

sync.Map特性：
1、无须初始化，直接声明即可。
2、sync.Map不能使用map的方式进行取值和设置等操作，二十使用sync.Map的方法进行调用。Store表示存储，Load表示获取，Delete表示删除。
3、使用Range配合一个回调函数进行遍历操作，通过回调函数返回内部遍历出来的值。Range参数中的回调函数的返回值功能是：需要继续迭代遍历是，返回true；终止迭代遍历是，返回false。
*/

package main

import (
	"fmt"
	"sync"
)

func main() {

	var scene sync.Map

	//存储
	scene.Store("greece", 97)
	scene.Store("loddon", 100)
	scene.Store("egy", 220)

	//获取
	fmt.Println(scene.Load("loddon"))

	//删除
	scene.Delete("loddon")

	//遍历
	scene.Range(func(k, v interface{}) bool {

		fmt.Println("iterate:", k, v)
		return true
	})

}

/*
输出结果为：
100 true
iterate: greece 97
iterate: egy 220

sync.Map没有提供获取map数量的方法，替代方法是获取时遍历自行计算。
sync.Map为了保证并发安全有一些性能损失，因此在非并发的情况下，使用map相比使用sync.Map会有更好的性能。
*/
