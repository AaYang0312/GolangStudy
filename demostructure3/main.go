package main

import "fmt"

// 结构体的匿名字段
// type Person struct {
// 	string
// 	int
// }

type Person struct {
	Name  string
	Age   int
	Hobby []string
	map1  map[string]string
}

func (p Person) printPersonInfo() {
	fmt.Printf("姓名：%v\n年龄：%d\n爱好：%v\n补充信息：%v", p.Name, p.Age, p.Hobby, p.map1)
}

func main() {
	// p1 := Person{
	// 	"张三",
	// 	20,
	// }
	// fmt.Println(p1)
	var p Person
	p.Name = "李四"
	p.Age = 18
	p.Hobby = make([]string, 3)

	p.Hobby[0] = "写代码"
	p.Hobby[1] = "打篮球"
	p.Hobby[2] = "打游戏"

	p.map1 = make(map[string]string)

	p.map1["adress1"] = "北京"
	p.map1["adress2"] = "伯明翰"

	p.printPersonInfo()
}
