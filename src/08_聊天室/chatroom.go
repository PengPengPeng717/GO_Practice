package main

import (
	"fmt"
	"net"
)

func main() {
	// 创建服务器
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("listen err:", err)
		return
	}
	// defer listener.Close()
	fmt.Println("服务器启动成功")
	for {
		//监听
		fmt.Println("监听中")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept err:", err)
			return
		}
		// defer conn.Close()

		//建立连接
		fmt.Println("客户端已连接：", conn.RemoteAddr().String())

		//启动处理业务的go程
		go handler(conn)
	}

}

// 处理具体业务
func handler(conn net.Conn) {
	fmt.Println("处理业务")
	//TODO
	// 创建一个缓冲区
	buf := make([]byte, 1024)
	// 循环读取客户端发送的数据
	for {
		// 读取数据
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read err:", err)
			return
		}
		fmt.Println("接收到的数据：", string(buf[:n-1]), "长度：", n)
	}
}
