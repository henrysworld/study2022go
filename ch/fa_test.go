package ch

import (
	"fmt"
	"testing"
)

// OperatorFactory 是工厂接口
type OperatorFactory interface {
	Create()
}
type PlusOperatorFactory struct{}

func (PlusOperatorFactory) Create() {
	fmt.Println("aaaa")
	return
}

func TestNamae(t *testing.T) {
	var factory OperatorFactory

	factory = PlusOperatorFactory{}

	factory.Create()
}
