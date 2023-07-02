package sync

import "sync"

type SafeMap1[K comparable, V any] struct {
	m     map[K]V
	mutex sync.RWMutex
}

// LoadOrStore loaded 代表是返回老的对象，还是返回了新的对象
func (s *SafeMap[K, V]) LoadOrStore1(key K, newVale V) (V, bool) {
	s.mutex.RLock()
	val, ok := s.m[key]
	s.mutex.RUnlock()
	if ok {
		return val, true
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()
	val, ok = s.m[key]
	if ok {
		return val, true
	}

	s.m[key] = newVale
	return newVale, false
}

type valProvider[V any] func() V

func (s *SafeMap[K, V]) LoadOrStoreHeavy(key K, p valProvider[V]) (interface{}, bool) {
	s.mutex.RLock()
	val, ok := s.m[key]
	s.mutex.RUnlock()
	if ok {
		return val, true
	}

	s.mutex.Lock()
	defer s.mutex.Unlock()
	val, ok = s.m[key]
	if ok {
		return val, true
	}

	val = p()
	s.m[key] = val
	return val, false
}
