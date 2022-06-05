package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		if p, ok := hashTable[another]; ok {
			return []int{p, i}
		}
		hashTable[nums[i]] = i
	}
	return nil
}

func main() {
	nums := [...]int{2, 7, 11, 15}
	result := twoSum(nums[:], 9)
	fmt.Println(result)
}
