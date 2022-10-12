package main

import "fmt"

// Hero大写表示外部可访问
type Hero struct {
	// 字段大写也是。。
	Name string
	Age  int
}

// (this Hero) 表示当前Hero结构体与该方法绑定
func (this Hero) GetName() string {
	return this.Name
}

// 不加*表示指针的话，相当于是拷贝。。相当无语
func (this *Hero) SetName(name string) {
	this.Name = name
}

func main() {
	hero := Hero{Age: 21, Name: "goodboy"}
	fmt.Println(hero.GetName())
	hero.SetName("badboy")
	fmt.Println(hero.GetName())
}
