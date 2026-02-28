package main

import "fmt"

type Number interface {
	int | float64
}

// 定义泛型方法
func Add[T Number](a, b T) T {
	return a + b
}

func main() {
	fmt.Printf("%d + %d = %d\n", 1, 2, Add(1, 2))
	fmt.Println("2 + 2.5 = ", Add(2, 2.5))
}
