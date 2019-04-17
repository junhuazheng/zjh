package main

import (
	"fmt"
	"runtime"
	"sync"
)

func init() {
	fmt.Println("Current Go Version:", runtime.Version())
}

func main() {

	runtime.GOMAXPROCS(1)

	count := 10
	wg := sync.WaitGroup{}
	wg.Add(count * 2)

	for i := 0; i < count; i++ {
		go func() {
			fmt.Printf("[%d]", i)
			wg.Done()
		}()
	}

	for i := 0; i < count; i++ {
		go func(i int) {
			fmt.Printf("-%d-", i)
			wg.Done()
		}(i)
	}

	wg.Wait()
}

/*两个for循环内部 go func 调用参数i的方式是不同的，导致结果完全不同。

第一个 go func 中 i 是外部for的一个变量，地址不变。
遍历完成后，最终i=10.故go func 执行时，i的值始终是10（10次遍历很快完成）

第二个 go func 中 i 是函数参数，与外部for中的i完全是两个变量。
尾部（i）将发生值拷贝，go func 内部指向值拷贝地址。*/
