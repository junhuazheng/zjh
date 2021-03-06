//Go语言goroutine（轻量级线程）

//线程
//线程（thread）是操作系统能够进行运算调度的最小单位。
//他被包含在进程之中，是进进程中的实际运作单位。
//一条线程指的是进程中一个单一顺序的控制流程，一个进程中可以并发多个线程，每条线程并行执行不同的任务
//在Unix System V 及SunOS中也被称为轻量进程，但轻量进程更多指内核线程，而吧用户线程称为线程

//线程是独立调度和分派的基本单位。线程可以为操作系统内核调度的内核线程，如Win32线程
//一个进程可以有很多线程，每条线程并行执行不同的任务

//在编写Socket网络程序时，需要提前准备一个线程池尾每个Socket的收发包分配一个线程。
//开发人员需要在线程数量和CPU数量间建立一个对应关系，以保证每个任务能及时地分配到CPU上进行处理，同事
//同时避免多个任务频繁地在线程间切换执行而损失效率

//虽然线程池为逻辑编写者提供了线程分配的抽象机制。
//但是，如果面对随时随地可能发生的并发和现场呢过处理任务，线程池就不是非常直观和方便了
//能否有一种机制：使用者分配足够多的任务，系统能自动使用者把任务分配到CPU上，让这些任务尽量并发运作。
//这种机制在Go语言中被称为goroutine

//goroutine的概念类似于线程，但goroutine由Go程序运行时的调度和管理
//Go程序会只能地将goroutine中的任务合理地分配给每个CPU

//Go程序从main包的main()函数开始，在程序启动时，Go程序就会为main()函数创建一个默认的goroutine

//使用普通函数创建goroutine
//Go程序中使用go关键字为一个函数创建一个goroutine。一个函数可以被创建多个goroutine，一个goroutine必定对应一个函数

//格式：
//为一个普通函数创建goroutine的写法如下：
go 函数名(参数列表)
//函数名：要调用的函数名
//参数列表：调用函数需要传入的参数

//使用go关键字创建goroutine时，被调用函数的返回值会被忽略。
//如果需要在goroutine中返回数据，请使用后面介绍的通道（channel）特性，通过通道把数据从goroutine中作为返回值传出