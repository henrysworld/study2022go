package demo

import (
	"errors"
	"fmt"
	"github.com/magiconair/properties/assert"
	"reflect"
	"testing"
	"unsafe"
)

func TestName(t *testing.T) {
	// test
	perPage := 3
	rids := [...]int{1, 2, 3, 4, 5, 6, 7}
	var page *[]int
	pageNum := len(rids) / perPage
	if len(rids)%perPage != 0 {
		pageNum = (len(rids) / perPage) + 1
	}
	result := make([][]int, 0, pageNum)
	for i, id := range rids {
		if i == 0 || i%perPage == 0 {
			//page = make([]int, 0, perPage)
			p := make([]int, 0, perPage)
			page = &p
			fmt.Printf("%p\n", p)
			fmt.Printf("%p\n", *page)

			result = append(result, *page)
		}

		*page = append(*page, id)
		fmt.Printf("%p\n", *page)
		fmt.Printf("%p\n", result[0])
		fmt.Println(*page)
		fmt.Println(result[0])
		fmt.Println(result)
	}

	fmt.Println(result)
}

func TestName1(t *testing.T) {
	perPage := 3
	rids := []int{1, 2, 3, 4, 5, 6, 7}
	var result [][]int
	for i := 0; i < len(rids); i += perPage {
		end := i + perPage
		if end > len(rids) {
			end = len(rids)
		}
		result = append(result, rids[i:end])
	}
	fmt.Println(result)

}

func TestName2(t *testing.T) {
	perPage := 3
	rids := []int{1, 2, 3, 4, 5, 6, 7}
	//var result [][]int
	//for i := 0; i < len(rids); i += perPage {
	//	end := i + perPage
	//	if end > len(rids) {
	//		end = len(rids)
	//	}
	//	result = append(result, rids[i:end])
	//}
	result := arrayPartition(rids, perPage)
	fmt.Println(result)

}
func arrayPartition(array []int, limit int) [][]int {
	groupIndex := 0
	var newArray [][]int
	for groupIndex < len(array) {
		end := groupIndex + limit
		if end >= len(array) {
			newArray = append(newArray, array[groupIndex:len(array)])
		} else {
			newArray = append(newArray, array[groupIndex:end])
		}
		groupIndex = end
	}
	return newArray
}

func TestName3(t *testing.T) {
	// 定义一个 2x2 的数组
	arr := [][]int{{1, 2}, {3, 4}}
	fmt.Println("Before:", arr)

	// 将数组的指针传递给 modifyArray 函数
	modifyArray(&arr)

	fmt.Println("After:", arr)
}

func modifyArray(arr *[][]int) {
	for i := 0; i < len(*arr); i++ {
		for j := 0; j < len((*arr)[i]); j++ {
			(*arr)[i][j] = (*arr)[i][j] * 2 // 修改元素的值
		}
	}
}

func TestName5(t *testing.T) {
	var page *[]int
	result := make([][]int, 0, 3)

	p := make([]int, 0, 3)
	page = &p
	result = append(result, p)
	*page = append(*page, 100)

	fmt.Printf("%p\n", &p)
	fmt.Printf("%p\n", p)
	fmt.Printf("%p\n", &result[0])
	fmt.Println(result)
}

func TestArrayList_Add(t *testing.T) {

	perPage := 3
	rids := [...]int{1, 2, 3, 4, 5, 6, 7}
	var page []int
	pageNum := len(rids) / perPage
	if len(rids)%perPage != 0 {
		pageNum = (len(rids) / perPage) + 1
	}
	result := make([][]int, 0, pageNum)
	for i, id := range rids {
		if i == 0 || i%perPage == 0 {
			page = make([]int, 0, perPage)

			result = append(result, page)
		}

		page = append(page, id)
		sh := (*reflect.SliceHeader)(unsafe.Pointer(&result[len(result)-1]))
		sh.Len = len(page)
		//sh :=（*reflect.SliceHeader)(unsafe.Pointer(&page))
		//fmt.Print(&sh)
		//fmt.Printf("%p\n", page)
		//fmt.Println(*page)
		//fmt.Println(result)
	}

	fmt.Println(result)

	//j := -1
	//for i, id := range rids {
	//	if i == 0 || i%perPage == 0 {
	//		page = make([]int, 0, perPage)
	//		result = append(result, page)
	//		j++
	//	}
	//
	//	result[j] = append(result[j], id)
	//}

	fmt.Println(result)

	//a := make([]int, 6)
	//a = append(a, 1)
	//fmt.Println(a)
	testCases := []struct {
		name      string
		list      *ArrayList[int]
		index     int
		newVal    int
		wantSlice []int
		wantErr   error
	}{
		{
			name:      "Add num to index left",
			list:      NewArrayListOf([]int{1, 2, 3}),
			index:     0,
			newVal:    100,
			wantSlice: []int{100, 1, 2, 3},
		},
		{
			name:      "Add num to index right",
			list:      NewArrayListOf([]int{1, 2, 3}),
			index:     3,
			newVal:    100,
			wantSlice: []int{1, 2, 3, 100},
		},
		{
			name:      "Add num to index middle",
			list:      NewArrayListOf([]int{1, 2, 3}),
			index:     1,
			newVal:    100,
			wantSlice: []int{1, 100, 2, 3},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.list.Add(tc.index, tc.newVal)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}

			assert.Equal(t, tc.list.vals, tc.wantSlice)
		})
	}
}

func TestArrayList_Cap(t *testing.T) {
	testCases := []struct {
		name    string
		wantRes int
		list    *ArrayList[int]
	}{
		{
			name:    "Cap empty list",
			wantRes: 5,
			list:    NewArrayListOf([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cap := tc.list.Cap()

			assert.Equal(t, tc.wantRes, cap)
		})
	}
}

func TestArrayList_Get(t *testing.T) {
	var a int
	fmt.Println(a)
	testCases := []struct {
		name    string
		list    *ArrayList[int]
		index   int
		wantRes int
		wantErr error
	}{
		{
			name:    "Get index error",
			list:    NewArrayListOf([]int{1, 2, 3, 4}),
			index:   -1,
			wantRes: 0,
			wantErr: errors.New("index out of bounds"),
		},
		{
			name:    "Get index",
			list:    NewArrayListOf([]int{1, 2, 3, 4}),
			index:   0,
			wantRes: 1,
			wantErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := tc.list.Get(tc.index)

			assert.Equal(t, tc.wantRes, res)
			assert.Equal(t, tc.wantErr, err)
		})
	}
}
