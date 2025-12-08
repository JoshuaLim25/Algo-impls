package main

import (
	"fmt"
	"os"
)


func nextGreaterElement(nums1, nums2 []int) []int {
	m := make(map[int]int, len(nums1))
	// ans := make([]int, 0, len(nums1))
	ans := make([]int, len(nums1))
	s := []int {}

	for i, v := range nums1 {
		m[v] = i  // values: idx
	}

	for i := range ans {
		ans[i] = -1
	}

	for _, num := range nums2 {
		for len(s) > 0 && num > s[len(s)-1] {  // > means nextGreaterElement
			// pop
			top := s[len(s)-1]
			s = s[:len(s)-1]
			// get idx
			insertIdx := m[top]
			ans[insertIdx] = num
		}

		if _, ok := m[num]; ok {
			s = append(s, num)
		}
	}

	return ans

}

func nextGreaterElt(nums1, nums2 []int) []int {
	m := make(map[int]int, len(nums1))
	res := make([]int, len(nums1))
	s := []int{}

	for i, num := range nums1 {
		m[num] = i // nums: idx, because indices are what lookups are for
	}

	for i := range res {
		res[i] = -1
	}
	fmt.Fprintf(os.Stderr, "DEBUG[1]: next_greater.go:51: res=%+v\n", res)

	for _, num := range nums2 {
		for len(s) > 0 && num > s[len(s)-1] {  // num == nextGreaterElt
			// pop
			popped := s[len(s)-1]
			s = s[:len(s)-1]
			// know from below that top is guaranteed in nums1
			insertIdx := m[popped]
			res[insertIdx] = num
		}
		fmt.Printf("stack: %+v\n", s)
		if _, ok := m[num]; ok {  // if num in nums1
			s = append(s, num)
		}
	}
	return res
}


func main() {
	// nums1 := []int{4, 1, 2}
	// nums2 := []int{1, 3, 4, 2}
	nums1 := []int{2,4}
	nums2 := []int{1,2,3,4}
	// ans := nextGreaterElement(nums1, nums2)
	ans := nextGreaterElt(nums1, nums2)
	fmt.Fprintf(os.Stderr, "DEBUG[1]: next_greater.go:27: ans=%+v\n", ans)
}
