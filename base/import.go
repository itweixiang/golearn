package main

import (
	// 直接导入，不使用会报错
	//"golearn/lib"
	// 匿名导入，不使用也会调用init方法
	//_ "golearn/lib"
	// 静态导入
	. "golearn/lib"
)

func main() {
	//直接导入
	//lib.Method()
	// 静态导入
	Method()
}
