package main

import "fmt"

type myInt int

// type fn func(int, int) int
type myFloat = float64
type myStru struct {
	age    myInt
	height myFloat
	name   string
	gender string
}

func main() {
	var a myInt = 10
	var b myFloat = 22.3
	/*
		var zhangsan myStru // 1.
		zhangsan := new(myStru) // 2. 创建后为结构体指针，但是可以通过zhangsan.name取值，底层实际为(*zhangsan).name
		zhangsan := &myStru{} // 3. 和new差不多

		zhangsan.name = "ZHANGSAN"
		zhangsan.age = 12
		zhangsan.height = 165.5
		zhangsan.gender = "male"
	*/

	var zhangsan = myStru{
		name:   "ZHANGSAN",
		age:    12,
		height: 165.5,
		gender: "male",
	} // 4. 键值对，等价于var后赋值，也可以不写key，但需要按顺序

	fmt.Printf("%v %T\n", a, a)                     // 10 main.myInt
	fmt.Printf("%v %T\n", b, b)                     // 22.3 float64
	fmt.Printf("值：%v，类型：%T\n", zhangsan, zhangsan)  // 值：{12 165.5 ZHANGSAN male}，类型：main.myStru
	fmt.Printf("值：%#v，类型：%T\n", zhangsan, zhangsan) // 值：main.myStru{age:12, height:165.5, name:"ZHANGSAN", gender:"male"}，类型：main.myStru
}
