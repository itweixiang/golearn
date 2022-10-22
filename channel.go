package main

import "fmt"

func main() {
	//创建一个channel ，指定int类型，cap默认0
	//ch := make(chan int)
	ch := make(chan int, 10)
	fmt.Println("length is ", len(ch), " cap is ", cap(ch))
	go func() {
		defer fmt.Println("go runtime is over")
		fmt.Println("go runtime is running ")
		// 将值赋给channel。如果没有读取、channel也没有缓冲区的话，那么协程会一直阻塞
		// 当然，缓冲区满了也会阻塞
		ch <- 100
		// 关闭channel。如果不关闭，其他读取的地方可能会一直阻塞，形成死锁
		close(ch)
	}()
	// <- 表示读取channel的数据，并赋值。也可以直接读不赋值
	// 该操作会阻塞，直到读取到ch的值
	num := <-ch
	fmt.Println("num is ", num)

	// ok=true，表示channel没有关闭。false表示已经关闭
	if data, ok := <-ch; ok {
		fmt.Println("data is ", data)
	} else {
		fmt.Println("channel was closed")
	}

	// 可以用range直接读取channel中的值
	for data := range ch {
		fmt.Println(data)
	}
}
