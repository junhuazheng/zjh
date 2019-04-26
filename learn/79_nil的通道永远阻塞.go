/*nil的通道永远阻塞

当case上读一个通道时，如果这个通道是nil，则该case永远阻塞。
这个功能有一个妙用，select通常处理的是多个通道，当某个读通道关闭了
但不想select再继续关注此case，二十关注其他case，把该通道设置为nil即可。

下面是一个合并程序等待两个输入通道都关闭后才退出的例子，就是使用了这个特性*/

package main

func combine(inCh1, inCh2 <-chan int) <-chan int {

	//输出通道
	out := make(chan int)

	//启动协程合并数据
	go func() {
		defer close(out)
		for {
			select {
			case x, open := <-inCh1:
				if !open {
					inCh1 = nil
					continue
				}
				out <- x
			case x, open := <-inCh2:
				if !open {
					inCh2 = nil
					continue
				}
				out <- x
			}

			//ch1和ch2都关闭才是退出
			if inCh1 == nil && inCh2 == nil {
				break
			}
		}
	}()

	return out
}
