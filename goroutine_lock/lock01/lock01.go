package main

import (
	"fmt"
	"sync"
	"time"
)

// 互斥锁
var count = 0
var wg sync.WaitGroup
var mutex sync.Mutex

func test() {
	defer wg.Done()
	mutex.Lock()
	count++
	fmt.Println("the count is: ", count)
	time.Sleep(time.Millisecond)
	mutex.Unlock()
}
func main() {
	// 可能同时操作 count
	for r := 0; r < 200; r++ {
		wg.Add(1)
		go test()
	}
	wg.Wait()
}
