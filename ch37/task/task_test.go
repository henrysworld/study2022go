package task

import (
	"fmt"
	"sync"
	"testing"
)

type task struct {
	Index int64
	Url   string
	times []string
	temps []string
}

func TestName(t *testing.T) {

}

func TestTask(t *testing.T) {
	tasksMap := map[string]chan task{
		"FSDA": make(chan task, 10),
		"FSDB": make(chan task, 10),
	}

	done := make(chan bool, 2) // 用于通知主协程所有任务组完成

	// 启动worker
	for key, tasksChan := range tasksMap {
		go worker(key, tasksChan, done)
	}

	// 添加任务
	addTasksToChannel("FSDA", tasksMap["FSDA"])
	addTasksToChannel("FSDB", tasksMap["FSDB"])

	// 等待所有worker完成
	for i := 0; i < 2; i++ {
		<-done
	}
}

func addTasksToChannel(key string, tasksChan chan task) {
	for i := int64(0); i < 5; i++ {
		tasksChan <- task{Index: i, Url: fmt.Sprintf("http://example.com/%s/%d", key, i)}
	}
	close(tasksChan)
}

func worker(key string, tasksChan chan task, done chan bool) {
	subtasks := make(chan task, 5) // subtask队列

	// 启动五个subworker
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go subworker(subtasks, &wg)
	}

	for t := range tasksChan {
		subtasks <- t
	}
	close(subtasks)

	wg.Wait() // 等待所有subworker完成
	done <- true
}

func subworker(subtasks chan task, wg *sync.WaitGroup) {
	defer wg.Done()
	for t := range subtasks {
		// 这里处理任务
		fmt.Printf("Processing task: %v\n", t)
	}
}
