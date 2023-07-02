package net

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net"
	"testing"
	"time"
)

func TestSimplePool(t *testing.T) {
	p := NewSimplePool(func() (net.Conn, error) {
		return &mockConn{}, nil
	}, WithMaxIdleCnt(2), WithMaxCnt(3))
	// 这三次都能正常拿出来
	c1, err := p.Get()
	assert.Nil(t, err)
	_, err = p.Get()
	assert.Nil(t, err)
	_, err = p.Get()
	assert.Nil(t, err)

	now := time.Now()

	go func() {
		// 睡两秒
		time.Sleep(time.Second)
		p.Put(c1)
	}()
	// 直接阻塞
	c4, err := p.Get()
	assert.Nil(t, err)
	// 就是我们放回去的那个
	assert.Equal(t, c1, c4)
	// 确认被阻塞过
	assert.Greater(t, time.Now().Sub(now), time.Second)
}

// mockConn 用于辅助测试
type mockConn struct {
	closed bool
}

func (m *mockConn) Read(b []byte) (n int, err error) {
	// TODO implement me
	panic("implement me")
}

func (m *mockConn) Write(b []byte) (n int, err error) {
	// TODO implement me
	panic("implement me")
}

func (m *mockConn) Close() error {
	// 用于辅助测试
	fmt.Println("connection closing")
	m.closed = true
	return nil
}

func (m *mockConn) LocalAddr() net.Addr {
	// TODO implement me
	panic("implement me")
}

func (m *mockConn) RemoteAddr() net.Addr {
	// TODO implement me
	panic("implement me")
}

func (m *mockConn) SetDeadline(t time.Time) error {
	// TODO implement me
	panic("implement me")
}

func (m *mockConn) SetReadDeadline(t time.Time) error {
	// TODO implement me
	panic("implement me")
}

func (m *mockConn) SetWriteDeadline(t time.Time) error {
	// TODO implement me
	panic("implement me")
}
