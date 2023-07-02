package reflect

import (
	"errors"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestIterate(t *testing.T) {
	// f := []byte("a")
	// fmt.Println(f)
	testCases := []struct {
		name    string
		input   any
		wantRes []any
		wantErr error
	}{
		{
			name: "slice",
			// input:   []int{1, 2, 3},
			// wantRes: []any{1, 2, 3},
			input:   "abc",
			wantRes: []any{[]byte("a"), []byte("b"), []byte("c")},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := Iterate(tc.input)
			// assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantRes, res)
		})
	}
}

func TestIterateMap(t *testing.T) {
	testCases := []struct {
		name       string
		input      any
		wantKeys   []any
		wantValues []any
		wantErr    error
	}{
		{
			name:    "nil",
			input:   nil,
			wantErr: errors.New("非法类型"),
		},
		{
			name: "happy case",
			input: map[string]string{
				"a_k": "a_v",
			},
			wantKeys:   []any{"a_k"},
			wantValues: []any{"a_v"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			keys, vals, err := IterateMapV1(tc.input)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantKeys, keys)
			assert.Equal(t, tc.wantValues, vals)

			keys, vals, err = IterateMapV2(tc.input)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantKeys, keys)
			assert.Equal(t, tc.wantValues, vals)
		})
	}
}
