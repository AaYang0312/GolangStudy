package main

import "fmt"

func getUserinfo() (string, int) {
	return "zhangsan", 10
}
func main() {

	/*
		var 变量名称 类型

	*/
	// var username string
	// username = "zhangsan"
	// var username string = "zhangsan"
	// var a1 = "zhangsan"

	// fmt.Println(username)
	// fmt.Println(a1)

	// var m_ = "lisi"
	// fmt.Println(username)

	// var username = "zhangsan"
	// var age = 20
	// var sex = "nan"
	// fmt.Println(username, age, sex)

	// var a1, a2 string

	// a1 = "aaa"
	// a2 = "aaaaaaaaaaaa"

	// fmt.Println(a1, a2)

	/*
			var (
			username string
			age      int
			sex      string
		)

		username = "张三"
		age = 20
		sex = "男"

		fmt.Println(username, age, sex)
	*/

	/*
		username := "张三" //短变量声明法，只能用于func内
		fmt.Println("类型：%T", username)
	*/

	a1, a2, a3 := 12, 13, "C"
	fmt.Printf("a1的类型:%T，a2的类型:%T，a3的类型:%T", a1, a2, a3)
	fmt.Printf("a1:%v，a2:%v，a3:%v", a1, a2, a3)
	var username, _ = getUserinfo() //匿名变量
	var _, age = getUserinfo()      //匿名变量

	fmt.Println(username) //zhangsan
	fmt.Println(age)      //10
}
