package main

import (
	"fmt"
	"sync"
	"time"
)

// 定义函数
type Task func()

type TaskScheduler struct {
	tasks []Task
	wg    sync.WaitGroup
}

// 本函数没有完全看懂，期待下次再看
// 使用的方法就是结构体
func NewTaskScheduler() *TaskScheduler {
	return &TaskScheduler{
		tasks: make([]Task, 0),
	}
}

func (ts *TaskScheduler) AddTask(task Task) {
	ts.tasks = append(ts.tasks, task)
}

func (ts *TaskScheduler) Run() {
	startTime := time.Now()
	fmt.Printf("开始执行 %d 个任务\n", len(ts.tasks))

	ts.wg.Add(len(ts.tasks))

	for i, task := range ts.tasks {
		go func(taskIndex int, t Task) {
			defer ts.wg.Done()

			taskStart := time.Now()
			//执行函数(func() {
			//		//time.Sleep(1 * time.Second)
			//		fmt.Println("任务1完成")
			//	}
			t()
			duration := time.Since(taskStart)

			fmt.Printf("任务 %d 执行时间: %v\n", taskIndex+1, duration)
		}(i, task)
	}

	ts.wg.Wait()
	totalDuration := time.Since(startTime)
	fmt.Printf("所有任务执行完毕，总耗时: %v\n", totalDuration)
}

func main() {
	scheduler := NewTaskScheduler()

	// 添加示例任务
	scheduler.AddTask(func() {
		//time.Sleep(1 * time.Second)
		fmt.Println("任务1完成")
	})

	scheduler.AddTask(func() {
		//time.Sleep(2 * time.Second)
		fmt.Println("任务2完成")
	})

	scheduler.AddTask(func() {
		//time.Sleep(500 * time.Millisecond)
		fmt.Println("任务3完成")
	})

	scheduler.Run()
}
