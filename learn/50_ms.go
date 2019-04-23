package main

// const c1 = 100

var b1 = 123

func main() {
	println(&b1, b1)
	println(&c1, c1)
}

/*输出结果为:cannot take the adress of c1
常量
常量不同于变量的运行期分配内存
常量通常会被编译器在预处理阶段直接展开
作为指令数据使用
*/
/*修改*/
var c1 = 100

/*输出结果为:
0x4b2138 123
0x4b2140 100
*/
