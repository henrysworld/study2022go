package net

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

func Serve(addr string) error {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}

		go func() {
			handleConn(conn)
		}()
	}
}

func handleConn(conn net.Conn) {
	for {
		bs := make([]byte, lenBytes)
		_, err := conn.Read(bs)
		if err == io.EOF || err == net.ErrClosed || err == io.ErrUnexpectedEOF {
			_ = conn.Close()
			return
		}

		if err != nil {
			continue
		}

		res := handleMsg(bs)
		_, err = conn.Write(res)
		if err == io.EOF || err == net.ErrClosed ||
			err == io.ErrUnexpectedEOF {
			_ = conn.Close()
			return
		}
	}
}

type Server struct {
	addr string
}

func handleMsg(bs []byte) []byte {
	return []byte("world")
}

func (s *Server) StartAndServe() error {
	listener, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}
		go func() {
			er := s.handleConn(conn)
			if er != nil {
				conn.Close()
				fmt.Printf("connect error: %v", er)
			}
		}()
	}
}

func (s *Server) handleConn(conn net.Conn) error {
	for {
		bs := make([]byte, lenBytes)
		_, err := conn.Read(bs)
		if err != nil {
			return err
		}

		reqBs := make([]byte, binary.BigEndian.Uint64(bs))
		_, err = conn.Read(reqBs)
		if err != nil {
			return err
		}

		res := string(reqBs) + ", from response"
		bs = make([]byte, lenBytes, len(res)+lenBytes)
		binary.BigEndian.PutUint64(bs, uint64(len(res)))
		bs = append(bs, res...)
		_, err = conn.Write(bs)
		if err != nil {
			return err
		}

	}
}
