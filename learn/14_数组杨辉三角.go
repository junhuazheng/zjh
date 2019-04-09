package main

import "fmt"

func main() {
	nums := []int{}
	var i int
	for i = 0; i < 10; i++ {
		nums = Yang(nums)
		fmt.Println(nums)
	}
}

func Yang(inArr []int) []int {
	var out []int
	var i int
	arrlen := len(inArr)
	out = append(out, 1)
	if 0 == arrlen {
		return out
	}
	for i = 0; i < arrlen-1; i++ {
		out = append(out, inArr[i]+inArr[i+1])
	}
	out = append(out, 1)
	return out
}
