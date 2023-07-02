package _05_array

import (
	"fmt"
	"testing"
	"time"
)

func TestBSearch(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 0
	fmt.Println(search(nums, target))
	fmt.Println("time.Now: ", time.Now())
	fmt.Println("time.Now.UTC: ", time.Now().UTC())
}

//func search(nums []int, target int) int {
//	left, right := 0, len(nums)-1
//	for left <= right {
//		mid := (left + right) / 2
//		num := nums[mid]
//		if num == target {
//			return mid
//		} else if num > target {
//			right = mid - 1
//		} else {
//			left = mid + 1
//		}
//	}
//
//	return -1
//}

//func search(nums []int, target int) int {
//	left, right := 0, len(nums)-1
//	for left <= right {
//		mid := (left + right) / 2
//		ret := nums[mid]
//		if ret == target {
//			return mid
//		} else if ret < target {
//			left = mid + 1
//		} else {
//			right = mid - 1
//		}
//	}
//
//	return -1
//}

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
