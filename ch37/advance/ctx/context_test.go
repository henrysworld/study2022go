package ctx

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"golang.org/x/sync/errgroup"
)

type Cache interface {
	Get(key string) (string, error)
}

type OtherCache interface {
	GetValue(ctx context.Context, key string) (any, error)
}

type CacheAdapter struct {
	Cache
}

func (c *CacheAdapter) GetValue(ctx context.Context, key string) (any, error) {
	return c.Cache.Get(key)
}

// 已有的，不是线程安全的
type memoryMap struct {
	// 如果你这样添加锁，那么就是一种侵入式的写法
	// 那么你就需要测试这个类
	// 而且有些时候，这个事第三方的依赖，你都改不了
	// lock sync.RWMutex
	m map[string]string
}

func (m *memoryMap) Get(key string) (string, error) {
	return m.m[key], nil
}

var s = &SafeCache{
	Cache: &memoryMap{},
}

// SafeCache 我要改造为线程安全的
// 无侵入式地改造

type SafeCache struct {
	Cache
	lock sync.RWMutex
}

func (s *SafeCache) Get(key string) (string, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.Cache.Get(key)
}

func TestErrgroup(t *testing.T) {
	eg, ctx := errgroup.WithContext(context.Background())
	var result int64 = 0
	for i := 0; i < 10; i++ {
		delta := i
		eg.Go(func() error {
			atomic.AddInt64(&result, int64(delta))
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		t.Fatal(err)
	}

	ctx.Err()
	fmt.Println(result)
}

func TestBusinessTimeout1(t *testing.T) {
	ctx := context.Background()
	toCtx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	end := make(chan struct{}, 1)

	go func() {
		MyBusiness1()
		end <- struct{}{}
	}()
	ch := toCtx.Done()
	err := toCtx.Err()
	fmt.Println("ch=", err)

	select {
	case <-ch:
		fmt.Println("timeout")
	case <-end:
		fmt.Println("biz finish")
	}

}

func MyBusiness1() {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("test")
}

type mkey string

func TestParentValueCtx1(t *testing.T) {
	ctx := context.Background()
	key1 := mkey("key1")
	key2 := mkey("key2")
	childCtx := context.WithValue(ctx, key1, "value1")
	ccChildCtx := context.WithValue(childCtx, key2, "value2")

	val1 := ccChildCtx.Value(key1)
	fmt.Println(val1)
	val2 := ccChildCtx.Value(key2)
	fmt.Println(val2)
}
