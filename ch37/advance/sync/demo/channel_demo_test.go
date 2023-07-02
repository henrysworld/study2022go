package demo

import (
	"fmt"
	"testing"
	"time"

	"github.com/magiconair/properties/assert"
)

func TestChannel(t *testing.T) {
	ch := make(chan string, 4)

	go func() {
		str := <-ch
		fmt.Println(str)
	}()

	go func() {
		str := <-ch
		fmt.Println(str)
	}()

	go func() {
		str := <-ch
		fmt.Println(str)
	}()

	ch <- "Hello, world!"
	ch <- "Hello, world!"

	time.Sleep(2 * time.Second)
}

func TestBroker(t *testing.T) {
	b := &Broker{
		consumers: make([]*Consumer, 0, 10),
	}

	c1 := &Consumer{
		ch: make(chan string, 1),
	}

	c2 := &Consumer{
		ch: make(chan string, 1),
	}

	b.Subscribe(c1)
	b.Subscribe(c2)

	b.Produce("Hello, world! Broker!")

	fmt.Println("C1", <-c1.ch)
	fmt.Println("C2", <-c2.ch)
}

type Broker struct {
	consumers []*Consumer
}

func (b *Broker) Produce(msg string) {
	for _, c := range b.consumers {
		c.ch <- msg
	}
}

func (b *Broker) Subscribe(c *Consumer) {
	b.consumers = append(b.consumers, c)
}

type Consumer struct {
	ch chan string
}

type Broker1 struct {
	ch        chan string
	consumers []func(s string)
}

func (b *Broker1) Produce(msg string) {
	b.ch <- msg
}

func (b *Broker1) Subscribe(consumer func(s string)) {
	b.consumers = append(b.consumers, consumer)
}

func (b *Broker1) Start() {
	go func() {
		for {
			s, ok := <-b.ch
			if !ok {
				return
			}
			for _, c := range b.consumers {
				c(s)
			}
		}
	}()
}

func NewBroker1() *Broker1 {
	b := &Broker1{
		ch:        make(chan string, 10),
		consumers: make([]func(s string), 0, 10),
	}

	go func() {
		for {
			s, ok := <-b.ch
			if !ok {
				return
			}

			for _, c := range b.consumers {
				c(s)
			}
		}
	}()

	return b
}

func TestBroker1(t *testing.T) {
	b := NewBroker1()

	str1 := ""
	b.Subscribe(func(s string) {
		str1 = str1 + s
	})

	str2 := ""
	b.Subscribe(func(s string) {
		str2 = str2 + s
	})

	b.Produce("hello")
	b.Produce(" ")
	b.Produce("world")

	time.Sleep(time.Second)
	assert.Equal(t, "hello world", str1)
	assert.Equal(t, "hello world", str2)

}

type Broker3 struct {
	consumers []*Consumer3
}

func (b *Broker3) Produce(msg string) {
	go func() {
		for _, c := range b.consumers {
			c.ch <- msg
		}
	}()
}

func (b *Broker3) Subscribe(c *Consumer3) {
	b.consumers = append(b.consumers, c)
}

type Consumer3 struct {
	ch chan string
}

func TestBroker3(t *testing.T) {
	b := &Broker3{
		consumers: make([]*Consumer3, 0, 10),
	}

	c1 := &Consumer3{
		ch: make(chan string, 1),
	}

	c2 := &Consumer3{
		ch: make(chan string, 1),
	}

	c3 := &Consumer3{
		ch: make(chan string, 1),
	}

	b.Subscribe(c1)
	b.Subscribe(c2)
	b.Subscribe(c3)

	b.Produce("hello")
	b.Produce(" ")
	b.Produce("world ")

	for {
		fmt.Println("C1", <-c1.ch)
		fmt.Println("C2", <-c2.ch)
		fmt.Println("C3", <-c3.ch)
	}

}
