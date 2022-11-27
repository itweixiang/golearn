package main

import "fmt"

// 声明int类型的别名
type myint int

// 定义一个结构器
type Book struct {
	number int
	title  string
}

func main() {
	var i myint = 1
	//也算是新的类型
	fmt.Printf("type of i = %T", i)
	fmt.Println()

	// 结构体声明，感觉像Java的子类
	var book Book
	book.title = "论语"
	book.number = 1

	fmt.Printf("%v", book)
}
