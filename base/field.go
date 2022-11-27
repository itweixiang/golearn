package main

import "fmt"

var aa int
var bb int = 200
var cc = 300

// ！ 全局变量不能省略var关键字
// dd := 400

func main() {
	// 不赋值 默认值是0
	var a int
	fmt.Println(a)

	// 指定类型
	var b int = 100
	fmt.Println(b)

	// 省略类型
	var c = 1000
	fmt.Println(c)
	// 打印类型
	fmt.Printf("c type is T%", c)
	fmt.Println()

	// 省略var 关键字
	d := 10000
	fmt.Println(d)

	// 多个变量
	var xx, yy int = 100, 200
	fmt.Println("xx =", xx, "yy =", yy)

	var kk, ll = 100, "abcd"
	fmt.Println("kk =", kk, "ll =", ll)

	// 也可以括号包起来
	var (
		aaa int    = 100
		bbb string = "abcd"
	)
	fmt.Println("aaa =", aaa)
	fmt.Println("bbb =", bbb)
}
