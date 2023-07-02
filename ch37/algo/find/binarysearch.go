package find

//func BinarySearch(nums []int, target int) int {
//	left, right := 0, len(nums)-1
//	for left <= right {
//		mid := (left + right) / 2
//		if nums[mid] == target {
//			return mid
//		} else if nums[mid] < target {
//			left = mid + 1
//		} else {
//			right = mid - 1
//		}
//	}
//
//	return -1
//}

func BinarySearch(nums []int, target int) int {
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

func BinarySearch1(nums []int, target int) int {
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

func BinarySearch2(nums []int, target int) int {
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
