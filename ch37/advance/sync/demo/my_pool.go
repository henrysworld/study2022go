package demo

import (
	"sync"
	"unsafe"
)

type MyPool struct {
	p      sync.Pool
	maxCnt int
	cnt    int
}

func (m *MyPool) Get() any {
	return m.p.Get()
}

func (m *MyPool) Set(a any) {
	if unsafe.Sizeof(a) > 1024 {
		return
	}
	m.p.Put(a)
}
