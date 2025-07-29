package main

import (
	"fmt"
	"time"
)

func main() {
	//numsChan := make(chan int, 10)
	//1. 当缓冲写满的时候，写阻塞，当被读取后，再恢复写入
	//2. 当缓冲区读取完毕，读阻塞
	//3. 如果管道没有使用make分配空间，那么管道默认是nil的，读取、写入都会阻塞
	//4. 对于一个管道，读与写的次数，必须对等

	var names chan string //默认是nil的
	names = make(chan string, 10)

	go func() {
		fmt.Println("names:", <-names)
	}()

	names <- "hello" //由于names是nil的，写操作会阻塞在这里
	time.Sleep(1 * time.Second)

	numsChan1 := make(chan int, 10)

	//写
	go func() {
		for i := 0; i < 60; i++ {
			// for i := 0; i < 50; i++ {
			numsChan1 <- i
			fmt.Println("写入数据:", i)
		}
	}()

	//读,当主程序被管道阻塞时，那么程序将锁死崩溃
	//要求我们一定要读写次数保持一致
	func() {
		for i := 0; i < 60; i++ {
			fmt.Println("主程序准备读取数据.....")
			data := <-numsChan1
			fmt.Println("读取数据:", data)
		}
	}()

	for {

	}
}
