package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8848")
	if err != nil {
		fmt.Println("Error dialing:", err.Error())
		return // 终止程序
	}
	defer conn.Close()
	fmt.Println("Connect to 127.0.0.1:8080")
	sendData := []byte("hello, server")
	// 发送数据
	cnt, err := conn.Write(sendData)
	if err != nil {
		fmt.Println("Write err:", err)
		return
	}
	fmt.Println("真正发送给服务器的数据长度：", cnt)
	fmt.Println("发送给服务端的数据：", string(sendData))
	// 接收数据
	buf := make([]byte, 1024)
	cnt, err = conn.Read(buf)
	if err != nil {
		fmt.Println("Read err:", err)
		return
	}
	fmt.Println("真正读取server发来的数据长度：", cnt)
	fmt.Println("接收到server发来的数据：", string(buf[:cnt]))
}
