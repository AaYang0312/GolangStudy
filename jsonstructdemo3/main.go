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
	// var jsonC1, _ = json.Marshal(c1)
	// var jsonStr = string(jsonC1)
	// fmt.Println(jsonStr)

	// 反向转换成结构体
	var strJson = `{"Title":"2201","Students":[{"Id":0,"Gender":"Male","Name":"stu_0"},{"Id":1,"Gender":"Male","Name":"stu_1"},{"Id":2,"Gender":"Male","Name":"stu_2"},{"Id":3,"Gender":"Male","Name":"stu_3"},{"Id":4,"Gender":"Male","Name":"stu_4"},{"Id":5,"Gender":"Male","Name":"stu_5"},{"Id":6,"Gender":"Male","Name":"stu_6"},{"Id":7,"Gender":"Male","Name":"stu_7"},{"Id":8,"Gender":"Male","Name":"stu_8"},{"Id":9,"Gender":"Male","Name":"stu_9"}]}`

	var c1r = &Class{}
	err := json.Unmarshal([]byte(strJson), c1r)
	if err != nil {
		fmt.Println("Fail")
	}
	fmt.Printf("%v", c1r.Title)
}
