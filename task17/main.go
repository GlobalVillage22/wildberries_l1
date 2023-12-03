package main

import (
	"fmt"
	"sort"
)

/*
Реализовать бинарный поиск встроенными методами языка.
*/

func binarySearch(arr []int, x int) int {

	return -1
}

func main() {
	nums := []int{5, 2, 1, 6, 3, 24, -91}

	target := 24

	index := sort.Search(len(nums), func(index int) bool {
		return nums[index] >= target
	})

	if index < len(nums) && nums[index] == target {
		fmt.Println(nums[index])
	} else {
		fmt.Println("элемент не найден")
	}
}
