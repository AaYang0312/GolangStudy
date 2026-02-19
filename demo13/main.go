package main

func main() {
	// 1、break语句：跳出循环，在switch语句中执行case后跳出，在多重循环中，可用标号label标出想break的循环
	// for i := 0; i <= 9; i++ {
	// 	fmt.Println(i)
	// 	if i == 5 {
	// 		break
	// 	}
	// }
	// fmt.Println("继续执行")

	// 跳出当前循环
	// for i := 0; i < 10; i++ {
	// 	for j := 0; j < 10; j++ {

	// 		if j == 3 {
	// 			break
	// 		}
	// 		fmt.Printf("i=%v, j=%v\n", i, j)
	// 	}
	// }

	// label

	// for i := 0; i < 2; i++ {
	// label1:
	// 	for j := 0; j < 10; j++ {
	// 		if j == 3 {
	// 			break label1
	// 		}
	// 		fmt.Printf("i=%v, j=%v\n", i, j)
	// 	}
	// }
	// fmt.Printf("继续执行")

	// 2、continue 语句可以跳出当前循环，开始下一次循环，仅限for语句内
	// for i := 0; i < 2; i++ {
	// 	if i == 0 {
	// 		continue
	// 	}
	// 	for j := 0; j < 10; j++ {
	// 		fmt.Printf("i=%v, j=%v\n", i, j)
	// 	}
	// }
	// fmt.Printf("继续执行")

	// 可以和标签一起使用
	// label1:
	// 	for i := 0; i < 2; i++ {
	// 		for j := 0; j < 10; j++ {
	// 			if j == 3 {
	// 				continue label1
	// 			}
	// 			fmt.Printf("i=%v, j=%v\n", i, j)
	// 		}
	// 	}
	// 	fmt.Printf("继续执行")

	// 3、goto 语句通过标签进行代码间无条件跳转，可以跳出循环，避免重复退出
	// 	var i int = 3
	// 	if i < 4 {
	// 		fmt.Println("比四小")
	// 		goto label2
	// 	}

	// 	fmt.Println("aaa")
	// 	fmt.Println("bbb")
	// label2:
	// 	fmt.Println("ccc")
	// 	fmt.Println("ddd")
}
