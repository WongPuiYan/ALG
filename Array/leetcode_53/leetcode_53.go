package main

import (
	"fmt"
)

func maxSubArray(nums []int) int {
	max_sum := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1] // num[i]记录的是以num[i]结尾的最大连续子数组和
		}
		if nums[i] > max_sum {
			max_sum = nums[i]
		}
	}
	return max_sum
}

type Status struct {
	lSum, rSum, mSum, iSum int
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func pushUp(l, r Status) Status {
	iSum := l.iSum + r.iSum
	lSum := max(l.lSum, l.iSum+r.lSum)
	rSum := max(r.rSum, r.iSum+l.rSum)
	mSum := max(max(l.mSum, r.mSum), l.rSum+r.lSum)
	return Status{lSum, rSum, mSum, iSum}
}
func get(nums []int, l, r int) Status {
	if l == r {
		return Status{nums[l], nums[l], nums[l], nums[l]}
	}
	m := (l + r) >> 1
	lSub := get(nums, l, m)
	rSub := get(nums, m+1, r)
	return pushUp(lSub, rSub)
}
func maxSubArray2(nums []int) int {
	return get(nums, 0, len(nums)-1).mSum
}

func main() {
	nums := [...]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	result := maxSubArray(nums[:])
	fmt.Println(result)
	nums2 := [...]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	result2 := maxSubArray2(nums2[:])
	fmt.Println(result2)
}
