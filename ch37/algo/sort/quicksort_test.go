package sort

import "testing"

func TestQuickSort(t *testing.T) {
	//arr := []int{5, 4}
	//QuickSort(arr)
	//t.Log(arr)

	//arr := []int{5, 4, 3, 2, 1}
	arr := []int{6, 11, 3, 8, 10, 9, 8}
	ret := QuickSortK(arr, 6)
	t.Log(arr)
	t.Log(ret)
	arr = []int{6, 11, 3, 8, 10, 9, 8}
	//QuickSort(arr)
	QuickSort(arr)
	t.Log(arr)
}
