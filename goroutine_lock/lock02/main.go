package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var mutex sync.RWMutex

// 读锁和写锁
func write() {
	defer wg.Done()
	mutex.Lock()
	fmt.Println("执行写操作")
	time.Sleep(time.Second * 2)
	mutex.Unlock()

}
func read() {
	defer wg.Done()
	mutex.RLock()
	fmt.Println("执行读操作")
	time.Sleep(time.Second * 2)
	mutex.RUnlock()
}

/*
读锁（RLock/RUnlock）
允许多个 goroutine 同时持有读锁。
不允许与写锁同时存在。
适用于只读操作，多个 goroutine 可以并发读取共享资源，但不能修改。
特点：
多个 goroutine 可以同时获取读锁。
如果有 goroutine 持有写锁，则其他 goroutine 无法获取读锁（必须等待写锁释放）。
读锁之间是共享的，不会互相阻塞。
*/
/*
写锁（Lock/Unlock）
同一时间只能有一个 goroutine 持有写锁。
不允许与其他读锁或写锁同时存在。
适用于写操作，确保对共享资源的独占访问。
特点：
写锁是排他的，只有一个 goroutine 可以持有写锁。
如果有 goroutine 持有读锁或写锁，则其他 goroutine 无法获取写锁（必须等待所有锁释放）。
写锁会阻塞所有其他读锁和写锁的获取。
*/
func main() {
	// 10个读10个写
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go read()
	}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}
	wg.Wait()
}
