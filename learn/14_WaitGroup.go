/*
WaitGroup用于等待一组goroutine结束，用法比较简单。它有三个方法：
func (wg *WaitGroup) Add(delta int)
func (wg *WaitGroup) Done()
func (wg *WaitGroup) Wait()

Add用来添加groutine的个数。
Done执行一次数量减一。
Wait用来等待结束。
*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i, s := range seconds {

		//计数加1
		wg.Add(1)
		go func(i, s int) {

			//计数减一
			defer wg.Done()
			fmt.Printf("goroutine %d 结束\n", i)
		}(i, s)
	}

	//等待执行结束
	wg.Wait()
	fmt.Println("所有goroutine执行结束")
}

//注意，wg.Add()方法一定要在goroutine开始前执行。
