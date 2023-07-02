package tree

import (
	"fmt"
	"sort"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}

	v := groupAnagrams1(strs)
	fmt.Println(v)
	//var a *interface{} = var b *int
	//var a *interface{}
	//var b *int
	//a = b
	// fmt.Println('c' - 'a')
}

func groupAnagrams(strs []string) [][]string {
	hashTable := map[string][]string{}
	for _, cur := range strs {
		s := []byte(cur)
		sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
		key := string(s)
		hashTable[key] = append(hashTable[key], cur)
	}

	ans := make([][]string, 0, len(hashTable))
	for _, v := range hashTable {
		ans = append(ans, v)
	}

	return ans
}

func isAnagram(s string, t string) bool {
	sb := []byte(s)
	st := []byte(t)
	sort.Slice(sb, func(i, j int) bool { return sb[i] < sb[j] })
	sort.Slice(st, func(i, j int) bool { return st[i] < st[j] })

	return string(sb) == string(st)
	// return bytes.Equal(sb, st)
}

func groupAnagrams1(strs []string) [][]string {
	hm := map[string][]string{}

	for _, s := range strs {
		sb := []byte(s)
		sort.Slice(sb, func(i, j int) bool { return sb[i] < sb[j] })
		key := string(sb)
		hm[key] = append(hm[key], s)
	}

	ans := make([][]string, 0, len(hm))
	for _, v := range hm {
		ans = append(ans, v)
	}

	return ans

}
