package main

import "fmt"

// 父结构体
type Animal struct {
	Name string
}

func (a Animal) run() {
	fmt.Printf("%v在奔跑\n", a.Name)
}

//子结构体
type Dog struct {
	Age int
	//Animal // 结构体嵌套，也叫继承
	*Animal // 可以传结构体指针
}

func (d Dog) bark() {
	fmt.Printf("%v: wangwang", d.Name)
}

func main() {
	var d = Dog{
		Age: 8,
		Animal: &Animal{
			Name: "雪球",
		},
	}
	d.run()
	d.bark()

}
