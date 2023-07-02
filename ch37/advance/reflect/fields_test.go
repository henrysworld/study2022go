package reflect

import (
	"fmt"
	"testing"

	"github.com/henrysworld/study2022go/ch37/advance/reflect/types"
	"github.com/magiconair/properties/assert"
)

func TestIterateFields(t *testing.T) {
	up := &types.User{}
	up2 := &up
	testCases := []struct {
		name       string
		input      any
		wantFields map[string]any
		wantErr    error
	}{
		{
			name: "normal struct",
			input: types.User{
				Name: "Tom",
			},
			wantFields: map[string]any{
				"Name": "Tom",
				"age":  0,
			},
		},
		{
			// 指针
			name: "array_a",
			input: &types.User{
				Name: "Tom",
			},
			wantFields: map[string]any{
				"Name": "Tom",
				"age":  0,
			},
		},
		{
			// 多重指针
			name:  "multiple array_a",
			input: up2,
			wantFields: map[string]any{
				"Name": "",
				"age":  0,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := iterateFields(tc.input)
			if err != nil {
				assert.Equal(t, tc.wantErr, err)
				return
			}
			assert.Equal(t, tc.wantFields, res)
		})
	}

}

func TestSetField(t *testing.T) {
	testCases := []struct {
		name string

		field  string
		entity any
		newVal any

		wantErr error
	}{
		// {
		// 	name:    "struct",
		// 	entity:  types.User{},
		// 	field:   "Name",
		// 	wantErr: errors.New("非法类型"),
		// },
		// {
		// 	name:    "private field",
		// 	entity:  &types.User{},
		// 	field:   "age",
		// 	wantErr: errors.New("不可修改字段"),
		// },
		// {
		// 	name:    "invalid field",
		// 	entity:  &types.User{},
		// 	field:   "invalid_field",
		// 	wantErr: errors.New("字段不存在"),
		// },
		{
			name: "pass",
			entity: &types.User{
				Name: "",
			},
			field:  "Name",
			newVal: "Tom",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := SetField(tc.entity, tc.field, tc.newVal)
			fmt.Println(tc.entity.(*types.User).Name)
			assert.Equal(t, tc.wantErr, err)
		})
	}
}
