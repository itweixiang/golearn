package main

import "fmt"

// 单个参数
func func1(a int) int {
	// 单返回值
	return a + 1
}

// 多参
func func2(a int, b string) (int, string) {
	// 多返回值
	return a + 1, b + "d"
}

// 可以给返回值取个别名，这个好一点
func func3(a int) (an int) {
	an = a + 1
	// 自动返回an
	return
}

// 多返回值，多别名，a1，a2默认值为0
func func4(a int, b int) (a1, a2 int) {
	fmt.Println(a1)
	fmt.Println(a2)
	a1 = a + 1
	a2 = b + 1
	return
}

func main() {
	fmt.Println(func1(111))

	// 接受返回值
	var a, b = func2(111, "abc")
	aa, bb := func2(111, "abc")
	fmt.Println(a, b)
	fmt.Println(aa, bb)

	an := func3(111)
	fmt.Println(an)
}
