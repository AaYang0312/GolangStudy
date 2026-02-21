package main

import (
	"fmt"
	"gomoddemo1/calc"
)

func main() {
	sum := calc.Add(1, 2)
	fmt.Println(sum)
	//fmt.Println(calc.aaa) // 首字母小写无法访问
	fmt.Println(calc.AAA)
}
