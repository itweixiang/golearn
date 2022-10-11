package main

import "fmt"

func main() {
	defer fmt.Println("main is run4")
	fmt.Println("main is run1")
	// 在函数结束时执行，无论在函数中的哪个位置。多个的话，根据倒序执行
	defer fmt.Println("main is run3")
	fmt.Println("main is run2")
}
