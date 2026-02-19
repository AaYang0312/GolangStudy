package main

import "fmt"

func main() {
	// var num float32 = 3.12

	// fmt.Printf("num=%f, type=%T\n", num, num)
	// fmt.Println(unsafe.Sizeof(num))

	// fmt.Println(unsafe.Sizeof(float64(num)))

	// var a float32 = 3.14 //4字节
	// var b float64 = 3.14 //8字节
	// var c float64 = 3.1415925535

	// fmt.Printf("a=%v, type=%T\nmem=", a, a)
	// fmt.Println(unsafe.Sizeof(a))
	// fmt.Printf("b=%v, type=%T\nmem=", b, b)
	// fmt.Println(unsafe.Sizeof(b))
	// fmt.Printf("c(float32)=%.4f, c(float64)=%v", c, c)

	// 64位系统中 go语言 float 类型默认 float64
	// f1 := 3.14
	// fmt.Printf("f1=%v, type=%T", f1, f1) //f1=3.14, type=float64

	// 科学计数法
	// var f2 float32 = 3.14e2      //表示 3.14*10的2次方
	// fmt.Printf("%v--%T", f2, f2) // 314--float32

	// 精度损失
	// var f4 float64 = 1129.6

	// fmt.Println(f4 * 100) // 112959.99999999999

	// int类型转换成float
	// a := 10
	// b := float64(a)

	// fmt.Printf("a=%d, type=%T\nb=%v, type=%T", a, a, b, b)
	// float类型转换成int，会截断，不建议

	// 布尔型，go语言中bool变量无法参与数值运算，也不能强制转换
	// 不赋值默认false
	// var flag bool = true
	// fmt.Printf("%v--%T", flag, flag) // true--bool

	// string型变量默认值为空
	// var s string
	// fmt.Printf("%v", s)

	// int型变量默认值为0
	// var i int
	// fmt.Printf("%v--%T", i, i) // 0--int

	// float型变量默认值为0
	// var f float32
	// fmt.Printf("%v--%T", f, f) // 0--float32

	var a = true
	if a {
		fmt.Println("1")
	}
}
