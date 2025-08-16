package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// https//127.0.0.1:8080/name
	// http.ResponseWriter 通过http.ResponseWriter将数据返回给客户端
	// http.Request 包含客户端发来的数据
	http.HandleFunc("/name", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("request:", r)
		io.WriteString(w, "这是/name页面")
	})
	// https//127.0.0.1:8080/age
	http.HandleFunc("/age", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintf(w, "18")
		fmt.Println("request:", r)
		io.WriteString(w, "这是/age页面")
	})
	// https//127.0.0.1:8080/id
	http.HandleFunc("/id", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintf(w, "Hello World!")
		fmt.Println("request:", r)
		io.WriteString(w, "这是/id页面")
	})

	fmt.Print("http server start at 127.0.0.1:8080")
	if err := http.ListenAndServe("127.0.0.1:8080", nil); err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

}
