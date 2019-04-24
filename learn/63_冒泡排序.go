package main

import (
	"fmt"
	"sort"
	"time"
)

func bubbleSort(items []int) {
	var (
		n       = len(items)
		swapped = true
	)

	for swapped {
		swapped = false
		for i := 0; i < n-1; i++ {
			if items[i] > items[i+1] {
				items[i+1], items[i] = items[i], items[i+1]
				swapped = true
			}
		}
		n = n - 1
	}
}

//使用sort包
func bubbleSortUsingSortPackage(data sort.Interface) {
	r := data.Len() - 1
	for i := 0; i < r; i++ {
		for j := r; j > i; j-- {
			if data.Less(j, j-1) {
				data.Swap(j, j-1)
			}
		}
	}
}
