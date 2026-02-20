package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Id     int
	Gender string
	Name   string
}
type Class struct {
	Title    string
	Students []Student
}

func main() {
	// 实例化Class
	var c1 = Class{
		"2201",
		make([]Student, 0, 40),
	}
	for i := 0; i < 10; i++ {
		s := Student{
			i,
			"Male",
			fmt.Sprintf("stu_%v", i),
		}
		c1.Students = append(c1.Students, s)
	}
	// fmt.Printf("%#v", c1)
	// 把嵌套的结构体转换成 json
	var jsonC1, _ = json.Marshal(c1)
	var jsonStr = string(jsonC1)
	fmt.Println(jsonStr)
}
