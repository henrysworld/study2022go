package task2

import (
	"fmt"
	"sort"
	"sync"
	"testing"
)

func TestName(t *testing.T) {

}

type task struct {
	Position int // 任务在原始列表中的位置
	Index    int64
	Url      string
	times    []string
	temps    []string
}

var processedTasksFSDA []task
var processedTasksFSDB []task
var mu sync.Mutex

func TestTask1(t *testing.T) {
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
	addTasksToChannel("FSDA", tasksMap["FSDA"], 0)
	addTasksToChannel("FSDB", tasksMap["FSDB"], 5)

	// 等待所有worker完成
	for i := 0; i < 2; i++ {
		<-done
	}

	// 对每组任务单独排序
	sortTasks(processedTasksFSDA)
	sortTasks(processedTasksFSDB)

	fmt.Println("Processed tasks for FSDA:", processedTasksFSDA)
	fmt.Println("Processed tasks for FSDB:", processedTasksFSDB)
}

func addTasksToChannel(key string, tasksChan chan task, startPos int) {
	for i := 0; i < 5; i++ {
		tasksChan <- task{Position: startPos + i, Index: int64(i), Url: fmt.Sprintf("http://example.com/%s/%d", key, i)}
	}
	close(tasksChan)
}

func worker(key string, tasksChan chan task, done chan bool) {
	subtasks := make(chan task, 5) // subtask队列

	// 启动五个subworker
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go subworker(key, subtasks, &wg)
	}

	for t := range tasksChan {
		subtasks <- t
	}
	close(subtasks)

	wg.Wait() // 等待所有subworker完成
	done <- true
}

func subworker(group string, subtasks chan task, wg *sync.WaitGroup) {
	defer wg.Done()
	for t := range subtasks {
		// 这里处理任务
		t.Url = t.Url + "-processed"
		mu.Lock()
		if group == "FSDA" {
			processedTasksFSDA = append(processedTasksFSDA, t)
		} else if group == "FSDB" {
			processedTasksFSDB = append(processedTasksFSDB, t)
		}
		mu.Unlock()
	}
}

func sortTasks(taskSlice []task) {
	sort.Slice(taskSlice, func(i, j int) bool {
		return taskSlice[i].Position < taskSlice[j].Position
	})
}

func TestInterpolation(t *testing.T) {
	x0, y0 := 2.8, 8.23 // Starting point
	x1, y1 := 3.5, 10.0 // Ending point
	x := 2.7            // The point where we want to interpolate

	y := LinearInterpolation(x0, y0, x1, y1, x)
	fmt.Printf("The interpolated y value at x=%.2f is: %.2f\n", x, y)
}

func LinearInterpolation(x0, y0, x1, y1, x float64) float64 {
	if x0 == x1 {
		return (y0 + y1) / 2 // Avoid division by zero and return the average of y0 and y1
	}
	return y0 + (x-x0)*(y1-y0)/(x1-x0)
}
