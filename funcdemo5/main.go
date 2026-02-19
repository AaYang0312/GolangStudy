package main

import (
	"context"
	"fmt"
	"time"
)

// 传统 Ticker 用法
func traditionalTickerDemo() {
	fmt.Println("传统Ticker用法")
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop() // 记得停止
	count := 0
	for range ticker.C {
		count++
		fmt.Printf("传统Ticker触发, #%d", count)
		if count >= 3 {
			break
		}
	}
	fmt.Println("传统Ticker完成")
}
func modernContextTickerDemo() {
	fmt.Println("\n--- 2. Context控制的Ticker ---")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	ticker := time.NewTicker(300 * time.Millisecond)
	defer ticker.Stop()

	count := 0
	for {
		select {
		case <-ticker.C:
			count++
			fmt.Printf("Context Ticker触发 #%d\n", count)
		case <-ctx.Done():
			fmt.Printf("Context结束: %v\n", ctx.Err())
			return
		}
	}
}

// 场景3：速率限制
func scenario3_rateLimiting() {
	fmt.Println("\n场景3：API调用速率限制")

	// 限制每秒最多3个请求
	rateLimiter := time.NewTicker(333 * time.Millisecond) // ~3次/秒
	defer rateLimiter.Stop()

	requests := []string{"req1", "req2", "req3", "req4", "req5"}

	for i, req := range requests {
		<-rateLimiter.C // 等待令牌
		fmt.Printf("处理请求 %s (#%d)\n", req, i+1)
	}

	fmt.Println("所有请求处理完成")
}
func main() {
	// 字符串转换时间戳
	// str := "2026-2-16 15:17:00"
	// layout := "2006-1-2 15:04:05"
	// t, _ := time.Parse(layout, str)
	// timeStamp := t.Unix()
	// fmt.Println(timeStamp)

	// // 一些常量
	// fmt.Println(time.Nanosecond)
	// fmt.Println(time.Second)
	// fmt.Println(time.Millisecond)
	// fmt.Println(time.Minute)
	// fmt.Println(time.Hour)

	// 时间运算
	// now := time.Now()
	// later1 := now.Add(time.Hour) // 一小时后
	// later2 := later1.Sub(now)    // 时间差
	// equal := now.Equal(later1)   // 是否相同时间
	// before := later1.Before(now) // 是否早于
	// fmt.Println(now)
	// fmt.Println(later1)
	// fmt.Println(later2)
	// fmt.Println(equal)
	// fmt.Println(before)

	// 定时器
	// func() {
	// 	fmt.Println("任务开始执行")
	// 	time.Sleep(2 * time.Second)
	// 	fmt.Println("定时器任务结束")
	// }()

	// func() {
	// 	fmt.Println("任务开始执行")
	// 	<-time.After(2 * time.Second)
	// 	fmt.Println("定时器任务结束")
	// }()

	// fmt.Println("创建NewTimer定时器")
	// timer1 := time.NewTimer(2 * time.Second)
	// func() {
	// 	<-timer1.C
	// 	fmt.Println("NewTimer定时器任务完成")
	// }()

	// fmt.Println("创建NewTimer定时器")
	// timer1 := time.NewTimer(6 * time.Second)
	// func() {
	// 	select {
	// 	case <-timer1.C:
	// 		fmt.Println("NewTimer定时器任务完成")
	// 	case <-time.After(4 * time.Second):
	// 		fmt.Println("超时退出")
	// 	}
	// }()
	// time.Sleep(100 * time.Millisecond)
	// timer1.Stop()
	// fmt.Println("定时器已停止")

	// Ticker
	//traditionalTickerDemo()
	//modernContextTickerDemo()
	scenario3_rateLimiting()
}
