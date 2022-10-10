package main

import "fmt"

func main() {
	// 常量，将var替换成const
	const aaa int = 111

	// 批量定义常量
	const (
		// iota 每行都会递增 , iota初始值为0
		// iota 只有在const中才能使用
		a = 10 * iota
		b
		c
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// 可读性灾难
	const (
		aa, bb = iota + 1, iota + 2
		cc, dd
		ee, ff = iota * 2, iota * 3
	)
	fmt.Println("aa=", aa)
	fmt.Println("bb=", bb)
	fmt.Println("cc=", cc)
	fmt.Println("dd=", dd)
	fmt.Println("ee=", ee)
	fmt.Println("ff=", ff)

}
