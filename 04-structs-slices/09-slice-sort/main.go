package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("sort slice")
	nums := []int{5, 4, 18, 47, 50, 9, 15}
	sort.Ints(nums)
	for _, n := range nums {
		fmt.Printf("%d \n", n)
	}
}
