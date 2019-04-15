/*假设要创建一个文件，写入内容，然后在完成之后关闭。这里可以这样使用延迟（defer）处理。
在使用createFile获取文件对象后，立即使用closeFile推迟该文件的关闭。这将在writeFile()完成后封装函数（main）结束时执行。*/

package main

import (
	"fmt"
	"os"
)

func main() {

	f := createFile("defer-test.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	f.Close()
}
