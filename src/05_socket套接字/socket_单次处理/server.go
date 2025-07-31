package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	//创建监听
	ip := "127.0.0.1"
	port := 8848
	adrress := fmt.Sprintf("%s:%d", ip, port)

	listener, err := net.Listen("tcp", adrress)
	// net.Listen("tcp", "127.0.0.1:8848")//简写，冒号前面默认是本机：127.0.0.1
	if err != nil {
		fmt.Println("listen err:", err)
		return
	}
	fmt.Println("监听中")
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Accept err:", err)
		return
	}
	defer conn.Close()
	fmt.Println("客户端已连接：", conn.RemoteAddr().String())
	//创建一个容器，用于接收读取到的数据
	buf := make([]byte, 1024) //使用make来创建字节切片
	// cnt:真正读取client发来的数据的长度
	cnt, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Read err:", err)
		return
	}
	fmt.Println("真正读取client发来的数据长度：", cnt)
	data := string(buf[:cnt])
	fmt.Println("接收到的数据：", data)
	// fmt.Println("接收到的数据：", string(buf[:cnt]))

	//服务器对客户端请求进行响应，将数据转换成大写
	upperData := strings.ToUpper(string(buf[:cnt]))

	cnt, err = conn.Write([]byte(upperData))
	if err != nil {
		fmt.Println("Write err:", err)
		return
	}
	fmt.Println("真正发送给客户端的数据长度：", cnt)
	fmt.Println("发送给客户端的数据：", upperData)
	// for {
	// 	fmt.Println("等待客户端连接...")
	// 	conn, err := listener.Accept()
	// 	if err != nil {
	// 		fmt.Println("Accept err:", err)
	// 		continue
	// 	}

	// 	fmt.Println("客户端已连接：", conn.RemoteAddr().String())

	// 	go func(conn net.Conn) {
	// 		defer conn.Close()

	// 		for {
	// 			buf := make([]byte, 1024)
	// 		}
	// 	}
	// }

}
