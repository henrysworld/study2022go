package _05_array

import (
	"fmt"
	"testing"
)

func TestRing(t *testing.T) {
	ringSize := 4
	r := NewRing(ringSize)

	// Test empty ring
	_, err := r.Get()
	if err == nil || err.Error() != "ring buffer empty" {
		t.Error("Expected error 'ring buffer empty' when getting from an empty ring")
	}

	// Test adding elements and getting elements
	for i := 0; i < ringSize; i++ {
		proto, err := r.Set()
		if err != nil {
			t.Errorf("Unexpected error when setting a proto at index %d: %v", i, err)
		}
		proto.Ver = int32(i)
		r.SetAdv()
	}

	// Test full ring
	_, err = r.Set()
	if err == nil || err.Error() != "ring buffer full" {
		t.Error("Expected error 'ring buffer full' when adding element to a full ring")
	}

	for i := 0; i < ringSize; i++ {
		proto, err := r.Get()
		if err != nil {
			t.Errorf("Unexpected error when getting a proto at index %d: %v", i, err)
		}
		fmt.Println(proto.Ver)
		if proto.Ver != int32(i) {
			t.Errorf("Expected proto.Ver to be %d, but got %d", i, proto.Ver)
		}
		r.GetAdv()
	}

	// Test reset
	r.Reset()
	_, err = r.Get()
	if err == nil || err.Error() != "ring buffer empty" {
		t.Error("Expected error 'ring buffer empty' when getting from a reset ring")
	}
}
