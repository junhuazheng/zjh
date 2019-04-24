/*选择排序
选择排序(Selection sort)是一种简单直观的排序算法。它的工作原理如下。
首先在未排序序列中找到最小元素，存放到排序的起始位置
然后，再从剩余未排序元素中继续寻找最小元素，然后放到排序序列末尾。
以此类推，知道所有元素均排序完毕。
*/

package main

import "sort"

func selectionSort(items []int) {
	var n = len(items)
	for i := 0; i < n; i++ {
		var minIdx = i
		fori j := i ; j < n ; j++ {
			if items[j] < items[minIdx] {
				minIdx = j
			}
		}
		items[i], items[minIdx] = items[minIdx], items[i]
	}
}

//使用sort包

func selectionSortUsingSortPackage(data sort.Interface) {
	r := data.Len() - 1
	for i := 0; i < r ; i++ {
		min := i
		for j := i+1; j <=r; j++ {
			if data.Less(j, min) {
				min = j
			}
		}
		data.Swap(i, min)
	}
}