package main

import (
	"fmt"
	"runtime"
	"time"
)

func m1() {
	for i := 0; i < 100; i++ {
		fmt.Printf("i = %d\n", i)
		time.Sleep(1 * time.Second)
	}
	func() {
		fmt.Println("匿名函数退出")
		// 仅仅会退出当前的协程
		return
	}()
	// 退出当前程序
	runtime.Goexit()
}
func main() {
	// 创建一个协程
	go m1()

	// 创建一个匿名函数，但是跑在协程上
	go func(a int, b int) {
		for k := 0; k < 100; k++ {
			fmt.Printf("k = %d\n", k)
			time.Sleep(1 * time.Second)
		}
		// () 表示调用。1，2表示传参
	}(1, 2)
	// for {} 表示死循环
	for {
		fmt.Println("main thread is running")
		time.Sleep(1 * time.Second)
	}
}
