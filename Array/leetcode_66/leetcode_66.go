package main

import (
	"fmt"
)

func plusOne(digits []int) []int {
	var (
		n   = len(digits)
		tmp = 1
	)

	for i := n - 1; i >= 0; i-- {
		tmp = tmp + digits[i]
		digits[i] = tmp % 10
		tmp = tmp / 10
		if tmp == 0 {
			break
		}
	}
	if tmp != 0 {
		digits = make([]int, n+1)
		digits[0] = 1
	}
	return digits
}

func main() {
	nums := [...]int{9}
	result := plusOne(nums[:])
	fmt.Println(result)
}
