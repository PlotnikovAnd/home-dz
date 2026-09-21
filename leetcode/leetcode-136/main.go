package main

import "fmt"

func singleNumber(nums []int) int {
	res := 0
	for _, val := range nums {
		res = res ^ val
	}
	return res
}

func main() {
	fmt.Println(singleNumber([]int{1, 2, 3, 1, 2, 3, 5, 6, 6, 7, 7}))
}
