package main

import "fmt"

func main() {
	// 仅做声明，但还没有开辟空间，直接赋值会报错
	// make 开辟容量3的空间
	var array1 []int = make([]int, 3)
	fmt.Println("length =", len(array1), array1)

	// 简写
	// var array1  = make([]int,3)

	// 声明并赋初值
	array2 := []int{1, 2, 3}
	if array2 == nil {
		fmt.Println("array2 is nil")
	}
}
