package sync

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestArrayList_DeleteAt(t *testing.T) {
	testCases := []struct {
		name     string
		index    int
		input    []int
		wantVals []int
	}{
		{
			name:     "first",
			index:    0,
			input:    []int{1, 2, 3},
			wantVals: []int{2, 3},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			a := NewArrayList[int](12)
			a.vals = tc.input
			_ = a.DeleteAt(tc.index)
			assert.Equal(t, tc.wantVals, a.vals)
		})
	}
}

func TestSafeArrayList_DeleteAt(t *testing.T) {
	testCases := []struct {
		name      string
		args      int
		input     []int
		wantVals  []int
		wantError error
	}{
		{
			name:      "middle",
			args:      1,
			input:     []int{1, 2, 3},
			wantVals:  []int{1, 3},
			wantError: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sa := NewSafeListDecorator[int](12)
			a := sa.l.(*ArrayList[int])
			a.vals = tc.input
			_ = sa.DeleteAt(tc.args)
			assert.Equal(t, tc.wantVals, a.vals)
		})
	}
}
