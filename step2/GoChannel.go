package main

import (
	"fmt"
	"time"
)

func main() {

	// cache channel
	ch := make(chan int, 2)

	go func() {
		for i := 0; i < 4; i++ {
			ch <- i
			fmt.Println("sub goroutine running")
			fmt.Println("send data: ", i, " length: ", len(ch), " capacity: ", cap(ch))
		}
		defer fmt.Println("sub goroutine shutting")
	}()
	time.Sleep(3 * time.Second)

	for i := 0; i < 4; i++ {
		num := <-ch
		fmt.Println("get data :", num)
	}
	fmt.Println("main goroutine shutting")
}
