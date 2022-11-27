package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	// json需要加上特定的注解或者叫做标签
	Title string  `json:"title"`
	Price float32 `json:"price"`
}

func main() {
	movie := Movie{Title: "功夫", Price: 32.5}
	// 对象转json字符串，匿名的是错误
	str, _ := json.Marshal(movie)
	fmt.Printf("%s", str)
	fmt.Println()

	// 反序列化
	empty_movie := Movie{}
	json.Unmarshal(str, &empty_movie)
	fmt.Println("title =", empty_movie.Title, "price =", empty_movie.Price)

}
