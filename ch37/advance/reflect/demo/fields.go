package demo

import (
	"errors"
	"fmt"
	"reflect"
)

func IterateFields(val any) {
	// 复杂逻辑
	res, err := iterateFields(val)

	// 简单逻辑
	if err != nil {
		fmt.Println(err)
		return
	}
	for k, v := range res {
		fmt.Println(k, v)
	}
}

func iterateFields(val any) (map[string]any, error) {
	if val == nil {
		return nil, errors.New("不能为nil")
	}
	typ := reflect.TypeOf(val)
	value := reflect.ValueOf(val)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
		value = value.Elem()
	}

	if typ.Kind() != reflect.Struct {
		return nil, errors.New("无效的类型")
	}

	num := typ.NumField()
	result := make(map[string]any, num)

	for i := 0; i < num; i++ {
		ft := typ.Field(i)
		vf := value.Field(i)
		result[ft.Name] = vf.Interface()
	}

	return result, nil
}
