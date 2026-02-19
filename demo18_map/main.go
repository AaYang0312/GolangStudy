package main

import "fmt"

func main() {
	// map 是基于 key, value 的数据
	// map[key]value

	// 1、创建和初始化
	// var userinfo = make(map[string]string)

	// userinfo["username"] = "David"
	// userinfo["age"] = "20"
	// userinfo["sex"] = "male"

	// fmt.Println(userinfo["username"])

	// 2、创建和初始化
	// var userinfo = map[string]string{
	// 	//userinfo := map...
	// 	"username": "David",
	// 	"age":      "20",
	// 	"sex":      "male",
	// }
	//fmt.Println(userinfo)

	// for range循环遍历map类型的数据
	// for key, value := range userinfo {
	// 	fmt.Println(key, "--", value)
	// }

	// map 类型的curd

	// 1、创建并修改 map 类型的数据
	// var userinfo = make(map[string]string)
	// userinfo["username"] = "zhangsan"
	// userinfo["age"] = "20"
	// // 2、获取查找 map 类型的数据
	// fmt.Println(userinfo["username"]) // 获取

	// v, ok := userinfo["xxxx"] // 查找
	// fmt.Println(v, ok) // 空 和 false

	// 3、删除 map 数据里的 key 以及对应的值

	var userinfo = map[string]string{
		"username": "David",
		"age":      "20",
		"sex":      "male",
		"height":   "180cm",
	}
	fmt.Println(userinfo)

	delete(userinfo, "sex")
	v, ok := userinfo["sex"]
	fmt.Println(userinfo)
	fmt.Println(v, "--", ok)
}
