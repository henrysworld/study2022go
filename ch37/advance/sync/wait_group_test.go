package sync

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"

	"golang.org/x/sync/errgroup"
)

func TestWaitGroup(t *testing.T) {
	wg := sync.WaitGroup{}

	var result int64 = 0
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(index int) {
			atomic.AddInt64(&result, int64(index))
			fmt.Println(index)
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println(result)
}

func TestErrorGroup(t *testing.T) {
	eg := errgroup.Group{}
	var result int64 = 0
	for i := 0; i < 10; i++ {
		delta := i
		eg.Go(func() error {
			fmt.Println(delta)
			atomic.AddInt64(&result, int64(delta))
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		t.Fatal(err)
	}

	fmt.Println(result)
}
