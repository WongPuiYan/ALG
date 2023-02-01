package main

import (
	// "crypto/rand"
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// 随机数组生成器
func generateRandomArray(maxSize int, maxValue int) []int {
	rand.Seed(time.Now().UnixNano())
	arr := []int{}
	for i := 0; i < int(rand.Float64()*float64(maxSize+1)); i++ {
		arr = append(arr, int(rand.Float64()*float64(maxValue+1))-int(rand.Float64()*float64(maxValue+1)))
	}
	return arr
}

func copyArray(arr []int) []int {
	copyArr := []int{}
	for i := 0; i < len(arr); i++ {
		copyArr = append(copyArr, arr[i])
	}
	return copyArr
}

func isEqual(arr1 []int, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

// 冒泡排序
func bubbleSort(arr []int) []int {
	numsLength := len(arr)
	if numsLength <= 1 {
		return arr
	}
	for i := 0; i < numsLength; i++ {
		for j := 0; j < numsLength-1-i; j++ {
			// 大数移到队列后面
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

// 选择排序
func selectionSort(arr []int) []int {
	numsLength := len(arr)
	if numsLength <= 1 {
		return arr
	}
	for i := 0; i < numsLength-1; i++ {
		minIndex := i
		for j := i + 1; j < numsLength; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		// 最小数直接与队列前面的元素交换
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
	return arr
}

// 插入排序
func insertionSort(arr []int) []int {
	numsLength := len(arr)
	if numsLength <= 1 {
		return arr
	}
	for index := range arr {
		preIndex := index - 1
		current := arr[index]
		// 小数逐个对比移到队列前面
		for ; preIndex >= 0 && arr[preIndex] > current; preIndex-- {
			arr[preIndex+1] = arr[preIndex]
		}
		arr[preIndex+1] = current
	}
	return arr
}

func main() {
	testTime := 1000000
	maxSize := 1000000
	maxValue := 1000000
	for i := 0; i < testTime; i++ {
		arr := generateRandomArray(maxSize, maxValue)
		compareArray := copyArray(arr)
		sort.Ints(compareArray)
		// fmt.Println("origin array: ", arr)
		// fmt.Println("result array: ", compareArray)

		resultArray := []int{}
		resultArray = bubbleSort(copyArray(arr))
		// fmt.Println("bubble result:", bubbleResult)
		if !isEqual(resultArray, compareArray) {
			fmt.Println("result array:", resultArray)
			fmt.Println("compare array:", compareArray)
			break
		} else {
			// fmt.Println("pass")
		}
		// fmt.Println(isEqual(resultArray, compareArray))

		resultArray = selectionSort(copyArray(arr))
		if !isEqual(resultArray, compareArray) {
			fmt.Println("result array:", resultArray)
			fmt.Println("compare array:", compareArray)
			break
		} else {
			// fmt.Println("pass")
		}
		// fmt.Println("select result:", selectionResult)
		// fmt.Println(isEqual(resultArray, compareArray))

		resultArray = insertionSort(copyArray(arr))
		if !isEqual(resultArray, compareArray) {
			fmt.Println("result array:", resultArray)
			fmt.Println("compare array:", compareArray)
			break
		} else {
			// fmt.Println("pass")
		}
		// fmt.Println("insert result:", insertionResult)
		// fmt.Println(isEqual(resultArray, compareArray))
		fmt.Println(i)
	}

}
