package demo

import (
	"errors"
	"reflect"
)

func IterateFuncs(val any) (map[string]*FuncInfo, error) {
	if val == nil {
		return nil, errors.New("输入 nil")
	}

	typ := reflect.TypeOf(val)
	value := reflect.ValueOf(val)
	for typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
		value = value.Elem()
	}
	if typ.Kind() != reflect.Struct {
		return nil, errors.New("输入 nil")
	}

	funcNum := typ.NumMethod()
	res := make(map[string]*FuncInfo, funcNum)
	for i := 0; i < funcNum; i++ {
		method := typ.Method(i)
		inNum := method.Type.NumIn()
		in := make([]reflect.Type, inNum)
		for j := 0; j < inNum; j++ {
			cIn := method.Type.In(j)
			in = append(in, cIn)
		}

		outNum := method.Type.NumOut()
		out := make([]reflect.Type, outNum)
		for k := 0; k < inNum; k++ {
			cOut := method.Type.Out(k)
			out = append(out, cOut)
		}

		callRes := method.Func.Call([]reflect.Value{reflect.ValueOf(val)})
		retVals := make([]any, 0, len(callRes))
		for _, cr := range callRes {
			retVals = append(retVals, cr.Interface())
		}
		res[method.Name] = &FuncInfo{
			Name:   method.Name,
			In:     in,
			Out:    out,
			Result: retVals,
		}
	}
	return res, nil
}

type FuncInfo struct {
	Name string
	In   []reflect.Type
	Out  []reflect.Type

	// 反射调用得到的结果
	Result []any
}
