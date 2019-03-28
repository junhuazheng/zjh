//goroutine和coroutine的区别

//coroutine与goroutine在名字上类似，都可以将函数或者语句在独立的环境中运行，单是他们之剑有两点不同：
//1、goroutine可能发生并行执行
//2、但coroutine始终顺序执行

//狭义地说，goroutine可能法神刚在多线程环境下，goroutine无法控制自己获得高有限度支持；
//coroutine始终发生在单线程，coroutine程序需要主动交出控制权，宿主才能获得控制权并将控制权交给其他coroutine

//goroutine间使用channel通信，coroutine使用yield和resume操作。

//goroutine和coroutine的概念和运行机制都是脱胎于早期的操作系统

//coroutine的运行机制属于协作式任务处理，早期的操作系统要求每一个应用必须遵守操作系统的任务处理规则
//应用程序在不需要使用CPU时，会主动交出CPU使用权。如果开发者无意间或者故意让应用程序长时间占用CPU
//操作系统也无能为力，表现出来的效果就是计算机很容易失去响应或者司机

//goroutine属于抢占式任务处理，已经和现有的多线程和多进程任务处理非常类似。
//应用程序对CPU的控制最终还需要操作系统来管理，操作系统如果发现一个应用程序长时间大量的占用CPU，那么用户有权终止这个任务
