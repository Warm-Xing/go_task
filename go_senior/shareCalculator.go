package main

import (
	"fmt"
	"sync"
)

type SafeCounter struct {
	mu    sync.Mutex
	value int
}

func (c *SafeCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *SafeCounter) Value() int {
	c.mu.Lock()         // 1. 获取锁，同时读取值，
	defer c.mu.Unlock() // 2. 注册解锁操作（函数返回前执行）2. 注册解锁操作，函数退出前自动执行
	return c.value      // 3. 读取值并返回（此时锁仍持有）
}

func worker(counter *SafeCounter, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 1000; i++ {
		counter.Increment()
	}
}

func main() {
	var wg sync.WaitGroup
	counter := &SafeCounter{}

	// 启动10个协程
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(counter, &wg)
	}

	wg.Wait()
	fmt.Printf("最终计数器值: %d\n", counter.Value()) // 输出: 最终计数器值: 10000
}
