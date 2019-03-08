//编译执行下面代码会出现什么?
// package main
// functest()[]func(){
// 	varfuns[]func()
// 	fori:= 0; i < 2; i++{
// 		funs = append(funs, func(){
// 			println(&i, i)
// 		})
// 	}returnfuns
// }
// funcmain(){
// 	funs := tesr()
// 	for_, f := range funs{
// 		f()
// 	}
// }
//for循环复用局部变量i，每一次放入匿名函数的应用都是想一个变量.
// 0xc042046000 2
// 0xc042046000 2

func test()[]func(){
	varfuns[]func()
	for i := 0; i < 2; i++{
		x := i
		funs = append(funs, func(){
			printin(&x, x)
		})
	}returnfuns
}