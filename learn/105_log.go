/*Golang's log模块主要提供了3类接口。
分别是“Print。Panic、Fatal”， 对每一类接口其提供了3种调用方式
分别是“Xxxx、Xxxln、Xxxxf”，基本和fmt中的相关函数类似，下面是一个Print的示例：*/
package main

import (
	"fmt"
	"log"
)

/*输出结果为：
2016/12/15 19:46:19 Print array [2 3]
2016/12/15 19:46:19 Println array [2 3]
2016/12/15 19:46:19 Printf array with item [2,3]
*/
/*
对于log.Fatal接口，会先将日志内容打印到标准输出，接着调用系统的os.exit(1)接口，退出程序并返回状态1
但是有一点需要注意，由于是直接调用系统接口退出，defer函数不会被调用，下面是一个Fatal示例：
*/

func test_deferfatal() {
	defer func() {
		fmt.Println("--first--")
	}()
	log.Fatalln("test for defer Fatal")
}

func main() {
	arr := []int{2, 3}
	log.Print("Print array", arr, "\n")
	log.Println("Println array", arr)
	log.Printf("Printf array with item [%d,%d]\n", arr[0], arr[1])

	test_deferfatal()

	test_deferpanic()
}

/*输出结果为：
2016/12/15 19:46:45 test for defer Fatal
可以看到并没有调用defer函数。
*/

//对于log.Panic接口，该函数把日志内容刷到标准错误后调用panic函数，下面是一个panic的示例：

func test_deferpanic() {
	defer func() {
		fmt.Println("--first--")
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	log.Panicln("test for defer Panic")
	defer func() {
		fmt.Println("--second--")
	}()
}

/*输出结果为：

2016/12/15 19:59:30 test for defer Panic
--first--
test for defer Panic

可以看到收线输出了“test for defer Panic”，然后第一个defer函数被调用并输出了“--first--"
但是第二个defer函数并没有输出，可见在Panic之后声明的defer是不会执行的。*/
