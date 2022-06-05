package main

import (
	"fmt"
)

func removeElement(nums []int, val int) int {
	var (
		left  = 0
		right = len(nums) - 1
	)
	if left == right && nums[left] != val {
		return 1
	}
	for ; left <= right; left++ {
		if nums[left] == val {
			for ; right > left && nums[right] == val; right-- {
			}
			if left == right {
				break
			}
			nums[left], nums[right] = nums[right], nums[left]
		}
	}
	return left
}

func main() {
	nums := [...]int{3}
	// nums := [...]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	target := 3
	result := removeElement(nums[:], target)
	fmt.Println(result)
}
