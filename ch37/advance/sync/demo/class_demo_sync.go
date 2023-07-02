package demo

import "sync"

type SafeMap[K comparable, V any] struct {
	m     map[K]V
	mutex sync.RWMutex
}

func (s *SafeMap[K, V]) LoadOrStoreV3(key K, value V) (V, bool) {
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

	s.m[key] = value

	return value, false
}
