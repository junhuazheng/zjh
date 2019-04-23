package main

// func test() []func() {
// 	var funs []func()
// 	for i := 0; i < 2; i++ {
// 		funs = append(funs, func() {
// 			println(&i, i)
// 		})
// 	}
// 	return funs
// }

func main() {
	funs := test()
	for _, f := range funs {
		f()
	}
}

/*
闭包延迟求值
输出结果为:
0xc00003e00 2
0xc00003e00 2
for循环复用局部变量i，每一次放入匿名函数的都是同一个变量。
*/
//如果想不一样，可以修改为：
func test() []func() {
	var funs []func()
	for i := 0; i < 2; i++ {
		x := i
		funs = append(funs, func() {
			println(&x, x)
		})
	}
	return funs
}

/*
输出结果为：
0xc00003e000 0
0xc00003e008 1
*/
