package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type AtomicCounter struct {
	value int64
}

func (c *AtomicCounter) Increment() {
	atomic.AddInt64(&c.value, 1) //使用atomic.AddInt64原子地将计数器加1
}

func (c *AtomicCounter) Value() int64 {
	return atomic.LoadInt64(&c.value) //使用atomic.LoadInt64原子地读取计数器当前值
}

func worker(counter *AtomicCounter, wg *sync.WaitGroup) {
	defer wg.Done() // 函数跑完之后就开始注册减1

	for i := 0; i < 1000; i++ {
		counter.Increment()
	}
}

func main() {
	var wg sync.WaitGroup
	counter := &AtomicCounter{}

	// 启动10个协程
	for i := 0; i < 10; i++ {
		wg.Add(1) //注册加1
		go worker(counter, &wg)
	}

	wg.Wait()
	fmt.Printf("最终计数器值: %d\n", counter.Value()) // 输出: 最终计数器值: 10000
}
