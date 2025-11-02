package main

import (
	"fmt"
	"sync"
)

func producer(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done() //任务完成时通知WaitGroup
	defer close(ch) // 函数退出时关闭通道

	for i := 1; i <= 10; i++ {
		ch <- i
	}
	fmt.Println("生产者完成")
}

func consumer(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for num := range ch {
		fmt.Printf("消费者接收: %d\n", num)
	}
	fmt.Println("消费者完成")
}

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup

	wg.Add(2)

	go producer(ch, &wg)
	go consumer(ch, &wg)

	wg.Wait()
	fmt.Println("程序结束")
}
