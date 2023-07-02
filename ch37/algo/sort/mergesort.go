package sort

func MergeSort(arr []int) {
	arrLen := len(arr)
	if arrLen <= 1 {
		return
	}

	mergeSort(arr, 0, arrLen-1)
}

//func mergeSort(arr []int, left int, right int) {
//	if left >= right {
//		return
//	}
//	mid := (left + right) / 2
//
//	mergeSort(arr, left, mid)
//	mergeSort(arr, mid+1, right)
//	merge(arr, left, mid, right)
//}
//
//func merge(arr []int, left, mid, right int) {
//	//我们申请一个临时数组 tmp，大小与 A[p...r]相同。
//	tmp := make([]int, right-left+1)
//	//我们用两个游标 i 和 j，分别指向 A[p...q]和 A[q+1...r]的第一个元素
//	i := left
//	j := mid + 1
//	k := 0
//
//	//比较这两个元素 A[i]和 A[j]，如果 A[i]<=A[j]，我们就把 A[i]放入到临时数组 tmp，并且 i 后移一位，否则将 A[j]放入到数组 tmp，j 后移一位。
//	for i <= mid && j <= right {
//		if arr[i] <= arr[j] {
//			tmp[k] = arr[i]
//			i++
//		} else {
//			tmp[k] = arr[j]
//			j++
//		}
//
//		k++
//	}
//
//	// 判断哪个子数组中有剩余的数据
//	// 继续上述比较过程，直到其中一个子数组中的所有数据都放入临时数组中，再把另一个数组中的数据依次加入到临时数组的末尾，这个时候，临时数组中存储的就是两个子数组合并之后的结果了
//	for i <= mid {
//		tmp[k] = arr[i]
//		k++
//		i++
//	}
//
//	for j <= right {
//		tmp[k] = arr[j]
//		k++
//		j++
//	}
//
//	//最后再把临时数组 tmp 中的数据拷贝到原数组 A[p...r]中。
//	copy(arr[left:right+1], tmp)
//}
//
//func mergeSort(arr []int, left, right int) {
//	if left >= right {
//		return
//	}
//
//	mid := (left + right) / 2
//
//	mergeSort(arr, left, mid)
//	mergeSort(arr, mid+1, right)
//
//	merge(arr, left, mid, right)
//}
//
//func merge(arr []int, left, mid, right int) {
//	tmp := make([]int, len(arr))
//	i := left
//	j := mid + 1
//	k := 0
//	for i <= mid && j <= right {
//		if arr[i] <= arr[j] {
//			tmp[k] = arr[i]
//			i++
//		} else {
//			tmp[k] = arr[j]
//			j++
//		}
//		k++
//	}
//
//	for i <= mid {
//		tmp[k] = arr[i]
//		i++
//		k++
//	}
//
//	for j <= right {
//		tmp[k] = arr[j]
//		j++
//		k++
//	}
//
//	copy(arr[left:right+1], tmp)
//}

func mergeSort(arr []int, left, right int) {
	if left >= right {
		return
	}

	mid := (left + right) / 2
	mergeSort(arr, left, mid)
	mergeSort(arr, mid+1, right)
	merge(arr, left, mid, right)
}

func merge(arr []int, left, mid, right int) {
	tmp := make([]int, len(arr))

	i := left
	j := mid + 1
	k := 0

	for i <= mid && j <= right {
		if arr[i] <= arr[j] {
			tmp[k] = arr[i]
			i++
			k++
		} else {
			tmp[k] = arr[j]
			j++
			k++
		}
	}

	for i <= mid {
		tmp[k] = arr[i]
		i++
		k++
	}

	for j <= right {
		tmp[k] = arr[j]
		j++
		k++
	}

	copy(arr[left:right+1], tmp)
}
