package main

import (
	"fmt"
	"strconv"
	"strings"
)

//定义一个处理错误的结构体
type ParseError struct {
	key   int
	Value string
	Err   error
}

//Sprintf是字符串格式化命令，主要功能是把格式化的数据写入某个字符串中
func (p *ParseError) String() string {
	return fmt.Sprintf("[%q] type is not int", p.Value)
}

func JudgmentType(fields []string) (numbers []int) { //这个函数是判断fields切片中的每个元素是否都可以转换成INT类型
	if len(fields) == 0 {
		panic("Nothing can be explained")
	}

	for key, value := range fields {
		num, err := strconv.Atoi(value) //这里是将每个切片元素中的字符串进行转换
		if err != nil {
			panic(&ParseError{key, value, err}) //如果解析出错就将自定义的ParseError结构体的error对象返回。
		}
		numbers = append(numbers, num) //如果转换成新类型没有出错的话就会被追加到一个提前定义好的切片中
	}
	return
}

func StringParse(Input string) (numbers []int, err error) { //解析字符串的函数

	defer func() { //recover函数来接受panic抛出的错误信息。
		if ErrorOutput := recover(); ErrorOutput != nil {
			var ok bool
			err, ok = ErrorOutput.(error) //这里是一种断言操作，即判断是否有error类型出现
			if !ok {
				err = fmt.Errorf("Parse error: %v", ErrorOutput)
			}
		}
	}()
	fields := strings.Fields(Input)
	numbers = JudgmentType(fields)
	return
}

func main() {
	var zjh = []string {
		"100 200 300",
		"1 2 2.5 3",
		"30 * 40", 
		"zhengjunhua Golang",
		"",
	}
	
	for _. ex := range zjh {
		fmt.Printf("正在解析[%q]:\n", ex)
		result, err := StringParse(ex)
		if err != nil {
			fmt.Println("解析结果:", err)
			continue
		}
		fmt.Println("解析结果:", result)
	}
}
