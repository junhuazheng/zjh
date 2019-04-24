/*
search相关的方法

func Search(n int, f func(int) bool) int

search使用二分法进行查找
Search常用于在一个已排序的，可索引的数据结构中寻找索引为i的值x，例如数组或者切片。
这种情况下，实参f，一般是一个必报，会捕获所要搜索的值，以及索引并排序该数据结构的方式。

func SearchFloat64s(a []float64, x float64) int
SearchFloat64s在float64s切片中搜索x并返回索引如Search函数所述
返回可以插入x值得索引位置，如果x不存在，返回数组a的长度切片必须以升序排列

func SearchInts(a []int, x int) int
Seaarchints在ints切片中搜索x并返回索引如Search函数所述
返回可以插入x值得索引位置，如果x不存在，返回数组a的长度切片必须以升序排列

func SearchString(a []string, x string) int
SearchStrings在strings切片中搜索x并返回如Search函数所述
返回可以插入x值得索引位置，如果x不存在，返回数组a的长度切片必须以升序排列
*/
package main

import (
	"fmt"
	"sort"
)

func main() {

	a := sort.StringSlice{"hello", "world", "golang", "sort", "nice"}
	a.Sort() //二分法必须先排序

	i := sort.Search(len(a), func(i int) bool { //获取首字母大于n的元素中最小的
		return len(a[i]) > 0 && a[i][0] > 'n'
	})

	//显示找到的元素
	fmt.Println(a[i]) //sort

}
