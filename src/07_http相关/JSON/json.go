package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Id     int
	Name   string
	Age    int
	gender string //注意：gender字段首字母必须大写，否则无法正常序列化
}

func main() {
	lily := Person{1, "Lily", 18, "female"}
	// json编码
	data, err := json.Marshal(&lily)
	if err != nil {
		fmt.Println("json err:", err)
		return
	}
	fmt.Println(string(data))
	// json解码
	var lily2 Person
	err = json.Unmarshal([]byte(data), &lily2) // 修正：正确语法
	if err != nil {                            // 修正：分离赋值和条件判断
		fmt.Println("json err:", err)
		return
	}
	fmt.Println(lily2.Name)
	fmt.Println(lily2.Age)
	fmt.Println(lily2.Id)

	fmt.Println(lily2.gender)

}
