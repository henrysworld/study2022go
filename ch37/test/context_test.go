package test

import (
	"fmt"
	"testing"
)

type key string

func TestC(t *testing.T) {
	println(f1())
	println(f2())
	println(f3())
}

type number int

func (n number) print() {
	fmt.Printf("输出 number 值 print: %v\n", n)
}

func (n *number) pprint() {
	fmt.Printf("输出 number 值 pprint: %v\n", *n)
}

func f1() (r int) {
	defer func() {
		r++
	}()
	return 0
}

func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}
