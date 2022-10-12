package main

import "fmt"

type Human struct {
	Sex  int
	Name string
}

type Man struct {
	// 继承Human类的属性
	Human
	length int
}

func (this *Human) Eat() {
	fmt.Println(this.Name, "is eating")
}

func main() {
	// 看着很像组合
	// man := Man{Human{1, "goodboy"}, 10}
	// 稍微不那么变扭的方式。。
	var man Man
	man.Name = "goodboy"
	man.Eat()
}
