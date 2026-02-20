package main

import (
	"encoding/json"
	"fmt"
)

// 将小写属性转换成json方法(结构体标签)
type Student struct {
	Id     int    `json:"id"`
	Gender string `json:"gender"`
	Name   string `json:"name"`
	Sno    string `json:"sno"`
}

func main() {
	var s1 = Student{
		12,
		"female",
		"Alice",
		"s00001",
	}
	byteS1, _ := json.Marshal(s1)
	strS1 := string(byteS1)
	fmt.Printf("%v", strS1) // {"id":12,"gender":"female","name":"Alice","sno":"s00001"}
}
