package main

func main() {
	// 1、整型和整型转换
	// var a int8 = 20
	// var b int16 = 40
	// fmt.Println(int16(a) + b)

	// 2、浮点和浮点，省略
	// 3、整型和浮点转换
	// var a float32 = 3.12
	// var b int = 6
	// fmt.Println(a + float32(b))
	// 建议低位转换为高位，否则可能丢失位数

	// 4、转换为string类型
	// 4.1、使用Sprintf()
	// var a int = 123
	// var b string = fmt.Sprintf("%d", a) // (float-%f, bool-%t, byte-%c)
	// fmt.Printf("%v--%T", b, b) // 123--string

	// 4.2、使用strconv包

}
