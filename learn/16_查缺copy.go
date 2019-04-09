//copy(desrSlice, srcSlice []T)int
//SrcSlice为数据来源切片
//destSlice为复制的目标。目标切片必须分配过空间且足够承载复制的元素个数
//来源和目标的类型一直，copy的返回值表示实际发生复制的元素个数。

package main

import "fmt"

func main() {

	const eleC = 1000

	srcData := make([]int, eleC)

	for i := 0; i < eleC; i++ {
		srcData[i] = i
	}

	refData := srcData

	copyData := make([]int, eleC)

	copy(copyData, srcData)

	srcData[0] = 999

	fmt.Println(refData[0])

	fmt.Println(copyData[0], copyData[eleC-1]) //复制数据的首位数据，由于数据是复制的，因此不会法神变化。

	copy(copyData, srcData[4:6]) //强srcData的局部数据复制到copyData中

	for i := 0; i < 5; i++ {
		fmt.Printf("%d\n", copyData[i])
	}

	//删除切片元素
	seq := []string{"a", "b", "c", "d", "e"}

	boy := 2

	fmt.Println(seq[:boy], seq[boy:])

	seq = append(seq[:boy], seq[boy+1]...)
	fmt.Println(seq)
}
