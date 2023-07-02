package array_a

import (
	"context"
	"fmt"
	"runtime"
	"sort"
	"sync"
	"testing"
	"time"
)

func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)

	// 枚举 a
	for first := 0; first < n; first++ {
		// 需要和上一次枚举的数不相同
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		// c 对应的指针初始指向数组的最右端
		third := n - 1
		target := -1 * nums[first]
		// 枚举 b
		for second := first + 1; second < n; second++ {
			// 需要和上一次枚举的数不相同
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			// 需要保证 b 的指针在 c 的指针的左侧
			for second < third && nums[second]+nums[third] > target {
				third--
			}
			// 如果指针重合，随着 b 后续的增加
			// 就不会有满足 a+b+c=0 并且 b<c 的 c 了，可以退出循环
			if second == third {
				break
			}
			if nums[second]+nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return ans
}

//func tt(nums []int) [][]int {
//	n := len(nums)
//	sort.Ints(nums)
//	ans := make([][]int, 0)
//
//	for first := 0; first < n; first++ {
//		if first > 0 && nums[first] == nums[first-1] {
//			continue
//		}
//		third := n - 1
//		target := -1 * nums[first]
//
//	}
//}

func TestThreeSum(t *testing.T) {

}

//一随机数f(x)以概率 p 生成0,
//那么设g(x)=f(x)>0?0:1;  刚g(x)以概率 1-p 生成0.
//所以f(x),g(x)同时生成0的概率为p(1-p)等于同时生成1的概率.
//得等概率随机数
//function g(x){
//int v=f(x)+g(x);
//if(v==0){
//return 0;  //1.f(x)g(x)同时为0
//else if(v==2){
//return 1;  //2.f(x)g(x)同时为1
//}else{
//g(x);  //3.f(x)g(x)一个为0一个为1,重新生成随机数
//}
//}
//
//上面第3步的概率为p^2+(1-p)^2

//func TestName(t *testing.T) {
//	//for i := 0; i < 10; i++ {
//	//	fmt.Println(generateHalf())
//	//}
//
//	//var i any
//	//i = 1
//	//switch i.(type) {
//	//case int:
//	//	fmt.Println("ture")
//	//}
//
//	ch := make(chan int, 4)
//	//var ch chan int
//
//	//fmt.Printf("%v", ch)
//
//	//close(ch)
//
//	//ch <- 1
//	go func() {
//		ch <- 1
//		close(ch)
//	}()
//	//<-ch
//
//	time.Sleep(time.Second)
//	go func() {
//		for i := 0; i < 10; i++ {
//			sum := <-ch
//			fmt.Println(sum)
//		}
//	}()
//
//	//defer func() {
//	//	fmt.Println("goroutines: ", runtime.NumGoroutine())
//	//}()
//
//	//var ch chan int
//	//go func() {
//	//	ch <- 10
//	//}()
//
//	//time.Sleep(100 * time.Second)
//
//}
//
//func generate(p float64) int {
//	rand.Seed(time.Now().UnixNano())
//	if rand.Float64() < p {
//		return 0
//	} else {
//		return 1
//	}
//}
//
//func generateHalf() int {
//	return generate(0.5) + generate(0.5)
//}

func Test1(t *testing.T) {
	//fmt.Println(test1())
	//fmt.Println(test2())
	//fmt.Println(test3())
	//fmt.Println(test4())

	//var s []int
	s := make([]int, 0, 256)
	for i := 0; i < 513; i++ {
		s = append(s, i)

	}

	fmt.Println(len(s), cap(s))

	newcap := 512
	threshold := 256
	newcap += (newcap + 3*threshold) / 4

	fmt.Println(newcap)

	return
}

func TestMap(t *testing.T) {
	m := make(map[int]string, 10)
	ks := make([]int, 0, 10)
	m[1] = "string1"
	m[2] = "string2"
	m[3] = "string3"
	m[4] = "string4"
	m[5] = "string5"
	m[6] = "string6"
	m[7] = "string7"

	for k := range m {

		ks = append(ks, k)
	}

	fmt.Printf("%v\n", ks)

	sort.Ints(ks)

	for _, k := range ks {
		v, _ := m[k]
		fmt.Println(k, v)
	}
}

func TestChannel(t *testing.T) {

	ch := make(chan func(index int), 1)
	go func() {
		for {
			ch <- task1
			ch <- task2
			ch <- task3
			ch <- task4
		}
	}()

	go func() {
		count := 1
		for {
			select {
			case f := <-ch:
				f(count)
				count++
				if count > 4 {
					count = 1
				}

				time.Sleep(time.Second)
			}
		}
	}()

	time.Sleep(100 * time.Second)
}

func TestC(t *testing.T) {
	ch := make(chan int, 2)
	go func() {
		select {
		case num := <-ch:
			fmt.Println("1 finish ", num)
		}
	}()

	go func() {
		select {
		case num := <-ch:
			fmt.Println("2 finish ", num)
		}
	}()

	go func() {
		select {
		case num := <-ch:
			fmt.Println("3 finish ", num)
		}
	}()

	time.AfterFunc(1*time.Second, func() {
		//close(ch)
		//ch <- 1
		//ch <- 2
		//time.Sleep(1 * time.Second)
		close(ch)
		//ch <- 3
	})

	time.Sleep(100 * time.Second)
}

func TestMapsss(t *testing.T) {
	m := sync.Map{}
	var wg sync.WaitGroup
	wg.Add(30)
	for i := 0; i < 30; i++ {
		go func(in int) {
			defer wg.Done()
			len := 3
			if in >= 20 {
				len = 4
			}
			for j := 0; j < len; j++ {
				if in >= 20 {
					fmt.Printf("%v s\n", (20*3)+(in-20)*4+j)
					m.Store((20*3)+(in-20)*4+j, struct{}{})
				} else {
					fmt.Printf("%v \n", in*3+j)
					m.Store(in*3+j, struct{}{})
				}

			}
		}(i)
	}
	wg.Wait()
	fmt.Println()
	flag := 0
	for i := 0; i < 100; i++ {
		if _, ok := m.Load(i); ok {
			flag++
		}
	}
	fmt.Println("总数：", flag)
}

//func TestMap2(t *testing.T) {
//	var wg sync.WaitGroup
//	var mu sync.Mutex
//	for i := 0; i < 30; i++ {
//		go func(index int) {
//			if index == 0 {
//				for i := 0; i < 70; i++ {
//
//				}
//			}
//		}(i)
//	}
//}

func TestPrint(t *testing.T) {
	defer func() {
		fmt.Println("goroutines: ", runtime.NumGoroutine())
	}()

	ch := make(chan int, 100)
	done := make(chan int)
	for i := 0; i <= 99; i++ {
		ch <- i
	}
	for i := 0; i < 30; i++ {
		go pm2(ch, done, i)
	}
	time.Sleep(5 * time.Second)
}
func pm2(ch chan int, done chan int, index int) {
	for {
		select {
		case num, ok := <-ch:
			fmt.Println("go = ", index, " value = ", num)
			if !ok {
				fmt.Println("finish = ", index)
				return
			}
			//default:
			//	fmt.Println("finish")
			//	break
		}
	}
}

func TestPrint3(t *testing.T) {
	defer func() {
		fmt.Println("goroutines: ", runtime.NumGoroutine())
	}()

	var m sync.Map
	var wg sync.WaitGroup

	ch := make(chan int, 1)
	done := make(chan int)

	for i := 0; i < 30; i++ {
		wg.Add(1)
		go pm3(ch, done, i, &m, &wg)
	}

	ch <- 0
	wg.Wait()

	sum := 0
	for i := 0; i < 100; i++ {
		if _, ok := m.Load(i); ok {
			sum++
		}
	}

	fmt.Println("总数 = ", sum)
}

func pm3(ch chan int, done chan int, index int, m *sync.Map, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case num := <-ch:
			fmt.Println("go = ", index, " value = ", num)
			m.Store(index, struct{}{})
			num++
			if num == 100 {
				fmt.Println("finish = ", index)
				close(done)
				return
			}
			ch <- num
		case <-done:

			return
		}
	}
}

func TestPrint4(t *testing.T) {
	defer func() {
		fmt.Println("goroutines: ", runtime.NumGoroutine())
	}()

	for i := 0; i < 1; i++ {
		print4()
	}

}

func print4() {
	ch := make(chan int, 100)
	var wg sync.WaitGroup
	var m sync.Map
	//ctx, _ := context.WithCancel(context.Background())
	ctx := context.Background()
	for i := 0; i < 100; i++ {
		ch <- i
	}

	for i := 0; i < 30; i++ {
		wg.Add(1)
		go pm4(ctx, ch, i, &m, &wg)
	}

	wg.Wait()

	sum := 0
	for i := 0; i < 100; i++ {
		if _, ok := m.Load(i); ok {
			sum++
		}
	}

	fmt.Println(sum)
}

func pm4(ctx context.Context, ch chan int, index int, m *sync.Map, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ch:
			//fmt.Println("go = ", index, " value = ", v)
			m.Store(index, struct{}{})
		//case v := <-ch:
		//	fmt.Println("go = ", index, " value = ", v)
		//	m.Store(index, struct{}{})
		case <-ctx.Done():
			return
		default:
			if len(ch) == 0 {
				return
			}
		}
	}
}

//func TestPrint5(t *testing.T) {
//	var wg sync.WaitGroup
//	chFuncs := make(chan func(ctx context.Context, ch chan int, index int, wg *sync.WaitGroup), 30)
//	ch := make(chan int, 100)
//	ctx, cancel := context.WithCancel(context.Background())
//	for i := 0; i < 100; i++ {
//		ch <- i
//	}
//	for i := 0; i < 30; i++ {
//		chFuncs <- gpm5
//	}
//
//	for f := range chFuncs {
//		wg.Add(1)
//		go f(ctx, ch, 0, &wg)
//	}
//
//	wg.Wait()
//
//}

func gpm5(ctx context.Context, ch chan int, index int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case v := <-ch:
			fmt.Println("go = ", index, " value = ", v)
		case <-ctx.Done():
			return
		}
	}
}

type Token int

func TestPrint6(t *testing.T) {
	defer func() {
		fmt.Println("goroutines: ", runtime.NumGoroutine())
	}()

	maxNum := 100
	chs := [30]chan Token{}
	done := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())
	//defer cancel()

	for i := 0; i < 30; i++ {
		chs[i] = make(chan Token)
	}

	for i := 0; i < 30; i++ {
		go gpm6(ctx, chs[i], i, chs[(i+1)%30], done, maxNum)
	}

	chs[0] <- 0

	select {
	case <-done:
		fmt.Println("finish====================")
		cancel()
	}

	time.Sleep(1 * time.Second)

}

func gpm6(ctx context.Context, ch chan Token, index int, nextCh chan Token, done chan struct{}, maxNum int) {
	for {
		select {
		case v := <-ch:
			fmt.Println("go = ", index, " value = ", v)
			v++
			if int(v) == maxNum {
				done <- struct{}{}
				return
			}
			nextCh <- v
		case <-ctx.Done():
			fmt.Println("done====================")
			return
		}
	}
}

var a, b int

func TestCache(t *testing.T) {

	for i := 0; i < 100; i++ {
		go f() //g1
		g()    //g2
	}
}

func f() {
	a = 1 // w之前的写操作
	b = 2 // 写操作w
}

func g() {
	print(b)  // 读操作r
	print(a)  // ???
	println() // ???
}

func main() {

}

//func pm(ch chan int, index int) {
//	for {
//		num := <-ch
//		fmt.Println("go = ", index, " value = ", num)
//		//select {
//		//case num := <-ch:
//		//	fmt.Println("go = ", index, " value = ", num)
//		//	if num >= 99 {
//		//		fmt.Println("finish = ", index)
//		//		close(done)
//		//	}
//		//case <-done:
//		//	fmt.Println("end = ", index)
//		//}
//	}
//}

func task1(index int) {
	fmt.Println(index)
}
func task2(index int) {
	fmt.Println(index)
}
func task3(index int) {
	fmt.Println(index)
}
func task4(index int) {
	fmt.Println(index)
}

func test1() (v int) {
	defer fmt.Println(v)
	return v
}

func test2() (v int) {
	defer func() {
		fmt.Println(v)
	}()
	return 3
}

func test3() (v int) {
	defer fmt.Println(v)
	v = 3
	return 4
}

func test4() (v int) {
	defer func(n int) {
		fmt.Println(n)
	}(v)
	return 5
}
