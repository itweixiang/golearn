package main

import "fmt"

func main() {
	var a int = 1
	changeValue(a)
	// 跟java一样
	fmt.Println("a =", a) // a = 1

	// &表示传入地址
	changeValue1(&a)
	fmt.Println("a =", a) // a = 10

	var p = &a
	// p = 20指针类型，无法直接赋值。
	*p = 20
	// &地址，*值
	fmt.Println(*p)
}

func changeValue(p int) {
	p = 10
}

// *int int类型的指针类型
func changeValue1(p *int) {
	// p 为指针类型，无法直接赋值。
	// 需要使用*，获取指针类型指向的地址，才能赋值。
	*p = 10
}
