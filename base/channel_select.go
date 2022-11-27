package main

import "fmt"

func main() {
	data := make(chan int)
	flag := make(chan int)

	go func() {
		for i := 0; i <= 100; i++ {
			data <- i
		}
	}()

	go func() {
		for i := 0; i <= 100; i++ {
			fmt.Println("data is ", <-data)
		}
		flag <- 0
	}()

	for {
		// 一个select可以监控多个channel
		select {
		case <-flag:
			close(flag)
			close(data)
			return
		}
	}
}
