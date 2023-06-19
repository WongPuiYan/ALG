package main

import (
	"fmt"
)

func climbStairs(n int) int {
	p, q := 1, 2
	if n == 1 {
		return p
	}
	if n == 2 {
		return q
	}
	result := 0
	for i := 3; i <= n; i++ {
		result = p + q
		p = q
		q = result
	}
	return result
}

func main() {
	result := climbStairs(4)
	fmt.Println(result)
}
