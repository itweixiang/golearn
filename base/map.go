package main

import "fmt"

func main() {
	// string 表示key ,表示value ,5表示key的长度
	//var kv map[string]int = make(map[string]int,5)
	// 省略写法
	var kv = make(map[string]int, 5)
	fmt.Println(kv)
	kv["h"] = 1
	// 容量不够的时候会自动扩容
	kv["e"] = 2
	fmt.Println(kv)

	// 不写容量也可以
	kv2 := make(map[string]int)
	kv2["l"] = 3
	fmt.Println(kv2)

	// 直接赋值的写法
	kv3 := map[string]int{
		//需要有逗号。。。。
		"l": 4,
		"o": 5,
	}
	fmt.Println(kv3)
}
