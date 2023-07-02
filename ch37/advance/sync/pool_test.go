package sync

import (
	"fmt"
	"sync"
	"testing"
)

func TestPool(t *testing.T) {
	p := sync.Pool{
		New: func() interface{} {
			return nil
		},
	}

	obj := p.Get()

	p.Put(obj)
	val := 5
	array := []int{1, 2, 3}
	array1 := []int{111, 222, 333}
	array = append(array, val)
	index := 1
	fmt.Println(array)
	fmt.Println(array[index+1:])
	fmt.Println(array[index:])
	copy(array[index+1:], array[index:])
	// array_a[index] = val
	fmt.Println(array)

	fmt.Println(array1)
}
