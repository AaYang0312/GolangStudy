package main

import "fmt"

// type calc func(int, int) int // 表示定义一个calc的类型
// type myInt int

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

// type calcType func(int, int) int

// func calc(x, y int, cb calcType) int { // 调用时把函数赋值给cb，如cb = add()
// 	fmt.Printf("%T", cb)
// 	return cb(x, y)
// }

// 定义一个方法类型
//type calcType func(int, int) int

//封装匿名方法
func do(o string) func(int, int) int {
	switch o {
	case "+":
		return add
	case "-":
		return sub
	case "*":
		return func(x, y int) int {
			return x * y
		}
	default:
		return nil
	}
}

func main() {
	// fmt.Println(calc(10, 5, add))
	// 匿名函数
	// j := calc(3, 4, func(x, y int) int {
	// 	return x * y
	// })
	var a = do("+")
	var s = do("-")
	var f = do("*")
	fmt.Println(a(12, 4))
	fmt.Println(s(12, 4))
	fmt.Println(f(12, 4))
}

// func main() {
// 	var c calc
// 	c = add
// 	f := sub
// 	fmt.Printf("c-type: %T, f-type: %T\n", c, f) //c-type: main.calc, f-type: func(int, int) int

// 	var a int = 10
// 	var b myInt = 20
// 	fmt.Printf("a-type: %T, b-type: %T\n", a, b) //a-type: int, b-type: main.myInt
// 	// int 和 myInt 不能一起运算
// 	fmt.Println(a + int(b)) // 需要类型转换
// }

// func main() {
// 	// 匿名自执行函数
// 	func() {
// 		fmt.Println("test....")
// 	}()
// 	// 匿名函数
// 	var fn = func(x, y int) int {
// 		return x + y
// 	}
// 	fmt.Println(fn(3, 4))
// 	// 匿名自执行函数接收参数
// 	func(x, y int) {
// 		fmt.Println(x + y)
// 	}(10, 30)
// }
