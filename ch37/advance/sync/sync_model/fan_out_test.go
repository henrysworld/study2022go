package sync_model

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

//func TestFanOut(t *testing.T) {
//	var wg sync.WaitGroup
//	wg.Add(36)
//	go pool(&wg, 36, 50)
//	wg.Wait()
//}
//
//func worker(tasksCh <-chan int, wg *sync.WaitGroup) {
//	defer wg.Done()
//	for {
//		task, ok := <-tasksCh
//		if !ok {
//			return
//		}
//
//		d := time.Duration(task) * time.Second
//		time.Sleep(d)
//		fmt.Println("processing task", task, "duration", d)
//	}
//}
//
//func pool(wg *sync.WaitGroup, workers, tasks int) {
//	tasksCh := make(chan int)
//
//	for i := 0; i < workers; i++ {
//		go worker(tasksCh, wg)
//	}
//
//	for i := 0; i < tasks; i++ {
//		tasksCh <- i
//	}
//
//	close(tasksCh)
//}
//
//const (
//	WORKERS    = 5
//	SUBWORKERS = 3
//	TASKS      = 20
//	SUBTASKS   = 10
//)
//
//func TestFanOutSub(t *testing.T) {
//	var wg sync.WaitGroup
//	wg.Add(WORKERS)
//	tasks := make(chan int)
//
//	for i := 0; i < WORKERS; i++ {
//		go workerSub(tasks, &wg)
//	}
//
//	for i := 0; i < TASKS; i++ {
//		tasks <- i
//	}
//
//	close(tasks)
//	wg.Wait()
//}
//
//func subworker(subtasks chan int) {
//	for {
//		task, ok := <-subtasks
//		if !ok {
//			return
//		}
//		time.Sleep(time.Duration(task) * time.Millisecond)
//		fmt.Println(task)
//	}
//}
//
//func workerSub(tasks <-chan int, wg *sync.WaitGroup) {
//	defer wg.Done()
//	for {
//		task, ok := <-tasks
//		if !ok {
//			return
//		}
//
//		subtasks := make(chan int, 0)
//		for i := 0; i < SUBWORKERS; i++ {
//			go subworker(subtasks)
//		}
//
//		for i := 0; i < SUBTASKS; i++ {
//			task1 := task + i
//			subtasks <- task1
//		}
//
//		close(subtasks)
//	}
//}

const (
	WORKERS    = 5
	SUBWORKERS = 3
	TASKS      = 20
	SUBTASKS   = 10
)

func subworker(subtasks chan int) {
	for {
		task, ok := <-subtasks
		if !ok {
			return
		}
		time.Sleep(time.Duration(task) * time.Millisecond * 100)
		fmt.Println(task)
	}
}

func worker(tasks <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		task, ok := <-tasks
		if !ok {
			return
		}

		subtasks := make(chan int)
		for i := 0; i < SUBWORKERS; i++ {
			go subworker(subtasks)
		}
		for i := 0; i < SUBTASKS; i++ {
			task1 := task * i
			subtasks <- task1
		}
		close(subtasks)
	}
}

func TestFanOut(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(WORKERS)
	tasks := make(chan int)

	for i := 0; i < WORKERS; i++ {
		go worker(tasks, &wg)
	}

	for i := 0; i < TASKS; i++ {
		tasks <- i
	}

	close(tasks)
	wg.Wait()
}
