package sync

import (
	"fmt"
	"reflect"
	"testing"
)

type MyStruct struct {
	// 省略其他字段
}

func (m *MyStruct) MyMethod() {
	fmt.Println("MyMethod is called.")
}

func TestReflect(t *testing.T) {

	s := MyStruct{}

	// 获取 MyStruct 类型的反射值
	val := reflect.ValueOf(s)

	// 获取 MyMethod 方法
	method := val.MethodByName("MyMethod")

	// 调用 MyMethod 方法
	method.Call(nil) // 输出: MyMethod is called.
}
