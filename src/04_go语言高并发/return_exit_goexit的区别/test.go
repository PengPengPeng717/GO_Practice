package main

import (
	"fmt"
	"time"
)

// 这个将用于子go程使用
func display(num int) {
	count := 1
	for {
		fmt.Println("=============> 这是子go程:", count, "num:", num)
		count++
		time.Sleep(1 * time.Second)
	}
}

func main() {
	// fmt.Println("Hello World!")
	//启动子go程
	// go display()
	for i := 0; i < 3; i++ {
		go display(i)
	}

	//主go程
	count := 1
	for {
		fmt.Println(" 这是主go程:", count)

		//休眠1秒
		count++
		time.Sleep(1 * time.Second)
	}

}
