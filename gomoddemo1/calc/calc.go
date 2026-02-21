package calc

import "fmt"

var aaa = "私有变量"
var AAA = "公有变量"

func init() {
	fmt.Println("执行 init 函数")
}
func Add(x, y int) int { // 大写表示公有方法
	return x + y
}

func Sub(x, y int) int {
	return x - y
}
