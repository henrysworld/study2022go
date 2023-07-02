package sort

func QuickSort(arr []int) {
	quickSort1(arr, 0, len(arr)-1)
}

func QuickSortK(arr []int, k int) int {
	return quickSortK(arr, 0, len(arr)-1, k)
}

//func quickSort(arr []int, left, right int) {
//	if left >= right {
//		return
//	}
//
//	pivot := partition(arr, left, right)
//	quickSort(arr, left, pivot-1)
//	quickSort(arr, pivot+1, right)
//}

func quickSortK(arr []int, left, right, k int) int {
	//if left >= right {
	//	return
	//}

	pivot := partitionK(arr, left, right)
	if pivot == len(arr)-k {
		return arr[pivot]
	} else if pivot > len(arr)-k {
		return quickSortK(arr, left, pivot-1, k)
	} else {
		return quickSortK(arr, pivot+1, right, k)
	}
}

//func partition(arr []int, left, right int) int {
//	pivot, counter := right, left
//
//	for i := left; i < right; i++ {
//		if arr[i] < arr[pivot] {
//			arr[i], arr[counter] = arr[counter], arr[i]
//			counter++
//		}
//	}
//	arr[pivot], arr[counter] = arr[counter], arr[pivot]
//
//	return counter
//}

func partitionK(arr []int, left, right int) int {
	pivot, counter := right, left

	for i := left; i < right; i++ {
		if arr[i] < arr[pivot] {
			arr[i], arr[counter] = arr[counter], arr[i]
			counter++
		}
	}

	arr[pivot], arr[counter] = arr[counter], arr[pivot]
	return counter
}

func quickSort(arr []int, left, right int) {
	if left >= right {
		return
	}

	pivot := partition(arr, left, right)
	quickSort(arr, left, pivot-1)
	quickSort(arr, pivot+1, right)
}

func partition(arr []int, left, right int) int {
	pivot, counter := right, left

	for i := left; i < right; i++ {
		if arr[i] < arr[pivot] {
			arr[i], arr[counter] = arr[counter], arr[i]
			counter++
		}
	}

	arr[pivot], arr[counter] = arr[counter], arr[pivot]
	return counter
}

func quickSort1(arr []int, left, right int) {
	if left >= right {
		return
	}
	pivot := partition1(arr, left, right)
	quickSort1(arr, left, pivot-1)
	quickSort1(arr, pivot+1, right)
}

func partition1(arr []int, left, right int) int {
	pivot, counter := right, left
	for i := left; i < right; i++ {
		if arr[i] < arr[pivot] {
			arr[i], arr[counter] = arr[counter], arr[i]
			counter++
		}
	}

	arr[counter], arr[pivot] = arr[pivot], arr[counter]

	return counter
}
