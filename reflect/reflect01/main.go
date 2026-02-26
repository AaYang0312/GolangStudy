package main

import (
	"fmt"
	"reflect"
)

/*
反射：
1. 可在程序运行期间动态获取变量的各种信息，比如变量类型 类别
2. 如果是结构体，通过反射可以获取结构体本身信息，结构体字段，结构体方法
3. 可以修改变量的值，调用关联的方法

golang中的变量分为两个部分:
1. 类型信息 reflect.TypeOf()
2. 值信息	reflect.ValueOf()
*/

type myInt int
type Person struct {
	Name string
	Age  int
}

func (p Person) eat() {
	fmt.Println("吃饭")
}

// 反射获取变量类型

func reflectFn(x interface{}) {
	v1 := reflect.TypeOf(x)
	// v2 := reflect.ValueOf(x)
	// v1.Name() // 类型名称
	// v1.Kind() // 底层种类
	fmt.Printf("类型: %v, 类型名称: %v, 类型种类: %v\n", v1, v1.Name(), v1.Kind())
}

// func main() {
// 	a := 10
// 	reflectFn(a)
// }

// 反射获取自定义类型和结构体信息

func main() {
	var a myInt
	p := Person{
		"张三",
		20,
	}
	var h = 25
	var i = [3]int{1, 2, 3}
	var j = []int{11, 22, 33}
	reflectFn(a)
	reflectFn(p)
	reflectFn(p.Age)
	reflectFn(&h)
	reflectFn(i)
	reflectFn(j)
	/*
		类型: main.myInt, 类型名称: myInt, 类型种类: int
		类型: main.Person, 类型名称: Person, 类型种类: struct
		类型: int, 类型名称: int, 类型种类: int
		类型: *int, 类型名称: , 类型种类: ptr
		类型: [3]int, 类型名称: , 类型种类: array
		类型: []int, 类型名称: , 类型种类: slice
	*/
}
