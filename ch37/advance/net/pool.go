package net

import (
	"net"
	"sync"
	"sync/atomic"
	"time"
)

type Option func(p *SimplePool)

type SimplePool struct {
	idleChan chan conn
	waitChan chan *conReq

	factory     func() (net.Conn, error)
	idleTimeout time.Duration
	maxCnt      int32
	cnt         int32
	l           sync.Mutex
}

func NewSimplePool(factory func() (net.Conn, error), opts ...Option) *SimplePool {
	res := &SimplePool{
		idleChan: make(chan conn, 16),
		waitChan: make(chan *conReq, 128),
		factory:  factory,
		maxCnt:   128,
	}

	for _, opt := range opts {
		opt(res)
	}
	return res
}

func (s *SimplePool) Get() (net.Conn, error) {
	for {
		select {
		case c := <-s.idleChan:
			if c.lastActive.Add(s.idleTimeout).Before(time.Now()) {
				atomic.AddInt32(&s.cnt, -1)
				_ = c.c.Close()
				continue
			}
			return c.c, nil
		default:
			cnt := atomic.AddInt32(&s.cnt, 1)
			if cnt <= s.maxCnt {
				return s.factory()
			}
			atomic.AddInt32(&s.cnt, -1)
			req := &conReq{
				con: make(chan conn, 1),
			}

			s.waitChan <- req
			c := <-req.con
			return c.c, nil
		}
	}
}

func (s *SimplePool) Put(c net.Conn) {
	s.l.Lock()

	if len(s.waitChan) > 0 {
		req := <-s.waitChan
		s.l.Unlock()
		req.con <- conn{c: c, lastActive: time.Now()}
		return
	}

	s.l.Unlock()
	select {
	case s.idleChan <- conn{c: c, lastActive: time.Now()}:
	default:
		defer func() {
			atomic.AddInt32(&s.maxCnt, -1)
		}()
		_ = c.Close()

	}
}

func WithMaxIdleCnt(maxIdleCnt int32) Option {
	return func(p *SimplePool) {
		p.idleChan = make(chan conn, maxIdleCnt)
	}
}

func WithMaxCnt(maxCnt int32) Option {
	return func(p *SimplePool) {
		p.maxCnt = maxCnt
	}
}

type conn struct {
	c          net.Conn
	lastActive time.Time
}

type conReq struct {
	con chan conn
}
