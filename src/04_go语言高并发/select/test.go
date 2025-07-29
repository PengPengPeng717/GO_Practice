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
		close(chan1) // 发送完关闭
	}()
	go func() {
		for i := 0; i < 10; i++ {
			chan2 <- i
			time.Sleep(time.Second)
		}
		close(chan2) // 发送完关闭
	}()

	chan1Closed, chan2Closed := false, false
	for {
		if chan1Closed && chan2Closed { // 两个管道都关闭了
			break
		}
		fmt.Println("监听中。。。")
		select {
		case v, ok := <-chan1:
			if ok {
				fmt.Println("chan1:", v)
			} else {
				chan1Closed = true
			}
		case v, ok := <-chan2:
			if ok {
				fmt.Println("chan2:", v)
			} else {
				chan2Closed = true
			}
		}
	}
	fmt.Println("over")
}
