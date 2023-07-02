package ctx

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

type key string

func TestContext(t *testing.T) {
	ctx := context.Background()
	timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	key1 := key("key1")
	vCtx := context.WithValue(timeoutCtx, key1, "value1")
	value := vCtx.Value(key1)
	fmt.Println(value)
	dCtx, dcancel := context.WithDeadline(vCtx, time.Now().Add(time.Second))
	defer dcancel()

	time.Sleep(2 * time.Second)

	err := timeoutCtx.Err()
	derr := dCtx.Err()
	fmt.Println("t = ", err)
	fmt.Println("d = ", derr)

}

func TestContextValue(t *testing.T) {
	parentkey := key("parent")
	ctx := context.Background()
	parent := context.WithValue(ctx, parentkey, "parent")
	subkey := key("subkey")
	sub := context.WithValue(parent, subkey, "sub")

	err := sub.Err()
	if err != nil {
		fmt.Println(err)
	}

	valps := parent.Value(subkey)
	valp := parent.Value(parentkey)
	valsp := sub.Value(parentkey)
	vals := sub.Value(subkey)
	fmt.Println(valps)
	fmt.Println(valp)
	fmt.Println(valsp)
	fmt.Println(vals)
}

// func TestErrgroup(t *testing.T) {
// 	eg, ctx := errgroup.WithContext(context.Background())

// 	var result int64 = 0
// 	for i := 0; i < 10; i++ {
// 		delta := i
// 		eg.Go(func() error {
// 			atomic.AddInt64(&result, int64(delta))
// 			return nil
// 		})
// 	}

// 	if err := eg.Wait(); err != nil {
// 		t.Fatal(err)
// 	}
// 	ctx.Err()
// 	fmt.Println(result)

// }

func TestBusinessTimeout(t *testing.T) {
	ctx := context.Background()
	timeoutCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	end := make(chan struct{}, 1)
	ret := make(chan struct{}, 1)
	go func() {
		select {
		case <-ret:
			return
		}
		MyBusiness()
		end <- struct{}{}
	}()

	ch := timeoutCtx.Done()

	select {
	case <-ch:
		fmt.Println("Timeout")
	case <-end:
		fmt.Println("Business end")
	}
}

func MyBusiness() {
	time.Sleep(5 * time.Millisecond)
	fmt.Println("MyBusiness")
}

func TestParentValueCtx(t *testing.T) {
	ctx := context.Background()
	pKey := key("parent")
	sKey := key("sub")
	parnetCtx := context.WithValue(ctx, pKey, map[string]string{})
	childCtx := context.WithValue(parnetCtx, sKey, "sub")
	m := childCtx.Value(pKey).(map[string]string)
	m["level"] = "111"

	val := parnetCtx.Value(pKey)
	fmt.Println(val)

	vals := childCtx.Value(sKey)
	fmt.Println(vals)
}

func TestParentCtx(t *testing.T) {
	ctx := context.Background()
	// deadlineCtx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second))
	deadlineCtx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	key1 := key("key1")
	childCtx := context.WithValue(deadlineCtx, key1, "value1")
	// cancel()
	time.Sleep(2 * time.Second)

	err := childCtx.Err()
	fmt.Println(err)
}

func TestContext_timeout(t *testing.T) {
	bg := context.Background()
	timeoutCtx, cancel1 := context.WithTimeout(bg, time.Second)
	subCtx, cancel2 := context.WithTimeout(timeoutCtx, time.Second*3)

	go func() {
		<-subCtx.Done()
		fmt.Println("timeout")
	}()

	time.Sleep(1 * time.Second)
	cancel1()
	cancel2()

	fmt.Println(timeoutCtx.Err())
	fmt.Println(subCtx.Err())
}

func TestTimeoutTimeAfter(t *testing.T) {
	bsChan := make(chan struct{})

	go func() {
		slowBusiness()
		bsChan <- struct{}{}
	}()

	timer := time.AfterFunc(time.Second, func() {
		fmt.Println("timeout")
	})

	<-bsChan
	fmt.Println("business end")
	timer.Stop()
}

func slowBusiness() {
	time.Sleep(2 * time.Second)
}

func TestTimeoutTimeAfter11(t *testing.T) {
	ctx := context.Background()
	timeoutCtx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	end := make(chan struct{}, 1)

	go func() {
		slowBusiness()
		end <- struct{}{}
	}()

	status := timeoutCtx.Done()

	select {
	case <-end:
		fmt.Println("finish")
	case <-status:
		fmt.Println("timeout")
	}

}

func Test11(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 30; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				fmt.Printf("%d \n", j)
			}
		}()
	}
	wg.Wait()
}

func TestParentValueCtx11(t *testing.T) {
	ctx := context.Background()
	childCtx := context.WithValue(ctx, "map", map[string]string{})
	ccChild := context.WithValue(childCtx, "key1", "value1")
	m := ccChild.Value("map").(map[string]string)
	m["key1"] = "val1"
	val := childCtx.Value("key1")
	fmt.Println(val)
	val = childCtx.Value("map")
	fmt.Println(val)
}

func Test1(t *testing.T) {
	now := time.Now()
	fmt.Println(now)
	v := now.Format("2006-01-02 15:04:05")
	fmt.Println(v)
}

func Test2(t *testing.T) {
	now := time.Now().UnixMilli()
	now1 := time.Now().UnixMicro()
	now2 := time.Now().UnixNano()
	fmt.Println(now)
	fmt.Println(now1)
	fmt.Println(now2)
	fmt.Println(time.UnixMilli(now).Format("2006-01-02 15:04:05"))
	// time.Sleep(time.Second)
	// end := time.Now()

	// elapsed := end.Sub(now)
	// fmt.Println(elapsed)
}
