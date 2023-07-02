package sync_model

import (
	"fmt"
	"testing"
	"time"
)

func TestFanIn(t *testing.T) {

	ch := make(chan string, 6)
	go search(ch, "jonson")
	go search(ch, "olaya")
	go search(ch, "boq")

	for i := range ch {
		fmt.Println(i)
	}

}

func search(ch chan string, msg string) {
	var i int
	for {
		ch <- fmt.Sprintf("get %s %d", msg, i)
		i++
		time.Sleep(1000 * time.Millisecond)
	}
}

func TestFanIn2(t *testing.T) {

	ch1 := search2("jonson")
	ch2 := search2("olaya")
	ch3 := search2("boq")

	for {
		select {
		case msg := <-ch1:
			fmt.Println(msg)
		case msg := <-ch2:
			fmt.Println(msg)
		case msg := <-ch3:
			fmt.Println(msg)
		}
	}

}

func search2(msg string) chan string {
	ch := make(chan string, 0)
	go func() {
		var i int
		for {
			ch <- fmt.Sprintf("get %s %d", msg, i)
			i++
			time.Sleep(1000 * time.Millisecond)
		}
	}()

	return ch
}
