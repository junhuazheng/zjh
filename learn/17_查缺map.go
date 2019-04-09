/*映射的关系容器为map，map使用散列表（hash）实现。
大多数语言中映射关系容器使用两种算法：散列表和平衡树。
散列表可以简单描述为一个数组（俗称”桶“），数组的每个元素是一个列表。
根据散列函数获得每个元素的特征值，将特征值作为映射的键。如果特征值重复，表示元素发生碰撞。
碰撞的元素将被放在同一个特征值的列表中进行保存。散列表查找复杂度为O(1)，和数组一致。
最坏的情况为O(n)，n为元素总数，散列需要尽量避免元素碰撞以提高查找效率，这样就需要对“桶”进行扩容
每次扩容，元素需要重新放入桶中，较为耗时。

平衡树类似于有斧子关系的一棵数据树，每个元素在放入树时，都要与一些节点进行比较。
map[KeyType]ValueType
KeyType为键类型
ValueTye是键对应的值得类型。
*/
package main

import (
	"fmt"
	"sort"
)

func main() {
	scenc := make(map[string]int)

	scenc["route"] = 66

	fmt.Println(scenc["route"])

	v := scenc["route2"]
	fmt.Println(v)

	//使用大括号进行内容定义，就像JSON格式一样，冒号的左边是key，右边是值，键值对之间使用逗号分隔。
	m := map[string]string{
		"A": "lef",
		"D": "rge",
		"S": "fasda",
	}

	fmt.Println(m["A"])

	//遍历map
	scene := make(map[string]int)

	scene["boy"] = 11
	scene["girl"] = 22
	scene["man"] = 33

	var sList []string

	for k := range scene {
		sList = append(sList, k)
	}

	sort.Strings(sList)

	fmt.Println(sList)

	//使用delete(map, 键)从map中删除一组键值对
	delete(scene, "boy")

	for k, v := range scene {
		fmt.Println(k, v)
	}
}
