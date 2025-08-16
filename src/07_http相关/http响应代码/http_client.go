package main

import (
	"fmt"
	"net/http"
)

func main() {

	client := http.Client{}

	resp, err := client.Get("https://www.baidu.com")
	if err != nil {
		fmt.Println("Get err:", err)
		return
	}

	ct := resp.Header.Get("Content-Type")
	date := resp.Header.Get("Date")
	server := resp.Header.Get("Server")
	fmt.Println("Ct:", ct)
	fmt.Println("Date:", date)
	fmt.Println("Server:", server)

}
