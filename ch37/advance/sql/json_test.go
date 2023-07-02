package sql

import (
	"database/sql"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJsonColumn_Value(t *testing.T) {
	js := JsonColumn[User]{Valid: true, Val: User{Name: "Tom"}}
	value, err := js.Value()
	assert.Nil(t, err)
	assert.Equal(t, []byte(`{"Name":"Tom"}`), value)
	js = JsonColumn[User]{}
	value, err = js.Value()
	assert.Nil(t, err)
	assert.Nil(t, value)
}

func TestJsonColumn_Scan(t *testing.T) {
	testCases := []struct {
		name    string
		src     any
		wantErr error
		wantVal User
	}{
		{
			name:    "nil",
			wantErr: errors.New("ekit：JsonColumn.Scan 不支持 src 类型 <nil>"),
		},
		{
			name:    "string",
			src:     `{"Name":"Tom"}`,
			wantVal: User{Name: "Tom"},
		},
		{
			name: "string array_a",
			src: func() string {
				return `{"Name":"Tom"}`
			}(),
			wantVal: User{Name: "Tom"},
		},
		{
			name:    "bytes",
			src:     []byte(`{"Name":"Tom"}`),
			wantVal: User{Name: "Tom"},
		},
		{
			name: "bytes array_a",
			src: func() *[]byte {
				res := []byte(`{"Name":"Tom"}`)
				return &res
			}(),
			wantVal: User{Name: "Tom"},
		},
		{
			name:    "sql.RawBytes",
			src:     sql.RawBytes(`{"Name":"Tom"}`),
			wantVal: User{Name: "Tom"},
		},
		{
			name: "sql.RawBytes array_a",
			src: func() *sql.RawBytes {
				res := sql.RawBytes(`{"Name":"Tom"}`)
				return &res
			}(),
			wantVal: User{Name: "Tom"},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			js := &JsonColumn[User]{}
			err := js.Scan(tc.src)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantVal, js.Val)
			assert.True(t, js.Valid)
		})
	}
}

type User struct {
	Name string
}
