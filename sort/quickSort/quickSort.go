package main

import (
	// "crypto/rand"
	"fmt"
	// "math/big"
	"math/rand"
	"time"
)

func partition(nums []int, left int, right int) int {
	rand.Seed(time.Now().UnixNano())
	pos := rand.Intn(right-left) + left
	tar := nums[pos]
	nums[left], nums[pos] = nums[pos], nums[left]

	for left < right {
		for right > left && nums[right] >= tar {
			right -= 1
		}
		nums[left] = nums[right]
		for left < right && nums[left] <= tar {
			left += 1
		}
		nums[right] = nums[left]
	}
	nums[left] = tar

	return left
}

func qsort(nums []int, left int, right int) []int {
	if left >= right {
		return nums
	}

	p_index := partition(nums, left, right)
	if p_index > left+1 {
		for p_index > left && nums[p_index] == nums[p_index-1] {
			p_index -= 1
		}
		qsort(nums, left, p_index-1)
	}
	if p_index < right-1 {
		for p_index < right && nums[p_index] == nums[p_index+1] {
			p_index += 1
		}
		qsort(nums, p_index+1, right)
	}
	return nums
}

func quickSort(nums []int, left int, right int) []int {
	return qsort(nums, 0, len(nums)-1)
}

func main() {
	// nums := []int{3, 2, 1, 8, 2, 1, 6, 7, 9}
	nums := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}

	bubbleResult := quickSort(nums, 0, len(nums)-1)
	fmt.Println(bubbleResult)
}
