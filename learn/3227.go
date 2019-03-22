//time包中的类型方法

//Go语言提供的time包主要用于时间的获取和计算等。在这个包中，也使用了类型方法：
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Second.String())
}

//time.Second是一个常量，下面代码就是time.Second的定义：
const (
	Nanosecond  Duration = 1
	Microsecond          = 1000 * Nanosecond
	Millisecond          = 1000 * Microsecond
	Second               = 1000 * Millisecond
	Minute               = 60 * Second
	Hour                 = 60 * Minute
)

//Second的类型为Duration，而Duration实际是一个int64类型，定义如下：
type Duration int64
//它拥有一个String的方法，部分定义如下：
func (d Duration) string() string {
	
	//一系列生成buf的代码
	...
	
	return string(buf[w:])
}
//Duration.String可以将Duration的值转换为字符串。
