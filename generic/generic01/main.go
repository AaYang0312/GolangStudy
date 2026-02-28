package main

import "fmt"

func getValue(value interface{}) interface{} {
	return value
}
func main() {
	str := "123"
	resultTemp := getValue(str)
	result, _ := resultTemp.(string)
	fmt.Printf("类型: %T, 长度: %d,值: %v\n", result, len(result), result) // 空接口类型不能直接获取长度，需要先类型断言
}
