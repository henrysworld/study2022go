package sync

import (
	"fmt"
	"math/rand"
	_ "net/http/pprof"
	"sync"
	"testing"
	"time"
)

func Test(t *testing.T) {
	// 2^N
	num := 5
	if num&(num-1) != 0 {
		for num&(num-1) != 0 {
			num &= num - 1
		}
		num <<= 1
	}
	//data := make([]int, num)
	fmt.Println(num)
}

func handle(v int) {
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < v; i++ {
		fmt.Println("蠢哭了")
		wg.Done()
	}
	wg.Wait()
}

func queryAll() int {
	ch := make(chan int)
	for i := 0; i < 3; i++ {
		go func(i int) {
			ch <- i
		}(i)
	}

	sum := 0
	//for vi2 := range ch {
	//	sum += vi2
	//}
	//select {
	//case i := <-ch:
	//	sum += i
	//}
	return sum
}

func query() int {
	n := rand.Intn(100)
	time.Sleep(time.Duration(n) * time.Millisecond)
	return n
}

func Test3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	panic("a")
	//var m = make(map[int]int, 10) // 初始化一个map
	//go func() {
	//	for {
	//		m[1] = 1 //设置key
	//	}
	//}()
	//
	//go func() {
	//	for {
	//		_ = m[2] //访问这个map
	//	}
	//}()
	//select {}
}
