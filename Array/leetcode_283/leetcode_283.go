package main

import(
	"fmt"
)

func moveZeroes(nums []int)  {
	var slow int
	for fast := 0; fast < len(nums); fast++{
		if nums[fast] != 0{
			nums[slow] = nums[fast]
			slow++
		}
	}

	for ; slow < len(nums); slow++{
		nums[slow] = 0
	}
}

func main(){
	var 
	moveZeroes()
}