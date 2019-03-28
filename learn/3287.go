//Go语言GOMAXPROCS（调整并发的运行性能）

//在Go程序运行时（runtime）实现了一个小型的任务调度器。
//这套调度器的工作原理类似于操作系统调度线程，Go程序调度器可以高效将CPU资源分配给每一个任务。
//传统逻辑中，开发者需要维护线程池中线程与CPU核心数量的对应关系，同样的，Go地中也可以通过
//runtime.GOMAXPROCS()函数做到：
//格式为：
runtime.GOMAXPROCS(逻辑CPU数量)

//这里的逻辑CPU数量可以有如下几种数值：
// <1 : 不修改任何数值
// =1 ： 单核心执行
// >1 : 多核并发执行

//一般情况下，可以使用runtime.NumCPU()查询CPU数量，并使用runtime.GOMAXPROCS()函数进行设置：
runtime.GOMAPROCS(runtime.NumCPU())

//GOMAXPROCS同时也是一个环境变量，在应用程序启动前设置环境变量也可以起到相同的作用

//并发和并行的区别
//并发（concurrency）：把任务在不同的时间点交给处理器进行处理。在同一时间点，任务并不会同时运行
//并行（parallelism）：把每一个任务分配给每一个处理器独立完成。在同一时间点，任务一定是同时运行。

//Go语言在GOMAXPROCS数量与任务数量相等时，可以做到并行执行，但一般情况下都是并发执行。
