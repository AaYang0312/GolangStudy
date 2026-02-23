package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func thread(num int) {
	wg.Add(1)
	for i := 0; i < 10; i++ {
		fmt.Printf("协程%v: %v\t", num, i)
		time.Sleep(time.Millisecond * 100)
	}
	wg.Done()
}
func main() {
	for i := 0; i < 10; i++ {
		go thread(i)
	}
	wg.Wait()
	fmt.Println("关闭主线程....")
}
