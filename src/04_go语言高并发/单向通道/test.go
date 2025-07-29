package main

import (
	"fmt"
	"time"
)

func main() {
	// 单向读通道
	// var numChanReadOnly <- chan int;
	// 单向写通道
	// var numChanWriteOnly chan<- int;

	// 生产者消费者模型
	// consumer -> 提供只读通道
	// producer -> 提供只写通道

	//1. 在主函数中创建一个双向通道
	numChan := make(chan int, 5)
	// 将numChan 传递给producer
	// 将numChan 传递给consumer
	go producer(numChan)
	go consumer(numChan)
	time.Sleep(time.Second * 2)
	fmt.Println("main goroutine done")
}

func producer(numChan chan<- int) {
	for i := 0; i < 10; i++ {
		numChan <- i
		fmt.Println("生产数据：", i)
	}
}

func consumer(numChan <-chan int) {
	for v := range numChan {
		fmt.Println("消费数据：", v)
	}
}
