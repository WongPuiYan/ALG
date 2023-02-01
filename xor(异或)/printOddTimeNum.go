package main

import (
	"fmt"
)

// 一种数出现奇数次，其他数出现偶数次
func printOddTimeNum1(nums []int) {
	var eor = 0
	for _, element := range nums {
		eor ^= element
	}
	fmt.Println(eor)
}

// 两种数出现奇数次，其他数出现偶数次
func printOddTimeNum2(nums []int) {
	var eor = 0
	var _eor = 0
	for _, element := range nums {
		eor ^= element
	}
	// 提取出最右边的1
	var rightOne = eor & (^eor + 1)

	for _, element := range nums {
		if element&rightOne == 1 {
			fmt.Println("ele", element)
			_eor ^= element
		}
	}
	// fmt.Println(rightOne)
	eor = eor ^ _eor
	// _eor = eor ^ _eor
	fmt.Println(eor, _eor)
}

func main() {
	// oddNums1 := []int{1, 2, 2, 4, 4, 6, 6}
	// printOddTimeNum1(oddNums1)
	oddNums2 := []int{1, 2, 2, 3, 4, 4, 6, 6}
	printOddTimeNum2(oddNums2)
}
