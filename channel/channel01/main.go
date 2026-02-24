package main

import "fmt"

func main() {
	// ch := make(chan int, 3)
	// ch <- 10
	// ch <- 32
	// ch <- 44
	// fmt.Println(<-ch)
	// // 管道的长度和容量
	// fmt.Printf("值: %v, 容量: %v, 长度: %v", ch, cap(ch), len(ch)) // 值: 0x4812f9d4000, 容量: 3, 长度: 2

	// ch1 := make(chan int, 4)
	// ch1 <- 10
	// ch1 <- 20
	// ch1 <- 30
	// ch2 := ch1
	// ch2 <- 69
	// <-ch1
	// <-ch1
	// <-ch1
	// d := <-ch1
	// fmt.Println(d) // 69

	// 循环遍历管道数据，注意管道没有key
	var ch3 = make(chan int, 5)
	ch3 <- 1
	ch3 <- 2
	ch3 <- 3
	ch3 <- 4
	ch3 <- 5
	close(ch3) // 关闭一个管道
	for val := range ch3 {
		fmt.Printf("%v ", val)
	}
	// 通过for循环遍历则不需要关闭管道
}
