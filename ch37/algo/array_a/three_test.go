package array_a

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"testing"
)

func TestThree(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}

	//i, j, k := 0, 1, 2
	l := len(nums)
	var ret [][]int
	for i := 0; i < len(nums); i++ {
		obj1 := nums[i]
		for j := i + 1; j < l; j++ {
			obj2 := nums[j]
			for k := j + 1; k < l; k++ {
				obj3 := nums[k]
				if (obj1 + obj2 + obj3) == 0 {
					var retSub []int
					retSub = append(retSub, obj1)
					retSub = append(retSub, obj2)
					retSub = append(retSub, obj3)
					ret = append(ret, retSub)
				}
			}
		}
	}

	var rets [][]int
	mp := make(map[string][]int)
	for _, intSlice := range ret {
		sort.Ints(intSlice)
		var strSlice []string
		for _, num := range intSlice {
			strSlice = append(strSlice, strconv.Itoa(num))
		}
		key := strings.Join(strSlice, "_")

		mp[key] = intSlice
	}

	for _, v := range mp {
		rets = append(rets, v)
	}

	fmt.Printf("%v", rets)
}

// TestThree2
// 标签：数组遍历
// 首先对数组进行排序，排序后固定一个数 nums，再使用左右指针指向nums[i]后面的两端，数字分别为 nums[L]和nums[R]，计算三个数的和sum判断是否满足为0，满足则添加进结果集。
// 如果 nums[i]大于 0，则三数之和必然无法等于 0，结束循环
// 如果 nums[i] == nums[i]- 1，则说明该数字重复，会导致结果重复，所以应该跳过。
// 当 sum ==0 时，nums[L] == nums[L]+1 则会导致结果重复，应该跳过，L++
// 当 sum ==0 时，nums[R] == nums[R]-1 则会导致结果重复，应该跳过，R--
// 时间复杂度: 0(m2)，n 为数组长度

func TestThree2(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	sort.Ints(nums)
	length := len(nums)
	rets := make([][]int, 0)

	for i := 0; i < length-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		L := i + 1
		R := length - 1

		for L < R {
			sum := nums[i] + nums[L] + nums[R]
			if sum == 0 {
				rets = append(rets, []int{nums[i], nums[L], nums[R]})
				for L < R && nums[L] == nums[L+1] {
					L++
				}

				for L < R && nums[R] == nums[R-1] {
					R--
				}
				L++
				R--
			} else if sum < 0 {
				L++
			} else {
				R--
			}
		}

	}

	fmt.Printf("%v", rets)

}

func TestThree3(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	sort.Ints(nums)
	length := len(nums)
	rets := make([][]int, 0)

	for i := 0; i < length-2; i++ {
		if nums[i] > 0 {
			break
		}
		if nums[i] == nums[i+1] {
			continue
		}
		L := i + 1
		R := length - 1
		for L < R {
			sum := nums[i] + nums[L] + nums[R]
			if sum == 0 {
				rets = append(rets, []int{nums[i], nums[L], nums[R]})
				for nums[L] == nums[L+1] {
					L++
				}

				for nums[R] == nums[R-1] {
					R--
				}

				L++
				R--
			} else if sum < 0 {
				L++
			} else {
				R--
			}
		}

	}
}
