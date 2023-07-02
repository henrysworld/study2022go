package tree

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	r := twoSum6(nums, 10)
	fmt.Println(r)
}

func twoSum(nums []int, target int) []int {
	var ans []int
	for i := 0; i < len(nums); i++ {
		iVal := nums[i]
		for j := i + 1; j < len(nums); j++ {
			jVal := nums[j]
			if iVal+jVal == target {
				ans = append(ans, i)
				ans = append(ans, j)
			}
		}
	}

	return ans
}

func twoSum1(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, cur := range nums {
		if p, ok := hashTable[target-cur]; ok {
			return []int{p, i}
		}
		hashTable[cur] = i
	}

	return nil
}

func twoSum2(nums []int, target int) []int {
	ht := map[int]int{}

	for i, v := range nums {
		if p, ok := ht[target-v]; ok {
			return []int{p, i}
		}
		ht[v] = i
	}

	return nil
}

func towSum4(nums []int, target int) []int {
	mp := make(map[int]int, 0)
	for i, num := range nums {
		if p, ok := mp[target-num]; ok {
			return []int{p, i}
		}
		mp[num] = i
	}

	return nil
}

func twoSum3(nums []int, target int) []int {
	mp := map[int]int{}

	for i, cur := range nums {
		if p, ok := mp[target-cur]; ok {
			return []int{p, i}
		}

		mp[cur] = i
	}

	return nil
}

func twoSum4(nums []int, target int) []int {
	mp := make(map[int]int, len(nums))

	for i, num := range nums {
		if value, ok := mp[target-num]; ok {
			return []int{value, i}
		} else {
			mp[num] = i
		}
	}

	return nil
}

func twoSum5(nums []int, target int) []int {
	mp := make(map[int]int, len(nums))
	for i, num := range nums {
		if value, ok := mp[target-num]; ok {
			return []int{i, value}
		} else {
			mp[num] = i
		}
	}

	return nil
}

func twoSum6(nums []int, target int) []int {
	mp := make(map[int]int, len(nums))

	for _, num := range nums {
		if v, ok := mp[target-num]; ok {
			return []int{num, v}
		}

		mp[num] = num
	}

	return nil
}
