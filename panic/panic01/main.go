package main

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 0; i < 10; i++ {
		fmt.Printf("Hello World!\n")
	}
}

func test() {
	// 定义一个 defer + recover 匿名函数用于捕获 panic
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("test()发生错误: ", err) // test()发生错误:  assignment to entry in nil map， 要用 make 来创建 map
		}
	}()
	var myMap map[int]string
	myMap[0] = "golang"
}
func main() {
	go sayHello()
	go test()

	time.Sleep(time.Second) // 防止主进程退出
}
