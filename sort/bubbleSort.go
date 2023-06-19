package main

import (
	"fmt"
)

// 冒泡排序
func bubbleSort(nums []int) []int {
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

// 选择排序
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

// 插入排序
func insertionSort(nums []int) []int {
	numsLength := len(nums)
	if numsLength <= 1 {
		return nums
	}
	for item := range nums {
		for preIndex := item - 1; preIndex >= 0 && nums[preIndex] > nums[item]; preIndex-- {
			nums[preIndex+1] = nums[preIndex]
		}
		// nums[preIndex+1] = nums[item]
	}
	return nums
}

func main() {
	nums := []int{3, 2, 1, 8, 2, 1, 6, 7, 9}

	bubbleResult := bubbleSort(nums)
	fmt.Println(bubbleResult)

	selectionResult := selectionSort(nums)
	fmt.Println(selectionResult)

	insertionResult := insertionSort(nums)
	fmt.Println(insertionResult)
}
