package main

import "fmt"

func main() {
	city := make(map[string]string)
	city["China"] = "beking"
	city["Japan"] = "Tokyo"
	city["USA"] = "NewYork"
	// 遍历
	for key, value := range city {
		fmt.Println("key =", key, ", value =", value)
	}
	// 删除
	delete(city, "Japan")
}
