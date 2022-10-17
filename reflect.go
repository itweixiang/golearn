package main

import (
	"fmt"
	"reflect"
)

type Phone struct {
	//`` 给字段加上注解
	cpu int `info:"cpu tag",example:"3500g"`
}

func (p Phone) play() {
	fmt.Println("play phone on ", p.cpu)
}

func main() {
	array := make([]int, 5)
	array[0] = 1
	//获取变量的类型和值
	fmt.Println(reflect.ValueOf(array))
	fmt.Println(reflect.TypeOf(array))

	phone := Phone{1}
	valueOf := reflect.ValueOf(phone)
	typeOf := reflect.TypeOf(phone)
	//获取类的字段
	for i := 0; i < typeOf.NumField(); i++ {
		field := typeOf.Field(i)
		fmt.Println("tag =", field.Tag.Get("info"))
		fmt.Println("name =", field.Name, "type =", field.Type, "value =", valueOf.Field(i))
	}

	//调用类的方法
	for i := 0; i < typeOf.NumMethod(); i++ {
		method := typeOf.Method(i)
		fmt.Println(method.Name, method.Type)
		//call := method.Func.Call([]reflect.Value{})
	}

}
