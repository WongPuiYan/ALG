package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateRandomArray(maxSize int, maxValue int) []int {
	rand.Seed(time.Now().UnixNano())
	arr := []int{}
	for i := 0; i < int(rand.Float32()*float32(maxSize+1)); i++ {
		arr = append(arr, int(rand.Float32()*float32(maxValue+1))-int(rand.Float32()*float32(maxValue+1)))
	}
	return arr
}

func binarySearch(arr []int) int {
	left, right := 0, len(arr)-1
	mid := left + (right-left)>>1
	fmt.Println(left, mid, right)
	return 0
}

func main() {
	testTime := 10
	maxSize := 10
	maxValue := 100
	for i := 0; i < testTime; i++ {

		arr := generateRandomArray(maxSize, maxValue)
		fmt.Println(arr)
		binarySearch(arr)
		break
	}
}
