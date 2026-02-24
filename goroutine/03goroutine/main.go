package main

import (
	"fmt"
	"sync"
)

// 练习：统计1-120000以内的素数，for循环实现

// func main() {
// 	start := time.Now().Unix()
// 	for num := 2; num < 120000; num++ {
// 		var flag = true
// 		for i := 2; i < num; i++ {
// 			if num%i == 0 {
// 				flag = false
// 				break
// 			}
// 		}
// 		if flag == true {
// 			fmt.Printf("%d 是素数\n", num)
// 		}
// 	}
// 	end := time.Now().Unix()
// 	fmt.Println(end - start)
// }

// 协程实现
var wg sync.WaitGroup

func sushu(start int, end int) {
	defer wg.Done()
	for num := start; num <= end; num++ {
		if num < 2 {
			continue
		}
		var flag = true
		for i := 2; i*i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag == true {
			fmt.Printf("%d是素数\n", num)
		}
	}
}
func main() {
	for n := 0; n < 4; n++ {
		wg.Add(1)
		go sushu(n*30+1, (n+1)*30)
	}
	wg.Wait()
	fmt.Println("执行完毕")
}
