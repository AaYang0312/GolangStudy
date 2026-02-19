package main

func main() {
	// 1、定义string类型

	// var str1 string = "nihao"
	// var str2 = "my name is Zhangsan"
	// str3 := "Nice to meet you"

	// fmt.Printf("%v,\r,%v\n,%v\t", str1, str2, str3)

	// 2、字符转义符
	// var str2 = "C:\\Go\\bin" // C:\Go\bin
	// var str3 = "\'hello\'" // 'hello'
	// var str4 = "\"hello\"" // "hello"
	// fmt.Printf("%s", str2)

	// 3、多行字符串
	// str1 := `this is str
	// this is str
	// this is str
	// this is str
	// this is str
	// `
	// fmt.Println(str1)

	// 4、求长度
	// str := "Hello world"
	// fmt.Println(len(str)) // 11

	// 5、拼接字符串
	// str1 := "Hello"
	// str2 := "world"
	// str3 := str1 + str2
	// str4 := fmt.Sprintf("%v, %v", str1, str2)
	// fmt.Println(str3)
	// fmt.Println(str4)

	// 6、分割字符串
	// var str1 = "123-456-789"
	// arr := strings.Split(str1, "-")

	// fmt.Println(arr) // [123 456 789] 切片，和数组有一些区别

	// 7、合并字符串，把切片连接成字符串
	// var str1 = "123-456-789"
	// arr := strings.Split(str1, "-")
	// str2 := strings.Join(arr, "")
	// fmt.Println(str2)

	// arr := []string{"php", "java", "golang"}
	// fmt.Println(arr)
	// str3 := strings.Join(arr, "-")
	// fmt.Println(str3)

	// 7、strings.contains 判断是否包含
	// str1 := "this is str"
	// str2 := "this"
	// flag := strings.Contains(str1, str2)

	// fmt.Printf("%v--%T", flag, flag) //true--bool

	// 8、strings.HasPrefix,strings.HasSuffix 前缀后缀判断
	// str1 := "this is str"
	// str2 := "this"
	// str3 := "str"

	// flag1 := strings.HasPrefix(str1, str2)
	// flag2 := strings.HasSuffix(str1, str3)
	// fmt.Printf("%v--%T\n", flag1, flag1)
	// fmt.Printf("%v--%T\n", flag2, flag2)

	// 9、strings.Index(), strings.LastIndex() 子串出现的位置，返回下标位置,Index从前往后，LastIndex从后往前，查不到返回-1
	// str1 := "this is str"
	// str2 := "this"
	// str3 := "str"
	// pzt1 := strings.Index(str1, str2)
	// pzt2 := strings.LastIndex(str1, str3)
	// fmt.Printf("%v--%T\n", pzt1, pzt1)
	// fmt.Printf("%v--%T\n", pzt2, pzt2)
}
