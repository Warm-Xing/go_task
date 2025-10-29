package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	// 启动奇数打印协程
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 100; i += 2 {
			fmt.Printf("奇数: %d\n", i)
			time.Sleep(10 * time.Millisecond)
		}
	}()

	// 启动偶数打印协程
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 2; i <= 100; i += 2 {
			fmt.Printf("偶数: %d\n", i)
			time.Sleep(10 * time.Millisecond)
		}
	}()

	// 等待所有协程完成
	wg.Wait()
}
