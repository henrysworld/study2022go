package sync_model

import (
	"fmt"
	"testing"
)

func TestPipeline(t *testing.T) {
	multiply := func(value, multiplier int) int {
		return value * multiplier
	}

	add := func(value, additive int) int {
		return value + additive
	}

	ints := []int{1, 2, 3, 4}
	for _, v := range ints {
		fmt.Println(multiply(add(multiply(v, 2), 1), 2))
	}
}

func TestPipeline1(t *testing.T) {
	generator := func(done <-chan interface{}, integers ...int) <-chan int {
		intStream := make(chan int)
		go func() {
			defer close(intStream)
			for _, v := range integers {
				select {
				case <-done:
					return
				case intStream <- v:

				}
			}
		}()

		return intStream
	}

	mutiply := func(done <-chan interface{}, intStream <-chan int, mutiplier int) <-chan int {
		multipliedStream := make(chan int)
		go func() {
			defer close(multipliedStream)
			for v := range intStream {
				select {
				case <-done:
					return
				case multipliedStream <- v * mutiplier:

				}
			}
		}()

		return multipliedStream
	}

	add := func(done <-chan interface{}, intStream <-chan int, additive int) <-chan int {
		addedStream := make(chan int)
		go func() {
			defer close(addedStream)
			for v := range intStream {
				select {
				case <-done:
					return
				case addedStream <- v + additive:

				}
			}
		}()

		return addedStream
	}

	done := make(chan interface{})
	defer close(done)

	intStream := generator(done, 1, 2, 3, 4)
	pipeline := mutiply(done, add(done, mutiply(done, intStream, 2), 1), 2)

	for v := range pipeline {
		fmt.Println(v)
	}
}
