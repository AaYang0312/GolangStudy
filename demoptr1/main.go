package main

import "fmt"

// 传入指针修改值
func fn1(x int) {
	x = 20
}
func fn2(x *int) {
	*x = 40
}
func main() {
	var a int = 10
	var p = &a
	fmt.Printf("a的值为%d, a的类型为%T, a的地址为%p\n", a, a, &a)
	fn2(&a)
	fmt.Println(*p)

	// 定义map这样的引用数据类型时需要使用make，分配内存空间，切片也需要
	var userinfo = make(map[string]string)
	userinfo["name"] = "zhangsan"
	fmt.Println(userinfo["name"])

	var a1 = make([]int, 4)
	a1[0] = 1
	fmt.Println(a1)

	// 用new函数给指针分配内存
	b := new(int)
	c := new(bool)
	fmt.Printf("%T\n", b) // *int
	fmt.Printf("%T\n", c) // *bool
	fmt.Println(*b)       // 0
	fmt.Println(*c)       // false

	// 总结: make 用于 map, slice, channel 类型
	// new 用于 指针类型（较少用到）
}
