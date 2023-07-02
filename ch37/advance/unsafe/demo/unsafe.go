package demo

import (
	"errors"
	"reflect"
	"unsafe"
)

type FieldAccessor interface {
	Field(field string) (int, error)
	SetField(field string, val int) error
}

type UnsafeAccessor struct {
	fields     map[string]FieldMeta
	entityAddr unsafe.Pointer
}

func NewUnsafeAccessor(entity interface{}) (*UnsafeAccessor, error) {
	if entity == nil {
		return nil, errors.New("invalid entity")
	}

	typ := reflect.TypeOf(entity)
	if typ.Kind() != reflect.Pointer || typ.Elem().Kind() != reflect.Struct {
		return nil, errors.New("invalid entity")
	}
	elemType := typ.Elem()
	typeNum := elemType.NumField()
	fields := make(map[string]FieldMeta, typeNum)
	for i := 0; i < typeNum; i++ {
		fd := elemType.Field(i)
		fields[fd.Name] = FieldMeta{
			typ:    fd.Type,
			offset: fd.Offset,
		}
	}

	val := reflect.ValueOf(entity)
	return &UnsafeAccessor{
		fields:     fields,
		entityAddr: val.UnsafePointer(),
	}, nil
}

func (u *UnsafeAccessor) Field(field string) (int, error) {
	meta, ok := u.fields[field]
	if !ok {
		return 0, errors.New("不存在字段")
	}
	res := *(*int)(unsafe.Pointer(uintptr(u.entityAddr) + meta.offset))
	return res, nil
}

func (u *UnsafeAccessor) SetField(field string, val int) error {
	meta, ok := u.fields[field]
	if !ok {
		return errors.New("不存在字段")
	}
	*(*int)(unsafe.Pointer(uintptr(u.entityAddr) + meta.offset)) = val
	return nil
}

func (u *UnsafeAccessor) FieldAny(field string) (any, error) {
	meta, ok := u.fields[field]
	if !ok {
		return 0, errors.New("不存在字段")
	}
	res := reflect.NewAt(meta.typ, unsafe.Pointer(uintptr(u.entityAddr)+meta.offset)).Elem()
	return res.Interface(), nil
}

func (u *UnsafeAccessor) SetFieldAny(field string, val any) error {
	meta, ok := u.fields[field]
	if !ok {
		return errors.New("不存在字段")
	}
	res := reflect.NewAt(meta.typ, unsafe.Pointer(uintptr(u.entityAddr)+meta.offset)).Elem()
	if res.CanSet() {
		res.Set(reflect.ValueOf(val))
	}
	return nil
}

type FieldMeta struct {
	typ reflect.Type
	// offset 后期在我们考虑组合，或者复杂类型字段的时候，它的含义衍生为表达相当于最外层的结构体的偏移量
	offset uintptr
}
