package sync

import "sync"

type List[T any] interface {
	Get(index int) T
	Set(index int, val T)
	DeleteAt(index int) T
	Append(val T)
}

type ArrayList[T any] struct {
	vals []T
}

func NewArrayList[T any](initCap int) *ArrayList[T] {
	return &ArrayList[T]{
		vals: make([]T, 0, initCap),
	}
}

var _ List[int] = NewArrayList[int](10)

// Append implements List
func (a *ArrayList[T]) Append(val T) {
	a.vals = append(a.vals, val)
}

// DeleteAt implements List
func (a *ArrayList[T]) DeleteAt(index int) T {
	if index >= len(a.vals) || index < 0 {
		panic("index out of bounds")
	}
	res := a.vals[index]
	a.vals = append(a.vals[:index], a.vals[index+1:]...)
	return res
}

// Get implements List
func (a *ArrayList[T]) Get(index int) T {
	return a.vals[index]
}

// Set implements List
func (a *ArrayList[T]) Set(index int, val T) {
	if index >= len(a.vals) || index < 0 {
		panic("index out of bounds")
	}
	a.vals[index] = val
}

type SafeListDecorator[T any] struct {
	l     List[T]
	mutex sync.RWMutex
}

func NewSafeListDecorator[T any](initCap int) *SafeListDecorator[T] {
	return &SafeListDecorator[T]{
		l: NewArrayList[T](initCap),
	}
}

var _ List[int] = NewSafeListDecorator[int](10)

// Append implements List
func (s *SafeListDecorator[T]) Append(val T) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.l.Append(val)
}

// DeleteAt implements List
func (s *SafeListDecorator[T]) DeleteAt(index int) T {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.l.DeleteAt(index)
}

// Get implements List
func (s *SafeListDecorator[T]) Get(index int) T {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.l.Get(index)
}

func (s *SafeListDecorator[T]) GetAll(index int) T {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.l.Get(index)
}

// Set implements List
func (s *SafeListDecorator[T]) Set(index int, val T) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.l.Set(index, val)
}
