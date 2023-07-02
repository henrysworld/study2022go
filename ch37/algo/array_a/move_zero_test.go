package array_a

import (
	"fmt"
	"testing"
)

func TestMoveZero(t *testing.T) {
	arr := []int{1, 5, 0, 0, 4, 0, 0, 0, 6, 3, 9, 6, 0, 11}
	fmt.Println(moveZero(arr))
}

func moveZero(arr []int) []int {
	left, right, end := 0, 0, len(arr)

	for right < end {
		if arr[right] != 0 {
			arr[left], arr[right] = arr[right], arr[left]
			left++
		}
		right++
	}

	return arr
}

func moveZero1(arr []int) []int {
	left, right, end := 0, 0, len(arr)

	for right < end {
		if arr[right] != 0 {
			arr[left], arr[right] = arr[right], arr[left]
			left++
		}

		right++
	}

	return arr
}
