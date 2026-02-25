package main

import (
	"fmt"
	"time"
)

func main() {
	// 有时我们需要同时从多个管道获取数据，这个时候可以使用golang提供的select语句

	// 1. 定义一个 int 管道
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}

	// 2. 定义一个 string 管道有5个数据
	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- fmt.Sprintf("person%d", i)
	}

	// select 需要一个 for 死循环
	for {
		// 每次循环随机选取一个管道取数据
		select {
		case v := <-intChan:
			fmt.Printf("从 intChan 读取的数据%d\n", v)
			time.Sleep(time.Millisecond * 50)
		case v := <-stringChan:
			fmt.Printf("从 stringChan 读取的数据%v\n", v)
			time.Sleep(time.Millisecond * 50)
		default:
			fmt.Printf("数据获取完毕")
			return // 注意退出
		}
	}
}
