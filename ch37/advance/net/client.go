package net

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

const lenBytes = 8

func Connect(addr string) error {
	conn, err := net.DialTimeout("tcp", addr, 3*time.Second)
	if err != nil {
		return err
	}

	defer func() {
		_ = conn.Close()
	}()

	for {
		_, err := conn.Write([]byte("hello"))
		if err != nil {
			return err
		}

		res := make([]byte, 8)
		_, err = conn.Read(res)
		if err != nil {
			return err
		}

		fmt.Println(string(res))
		time.Sleep(time.Second)
	}

}

type Client struct {
	addr string
}

func (c *Client) Send(msg string) (string, error) {
	conn, err := net.DialTimeout("tcp", c.addr, 3*time.Second)
	if err != nil {
		return "", err
	}
	defer func() {
		_ = conn.Close()
	}()

	//总长度
	bs := make([]byte, lenBytes, len(msg)+lenBytes)
	//写入消息长度
	binary.BigEndian.PutUint64(bs, uint64(len(msg)))
	bs = append(bs, msg...)
	_, err = conn.Write(bs)
	if err != nil {
		return "", err
	}

	lenBs := make([]byte, lenBytes)
	_, err = conn.Read(lenBs)
	if err != nil {
		return "", err
	}

	resLength := binary.BigEndian.Uint64(lenBs)

	//读取相应
	reBs := make([]byte, resLength)
	_, err = conn.Read(reBs)
	if err != nil {
		return "", err
	}

	return string(reBs), nil
}
