package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	var (
		n    = len(nums)
		left = 0
	)
	if n == 0 {
		return left
	}
	left++
	for right := 1; right < n; right++ {
		if nums[right] != nums[right-1] {
			nums[left] = nums[right]
			left++
		}
	}
	return left
}

func main() {
	nums := [...]int{1, 1, 2, 2}
	// nums := [...]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	result := removeDuplicates(nums[:])
	fmt.Println(result)
}
