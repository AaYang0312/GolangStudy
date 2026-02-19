package main

import "fmt"

func main() {
	// 1、golang中定义字符 字符属于int类型

	// a := 'a'
	// b := 'b'
	// fmt.Printf("value:%v, type:%T\n", a, a) // value:97, type:int32
	// fmt.Printf("value:%v, type:%T\n", b, b) // value:98, type:int32
	// // 默认输出ASCII码值

	// // 2、原样输出
	// fmt.Printf("value:%c type:%T\n", a, a) // value:a type:int32
	// fmt.Printf("value:%c type:%T\n", b, b) // value:b type:int32

	// // 3、定义一个字符串，输出里面的字符
	// str := "this"
	// fmt.Printf("value:%v, 原样输出:%c type:%T\n", str[2], str[2], str[2]) // value:105, 原样输出:i type:uint8

	// 4、一个汉字占用3字节(utf-8), 一个字母占用一个字节
	// unsafe.Sizeof() 没法查看string类型数据所占用的存储空间
	// strhan := "你"
	// str := "this" // 4个字节
	// str1 := "你好ok"
	// fmt.Println(len(strhan)) // 3
	// fmt.Println(len(str))    // 4
	// fmt.Println(len(str1))   // 8 = 3 + 3 + 1 + 1

	// 5、通过循环输出字符串中的字符
	// str := "你好 golang"
	// count := len(str)
	// for i := 0; i < count; i++ { // byte(用于处理ASCII码)
	// 	fmt.Printf("%v(%c) ", str[i], str[i])
	// }

	// for _, v := range str { // rune(用于处理中文日文等复合文本，即Unicode编码，兼容ASCII码)
	// 	fmt.Printf("%v(%c) ", v, v)
	// }

	// 6、修改字符串，先将其转换成[]rune 或[]byte，完成后再转换为 string。需要重新分配内存，复制字节数组
	s1 := "big"
	s2 := "你好"
	// 强制类型转换
	byteS1 := []byte(s1)
	byteS1[0] = 'p'
	fmt.Println(string(byteS1))

	runeS2 := []rune(s2)
	runeS2[0] = '我'
	fmt.Println(string(runeS2))
}
