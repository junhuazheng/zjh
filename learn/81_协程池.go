package main

import (
	"fmt"
	"time"
)

/*协程池
使用协程池是为了更好并发，程序鲁棒性、容错性等。
协程池指的是预先分配固定数量的goroutine处理相同的任务。

最简单的协程池模型 ：

        goroutine0
        goroutine1
 jobCh{ goroutine2 }retCh
		goroutine3
		goroutine4

上面这个就是最简单的协程池的样子。
先把协程池作为一个整体看，他使用两个通道

左边的jobCh是任务通道，任务会从这个通道中流进来
右边的retCh是结果通道，协程池处理任务后得到的结果会写入这个通道。

看一下协程池的内部，图中有5个goroutine，实际goroutine的数量是一具体情况而定的
协程池内每个协程都从jobCh读任务、处理任务，然后将结果写入到retCh。*/

//例子1
func workerpool(n int, jobCh <-chan int, retCh chan<- string) {
	for i := 0; i < n; i++ {
		go worker(i, jobCh, retCh)
	}
}

func worker(id int, jobCh <-chan int, retCh chan<- string) {
	for job := range jobCh {
		ret := fmt.Sprintf("worker %d processed job: %d", id, job)
		retCh <- ret
	}
}

/*
workerPool()会创建1个简单的协程池，协程的数量可以通过入参n执行，并且还制定了jobCh和retCh两个参数。

wroker()是协程池中的协程，入参分布是他的id、job通道和结果通道。
使用for-range从jobch读取任务，直到jobch关闭，然后一个最简单的任务：生成1个字符串，证明自己处理了某个人物，并把字符串作为结果写入retCh。
*/

func genJob(n int) <-chan int {
	jobCh := make(chan int, 200)
	go func() {
		for i := 0; i < n; i++ {
			jobCh <- i
		}
		close(jobCh)
	}()
	return jobCh
}
func main() {
	jobCh := genJob(10)
	retCh := make(chan string, 200)
	workerpool(5, jobCh, retCh)

	time.Sleep(time.Second)
	close(retCh)
	for ret := range retCh {
		fmt.Println(ret)
	}
}
