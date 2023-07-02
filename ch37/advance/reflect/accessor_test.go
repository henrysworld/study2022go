package reflect

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestReflectAccessor_Field(t *testing.T) {
	testCases := []struct {
		name   string
		entity interface{}
		field  string

		wantVal int
		wantErr error
	}{
		{
			name:    "normal case",
			entity:  &User{Age: 18},
			field:   "Age",
			wantVal: 18,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accessor, err := NewReflectAccessor(tc.entity)
			if err != nil {
				assert.Equal(t, tc.wantErr, err)
			}

			val, err := accessor.Field(tc.field)
			if err != nil {
				assert.Equal(t, tc.wantErr, err)
			}

			assert.Equal(t, tc.wantVal, val)

		})
	}
}

func TestReflectAssessor_SetField(t *testing.T) {
	testCases := []struct {
		name     string
		entity   *User
		field    string
		inputVal int

		wantVal int
		wantErr error
	}{
		{
			name:     "normal case",
			entity:   &User{Age: 18},
			field:    "Age",
			inputVal: 110,
			wantVal:  110,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			accessor, err := NewReflectAccessor(tc.entity)
			if err != nil {
				assert.Equal(t, tc.wantErr, err)
			}

			err = accessor.SetField(tc.field, tc.inputVal)
			if err != nil {
				assert.Equal(t, tc.wantErr, err)
			}

			fmt.Println(tc.entity.Age)
			assert.Equal(t, tc.wantVal, tc.entity.Age)

		})
	}
}

func TestTypeOf(t *testing.T) {
	// Define variable
	// var x int = 123
	var u1 = User{
		Age: 180,
	}
	var u2 = &User{
		Age: 18,
	}

	// Get type of variable
	zu1 := reflect.TypeOf(u1)
	zu2 := reflect.TypeOf(u2)
	val1 := reflect.ValueOf(u1)
	val2 := reflect.ValueOf(u2)
	// vale1 := reflect.ValueOf(u1).Elem()
	vale2 := reflect.ValueOf(u2).Elem()

	if zu1.Kind() == reflect.Struct {
		fmt.Println("u1 is struct")
	}

	if zu2.Kind() == reflect.Pointer {
		fmt.Println("u2 is array_a")
	}

	// Print results
	fmt.Println(zu1) // "int"
	fmt.Println(zu2)
	fmt.Println(val1) // "int"
	fmt.Println(val2)
	// fmt.Println(vale1) // "int"
	fmt.Println(vale2)
}

type User struct {
	Age int
}
