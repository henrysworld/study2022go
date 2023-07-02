package reflect

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/henrysworld/study2022go/ch37/advance/reflect/types"
	"github.com/magiconair/properties/assert"
)

func TestIterateFuncs(t *testing.T) {
	testCases := []struct {
		name string

		input any

		wantRes map[string]*FuncInfo
		wantErr error
	}{
		// {
		// 	// 普通结构体
		// 	name:  "normal struct",
		// 	input: types.User{},
		// 	wantRes: map[string]*FuncInfo{
		// 		"GetAge": {
		// 			Name:   "GetAge",
		// 			In:     []reflect.Type{reflect.TypeOf(types.User{})},
		// 			Out:    []reflect.Type{reflect.TypeOf(0)},
		// 			Result: []any{0},
		// 		},
		// 	},
		// },
		{
			// 指针
			name:  "array_a",
			input: &types.User{},
			wantRes: map[string]*FuncInfo{
				"GetAge": {
					Name:   "GetAge",
					In:     []reflect.Type{reflect.TypeOf(&types.User{})},
					Out:    []reflect.Type{reflect.TypeOf(0)},
					Result: []any{0},
				},
				"ChangeName": {
					Name:   "ChangeName",
					In:     []reflect.Type{reflect.TypeOf(&types.User{}), reflect.TypeOf("")},
					Out:    []reflect.Type{},
					Result: []any{},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := IterateFuncs(tc.input)
			fmt.Println(tc.input.(*types.User).Name)
			// assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantRes, res)
		})
	}
}
