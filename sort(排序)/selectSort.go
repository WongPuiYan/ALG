package main

import (
	"fmt"
)

func selectionSort(nums []int) []int {
	numsLength := len(nums)
	if numsLength <= 1 {
		return nums
	}
	for i := 0; i < numsLength-1; i++ {
		minIndex := i
		for j := i + 1; j < numsLength; j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
		nums[i], nums[minIndex] = nums[minIndex], nums[i]
	}
	return nums
}

func main() {
	nums := []int{3, 2, 1}
	nums = selectionSort(nums)
	fmt.Println(nums)
}
