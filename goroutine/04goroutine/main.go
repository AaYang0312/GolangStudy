package main

import (
	"fmt"
	"sort"
	"sync"
)

// 获取 1-120000 之间的素数

var wg sync.WaitGroup

// 向 intChan 放入 1-120000 的数
func putNum(intChan chan int) {
	defer wg.Done()
	for i := 2; i < 120000; i++ {
		intChan <- i
	}
	close(intChan)
}

// 从 intChan 取出数据，并判断是否为素数，如果是，就放入primeChan
func putPrime(intChan chan int, primeChan chan int, exitChan chan bool) {
	defer wg.Done()
	for num := range intChan {
		prime := true
		if num%2 == 0 && num != 2 {
			prime = false
		} else {
			for i := 2; i*i < num; i++ {
				if num%i == 0 {
					prime = false
					break
				}
			}
		}
		if prime == true {
			primeChan <- num
		}
	}
	exitChan <- true
}
func printPrime(primeChan chan int) {
	defer wg.Done()
	var primes []int
	for v := range primeChan {
		// fmt.Println(v) // 打印素数
		primes = append(primes, v)
	}
	sort.Ints(primes)
	for _, p := range primes {
		fmt.Printf("%d\t", p)
	}
}
func main() {
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 1000)
	exitChan := make(chan bool, 16) // 标识符

	// 存放数字的协程
	wg.Add(1)
	go putNum(intChan)

	// 存放素数的协程
	for i := 0; i < 16; i++ {
		wg.Add(1)
		go putPrime(intChan, primeChan, exitChan)
	}

	// 根据 exitChan 判断是否关闭 primeChan，原理是读取协程在有协程写入管道时会等待
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 16; i++ {
			<-exitChan
		}
		close(primeChan)
	}()

	// 打印素数的协程
	wg.Add(1)
	go printPrime(primeChan)

	wg.Wait()
	fmt.Println("主线程执行完毕")
}
