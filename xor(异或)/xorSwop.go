package main

import (
	"fmt"
)

func main() {
	var (
		a, b int = 1, 2
	)
	fmt.Println(a, b)
	// 要保证a，b的指针不指向同一个内存地址。同一个内存位置的数据跟自己异或会将原数据洗成0
	a = a ^ b
	b = a ^ b
	a = a ^ b
	fmt.Println(a, b)
}
