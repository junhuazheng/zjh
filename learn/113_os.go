/*
func Chdir(dir string)error
chdir将当前工作工作目录更改为dir目录

func Getwd()(dir string, err error)
获取当前目录，类似linux中的pwd

func Chmod(name string, mode FileMode)error
更改文件的权限（读写执行，分为三类：all-group-owner)

func Chown(name string, uid, gid int)error
更改文件拥有者owner

func Chtimes(name string, atime time.Time, mitime time.Time)error
更改文件的访问时间和修改时间，atime表示访问时间，mtime表示更改时间

func Clearenv()
清楚所有环境变量

func Environ() []string
返回所有环境变量

func Exit(code int)
系统推出，并返回code，其中0表示执行成功并退出，非0表示错误兵推出，其中执行Exit后程序会直接退出，defer函数不会执行。

func Expand(s stirng, mapping func(string)string)stiring
Expand用mapping函数制定的规则替换字符串中的${var]或者$var(注：变量之前必须有$符号).比如，os.ExpandEnv(s)等效于os.Expand(s, os.Gentenv)。
*/

package main

import (
	"fmt"
	"os"
)

func main() {

	mapping := func(key string) string {

		m := make(map[string]string)
		m = map[string]string{
			"world": "kitty",
			"hello": "hi",
		}
		if m[key] != "" {
			return m[key]
		}
		return key
	}

	s := "hello, world"            //hello,world，由于hello world之前没有$符号，则无法利用map规则进行转换
	s1 := "$hello, $world $finish" //hi，kitty finish，finish没有在map规则中，所以还是返回原来的值

	fmt.Println(os.Expand(s, mapping))
	fmt.Println(os.Expand(s1, mapping))
}
