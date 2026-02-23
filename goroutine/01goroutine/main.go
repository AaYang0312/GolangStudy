package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("test() golang-", i)
		time.Sleep(time.Millisecond * 100)
	}
	wg.Done() // 协程计数器减一
}
func main() {
	wg.Add(1) // 协程计数器加一
	go test() // 创建一个协程
	fmt.Println("可使用的CPU个数为:", runtime.NumCPU())
	// runtime.GOMAXPROCS(-1) 设置最大可用cpu数
	for i := 0; i < 10; i++ {
		fmt.Println("main() golang-", i)
		time.Sleep(time.Millisecond * 50)
	}
	wg.Wait() // 等待所有协程执行完毕
}
