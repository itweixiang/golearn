package main

import "fmt"

func main() {
	// 固定长度的数组声明
	var array [10]int
	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}
	fmt.Println("---------")

	// 前3个元素时1，2，3，后面都是0
	array2 := [10]int{1, 2, 3}
	printArray1(array2)
	for index := range array2 {
		fmt.Println(array2[index])
	}
	fmt.Println("---------")

	// 动态数组，没有长度
	array3 := []int{1, 2, 3}
	//array3[100] = 100
	printArray2(array3)
	for index, value := range array3 {
		fmt.Println("index = ", index, ", value = ", value)
	}
	fmt.Println("---------")
}

// 普通数组是值传递
func printArray1(array [10]int) {
	array[0] = 100
	for _, index := range array {
		fmt.Println(index)
	}
	fmt.Println("---------")
}

// 动态数组是引用传递
func printArray2(array []int) {
	array[0] = 100
	for _, index := range array {
		fmt.Println(index)
	}
	fmt.Println("---------")
}
