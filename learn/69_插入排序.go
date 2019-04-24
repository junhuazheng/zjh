/*
插入排序（Insertion Sort）算法描述是一种简单直观的排序算法。
它的工作原理是通过构建有序序列，对未排序数据，在已排序序列中从后向前扫面，找到相应的位置并插入。
插入排序在实现上，通常采用in-place排序（即只需要用到O(1)的额外空间的排序）
因而在从后向前扫描过程中，需要反复把已排序元素逐步向后挪位，为最新元素提供插入空间。
*/

package main

import "sort"

func insertionSort(items []int) {
	var n = len(items)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if items[j-1] > items[j] {
				items[j-1], items[j] = items[j], items[j-1]
			}
			j = j - 1
		}
	}
}

//使用sort包

func insertionSortUsingSortPackage(data sort.Interface) {
	r : =data.Len() - 1
	for i := 1; i <= r; i++ {
		for j := i; j > 0 && data.Less(j, j-1); j-- {
			data.Swap(j, j-1)
		}
	}
}
