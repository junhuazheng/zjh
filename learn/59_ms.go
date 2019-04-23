package main

func test(x int) (func(), func()) {
	return func() {
			println(x)
			x += 110
		}, func() {
			println(x)
		}
}

func main() {
	a := test(100)
	a()
	b()
}

/*
输出结果为：
100
210
*/
