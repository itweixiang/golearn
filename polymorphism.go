package main

import "fmt"

// interface 本质是一个指针
type Animal interface {
	Sleep()
}

type Cat struct {
	//Animal
	Name string
}

// 拥有和Animal中的所有方法，则为其子类。真的是zz
// 为了更好的工程化，最好在strut中显示Animal
func (this *Cat) Sleep() {
	fmt.Println(this.Name, "is sleep")
}

func main() {
	// 多态
	// &获取地址给animal
	var cat Animal = &Cat{"Puff"}
	cat.Sleep()
}
