package sort

import (
	"fmt"
	"testing"
)

func TestBubble(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}

	length := len(nums)
	for i := 0; i < length; i++ {

		flag := false
		for j := 0; j < length-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}

	fmt.Printf("%v", nums)

}

func TestBubble1(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}

	length := len(nums)
	for i := 0; i < length; i++ {
		flag := false
		for j := 0; j < length-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				flag = true
			}
		}

		if !flag {
			break
		}
	}

	fmt.Printf("%v", nums)

}
