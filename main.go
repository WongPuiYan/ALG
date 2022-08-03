package main

import "fmt"

// 交换函数
func swap(a, b *int) {
	// 取a指针的值, 赋给临时变量t
	// t := *a
	// 取b指针的值, 赋给a指针指向的变量
	// *a, *b = *b, *a
	a, b = b, a
	// 将a指针的值赋给b指针指向的变量
	// *b = t
}
func main() {
	fmt.Println("Hello World!")
	// 准备两个变量, 赋值1和2
	x, y := 1, 2
	// 交换变量值
	// x, y = y, x
	swap(&x, &y)
	// 输出变量值
	fmt.Println(x, y)
}
