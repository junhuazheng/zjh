//var team = [...]string{"ham", "sold", "mum"}
package main

import "fmt"

func main() {

	var team [3]string
	team[0] = "ham"
	team[1] = "sold"
	team[2] = "mun"

	//使用for循环，遍历team数组，遍历出的键k为数组的索引
	//值v为数组的每个元素值
	for k, v := range team {
		fmt.Println(k, v)
	}

	var n [10]int
	var i, j int

	for i = 0; i < 10; i++ {
		n[i] = i + 100
	}

	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}
}

/*数组是具有相同唯一类型的一组已编号且长度固定的数据项序列，这种类型可以使任意的原始类型或者自定义类型
数组元素可以通过索引（位置）来读取（或者修改）*/
