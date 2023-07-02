package demo

import (
	"encoding/binary"
	"fmt"
	"net"
	"testing"
)

func TestClient(t *testing.T) {
	conn, err := net.Dial("tcp", ":8081")
	if err != nil {
		t.Fatal(err)
	}

	msg := "how are you"
	msgLen := len(msg)
	msgBs := make([]byte, 8)
	binary.BigEndian.PutUint64(msgBs, uint64(msgLen))
	data := append(msgBs, []byte(msg)...)
	_, err = conn.Write(data)
	if err != nil {
		conn.Close()
		return
	}

	respBs := make([]byte, 16)
	_, err = conn.Read(respBs)
	if err != nil {
		conn.Close()
	}

	fmt.Println(string(respBs))

}
