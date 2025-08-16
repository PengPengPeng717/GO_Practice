package main

import (
	"fmt"
	"time"
)

func main() {
	chan1 := make(chan int)
	chan2 := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			chan1 <- i
			time.Sleep(time.Second / 2)
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			chan2 <- i
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for {
			fmt.Println("监听中。。。。")
			select {
			case data1 := <-chan1:
				fmt.Println("从chan1获取数据:", data1)
			case data2 := <-chan2:
				fmt.Println("从chan2获取数据:", data2)
			default:
				fmt.Println("11111")
				time.Sleep(time.Second)
			}
		}
	}()
	for {
		fmt.Print("over")
		time.Sleep(time.Second * 10)
	}
}
