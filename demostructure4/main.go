package main

import "fmt"

// 结构体的嵌套
type User struct {
	Username string
	Password string
	Gender   string
	Age      int
	Address  // User 结构体嵌套 Address 结构体,只写一个可以匿名
}
type Address struct {
	Name  string
	Phone string
	City  string
}

func main() {
	var u User

	u.Username = "Louis"
	u.Password = "123456"
	u.Address.Name = "Aemeath"
	u.Address.Phone = "88888888"
	u.Address.City = "Lahai roi"

	u.City = "Startorch" // 嵌套的结构体为匿名结构体时可以这样写，使用时程序先查找结构内变量，再查找嵌套的结构体

	fmt.Printf("%#v", u)
}
