package reflect

import (
	"errors"
	"reflect"
)

type FuncInfo struct {
	Name string
	In   []reflect.Type
	Out  []reflect.Type

	Result []any
}

func IterateFuncs(val any) (map[string]*FuncInfo, error) {
	typ := reflect.TypeOf(val)
	// v := reflect.ValueOf(val)
	if typ.Kind() != reflect.Struct && typ.Kind() != reflect.Ptr {
		return nil, errors.New("非法类型")
	}

	num := typ.NumMethod()
	fm := make(map[string]*FuncInfo, num)
	for i := 0; i < num; i++ {
		m := typ.Method(i)
		// mv := v.Method(i)
		inNum := m.Type.NumIn()
		ps := make([]reflect.Value, 0, inNum)
		ps = append(ps, reflect.ValueOf(val))
		// ps = append(ps, reflect.ValueOf(mv))
		in := make([]reflect.Type, 0, inNum)
		for j := 0; j < inNum; j++ {
			p := m.Type.In(j)
			in = append(in, p)
			if j > 0 {
				if p.Kind() == reflect.String {
					ps = append(ps, reflect.ValueOf("Henry"))
				} else {
					ps = append(ps, reflect.Zero(p))
				}
			}
		}
		outNum := m.Type.NumOut()
		out := make([]reflect.Type, 0, outNum)

		ret := m.Func.Call(ps)
		result := make([]any, 0, outNum)

		for k := 0; k < outNum; k++ {
			out = append(out, m.Type.Out(k))
			result = append(result, ret[k].Interface())
		}

		fm[m.Name] = &FuncInfo{
			Name:   m.Name,
			In:     in,
			Out:    out,
			Result: result,
		}
	}

	return fm, nil
}
