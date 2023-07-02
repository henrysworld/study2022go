package find

import (
	"fmt"
	"testing"
)

func TestBSearch(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 12
	fmt.Println(BinarySearch(nums, target))
}
