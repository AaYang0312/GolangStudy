package main

import "fmt"

// type A interface{} // 空接口
// func show(a interface{}) {
// 	fmt.Printf("值:%v, 类型:%T\n", a, a)
// }
func main() {
	// 空接口可以表示任何类型
	// var a interface{}
	// fmt.Printf("a的值:%v, a的类型:%T\n", a, a) // a的值:<nil>, a的类型:<nil>
	// a = 12
	// fmt.Printf("a的值:%v, a的类型:%T\n", a, a) // a的值:12, a的类型:int

	// a = "Hello World!"
	// fmt.Printf("a的值:%v, a的类型:%T\n", a, a) // a的值:Hello World!, a的类型:string

	// a = false
	// fmt.Printf("a的值:%v, a的类型:%T\n", a, a) // a的值:false, a的类型:bool

	// var b = false
	// show(b)
	// slice := []int{1, 12, 3, 4, 5}
	// show(slice)

	// 空接口在 map 中使用可以存储多种类型的 key
	var m1 = make(map[string]interface{})
	m1["username"] = "Cartethyia"
	m1["height"] = 160
	m1["tabemono"] = "Pizza Margherita"

	var s1 = []interface{}{1, "Hello", true}
	fmt.Println(m1)
	fmt.Println(s1)

	// golang 1.18+ 增加了泛型也可以实现
	var m2 = make(map[string]any)
	m2["username"] = "Cartethyia"
	m2["height"] = 160
	m2["tabemono"] = "Pizza Margherita"

	var s2 = []any{1, "Hello", true}
	fmt.Println(m2)
	fmt.Println(s2)
}
