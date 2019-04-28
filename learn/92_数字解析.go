/*从字符串中解析数字在很多程序中是一个基础常见的任务，在Go中是这样处理的

内置的strconv包提供了数字解析功能*/
package main

import (
	"fmt"
	"strconv"
)

func main() {

	f, _ := strconv.ParseFloat("1.234", 64) //使用ParseFloat解析浮点数，这里的64表示解析的数的位数。
	fmt.Println(f)

	i, _ := strconv.ParseInt("123", 0, 64) //使用ParseInt解析整形数时，例子中的参数0表示自动推断字符串索表示的数字的进制。64表示返回的整形数是以64位存储的。
	fmt.Println(i)

	d, _ := strconv.ParseInt("0x1c8", 0, 64) //ParseInt会自动识别出十六进制数。
	fmt.Println(d)

	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	k, _ := strconv.Atoi("135") //Atoi是一个基础的10进制整形数转换函数。
	fmt.Println(k)

	_, e := strconv.Atoi("wat") //在输入错误时，解析函数会返回一个错误。
	fmt.Println(e)
}
