package main

import "testing"

func TestName(t *testing.T) {
	p := Person{
		Name: "Tom",
	}
	p.fn = testFunc("hello ")
	t.Log(p.fn("world"))
}

type exec func(str string) string
type Person struct {
	Name string
	fn   exec
}

func testFunc(last string) exec {
	return func(str string) string {
		return last + str
	}
}
