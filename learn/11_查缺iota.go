//C#中Week枚举值Mondy为1.那么可以通过Week.Monday.ToString()函数获得Monday字符串
//Go语言也可以实现这一功能

package main

import "fmt"

type ChTy int

const (
	No ChTy = iota
	Cpu
	Gpu
)

func (c ChTy) String() string {
	switch c {
	case No:
		return "No"
	case Cpu:
		return "Cpu"
	case Gpu:
		return "Gpu"
	}

	return "A"
}

func main() {
	fmt.Printf("%s %d", Cpu, Cpu)
}

//使用String()方法的ChTy在使用上和普通的常量没有区别。当这个类型需要显示为字符串时，Go语言会自动寻找String()方法进行调用。
