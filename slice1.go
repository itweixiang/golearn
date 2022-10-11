package main

import "fmt"

func main() {
	// 开辟一个长度3，容量为5的数组。没有声明容量的话，容量就和长度一样
	var array = make([]int, 3, 5)
	fmt.Printf("len = %d , cap = %d , ", len(array), cap(array), array)
	fmt.Println()

	// 向数组追加一个元素1
	array = append(array, 1)
	array = append(array, 1)

	// 当容量不够时，会扩容一个初始容量
	array = append(array, 1)
	fmt.Printf("len = %d , cap = %d , ", len(array), cap(array), array)
	fmt.Println()

	// : 左右可以填也可以不填，指定起始位置和结束位置
	// ! 注意，不是开辟了新的地址空间，而是在原先的地址上操作。相当于浅拷贝
	split := array[:7]
	fmt.Printf("len = %d , cap = %d , ", len(split), cap(split), split)

	copy := copy(array)
}
