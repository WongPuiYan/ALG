package main

import (
	"fmt"
)

func insertionSort(nums []int) []int {
	numsLength := len(nums)
	if numsLength <= 1 {
		return nums
	}
	for i := 0; i < numsLength; i++ {
		for j := 0; j < numsLength-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}

func main() {
	nums := []int{3, 2, 1}
	nums = insertionSort(nums)
	fmt.Println(nums)
}
