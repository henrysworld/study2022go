package reflect

import (
	"errors"
	"fmt"
	"reflect"
)

func IterateFields(val any) {
	res, err := iterateFields(val)
	if err != nil {
		fmt.Println(err)
		return
	}

	for k, v := range res {
		fmt.Println(k, v)
	}
}

func iterateFields(input any) (map[string]any, error) {
	typ := reflect.TypeOf(input)
	val := reflect.ValueOf(input)

	// if typ.Kind() == reflect.Ptr {
	for typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
		val = val.Elem()
	}

	if typ.Kind() != reflect.Struct {
		return nil, errors.New("非法类型")
	}

	num := typ.NumField()
	ret := make(map[string]any, num)
	for i := 0; i < num; i++ {
		f := typ.Field(i)
		v := val.Field(i)
		if f.IsExported() {
			ret[f.Name] = v.Interface()
		} else {
			ret[f.Name] = reflect.Zero(f.Type).Interface()
		}
	}

	return ret, nil
}

func SetField(entity any, field string, val any) error {
	value := reflect.ValueOf(entity)
	typ := value.Type()
	if typ.Kind() != reflect.Ptr || typ.Elem().Kind() != reflect.Struct {
		return errors.New("非法类型")
	}

	typ = typ.Elem()
	value = value.Elem()

	_, found := typ.FieldByName(field)
	if !found {
		return errors.New("字段不存在")
	}

	fd := value.FieldByName(field)

	if fd.CanSet() {
		fd.Set(reflect.ValueOf(val))
	} else {
		return errors.New("不可修改字段")
	}

	return nil
}
