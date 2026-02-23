package main

import "fmt"

type Animaler1 interface {
	SetName(string)
	GetName() string
}
type Animaler2 interface {
	run()
}

// 接口的嵌套
type Animaler interface {
	Animaler1
	Animaler2
}

type Dog struct {
	Name string
}

func (d *Dog) SetName(name string) { // 指针类型的接收者
	d.Name = name
}
func (d Dog) GetName() string {
	return d.Name
}
func (d Dog) run() {
	fmt.Printf("%v在奔跑\n", d.Name)
}
func main() {
	// 一个结构体实现多个接口
	// d1 := &Dog{
	// 	"snowball",
	// }
	// var a1 Animaler1
	// var a2 Animaler2
	// a1 = d1
	// a2 = d1
	// a1.SetName("wangcai")
	// fmt.Println(a1.GetName()) // wangcai
	// a2.run()                  // wangcai在奔跑

	// 接口嵌套
	d1 := &Dog{
		"wangcai",
	}
	var a Animaler
	a = d1
	a.run()
	a.SetName("雪球\n")
	fmt.Printf("%v\n", a.GetName())
}
