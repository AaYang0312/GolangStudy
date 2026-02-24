package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func WriteChan(ch chan int) {
	defer wg.Done()
	for i := 1; i <= 100; i++ {
		ch <- i
		fmt.Printf("[写入]数据%v成功\n", i)
		time.Sleep(time.Millisecond * 500)
	}
	close(ch)
}

func ReadChan(ch chan int) {
	defer wg.Done()
	for v := range ch {
		time.Sleep(time.Millisecond * 50)
		fmt.Printf("[读取]数据%v成功\n", v)
	}
}
func main() {
	var ch = make(chan int, 100)

	wg.Add(1)
	go WriteChan(ch)
	wg.Add(1)
	go ReadChan(ch)
	wg.Wait()
}
