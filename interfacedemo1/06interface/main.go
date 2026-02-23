package main

import "fmt"

type Address struct {
	City  string
	Phone int
}

func main() {
	var userinfo = make(map[string]any)
	var address = Address{
		"Beijing",
		13555555555,
	}

	userinfo["name"] = "Jimmy"
	userinfo["age"] = 28
	userinfo["hobby"] = []string{"睡觉", "吃饭", "开车"} // 空接口不支持索引，无法精确调用
	userinfo["address"] = address                  // 可通过外部结构体对象访问

	// fmt.Printf("%#v", userinfo)
	fmt.Println(address.City)

	// 解决空接口类型不能使用索引的方法，把接口类型强制转换为对应类型
	hobby2, _ := userinfo["hobby"].([]string)
	fmt.Println(hobby2[0]) // 睡觉
}
