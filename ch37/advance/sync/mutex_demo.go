package sync

import "sync"

type SafeMap[K comparable, V any] struct {
	m     map[K]V
	mutex sync.RWMutex
}

func (s *SafeMap[K, V]) LoadOrStore(key K, newValue V) (V, bool) {
	val, ok := s.get(key)
	if ok {
		return val, true
	}
	s.mutex.Lock()
	defer s.mutex.Unlock()
	val, ok = s.get(key)
	if ok {
		return val, true
	}

	s.m[key] = newValue
	return newValue, false
}

func (s *SafeMap[K, V]) get(key K) (V, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	val, ok := s.m[key]
	return val, ok
}

type ConcurrentArrayList[T any] struct {
	vals  []T
	mutex sync.RWMutex
	sync.Mutex
}

func NewConcurrentArrayList[T any](initCap int) *ConcurrentArrayList[T] {
	return &ConcurrentArrayList[T]{
		vals: make([]T, initCap),
	}
}

func (c *ConcurrentArrayList[T]) Get(index int) T {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return c.vals[index]
}

func (c *ConcurrentArrayList[T]) DeleteAt(index int) T {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	res := c.vals[index]
	c.vals = append(c.vals[:index], c.vals[index+1:]...)
	return res
}

func (c *ConcurrentArrayList[T]) Append(val T) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.vals = append(c.vals, val)
}
