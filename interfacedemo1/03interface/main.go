package main

import "fmt"

func main() {
	// 类型断言
	// var a interface{}
	// var b any

	// a = "Hello World!"
	// b = "Hello World!"

	// v1, ok1 := a.(string)
	// v2, ok2 := b.(string)

	// fmt.Println(v1, ok1, v2, ok2) // Hello World! true Hello World! true

	// 根据不同类型实现不同功能
	var a interface{}
	a = "123"
	switch a.(type) {
	case int:
		fmt.Println("int类型")
	case string:
		fmt.Println("string类型")
	case bool:
		fmt.Println("bool类型")
	default:
		fmt.Println("不受支持的类型")
	}
}
