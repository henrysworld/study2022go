package sync

import (
	"fmt"
	"testing"
	"time"
)

func TestChannelReceive(t *testing.T) {
	ch := make(chan string, 1)
	go func() {
		data := <-ch
		fmt.Printf("g1 receiver %s", string(data))
	}()

	go func() {
		data := <-ch
		fmt.Printf("g2 receiver %s", string(data))
	}()

	ch <- "data streaming"
	time.Sleep(3 * time.Second)
}
