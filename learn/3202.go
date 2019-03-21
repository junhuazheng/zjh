package main

//匿名安徽省农户作为回调函数的设计在Go语言的系统包中也不叫常见，strings包中就有如下代码：
func TrimFunc(s string, f func(rune) bool) string {
	return TrimRightFunc(TrimLeftFunc(s, f), f)
}

//上个例子
func main() {

	visit([]int{1, 2, 3, 4}, func(v int) {
		fmt.Println(v)
	})
}

//定义了一个匿名函数，作用是将遍历的每个值打印出来
