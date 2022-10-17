package main

import "fmt"

// interface{}表示万能类型
func method(args interface{}) {
	// 断言语法，只有在万能类型时才能使用
	value, ok := args.(string)
	if ok {
		fmt.Println(value)
	} else {
		fmt.Printf("value = %T , args = %T", value, args)
	}
}

type Dog struct {
	Name string
}

func main() {
	method("123")
	array := []string{}
	method(array)
	fmt.Println()
	dog := Dog{"dog"}
	method(dog)
}
