package main

func main() {
	// 1、打印：0-50 所有偶数
	// 2、求 1 + 2 + 3 + ... + 100 的和
	// 3、打印 1~100 之间所有是 9 的倍数的整数的个数及总和
	// 4、计算 5 的阶乘（ 12345 n 的阶乘 12......n）
	// 5、打印一个矩形

	// 1

	// for i := 0; i <= 50; i++ {
	// 	if (i != 0) && (i%2 == 0) {
	// 		fmt.Println(i)
	// 	}
	// }

	// 2
	// var num = 1
	// var sum = 0
	// for num = 1; num < 101; num++ {
	// 	sum += num
	// }
	// fmt.Println(sum)

	// 3
	// var num = 1
	// var count = 0
	// var sum = 0
	// for num = 1; num <= 100; num++ {
	// 	if num%9 == 0 {
	// 		count++
	// 		sum += num
	// 	}
	// }
	// fmt.Println(count)
	// fmt.Println(sum)

	// 4
	// var sum = 1
	// for i := 1; i <= 5; i++ {
	// 	sum *= i
	// }
	// fmt.Println(sum)

	// 5
	// for i := 0; i < 10; i++ {
	// 	for j := 0; j <= 10; j++ {
	// 		fmt.Printf("#")
	// 	}
	// 	fmt.Printf("\n")
	// }

	// 6、打印一个三角形
	// var i = 5
	// var j = 0
	// for i = 1; i <= 5; i++ {
	// 	for j = 1; j <= i; j++ {
	// 		fmt.Printf("*")
	// 	}
	// 	fmt.Printf("\n")
	// }

	// 7、打印九九乘法表
	// var row = 9
	// var col = 9
	// for row = 1; row <= 9; row++ {
	// 	for col = 1; col <= row; col++ {
	// 		fmt.Printf("%dx%d=%d ", row, col, row*col)
	// 	}
	// 	fmt.Printf("\n")
	// }
}
