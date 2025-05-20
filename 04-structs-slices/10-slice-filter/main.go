package main

import (
	"fmt"
)

func filters_ints(slice []int, pred func(int) bool) []int {

	result := []int{}
	for _, v := range slice {
		if pred(v) {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	fmt.Println("")
	nums := []int{2, 9, 8, 15, 4, 75, 90, 1, 6}
	start := 3
	even := filters_ints(nums, func(c int) bool {
		return c > start
	})

	fmt.Println(even)
}
