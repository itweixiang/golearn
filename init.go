package main

import "fmt"

// 对外可见的方法首字母大写
func Method1() {
	fmt.Println("lib1 method is run")
}

// 其实相当于Java中的无参构造
func init() {
	fmt.Println("lib1 init is run")
}
