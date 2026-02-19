package main

import "fmt"

type Person struct {
	Name string
	Age  int
	Sex  string
}

// 定义结构体方法
func (p Person) PrintInfo() {
	fmt.Printf("姓名：%v, 年龄：%d", p.Name, p.Age)
}

// 要修改属性必须使用指针
func (p *Person) SetInfo(name string, age int) {
	p.Age = age
	p.Name = name
}
func main() {
	var p1 = Person{
		Name: "张三",
		Age:  18,
		Sex:  "男",
	}

	// p2 := &p1
	// p2.Age = 20 // 通过指针修改对应值

	// fmt.Printf("%#v", p1)
	// fmt.Printf("%#v", p2)

	p1.SetInfo("李四", 22)
	p1.PrintInfo()
}
